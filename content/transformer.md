+++
Categories = ["Learning", "Activation", "Computation"]
bibfile = "ccnlab.json"
+++

The **transformer** architecture, developed by researchers at Google ([[@VaswaniShazeerParmarEtAl17]]), powers [[large language models]] (LLMs) and contains two distinct mechanisms, _self-attention_ and high-dimensional _feedforward network_ (FFN) (i.e., standard [[error-backpropagation]]). The [[attention]] component gets all the attention, so to speak, but apparently the FFN component is responsible for the considerable memorization abilities of this architecture, as evident in LLMs ([[@NandaRajamanoharanKramarEtAl23]]). It also depends on a number of other widely-adopted [[abstract neural network]] mechanisms including critically the ResNet (residual network) [[@HeZhangRenEtAl15]] and normalization mechanisms.

{id="figure_transformer-attn" style="height:20em"}
![Illustration of the attentional modulation performed in a transformer. Each input token (word) serves as a Query that is multiplied (dot product) with the Keys for each other token, with the normalized (softmax) magnitude multiplying the Value of each token, which is the result that goes on to the next stage. Instead of literally using the token representation, each attention head learns to transform the token to shape the Query, Key, and Value representations so that they end up being useful (via error backpropagation) to solving the problem. Multiple such heads are applied at each stage.](media/fig_transformer_attention.png)

The attentional mechanism is illustrated in [[#figure_transformer-attn]], which is computed for each input token serving as a _Query_ that is compared (via a dot product) with the _Key_ values for each of the other words. The representation of the token as a Query and Key is determined by separate trained weights (via error backpropagation through the entire network), which allows the attention to be somewhat flexible in terms of what it extracts from each token, and how it tries to match against the other tokens as Keys.

The limited, focal nature of attention is captured by the application of a _SoftMax_ normalizing function that is the normalized exponential function of the Query-Key dot products (which are single scalar values):

{id="eq_attn" title="Attention"}
$$
\rm{Attention}(Q, K) = \rm{softmax} \left( \frac{Q \cdot K}{\sqrt{d_k}} \right)
$$

{id="eq_softmax" title="SoftMax"}
$$
\rm{softmax}(\bf{X}) = \left[ \frac{e^{x_i}}{\sum_j e^{x_j}}, ... \right]_i
$$

The resulting normalized scalar attention value for each _Key_ (represented by the downward arrows in [[#figure_transformer-attn]]) then multiplies the _Value_ vector for that token, which again is transformed by linear weights, and it is this transformed, attention-modulated value that is the output result of the attention mechanism (shown by the big arrows flowing from input to output in the figure).

Intuitively, this attentional mechanism allows each input word to select the other words that are most relevant to it, so in the example in the figure, the input word "orange" modifies "cat" so it gives a strong attentional weight to cat.

Because of the exponential SoftMax normalization, the attentional mechanism typically focuses on relatively few other words, so multiple such attentional _heads_ are used in parallel to allow multiple attentional foci to be processed in the same feedforward sweep.

At a cognitive level, this multi-head attention likely reflects some element of both parallel [[constraint satisfaction]] and serial, sequential attention. Because the model is trained with standard [[error backpropagation]], it is critical for the credit assignment process to have everything of relevance present all at once in parallel -- otherwise a much more complex backpropagation through time process would be required to accomplish [[temporal credit assignment]]. 

As noted above, it is actually the FFN component that accounts for much of the memorization ability of a transformer, which is essential for these models to be able to digest such a huge amount of information and generate any kind of prediction at an above-chance level. The FFN employs a widely-used strategy that is similar to what happens in the [[hippocampus]] and [[cerebellum]], by projecting the attentionally-weighted tokens into a _high dimensional space_, which allows activity patterns to become relatively more separated from each other (i.e., _pattern separation_). In the standard GPT architecture, this hidden layer is 4x larger than the size of the token input vectors, and overall about 70% of the total LLM learnable parameters are contained in this FFN component.

The ResNet property of transformers is critical for essentially providing a bus-like pathway that the attention heads can partition and manipulate as the information flows forward through the network ([[@ElhageNandaOlssonEtAl21]]). Specifically, each level of the network passes its outputs directly through to the final output stage of the entire network, and additional levels just add their contributions to this "residual stream" of information.

Each higher layer only gets information from the level just below it, so that contextualizes the processing performed by that layer, but the ability to directly pass lower-level information straight through to the output is essential not only for making the error backpropagation work effectively (otherwise the error gradients would be massively diluted by the time they filtered through all the higher layers), but also for enabling higher layers to serve as "dynamic editors" of an accumulating flow of transformed information extracted from the inputs.

