+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.json"
+++

**Error-driven learning** is a powerful form of learning that drives changes in synaptic weights to reduce _errors_. In [[axon]], these errors are generally _prediction errors_: the differences between a prediction and what actually happens (see [[predictive learning]]). Prediction-error-driven learning is what drives learning in the [[large language models]] (LLMs) that power Chat GPT and related models.

Error-driven learning requires a mechanism for these error signals to drive synaptic changes. In LLMs and other current "AI" / "ML" (machine learning) models, [[error backpropagation]] is used to drive synaptic learning. However, this mechanism is not directly compatible with known neurobiological mechanisms. Thus, we use [[temporal derivative|temporal derivatives]], which are compatible with neurobiology to drive error-driven learning in [[axon]], via the [[kinase algorithm]].


