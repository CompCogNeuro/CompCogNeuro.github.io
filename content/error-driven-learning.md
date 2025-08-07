+++
Name = "Error-driven learning"
Categories = ["Learning", "Computation"]
bibfile = "ccnlab.json"
+++

**Error-driven learning** is a powerful form of learning that drives changes in synaptic weights to reduce _errors_. In [[Axon]], these errors are generally _prediction errors_: the differences between a prediction and what actually happens (see [[predictive learning]]). Prediction-error-driven learning is what drives learning in the [[large language models]] (LLMs) that power ChatGPT and related models.

Error-driven learning requires a mechanism for these error signals to drive synaptic changes. In LLMs and other current [[abstract neural network]] models, [[error backpropagation]] is used to drive synaptic learning. However, this mechanism is not directly compatible with known neurobiological mechanisms. Thus, Axon uses [[temporal derivative]]s, which can be computed locally using biologically based mechanisms in the context of [[bidirectional connectivity]], to drive error-driven learning using the [[kinase algorithm]].

Because error-driven learning is mathematically derived to minimize performance errors on any given task, it provides a very general and flexible form of learning. By contrast, [[Hebbian learning]], which is more obviously supported by the known [[synaptic plasticity]] mechanisms, is mathematically related to [[principal components analysis]], and generally learns to encode the strong correlational structure present in its inputs. Empirically, it does not learn to solve arbitrary tasks, and is thus not a plausible basis for a general-purpose learning mechanism in the [[neocortex]].

See [[combinatorial vs conjunctive]] for an analysis of the tradeoffs at play in error-driven learning, that help clarify the relationship to Hebbian learning.


