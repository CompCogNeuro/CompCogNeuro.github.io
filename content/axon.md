+++
Categories = ["Learning", "Rubicon"]
bibfile = "ccnlab.json"
+++

**Axon** is an attempt to build a biologically accurate, yet still abstracted and computationally tractable model of neural computation in the mammalian brain (e.g., in rodents, monkeys, and people). The goal is to commit as few _errors of commission_ as possible: the features of the model should not violate any known, well-established properties of the brain. There is still plenty of room for _errors of omission_: not every detail can or should be included.

The approach is guided by an in-depth consideration of the available neuroscience literature, as well as a strong consideration for the functional and computational levels of analysis, informed by an understanding of the known cognitive and behavioral capabilities of the respective organisms. In other words, it is [[Computational Cognitive Neuroscience]], leveraging notable [[synergies]] across these different levels of analysis. For an introductory textbook-like coverage of this content, see [[CCN Intro]].

Axon is the successor to [[Leabra]], featuring more realistic discrete spiking dynamics and associated learning mechanisms, and more advanced brain area models to implement a functional version of the [[Rubicon]] goal-driven learning system, based on the interactions of a large number of [[limbic system|limbic]] brain areas. Axon is used for all of the newer content here on `CompCogNeuro.org`, and all of this content is presented in the context of the scientific hypotheses behind this framework (i.e., we do not aspire to a wikipedia-like "neutral point of view"). Nevertheless, we aspire to a high standard of scholarly rigor, and appreciate any feedback about potential inaccuracies with any of the information presented here.

These are the central elements of Axon in terms of general neural mechanisms which are well established properties of the mammalian [[neocortex]]:

* [[neuron|Spiking neurons]] with long-lasting [[neuron channels#NMDA]] and [[neuron channels#GABA-B]] channels that support relatively [[stable activation]] states over the course of a roughly 200 msec [[theta cycle]], which is essential for establishing a coherent representation of the current input state. This stability is necessary to drive effective learning as described next. 

	Spiking also enables effective graded information integration over time in a way that continuous [[rate code activation]] communication does not, by allowing many different signals to be communicated over time, competing for the overall control of the network activation state as a function of the collective integration of spikes within the neurons in the network. As a result, Axon models are overall much more robust and well-behaved overall compared to their Leabra rate-code based counterparts.

* [[Error-driven learning]] based on the encoding of errors using a [[temporal derivative]] that naturally supports [[predictive learning]] in terms of the difference over time of a network state representing a prediction followed by an outcome. Local [[synaptic plasticity]] based on the competition between kinases updating at different rates, i.e., the [[kinase algorithm]], naturally computes the error gradient via the [[temporal derivative]] property, supporting a fully biologically plausible form of the computationally powerful (and ubiquitous) [[error backpropagation]] algorithm. Initial empirical support for this mechanism is reported in [[Jiang et al 2025]].

* [[Bidirectional connectivity]] among excitatory neurons, which is necessary for propagating error signals throughout the network, and pooled [[inhibition]] which is necessary for controlling the effects of bidirectional excitatory connectivity, while also having beneficial computational effects in terms of [[attention]] and [[competition]]. Bidirectional connectivity also supports multiple [[constraint satisfaction]] dynamics that can efficiently search through large high-dimensional knowledge spaces to find (and synthesize) the most relevant information given the current bottom-up (sensory) and top-down (goals) constraints.

   Perhaps most importantly, this bidirectional connectivity is widely thought to be essential for [[conscious awareness]] ([[@Lamme06]]), which is likely critical for the system to access its own internal state of knowledge. This ability is notably absent in current [[abstract neural network]] models that drive the widely-used [[large language model]]s for example, which are notorious for their inability to accurately evaluate their own knowledge states, resulting in significant _confabulation_. In addition, most experts do not think these models are conscious. Notably, these models are based exclusively on [[feedforward connectivity]], consistent with the idea that bidirectional connectivity is essential for consciousness, and the functional benefits associated with it.

    The central role of bidirectional connectivity in Axon represents one of the most important points of divergence relative to the vast majority of existing neural network models, and testing the functional importance of this property is a major overarching goal of this research.

These basic neural mechanisms are sufficient to learn well-established functions of the [[posterior neocortex]], including spatially invariant [[object recognition]]. However, another critical hypothesis for the Axon framework is that successful learning and performance in complex real-world environments requires significant additional neural circuits and systems to support _goal-driven_ learning and performance across longer time scales in a coordinated manner, which is captured in the [[Rubicon]] framework.

The goal-driven system must be sensitive to the overall costs and benefits of actions, in order to commit time and effort to a consistent plan of behavior to accomplish a given goal. Thus, the brain systems in the Rubicon model include various subcortical areas involved in representing reward, punishment, and effort. Overall, this goal-driven learning provides critical error signals that shape learning throughout the network, going beyond what is possible with basic predictive error-driven learning.

The [[Rubicon]] framework provides an implementation of this goal-driven learning system, based on the detailed biology of the [[prefrontal cortex]], [[basal ganglia]], [[amygdala]], [[dopamine]], and the [[hippocampus]] (traditionally known as the [[limbic system]]), among others.

## Axon pages

