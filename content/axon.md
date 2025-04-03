+++
Categories = ["Learning", "Rubicon"]
bibfile = "ccnlab.json"
+++

**Axon** is the successor to [[leabra]], featuring more realistic discrete spiking dynamics and associated learning mechanisms and more advanced brain area models to implement a functional version of the [[Rubicon]] goal-driven learning system, based on the interactions of a large number of "limbic" brain areas. Axon is used for all of the newer content here on CompCogNeuro.

These are the central elements of axon in terms of general neural mechanisms:

* [[Spiking neuron|Spiking neurons]] with long-lasting [[NMDA]] and [[GABA-B]] channels that support relatively [[stable activation]] states over the course of a roughly 200 msec [[theta cycle]], which is essential for allowing a meaningful _representation_ of the current input state, that is needed to drive learning. Spiking enables effective graded information integration over time in a way that continuous rate code activation communication does not, by allowing many different signals to be communicated over time, competing for the overall control of the network activation state as a function of the collective [[neural integration|integration]] of spikes within all the neurons in the network.

* [[Error-driven learning]] based on the encoding of errors using a [[temporal derivative]] that naturally supports [[predictive learning]] in terms of the difference over time of a network state representing a prediction followed by an outcome. Local [[synaptic plasticity]] based on the competition between kinases updating at different rates, i.e., the [[kinase algorithm]], naturally computes the error gradient via the [[temporal derivative]] property, supporting a fully biologically plausible form of the computationally powerful (and ubiquitous) [[error backpropagation]] algorithm.

* [[Bidirectional excitatory connectivity]] which is necessary for propagating error signals throughout the network, and [[surround inhibition]] which is necessary for controlling the effects of bidirectional excitatory connectivity, while also having beneficial computational effects in terms of [[attention]] and [[competition]], including likely playing an essential role in [[conscious awareness]].

Another critical hypothesis for the axon framework is that locally-computed error-driven learning and processing using these general neural mechanisms is _insufficient_ for successful learning and performance in complex real-world environments. In addition, organized neural networks are required to drive goal-driven learning and performance across longer time scales in a coordinated manner, that is sensitive to the overall costs and benefits of actions. This goal-driven learning provides critical error signals that shape overall network learning. The [[Rubicon]] framework provides an implementation of this goal-driven learning system, based on the detailed biology of the [[prefrontal cortex]], [[basal ganglia]], [[amygdala]], [[dopamine]], and the [[hippocampus]] (traditionally known as the [[limbic system]]), among others.

## Axon pages

