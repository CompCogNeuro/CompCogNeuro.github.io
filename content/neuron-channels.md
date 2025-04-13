+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

This page provides details for the full range of channel types that are available in an [[Axon]] [[neuron]] to drive specific biologically-based and functionally important behavior in specific neuron types. Every neuron uses the basic excitatory, inhibitory, and leak channels discussed in detail in the [[neuron]] page, and some of the following channels are used in most other neurons, while some are only used in specific neuron types where they are particularly important.

## AMPA

The AMPA (α-Amino-3-hydroxy-5-methyl-4-isoxazolepropionic acid) channel is the primary excitatory synaptic input channel, discussed extensively in [[neuron]]. The channel is opened by the binding of the neurotransmitter **glutamate**, which causes the channel structure to twist open, allowing primarily $Na^+$ ions to enter the cell, causing the electrical potential to increase (excitation).

The AMPA receptor conductance can be modeled using the _double-exponential_ function, where _t_ is the time since the binding of glutamate to the receptor: 

{id="eq_double_e"}
$$
g(t) = e^{-t / \tau_1} - e^{-t / \tau_2}
$$

$\tau_1$ is a fast _rise_ time constant for the increase in conductance when the glutamate first binds to the AMPA receptor (less than 1 ms according to [[@HestrinNicollPerkelEtAl90]]), and $\tau_2$ is a slower _decay_ time constant reflecting the inactivation of the AMPA receptor over time, estimated at 4.4 ms by [[@^HestrinNicollPerkelEtAl90]].

The _alpha_ function as introduced by [[@^Rall67]] has also been used to model relatively fast conductances:

{id="eq_alpha"}
$$
g(t) = \frac{t}{\tau} e^{-t / \tau}
$$

In the [[Axon]] model we use a time step of 1 ms for integrating all of the [[neuron]] level equations, so the relatively fast rise time constant happens too quickly to be of relevance. Thus, the AMPA conductance increases discretely in the 1 ms time step, and we use a single exponential decay function with a time constant of 5 ms, which can be computed using an online exponential decay function:

{id="eq_ampa_g"}
$$
g_{ampa}(t) = g_{ampa}(t-1) \left(1 - \frac{1}{\tau_{ampa}} \right)
$$

As with all channels, this conductance then drives a corresponding current as a function of the reversal potential for AMPA, which is estimated at 0 mV:

{id="eq_ampa_i"}
$$
I_{ampa} = g_{ampa} \left(E_{ampa} - Vm\right)
$$

## GABA-A

## Leak

## NMDA

{id="plot_nmda" title="NMDA Channels" collapsed="true"}
```Goal
pl := &chanplots.NMDAPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "NMDA", "NMDA Channel", "NMDA Channel equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
```

NMDA (N-methyl-D-aspartate) channels are found throughout the brain, and play a critical role in learning as captured in the [[kinase algorithm]] used in [[Axon]]. The opening of NMDA channels is typically blocked by positively-charged magnesium ions ($Mg^{++}$) when the membrane potential is close to the resting potential, and the removal of this block as a result of membrane depolarization above this resting potential is one of two key functional features of this channel. This unblocking dynamic is important both for learning and for its important contributions to the activation dynamics of the neuron.

The other critical functional property of the NMDA channel is that it also requires _glutamate_ neurotransmitter binding to open, in addition to $Mg^{++}$ unblocking. Thus, unlike many other channels that just have one (or none) _gating_ factors, NMDA requires both of these factors. Furthermore, the $Mg^{++}$ factor is a function of the _postsynaptic_ activity (depolarization), while glutamate is released by the  _presynaptic_ neuron. Thus, the NMDA is in a unique position to respond only to the conjunction of both pre and postsynaptic activity.

To see the unblocking in action, press the [[#plot_nmda:GV run]] button above, which generates a conductance (_g_) (and current _I_) vs. voltage (_v_) plot using the parameters shown to the left of the plot, assuming that there is plenty of glutamate around so that factor is not relevant. As the voltage increases above the -90 hyperpolarized starting point, the conductance steadily rises, reflecting the progressively increased likelihood that the $Mg^{++}$ ions will not be blocking the channel opening. The _reversal potential_ for the NMDA channel is around 0 mV, so as the voltage approaches this point, the net force pulling the ions through the channel gets progressively weaker (as explained by the tug-of-war analogy in [[neuron]]), so the current _I_ decreases as the voltage approaches 0.

NMDA channels mostly allow _calcium_ ions ($Ca^{++}$) to flow into the cell, and the learning effects of this channel are due to the ability of calcium to trigger various postsynaptic chemical reactions as described in [[kinase algorithm]]. The activation effects are due to positive charges on this ion, which therefore has a net excitatory (depolarizing) effect on the cell.

[[#plot_nmda:Time run]] shows the other critical feature of the NMDA channel, which is that the _Tau_ time constant parameter (greek $\tau$) is much longer than most other channels, on the order of 100 ms or more. This parameter describes the amount of time that the channel stays open and conducting, once it has been unblocked and activated by glutamate. This relatively long time constant is critical for the activation contributions of the NMDA channel, because it creates a [[stable activation]] pattern over time (see that page for more discussion and a demonstration).

The equations we use are the same as in the widely-cited [[@BrunelWang01]] model:

{id="eq_nmda_g"}
$$
g_{nmda} = \overline{g}_{nmda} \frac{1}{1 + \frac{[Mg^{++}]}{3.57} e^{-0.062 V}}
$$

where $\overline{g}_{nmda}$ is the overall "g-bar" maximum conductance factor, and $[Mg^{++}]$ is the extracellular magnesium concentration, which is typically between 1 and 1.5. This function is of a sigmoidal, "S-shaped" form, increasing to an asymptotic value as the voltage increases; the relevant portion of this function is on the left side of the sigmoid as you can see in [[#plot_nmda:GV run]]. The decrease that we see in _I_ in is due to the standard tug-of-war Ohm's law multiplier that applies to all channel conductances:

{id="eq_nmda_i"}
$$
I_{nmda} = g_{nmda} \left(E_{nmda} - Vm\right)
$$

## GABA-B

GABA-B is an inhibitory channel activated by the usual GABA inhibitory neurotransmitter, which is coupled to the GIRK *G-protein coupled inwardly rectifying potassium (K) channel*.  It is ubiquitous in the brain, and is likely essential for basic neural function (especially in spiking networks from a computational perspective).  The inward rectification is caused by a Mg+ ion block *from the inside* of the neuron, which means that these channels are most open when the neuron is hyperpolarized (inactive), and thus it serves to *keep inactive neurons inactive*. Implementation based on [Thomson & Destexhe, 1999](#references).

In the original Leabra rate-code neurons using FFFB inhibition, the continuous nature of the GABA-A type inhibition serves this function already, so these GABA-B channels have not been as important, but whenever a discrete spiking function has been used along with FFFB inhibition or direct interneuron inhibition, there is a strong tendency for every neuron to fire at some point, in a rolling fashion, because neurons that are initially inhibited during the first round of firing can just pop back up once that initial wave of associated GABA-A inhibition passes.  This is especially problematic for untrained networks where excitatory connections are not well differentiated, and neurons are receiving very similar levels of excitatory input.  In this case, learning does not have the ability to further differentiate the neurons, and does not work effectively.

## Sodium-Gated Potassium Channels for Adaptation (kNa Adapt)

The longer-term adaptation (accommodation / fatigue) dynamics of neural firing in our models are based on sodium (Na) gated potassium (K) currents.  As neurons spike, driving an influx of Na, this activates the K channels, which, like leak channels, pull the membrane potential back down toward rest (or even below).  Multiple different time constants have been identified and this implementation supports 3: M-type (fast), Slick (medium), and Slack (slow) ([[@Kaczmarek13]]; [[@Kohn07]]; [[@Sanchez-VivesNowakMcCormick00a]]; [[@BendaMalerLongtin10]]).

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

The default parameters, which were fit to various empirical firing patterns and also have proven useful in simulations, are:

| Channel Type     | Tau (ms) | Rise  |  Max  |
|------------------|----------|-------|-------|
| Fast (M-type)    | 50       | 0.05  | 0.1   |
| Medium (Slick)   | 200      | 0.02  | 0.1   |
| Slow (Slack)     | 1000     | 0.001 | 1.0   |


