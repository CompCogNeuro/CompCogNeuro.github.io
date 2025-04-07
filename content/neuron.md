+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

This page describes the computational model of spiking neurons used in [[Axon]], which accurately characterizes the behavior of neurons in the [[neocortex]] and other brain areas, and enables many different types of [[channels]] to be used to modify [[neural integration]] behavior to capture a wide range of neurobiologically identified neuron types.

Most biological neurons have a system of ion channels that drive a brief (&lt; 1 msec) **spike** in electrical potential, followed by an _after hyperpolarization (AHP)_ that resets the potential back down to or below the resting potential. This spike triggers the **action potential** by initiating a travelling wave of depolarization down the **axon**, resulting in the release of _neurotransmitter(s)_ that then propagate the neural communication on to other neurons. Critically, there is an effective **threshold** for this spiking dynamic to be initiated, so that electrical potentials below this threshold do not result in a spike or the consequent signal being sent.

The $Na^+$ (sodium) and $K^+$ (potassium) channels underlying neural spiking were first described by [[@HodgkinHuxley52]], and have remained a cornerstone of neuroscience since then. However, the actual "HH" channel dynamics require a very fast rate of numerical integration because a lot happens in a very short period of time, so they are not computationally efficient to use directly. Instead, we adopt in axon a widely-used and well-established approximation called [[AdEx]] ( _Adaptive Exponential_; [[@BretteGerstner05]]), that uses an exponential function to approximate the voltage spike, and it also captures the spike rate [[adaptation]] dynamics of the actual HH equations.

To explore the full behavior of Axon spiking neurons interactively, see the [[neuron sim]], which allows you to observe the behavior of the different channels.

Spiking neurons have several important differences from [[rate code neurons]], which are dominant in more [[abstract neural network]] models such as those used in [[large language model]]s (LLMs), and were used in the [[Leabra]] model. In a rate code, neurons continuously communicate a floating point value representing something like the instantaneous rate of spiking.

When used in a biologically-realistic context where neural signals are being updated and communicated continuously over time, typically at a roughly 1 msec resolution, this means that a rate code neuron is constantly sending its signal to influence other neurons, with no gaps or pauses. By contrast, discrete spiking naturally creates significant periods of _silence_ in terms of the output of a given neuron, and this silence turns out to be golden, because it allows other neurons to send their signals in turn, without every neuron constantly being influenced by every other neuron. In practice, this allows spiking networks to much more robustly integrate graded and high-dimensional signals over time, compared to rate code neurons. TODO: simple sim demo!

The [[neural integration]] dynamics of biological neurons is well-described using simple electronic circuit equations, reflecting the _conductance_ of ions into and out of the cell across _ion channels_, and the resulting effects of this electrical current on the overall _electrical potential_ of the neuron, as measured across its lipid membrane (i.e., the _membrane potential_, $Vm$ ). Axon uses this standard _conductance model_ to update the membrane potential of neurons, incorporating a number of more complex ion channels with various modulatory properties, that shape the overall information integration properties of the neuron across time. See the [[neural integration]] page for full details.

Conceptually these neural integration dynamics can be understood in terms of the [[detector model]] of the neuron, where each neuron is continuously monitoring its synaptic inputs, looking for specific patterns that, when detected, cause it to signal the finding to other neurons.



