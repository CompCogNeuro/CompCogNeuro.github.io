+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.json"
+++

The famous Canadian psychologist Donald O. Hebb predicted the nature of the NMDA channel many years in advance of its discovery, just by thinking about how learning should work at a functional level. Here is a key quote ([[@Hebb49]]):

> _Let us assume that the persistence or repetition of a reverberatory activity (or "trace") tends to induce lasting cellular changes that add to its stability.… When an axon of cell A is near enough to excite a cell B and repeatedly or persistently takes part in firing it, some growth process or metabolic change takes place in one or both cells such that A's efficiency, as one of the cells firing B, is increased._

This can be more concisely summarized as _cells that fire together, wire together._ As discussed in detail in [[synaptic plasticity]], the [[neuron channels#NMDA]] channel is essential for this process, because it requires both pre and postsynaptic activity to allow $Ca^{++}$ to enter and drive learning. It can detect the _coincidence_ of neural firing. Interestingly, Hebb is reputed to have said something to the effect of "big deal, I knew it had to be that way already" when someone told him that his learning principle had been discovered in the form of the NMDA receptor.

Mathematically, we can summarize Hebbian learning as:

{id="eq_hebb" title="Basic Hebbian learning"}
$$
\Delta w \propto x y
$$

where $\Delta w$ is the change in synaptic weight _w_, as a function of sending activity _x_ and receiving activity _y_. If you only use this equation, and the activity values are biologically-realistic firing rates, which are only positive numbers, then the synaptic weight can only go up in value. Thus, some additional factors are required in this case to produce a more balanced pattern of weight increases (LTP) and decreases (LTD). A few such factors are reviewed below.

However, even with a well-balanced form of Hebbian learning, and including the widely-investigated  [[synaptic plasticity#spike timing dependent plasticity]] (STDP; [[@BiPoo98]]) variant of Hebbian learning, the primary issue is that there is no mathematical or computational basis to expect that it will actually enable a network of neurons to learn the kinds of cognitive and behavioral functionality that humans and other animals exhibit. It is simply a "heuristic" principle that might make some intuitive sense, but all of the available evidence strongly indicates that it lacks the general purpose flexibility of [[error-driven learning]] mechanisms. See the [[pattern associator sim]] and [[hidden layer sim]] simulations for simple examples demonstrating this difference between Hebbian and error-driven learning.

The difference between Hebbian learning and error-driven learning at a mathematical level comes down to the presence of some kind of _difference-based_ element in equation, that drives learning based on a _contrast_ between two factors, rather than just the "main effect" of raw association. For example, the [[generec]] version of [[error-backpropagation]] learning results in the **Contrastive Hebbian Learning** (CHL) equation:

{id="eq_chl" title="Contrastive Hebbian learning"}
$$
\Delta w \propto (x^+ y^+) - (x^- y^-)
$$

You can see that this is just a difference between two Hebbian factors as in [[#eq_hebb]], where the _+_ superscript indicates neural activity in the _plus phase_ when the correct answer is present, while the _-_ indicates the _minus phase_ where the network is just guessing what the answer should be. In the context of [[predictive learning]], the minus phase immediately precedes the plus phase, as the prediction of what is going to happen, followed by what actually does happen.

This CHL equation provides a good qualitative summary of the actual [[kinase algorithm]] equations used in Axon, and illustrates how the intuitive core of Hebbian associative learning and the pre-post contingency of the NMDA receptor can all be carried over into the much more computationally-powerful error-driven learning domain.

Furthermore, a central premise of the [[Leabra]] algorithm is that a very small amount of Hebbian learning can be combined with this CHL-based error-driven learning, where the Hebbian factor acts as a kind of _bias_ or _regularization factor_ that shapes the learning toward encoding statistical regularities (see [[bias-variance tradeoff]] for general principles). This is similar to [[weight decay]] which is used in [[abstract neural network]] models and in statistical regression, except with Hebbian learning you also have weight values going up and not just decaying toward zero.

In the remainder of this page, various standard forms of Hebbian learning are reviewed.

## Hebbian learning and inhibitory competition

There are a number of _self organizing_ models that combine Hebbian learning with various forms of [[inhibition]] among neurons within a hidden layer, which are synergistic. The inhibitory competition forces individual neurons to _specialize_ on encoding different _separable features_ of the input environment. At an abstract mathematical level, 

* Kohonen, SOM, V1, etc.


## Hebbian demo

This section provides a detailed treatment of Hebbian learning and popular variants thereof.

{id="figure_hebb-demo" style="height:20em"}
![Simple Hebbian learning demonstration across 4 time steps (t=0 thru 3).  Bottom row of network has 3 input units, the last of which fires in an ''uncorrelated'' fashion with the other two.  They all start out with weights w = .1.  Receiving activity is just a linear sum of the sending activations times weights: $y = \sum x w = .1$ for the first time step.  Learning is simple Hebbian: $\Delta w = xy$.  As you fill in the remainder of the activations, weights, and weight changes, you will find that the two correlated input units dominate the receiving unit activation, and thus they end up being correlated in their activity, causing their weights to always increase.  The third unit sometimes goes up and sometimes down, with no net increase over time.  Thus, Hebbian learning discovers correlations in the inputs.](media/fig_hebb_demo_blank.png)

[[#figure_hebb-demo]] shows a simple demonstration of how Hebbian learning causes the receiving network to discover correlations in the patterns of input unit activation.  The input units that are correlated end up dominating the receiving unit activity, and thus the receiving unit ends up being correlated with this subset of correlated inputs, and their weights always increase under the Hebbian learning function.  Uncorrelated inputs bounce around without a systematic trend.  If you keep going, you'll see that the weights grow quickly without bound, so this is not a practical learning function, but it illustrates the essence of Hebbian learning.

Next, we do some math to show that the simplest version of Hebbian correlational learning, in the case of a single linear receiving unit that receives input from a set of input units, result in the unit extracting the first **principle component** of correlation in the patterns of activity over
the input units.

Because it is linear, the receiving unit's activation function is just the weighted sum of its inputs

$$
y_j = \sum_k x_k w_{kj}
$$

where $k$ (rather than the usual $i$) indexes over input units, for reasons that will become clear (and all of the variables are a function of the current time step $t$ reflecting different inputs). The weight change is:

$$
\Delta_t w_{ij} = \epsilon x_i y_j
$$

where $\epsilon$ is the *learning rate* and $i$ is the index of a particular input unit, and weights just increment these changes over time:

$$
w_{ij}(t+1) = w_{ij}(t) + \Delta_t w_{ij}
$$

To understand the aggregate effects of learning over many patterns, we can just sum the changes over time:

$$
\Delta w_{ij} = \epsilon \sum_t x_i y_j
$$

and we assume that $\epsilon = 1 / N$, where $N$ is the total number of patterns in the input. This turns the sum into an *average*:

$$
\Delta w_{ij} = \langle x_i y_j \rangle_t
$$

Next, substitute into this equation the formula for $y_j$, showing that the weight changes are a function of the *correlations* between the input units:

$$
\Delta w_{ij}  = \langle x_i \sum_k x_k w_{kj} \rangle_t 
$$

$$
= \sum_k \langle x_i x_k \rangle_t \langle w_{kj} \rangle_t
$$

$$
= \sum_k \bf{C}_{ik} \langle w_{kj} \rangle_t
$$

This new variable $\bf{C}_{ik}$ is an element of the *correlation matrix* between the two input units $i$ and $k$, where correlation is defined here as the expected value (average) of the product of their activity values over time ($\bf{C}_{ik} = \langle x_i x_k \rangle_t$).  You might be familiar with the more standard correlation measure:

$$
\bf{C}_{ik} = \frac{\langle (x_i - \mu_i)(x_k - \mu_k) \rangle_t} {\sqrt{\sigma^2_i \sigma^2_k}}
$$

which subtracts away the mean values ($\mu$) of the variables before taking their product, and normalizes the result by their variances ($\sigma^2$).  Thus, an important simplification in this form of Hebbian correlational learning is that it assumes that the activation variables have zero mean and unit variance.

The implication of all this is that where strong correlations exist across input units, the weights for those units will increase because this average correlation value will be relatively large. Interestingly, if we run this learning rule long enough, the weights will become dominated by the strongest set of correlations present in the input, with the gap between the strongest set and the next strongest becoming increasingly large.  Thus, this simple Hebbian rule learns the *first* (strongest) principal component of the input data.

One problem with the simple Hebbian learning rule is that the weights become infinitely large as learning continues.  One solution to this problem was proposed by [[@^Oja82]], known as *subtractive normalization*:

$$
\Delta w_{ij} = \epsilon (x_i y_j - y^2_j w_{ij})
$$

As we did in Chapter 2, you just set the equation equal to zero and solve for the *equilibrium* or *asymptotic* weight values:

$$
0 = \epsilon(x_i y_j - y^2_j w_{ij})
$$

$$
w_{ij} = \frac{ x_i}{y_j}
$$

$$
w_{ij} = \frac{ x_i}{\sum_k x_k w_{kj}}
$$

Thus, the weight from a given input unit will end up representing the proportion of that input's activation relative to the total weighted activation over all the other inputs.  This will keep the weights from growing without bound.  Finally, because it is primarily based on the same correlation terms $\bf{C}_{ik}$ as the previous simple Hebbian learning rule, this Oja rule still computes the first principal component of the input data (though the proof of this is somewhat more involved, see [[@^HertzKroghPalmer91]] for a nice treatment).

Moving beyond a single hidden unit, there are ways of configuring inhibition so that the units end up learning the sequence of PCA values of the correlation matrix in eigenvalue order ([[@Sanger89]]; [[@Oja89]]). In [[@^OReillyMunakata00]], we developed a different alternative known as **conditional principal components analysis** or CPCA, which assumes that we want the weights for a given input unit to represent the conditional probability that the input unit ($x_i$) was active given that the receiving unit ($y_j$) was also active:

$$
w_{ij} = P(x_i = 1 | y_j = 1)
$$

$$
w_{ij} = P(x_i | y_j)
$$

where the second form uses simplified notation that will continue to be used below. 

The important characteristic of CPCA is that the weights will reflect the extent to which a given input unit is active across the subset of input patterns represented by the receiving unit (i.e., conditioned on this receiving unit).  If an input pattern is a very typical aspect of such inputs, then the weights from it will be large (near 1), and if it is not so typical, they will be small (near 0).

Following the analysis of [[@^RumelhartZipser85]], the CPCA learning rule can be derived as:

$$
\Delta w_{ij} = \epsilon [y_j x_i - y_j w_{ij}]
$$

$$
= \epsilon y_j (x_i - w_{ij})
$$

The two equivalent forms of this equation are shown to emphasize the similarity of this learning rule to Oja's normalized PCA learning rule, while also showing its simpler form, which emphasizes that the weights are adjusted to match the value of the sending unit activation $x_i$ (i.e., minimizing the difference between $x_i$ and $w_{ij}$), weighted in proportion to the activation of the receiving unit ($y_j$).  

We use the expression $P(y_j | t)$ to represent the probability that the receiving unit $y_j$ is active given that some particular input pattern $t$ was presented.  $P(x_i | t)$ represents the corresponding thing for the sending unit $x_i$. Substituting these into the learning rule, the total weight update computed over all the possible patterns $t$ (and multiplying by the probability that each pattern occurs, $P(t)$) is:

$$
\Delta w_{ij} = \epsilon \sum_t [P(y_j | t) P(x_i | t) - P(y_j | t) w_{ij}] P(t)
$$

$$
= \epsilon \left( \sum_t P(y_j | t) P(x_i | t) P(t) - \sum_t P(y_j | t) P(t) w_{ij} \right)
$$

As usual, we set $\Delta w_{ij}$ to zero and solve:

$$
w_{ij} = \frac{\sum_t P(y_j | t) P(x_i | t) P(t)} {\sum_t P(y_j | t) P(t)}
$$

Interestingly, the numerator is the definition of the joint probability of the sending and receiving units both being active together across all the patterns $t$, which is just $P(y_j, x_i)$. Similarly, the denominator gives the probability of the receiving unit being active over all the patterns, or $P(y_j)$.  Thus, we can rewrite the preceding equation as:

$$
w_{ij} = \frac{P(y_j, x_i)}{P(y_j)}
$$

$$
w_{ij} = P(x_i | y_j)
$$

at which point it becomes clear that this fraction of the joint probability over the probability of the receiver is just the definition of the conditional probability of the sender given the receiver.

Although CPCA is effective and well-defined mathematically, it suffers one major problem relative to the BCM formulation that we now use: it drives significant LTD (weight decrease) when a sending neuron is *not* active, and the receiving unit is active.  This results in a significant amount of interference of learning across time.  By contrast, the XCAL dWt function specifically returns to zero when either sending or receiving neuron has zero activity, and that significantly reduces interference, preserving existing weight values for inactive neurons.
