+++
Categories = ["Computation", "Activation", "Axon"]
bibfile = "ccnlab.json"
+++

The use of **optimized representations** means that the system computes dynamically updated novel representations based on [[constraint satisfaction]] and/or [[error backpropagation#backpropagation to activations]] as the basis for processing inputs and driving cognition and behavior. The use of [[bidirectional connectivity]] in [[Axon]] automatically generates optimized representations, where the [[GeneRec]] analysis shows that the activity states automatically compute the _error gradient_ in the activation states of neurons.

There is a tradeoff in computation expended for generating these optimized representations on each trial, versus optimizing the number of trials for learning, using a much faster mode of per-trial processing, such as the standard single-iteration feedforward pass in a standard [[abstract neural network]]. Given the biologically-realistic approach taken in the [[Axon]] model, which uses bidirectional connectivity to perform [[error-driven learning]], these models are effectively committed to at least around 200 iterations worth of constraint-satisfaction processing per trial.

Therefore, the central question here is whether there are significant functional benefits for using optimized representations in this way? Given that the mammalian brain is likewise committed to this type of processing, it is reasonable to conclude that this is likely, but further research is needed to more definitively answer this question.

Further motivation for the potential benefits of using optimized representations is provided in the discussion on [[search]] as a universal computational process, and the ability of parallel, iterative constraint-satisfaction processing to efficiently search through high-dimensional representational space to find particularly useful ways of encoding the current situation. Some potential computational benefits of doing so are enumerated in the discussion of [[reinforcement learning#model-based]] reinforcement learning, including activating plan and goal representations that satisfy external and internal constraints, and shaping lower-level perceptual and motor processing to optimize pursuit of these goals.

In this context, there is a broader continuum in terms of the amount of time and effort spent evaluating decisions and actions at any given point in time, which can be characterized in terms of the spectrum between _habitual_ vs. _controlled_ processing. Thus, any given set of representations can be subject to more or less optimization, depending on the relative importance and risk involved, for example.

Recent versions of [[large language models]] (LLMs) are now being optimized to perform more of this "reasoning" processing, versus just going with the first feedforward pass of output that they generate. This generally produces more "intelligent" and effective processing of more complex and challenging problems, consistent with the idea that an overall bias toward using optimized representations will be generally beneficial. See [[@OswaldNiklassonRandazzoEtAl23]] for a consistent analysis of the way that "in context" learning operates within an LLM to produce gradient-optimized representations.

From the perspective of subjective human experience, most theories of [[conscious awareness]] emphasize the importance of bidirectional (recurrent) dynamics to integrate and coordinate multiple brain areas on the current focus of conscious processing. Thus, our subjective conscious experience may reflect this optimized representational processing that we automatically perform on a continuous basis, with modulations in the amount of time and effort spent thinking about different topics.


