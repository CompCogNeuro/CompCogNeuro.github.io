+++
Categories = ["Computation", "Learning"]
bibfile = "ccnlab.json"
+++

A **genetic algorithm** (GA) captures key elements of the process of evolution to perform [[search]] through high-dimensional spaces. In the space of adaptive learning algorithms, gradient-based [[error-driven learning]] is the best, trial-and-error [[reinforcement learning]] is the second best, and genetic algorithms are the third best, in terms of speed of convergence.

This hierarchy is also defined in terms of the generality of the mechanism: error-driven learning has the strongest requirements (needing target output patterns), while reinforcement learning only requires a scalar reward signal, and genetic algorithms likewise only require a definition of _fitness_, without the need for the ability to compute a gradient. Thus, it is the most general, robust form of learning. But like evolution, it can take a while.

The essential ingredients of a GA include:

* A _genotype_, which is a compact modifiable plan for how to build a phenotype (organism, entity, etc).
* A _fitness function_ that evaluates the performance of individual phenotypes.

The process of evolution thus proceeds by randomly generating genotypes, creating a population of phenotypes from them, and then evaluating the fitness of the individuals in the current population. The top-performing subset is then "mated" to form a new set of genotypes by mixing elements among them, and the process iterates over multiple generations. There are many variations in each of these different components, as summarized in the [wikipedia page](https://en.wikipedia.org/wiki/Genetic_algorithm).

From a pragmatic perspective in relation to the [[Axon]] and [[Rubicon]] models, GAs are typically too computationally expensive to be of significant value. Simulating a single brain-level model of any complexity is already a very computationally expensive process, so iterating over large populations of such models for many generations is not an efficient way to proceed.

Instead, we take the approach of [[computational-cognitive-neuroscience#reverse engineering the brain]] through [[neuroscience]] informed by research on [[cognition]] and [[computation]] to conduct a more directed, hypothesis-driven search for the key "discoveries" that millions of years of evolution on our planet has produced.

Nevertheless, in specific narrower contexts, GA algorithms are the only way to proceed.

