+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

Here we describe in full detail how the excitatory conductance $Ge$ or *net input* to the neuron is computed, taking into account differences across different sources of input to a given neuron.  In the main chapter, the core computation is summarized, as an average of the weights times sending activations:

{id="eq_gesum"}
$$
g_e(t) = \frac{1}{n} \sum_i x_i w_i
$$

where *n* is the total number of channels, and $x_i$ is the **activity** of a particular sending neuron indexed by the subscript *i*, and $w_i$ is the **synaptic weight strength** that connects sending neuron *i* to the receiving neuron.

The overall goal of the more elaborate net input calculation described here, which is what is actually used in the *emergent* software, is to ensure that inputs from different layers having different overall levels of activity have a similar impact on the receiving neuron in terms of overall magnitude of net input, while also allowing for the strength of these different inputs to be manipulated in ways that continue to work functionally.  For example, a "localist" input layer may have only one unit active out of 100 (1%) whereas a hidden layer may have 25% activity (e.g., 25 out of 100) --- this vast difference in overall activity level would make these layers have very disparate influence on a receiving layer if not otherwise compensated for.  Terminologically, we refer to the set of connections from a given sending layer as a **projection**.

The full equation for the net input is as follows, which contains a double sum, first over the different projections, indexed by the letter *k*, and then within that by the receiving connections for each projection, indexed by the letter *i* (where these are understood to vary according to the outer projection loop):

{id="eq_get"}
$$
g_e(t) = \sum_k \left[ s_k \left(\frac{r_k}{\sum_p r_p}\right) \frac{1}{\alpha_k} \frac{1}{n_k} \sum_i \left( x_i w_i \right) \right]
$$

The factors in this equation are:

* $s_k$ = absolute scaling parameter (set at the user's discretion) for the projection, represented by `WtScale.Abs` parameter in the LeabraConSpec in *emergent*.

* $r_k$ = relative scaling parameter for the projection, which is always normalized by the sum of the relative parameters for all the other projections, which is what makes it relative --- the total is constant and one can only alter the relative contributions --- represented by `WtScale.Rel` in *emergent*.

* $\alpha_k$ = effective expected activity level for the sending layer, computed as described below, which serves to equalize projections regardless of differences in expected activity levels.

* $n_k$ = number of connections within this projection

The equations for computing the effective expected activity level $\alpha_k$ are based on the integer counts of numbers of expected active inputs on a given projection --- this takes into account both the sending layer expected activation, and the number of connections being received.  For example, consider a projection from a layer having 1% activity (1 out of 100 units active), with only a single incoming connection from that layer.   Even though the odds of this single incoming connection having an active sending unit are 1% on average, *some* receiving unit in the layer is highly likely to be getting that 1 sending unit active.  Thus, we use the "highest expected activity level" on the layer, which is 1, rather than the average expected sending probability, which is 1%.

Specifically, the equations, using pseudo-programming variables with longer names instead of obscure mathematical symbols, are:

* `alpha_k = MIN(%_activity * n_recv_cons + sem_extra, r_max_act_n)`

    + `%_activity` = % expected activity on sending layer
    + `n_recv_cons` = number of receiving connections in projection
    + `sem_extra` = standard error of the mean (SEM) extra buffer, set to 2 by default --- this makes it the highest expected activity level by including effectively 4 SEM's above the mean, where the real SEM depends on `%_activity` and is a maximum of .5 when `%_activity` = .5.
    + `r_max_act_n = MIN(n_recv_cons, %_activity * n_units_in_layer)` = hard upper limit maximum on number of active inputs --- can't be any more than either the number of connections we receive, or the total number of active units in the layer

See the [emer/leabra](https://github.com/emer/leabra) README docs for more detailed information about parameters, monitoring the actual relative net input contributions from different projections, etc.



