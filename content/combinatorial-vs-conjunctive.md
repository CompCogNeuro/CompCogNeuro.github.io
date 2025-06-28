+++
Categories = ["Axon", "Learning", "Computation"]
bibfile = "ccnlab.json"
+++

There is a direct tension between the benefit of **combinatorial** codes for [[generalization]], versus **conjunctive** codes for arbitrary task learning. Understanding the nature of this tension and the tradeoffs involved is key for understanding the relative successes and failures of different network architectures.

We can understand this tension in the context of a classic "3 layer perceptron" [[abstract neural network]], with `Input` $\rightarrow$ `Hidden` $\rightarrow$ `Output` layers, which needs to accomplish two conflicting goals in order to both learn arbitrary problems and generalize successfully to novel inputs:

1. Capture the _similarity structure_ of the Inputs, so that similar inputs generally map to similar outputs (i.e., _similarity-based generalization_), which is critical for systematic behavior in novel situations. This is best accomplished by a [[distributed representations|distributed]], _combinatorial_ code, where the pattern overlap of the Hidden activity mirrors that of the Input, and the individual Hidden units _"carve the input at its joints"_ so that novel combinations of Hidden activity systematically and sensibly represent novel states of the Input.

    This objective can be measured by the [[information theory#mutual information]] between the Input and Hidden layers ([[@Shwartz-ZivTishby17]]). The most efficient code from an information-theoretic sense is achieved when the individual Hidden units extract something like the [[principal components analysis|principal components]] (PCA) of variance across the inputs, such that each hidden unit is uncorrelated with the others, and thus contributes unique independent information.

2. Learn whatever _arbitrary_ task is required (to survive), which is the goal of [[error-driven learning]], and [[error backpropagation]] specifically. This can be measured in terms of the mutual information from the Hidden to the Output.

Thus, the Hidden layer has two different masters, each of which pull it in different directions, and its job is to synthesize a mapping between them that not only gets the right answers (Outputs), but does so with a _deeper understanding_ of the Input space that allows it to systematically generalize to novel situations. Often this magic mapping cannot be done in a single step, so multiple Hidden layers are required to develop appropriate abstractions of the input (as discussed in [[categorization]]) (i.e., _deep networks_).

The "no free lunch" [[bias-variance-tradeoff]] analysis ([[@GemanGeman84]] and [[@VapnikChervonenkis71]]) shows that a _lookup table_ is the most flexible arbitrary function learning mechanism, placing no constraints (biases) at all on the shape of the function. But it requires a maximum number of parameters, namely an entire table entry per each data point. In neural terms, this is a _one-hot_ or _localist_, _conjunctive_ code, where a single, distinct Hidden unit is activated for each distinct Input pattern (i.e., that unit is exquisitely sensitive to the full _conjunction_ of input activity). This code generalizes poorly, because it doesn't tell you what to do if something is not already in the table (i.e., a newly activated Hidden unit has no learned weights to the Output).

With limited amounts of training data, generalization in networks benefits from _biases_ that enhance the mutual information between Input and Hidden, and promote combinatorial codes. For example, adding an explicit auto-encoder term to the objective function directly forces the Hidden layer to capture the Input (and approximates the results of PCA; [[@BaldiHornik89]]). PCA can also be accomplished using [[Hebbian learning]] in addition to error-driven learning; this was the key point of the [[Leabra]] algorithm. There are many other examples of these bias / variance tradeoff constraints.

As the amount of training data increases, the benefits of these biases progressively diminish because the data itself starts to cover the entire space anyway. Intuitively, it fills out the lookup table sufficiently densely that simple interpolation between nearest-neighbors in the table becomes effective. This is the  [bitter lesson](http://www.incompleteideas.net/IncIdeas/BitterLesson.html) from Rich Sutton, and the reason why "big data" is so successful.

The [[transformer]] architecture and the scaling properties of [[large language models]] provide a clear demonstration: transformers generalize _terribly_ with small amounts of data, but have the lookup-table like capacity to learn successfully with huge amounts of data. The key-value matching aspect of transformers is a soft version of a one-hot conjunctive code, making these networks very lookup-table like, while also supporting similarity-based generalization.

Although the transformer was billed as a mechanism for [[attention]], in fact it is much more of an [[episodic memory]]-like system capturing the function of the [[hippocampus]], but instead of being a separate brain system, this lookup-table-like functionality is embedded at each level of processing within the transformer. Consistent with this analysis, LLMs have been shown to be capable of directly memorizing long passages of their training corpus (e.g., [[@HuangYangPotts24]]).

## The binding problem

There is also a time domain aspect to the combinatorial vs. conjunctive tradeoff, in addition to the above learning and generalization aspect, in terms of the [[binding problem]]: combinatorial codes directly lead to binding errors, which are minimized by conjunctive codes. For example, if you have separate Hidden units representing color vs. shape, and you have Red, Green, Triangle, and Square all activated, you don't know whether you're seeing a Red Triangle and a Green Square or the other combination.

The binding problem has significant implications for what can be done in _parallel_ vs. _serial_ (see [[search]] for a detailed discussion). Anne Treisman showed that visual search on separably represented features (color vs. shape) can happen efficiently in parallel, but conjunctions require slower serial processing, constrained by top-down spatial attention to different regions of the visual space, to reduce binding error ([[@Treisman77]]; [[@TreismanGelade80]]).

If the brain could represent these conjunctions in distinct neural populations, then parallel search would be possible, but the [[curse of dimensionality]] prevents this, as the number of neurons required would grow exponentially.

Similar conclusions were reached in an analysis of visual binding errors made by people and LLM models ([[@CampbellRaneGiallanzaEtAl24]]), and in the case of multitasking and cognitive control ([[@MusslickCohen21]]; [[@MusslickSaxeHoskinEtAl20]]). In this latter case, the [[prefrontal cortex]] is thought to provide a top-down [[attention]]al focus on the features relevant for one task vs. another, imposing a strong limit on what can be done in parallel.

## Catastrophic interference 

Combinatorial, distributed representations also suffer from [[catastrophic interference]] because learning on any given input affects a large number of weights in such a distributed system, whereas a one-hot conjunctive code minimizes interference (in the limit, an entirely new unit could be allocated for each new input). 

The _complementary learning systems (CLS)_ framework for understanding the roles of the neocortex vs. [[hippocampus]] is based on this tradeoff, with the idea that tradeoffs can be minimized by having _two separate systems_ that optimize each end of the continuum. Cortex uses distributed, combinatorial codes to support similarity-based generalization, while the hippocampus uses sparse, conjunctive codes to minimize interference and support lookup-table like learning of specific episodes.

By contrast, transformers (LLMs) integrate this conjunctive coding ability throughout the network, and therefore do not have the equivalent of a more purely combinatorially-biased posterior neocortex. Nevertheless, with sufficiently huge amounts of training data and very deep layered architectures, they are able to develop sufficient abstractions and dynamics that unfold over layers and across many sequential iterations (see [[optimized-representations]] and [[@OswaldNiklassonRandazzoEtAl23]]), that end up supporting their impressive levels of generalization performance.

Reinforcing this point, the key-value softmax mechanism that is integral to transformers has recently been re-branded as a _modern Hopfield network_, and used to simulate the episodic memory functions of the hippocampus ([[@RamsauerSchaflLehnerEtAl21]]; [[@KrotovHopfield21]])

## The neocortex is more combinatorial than conjunctive

The above phenomena, along with a huge amount of direct neural recording and other data, strongly support the idea that the neocortex of the mammalian (including human) brain is strongly weighted toward the combinatorial, distributed end of the spectrum, which is consistent with the idea that real-world survival depends critically on dealing with the _small data, strongly biased_ domain.

In other words, it is much more important to be approximately correct based on very little data, than it is to be perfectly correct with large amounts of data. This conclusion is also consistent with a large literature on the use of simple heuristics instead of more accurate but expensive processing of statistical data ([[@KahnemanTversky84]).

This conclusion thus suggests that LLMs do not accurately capture the way that the human brain learns, because they are strongly weighted toward the conjunctive end of the spectrum. They are good at using huge amounts of data to develop accurate models of what the human brain has accomplished, but they are not a good model of how the human brain actually came up with all these ideas in the first place, which requires the kind of "deep" conceptual understanding of the world that only the combinatorially-biased neocortex can provide.

In summary, there is a corollary to the "bitter lesson": the _boring lesson_:

> Any system that relies largely on big data instead of strong biases is fundamentally boring, because it does not provide a generative solution to where all this data came from in the first place.

Humans are strongly biased and motivated to seek low-dimensional representations of the world, to satisfy a desire to understand how it all works, and make it work for us.

