+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.json"
+++

The **kinase** learning algorithm is an abstraction of the detailed chemical pathways involved in [[synaptic plasticity]], which are mediated by _kinases_ including _CaMKII_ and _DAPK1_ that play central roles in these learning processes. This algorithm accomplishes [[error-driven learning]] via a [[temporal derivative]] computed between the faster CaMKII-mediated pathway and the slower DAPK1-mediated pathway, which are known to be in competition with each other ([[@GoodellZaegelCoultrapEtAl17]]; [[@GoodellTullisBayer21]]; [[@CookBuonaratiCoultrapEtAl21]]; [[@TullisBayer23]]; [[@BayerGiese25]]).

The relevant background for this algorithm is presented in the following pages: 

* [[Synaptic plasticity]] reviews the relevant neuroscience for how synapses change their effective strength (weight), and the critical contributions of kinases to this process.

* [[Temporal derivative]] provides a high-level account for the essential computational principles behind this algorithm, including an interactive simulation of how a competitive interaction between fast and slow pathways can compute the _error gradient_ at the heart of error-driven learning.

* [[GeneRec]] derives a concrete learning algorithm directly from the mathematics of [[error backpropagation]], which uses [[bidirectional connectivity]] to propagate error gradients throughout the [[neocortex]]. The kinase algorithm is closely related to GeneRec.

* [[Jiang et al 2025]] presents initial direct evidence showing that the direction of synaptic plasticity in neurons recorded in the mouse CA1 area is consistent with the temporal derivative hypothesis.

Here, we build on these foundations to describe the detailed mechanisms that actually drive learning in the Axon models, which represent an attempt to satisfy constraints from neuroscience, computational efficacy, and computational cost.

The starting point is to compute Ca++ via the biophysically accurate implementations of the [[neuron channels#NMDA]] and [[neuron channels#VGCC]] channels that are well-established as the primary initial drivers of synaptic plasticity. NMDA is sensitive to the conjunction of pre and postsynaptic activity, while VGCCs (voltage-gated calcium channels) are driven in a sharply phasic manner by backpropagating action potentials from the receiving neuron.

However, it would be very expensive computationally to compute this Ca++ value for each synapse individually, so we instead break out the computation into two subcomponents:

* A shared dendrite-level Ca++ value that reflects the overall dendritic membrane potential and the contributions of backpropagating action potentials from the receiving neuron on the NMDA and VGCC channels.

* An efficiently-computed synapse-specific multiplier, that reflects the specific coincident presynaptic and postsynaptic activity at each synapse.

In another example of the [[synergies]] between neuroscience and computation, these two terms can be directly associated with the two essential terms in the error backpropagation learning rule, which are the **error gradient** and the **credit assignment** factors:

{id="eq_err-cred" title="Error * Credit"}
$$
\Delta w \propto \rm{Error} * \rm{Credit}
$$

These are concretely expressed in terms of $\delta$ and the sending unit activity $x$ in error backpropagation:

{id="eq_bp" title="Backpropagation"}
$$
\Delta w \propto \delta x 
$$

In the kinase algorithm, the _Error_ factor is computed from the dendrite-level Ca++, and the synapse-specific multiplier provides the _Credit_ assignment. The computational-level properties of these two factors are overall consistent with the backpropagation versions, but also have important differences that are beneficial for the discrete spiking nature of the Axon framework, and also help reduce the _vanishing gradient_ problem in deep networks.

Due to the continuous-time nature of the Axon model, which simulates neural dynamics at the 1 ms timescale, it takes roughly 200 ms for each "trial" of processing to unfold, corresponding to a [[theta cycle]]. In the [[predictive learning]] framework typically used, this time window encompasses an iteration of prediction (minus phase) followed by an outcome (plus phase), and then learning occurs at the end of the plus phase. The specific events that could trigger this learning in a biologically-realistic manner are described below. Aside from this temporal discretization of the learning process (and external dynamics of the environment), all of the other equations in Axon operate continuously over time.

## Error gradient via dendritic Ca++

The overall Ca++ influx across all the dendrites of a simulated neuron is computed using the biophysically-based [[neuron channels#NMDA]] conductance model described in [[neuron channels]], and a simple spike-driven approximation to the VGCC conductance. The NMDA contribution starts with the conductance:

{id="eq_nmda_g" title="NMDA voltage-gated conductance"}
$$
g_{nmda}(t) = \frac{g_{e-raw}(t) \overline{g}_{nmda}}{1 + \frac{[Mg^{++}]}{3.57} e^{-0.062 V_d}} - \frac{1}{\tau_d} g_{nmda}(t-1)
$$

where $V_d$ is the dendritic membrane potential (see [[neuron dendrites]]) and the $g_{e-raw}(t)$ factor is the raw excitatory synaptic input described in [[neuron#Computing input conductances]], reflecting the total potential glutamate released from presynaptic neurons that could bind to and open NMDA receptors, which is mathematically equivalent to the _net input_ in [[abstract neural network]] models:

{id="eq_ge-raw" title="Excitatory raw glutamate"}
$$
g_{e-raw}(t) = \frac{1}{n} \sum_i x_i w_i
$$

The NMDA conductance drives a Ca++ influx level using the following equation from the [[@^UrakuboHondaFroemkeEtAl08]] model (see [model details](http://kurodalab.bs.s.u-tokyo.ac.jp/info/STDP/) and [[@SabatiniOertnerSvoboda02]]):

{id="eq_ca-nmda" title="NMDA Calcium"}
$$
\rm{Ca}_{nmda}(t) = \frac{- g_{nmda}(t) V_d}{1 - e^{0.0756 V_d}}
$$

The VGCC Ca++ influx occurs in a highly circumscribed time window around the spiking of the receiving neuron, and can be efficiently computed using a simple spike-driven impulse function with a decay factor ($\tau_d = 10 ms$):

{id="eq_ca-vgcc" title="VGCC Calcium"}
$$
\rm{Ca}_{vgcc}(t) = \rm{Spike} * 35 - \frac{1}{\tau_d} \rm{Ca}_{vgcc}(t-1)
$$

The total resulting dendritic Ca++ influx is just the sum of these two terms, with a normalization factor (80) that results in approximately normalized values that work well in practice in the learning rule:

{id="eq_ca-tot" title="Total receiver-based calcium"}
$$
\rm{Ca}_{tot}(t) = \frac{1}{80} \left( \rm{Ca}_{nmda}(t) + \rm{Ca}_{vgcc}(t) \right)
$$

### Kinase cascade

<!--- todo: rename sim -> simulation so it all reads better directly -->

As described in more detail in [[synaptic plasticity]] the influx of Ca++ ions via NMDA and VGCC channels then drives a complex cascade of chemical reactions involving various kinases and phosphatases, along with other molecules and critical binding dynamics, to ultimately drive changes in excitatory AMPA receptor number and efficacy, which is the end result of learning (see the [[Urakubo08 sim]] simulation for a detailed model of some of these processes).

{id="table_taus" title="Kinase time constants"}
| Parameter            | Value  |
|----------------------|--------|
| CaM $\tau_{cam}$     | 2 ms   |
| CaP $\tau_{cap}$     | 40 ms  |
| CaD $\tau_{cad}$     | 40 ms  |
| Syn $\tau_{syn}$     | 30 ms  |

Because it is computationally intractable to simulate these processes in biophysical detail at scale, we use a simple approximation involving a cascade of [[exponential integration]] steps (time constants shown in [[#table_taus]]), starting with the activation of _calcium calmodulin (CaM)_ by the raw Ca++ influx:

{id="eq_cam" title="CaM calmodulin"}
$$
\rm{CaM}(t) = \frac{1}{\tau_{cam}} \left( \rm{Ca}_{tot}(t) - \rm{CaM}(t) \right)
$$

This then drives the _potentiation_ factor, which reflects the extent to which the CaMKII kinase is bound to the GluN2B binding site on the NMDA receptor that then drives LTP (synaptic weight increases), according to the structural model ([[@BayerGiese25]]):

{id="eq_cap" title="CaP potentiation"}
$$
\rm{CaP}(t) = \frac{1}{\tau_{cap}} \left( \rm{CaM}(t) - \rm{CaP}(t) \right)
$$

The final, slowest step in the cascade is the _depression_ factor that reflects the activation of the DAPK1 kinase and binding at the Thr305/306 site on the CaMKII kinase, which drive LTD (synaptic weight decreases):

{id="eq_cad" title="CaD depression"}
$$
\rm{CaD}(t) = \frac{1}{\tau_{cad}} \left( \rm{CaP}(t) - \rm{CaD}(t) \right)
$$

In the [[temporal derivative]] framework, CaP is the faster, potentiation process while CaD is the slower, depression process, and the subtraction of the two provides an approximation to the backpropagation error gradient:

{id="eq_kinase-delta" title="Kinase error gradient"}
$$
\delta \approx \rm{CaP} - \rm{CaD}
$$

The kinase learning rule directly uses this delta factor, along with the _Syn_ synaptic activity credit assignment factor described below:

{id="eq_kinase-dw" title="Kinase learning rule"}
$$
\Delta w \propto (\rm{CaP} - \rm{CaD}) \rm{Syn}
$$

As noted above, this learning rule is applied after 200 ms of iterative updating of all of the above variables, which are plotted in [[#figure_ca++-integration]] across two 200 ms trials to provide a sense of their overall dynamics.

Note that because we are only using the difference between these factors, they are relatively robust to various missing components that would provide a constant offset contribution, such as a more detailed accounting of the Ca++ buffering dynamics. Also, while the direct cascading of CaD on top of CaP captures some of the temporal dynamics of CaMKII activation at Thr286 (which promotes potentiation) versus Thr305/306 (which promotes depression; [[@CookBuonaratiCoultrapEtAl21]]), there are also independent pathways modulating DAPK1 activity, but they are more indirect than the direct binding of CaM to activate CaMKII, and thus likely to be slower overall, which is the critical functional property.

Historically, the above cascaded exponential integration equations were originally developed for the [[Leabra]] XCAL learning rule (_temporally-eXtended Continuous Attractor Learning_; [[@OReillyMunakataFrankEtAl12]]), with some inspiration from biophysical cascades, but without as clear of an understanding of their direct mapping onto the CaMKII and DAPK1 kinase substrates.

### Computational properties of the kinase error gradient

The kinase error gradient ([[#eq_kinase-delta]]) has some important similarities and differences from the corresponding error backpropagation gradient, which can best be seen by comparison with the [[GeneRec]] approximation of this gradient:

{id="eq_generec-delta" title="GeneRec error gradient"}
$$
\delta \approx y^+ - y^-
$$

This simple difference in two [[rate-code activation]] states is used to approximate the corresponding difference in linear net input factors times the derivative of the activation function:

{id="eq_generec-netin" title="GeneRec gradient  netinput"}
$$
\delta = \left( \sum_i x_i^+ w_i - \sum_i x_i^- w_i \right) y' \approx y^+ - y^-
$$

This is overall similar to the kinase error gradient, where the Ca++ influx is driven largely by the net input-like $g_{e-raw}(t)$ term in [[#eq_ge-raw]], but it is also modulated by both the dendritic membrane potential $V_d$ (in both [[#eq_nmda_g]] and [[#eq_ca-nmda]]) and has a direct spiking-based contribution from VGCC ([[#eq_ca-vgcc]]). The subtraction of CaP $-$ CaD in [[#eq_kinase-delta]] implicitly computes the derivative of these _activation-based_ terms ($V_d$ and spiking), as explained in the [[GeneRec]] derivation.

A key difference however is that these kinase activation-based values are overall much more graded than the traditional sigmoidal logistic activation function used in simple GeneRec models, and the related _X-over-X+1_ function used in the [[Leabra]] framework. Therefore, the end result is something closer to the behavior of error backpropagation using the linear ReLU activation function, where the activation derivative $y'$ is just 1 for any active unit, such that the effective error gradient is effectively _linear in the net inputs_.

With a softer, more graded contribution of the activation factors in the kinase error gradient, they are also more linear in the underlying net input factors. The major consequence is that this error gradient should therefore more linearly reflect the error gradients across layers in deep networks with many such layers, thereby avoiding the _vanishing gradient_ problem that plagued earlier backprop networks that used sigmoidal-like functions with saturating nonlinearity (see [[error backpropagation]] for details).

The fact that these activation factors nevertheless do retain a graded representation of the receiving neuron's activity represents an important deviation from error backpropagation with ReLU activations, which only have a strictly binary contribution of receiving activity. If the receiving unit's net input is <= 0, then it learns nothing at all about the current trial, and, somewhat counter-intuitively, if it is > 0, then there is actually no contribution of the receiving unit's activity to the the magnitude of learning (all $y'$ are 1 in [[#eq_generec-netin]]). This was not true when saturating nonlinearities were used. This behavior is inconsistent with the basic mechanisms of [[synaptic plasticity]]. Note that it is still true that the magnitude of the weights going up to the next layer influence the sign and magnitude of the error gradients, but these do not contribute to the activity of the hidden unit in a feedforward backpropagation network.

{id="figure_ca++-integration" style="height:25em"}
![Example traces of Ca++ values across two different trials of 200 cycles (ms) each, showing the graded nature of the activation-based terms, which results in less of a vanishing gradient in kinase-based learning. The constant noisy background of presynaptic excitatory inputs (GeRaw; which is being offset by correspondingly high levels of pooled inhibition, not shown) results in a low-level elevation of the dendritic membrane potential (Vd) and thus a low level of total Ca++ (Ca). When the receiving neuron fires spikes, the NMDA and VGCC contributions to Ca increase dramatically, causing the cascaded integration of CaP and CaD to increase over time. However, in both of these trials, the final levels of CaP and CaD are roughly equal, so the error gradient is near zero and synaptic weights will not change.](media/fig_ca++_integration.png)

The graded nature of these integrated Ca++ signals is shown in [[#figure_ca++-integration]], which plots the relevant variables over two different trials of 200 ms each.

## Synaptic activity credit assignment

The second component of the kinase algorithm is the credit assignment factor, which captures the extent of coincident pre-post synaptic activity over time, which the NMDA receptor is strongly sensitive to. If we did not have real-world computational cost constraints, the aggregate NMDA Ca++ value computed in [[#eq_ca-nmda]] would instead be computed at each individual synapse, using local values, and this sensitivity would be directly manifest in the local synaptic Ca++ signal.

Instead, we use the following separate _Syn_ synapse-specific factor to multiply the aggregate error-gradient term as shown in [[#eq_kinase-dw]]. This activity factor starts with a spike-triggered decaying activity trace that is computed for each neuron, which reflects the net decay / buffering of intracellular Ca++ over time after a spike at the synaptic level, and is labeled `CaSyn` in the simulator:

{id="eq_casyn" title="Synaptic Ca++ spike trace"}
$$
\rm{CaSyn}(t) = \frac{1}{\tau_{syn}} (\rm{Spike}(t) - \rm{CaSyn}(t-1))
$$

At each synapse, the sending * receiving neuron product of these traces is accumulated, using the same cascade of exponential integration factors ([[#eq_cam]] -- [[#eq_cad]]) driven by this product:

{id="eq_sr" title="Synaptic product"}
$$
\rm{SR}(t) = \rm{CaSyn}_s(t) * \rm{CaSyn}_r(t)
$$

The longest time-integral of this SR value from the kinase cascade is then used as the _Syn_ credit assignment value in [[#eq_kinase-dw]] (in theory CaP would be multiplied by the CaP-timescale integral, and CaD would be multiplied by the CaD timescale, but we just have one value and use the longer timescale as it defines the full span of the relevant values).

### Neuron-level linear approximation version

The synapse-specific _SR(t)_ value in [[#eq_trace-syn]] and its exponential integrations, though very simple, is still very computationally-expensive to compute because the number of synapses is much greater than the number of neurons (e.g., $N^2$ in a fully-connected network). Therefore, in practice we compute these values using a set of linear regression coefficients based on a vector of binned _CaSyn_ values at the neuron level, every 10 ms (each bin has the average CaSyn value over that 10 ms).

At the point of learning, which occurs once at the end of a [[theta cycle]] of 200 ms typically, these binned values are multiplied for the sender * receiver at each synapse, and then the linear regression coefficients over these bins are applied to directly compute the effective CaP and CaD time integral value that would otherwise be computed according to the above equations. These coefficients were trained using a combination of ridge and lasso regression based on a full combinatorial space of 100 random trials of Poisson spike trains with minus and plus phase firing rates sampled in 10 hz increments from 0 -- 120 hz (12 * 12 * 12 * 12 = 20,736 frequency combinations * 100 trials = 2,073,600 total trials). The resulting coefficients have $r^2$ values of 0.991 and 0.996 for CaP and CaD respectively (i.e,. they account for that proportion of the variance in the data). Thus, this provides a highly accurate and significantly more performant way of computing these values at the synaptic level. See [kinase/linear](https://github.com/emer/axon/tree/kinase/linear) for details. 

## How does the synapse know when to learn?

TODO:

## Credit Assignment: Temporal Eligibility Trace

The extra mathematical steps taken in (O'Reilly, 1996) to get from backpropagation to the CHL algorithm end up eliminating the factorization of the learning rule into clear Error vs. Credit terms. While this produces a nice simple equation, it makes it essentially impossible to apply the results of Bellec et al., (2020), who showed that backprop-through-time can be approximated through the use of a credit-assignment term that serves as a kind of temporal \emph{eligibility trace}, integrating sender-times-receiver activity over a window of prior time steps.

By adopting the above form of error gradient term, we can now also adopt this trace-based credit assignment term as:

<!--- $$ -->
<!--- Credit = < x y >_t y' -->
<!--- $$ -->

where the angle-bracket expression indicates an exponential running-average time-integration of the sender and receiver activation product over time.


* Supporting a temporally-extended _eligibility trace_ factor that provides a biologically-plausible way of approximating the computationally-powerful backprop-through-time (BPTT) algorithm [[@BellecScherrSubramoneyEtAl20]].

The most computationally-effective form of learning goes one step further in computing this credit-assignment trace factor, by integrating spike-driven activity traces (representing calcium currents) within a given theta cycle of activity, in addition to integrating across multiple such theta-cycle "trials" of activity for the eligibility trace. This is significantly more computationally-expensive, as it requires synapse-level integration of the calcium currents on the millisecond-level timescale (a highly optimized form of computation is used so updates occur only at the time of spiking, resulting in a roughly 2x increase in computational cost overall).

The final term in the credit assignment factor is the derivative of the receiving activation function, \$y'\$, which would be rather difficult to compute exactly for the actual AdEx spiking dynamics used in Axon. Fortunately, using the derivative of a sigmoid-shaped logistic function works very well in practice, and captures the essential functional logic for this derivative: learning should be maximized when the receiving neuron is in its most sensitive part of its activation function (i.e., when the derivative is the highest), and minimized when it is basically "pinned" against either the upper or lower extremes. Specifically the derivative of the logistic is:

$$
y' = y (1-y)
$$

which is maximal at y = .5 and zero at either 0 or 1. In Axon, this is computed using a time-integrated spike-driven Ca-like term (\texttt{CaSpkD}), with the max value across the layer used instead of the fixed 1 constant. In addition, it is useful to use an additional factor that reflects the normalized difference in receiving spiking across the minus and plus phase, which can be thought of as an empirical measure of the sensitivity of the receiving neuron to changes over time: 

$$
y' = y (1-y) \frac{y^+ - y^-}{MAX(y^+, y^-)}
$$

## Stabilization and rescaling mechanisms

A collection of biologically-motivated mechanisms are used to provide a stronger "backbone" or "spine" for the otherwise somewhat "squishy" learning that emerges from the above error-driven learning mechanisms, serving to stabilize learning over longer time scales, and prevent parasitic positive feedback loops that otherwise plague these bidirectionally-connected networks. These positive feedback loops emerge because the networks tend to settle into stable attractor states due to the bidirectional, generally symmetric connectivity, and there is a tendency for a few such states to get broader and broader, capturing more and more of the "representational space". The credit assignment process, which is based on activation, contributes to this "rich get richer" dynamic where the most active neurons experience the greatest weight changes. We colloquially refer to this as the "hog unit" problem, where a small number of units start to hog the representational space, and it represents a major practical barrier to effective learning if not managed properly. Note that this problem does not arise in the vast majority of purely feedforward networks used in the broader neural network field, which do not exhibit attractor dynamics. However, this kind of phenomenon is problematic in other frameworks with the potential for such positive feedback loops, such as on-policy reinforcement learning or generative adversarial networks.

Metaphorically, various forms of equalizing taxation and wealth redistribution are required to level the playing field. The set of stabilizing, anti-hog mechanisms in Axon include:

* **SWt:** structural, slowly-adapting weights. In addition to the usual learning weights driven by the above equations, we introduce a much more slowly-adapting, multiplicative \texttt{SWt} that represents the biological properties of the dendritic \emph{spine} -\/- these SWts "literally" give the model a spine! Spines are structural complexes where all the synaptic machinery is organized, and they slowly grow and shrink via genetically-controlled, activity-dependent protein remodeling processes, primarily involving the \emph{actin} fibers also found in muscles. A significant amount of spine remodeling takes place during sleep -\/- so the SWt updating represents a simple model of sleep effects.

The SWt is multiplicative in the sense that larger vs. smaller spines provide more or less room for the AMPA receptors that constitute the adaptive weight value. The net effect is that the more rapid trial-by-trial weight changes are constrained by this more slowly-adapting multiplicative factor, preventing more extreme changes. Furthermore, the SWt values are constrained by a zero-sum dynamic relative to the set of receiving connections into a given neuron, preventing the neuron from increasing all of its weights higher and hogging the space. The SWt is also initialized with all of the randomness associated with the initial weights, and preserving this source of random variation, preventing weights from becoming too self-similar, is another important function.

* **Target activity levels:** There is extensive evidence from Gina Turrigiano and collaborators, among others, that synapses are homeostatically rescaled to maintain target levels of overall activity, which vary across individual neurons (e.g., {Torrado Pacheco et al., 2021}). Axon simulates this process, at the same slower timescale as updating the SWts (likewise associated with sleep), which are also involved in the rescaling process. The target activity levels can also slowly adapt over time, similar to an adaptive bias weight that absorbs the "DC" component of the learning signal {Schraudolph (1998)}, but this adaptation is typically subject to a zero-sum constraint, so any increase in activity in one neuron must be compensated for by reductions elsewhere.

This is similar to a major function performed by the BCM learning algorithm in the Leabra framework -\/- by moving this mechanism into a longer time-scale outer-loop mechanism (consistent with Turigiano's data), it worked much more effectively. By contrast, the BCM learning ended up interfering with the error-driven learning signal, and required relatively quick time-constants to adapt responsively as a neuron's activity started to change.

* **Zero-sum weight changes:** In some cases it can also be useful
to constrain the faster error-driven weight changes to be zero-sum, which is supported by an optional parameter. This zero-sum logic was nicely articulated by {Schraudolph (1998)}, and is implemented in the widely-used ResNet models.

* **Soft bounding and contrast enhancement:** To keep individual weight magnitudes bounded, we use a standard exponential-approach "soft bounding" dynamic (increases are multiplied by \$1-w\$; decreases by \$w\$). In addition, as developed in the Leabra model, it is useful to add a \emph{contrast enhancement} mechanism to counteract the compressive effects of this soft bounding, so that effective weights span the full range of weight values.


