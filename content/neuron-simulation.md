+++
Categories = ["Activation", "Simulations"]
bibfile = "ccnlab.json"
+++

{id="sim_neuron"}
<ccn-sim sim="neuron">

<div>

## Introduction

This simulation gives an in-depth view inside the processing within an individual [[neuron]], including the various [[neuron channels]] that shape its dynamics in important ways. See [axon/sims/neuron](https://github.com/emer/axon/tree/main/sims/neuron) for the source code.

In this model, the `Network` shows a single `Neuron`, which is "injected" with excitatory and inhibitory currents, as neuroscientists might do with an electrode injecting current into a single neuron. If you do `Run Cycles` in the toolbar you will see it get activated, but to really understand what is going on, we need to see the relationship among multiple variables as shown in the `Test Cycle Plot`.

## Overall phenomenology

* Click the [[#sim_neuron:Test Cycle Plot]] tab in the right panel to display the graph view display. If you haven't done [[#sim_neuron:Run Cycles]] yet, do it now so you can see the results of running with the default parameters, with the left panel showing the level of excitatory (`Ge`) and inhibitory (`Gi`) conductances being injected, along with some other parameters that we'll manipulate soon.

We'll start by first understanding the behavior of this neuron at an overall, qualitative level, with all the parameters at their standard default values, which are the values used in most of the other simulations in [[Axon]]. Although these neurons have a lot of parameters relative to the units in [[abstract neural network]]s, we almost never change these parameters from their default values, unless there is a clear biological or functional motivation to do so.

You can see that the level of injected input causes the neuron to fire a series of spikes, which are visible in the [[#sim_neuron:Test Cycle Plot/Vm]] membrane potential plot, and are also recorded discretely in the [[#sim_neuron:Test Cycle Plot/Spike]] variable (toggle that off to better see `Vm` alone). At the broadest level, you can see the periodic spikes that fire as the membrane potential gets over the firing threshold, and it is then reset back to the rest level, from which it then climbs back up again, to repeat the process again and again.

Because the excitation and inhibition are relatively closely balanced, you can see that the Vm starts to level off around the range just above -50 mV, due to the tug-of-war dynamics discussed in [[neuron]]. This is where the AdEx exponential function starts to contribute, due to the exponential offset value being set at -50. This allows you to see the gradual impact of the exponential spiking function over the course of a few cycles, as it drives what then becomes a clear spike in membrane potential as the exponential function does its exponential thing. This exponential dynamic is what the positive feedback loop of voltage-gated Na channels produces, as captured in the Hodgkin-Huxley model that the exponential function approximates.

The [[#sim_neuron:Test Cycle Plot/Act]] value plots the rate-code activation value that is computed from a running-average of the inter-spike-intervals (`ISIAvg`), which is the number of milliseconds between each spike. This `Act` value shows what you can hopefully see in the spikes themselves: the rate of spiking is going up over time!

Why would the rate of spiking increase, when we are only injecting a constant amount of excitation and inhibition?

* Click on [[#sim_neuron:Test Cycle Plot/Gnmda]] to see the overall NMDA channel conductance (see [[neuron channels#NMDA]] for details), which is steadily increasing over time. This channel has slower dynamics that result in [[stable activation]] patterns, and you can see that in the plot. This is what is driving the increase in firing rate over time. You can see the total excitatory conductance that includes this NMDA contribution by clicking on [[#sim_neuron:Test Cycle Plot/Ge]].

* To test the impact of NMDA, set the [[#sim_neuron:NmdaGe]] value to 0 instead of the default of 0.006, and then do [[#sim_neuron:Init]] (to clear the plot) and [[#sim_neuron:Run Cycles]] again.

This parameter determines the strength of the NMDA channel contribution to the overall `Ge(t)` value (i.e., the `Ge` value in the plot), which is the time-varying excitatory conductance (because NMDA channels have the same reversal potential as the standard excitatory AMPA channels, we just add their conductance to this overall `Ge(t)` value). When NMDA is now gone entirely, you should see that the neuron actually fails to spike at all.

* The neuron is also receiving a significant amount of inhibition from the GABA-B inhibitory channels (see [[neuron channels#GABA-B]] for details), so let's remove those too. Set the [[#sim_neuron:GababGk]] value to 0 instead of the default of 0.015, and then do [[#sim_neuron:Init]] and [[#sim_neuron:Run Cycles]] again. Note that this conductance goes into the overall `Gk(t)` potassium (K) conductance value because the GABA-B channel is coupled to a K channel.

## Adaptation

You should now see that the neuron can make a couple of spikes, but that its rate of spiking decreases over time, and soon stops. This is because the Neuron is also subject to [[adaptation]], which is driven by several different channels that operate over different time scales and in response to different activating signals.

* All of these adaptation channels contribute to the [[#sim_neuron:Test Cycle Plot/Gk]] conductance, which you can see increasing over time, and this additional leak current is what stops the neuron from firing.

* First, turn off [[#sim_neuron:KNa]] which will turn off the [[neuron channels#KNa]] sodium-gated K channels, which are one source of adaptation. [[#sim_neuron:Run Cycles]] to see the effect (if you don't hit Init then it will overlay on top of previous). You should see more spikes making it through.

* Next, set [[#sim_neuron:MahpGk]] to 0, which will turn off the [[neuron channels#mAHP]] M-type voltage-gated K channel, which drives medium timescale afterhyperpolarization (AHP) dynamics. [[#sim_neuron:Run Cycles]] to see the effect. You should see now that the rate of spiking is perfectly consistent over time, as you would expect from the constant level of excitatory and inhibitory inputs.

## Basic conductances

Now that things are simpler, we can drill down and understand the basic excitatory, inhibitory, and leak conductances, and how they drive the current which updates the membrane potential Vm.

* Turn on [[#sim_neuron:Test Cycle Plot/Ge]] = total excitatory input conductance to the neuron, which is generally a function of the number of open excitatory synaptic input channels at any given point in time (`Ge(t)`) and the overall strength of these input channels, which is given by `Gbar E`.  In this simple model, `Ge(t)` goes from 0 prior to cycle 10, to .75 from 10-160, and back to 0 thereafter.

The [[#sim_neuron:Ge]] parameter is set to 0.15 -- why are we getting 0.75 in the plot? The reason is that we're injecting 0.15 of _raw_ new synapse-level conductance that is supposed to be coming from AMPA channels. However, these AMPA channels stay open for 5 ms, i.e., 5 time steps, so they effectively accumulate 5x this per-step input, and indeed $5 * 0.15 = 0.75$.

{id="question_gi"}
> Now you should be able to explain the value of [[#sim_neuron:Test Cycle Plot/Gi]] in the plot, in relation to the [[#sim_neuron:Gi]] parameter of 0.1 (hint: the decay time constant for GABA-A inhibitory channels is 7 ms).

* Turn off `Ge` and `Gi`, and turn on [[#sim_neuron:Test Cycle Plot/Inet]] which shows the net current as the sum of excitatory, inhibitory and leak currents. Note that its values are rather large, especially during the spike, so it will swamp anything else on the left axis; Vm is plotted on the right axis, so you can see both at the same time.

{id="question_inet"}
> Based on what you've learned about the relationship between Inet and Vm, can you explain why Inet "spikes" as the Vm first rises when the input first comes on, and after each spike, and then it goes back toward 0 as the Vm stops changing as much, as it approaches the threshold level of -50 mV. 

## Balance of excitation and inhibition

From the tug-of-war model, you should expect that increasing the amount of excitation coming into the neuron will increase the rate of firing, by enabling the membrane potential to reach threshold faster, and conversely decreasing it will decrease the rate of firing. Furthermore, increasing the leak or inhibitory conductance will tug more strongly against a given level of excitation, causing it to reach threshold more slowly, and thus decreasing the rate of firing.

This intuitive behavior is the essence of what you need to understand about how the neuron operates. 

* Increase the [[#sim_neuron:Ge]] excitation from 0.15 to 0.16 (and then do Init and Run to see the effects). Then observe the effects of increasing [[#sim_neuron:Gi]] inhibition from 0.1 to 0.11. Go further and increase inhibition to 0.12.

{id="question_ge-gi"}
> Describe the qualitative effects on the rate of neural spiking of increasing Ge from 0.15 to 0.16, and then increasing Gi 0.1 to 0.11.

{id="question_gi-hi"}
> Is there a qualitative difference in the neural spiking when Gi is increased to 0.12 -- what important aspect of the neuron's behavior does this reveal?

### Driving / Reversal Potentials

* Set Ge back to .15 and Gi to .1, then set [[#sim_neuron:ErevE]] to -2 instead of 0, and Run. Surprisingly, this small change prevents the neuron from spiking, by reducing the driving potential of excitation. Likewise, lowering [[#sim_neuron:ErevI]] from -90 to -92 is sufficient to eliminate spiking.

## Noise

An important aspect of spiking in real neurons is that the timing and intervals between spikes can be quite random, generally obeying a Poisson distribution, which has the highest variance for a given rate of spiking. As in the example here, most excitatory neurons in the cortex receive closely balanced excitation and inhibition, due to the mechanisms covered in [[inhibition]], and thus even relatively small differences in the timing and relative strength of inputs can make a big difference in the neural response, due to the strong threshold nonlinearities present. For example, the parameter manipulations above demonstrated how relatively small changes resulted in the difference between spiking and not spiking at all.

To simulate a population of noisy inputs, the noise mechanism we use adds discrete excitatory and inhibitory conductances generated from a Poisson distribution of a given frequency (100 Hz for excitation, 200 Hz for inhibition, reflecting the general differences in their spike rates).

* First, do [[#sim_neuron:Defaults]] to return to default parameters, and then set the strength of these noisy conductances via the [[#sim_neuron:Noise]] parameter: set it to 0.05 instead of 0, and then do Run multiple times without doing Init.

You should observe that the timing of the spikes varies significantly across the different runs, such that the plot starts to fill in, especially after the first couple of spikes. You can switch to only plotting `Spike` to see this more clearly. Then switch to plotting `Ge` to see the magnitude of the noise being added to the excitatory conductance. Try doing Init and comparing Noise of 0 to 0.05 to better see the difference.

Thus, even a relatively modest amount of noise can produce significant differences in spike timing, due to the nonlinear positive feedback loops involved in neural spiking. If you increase the Noise conductance level further, you can get even more uniform (i.e., highly variable) distributions of spike timings across runs.

## Links

Next in [[Intro Book]]: [[Detector simulation]]

</div>


