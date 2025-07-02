+++
Categories = ["Learning", "Computation"]
bibfile = "ccnlab.json"
+++

**Large language models** (LLMs) are likely now familiar to most readers, as they are the technology behind the widely-used ChatGPT and related products, used by millions of people on a daily basis. After years of hyped expectations, these products represent the point at which [[artificial intelligence]] (AI) went fully mainstream, inevitably stimulating endless discussions about the end of human civilization etc, while demonstrably having at least the capability of changing the nature of work and education for everyone.

The underlying computational frameork behind LLMs is the [[transformer]], which combines a very standard [[abstract neural network]] component with a novel _self-attention_ mechanism that allows the system to dynamically modulate the activation strength of other elements of the input stream, based on learned transformations of those inputs ([[@VaswaniShazeerParmarEtAl17]]). The entire network is trained with end-to-end [[error backpropagation]], and it processes even relatively large chunks of text in _parallel_. This allows the backpropagation mechanism's _credit assignment_ process to develop all manner of complex circuits across many stacked layers, to transform the raw inputs into the next predicted word in this chunk of text. That word is then added to the input, and the process repeats, generating long and impressively-cogent text passages based on prompts.

As discussed in detail on the [[transformer]] page, the system captures the cognitive processes of systematic [[generalization]], semantic memory, and [[episodic memory]], using mechanisms that differ significantly from the biologically-based ones in the [[Axon]] and [[Rubicon]] framework.

Despite these mechanistic differences, the LLM transformer models provide the only working example of how these critical cognitive functions work in an artificial system outside the human brain. Extensive research is slowly peeling back the opaque "black box" cover of inscrutability that has made it difficult to leverage this amazing tool to better understand how the human brain accomplishes similar feats ([[@YangCampbellHuangEtAl25]]; [[@McGrathRussinPavlickEtAl24]]; [[@ElhageNandaOlssonEtAl21]]). Initial general insights are discussed in the [[transformer]] page, and other relevant insights are discussed throughout the text.

Some examples of key insights are:

* [[Combinatorial vs conjunctive]] analyzes LLMs in the context of general computational principles of learning, arguing that they often learn to function more like a conjunctive lookup table, and only when shaped by very large corpora do they exhibit more systematic [[generalization]].

* LLMs can learn to function like [[self-programmable]] [[Turing machine]]s, giving them the ability to follow a sequence of processing steps, assembled to drive behavior toward a desired gol.


