+++
Categories = ["Rubicon", "Simulations"]
bibfile = "ccnlab.json"
+++

{id="sim_inhib" collapsed="true"}
```Goal
// see https://github.com/emer/axon/tree/main/sims/bgdorsal for source code
bgdorsal.Embed(b)
```

<div>

This simulation explores the PCore model of [[basal ganglia]] (BG) function in the context of motor control, as supported by the dorsolateral striatum and associated BG pathways. It is best to do the [[BG ventral simulation]] first to understand the basic functionality of the PCore model.

* Final motor output in `MotorBS` is determined by a softmax: selection occurs very "late" in the spinal cord, not in the BG itself.

* Partial reward at the end, probability of reward and its magnitude are determined by number of correct actions. TODO: mastermind example-- partial information can potentially be very informative. Can learn without this but just much shorter sequences.

* Layers do not have any of the cortical [[inhibition]] functions active: all based on direct inhibitory synaptic transmission.

* Explain TD comparison model.


