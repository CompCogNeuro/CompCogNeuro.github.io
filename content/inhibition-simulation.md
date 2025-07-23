+++
Categories = ["Activation", "Simulations"]
bibfile = "ccnlab.json"
+++

{id="sim_inhib" collapsed="true"}
```Goal
// see https://github.com/emer/axon/tree/main/sims/inhib for source code
inhib.Embed(b)
```

<div>

This simulation explores how inhibitory interneurons can dynamically control overall activity levels within the network, by providing both feedforward and feedback inhibition to excitatory pyramidal neurons, with different time scales provided by PV neurons (fast spiking) and SST neurons (slow spiking). See [[inhibition]] for the biological basis and equations.

The `Network` has two layers of excitatory neurons (`Layer1`, `Layer2`) that each have an associated set of inhibitory interneurons (`Inhib1`, `Inhib2`) that are only relevant when not using the `FSFFFB` inhibition function. 

* Click [[#sim_inhib:Wts]] in the Network variables, and then on [[#sim_inhib:r.Wt]] to view the receiving weights into neurons as you click on them in the network.

You should see that the two excitatory layers have [[bidirectional connectivity]]. This positive feedback loop really pushes the need for strong inhibitory control so that the network does not spiral out of control into a state where all neurons are active (we will see how easy it is to get into such a state).

* Click [[#sim_inhib:Act]] (category then variable) and click on [[#sim_inhib:Step]] which will run one 200 ms `Trial` of updating. 

You should see that a _sparse_ subset of the neurons are active in response to a random input pattern that has 15% of the 100 input `Layer0` neurons active. Thus, the inhibition function is doing its job, maintaining sparse [[distributed representations]] despite the presence of the positive feedback loops.

* To see more of the time-course of activity across the trial, turn the [[#sim_inhib:Raster]] view on in the Network toolbar, and then [[#sim_inhib:Step]] again. Also try switching to viewing [[#sim_inhib:Spike]] to see the discrete spiking events.

You should see a relatively [[stable activation]] pattern over time, which depends critically on the properties of the inhibition function in addition to the [[neuron channels#NMDA]] and [[neuron channels#GABA-B]] channels that impart stability on individual neurons.

* Click on [[#sim_inhib:Test Cycle Plot]] to see a plot of the layer activity over the 200 ms trial. This is an average of the `Act` values in each layer, which is lower than the .15 (15%) value because the neural activity levels (rate code activations) are significantly less than 1 for the neurons that did get active. This is typical, especially prior to significant learning, as is the case with this network.

## Plotting the fast and slow components

To get a better sense of how the FS (fast spiking, PV) and SS (slow spiking, SST) components contribute to the overall Gi inhibitory conductance, you can click on the [[#sim_inhib:Layer1_FSGi]], [[#sim_inhib:Layer1_SSGi]], and [[#sim_inhib:Layer1_Gi]] plot lines. You should see that the FSGi conductance component consists of a rapidly and regularly oscillating sawtooth-like pattern (driven by the feedforward inputs from the input layer where the neurons fire relatively synchronously and regularly), while SSGi has a slow and smooth trajectory. The Layer2 versions of these values are qualitatively similar, although the FS component is more irregular.

Although the SS component of this function is not as obviously important in this simple test model (see below for parameter manipulations), it is critical for actual models that learn over time: it is essential for the inhibitory response to be responsive to the changes from learning, and also makes the inhibition more stable and robust overall.

## Manipulating parameters

You can explore the effects of the various parameters listed on upper left panel of the simulation.

**Be sure to press [[#sim_inhib:Init]] to get new parameters to be applied to network!**

* [[#sim_inhib:Gi]] controls overall inhibition: decreasing Gi will increase activity and vice-versa for increases. Explore effects on activity levels and also firing patterns as evident in the raster plot and the cycle plot.

{id="question_graded"}
> Does the network overall exhibit a graded, roughly linear response to the different Gi levels, or is there a clear point at which it just "blows up" or completely shuts down in a disproportional "nonlinear" way?

{id="question_oscillations"}
> As you decrease the inhibition, do you notice a point at which the activity levels start to oscillate?

Overall, the behavior of the discrete spiking neurons in [[Axon]] is considerably more robust relative to the [[rate-code activation]] in the [[Leabra]] model, which has a much more nonlinear response to inhibition parameters, and will fully saturate at maximal activity levels once it gets below some critical minimum level of inhibition. By contrast you should have noticed that the spiking neurons tend to start oscillating well before they exhibit fully saturated spiking. This is consistent with the phenomenology of epilepsy, which is characterized by strongly synchronized firing that comes in waves due to the refractory dynamics that shut off activity between waves.

* [[#sim_inhib:FSTau]] sets the [[exponential integration]] time constant for integrating the fast-spiking component, which also affects the total magnitude of the FS component overall. If you decrease this to 2 you'll see that the activity patterns become highly oscillatory with lots of peak activity. Even if you increase Gi to 3 to control the overall activity level better, it remains strongly oscillatory. Likewise increasing FSTau to 10 and reducing Gi to 0.6 to compensate for the larger magnitude also produces a more oscillatory pattern of activity. The default value of 6 aligns well with the excitatory and inhibitory conductance time constants of 5 and 7 ms respectively.

* [[#sim_inhib:SS]] sets the slow spiking (SST) contribution to the total Gi value. The default value of 30 may seem like it is the dominant factor, but if you click on the `SSi` values in the `Test Cycle Plot`, you can see that the raw input value is relatively small. If you set this value to 0, you can increase the overall Gi to mostly restore the overall activity level. However, in actual learning models, this SS component is critical.

You can also explore the other parameters and observe their effects on the SS and FS values to better understand their contributions.

## Inhibition from interneurons

To see how the results of the FS-FFFB inhibitory function compare with the use of actual inhibitory interneurons, you can click the [[#sim_inhib:FSFFFB]] switch of and do Init and Step. This exposes two new parameters, `InhibExcite` and `InhibInhib` which determine the strength of inhibitory connections onto excitatory and inhibitory neurons, respectively. 

The parameters for the `Inhib` inhibitory neurons have been set to drive the fast spiking behavior of PV neurons, so the contribution of the SST population is not represented here. Nevertheless, especially after the first trial, the activity patterns in the excitatory layers do not look too dissimilar from those with the FS-FFFB function. However, because these neurons only deliver inhibition when they spike, it would take a larger population with more diverse firing patterns to provide a reasonable approximation to the larger populations of such neurons in the real brain, which can provide a smoother, more continuous source of inhibition, which the FS-FFFB does a better job of providing. The key difference is in the strong tendency toward more oscillatory overall activity patterns with these discrete interneurons. 

In summary, the FS-FFFB function saves a lot of computational resources vs explicitly simulating inhibitory interneurons, and it produces results that better match those that would result from a larger population of such neurons. Therefore, it is used for all of the [[Axon]] models.

</div>

