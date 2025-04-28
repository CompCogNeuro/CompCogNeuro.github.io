+++
Categories = ["Axon", "Learning", "Activation"]
bibfile = "ccnlab.json"
+++

Bidirectional excitatory connectivity supports multiple important functional properties, including: [[constraint satisfaction]], [[attractor dynamics]], and [[top-down attention]]. At the broadest level, most theories of [[conscious awareness]] hinge on the ability to directly share information in an all-to-all manner, and either directly depend on bidirectional connectivity for this ([[@Lamme06]]) or indirectly through constructs such as a _global workspace_ or related constructs (e.g., [[@Baars02]], [[@DehaeneLauKouider17]]; [[@DehaeneKerszbergChangeux98]]; [[@SethBayne22]]).

Relative to most [[abstract neural network]] (ANN) models, [[Axon]] is unique in directly incorporating full bidirectional connectivity, and is thus in a position to capture the computational and functional benefits that this connectivity affords. However, there are also significant practical difficulties that come along with bidirectional excitatory connectivity. For example, it tends to produce complex activation dynamics that can easily lead to runaway positive feedback loops, which unfortunately also occur in the human brain in the form of epileptic seizures. A robust form of [[inhibition]] is critical for controlling these feedback loops. 

Learning is also much more difficult in the context of complex activation dynamics, and interestingly there are surprisingly impressive results from [[reservoir computing]] networks that eschew learning within bidirectionally connected networks entirely, using them instead as "reservoirs" of complex dynamical activity states from which signals can be decoded via simpler feedforward [[error-driven learning]] mechanisms.

By contrast, the form of learning in [[Axon]] depends critically on bidirectional excitatory connectivity for propagating error signals throughout the network, and can tune large, complex bidirectional networks to develop effective [[predictive learning]] representations of the environment, leveraging the principle of learning based on a [[temporal derivative]]. There is now experimental evidence consistent with this form of learning in at least one specific, widely-studied pathway involving cortical pyramidal neurons and synaptic mechanisms that exist throughout the neocortex [[JiangEtAlIP]].

From a computational cost perspective, bidirectional connectivity is very expensive because it doubles the number of synaptic connections, and requires roughly 200x iterations through the network to process a single input. This significantly limits the ability to scale up the models, which has been the primary driver of impressive computational performance in current feedforward ANN models. Nevertheless, as parallel compute hardware continues to improve, this limitation will hopefully be overcome (and the current version of [[Axon]] runs efficiently on any GPU, using WebGPU so it works through the browser too). For the time being, the models do focus more on capturing the principles rather than the kinds of raw performance improvements that come with scaling (see [[bias-variance tradeoff]] for more discussion).

## Biology of bidirectional excitatory connections

See [[neocortex#biology of the neocortex]] for more details on the biology of bidirectional excitatory connections in the neocortex.

The basic pattern of connectivity involves **bottom-up** or **feedforward** connections from more peripheral, sensory areas into higher-level more abstract areas. These pathways support the key process of [[categorization]]. The **top-down** or **feedback** connections from higher layers then provide the bidirectional aspect of the connectivity.

Intuitively, these top-down connections support the ability to resolve ambiguous inputs by bringing higher-level knowledge to bear on lower-level processing stages (e.g., [[@AngelucciBressloff06]]; [[@BarKassamGhumanEtAl06]]; [[@OReillyWyatteHerdEtAl13]]). For example, if you are trying to search for a companion in a big crowd of people (e.g., at a sporting event or shopping mall), you can maintain an image of what you are looking for (e.g., a red jacket), which helps to boost the relevant processing in lower-level stages. 

Thus, while the feedforward flow of excitation through multiple layers of the neocortex can make us intelligent by developing higher-level abstractions, the feedback flow of excitation in the opposite direction helps make us more _robust_, _flexible_, and _adaptive_. Without this feedback pathway, the system can only respond on the basis of whatever happens to drive the system most strongly in the feedforward, bottom-up flow of information. But often our first impression is wrong, or at least incomplete. 

There are also **lateral** excitatory connections which interconnect neurons at the same level of processing, and help provide mutual support for consistent patterns of activity. This is illustrated in the [[necker cube sim]] which shows how different visual features can support each other through lateral connections to drive a coherent overall interpretation of an otherwise ambiguous visual input. This is a particular example of the broader category of [[attractor dynamics]], which provides a more abstract, high-level characterization of the computational function of bidirectional connectivity.

## Simulations

* [[faces sim]] (Part II) demonstrates how top-down and bottom-up processing interact to produce imagery and help resolve ambiguous inputs (partially occluded faces).

* [[necker cube sim]] demonstrates how lateral excitatory connections can produce attractor dynamics in the case of a classic ambiguous visual stimulus.

* [[cats and dogs sim]] demonstrates bottom-up and top-down dynamics in a semantic network representing different levels of information about cats and dogs.

