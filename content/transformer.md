+++
Categories = ["Learning", "Activation", "Computation"]
bibfile = "ccnlab.json"
+++

The **transformer** architecture, developed by researchers at Google ([[@VaswaniShazeerParmarEtAl17]]), powers [[large language models]] (LLMs) and contains two distinct mechanisms, _self-attention_ and a high-dimensional _feedforward network_ (FFN) (i.e., standard [[error-backpropagation]]). The attention component gets all the [[attention]], so to speak, but apparently the FFN component is responsible for the considerable memorization abilities of this architecture, as evident in LLMs ([[@NandaRajamanoharanKramarEtAl23]]). The transformer also depends on a number of other widely-adopted [[abstract neural network]] mechanisms including critically the ResNet (residual network) [[@HeZhangRenEtAl15]] and normalization mechanisms.

Despite the fact that the transformer mechanisms have no plausible direct connections to neuroscience, it is of great importance to the [[Axon]] project, and any attempt to understand [[computational cognitive neuroscience]] more broadly, because it provides the only known working example of an artificial system based on neuron-like processing mechanisms that accomplishes significant levels of systematic, generative [[generalization]] at a human-like level ([[@YangCampbellHuangEtAl25]]; [[@McGrathRussinPavlickEtAl24]]; [[@McClellandHillRudolphEtAl20]]). After describing the basic mechanisms, we consider what can be learned by understanding how these mechanisms accomplish this impressive feat.

At an overview level, it emerges through a combination of structure-sensitive processing enabled by self-attention and the ResNet architecture, and a powerful memory system that combines features of both the [[neocortex]] and [[hippocampus]] in the feedforward network. Thus, the transform is more than just the sum of its parts: these components work together synergistically to support complex emergent dynamics ([[@ElhageNandaOlssonEtAl21]]).

## Self-attention

{id="figure_transformer-attn" style="height:20em"}
![Illustration of the attentional modulation performed in a transformer. Each input token (word) serves as a Query that is multiplied (dot product) with the Keys for each other token, with the normalized (softmax) magnitude multiplying the Value of each token, which is the result that goes on to the next stage. Instead of literally using the token representation, each attention head learns to transform the token to shape the Query, Key, and Value representations so that they end up being useful (via error backpropagation) to solving the problem. Multiple such heads are applied at each stage.](media/fig_transformer_attention.png)

The attentional mechanism is illustrated in [[#figure_transformer-attn]], which is computed for each input token serving as a _Query_ that is compared (via a dot product, i.e., [[linear algebra|matrix multiplication]]) with the _Key_ values for each of the other words. The representation of the token as a Query and Key is determined by separate trained weights (via error backpropagation through the entire network), which allows the attention to be somewhat flexible in terms of what it extracts from each token, and how it tries to match against the other tokens as Keys.

The limited, focal nature of attention is captured by the application of a _SoftMax_ function that is the normalized exponential function of the Query-Key dot products (which are single scalar values):

{id="eq_attn" title="Attention"}
$$
\rm{Attention}(Q, K) = \rm{softmax} \left( \frac{Q K^T}{\sqrt{d_k}} \right)
$$

{id="eq_softmax" title="SoftMax"}
$$
\rm{softmax}(\bf{X}) = \left[ \frac{e^{x_i}}{\sum_j e^{x_j}}, ... \right]_i
$$

The resulting normalized scalar attention value for each _Key_ (represented by the downward arrows in [[#figure_transformer-attn]], computed as a sum of the contributions from each Query) then multiplies the _Value_ vector for that token, which again is transformed by linear weights, and it is this transformed, attention-modulated value that is the output result of the attention mechanism (shown by the big arrows flowing from input to output in the figure).

Intuitively, this attentional mechanism allows each input word to select the other words that are most relevant to it, so in the example in the figure, the input word "orange" modifies "cat" so it gives a strong attentional weight to cat.

Because of the exponential SoftMax normalization, the attentional mechanism typically focuses on relatively few other words, so multiple such attentional _heads_ are used in parallel to allow multiple attentional foci to be processed in the same feedforward sweep. Because the model is trained with standard [[error backpropagation]], it is critical for the credit assignment process to have everything of relevance present all at once in parallel -- otherwise a much more complex backpropagation through time process would be required to accomplish [[temporal credit assignment]]. 

## Feedforward network: high-dimensional pattern separator

As noted above, it is actually the FFN component, which is just a standard feedforward backpropagation network, that accounts for much of the memorization ability of a transformer ([[@HuangYangPotts24]]), which is as critical to their functionality as the self-attention mechanism. This is what enables these models to be able to digest such a huge amount of information and generate any kind of prediction at an above-chance level over databases with billions of words.

The essential feature of the FFN that supports this memorization capacity is the size of the hidden layer, which is 4x larger than the size of the token input vectors, and accounts for roughly 70% of the total LLM learnable parameters in a typical GPT LLM model ([[@Aizi23]]). The FFN projects the attentionally-weighted token representations into this large hidden layer that allows activity patterns to become relatively more separated from each other (i.e., _pattern separation_). This is the same computational trick that is used in the _support vector machine_ (SVM; [[@CortesVapnik95]]), and in the brain in the [[hippocampus]] and [[cerebellum]].

After performing this pattern separation function, the representations are then compressed back down into the same dimensionality as the inputs, not as an autoencoder (i.e., using the same representation as the input) but rather creating new transformations that can build on the [[combinatorial vs conjunctive|conjunctive]] features discovered by the large hidden layer. This preservation of the same dimensionality throughout the entire network is critical for the ResNet property as discussed next.

## The ResNet residual stream

The ResNet property of transformers is critical for essentially providing a bus-like pathway that the attention heads can partition and manipulate as the information flows forward through the network ([[@ElhageNandaOlssonEtAl21]]). Specifically, each level of the network passes its outputs directly through to the final output stage of the entire network, and additional levels just add their contributions to this _residual stream_ of information.

Each higher layer only gets information from the level just below it, so that contextualizes the processing performed by that layer, but the ability to directly pass lower-level information straight through to the output is essential not only for making the error backpropagation work effectively (otherwise the error gradients would be massively diluted by the time they filtered through all the higher layers), but also for enabling higher layers to serve as "dynamic editors" of an accumulating flow of transformed information extracted from the inputs.

These higher layers therefore do not need to replicate all of the information _content_ extracted by lower layers, and instead can inject more _structural_ information that captures the relationships present in the input at various levels of abstraction, which is essential for their systematic, generative behavior as discussed in [[generalization]] ([[@YangCampbellHuangEtAl25]]; [[@WebbFranklandAltabaaEtAl24]]; [[@OReillyRanganathRussin22]]).

## Cognitive neuroscience implications

Despite some dismissive claims to the contrary, transformers do not get their impressive capabilities through pure rote memorization of the huge corpora upon which they are trained, even though they do tend to exhibit rather remarkable memorization abilities. Instead, at a cognitive level, the "secret sauce" behind their success is the ability to integrate this vast semantic and [[episodic memory]] capacity with abstract, structure-sensitive [[generalization]] abilities.

### Structure-based systematicity and optimized representations

At a cognitive level, this multi-head attention likely captures some of the dynamics that result from both parallel [[constraint satisfaction]] and serial, sequential attention in humans reading text ([[@McClellandHillRudolphEtAl20]]). For example, when there are mutually-compatible words that form a larger compound expression (e.g., "piece of cake"), then the attentional mechanism can produce mutual excitation and focus on those words as a unit, helping them to be processed as such. There is good evidence for these kinds of mechanisms operating in LLMS across multiple levels of abstraction, as represented in existing linguistic theories ([[@TenneyDasPavlick19]]; [[@ManningClarkHewittEtAl20]]; [[@WarstadtBowman22]]; [[@ChenShwartz-ZivChoEtAl25]]; [[@McGrathRussinPavlickEtAl24]]).

There is an extensive literature on the sequential processes involved in [[language]] comprehension, which naturally produce strong attention-like effects as one proceeds through the words in a sentence and uses the accumulated _context_ to interpret the subsequent words. When those subsequent words do not conform to these expectations, as in a _garden path_ sentence, then an active re-interpretation is often required ([[@FerreiraHenderson91]]; [[@Fujita21]]). Although transformers lack active working memory maintenance, they have increasingly large context input sizes, which effectively provide an "in context" active memory system that can be used to contextualize subsequent processing ([[@OlssonElhageNandaEtAl22]]). 

Overall, the attentional mechanism results in a form of [[optimized representations]] in the transformer, relative to a standard feedforward network, where these kinds of strong interactions among different streams of processing do not occur within a given layer, and can only accumulate through the accumulation of transformations across layers. Furthermore, even this layer-wise accumulation of interactions is often limited by the use of ReLU activation functions that are largely linear and thus tend to result in linear, additive effects. By contrast, attention in the transformer is more nonlinear due to the SoftMax normalization dynamic.

In comparison with the full [[bidirectional connectivity]] and consequent constraint satisfaction dynamics in [[Axon]], certainly the transformer attentional mechanism is a strong step in the right direction toward capturing these dynamics. However, the transformer remains a fundamentally feedforward architecture, so it cannot actually have higher-level representations directly feed back and modulate the activity of earlier layers. Instead, these higher-layers perform the dynamic editing of the outputs of lower layers in the residual stream.

Furthermore, the very large number of feedforward layers can be learning to effectively unroll the dynamics that would otherwise be happening in bidirectional networks (see [[bidirectional connectivity#feedforward unrolling]]). Nevertheless, unrolling the entire neocortical network in this way would be computationally intractable, so the transformer can only be capturing a much more limited aspect of what genuine bidirectional connectivity can accomplish in the neocortex.

### Semantic and episodic memory

Although the transformer does not have any kind of structural distinction corresponding to the anatomical and functional differences between the [[neocortex]] (for semantic memory) and [[hippocampus]] (for episodic memory), the high-dimensional nature of the feedforward network allows these models to encode both highly specific memories of specific texts and more generalized semantic knowledge ([[@NandaRajamanoharanKramarEtAl23]]). As discussed in [[combinatorial vs conjunctive]] representations, there are tradeoffs between these different types of representations, and the backpropagation mechanism can learn a useful mix of these different types of representations as needed for the [[predictive learning]] task that transformers are typically used for.

What the standard transformer definitely does lack is the ability to form _new_ episodic-like memories beyond its initial training set, so while it can exhibit remarkable veridical recall of training texts, it cannot remember in detail the subsequent events that it experiences, in the way that humans do. Many researchers are working on ways of adding this capacity to LLMs, and the additional levels of control over encoding and retrieval of episodic memory are interesting issues to explore (e.g., [[@ZhengWolfRanganathEtAl25]]).

In summary, the transformer can be seen as capturing essential properties of the [[cognitive architecture]] of the human brain, just using different underlying neural hardware. As such, the successes of these models provides a concrete demonstration that artificial versions of this cognitive architecture can actually replicate many of the most important aspects of human cognition, and that we have much to learn for the [[Axon]] models by better understanding how all of this stuff works in practice.


