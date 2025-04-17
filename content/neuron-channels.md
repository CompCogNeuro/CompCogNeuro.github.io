+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

This page provides details for the full range of channel types that are available in an [[Axon]] [[neuron]] to drive biologically-based and functionally important behavior in specific neuron types. Every neuron uses the basic excitatory, inhibitory, and leak channels discussed in detail in the [[neuron]] and [[neuron electrophysiology]] pages, and some of the following channels are used in most other neurons, while some are only used in specific neuron types where they are particularly important.

The implementation of several of these channels comes from standard biophysically detailed models such as [[@^MiglioreHoffmanMageeEtAl99]], [[@^PoiraziBrannonMel03]], and [[@^UrakuboHondaFroemkeEtAl08]]. See also [[@^BretteRudolphCarnevaleEtAl07]] and the [NEST model directory](https://nest-simulator.readthedocs.io/en/stable/models/index.html) for documented examples, including: [AdEx](https://nest-simulator.readthedocs.io/en/stable/models/aeif_cond_exp.html), [Traub HH](https://nest-simulator.readthedocs.io/en/stable/models/hh_cond_exp_traub.html).  The [Brian Examples](https://brian2.readthedocs.io/en/stable/examples/index.html) contain full easy-to-read equations for various standard models, including [Brunel & Wang, 2001](https://brian2.readthedocs.io/en/stable/examples/frompapers.Brunel_Wang_2001.html). Also see [Wikipedia: Biological neuron model](https://en.wikipedia.org/wiki/Biological_neuron_model) for a nice overview, and [ModelDB Currents](https://modeldb.science/NeuronDB/NeuronalCurrents), [ModelDB Current Search](https://modeldb.science/ModelDB/FindByCurrent), and [IonChannelGeneology](https://icg.neurotheory.ox.ac.uk) for standardized lists of currents included in biophysical models made in NEURON and related software.

## Normalized parameters

The [[neuron#normalized parameters]] for all of the channel types covered here are shown in the following tables.

{id="table_taus" title="Time constants"}
| Parameter                        | Bio value | Norm value |
|----------------------------------|-----------|------------|
| $\tau_{ampa}$                    | 5 ms      | 5          |
| $\tau_{gaba_a}$                  | 7 ms      | 7          |
| $\tau_{nmda}$                    | 100 ms    | 100        |
| $\tau_{gaba_b}$ rise ($\tau_r$)  | 45 ms     | 45         |
| $\tau_{gaba_b}$ decay ($\tau_d$) | 50 ms     | 50         |


{id="table_gbar-other" title="g-bar conductances"}
| Parameter        | Bio value | Norm value |
|---------------------------|-----------|------------|
| NMDA                      | 50 nS     | 0.50       |
| $\rm{VGCC}_{distal}$      | 146 nS    | 1.46       |
| $\rm{VGCC}_{proximal}$    | 3.2 nS    | 0.032      |
| $\rm{VGCC}_{soma}$        | 93 nS     | 0.93       |
| $aK_{distal}$             | 49 nS     | 0.49       |
| $aK_{proximal}$           | 7.5 nS    | 0.075      |

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

{id="plot_nmda" title="NMDA Channels" collapsed="true"}
```Goal
pl := &chanplots.NMDAPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "NMDA", "NMDA Channel", "NMDA Channel equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
```

NMDA (N-methyl-D-aspartate) channels are found throughout the brain, and play a critical role in learning as captured in the [[kinase algorithm]] used in [[Axon]]. The opening of NMDA channels is typically blocked by positively-charged magnesium ions ($Mg^{++}$) when the membrane potential is close to the resting potential. The removal of this block as a result of membrane depolarization above this resting potential (known as an _outward rectification_, because the $Mg^{++}$ ions are on the outside of the cell) is one of two key functional features of this channel. This unblocking dynamic is important both for learning and for its important contributions to the activation dynamics of the neuron.

The other critical functional property of the NMDA channel is that it also requires _glutamate_ neurotransmitter binding to open, in addition to $Mg^{++}$ unblocking. Thus, unlike many other channels that just have one (or none) _gating_ factors, NMDA requires both of these factors. Furthermore, the $Mg^{++}$ factor is a function of the _postsynaptic_ activity (depolarization), while glutamate is released by the  _presynaptic_ neuron. Thus, the NMDA is in a unique position to respond to the _conjunction_ of both pre and postsynaptic activity.

To see the unblocking in action, press the [[#plot_nmda:GV run]] button above, which generates a conductance (_g_) (and current _I_) vs. voltage (_v_) plot using the parameters shown to the left of the plot, assuming that there is plenty of glutamate around so that factor is not relevant. As the voltage increases above the -90 hyperpolarized starting point, the conductance steadily rises, reflecting the progressively increased likelihood that the $Mg^{++}$ ions will not be blocking the channel opening. The _reversal potential_ for the NMDA channel is around 0 mV, so as the voltage approaches this point, the net force pulling the ions through the channel gets progressively weaker (as explained by the tug-of-war analogy in [[neuron]]), so the current _I_ decreases as the voltage approaches 0.

NMDA channels mostly allow _calcium_ ions ($Ca^{++}$) to flow inanto the cell, and the learning effects of this channel are due to the ability of calcium to trigger various postsynaptic chemical reactions as described in [[kinase algorithm]]. The activation effects are due to positive charges on this ion, which therefore has a net excitatory (depolarizing) effect on the cell.

[[#plot_nmda:Time run]] shows the other critical feature of the NMDA channel, which is that the _Tau_ time constant parameter (greek $\tau$) is much longer than most other channels, on the order of 100 ms or more. This parameter describes the exponential decay constant (of the same form as [[#eq_ampa_g]]) for the NMDA conductance, which like AMPA has a sufficiently fast rise time that it can be ignored. This relatively long time constant is critical for the activation contributions of the NMDA channel, because it creates a [[stable activation]] pattern over time (see that page for more discussion and a demonstration).

The equation we use for the voltage-gated conductance is from the widely-cited [[@BrunelWang01]] model:

{id="eq_nmda_g" title="NMDA voltage-gated conductance"}
$$
g_{nmda}(V) = \overline{g}_{nmda} \frac{1}{1 + \frac{[Mg^{++}]}{3.57} e^{-0.062 V}}
$$

where $\overline{g}_{nmda}$ is the overall "g-bar" maximum conductance factor, and $[Mg^{++}]$ is the extracellular magnesium concentration, which is typically between 1 and 1.5. This function is of a sigmoidal, "S-shaped" form, increasing to an asymptotic value as the voltage increases; the relevant portion of this function is on the left side of the sigmoid as you can see in [[#plot_nmda:GV run]]. The decrease that we see in _I_ as voltage increases is due to the standard tug-of-war Ohm's law multiplier that applies to all channel conductances ([[#eq_ampa_i]]).

## GABA-B

{id="plot_gabab" title="GABA-B Channels" collapsed="true"}
```Goal
pl := &chanplots.GABABPlot{}
root, _ := tensorfs.NewDir("Root")
br := egui.NewGUIBody(b, pl, root, "GABA-B", "GABA-B Channel", "GABA-B Channel equations")
pl.Config(root, br.Tabs)
br.FinalizeGUI(false)
```

The GABA-B channel has a much slower decay time constant compared to the standard GABA-A inhibitory channel, because it is coupled to the GIRK channel: _G-protein coupled inwardly rectifying potassium (K) channel_. It is ubiquitous in the brain, and is likely essential for basic neural function. The _inward rectification_ is caused by a $Mg^{++}$ ion block from the _inside_ of the neuron, which means that these channels are most open when the neuron is hyperpolarized (inactive), and thus it serves to _keep inactive neurons inactive_. This is complementary to the effect of NMDA channels, and [[@^SandersBerendsMajorEtAl13]] emphasized the synergistic nature of these two channels in producing [[stable activation]] patterns: NMDA keeps active neurons active, while GABA-B keeps inactive neurons inactive.

Press the [[#plot_gabab:GV run]] button above to see the conductance (_g_) (and current _I_) vs. voltage (_v_) plot using the parameters shown to the left of the plot. In comparison with the comparable NMDA plot, you can see that GABA-B and NMDA are mirror-images of each other. Furthermore, the I value plotted here is the pabsolute value (positive) whereas the actual current has the opposite sign as NMDA, due to its reversal potential $E_{gaba_b}$ being that of chloride (-90 mV).

The implementation of the GABA-B channel is based on [[@^SandersBerendsMajorEtAl13]] and [[@^ThomsonDestexhe99]], with the following sigmoidal voltage-gated conductance function from [[@^YamadaInanobeKurachi98]]:

{id="eq_gabab_g" title="GABA-B voltage-gated conductance"}
$$
g_{gaba_b}(V) = \overline{g}_{gaba_b} \frac{1}{1 + e^{0.1(V-E_{gaba_b}+10)}}
$$

There is an additional sigmoidal function needed for computing the time dynamics of the GABA-B conductance as a function of the GABA binding from inhibitory input spikes over time (in Figure 16 of [[@ThomsonDestexhe99]]):

{id="eq_gabab_x" title="GABA-B spiking rate integration"}
$$
x = \frac{1}{1 + e^{-(s - 7.1) / 1.4}}
$$

Where _s_ is the rate of spiking over the recent time window, which we compute based on the $g_i$ inhibition factor from the [[inhibition]] function used in [[Axon]], and _x_ drives increases in GABA-B conductance according to this update equation with separate rise ($\tau_r$) and decay ($\tau_d$) factors (45 and 50 ms respectively; [[#table_taus]]):

{id="eq_gabab_dg" title="GABA-B conductance over time"}
$$
g_{gaba_b}(t) = \tau_r \left( \left[(\tau_d / \tau_r)^{(\tau_r / (\tau_d - \tau_r))} \right] x(t) - g_{gaba_b}(t-1) \right)
$$

$$
x(t) = \left( \frac{1}{1 + e^{-(s - 7.1) / 1.4}} - x(t-1) \right) - \tau_d x(t-1)
$$

The final GABA-B conductance is a product of this time-integration factor above in [[#eq_gabab_dg]] and the voltage gating factor shown in [[#eq_gabab_g]]:

{id="eq_gabab_tg" title="GABA-B net conductance over time"}
$$
g_{gaba_b}(t) = g_{gaba_b}(V) g_{gaba_b}(t)
$$

Do [[#plot_gabab:Time run]] to see these time dynamics play out over a 500 ms window with a pulse of input at the start.

## Sodium-gated potassium channels for adaptation (kNa adapt)

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


