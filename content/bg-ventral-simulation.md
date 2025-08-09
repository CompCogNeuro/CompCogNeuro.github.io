+++
Name = "BG ventral simulation"
Categories = ["Rubicon", "Simulations"]
bibfile = "ccnlab.json"
+++

{id="sim_vmbg" collapsed="true"}
```Goal
// see https://github.com/emer/axon/tree/main/sims/bgventral for source code
bgventral.Embed(b)
```

<div>

This simulation provides the simplest exploration of the PCore model of [[basal ganglia]] function, in the context of ventral and medial contributions to the goal-selection process, which is a central component of the [[Rubicon]] model.

The network receives input activations from two cortical areas that represent the net positive (benefits) vs. negative (costs) value associated with the current potential goal and action plan being evaluated. These are encoded using a simple gaussian "bump" over consecutive neurons, with low values in the lower left and value increasing to the right and up (back) in the layer. Because the initial weights are random, the striatal neurons do not know what these signals mean, and they must therefore learn the meaning by using [[dopamine]]-like reward prediction error (RPE) signals based on random exploration.

<!--- TODO: update thalamus to MD! -->

The BG model drives an overall disinhibitory _gating_ of the MD thalamus, which in turn projects back up into the frontal cortex to determine whether to lock-in the current goal being considered (i.e., to cross the Rubicon and transition into the goal-engaged state), or to skip this goal and move on to considering another. The dopamine feedback is computed (in the simulation code in this case -- see the [[PVLV simulation]] for a biologically-based model) as a function of this goal-gating in relation to the balance of positive vs. negative value represented in the input. If this balance is net positive, then the model is rewarded for disinhibitory gating, whereas it is punished if the balance is net negative. If nothing happens, then there is no dopamine feedback (i.e., "nothing wagered, nothing won").

* Click [[#sim_vmbg:Wts]] in the Network variables, and then on [[#sim_vmbg:r.Wt]] to view the receiving weights into neurons as you click on them in the network.

You should see the initial random weights associated with the pathways indicated by the arrows. The CT layer has recurrent self connections that allow it to better maintain information over time, so that it can leverage information from points even earlier than the prior trial.


TODO: look at weights, run a few trials, then run training, stop at around 20, step to see trained behavior, then continue to test.

{id="figure_bgventral-test" style="height:20em"}
![Testing results from the ventral BG model, trained with dopamine based on reward prediction error to do Go gating when the input signals indicate more positive reward versus negative costs are available, and No when the opposite is true. The testing sweeps through increments of negative costs in an inner loop, and positive rewards in the outer looop, as shown on the lower portion of the plot. The Gated line the proportion of times that the model did Go gating, which is strongly determined by the ratio of positive to negative, across the full range of these values. This demonstrates the balanced nature of the interactions between pathways. The RT line shows the normalized number of cycles taken when a Go gating outcome occurred, showing that the model was significantly slower in processing the cases with greatest conflict, where positive is very close to negative. Furthermore, the overall trend is that with stronger positive values, RT is overall faster. These patterns are widely observed in decision-making studies, and predicted by normative drift-diffusion models. This demostrates that the model naturally exhibits an information-integration dynamic to accumulate more input over time when the decision is more ambiguous.](media/fig_bgventral_test.png)


