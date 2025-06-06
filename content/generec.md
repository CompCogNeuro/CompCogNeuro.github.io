+++
Categories = ["Learning"]
bibfile = "ccnlab.json"
+++

The **GeneRec** (generalized recirculation) algorithm provides a more biologically plausible way of computing approximately the same _error gradients_ as the widely used but biologically implausible [[error backpropagation]] algorithm ([[@OReilly96]]).

For the present purposes, we can safely ignore these factors, which allows us to leverage all of the analysis that went into understanding GeneRec --- itself a large step towards biological plausibility relative to backpropagation. 

The core of this analysis revolves around the following simpler version of the GeneRec equation, which we call the *GeneRec delta equation*:
$$ \Delta w = x^- \left(y^+ - y^- \right) $$
where the weight change is driven only by the *delta* in activity on the receiving unit *y* between the plus (outcome) and minus (expectation) phases, multiplied by the sending unit activation *x*.  One can derive the full CHL equation from this simpler GeneRec delta equation by adding a constraint that the weight changes computed by the sending unit to the receiving unit be the same as those of the receiving unit to the sending unit (i.e., a *symmetry constraint* based on bidirectional connectivity), and by replacing the minus phase activation for the sending unit with the average of the minus and plus phase activations (which ends up being equivalent to the *midpoint method* for integrating a differential equation).  You can find the actual derivation later in the section *GeneRec and Activation Differences*, but you can take our word for it for the time being.

This equation has the desired property of credit assignment: the weight change is proportional to $ x^- $, which reflects how much this sender contributed to the error being learned from.

Interestingly, the GeneRec delta equation is equivalent in form to the *delta rule*, which we derive below as the optimal way to reduce error in a two layer network (input units sending to output units, with no hidden units in between).  The delta rule was originally derived by [@WidrowHoff60], and it is also basically equivalent to a gradient descent solution to linear regression.

But two-layer networks are very limited in what they can compute. As we discussed in the *Networks*  Chapter, you really need those hidden layers to form higher-level ways of re-categorizing the input, to solve challenging problems (you will also see this directly in the simulation explorations in this chapter).  As we discuss more below, the limitations of the delta rule and two-layer networks were highlighted in a very critical paper by [@MinskyPapert69], which brought research in the field of neural network models nearly to a standstill for nearly 20 years.

![Illustration of GeneRec/XCAL computation in three-layer network, for comparison with previous figure showing backpropagation.  Activations settle in the expectation/minus phase, in response to input activations presented to the input layer.  Activation flows bidirectionally, so that the hidden units are driven both by inputs and activations that arise on the output units.  In the outcome/plus phase, "target" values drive the output unit activations, and due to the bidirectional connectivity, these also influence the hidden units in the plus phase.  Mathematically, changing the weights based on the difference in hidden layer activation states between the plus and minus phases results in a close approximation to the delta value computed by backpropagation.  This same rule is then used to change the weights into the hidden units from the input units (delta times sending activation), which is the same form used in backpropagation, and identical in form to the delta rule.](media/fig_generec_compute_delta.png)

Before we unpack this equation a bit more, let's consider what happens at the *output* layer in a standard three-layer backprop network like that pictured in the Figure.  In these networks, there is no outcome/plus phase, but instead we just compare the output activity of units in the output layer (effectively the expectation) and compute externally the difference between these activities and the *target* activity values *t*. The difference is the *delta* value:
$$ \delta = t - z $$
and is used to drive learning by changing the weight from sending unit y in the hidden layer to a given output unit z is:
$$ \Delta w = y \delta = y (t - z) $$
You should recognize that this is exactly the *delta rule* as described above (where we keep in mind that y is now a sending activation to the output units).  The delta rule is really the essence of all error-driven learning methods.

Now let's get back to the delta backpropagation equation, and see how we can get from it to GeneRec (and thus to XCAL).  We just need to replace the $\delta_k$ term with the value for the output units, and then do some basic rearranging of terms, and we get very close to the GeneRec delta equation:
$$ \Delta w = x \left( \sum_k (t_k - z_k) w_k \right) y' $$
$$ \Delta w = x \left( \sum_k t_k w_k - \sum_k z_k w_k \right) y' $$
If you compare this last equation with the GeneRec delta equation, they would be equivalent (except for the *y'* term that we're still ignoring) if we made the following definitions:
$$ y^+ = \left. \sum_k t_k w_k \right. $$
$$ y^- = \left. \sum_k z_k w_k \right. $$
$$ x^- = x $$
Interestingly, these sum terms are identical to the *net input* that unit *y* would receive from unit *z* if the weight went the other way, or, critically, if *y* also received a *symmetric, bidirectional connection* from *z*, in addition to sending activity to *z*.  Thus, we arrive at the critical insight behind the GeneRec algorithm relative to the backpropagation algorithm:

> *Symmetric bidirectional connectivity can convey error signals as the difference between two activity states (plus/outcome vs. minus/expectation), instead of sending a single "delta" error value backward down a single weight in the opposite (backpropagation) direction.*

The only wrinkle in this argument at this point is that we had to assign the activation states of the receiving unit to be equal to those net-input like terms (even though we use non-linear thresholded activation functions), and also those net input terms ignore the other inputs that the receiving unit should also receive from the sending units in the input layer.  The second problem is easily dispensed with, because those inputs from the input layer would be common to both "phases" of activation, and thus they cancel out when we subtract $y^+ - y^-$.  The first problem can be solved by finally no longer ignoring the *y'* term --- it turns out that the difference between a function evaluated at two different points can be approximated as the difference between the two points, times the derivative of the function:
$$ f(a) - f(b) \approx f'(a) (a-b) $$
So we can now say that the activations states of y are a function of these net input terms: 
$$ y^+ =  f \left( \sum_k t_k w_k \right) $$
$$ y^- =  f \left( \sum_k z_k w_k \right) $$
and thus their difference can be approximated by the difference in net inputs times the activation function derivative:
$$ y^+ - y^- \approx y' \left( \sum_k t_k w_k - \sum_k z_k w_k \right) $$
Which gets us right back to the GeneRec delta equation as being a good approximation to the delta backpropagation equation:
$$ \Delta w = x^- \left(y^+ - y^- \right)  \approx x \left( \sum_k \delta_k w_k \right) y' $$

So if you've followed along to this point, you can now rest easy by knowing that the GeneRec (and thus XCAL) learning functions are actually very good approximations to error backpropagation. As we noted at the outset, XCAL uses bidirectional activation dynamics to communicate error signals throughout the network, in terms of averaged activity over two distinct states of activation (expectation followed by outcome), whereas backpropagation uses a biologically implausible procedure that propagates a single error value (outcome - expectation) backward across weight values, in the opposite direction of the way that activation typically flows.

