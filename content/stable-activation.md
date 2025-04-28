+++
Categories = ["Activation", "Axon", "Simulations"]
bibfile = "ccnlab.json"
+++

{id="sim_stability" collapsed="true"}
```Goal
// see https://github.com/CompCogNeuro/CompCogNeuro.github.io/sims/stability for source code
stability.EmbedSim(b)
```

<div>

## Introduction

The key development that finally enabled use of discrete spiking neurons within the context of the [[Leabra]] framework (resulting in the [[Axon]] framework that is the focus of this content), was adding relatively weak conductances of [[neuron channels#NMDA]] and [[neuron channels#GABA-B]] voltage-gated channels. This resulted in the emergence of relatively stable patterns of neural activity over time. The present simulation demonstrates this effect by simply presenting an input pattern and observing the resulting stability of the network response, while varying the strength of the NMDA and GABA-B channel conductances.

There is a long history of theorizing about the importance of NMDA channels for stabilizing neural activity in the context of [[working memory]] supported by the [[prefrontal cortex]] (PFC) ([[@Goldman-Rakic95]]; [[@LismanFellousWang99]]; [[@BrunelWang01]]). [[@^SandersBerendsMajorEtAl13]] further showed that the long time constants and inward rectification of the GABA-B inhibitory channel provide an important additional stabilizing force. In the context of implementing these mechanisms in Leabra for PFC working memory models, it became clear that these same dynamics in weaker form would solve the longstanding problem with all prior attempts to use spiking in Leabra: the representations were too unstable and constantly shifted over time.

A stable pattern of neural activity over roughly 200 ms in response to a given input stimulus is essential for [[predictive learning]] based on a [[temporal derivative]], because the learning mechanism is essentially comparing activity patterns across this 200 ms interval. If these patterns are unstable and constantly shifting, then this comparison is meaningless, and learning fails.

It is also the case that almost all theories of [[conscious awareness]] include a critical contribution for the stable integration of information across this same time scale (e.g., [[@Lamme06]]), and NMDA receptor agonists such as ketamine (a PCP analog) have significant effects on the state of consciousness and cortical function more generally, including inducing psychosis and hallucinations ([[@DriesenMcCarthyBhagwagarEtAl13]]; [[@NewcomerFarberJevtovic-TodorovicEtAl99]]; [[@KrystalAnticevicYangEtAl17]]). These NMDA manipulations affect both excitatory and inhibitory neurons throughout the brain, so they are a relatively "blunt" instrument, but nevertheless the effects are consistent with the current model.

The fact that these stabilizing mechanisms were _not_ needed with the [[rate-code activation]] used in [[Leabra]] is also indicative of the critical differences between discrete spiking and rate codes: because rate code neurons continuously send their activation state, they are intrinsically more stable. However, this stability can actually be a liability because it prevents a network from representing different interpretations of a given input stimulus.

## Exploration

The `Network` has three layers, which are functionally organized as an _Input_ layer and 2 _Hidden_ layers that represent internal [[neocortex|cortical]] processing levels. There are actually many more such layers in the cortex of most mammals, but we can see the issues already with just these two.

* Click [[#sim_stability:Step]] to run one trial of 200 ms cycles for a single randomly-generated input pattern. 

You will see a pattern of neural activity across the layers, which has a very clear differentiation between active and inactive neurons in the first, input layer, that gets more diffuse and graded in higher layers. The active neurons have progressively brighter, saturated, and more yellow colors, with the height of the neuron-level cubes set in proportion to their activity, while inactive ones are grey and flat. The synaptic connections in this network are completely random, so it is expected that the clear signal in the input is diluted as it progresses through this untrained network.

* Click [[#sim_stability:Network/Raster]] in the Network toobar to see a _raster_ plot of neural activity over time, and do [[#sim_stability:Step]]  to see the results.

Now you will see "streaks" of activity going from the front to back of each layer for each neuron that got active, which plot each of the 200 1 ms cycle steps of updating for each neuron, with the neurons arrayed horizontally across the layer. Switch to viewing [[#sim_stability:Network/Spike]] to see individual spikes, which will give you a better sense of the time dimension.

Hopefully you can see that overall there is a relatively consistent pattern of activity over time (i.e., from front to back of each layer), with more variability in the top-most layer.

* Now let's see what happens when we eliminate NMDA and GABA-B channels. Set [[#sim_stability:Nmda ge]] and [[#sim_stability:Gabab gk]] to 0, then do [[#sim_stability:Init]] and [[#sim_stability:Step]] again, while still viewing `Spike` state.

You should see a row of spiking activity at almost the same cycle in the top layer for most of the neurons, followed by a long pause and then some sporadic spiking after that. This increase in neural _synchrony_ is also characteristic of unconscious states, whereas waking, conscious brain states are characterized by much weaker levels of synchrony and sparser overall activity levels, with only a fraction of the excitatory neurons active (10-15% in general).

Because the NMDA channels provide additional excitation (as you can see in the [[neuron sim]]), we can try reducing the inhibition level, to allow more overall excitation.

* Set [[#sim_stability:Inhib gi]] to 0.8 instead of 0.85, which reduces the strength of the GABA-A conductance computed as described in [[inhibition]]. Then do `Init` and `Step` again. Keep reducing it, and also try doing multiple `Step`s to see what happens over different input patterns. The network activity state carries over across different inputs, which will sometimes disrupt the synchronous firing.

{id="question_gi"}
> Were you able to find a level of inhibition that resulted in level of activity and sparseness that was comparable to the state with NMDA and GABA-B conductances at their default values? Were you able to get it to exhibit a _stable_ pattern of activity, where the same sparse subset of neurons remains active across the 200 ms trial?

In summary, you should have observed that NMDA and GABA-B conductances are critical for supporting a stable activity pattern. This effect is especially critical as activity cascades across multiple layers, due to a diffusion-like process where activity spreads out as a result of the random initial synaptic weights. Thus, we saw in this model that the diffuse activity patterns were more pronounced the 2nd hidden layer.

If you repeat this experiment in any of the other learning models explored on other pages here, you'll readily see that turning off NMDA and GABA-B dramatically impairs learning, because these diffuse, unstable activity patterns undermine the learning signals.

</div>

