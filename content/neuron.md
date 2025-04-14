+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

This page describes the computational model of spiking neurons used in [[Axon]], which accurately characterizes the behavior of neurons in the [[neocortex]] and other brain areas, and enables many different types of [[channels]] to be used to modify the [[#neural-integration]] behavior to capture a wide range of neurobiologically identified neuron types.

Conceptually these neural integration dynamics can be understood in terms of the [[detector model]] of the neuron, where each neuron is continuously monitoring its synaptic inputs, looking for specific patterns that, when detected, cause it to signal the finding to other neurons.

{id="figure_cortical-neuron" style="height:30em"}
![Tracing of a cortical pyramidal neuron, showing the major components: dendrites where synaptic inputs come into the neuron; the cell body (soma) where these inputs are integrated, and the axon which communicates the output of the neuron.](media/fig_cortical_neuron.png)

Most biological neurons have a system of ion channels that drive a brief (&lt; 1 msec) **spike** in electrical potential, followed by an _after hyperpolarization (AHP)_ that resets the potential back down to or below the resting potential. This spike triggers the **action potential** by initiating a travelling wave of depolarization down the **axon**, resulting in the release of _neurotransmitter(s)_ that then propagate the neural communication on to other neurons. Critically, there is an effective **threshold** for this spiking dynamic to be initiated, so that electrical potentials below this threshold do not result in a spike or the consequent signal being sent.

The $Na^+$ (sodium) and $K^+$ (potassium) channels underlying neural spiking were first described by [[@HodgkinHuxley52]], and have remained a cornerstone of neuroscience since then. However, the actual "HH" channel dynamics require a very fast rate of numerical integration because a lot happens in a very short period of time, so they are not computationally efficient to use directly. Instead, we adopt in axon a widely-used and well-established approximation called [[AdEx]] ( _Adaptive Exponential_; [[@BretteGerstner05]]), that uses an exponential function to approximate the voltage spike, and it also captures the spike rate [[adaptation]] dynamics of the actual HH equations.

To explore the full behavior of Axon spiking neurons interactively, see the [[neuron sim]], which allows you to observe the behavior of the different channels.

Spiking neurons have several important differences from [[rate code neurons]], which are dominant in more [[abstract neural network]] models such as those used in [[large language model]]s (LLMs), and were used in the [[Leabra]] model. In a rate code, neurons continuously communicate a floating point value representing something like the instantaneous rate of spiking.

When used in a biologically-realistic context where neural signals are being updated and communicated continuously over time, typically at a roughly 1 msec resolution, a rate code neuron is _constantly_ sending its signal to influence other neurons, with no gaps or pauses. By contrast, discrete spiking naturally creates significant periods of _silence_ in terms of the output of a given neuron, and this silence turns out to be golden, because it allows other neurons to send their signals in turn, without every neuron constantly being influenced by every other neuron. In practice, this allows spiking networks to much more robustly integrate graded and high-dimensional signals over time, compared to rate code neurons. TODO: simple sim demo!

The [[#neural-integration]] dynamics of biological neurons described in detail below is well-characterized using simple electronic circuit equations, reflecting the _conductance_ of ions into and out of the cell across _ion channels_, and the resulting effects of this electrical current on the overall _electrical potential_ of the neuron, as measured across its lipid membrane (i.e., the _membrane potential_, $Vm$). Axon uses this standard _conductance model_ to update the membrane potential of neurons, incorporating a number of more complex ion channels with various modulatory properties, that shape the overall information integration properties of the neuron across time.

## Neural integration

The primary inputs to a neuron are **excitation** and **inhibition**, along with a constant **leak** channel, the dynamics of which can be understood using basic principles of electricity. We first provide a conceptual, intuitive understanding of this process, and then show how it relates to the underlying electrical properties of neurons. Then, we'll see how to translate this process into mathematical equations that can actually be simulated on the computer.

{id="figure_tug-of-war"}
![The neuron is a tug-of-war battleground between inhibition and excitation. The relative strength of each is what determines the membrane potential, Vm, which is what must get over threshold to fire an action potential output from the neuron.](media/fig_vm_as_tug_of_war.png)

The neural integration process can be understood in terms of a **tug-of-war** ([[#figure_tug-of-war]]). This tug-of-war takes place in the space of **electrical potentials** (measured in millivolts, _mV_) that exist in the neuron relative to the surrounding extracellular medium in which neurons live. Interestingly, this medium, and the insides of neurons and other cells, is basically salt water, with sodium ($Na^+$), chloride ($Cl^-$) and other ions floating around: we carry our ancient evolutionary environment around within us at all times. The electrical potentials (and differences in concentration) cause the electrically charged ions to flow in and out of the neuron through tiny pores called **ion channels**, creating small amounts of electrical **current**.

To see how this works, let's just consider excitation versus inhibition (inhibition and leak are effectively the same for our purposes at this time). The key point is that **the integration process reflects the relative strength of excitation versus inhibition:** if excitation is stronger than inhibition, then the neuron's electrical potential (voltage) increases, perhaps to the point of getting over threshold and firing an output action potential. If inhibition is stronger, then the neuron's electrical potential decreases, and thus moves further away from getting over the threshold for firing.

{id="sim_vm_gbar" title="Membrane potential tug-of-war: gbar" collapsed="true"}
```Goal
vmTau := 10.0 // time constant for vm integration
gbarE := 0.2
gbarI := 0.4
var gbarEStr, gbarIStr, vmTauStr string

##
totalTime := 100
gE := zeros(totalTime) // excitatory conductance
gI := zeros(totalTime) // inhibitory conductance
Vm := zeros(totalTime) // membrane potential
##

func vmRun() {
    gbarEStr = fmt.Sprintf("E: %7.4g", gbarE)
    gbarIStr = fmt.Sprintf("I: %7.4g", gbarI)
    vmTauStr = fmt.Sprintf("Vm Tau: %7.4g", vmTau)
    ##
    vm := 0.0 // current excitation
    tau := array(vmTau)
    gbE := array(gbarE)
    gbI := array(gbarI)
    ##
    for t := range 100 {
        ##
        ge := gbE * (1.0 - vm)
        gi := gbI * (0.0 - vm)
        dvm := (1.0 / tau) * (ge + gi)
		vm += dvm
        Vm[t] = vm
        gE[t] = gbE
        gI[t] = gbI
        ##
    }
}

vmRun()

plotStyler := func(s *plot.Style) {
    s.Range.SetMax(1).SetMin(0)
    s.Plot.XAxis.Label = "Time"
    s.Plot.XAxis.Range.SetMax(100).SetMin(0)
	s.Plot.Legend.Position.Left = true
}
plot.SetStyler(Vm, plotStyler) 

fig1, pw := lab.NewPlotWidget(b)
Vml := plots.NewLine(fig1, Vm)
gIl := plots.NewLine(fig1, gI)
gEl := plots.NewLine(fig1, gE)
fig1.Legend.Add("Vm", Vml)
fig1.Legend.Add("I", gIl)
fig1.Legend.Add("E", gEl)

func updt() {
    vmRun()
    Vml.SetData(Vm)
    gEl.SetData(gE)
    gIl.SetData(gI)
    pw.NeedsRender()
}

func addTauSlider(label *string, val *float64, mxVal float32) {
    tx := core.NewText(b)
    tx.Styler(func(s *styles.Style) {
        s.Min.X.Ch(40)  // clean rendering with variable width content
    })
    core.Bind(label, tx)
    core.Bind(val, core.NewSlider(b)).SetMin(1).SetMax(mxVal).
        SetStep(1).SetEnforceStep(true).SetChangeOnSlide(true).OnChange(func(e events.Event) {
    	updt()
        tx.UpdateRender()
    })
}

func addSlider(label *string, val *float64, mxVal float32) {
    tx := core.NewText(b)
    tx.Styler(func(s *styles.Style) {
        s.Min.X.Ch(40)  // clean rendering with variable width content
    })
    core.Bind(label, tx)
    core.Bind(val, core.NewSlider(b)).SetMin(0.02).SetMax(mxVal).
        SetStep(0.02).SetEnforceStep(true).SetChangeOnSlide(true).OnChange(func(e events.Event) {
    	updt()
        tx.UpdateRender()
    })
}

addSlider(&gbarIStr, &gbarI, 1)
addSlider(&gbarEStr, &gbarE, 1)
addTauSlider(&vmTauStr, &vmTau, 50)
```

[[#sim_vm_gbar]] provides an interactive exploration of this tug-of-war dynamic. As you drag the `E` (excitation) and `I` (inhibition) sliders, you control the strength of these two inputs, which are plotted also for easy visualization. The membrane potential `Vm` starts at 0, and is pushed up toward 1 by the E inputs, and down toward 0 by the I inputs. Thus, if you move `E` down toward 0, you can see that `Vm` barely gets off the ground, whereas if it is equal to `I` (e.g., both are 0.4) then `Vm` goes to exactly 0.5, reflecting an even balance between these opposing forces. When `E` is greater than `I` then, `Vm` goes increasingly higher, closer to 1. The `Vm Tau` slider controls the _rate_ at which `Vm` is updated, with larger values taking a longer time to converge on a stable final `Vm` value.

* What happens when `E` and `I` are both tied, but both at 0.2, or both at 0.8? Are these cases equivalent in all respects in terms of the resulting `Vm` plot? If not, in which ways do they differ?

Hopefully, you can see the _relative_ nature of this neural integration process: what matters most (though not entirely) is the relative balance between these values, not their absolute values.

The standard neuroscience notation in [[#figure_tug-of-war]] is as follows:

*  $g_i$ --- the **inhibitory conductance** (*g* is the symbol for a conductance, and *i* indicates inhibition) --- this is the total strength of the inhibitory input (i.e., how strong the inhibitory guy is tugging), and plays a major role in determining how strong of an inhibitory current there is. This corresponds biologically to the proportion of inhibitory ion channels that are currently open and allowing inhibitory ions to flow (these are **chloride** or **$Cl^-$** ions in the case of GABA **inhibition**, and **potassium** or **$K^+$** ions in the case of **leak** currents). For electricity buffs, the conductance is the inverse of resistance --- most people find conductance more intuitive than resistance, so we'll stick with it.

* $E_i$ --- the **inhibitory driving potential** --- in the tug-of-war metaphor, this just amounts to where the inhibitory guy happens to be standing relative to the electrical potential scale that operates within the neuron. Typically, this value is around -75mV where **mV** stands for **millivolts** --- one thousandth (1/1,000) of a volt. These are very small electrical potentials for very small neurons.

* $\Theta$ --- the **action potential threshold** --- this is the electrical potential at which the neuron will fire an action potential output to signal other neurons. This is typically around -50mV. This is also called the **firing threshold** or the **spiking threshold**, because neurons are described as "firing a spike" when they get over this threshold.

* $V_m$ --- the **membrane potential** of the neuron (V = voltage or electrical potential, and m = membrane). This is the current electrical potential of the neuron relative to the extracellular space outside the neuron. It is called the membrane potential because it is the cell membrane (thin layer of fat basically) that separates the inside and outside of the neuron, and that is where the electrical potential really happens. An electrical potential or voltage is a relative comparison between the amount of electric charge in one location versus another. It is called a "potential" because when there is a difference, there is the potential to make stuff happen. For example, when there is a big potential difference between the charge in a cloud and that on the ground, it creates the potential for lightning. Just like water, differences in charge always flow "downhill" to try to balance things out. So if you have a lot of charge (water) in one location, it will flow until everything is all level. The cell membrane is effectively a dam against this flow, enabling the charge inside the cell to be different from that outside the cell. The ion channels in this context are like little tunnels in the dam wall that allow things to flow in a controlled manner. And when things flow, the membrane potential changes! In the tug-of-war metaphor, think of the membrane potential as the flag attached to the rope that marks where the balance of tugging is at the current moment.

* $E_e$ --- the **excitatory driving potential** --- this is where the excitatory guy is standing in the electrical potential space (typically around 0 mV).

* $g_e$ --- the **excitatory conductance** --- this is the total strength of the excitatory input, reflecting the proportion of excitatory ion channels that are open (these channels pass **sodium** ($Na^+$) ions --- our deepest thoughts are all just salt water moving around).

{id="figure_tug-of-war-cases"}
![Specific cases in the tug-of-war scenario.](media/fig_vm_as_tug_of_war_cases.png)

[[#figure_tug-of-war-cases]] illustrates specific cases in the tug-of-war scenario. In the first case, the excitatory conductance $g_e$ is very low (indicated by the small size of the excitatory guy), which represents a neuron at rest, not receiving many excitatory input signals from other neurons. In this case, the inhibition/leak pulls much more strongly, and keeps the membrane potential ($Vm$) down near the -70mV territory, which is also called the **resting potential** of the neuron. As such, it is below the action potential threshold $\Theta$, and so the neuron does not output any signals itself. 

In the next case (b), the excitation is as strong as the inhibition, and this means that it can pull the membrane potential up to about the middle of the range. Because the firing threshold is toward the lower-end of the range, this is enough to get over threshold and fire a spike. The neuron will now communicate its signal to other neurons, and contribute to the overall flow of information in the brain's network.

The last case (c) illustrates how the integration process is sensitive to the _relative_ balance of excitation versus inhibition. If both are overall weaker, then neurons can still get over firing threshold. This is important for example in the visual system, which can experience huge variation in the overall amount of light depending on the environment (e.g., compare snowboarding on a bright sunny day versus walking through thick woods after sunset). The total amount of light coming into the visual system also drives a "background" level of inhibition, in addition to the amount of excitation that visual neurons experience. Thus, when it's bright, neurons get greater amounts of both excitation and inhibition compared to when it is dark. *This enables the neurons to remain in their sensitive range for detecting things* despite large differences in overall input levels.

## Spiking output

The membrane potential $Vm$ is not communicated directly to other neurons; instead it is subject to a **threshold**, so that only the strongest relative levels of excitation are then communicated, resulting in a more efficient and compact encoding of information in the brain. In human terms, neurons avoid sharing "TMI" (too much information), and instead communicate only relevant, important information, as if they were following ["Gricean maxims"](https://en.wikipedia.org/wiki/Cooperative_principle).

As described above, the firing of discrete spikes when $Vm$ gets above threshold occurs in biological neurons via the fast $Na^+$ and $K^+$ channels as first described by [[@HodgkinHuxley52]], and we use the [[AdEx]] ( _Adaptive Exponential_; [[@BretteGerstner05]]) model to approximate these fast dynamics. The overall cycle of spiking followed by after-hyperpolarization (AHP) and the subsequent rise in $Vm$ driven by continued excitation results in an overall **spike rate** that reflects the relative balance of excitation vs. inhibition.

In [[rate code activation]] models, this expected rate of spiking is computed directly from the inputs to a neuron, and is then computed as an overall **activation** value to other neurons. The validity of this rate code approximation is a matter of considerable debate, which is discussed further in the [[rate code activation]] page. Briefly, the [[Leabra]] model used rate code signaling, and direct comparisons with the discrete spiking [[Axon]] model show that rate codes can capture many of the same functional and cognitive phenomena as a discrete spiking model, but overall they are more brittle and require a significant tradeoff between representing graded, probabilistic information, and the speed and responsiveness of the network overall.

Specifically, a key functional advantage of discrete spiking is that considerable information about a new stimulus input can be rapidly propagated throughout the network via a cascade of _first spike_ responses [[@ThorpeDelormeVanRullen01]]. Electrophysiological recordings show that this initial wave of responding conveys significant information about the stimulus within a relatively short window of roughly 50-70 msec after it first hits the neocortex, with the relative timing of these first spikes being strongly correlated with the subsequent rate of firing. However, the subsequent firing is also critical for resolving many important properties of the stimulus, and provides a window for top-down and bottom-up signals to converge on a consistent interpretation. Thus, discrete spiking enables this "best of both worlds" of fast initial responding plus effective subsequent integration of more graded signals.

## Mathematical formulation

With the above intuitive understanding of how the neuron integrates excitation and inhibition, we can now see how a set of mathematical equations can be used to simulate this behavior in our models.

### Computing Inputs

We begin by formalizing the "strength" by which each side of the tug-of-war pulls, and then show how that causes the $Vm$ "flag" to move as a result. This provides explicit equations for the tug-of-war dynamic integration process. Then, we show how to actually compute the conductance factors in this tug-of-war equation as a function of the inputs coming into the neuron, and the synaptic weights (focusing on the excitatory inputs for now). Finally, we provide a summary equation for the tug-of-war which can tell you where the flag will end up in the end, to complement the dynamical equations which show you how it moves over time.

#### Neural Integration

The key idea behind these equations is that each side in the tug-of-war pulls with a strength that is proportional to both its overall strength (conductance), and how far the "flag" ($Vm$) is away from its position (indicated by the driving potential E). Imagine that the tuggers are planted in their position, and their arms are fully contracted when the $Vm$ flag gets to their position (E), and they can't re-grip the rope, such that they can't pull any more at this point. To put this idea into an equation, we can write the "force" or **current** that the excitatory side exerts as:

{id="eq_Ie"}
$$
I_e = g_e \left(E_e-V_m\right)
$$

The excitatory current is $I_e$ (_I_ is the traditional term for an electrical current, and _e_ again for excitation), and it is the product of the conductance $g_e$ times how far the membrane potential is away from the excitatory driving potential. If $V_m = E_e$ then the excitatory side has "won" the tug of war, and it no longer pulls anymore, and the current goes to zero (regardless of how big the conductance might be --- anything times 0 is 0). Interestingly, this also means that the excitatory side pulls the strongest when the $Vm$ "flag" is furthest away from it --- i.e., when the neuron is at its resting potential. Thus, it is easiest to excite a neuron when it's well rested.

This equation is known as **Ohm's Law**, one of the most basic laws of electricity, which you might have learned about in terms of _resistance_, which is 1 / conductance. Conductance is more intuitive in this case because it can be directly understood as the size and number of the channel openings that allow ions to flow.

The same basic equation can be written for the inhibition side, and also separately for the leak "side" (which we can now reintroduce as a clone of the inhibition term):

{id="eq_Ii"}
$$
I_i = g_i \left(E_i-V_m\right)
$$

leak current:

{id="eq_Il"}
$$
I_l = g_l \left(E_l-V_m\right)
$$

(only the subscripts are different in all of these equations).

Next, we can add together these three different currents to get the **net current**, which represents the net flow of charged ions across the neuron's membrane (through the ion channels):

{id="eq_Inet"}
$$
I_{net} = I_e + I_i + I_l
$$

$$
= g_e \left(E_e-V_m\right) + g_i \left(E_i-V_m\right) + g_l \left(E_l-V_m\right)
$$

So what good is a net current? Recall that electricity is like water, and it flows to even itself out. When water flows from a place where there is a lot of water to a place where there is less, the result is that there is less water in the first place and more in the second. The same thing happens with our currents: the flow of current changes the membrane potential (height of the water) inside the neuron:

{id="eq_Vm"}
$$
V_m\left(t\right) = V_m\left(t-1\right) + dt_{vm} I_{net}
$$

$V_m(t)$ is the current value of $Vm$, which is updated from value on the previous time step $V_m(t-1)$, and the $dt_{vm}$ is a **rate constant** that determines how fast the membrane potential changes. It mainly reflects the capacitance of the neuron's membrane.

The above two equations are the most essential tools we need to simulate a neuron on a computer. It tells us how the membrane potential changes as a function of the inhibitory, leak and excitatory inputs --- given specific numbers for these input conductances, and a starting $Vm$ value, we can then **iteratively** compute the new $Vm$ value according to the above equations, and this will accurately reflect how a real neuron would respond to similar such inputs.

To summarize, here's a single version of the above equations that does everything:

{id="eq_Vm-full"}
$$
V_m(t) = V_m(t-1) + dt_{vm} \left[ g_e (E_e-V_m) + g_i (E_i-V_m) + g_l (E_l-V_m) \right]
$$

For those of you who noticed the issue with the minus sign above, or are curious know more details about where these equations come from see , please see the Chapter Appendix section *Electrophysiology of the Neuron*. If you're happy enough with where we've come, feel free to move along to finding out how we compute these input conductances, and what we then do with the $Vm$ value to drive the output signal of the neuron.

#### Computing Input Conductances

The excitatory and inhibitory input conductances represent the total number of ion channels of each type that are currently open and thus allowing ions to flow. In real neurons, these conductances are typically measured in nanosiemens (nS), which is $10^{-9}$ siemens (a very small number --- neurons are very tiny). Typically, neuroscientists divide these conductances into two components:

* $\bar{g}$ ("g-bar") --- a constant value that determines the **maximum conductance** that would occur if every ion channel were to be open, and:

* $g\left(t\right)$ --- a dynamically changing variable that indicates at the present moment, what fraction of the total number of ion channels are currently open (goes between 0 and 1).

Thus, the total conductances of interest are written as: **excitatory conductance:**

{id="eq_gbar-e"}
$$
\overline{g}_e g_e(t)
$$

and the **inhibitory conductance:**

{id="eq_gbar-i"}
$$
\overline{g}_i g_i(t)
$$

and the **leak conductance:**

{id="eq_gbar-l"}
$$
\overline{g}_l
$$

(note that because leak is a constant, it does not have a dynamically changing value, only the constant g-bar value).

This separation of terms makes it easier to compute the conductance, because all we need to focus on is computing the proportion or fraction of open ion channels of each type. This can be done by computing the average number of ion channels open at each synaptic input to the neuron:

{id="eq_gbar-e-sum"}
$$
g_e(t) = \frac{1}{n} \sum_i x_i w_i
$$

where $x_i$ is the **activity** of a particular sending neuron indexed by the subscript *i*, $w_i$ is the **synaptic weight strength** that connects sending neuron *i* to the receiving neuron, and *n* is the total number of channels of that type (in this case, excitatory) across all synaptic inputs to the cell. As noted above, the synaptic weight determines what patterns the receiving neuron is sensitive to, and is what adapts with learning --- this equation shows how it enters mathematically into computing the total amount of excitatory conductance.

The above equation suggests that the neuron performs a very simple function to determine how much input it is getting: it just adds it all up from all of its different sources (and takes the average to compute a proportion instead of a sum --- so that this proportion is then multiplied by $\overline{g}_e$ to get an actual conductance value). Each input source contributes in proportion to how active the sender is, multiplied by how much the receiving neuron cares about that information --- the synaptic weight value. We also refer to this average total input as the **net input**.

The same equation holds for inhibitory input conductances, which are computed in terms of the activations of inhibitory sending neurons, times the inhibitory weight values.

There are some further complexities about how we integrate inputs from different categories of input sources (i.e., projections from different source brain areas into a given receiving neuron), which we deal with in the Chapter Appendix subsection *Net Input Detail*. But overall, this aspect of the computation is relatively simple and we can now move on to the next step, of comparing the membrane potential to the threshold and generating some output.

{id="sim_vm_g" title="Membrane potential tug-of-war: g's" collapsed="true"}
```Goal
vmTau := 10.0 // time constant for vm integration
gbarE := 0.2
gbarI := 0.4
var gbarEStr, gbarIStr, vmTauStr string

##
totalTime := 100
gE := zeros(totalTime) // excitatory conductance
gI := zeros(totalTime) // inhibitory conductance
Vm := zeros(totalTime) // membrane potential
##

func vmRun() {
    gbarEStr = fmt.Sprintf("gbar E: %7.4g", gbarE)
    gbarIStr = fmt.Sprintf("gbar I: %7.4g", gbarI)
    vmTauStr = fmt.Sprintf("Vm Tau: %7.4g", vmTau)
    ##
    vm := 0.0 // current excitation
    tau := array(vmTau)
    gbE := array(gbarE)
    gbI := array(gbarI)
    ##
    for t := range 100 {
        ##
        ge := gbE * (1.0 - vm)
        gi := gbI * (0.0 - vm)
        dvm := (1.0 / tau) * (ge + gi)
		vm += dvm
        Vm[t] = vm
        gE[t] = ge
        gI[t] = -gi
        ##
    }
}

vmRun()

plotStyler := func(s *plot.Style) {
    s.Range.SetMax(1).SetMin(0)
    s.Plot.XAxis.Label = "Time"
    s.Plot.XAxis.Range.SetMax(100).SetMin(0)
	s.Plot.Legend.Position.Left = true
}
plot.SetStyler(Vm, plotStyler) 

fig1, pw := lab.NewPlotWidget(b)
Vml := plots.NewLine(fig1, Vm)
gIl := plots.NewLine(fig1, gI)
gEl := plots.NewLine(fig1, gE)
fig1.Legend.Add("Vm", Vml)
fig1.Legend.Add("gI", gIl)
fig1.Legend.Add("gE", gEl)

func updt() {
    vmRun()
    Vml.SetData(Vm)
    gEl.SetData(gE)
    gIl.SetData(gI)
    pw.NeedsRender()
}

func addTauSlider(label *string, val *float64, mxVal float32) {
    tx := core.NewText(b)
    tx.Styler(func(s *styles.Style) {
        s.Min.X.Ch(40)  // clean rendering with variable width content
    })
    core.Bind(label, tx)
    core.Bind(val, core.NewSlider(b)).SetMin(1).SetMax(mxVal).
        SetStep(1).SetEnforceStep(true).SetChangeOnSlide(true).OnChange(func(e events.Event) {
    	updt()
        tx.UpdateRender()
    })
}

func addSlider(label *string, val *float64, mxVal float32) {
    tx := core.NewText(b)
    tx.Styler(func(s *styles.Style) {
        s.Min.X.Ch(40)  // clean rendering with variable width content
    })
    core.Bind(label, tx)
    core.Bind(val, core.NewSlider(b)).SetMin(0.02).SetMax(mxVal).
        SetStep(0.02).SetEnforceStep(true).SetChangeOnSlide(true).OnChange(func(e events.Event) {
    	updt()
        tx.UpdateRender()
    })
}

addSlider(&gbarIStr, &gbarI, 1)
addSlider(&gbarEStr, &gbarE, 1)
addTauSlider(&vmTauStr, &vmTau, 50)
```

[[#sim_vm_g]] shows the same excitatory vs inhibitory integration as in [[#sim_vm_gbar]], but instead plots the total current for each of these channels, with the inhibitory current shown as an absolute value (it actually has a negative sign).

## AdEx equations

To compute discrete action potential spiking behavior from the neural equations we have so far, we need to determine when the membrane potential gets above the firing threshold, and then emit a spike, and subsequently reset the membrane potential back down to a value, from which it can then climb back up and trigger another spike again, etc. This is actually best expressed as a kind of simple computer program:

``` C
if (Vm > Theta) then: y = 1; Vm = Vm_r; else y = 0
```

where y is the activation output value of the neuron, and $Vm_r$ is the *reset potential* that the membrane potential is reset to after a spike is triggered. Biologically, there are special potassium ($K^+$) channels that bring the membrane potential back down after a spike.

This simplest of spiking models is not *quite* sufficient to account for the detailed spiking behavior of actual cortical neurons. However, a slightly more complex model can account for actual spiking data with great accuracy (as shown by Gerstner and colleagues [[@BretteGerstner05]], even winning several international competitions!). This model is known as the *Adaptive Exponential* or AdEx model ([Scholarpedia Article on AdEx](http://www.scholarpedia.org/article/Adaptive_exponential_integrate-and-fire_model). We typically use this AdEx model when simulating discrete spiking, although the simpler model described above is also still an option. The critical feature of the AdEx model is that the effective firing threshold adapts over time, as a function of the excitation coming into the cell, and its recent firing history. The net result is a phenomenon called **spike rate adaptation**, where the rate of spiking tends to decrease over time for otherwise static input levels. Otherwise, however, the AdEx model is identical to the one described above.

## Normalized parameters

Managing the actual biological units for voltages, conductances and currents introduces a bit of additional complexity, which we avoid by using normalized values as in the following table:

{id="table_norms" title="Normalized units"}
| Dimension | Unit | Multiplier   | Norm range  |
|-----------|------------|--------------|--------|
| current   | amp | $10^{-8}$  | 1 = 1 nA |
| potential | volt  | 0.1 | 0..1 $\rightarrow$ -100..0 mV |
| time      | second, s $\rightarrow$ ms | 0.001 | 1 = 1 ms  |
| conductance | siemens = amp / volt | $10^{-7}$ | 1 = 100 nS |
| capacitance | farad = (sec * amp) / volt | $10^{-10}$ | 1 = 0.1 nF |

The key difference here is transforming the natural mV range of -100 to 0 mV into normalized units between 0 and 1. This transforms the reversal potentials for the standard channels as shown in the following table:

{id="table_erev" title="Reversal potentials"}
| Parameter | Bio value | Norm value |
|-----------|-----------|------------|
| Resting potential | -70 mV | 0.3   |
| Leak $E_l$        | -70 mV | 0.3   |
| Excitatory $E_e$  |   0 mV | 1.0   |
| Inhibition $E_i$  | -90 mV | 0.1   |
| Spiking threshold | -50 mV | 0.5   |

{id="table_gbar" title="g-bar conductances"}
| Parameter | Bio value | Norm value |
|-----------|-----------|------------|
| Leak $\overline{g}_l$ | 10 nS | 0.1 |
| Total excitatory $\overline{g}_e$ | 100 nS | 1.0 |
| Total inhibitory $\overline{g}_i$ | 100 nS | 1.0 |
| Nominal excitatory per synapse | 1 nS | 0.01 |


## Channel zoo

* all the other channel types and plots for each, and key discussion of issue of stability over time.

[[neuron channels]] has all the details.

## Additional details

There are a number of optional in-depth pages providing more details about biological and computational neurons.

* [[Neuron electrophysiology]]: more detailed description of the electrophysiology of the neuron, and how the underlying concentration gradients of ions give rise to the electrical integration properties of the neuron.

* [[Neuron input scaling]]: details on how excitatory and other neural inputs are computed and scaled across multiple different input projections.

* [[Neuron equilibrium potential]]: shows how to derive the _equilibrium_ (steady-state) $Vm$ equation, which clearly exhibits the relative tug-of-war dynamic.

* [[Neuron bayesian|Neuron as a Bayesian optimal detector]]: shows how the equilibrium membrane potential represents a Bayesian optimal way of integrating the different inputs to the neuron.



