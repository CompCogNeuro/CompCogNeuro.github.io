+++
Categories = ["Computation"]
bibfile = "ccnlab.json"
+++

Everything breaks down in high dimensional spaces, unless it can leverage the very power of high-dimensionality to its own advantage: fight dimensionality _with_ dimensionality.

This means that it needs to operate in _parallel_. Any sequential process dies from the curse of dimensionality.

Parallel processes are very special: much smaller space of viable algorithms.

Gradient descent is the paradigmatic example, and what allows huge network models to be constructed.

The magic ingredient in Axon is the ability to perform gradient descent in activation space, everywhere in the network. This is what bidirectional connections support.

This is really the key missing point from everything I've been saying!!  Backprop to activations is a thing, and sometimes it is used, but you need a more robust framework to support full backprop to activations everywhere, all the time.

This allows you to _optimize the representations_ in real time, not just the weights over the outer loop of training.


