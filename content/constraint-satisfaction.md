+++
Categories = ["Computation", "Activation", "Axon"]
bibfile = "ccnlab.json"
+++

**Constraint satisfaction** is one of the most important concepts for understanding the power of [[bidirectional connectivity]] in the [[neocortex]], and is central to the promise of the [[Axon]] approach, which is one of the few neural network models to incorporate extensive bidirectional connectivity in a way that tames its wild side while retaining its many benefits.

The constraint satisfaction problem (CSP) is defined as finding a set of values for N variables that satisfy a set of constraints defined over those variables ([[@Tsang14]]). A classic example is the N-queens problem, where you need to place N chess queens on a board such that no two queens should threaten each other. Another such problem is the _travelling salesman problem_ (TSP), analyzed by [[@^HopfieldTank85]] using a bidirectionally-connected Hopfield network, which involves finding a minimum distance route between _N_ cities.

Thus, the CSP is essentially a problem of [[search]]ing over all possible states to find the one(s) that best fit the set of constraints imposed. As the number of states increases, the number of possible states explodes exponentially due to the [[curse of dimensionality]], 

A bidirectionally-connected neural network can implement this search process in a highly efficient manner, by integrating all of the constraints _in parallel_ and essentially performing a _stochastic gradient descent_ process over possible solution states. This is essentially the same strategy used in [[error-backpropagation]] learning, to search over the high-dimensional space of possible representations, as explained in [[search]].

TODO: old:

In the context of the [[curse of dimensionality]] problem, bidirectional connectivity enables the network to rapidly converge through parallel processing on a set of representations that satisfy constraints communicated throughout the network, where each active neuron contributes its own constraint, and is in turn subject to the constraints from the other neurons. Given the _small world_ nature of neural connectivity, each neuron is effectively only a few synapses away from any other neuron, so effectively any constraint anywhere can be felt anywhere else in the network, in principle.

This parallel, distributed constraint satisfaction process is effectively performing a massive search over the entire space of possible ways of representing the current combination of external and internal constraints, in a way that would otherwise be impossible in any kind of more serial process, due to the combinatorial explosion of possible such representations.

Purely feedforward networks do not adapt their representations dynamically as they process the current set of inputs, and instead just generate a representation in one sweep, based on the learned weights. Thus, they are not optimizing these representations to find the most _satisfying_ way of interpreting the current situation. By contrast, the iterative back-and-forth interactions among bidirectionally-connected neurons ends up optimizing the active representation, which then provides the basis for subsequent learning.

Mathematically, it is effectively a process of [[error backpropagation#backpropagation to activations]], where the error gradient that is otherwise used to drive learning is actually used to drive the activation state of the network, and then this activation state, via the [[temporal derivative]] learning principle, drives learning according to this gradient.

