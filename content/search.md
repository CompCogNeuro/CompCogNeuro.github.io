+++
Categories = ["Computation", "Activation", "Learning", "Axon", "Bidrectional connectivity"]
bibfile = "ccnlab.json"
+++

**Search** is perhaps the single most general unifying concept in all of computation.

* _Problem solving_ can be defined as search through _problem space_, as in many classical symbolic [[artificial intelligence]] (AI) systems defined it ([[@NewellSimon72]]; [[Newell91]]).

* _Planning_ is search through _action space_ to accomplish a desired outcome, which is the focus of many approaches in [[reinforcement learning#model-based RL]].

* _Learning_ is search through _representation space_, to find the best [[linear algebra|basis]] for representing inputs, that supports the desired computational processes and behavioral outputs.

* _Evolution_ is search through _phenotype space_, and the widely-used [[genetic algorithm]]s provide an efficient way of searching these spaces.

This is not an empty tautology, because it provides a unified understanding about a central problem faced in all of these domains:

> Each of these spaces is plauged by the [[curse of dimensionality]], and effective solutions for real-world, large-scale problems must somehow tackle the exponential explosion problem.

## Serial vs. parallel tradeoffs

The complementary advantages and disadvantages of serial vs. parallel solutions to search apply across domains, and effective solutions must find an effective balance of each.

* A parallel search algorithm fights the curse of dimensionality by operating on dimensions in parallel, and is thus an essential element of any scalable solution (_fight dimensionality with dimensionality_).

* However, serial search algorithms are much more _flexible_, because serial processes allow arbitrary _combinations_ of dimensions to be processed for each case, whereas parallel algorithms require fixed combinations to be precomputed in parallel in an earlier preprocessing step. This essential advantage of serial processing is why the [[Turing machine]] is a universal computational device, whereas parallel computation is not.

For example, evolution in the natural world involves massively parallel search across "replicated" organisms, each living their lives in parallel, but the ontological development of each such organism is serial, which allows for considerable flexibility in the expression of the shared genomic programs that accumulate across individuals.

Likewise, search in the visual system employs both parallel and serial processes, which can be differentially engaged depending on the target properties. We can efficiently perform parallel search for basic visual features such as color and orientation, but increasingly complex combinations of such features require increasingly serial spatial [[attention]] to narrow the search space ([[@Treisman77]]; [[@Wolfe10]]; [[@HerdOReilly05]]).

## Learning as search

Perhaps the biggest divide between classical symbolic AI and the "new" neural-network based approaches is that neural networks enable a parallel, gradient-based learning strategy that allows learning to search through high-dimensional representational spaces, now reaching into the billions of parameters, which represents a practically infinite combinatorial space (with exponential size well beyond any human comprehension, exceeding the number of atoms on the universe by exponentially-large exponential factors!).

By contrast, the discrete symbolic nature of classical AI frameworks, where the models are essentially Turing machines with various AI-based mechanisms, is incompatible with such a parallel, gradient-based learning mechanism, and thus any attempt to learn in such models immediately collapses in the face of the curse of dimensionality.

The process of [[error-backpropagation]] using stochastic gradient descent is effectively exploring all possible directions to move in the high-dimensional parameter space in each iteration of learning, and selecting the direction of steepest descent, typically subject to a number of heuristic constraints that have significant practical benefits (e.g., the _AdaMax_ mechanism; [[@KingmaBa14]]). The _stochastic_ nature of this process derives from the use of small subsamples of the total problem space to compute these gradients, whereas using the full space ends up getting "stymied" by all the constraints imposed across this entire space.

These properties of stochastic, gradient-based search provide a useful template for any kind of high-dimensional search algorithm, even if the details might be implemented differently. For example, the _survival of the fittest_ principle of genetic algorithms provides a kind of local gradient computation, and _random mutation_ provides a source of stochastic sampling.

Interestingly, a number of people reportedly considered the use of something like error backpropagation, and dismissed it out of hand because it seemed inevitable that the algorithm would get stuck in highly suboptimal _local minima_, and thus be useless. Considerable research is now illuminating why it actually works as implausibly well as it does (e.g., [[@NakkiranKaplunBansalEtAl21]]; [[@Shwartz-ZivTishby17]]; [[@VidalBrunaGiryesEtAl17]]).

## Constraint satisfaction as search

The concept of [[constraint satisfaction]] provides another high-level unified way of understanding a wide range of computational processes, and this generality is not accidental from the present search-based perspective, because constraint satisfaction is essentially a form of parallel search. In this case, it is a search over possible states that solve or optimize a set of constraints. Many problems can be formulated in this way, including planning and problem solving, so constraint satisfaction provides a more general way of understanding the essential computations in these domains.

Here too the stochastic, gradient-based approach to search is the only one capable of handling high-dimensional state spaces, and it also is generally the most effective approach, and does not suffer too much from local minima problems ([[@HoosTsang06]]). One class of such algorithms is in fact a neural network, going back to the pioneering approach of [[@HopfieldTank85]], who showed that iterative updating of a bidirectionally-connected Hopfield network could provide good solutions to the travelling salesman problem, which is a classic example in the constraint satisfaction domain.

In the [[Axon]] framework, the use of the [[GeneRec]] approximation to [[error backpropagation]] via [[bidirectional connectivity]] and [[temporal derivative]] learning rules results in the ability to perform constraint satisfaction by effectively performing [[error backpropagation#backpropagation to activations]]. Thus, Axon networks are continuously generating [[optimized representations]] that reflect the results of this constraint-satisfaction process, and therefore simultaneously performing search through representation space dynamically in activation space, and over the course of learning through synaptic weight adjustments.

## Emergent serial processing

The most effective balance of parallel and serial processing is with a foundation of massively-parallel,  distributed processing that tackles the high-dimensional nature of the real-world with high-dimensional parallel processing, and in so doing, produces a much lower-dimensional high-level abstract summary of the situation. This low-dimensional abstract space should then be amenable to more flexible serial processing along the lines of a [[Turing machine]] architecture, which can operate at an emergent level on top of the underlying parallel hardware.

We believe that this overall configuration of parallel and serial processing is essential for intelligent behavior, and it provides a compelling description of many aspects of human cognition. In terms of the longstanding debate between symbolic vs. "subsymbolic" approaches to AI, this represents a synthesis of the two, with symbolic processing emerging out of fundamentally subsymbolic, neural-network hardware.

Interestingly, [[large language models]] have this same overall configuration as well, because they are extensively trained to predict the behavior of computer programs, and thus essentially learn to behave like a Turing machine ([[@YangCampbellHuangEtAl25]]). Indeed, research has shown that excluding the computer programming aspects of the standard training corpus significantly impacts the overall cognitive flexibility of the resulting system TODO CITES!!.


