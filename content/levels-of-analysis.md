+++
Categories = ["Axon", "Computational Cognitive Neuroscience"]
bibfile = "ccnlab.json"
+++

The goal of the [[Axon]] model is to provide a coherent theoretical framework for the overall field of [[Computational Cognitive Neuroscience]] (CCN), which involves three different **levels of analysis**:

* **Neuroscience** provides the "lowest level" grounding of the underlying biological mechanisms of the brain.

* **Cognitive Psychology** provides a higher level understanding of what the brain actually does at a functional level: how do people and animals actually think, and how can we understand the relationship between these cognitive functions and the underlying mechanisms from neuroscience?

* **Computation** provides a tool to build the bridges between neuroscience and cognition in a formal and precise way: instead of building a theory only on the basis of words and concepts, we also build it out of _computer programs_ that implement very specific mathematical and algorithmic mechanisms that attempt to describe what the neural mechanisms are doing.

These different levels of analysis in the CCN field can be mapped onto the levels of analysis articulated by David Marr ([[@Marr77]]), which are based on the _source_ domain of _computer science_, where a given computational device can be understood at three different levels:

* **Computational** is the highest level, characterizing the overall computations that are being performed, and the abstract nature of the information that is being processed. For example, a sorting algorithm can be described as accomplishing the computational outcome of ordering _items_ in a _list_ according to a given _metric_, where the italicized terms represent the critical computational-level parameters associated with this level of computation.

* **Algorithmic** is the next level down, providing a more detailed description of the underlying information processing steps required to accomplish a given computation. Given the universal nature of a [[turing machine]] computational device, this algorithmic description can be provided in any suitable _high level language_ such as Python, C, Go, or _pseudocode_ which tends to look a bit like Python.

* **Implementational** is the _hardware_ level of description, in terms of the actual bits of silicon that  actually implement the [[turing machine]], typically using a _Von Neumann_ architecture that underlies all modern computers.

It is essential to recognize that Marr's levels depend critically on the level-spanning role of the turing machine, which provides an abstract theoretical tool that bridges between these levels. Once you have this kind of abstract theoretical mechanism that accurately describes what the underlying hardware is actually doing at the implementational level, and provides the basis for your algorithmic-level description, then everything is very straightforward from a conceptual level.

However, we know for certain that _the brain is not a turing machine_ at the level of neurons, although it may end up operating one at an emergent, higher level through the interactions of multiple brain systems, and likely only in humans. Therefore, Marr's specific levels, which provide such a compelling source analogy, are not directly applicable to understanding the brain.

Instead, we need to find an alternative abstract theoretical mechanism that accurately describes what the underlying hardware of the brain is doing, and how this hardware accomplishes all of the higher-level cognitive functionality.

This is precisely what the [[Axon]] framework is intended to provide, through a highly accurate (though still abstracted) mechanistic implementation of the neural mechanisms of the [[neocortex]] and other brain areas (see the [[Rubicon]] model for details), that has a clear computational-level basis in terms of [[error-driven learning|error-driven]] [[predictive learning]] and the many functional properties of [[bidirectional connectivity]], etc. The models built from these mechanisms then provide an accurate explanation of a wide range of cognitive phenomena.


