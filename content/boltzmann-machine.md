+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.json"
+++

The **Boltzmann-machine** (BM) is a pioneering neural network learning algorithm from the 1980s, which was the first to use a [[temporal derivative]] to perform [[error driven learning]] ([[@AckleyHintonSejnowski85]]; [[@HintonSejnowski86]]). It was derived from a statistical physics framework in terms of the Boltzmann equation, which defines an overall _engergy function_ over a collection of interacting elements, which are typically atoms in statistical physics, and are neurons in this formulation. This framework is very similar to the **Hopfield network** ([[@Hopfield82]]; [[@Hopfield84]]).

Learning in the BM is a function of the difference between statistics computed in the **plus phase** vs. the **minus phase**, where the plus phase is the state of the network with the "right answer" present, while the minus phase has only an input pattern and the network is attempting to generate the right answer. Because the BM is based on a fully bidirectionally-connected network, any subset of neurons could be used to represent the input pattern, and any other subset could represent the desired output pattern. Thus, the BM nicely captures the generality and robustness of the [[temporal derivative]] learning framework.

However, in practice, the BM requires extensive periods of statistical sampling in each of the two phases, and does not learn very well when adding multiple different hidden layers (i.e., **deep networks**). In effect, it is highly susceptible to getting stuck in **local minima** of the overall error surface.

These limitations are overcome in the [[Axon]] framework in a variety ways, including the use of pooled [[inhibition]] (which keeps neurons in their sensitive range for responding to the differences across phases) and sparse representations, and differential weight strengths for feedback vs. feedforward connections. The [[GeneRec]] and [[kinase algorithm]] learning mechanisms build on the basic principles of the plus and minus phases as originally developed in the BM.

## Contrastive Hebbian Learning (CHL)

The specific form of the learning rule in the BM is the **Contrastive Hebbian Learning (CHL)** equation ([[@MovellanMcClelland93]]):

{id="eq_chl" title="Contrastive Hebbian Learning (CHL)"}
$$
\Delta w = \left(x^+ y^+\right) - \left(x^- y^-\right)
$$

Here, the first term is the activity of the sending and receiving units during the outcome (in the _plus phase_), while the second term is the activity during the prediction (in the _minus phase_). CHL is so-named because it involves the contrast or difference between two [[Hebbian learning#Hebbian]]-like terms.

The [[GeneRec]] algorithm ([[@OReilly96]]) derives this same CHL equation directly from [[error backpropagation]], along with a couple of other assumptions.

