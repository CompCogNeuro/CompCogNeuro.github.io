+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

This page provides details for the full range of channel types that are available in an [[Axon]] [[neuron]] to drive biologically-based and functionally important behavior in specific neuron types. Every neuron uses the basic excitatory, inhibitory, and leak channels discussed in detail in the [[neuron]] and [[neuron electrophysiology]] pages, and some of the following channels are used in most other neurons, while some are only used in specific neuron types where they are particularly important.

These biologically-grounded channels provide accurate fits to the detailed electrophysiological properties of real neurons, based on the sources listed below. Although this results in a large number of parameters relative to the units in [[abstract neural network]]s, we almost never change these parameters from their default values, unless there is a clear biological or functional motivation to do so. Furthermore, extensive testing across a wide range of models has shown that these biologically-grounded mechanisms, and parameter values, actually produce the best functional results.

The implementation of several of these channels comes from standard biophysically detailed models such as [[@^MiglioreHoffmanMageeEtAl99]], [[@^PoiraziBrannonMel03]], and [[@^UrakuboHondaFroemkeEtAl08]]. See also [[@^BretteRudolphCarnevaleEtAl07]] and the [NEST model directory](https://nest-simulator.readthedocs.io/en/stable/models/index.html) for documented examples, including: [AdEx](https://nest-simulator.readthedocs.io/en/stable/models/aeif_cond_exp.html), [Traub HH](https://nest-simulator.readthedocs.io/en/stable/models/hh_cond_exp_traub.html).  The [Brian Examples](https://brian2.readthedocs.io/en/stable/examples/index.html) contain full easy-to-read equations for various standard models, including [Brunel & Wang, 2001](https://brian2.readthedocs.io/en/stable/examples/frompapers.Brunel_Wang_2001.html). Also see [Wikipedia: Biological neuron model](https://en.wikipedia.org/wiki/Biological_neuron_model) for a nice overview, and [ModelDB Currents](https://modeldb.science/NeuronDB/NeuronalCurrents), [ModelDB Current Search](https://modeldb.science/ModelDB/FindByCurrent), and [IonChannelGeneology](https://icg.neurotheory.ox.ac.uk) for standardized lists of currents included in biophysical models made in NEURON and related software.

## Parameters

The default parameters for each of the channel types covered here are shown in the following tables, using the [[neuron#units and parameters|standard units]].

{id="table_taus" title="Time constants"}
| Parameter                        | Value    |
|----------------------------------|----------|
| $\tau_{ampa}$                    | 5 ms     |
| $\tau_{gaba_a}$                  | 7 ms     |
| $\tau_{nmda}$                    | 100 ms   |
| $\tau_{gaba_b}$ rise ($\tau_r$)  | 45 ms    |
| $\tau_{gaba_b}$ decay ($\tau_d$) | 50 ms    |
| mAHP $\tau_{max}$                | 1,000 ms |


{id="table_gs" title="Conductance scaling factors"}
| Parameter                 | Value   |
|---------------------------|---------|
| $\overline{g}_{nmda}      | 0.006   |
| $\overline{g}_{gabab}     | 0.015   |
| $\overline{g}_{vgcc}      | 0.02    |
| $\overline{g}_{ak}        | 0.1     |
| $\overline{g}_{mahp}      | 0.02    |
| $\overline{g}_{sahp}      | 0.05    |


{id="table_gs" title="Typical max conductances"}
| Parameter                 | Value     |
|---------------------------|-----------|
| NMDA posterior cortex     | 50 nS     |
| GABA-B                    | 50 nS     |
| VGCC                      | 4 nS      |
| aK                        | 10 nS     |
| Mahp                      | 2 nS      |
| Sahp                      | 5 nS      |


{id="table_gs-refs" title="Conductances from other models"}
| Parameter                            | Value            |
|--------------------------------------|------------------|
| NMDA PFC (Brunel & Wang, 2001)       | 0.327 nS per syn |
| VGCC distal (Urakubo et al., 2008)   | 146 nS   |
| VGCC proximal (Urakubo et al., 2008) | 3.2 nS   |
| VGCC soma (Urakubo et al., 2008)     | 93 nS    |
| aK distal (Urakubo et al., 2008)     | 49 nS    |
| aK proximal (Urakubo et al., 2008)   | 7.5 nS   |


## AMPA

The AMPA (α-Amino-3-hydroxy-5-methyl-4-isoxazolepropionic acid) channel is the primary excitatory synaptic input channel, discussed extensively in [[neuron]]. The channel is opened by the binding of the neurotransmitter **glutamate**, which causes the channel structure to twist open, allowing primarily $Na^+$ ions to enter the cell, causing the electrical potential to increase (excitation).

The AMPA receptor conductance can be modeled using the _double-exponential_ function, where _t_ is the time since the binding of glutamate to the receptor: 

{id="eq_double_e" title="Double exponential"}
$$
g(t) = e^{-t / \tau_1} - e^{-t / \tau_2}
$$

$\tau_1$ is a fast _rise_ time constant for the increase in conductance when the glutamate first binds to the AMPA receptor (less than 1 ms according to [[@HestrinNicollPerkelEtAl90]]), and $\tau_2$ is a slower _decay_ time constant reflecting the inactivation of the AMPA receptor over time, estimated at 4.4 ms by [[@^HestrinNicollPerkelEtAl90]].

The _alpha_ function as introduced by [[@^Rall67]] has also been used to model relatively fast conductances:

{id="eq_alpha" title="Alpha function"}
$$
g(t) = \frac{t}{\tau} e^{-t / \tau}
$$

In the [[Axon]] model we use a time step of 1 ms for integrating all of the [[neuron]] level equations, so the relatively fast rise time constant happens too quickly to be of relevance. Thus, the AMPA conductance increases discretely in the 1 ms time step, and we use a single exponential decay function with a time constant of 5 ms, which can be computed using an online exponential decay function:

{id="eq_ampa_g" title="AMPA conductance"}
$$
g_{ampa}(t) = g_{ampa}(t-1) \left(1 - \frac{1}{\tau_{ampa}} \right)
$$

As with all channels, this conductance then drives a corresponding current as a function of the reversal potential for AMPA ($E_{ampa}$), which is estimated at 0 mV:

{id="eq_ampa_i" title="AMPA current"}
$$
I_{ampa} = g_{ampa} \left(E_{ampa} - Vm\right)
$$

## GABA-A

The GABA-A channel is the standard inhibitory synaptic input channel discussed in [[neuron]] and [[inhibition]]. It is opened by the binding of the GABA (gamma-aminobutyric acid) neurotransmitter, released by special populations of inhibitory interneurons. It primarily allows negatively-charged chloride ions $Cl^-$ to flow into the cell, which act to keep the electrical potential negative.

We model GABA-A conductances in the same way as AMPA, with a single exponential decay function ([[#eq_ampa_g]], using a time constant $\tau$ of 7 ms [[@XiangHuguenardPrince98]]. The reversal potential for GABA-A ($E_{gaba_a}$) is -75 mV.

## K Leak

The $K^+$ leak channel is always open, so it just has a constant conductance parameterized by $\overline{g}_{gaba_a}$ and a reversal potential ($E_{leak}$) of -75 mV.

## NMDA

{id="plot_nmda" title="NMDA channel" collapsed="true"}
```Goal
pl := &chanplots.NMDAPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "NMDA", "NMDA Channel", "NMDA Channel equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
br.Splits.Styler(func(s *styles.Style) {
	s.Min.Y.Em(25)
})
pl.GVRun()
```

NMDA (N-methyl-D-aspartate) channels are found throughout the brain, and play a critical role in learning as captured in the [[kinase algorithm]] used in [[Axon]]. The opening of NMDA channels is typically blocked by positively-charged magnesium ions ($Mg^{++}$) when the membrane potential is close to the resting potential. The removal of this block as a result of membrane depolarization above this resting potential (known as an _outward rectification_, because the $Mg^{++}$ ions are on the outside of the cell) is one of two key functional features of this channel. This unblocking dynamic is important both for learning and for its important contributions to the activation dynamics of the neuron.

The other critical functional property of the NMDA channel is that it also requires _glutamate_ neurotransmitter binding to open, in addition to $Mg^{++}$ unblocking. Thus, unlike many other channels that just have one (or none) _gating_ factors, NMDA requires both of these factors. Furthermore, the $Mg^{++}$ factor is a function of the _postsynaptic_ activity (depolarization), while glutamate is released by the  _presynaptic_ neuron. Thus, the NMDA is in a unique position to respond to the _conjunction_ of both pre and postsynaptic activity.

To see the unblocking in action, press the [[#plot_nmda:GV run]] button above, which generates a conductance (_g_) (and current _I_) vs. voltage (_v_) plot using the parameters shown to the left of the plot, assuming that there is plenty of glutamate around so that factor is not relevant. As the voltage increases above the -90 hyperpolarized starting point, the conductance steadily rises, reflecting the progressively increased likelihood that the $Mg^{++}$ ions will not be blocking the channel opening. The _reversal potential_ for the NMDA channel is around 0 mV, so as the voltage approaches this point, the net force pulling the ions through the channel gets progressively weaker (as explained by the tug-of-war analogy in [[neuron]]), so the current _I_ decreases as the voltage approaches 0.

NMDA channels mostly allow _calcium_ ions ($Ca^{++}$) to flow into the cell, and the learning effects of this channel are due to the ability of calcium to trigger various postsynaptic chemical reactions as described in [[kinase algorithm]]. The activation effects are due to positive charges on this ion, which therefore has a net excitatory (depolarizing) effect on the cell.

[[#plot_nmda:Time run]] shows the other critical feature of the NMDA channel, which is that the _Tau_ time constant parameter (greek $\tau$) is much longer than most other channels, on the order of 100 ms or more. This parameter describes the exponential decay constant (of the same form as [[#eq_ampa_g]]) for the NMDA conductance, which like AMPA has a sufficiently fast rise time that it can be ignored. This relatively long time constant is critical for the activation contributions of the NMDA channel, because it creates a [[stable activation]] pattern over time (see that page for more discussion and a demonstration).

The equation we use for the voltage-gated conductance is due to [[@JahrStevens90]], and is used in the widely-cited [[@BrunelWang01]] model:

{id="eq_nmda_g" title="NMDA voltage-gated conductance"}
$$
g_{nmda}(V) = \overline{g}_{nmda} \frac{1}{1 + \frac{[Mg^{++}]}{3.57} e^{-0.062 V}}
$$

where $\overline{g}_{nmda}$ is the overall "g-bar" maximum conductance factor, and $[Mg^{++}]$ is the extracellular magnesium concentration, which is typically between 1 and 1.5. This function is of a sigmoidal, "S-shaped" form, increasing to an asymptotic value as the voltage increases; the relevant portion of this function is on the left side of the sigmoid as you can see in [[#plot_nmda:GV run]]. The decrease that we see in _I_ as voltage increases is due to the standard tug-of-war Ohm's law multiplier that applies to all channel conductances ([[#eq_ampa_i]]).

## GABA-B

{id="plot_gabab" title="GABA-B channel" collapsed="true"}
```Goal
pl := &chanplots.GABABPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "GABA-B", "GABA-B Channel", "GABA-B Channel equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
br.Splits.Styler(func(s *styles.Style) {
	s.Min.Y.Em(25)
})
pl.GVRun()
```

The GABA-B channel has a much slower decay time constant compared to the standard GABA-A inhibitory channel, because it is coupled to the GIRK channel: _G-protein coupled inwardly rectifying potassium (K) channel_. It is ubiquitous in the brain, and is likely essential for basic neural function. The _inward rectification_ is caused by a $Mg^{++}$ ion block from the _inside_ of the neuron, which means that these channels are most open when the neuron is hyperpolarized (inactive), and thus it serves to _keep inactive neurons inactive_. This is complementary to the effect of NMDA channels, and [[@^SandersBerendsMajorEtAl13]] emphasized the synergistic nature of these two channels in producing [[stable activation]] patterns: NMDA keeps active neurons active, while GABA-B keeps inactive neurons inactive.

Press the [[#plot_gabab:GV run]] button above to see the conductance (_g_) (and current _I_) vs. voltage (_v_) plot using the parameters shown to the left of the plot. In comparison with the comparable NMDA plot, you can see that GABA-B and NMDA are mirror-images of each other. Furthermore, the _I_ value plotted here is the absolute value (positive) whereas the actual current has the opposite sign as NMDA, due to its reversal potential $E_{gaba_b}$ being that of potassium (-90 mV).

The implementation of the GABA-B channel is based on [[@^SandersBerendsMajorEtAl13]] and [[@^ThomsonDestexhe99]], with the following sigmoidal voltage-gated conductance function from [[@^YamadaInanobeKurachi98]]:

{id="eq_gabab_gv" title="GABA-B voltage-gated conductance"}
$$
g_{gaba_b}(V) = \overline{g}_{gaba_b} \frac{1}{1 + e^{0.1(V-E_{gaba_b}+10)}}
$$

There is an additional sigmoidal function needed for computing the time dynamics of the GABA-B conductance as a function of the GABA binding from inhibitory input spikes over time (in Figure 16 of [[@ThomsonDestexhe99]]):

{id="eq_gabab_x" title="GABA-B spiking rate integration"}
$$
X = \frac{1}{1 + e^{-(s - 7.1) / 1.4}}
$$

Where _s_ is the rate of spiking over the recent time window, which we compute based on the $g_i$ inhibition factor from the [[inhibition]] function used in [[Axon]], and _X_ drives increases in GABA-B activation (labeled _M_) according to the following double-exponential update equations with separate rise ($\tau_r$) and decay ($\tau_d$) factors (45 and 50 ms respectively, which fit the timecourse data from [[@ThomsonDestexhe99]] well; [[#table_taus]]):

{id="eq_gabab_m" title="GABA-B activation M over time"}
$$
M(t) = \frac{1}{\tau_r} \left( \left[(\tau_d / \tau_r)^{(\tau_r / (\tau_d - \tau_r))} \right] X(t) - M(t-1) \right)
$$

$$
X(t) = \left( \frac{1}{1 + e^{-(s - 7.1) / 1.4}} - X(t-1) \right) - \tau_d X(t-1)
$$

The final GABA-B conductance is a product of the M activation factor above in [[#eq_gabab_m]] and the voltage gating factor shown in [[#eq_gabab_gv]]:

{id="eq_gabab_tg" title="GABA-B net conductance over time"}
$$
g_{gaba_b}(t) = g_{gaba_b}(V) M(t)
$$

Do [[#plot_gabab:Time run]] to see these time dynamics play out over a 500 ms window with a pulse of input at the start.

## VGCC: Voltage-gated calcium channels

Voltage gated calcium channels (VGCC) are similar to NMDA channels in that their conductance to $Ca^{++}$ has a voltage dependency, but they do _not_ have a neurotransmitter binding property, and their voltage dependency is typically at higher threshold than NMDA (and is not caused by $Mg^{++}$ block). Due to this higher threshold, the VGCC channels are typically only open during backpropagating action potentials (see [[neuron dendrites]] for details), and thus they provide a calcium signal that is closely tied to postsynaptic spiking. We leverage this property in the [[kinase algorithm]] learning rule. Because VGCCs also close very quickly once the spike is over, they do not have a big impact on activation dynamics --- they are mostly important for learning.

There are a large number of VGCCs types ([[@Dolphin18]]; [[@CainSnutch12]]) denoted by letters in descending order of the voltage threshold for activation: L, PQ, N, R, T, which have corresponding Ca_v names: Ca_v1.1, 1.2, 1.3. 1.4 are all L type, 2.1, 2.2, 2.3 are PQ, N, and R, respectively, and T type (low threshold) comprise 3.1, 3.2, and 3.3. Each channel is characterized by the voltage dependency and inactivation functions. 

{id="table_vgcc" title="VGCC channel types"}
| Letter | Ca_v    | V Threshold  | Inactivation | Location | Function              |
| ------ | ------- | ------------ | ------------ | -------- | --------------------- |
|  L     | 1.1-1.4 | high (-40mV) | fast         | Cortex + | closely tracks spikes |
|  PQ    | 2.1     | high         | ?            | Cerebellum (Purk, Gran) | ?      |
|  N     | 2.2     | high         | ?            | everywhere? | ?                  |
|  R     | 2.3     | med          | ?            | Cerebellum Gran | ?              | 
|  T     | 3.1-.3  | low          | ?            | 5IB, subcortical  | low-freq osc |

* The L type is the classic "VGCC" in dendritic spines in pyramidal cells, which we plot below.

* PQ and R are specific to cerebellum.

* The T type is the most important for low frequency oscillations, and is absent in pyramidal neurons outside of the 5IB layer 5 neurons, which are the primary bursting type. It is most important for subcortical neurons, such as in TRN. See [Destexhe et al, 1998 model in BRIAN](https://brian2.readthedocs.io/en/stable/examples/frompapers.Destexhe_et_al_1998.html) for an implementation.

{id="plot_vgcc" title="VGCC L-type channel" collapsed="true"}
```Goal
pl := &chanplots.VGCCPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "VGCC", "VGCC Channel", "VGCC L-type channel equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
br.Splits.Styler(func(s *styles.Style) {
	s.Min.Y.Em(25)
})
```

Our implementation of the L-type VGCC is based on [[@UrakuboHondaFroemkeEtAl08]], using source code available at this [link](http://kurodalab.bs.s.u-tokyo.ac.jp/info/STDP/Urakubo2008.tar.gz).

First, there is a temporally-invariant aspect of the voltage gating defined by a sigmoidal function similar to those seen above:

{id="eq_vgcc_gv" title="VGCC L voltage-gated conductance"}
$$
g_{vgcc}(V) = -\overline{g}_{vgcc} \frac{1}{1 - e^{0.0756 V}}
$$

And there are two additional opponent gating factors denoted _M_ (activating) and _H_ (inactivating) that have a strong time dependency, and sigmoidal driving functions as follows:

{id="eq_vgcc_m" title="VGCC M gate voltage-based max"}
$$
M_{max}(V) = \frac{1}{1 + e^{-(V + 37)}}
$$

{id="eq_vgcc_h" title="VGCC H gate voltage-based max"}
$$
H_{max}(V) = \frac{1}{1 + e^{2(V + 41)}}
$$

The update equations just move toward these max values with associated time constants:

{id="eq_vgcc_dm" title="VGCC M gate update"}
$$
M(t) = M(t-1) + \frac{1}{3.6} \left( M_{max}(V) - M(t-1) \right)
$$

{id="eq_vgcc_dh" title="VGCC H gate update"}
$$
H(t) = H(t-1) + \frac{1}{29} \left( H_{max}(V) - H(t-1) \right)
$$

The final conductance over time reflects the activation vs. inactivating binding sites in a 3 to 1 ratio:

{id="eq_vgcc_gt" title="VGCC L conductance over time"}
$$
g_{vgcc}(t) = g_{vgcc}(V) M^3(t) H(t)
$$

To see the static voltage-gated sigmoidal functions, do [[#plot_vgcc:GV run]]. To see the response of the M and H channels to discrete spiking inputs, do [[#plot_vgcc:Time run]]. In both cases you will need to deselect variables to be able to see the values with smaller ranges. You should observe that the M activating channel rises up quickly at every action potential, and drops quickly back down, consistent with its 3.6 ms time constant. By contrast, the H inactivating factor builds up over time and slowly decreases the overall conductance value.

## A-type K

{id="plot_ak" title="A-type K channel" collapsed="true"}
```Goal
pl := &chanplots.AKPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "AK", "A-type K Channel", "A-type K channel equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
br.Splits.Styler(func(s *styles.Style) {
	s.Min.Y.Em(25)
})
```

The A-type K channel is voltage-gated with maximal activation around -37 mV ([[@HoffmanMageeColbertEtAl97]]). It is particularly important for counteracting the excitatory effects of the VGCC L-type channels (with which they are co-localized) which can otherwise drive runaway excitatory currents. Think of it as an "emergency brake" and is needed for this reason whenever adding VGCC to a model.

It has two state variables, M (V-gated opening) and H (V-gated closing), which integrate with fast and slow time constants, respectively. H relatively quickly hits an asymptotic level of inactivation for sustained activity patterns, so we can actually ignore it with minimal consequences, and also simplify some of the faster time dynamics because we are not implementing the fast Hodgkin-Huxley spiking channels. Thus, in our simulations we just use a single sigmoidal function for the M activating component, with parameters that fit the rising portion of the more complex function:

{id="eq_ak_m" title="A-type M activation"}
$$
M = \frac{0.076}{1 + e^{0.075(V + 2)}}
$$

This function is missing the declining values beyond the -37 mV peak, but given our spiking dynamics and 1 ms time step of integration, this is not relevant. The comparison between this simplified conductance and the full model from [[@^HoffmanMageeColbertEtAl97]] and [[@^MiglioreHoffmanMageeEtAl99]] can be seen in the [[#plot_ak:GV run]] plot. You can also see the time dynamics in the [[#plot_ak:Time run]] plot.

## Adaptation channels

There are a number of different channels that drive an [[adaptation]] effect in neurons (also known as _accommodation_ or _neural fatigue_), over different time scales and in response to different activity signals. These channels cause neurons to decrease their responsiveness over time, as a function of how active they have been, which results in an overall suppression of responses to ongoing activity patterns, and a relative enhancement to novel ones (i.e., a _novelty filter_). All of these channels conduct $K^+$ ions, like the leak channel, and we add their conductances together in an overall `Gk(t)` value.

Functionally, these are also known as _afterhyperpolarization (AHP)_ channels, because they cause neurons to become refractory (less responsive) to further excitatory inputs for different time windows, from fast (fAHP; 2-5 ms), to medium (mAHP; 50-100 ms), and slow (sAHP; 0.1-2 s).

### M-type mAHP

{id="plot_mahp" title="mAHP M-type channel" collapsed="true"}
```Goal
pl := &chanplots.MahpPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "mAHP", "mAHP Channel", "mAHP M-type equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
br.Splits.Styler(func(s *styles.Style) {
	s.Min.Y.Em(25)
})
```

There are a large number of different K channels that were historically called _M-type_ due to their muscarinic acetylcholine (ACh) response in the bullfrog sympathetic ganglion cells. These are now classified as Kv7 KCNQ channels, which have been identified throughout the brain, with responsiveness to a wide range of different neurotransmitters and other factors ([[@GreeneHoshi17]]). One such M-type channel is a major contributor to the mAHP K current, which we describe here. It is voltage sensitive, but starts to open at low voltages (-60 mV), and can be closed by different neurotransmitters or other factors. In general it takes a while to activate, with a time constant of around 50 msec or so, and it also deactivates on that same timescale.

Relative to the kNA channels described below, which respond to $Na^+$ influx from spikes, the broadly-tuned voltage sensitivity of the M-type mAHP channel produces a stronger _anticipatory_ conductance prior to the spike. Thus, it will "head off" incipient spikes in a way that the KNa channels do not.

The original characterization of the M-type current in most models derives from [[@^GutfreundYaromSegev95]], as implemented in NEURON by [[@^MainenSejnowski96]], see these ModelDB entries: [2488](https://modeldb.science/2488?tab=2&file=cells/km.mod), and [181967](https://modeldb.science/181967?tab=2&file=CutsuridisPoirazi2015/km.mod) from [[@CutsuridisPoirazi15]], and [ICGeneology](https://icg.neurotheory.ox.ac.uk/viewer/?family=1&channel=1706) for the widespread use of this code.

There is a voltage gating factor _N_ (often labeled _M_ for other channels) which has an asymptotic drive value ($N_{infty}$) and a time constant $\tau$ which are both composed from two sigmoidal functions of potential V, centered at -30 mV with a slope of 9 mV:

{id="eq_mahp_ab" title="mAHP functions"}
$$
V_o = V + 30
$$

$$
A = \frac{V_o}{\tau_{max} \left(1 - e^{-V_o/9} \right)}
$$

$$
B = \frac{-V_o}{\tau_{max} \left(1 - e^{V_o/9} \right)}
$$

$$
N_{\infty} = \frac{A}{A + B}
$$

$$
\tau = \frac{1}{A + B}
$$

{id="eq_mahp_dn" title="mAHP N update"}
$$
N(t) = N(t-1) + \frac{1}{\tau} \left( N_{\infty} - N(t-1) \right)
$$

$$
g_{mahp} = \overline{g}_{mahp} 2.3^{(37-23)/10} N(t)
$$

You can see these functions in [[#plot_mahp:GV run]], and the time-course dynamics in [[#plot_mahp:Time run]].

### sAHP: slow afterhyperpolarization

{id="plot_sahp" title="sAHP M-type channel" collapsed="true"}
```Goal
pl := &chanplots.SahpPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "sAHP", "sAHP Channel", "sAHP M-type equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
br.Splits.Styler(func(s *styles.Style) {
	s.Min.Y.Em(25)
})
```

It is difficult to identify the origin of a slow, long-lasting sAHP current, which has been observed in hippocampal and other neurons ([[@Larsson13]]). It appears to be yet another modulator on the M-type channels driven by calcium sensor pathways that have longer time constants. There is more research to be done here, but we can safely use a mechanism that takes a long time to build up before activating the K+ channels, and then takes a long time to decay as well.

The above equations ([[#eq_mahp_ab]]) are used for sAHP, driven by a normalized integrated Ca value, with an offset of 0.8 and slope of 0.02. Unlike mAHP which is updated at the standard 1 ms time step, we update sAHP at the theta cycle interval, which automatically extends the temporal dynamics.

### Sodium-gated potassium channels for adaptation (kNa adapt)

The longer-term adaptation (accommodation / fatigue) dynamics of neural firing in our models are based on sodium (Na) gated potassium (K) currents. As neurons spike, driving an influx of Na, this activates the K channels, which, like leak channels, pull the membrane potential back down toward rest (or even below).  Multiple different time constants have been identified and this implementation supports 3: M-type (fast), Slick (medium), and Slack (slow) ([[@Kaczmarek13]]; [[@Kohn07]]; [[@Sanchez-VivesNowakMcCormick00a]]; [[@BendaMalerLongtin10]]).

The logic is simplest for the spiking case, and can be expressed in conditional program code:
```
	if spike {
		gKNa += Rise * (Max - gKNa)
	} else {
		gKNa -= 1/Tau * gKNa
	}
```

The KNa conductance ($g_{kna}$ in mathematical terminology, `gKNa` in the program) rises to a `Max` value with a `Rise` rate constant, when the neuron spikes, and otherwise it decays back down to zero with another time constant `Tau`.

The equivalent rate-code equation just substitutes the rate-coded activation variable in as a multiplier on the rise term:

```
	gKNa += act * Rise * (Max - gKNa) - (1/Tau * gKNa)
```

{id="table_kna" title="kNA adaptation time constants"}
| Channel Type     | Tau (ms) | Rise  |  Max  |
|------------------|----------|-------|-------|
| Fast (M-type)    | 50       | 0.05  | 0.1   |
| Medium (Slick)   | 200      | 0.02  | 0.1   |
| Slow (Slack)     | 1000     | 0.001 | 1.0   |

The default parameters, which were fit to various empirical firing patterns and also have proven useful in simulations, are shown in [[#table_kna]].


