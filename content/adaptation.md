+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

**Adaptation** (also known as **accommodation** or **neural fatigue**) is the process that causes neurons to decrease their responsiveness over time, as a function of how active they have been. It is driven biologically by several different [[neuron channels#adaptation channels]] that operate over different lengths of time, most of which affect potassium (K) channel conductances.

The functional benefit of adaptation is to produce a **novelty filter** dynamic, where novel inputs hitting "fresh", non-adapted neurons generate a stronger response relative to ongoing inputs that have already been processed, and are thus driving adapted neurons ([[@Benda21]]). You can explore this behavior in the individual [[neuron simulation]], by toggling the state of various adaptation channels.

In addition to the various K channels, **synaptic depression** ([[@AbbottVarelaSenEtAl97]]) has been proposed as an important mechanism of synapse-specific adaptation, which would potentially be significantly more selective than neuron-level adaptation mechanisms. However, studies in awake, non-anesthetized animals have shown that spontaneous levels of ongoing neural activity are sufficient to effectively put most synapses in a state of tonic depression, thus limiting the functional contributions of this mechanism ([[@BoudreauFerster05]], [[@Borst10]], [[@BuchholzGuilabertEhretEtAl23]]).

[[@^Borst10]] does a particularly thorough review of the available literature and concludes that synaptic depression is unlikely to play a significant role in neural computation in awake animals. He raises a number of points that also apply generally to the distinction between studies done on slices or cultures of neurons (_in vitro_) versus the awake, intact animal, which we have termed [[in activo]], as distinct from _in vivo_ which includes anesthetized animals. In general, as [[@^Borst10]] notes, the behavior of the awake brain tends to be more _linear_ overall, with less short-term plasticity and less bursting.

Another issue with synaptic depression is that it is largely driven by the spiking of the _presynaptic_ neuron, and has little dependence on the activity state of the postsynaptic neuron ([[@NeherSakaba08]]). Therefore, it doesn't really provide any more specificity than an adaptation mechanism operating on that presynaptic neuron as a whole. Computationally, this is convenient because it is much more expensive to simulate mechanisms at the synapse vs. those operating at the neuron level.



