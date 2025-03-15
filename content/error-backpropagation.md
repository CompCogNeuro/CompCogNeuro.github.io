+++
Categories = ["Learning"]
bibfile = "ccnlab.bib"
+++

**Error backpropagation** ("backprop") is a relatively straightforward mathematical procedure for performing [[error-driven-learning]], derived directly from the _chain rule_ of basic calculus. It starts with an _objective function_ that defines how to compute the _error_. To determine how to update [[synaptic-weights]] to minimize this error, the first-order _partial derivative_ of the objective function with respect to a given weight value can be computed using the chain rule, which thereby defines the **error gradient**. The synaptic weights are iteratively updated along the direction of this gradient. Many different techniques can be used to implement this error minimization procedure, including using higher-order derivatives. 

Consistent with its basis in fundamental principles of calculus and function optimization, backprop has been "invented" by a number of people over the years, including [@WidrowHoff60; @Werbos74]. However, [@RumelhartHintonWilliams86] were the _last_ to invent it, which means they were the most successful in articulating and demonstrating the important properties of this learning procedure in the context of a comprehensive overall framework for modeling neurally-inspired learning [@RumelhartMcClelland86; @McClellandRumelhart86]. All of the widely-used modern [[artificial-intelligence]] / [[machine-learning]] models today are based primarily on backprop, including [[large-language-models]] (LLMs), providing a clear demonstration of the computational power of this technique, and of [[error-driven-learning]] more broadly.

The literal mathematical procedure for computing backprop is inconsistent with the known properties of the brain, as famously pointed out by Francis Crick [@Crick89]. However, the [[generec]] (generalized recirculation) approximation to backprop [@OReilly96; @LillicrapSantoroMarrisEtAl20] provides a mechanism for powerful error-driven learning that is consistent with biological mechanisms, as elaborated in the even more biologically detailed [[kinase-algorithm]], which leverages the key principle of a [[temporal-derivative]] for computing error gradients.

The most commonly-used technique to perform gradient descent in backprop models, including in the original work of [@RumelhartHintonWilliams86], is **online, stochastic gradient descent**. In this case, small, randomly-chosen subsets of _input patterns_ are used to compute the error gradients and update the synaptic weights ("online") before proceeding to the next input patterns. This is contrasted with _batch mode_ learning where the gradients are computed across all input patterns prior to updating the synaptic weights. Intuitively, the online form allows the network to more rapidly explore promising gradient directions and it breaks a lot of "ties" that otherwise occur when considering all of the input patterns at once: it is a more "decisive" form of learning, as contrasted with trying to simultaneously satisfy all of the constraints across all patterns, all at once.

## Derivation of backprop

![Illustration of backpropgation computation in three-layer network.  First, the feedforward activation pass generates a pattern of activations across the units in the network, cascading from input, to hidden to output.  Then, "delta" values are propagated *backward* in the reverse direction across the same weights.  The delta sum is broken out in the hidden layer to facilitate comparison with the GeneRec algorithm as shown in the next figure.](media/fig_bp_compute_delta.png){#fig:fig-bp-compute-delta width=70% }

TODO: cleanup the following, convoluted derivation from CCN text:

The essence of the backpropagation (also called "backprop") algorithm is captured in this *delta backpropagation equation*:
$$ \Delta w = x \left( \sum_k \delta_k w_k \right) y' $$
where *x* is again the sending activity value, $\delta$ is the error derivative for the units in the next layer *above* the layer containing the current receiving unit *y* (with each such unit indexed by the subscript *k*), and $w_k$ is the weight *from* the receiving unit y to the k'th such unit in the next layer above (see [@fig:fig-bp-compute-delta]).  Ignore the $y'$ term for the time being  --- it is the derivative of the receiving unit's activation function, and it will come in handy in a bit.

So we're propagating this "delta" (error) value *backward* across the weights, in the opposite direction that the activation typically flows in the "feedforward" direction, which is from the input to the hidden to the output (backprop networks are typically feedforward, though bidirectional versions have been developed as discussed below).  This is the origin of the "backpropagation" name.

Before we unpack this equation a bit more, let's consider what happens at the *output* layer in a standard three-layer backprop network like that pictured in the Figure.  In these networks, there is no outcome/plus phase, but instead we just compare the output activity of units in the output layer (effectively the expectation) and compute externally the difference between these activities and the *target* activity values *t*. The difference is the *delta* value:
$$ \delta = t - z $$
and is used to drive learning by changing the weight from sending unit y in the hidden layer to a given output unit z is:
$$ \Delta w = y \delta = y (t - z) $$
You should recognize that this is exactly the *delta rule* as described above (where we keep in mind that y is now a sending activation to the output units).  The delta rule is really the essence of all error-driven learning methods.

For the time being, we assume a linear activation function of activations from sending units *y*, and that we just have a simple two-layer network with these sending units projecting directly to the output units:
$$ z_k = \left. \sum_j y_j w_{jk} \right. $$

Taking the negative of the derivative of SSE with respect to the weight *w*, which is more easily computed by breaking it down into two parts using the *chain rule* to first get the derivative of SSE with respect to the output activation *z*, and multiplying that by the derivative of *z* with respect to the weight:
$$ \Delta w_{jk} = -\frac{\partial SSE}{\partial w_{jk}} = -\frac{\partial SSE}{\partial z_k} \frac{\partial z_k}{\partial w_{jk}}$$
$$ = 2 (t_k - z_k) y_j $$
 
When you break down each step separately, it is all very straightforward:
$$ \frac{\partial SSE}{\partial z_k}  = -2 (t_k - z_k) $$
$$ \frac{\partial z_k}{\partial w_{jk}} = y_j $$
(the other elements of the sums drop out because the first partial derivative is with respect to $z_k$ so derivative for all other $z$'s is zero, and similarly the second partial derivative is with respect to $y_j$ so the derivative for the other $y$'s is zero.)

First, we define the _activation function_ in terms of $\eta_j$ as the _net input_ to unit $j$, i.e., the product of sending activity $x_i$ and weights:
$$ \eta_j = \sum x_i w_{ij} $$

and then the unit activity is a function of this net input:
$$ y_j = f(\eta_j) $$

The goal is to again minimize the error (SSE) as a function of the weights, 
$$ \Delta w_{ij} = -\frac{\partial SSE}{\partial w_{ij}} $$.

The chain rule expansion of the basic activation function through hidden units $j$ and output units $k$ is thus:

$$ = -\frac{\partial SSE}{\partial z_k} \frac{\partial z_k}{\partial \eta_k} \frac{\partial \eta_k}{\partial y_j} \frac{\partial y_j}{\partial \eta_j} \frac{\partial \eta_j}{\partial w_{ij}} $$

Although this looks like a lot, it is really just applying the same chain rule as above repeatedly. To know how to change the weights from input unit $x_i$ to hidden unit $y_j$, we have to know how changes in this weight $w_{ij}$ are related to changes in the SSE. This involves computing how the SSE changes with output activity, how output activity changes with its net input, how this net input changes with hidden unit activity $y_j$, how in turn this activity changes with its net input $\eta_j$, and finally, how this net input changes with the weights from sending unit $x_i$ to hidden unit $y_j$. Once all of these factors are computed, they can be multiplied together to determine how the weight $w_{ij}$ should be adjusted to minimize error, and this can be done for all sending units to all hidden units (and also as derived earlier, for all hidden units to all output units). 

We again assume a linear activation function at the output for simplicity, so that $\partial z_k / \partial \eta_k = 1$. We allow for non-linear activation functions in the hidden units *y*, and simply refer to the derivative of this activation function as $y'$ (which for the common sigmoidal activation functions turns out to be $y (1-y)$ but we leave it in generic form here so that it can be applied to any differentiable activation function.  The solution to the above equation is then, applying each step in order,
$$ -\frac{\partial SSE}{\partial w_{ij}} = \sum_k (t_k - z_k) * 1 * w_{jk} * y' * x_i $$
$$ = x \left( \sum_k \delta_k w_{jk} \right) y' $$
as specified earlier.

Thus, the negative of $\partial SSE / \partial w_{jk}$ is $2 (t_k -z_k)$ and since 2 is a constant, we can just absorb it into the learning rate parameter. 

Breaking down the error-minimization in this way, it becomes apparent that the weight change should be adjusted in proportion to both the error (difference between the target and the output) *and* the extent to which the sending unit *y* was active. This modulation of weight change by activity of the sending unit achieves a critical **credit assignment** function (or rather blame assignment in this case), so that when an error is made at the output, weights should only change for the sending units that contributed to that error. Sending units that were not active did not cause the error, and their weights are not adjusted.

