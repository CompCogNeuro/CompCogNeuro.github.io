+++
Categories = ["Learning", "Axon"]
bibfile = "ccnlab.json"
+++

The **GeneRec** (generalized recirculation) algorithm provides a more biologically plausible way of computing approximately the same _error gradients_ as the widely used but biologically implausible [[error backpropagation]] algorithm ([[@OReilly96]]). It does this by using [[bidirectional connectivity]] and a [[temporal derivative]] learning rule to effectively propagate the error signals. Although the Axon model does not directly use this algorithm, the [[kinase algorithm]] that it does use computes qualitatively the same error gradients, and GeneRec provides a more abstract and mathematically derived basis for understanding how and why these algorithms work.

{id="figure_bidir-err" style="height:25em"}
![How bidirectional activation propagation can communicate error signals within a network, in the case of a predictive learning network generating an initial minus phase prediction followed by an input signal that communicates the actual outcome to the top-most Prediction layer, in the plus phase. Time evolves from left to right, and the thick colored lines show the activation level of each of the three neurons over time, both in terms of the line height and the brightness and warmth of the color gradient. Snapshots of the activation state are shown at key points, using the same color scale. Initially, each neuron is inactive (blue). Then, sensory input excites the Input neuron, and a wave of bottom-up excitation propagates upward through the Hidden and Prediction layers. Critically, the Prediction and Hidden neurons mutually excite each other via bidirectional connections, which contributes to each of their activity levels. At the start of the plus phase, the actual Outcome arrives, which contradicts the prediction of "high activity", and directly inhibits the Prediction neuron, which comes to reflect the actual outcome in the plus phase. This decreased excitation from the Prediction neuron thus causes the Hidden neuron to also become less active, which is the key mechanism by which the top-down connections communicate the prediction error. All neurons learn based on the temporal difference between the plus phase state and the end of the minus phase state (where the activity snapshot is), as shown by the error brackets.](media/fig_generec_bidir_err.png)

All of the major dynamics that enable bidirectional excitatory connections coupled with a temporal derivative learning mechanism to compute an error gradient are shown in [[#figure_bidir-err]], which traces the activity of three representative neurons over time, in a [[predictive learning]] context where the network generates a "high activity" prediction on the Prediction layer neuron, which is contradicted by the "low activity" Actual outcome. At the start, all neurons are inactive (which is not typically the case in the [[neocortex]], but is useful for the illustration). Then, sensory input drives activity in the Input neuron, which then excites the Hidden and Prediction neurons in turn, assuming they have strong excitatory weights between them (in a real network, there are many neurons in each layer and richer patterns of weights, etc).

The result of this initial wave of activity is a briefly stable _Minus phase_ state of activity throughout the network, which is shown with the "snapshot" of the neurons at that point in time. This network state represents the collective Prediction state of the network. Critically, the bidirectional excitatory connectivity causes the Hidden and Prediction neurons to be more excited than they would with only feedforward connections: they are mutually reinforcing each other's activity. _It is this mutual interdependence that enables bidirectional connectivity to communicate error signals_.

You can see this when the _Plus phase_ arrives, which is when the Actual outcome that was being predicted actually happens, and drives neural activity directly in the Prediction layer. In the case illustrated in the figure, this Actual outcome is "not active" (blue), which contradicts the "high activity" prediction. Therefore, the Prediction layer is inhibited by this "negative" outcome, and its activity decreases. This decrease in activity then propagates down to the Hidden layer, which also becomes less active.

Thus, the Hidden neuron experiences a temporal derivative that reflects the corresponding temporal changes on the Prediction layer, by virtual of their bidirectional connectivity. The mathematical derivation below shows that this temporal derivative provides a good approximation to the error gradient computed by error backpropagation. In sum, the GeneRec algorithm shows how bidirectional connectivity, which is a prominent feature of the [[neocortex]], enables powerful error-driven learning in a biologically plausible manner.

To clarify a few issues with this example, the actual dynamics play out with distributed patterns of activity in the context of pooled [[inhibition]], so the difference between a prediction and an outcome is not generally about "high activity" vs. "low activity" but rather about different distributed patterns of activity.

## Mathematical derivation

{id="figure_generec-bp"}
![Illustration of GeneRec computation in three-layer network, for comparison with the corresponding figure from the error backpropagation page. Activations settle in the expectation / minus phase, in response to input activations presented to the input layer. Activation flows bidirectionally, so that the hidden units are driven both by inputs and activations that arise on the output units. In the outcome/plus phase, "target" values drive the output unit activations, and due to the bidirectional connectivity, these also influence the hidden units in the plus phase. Mathematically, changing the weights based on the difference in hidden layer activation states between the plus and minus phases results in a close approximation to the delta value computed by backpropagation. This same rule is then used to change the weights into the hidden units from the input units (delta times sending activation), which is the same form used in backpropagation, and identical in form to the delta rule.](media/fig_generec_compute_delta.png)

The core of this analysis revolves around the following simpler version of the GeneRec equation, which we call the *GeneRec delta equation*:

$$
\Delta w = x^- \left(y^+ - y^- \right)
$$

where the weight change is driven only by the *delta* in activity on the receiving unit *y* between the plus (outcome) and minus (expectation) phases, multiplied by the sending unit activation *x*.  One can derive the full CHL equation from this simpler GeneRec delta equation by adding a constraint that the weight changes computed by the sending unit to the receiving unit be the same as those of the receiving unit to the sending unit (i.e., a *symmetry constraint* based on bidirectional connectivity), and by replacing the minus phase activation for the sending unit with the average of the minus and plus phase activations (which ends up being equivalent to the *midpoint method* for integrating a differential equation).  You can find the actual derivation later in the section *GeneRec and Activation Differences*, but you can take our word for it for the time being.

This equation has the desired property of credit assignment: the weight change is proportional to $ x^- $, which reflects how much this sender contributed to the error being learned from.

Interestingly, the GeneRec delta equation is equivalent in form to the *delta rule*, which we derive below as the optimal way to reduce error in a two layer network (input units sending to output units, with no hidden units in between).  The delta rule was originally derived by [@WidrowHoff60], and it is also basically equivalent to a gradient descent solution to linear regression.

But two-layer networks are very limited in what they can compute. As we discussed in the *Networks*  Chapter, you really need those hidden layers to form higher-level ways of re-categorizing the input, to solve challenging problems (you will also see this directly in the simulation explorations in this chapter).  As we discuss more below, the limitations of the delta rule and two-layer networks were highlighted in a very critical paper by [@MinskyPapert69], which brought research in the field of neural network models nearly to a standstill for nearly 20 years.

Before we unpack this equation a bit more, let's consider what happens at the *output* layer in a standard three-layer backprop network like that pictured in the Figure.  In these networks, there is no outcome/plus phase, but instead we just compare the output activity of units in the output layer (effectively the expectation) and compute externally the difference between these activities and the *target* activity values *t*. The difference is the *delta* value:

$$ 
\delta = t - z
$$

and is used to drive learning by changing the weight from sending unit y in the hidden layer to a given output unit z is:

$$
\Delta w = y \delta = y (t - z)
$$

You should recognize that this is exactly the *delta rule* as described above (where we keep in mind that y is now a sending activation to the output units).  The delta rule is really the essence of all error-driven learning methods.

Now let's get back to the delta backpropagation equation, and see how we can get from it to GeneRec (and thus to XCAL).  We just need to replace the $\delta_k$ term with the value for the output units, and then do some basic rearranging of terms, and we get very close to the GeneRec delta equation:

$$
\Delta w = x \left( \sum_k (t_k - z_k) w_k \right) y'
$$

$$
\Delta w = x \left( \sum_k t_k w_k - \sum_k z_k w_k \right) y'
$$

If you compare this last equation with the GeneRec delta equation, they would be equivalent (except for the *y'* term that we're still ignoring) if we made the following definitions:

$$
y^+ = \left. \sum_k t_k w_k \right.
$$

$$
y^- = \left. \sum_k z_k w_k \right.
$$

$$
x^- = x
$$

Interestingly, these sum terms are identical to the *net input* that unit *y* would receive from unit *z* if the weight went the other way, or, critically, if *y* also received a *symmetric, bidirectional connection* from *z*, in addition to sending activity to *z*.  Thus, we arrive at the critical insight behind the GeneRec algorithm relative to the backpropagation algorithm:

> *Symmetric bidirectional connectivity can convey error signals as the difference between two activity states (plus/outcome vs. minus/expectation), instead of sending a single "delta" error value backward down a single weight in the opposite (backpropagation) direction.*

The only wrinkle in this argument at this point is that we had to assign the activation states of the receiving unit to be equal to those net-input like terms (even though we use non-linear thresholded activation functions), and also those net input terms ignore the other inputs that the receiving unit should also receive from the sending units in the input layer.  The second problem is easily dispensed with, because those inputs from the input layer would be common to both "phases" of activation, and thus they cancel out when we subtract $y^+ - y^-$.  The first problem can be solved by finally no longer ignoring the *y'* term --- it turns out that the difference between a function evaluated at two different points can be approximated as the difference between the two points, times the derivative of the function:

$$
f(a) - f(b) \approx f'(a) (a-b)
$$

So we can now say that the activations states of y are a function of these net input terms: 

$$
y^+ =  f \left( \sum_k t_k w_k \right)
$$

$$
y^- =  f \left( \sum_k z_k w_k \right)
$$

and thus their difference can be approximated by the difference in net inputs times the activation function derivative:

$$
y^+ - y^- \approx y' \left( \sum_k t_k w_k - \sum_k z_k w_k \right)
$$

Which gets us right back to the GeneRec delta equation as being a good approximation to the delta backpropagation equation:

$$
\Delta w = x^- \left(y^+ - y^- \right)  \approx x \left( \sum_k \delta_k w_k \right) y'
$$

So if you've followed along to this point, you can now rest easy by knowing that the GeneRec (and thus XCAL) learning functions are actually very good approximations to error backpropagation. As we noted at the outset, XCAL uses bidirectional activation dynamics to communicate error signals throughout the network, in terms of averaged activity over two distinct states of activation (expectation followed by outcome), whereas backpropagation uses a biologically implausible procedure that propagates a single error value (outcome - expectation) backward across weight values, in the opposite direction of the way that activation typically flows.

