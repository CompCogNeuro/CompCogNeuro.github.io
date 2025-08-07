+++
Categories = ["Rubicon", "Simulations"]
bibfile = "ccnlab.json"
+++

{id="sim_bg" collapsed="true"}
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

Notes from readme:

The DS is the input layer for the primary motor control part of the basal ganglia, and this model learns to execute a sequence of motor actions through reinforcement  learning (RL), getting positive reinforcement for correct actions and lack of reinforcement for incorrect ones.  Critically, there is no omnicient "teacher" input: the model has to discover the correct action sequence purely through trial and error, "online" learning (i.e., it learns on a trial-by-trial basis as it acts). This is the only biologically / ecologically realistic form of RL.

The model also has mechanisms to learn about the space of possible motor actions and their parameterization in the DS, which is driven by ascending pathways from the brainstem and spinal motor system, via the deep cerebellar nuclei (DCN) to the CL (central lateral) nucleus of the thalamus.  This pathway is not currently beneficial in the model, and will be revisited once a simple model of the cerebellum is implemented, in the context of more fine-grained parameterized motor control.  The descending pathway from motor cortex also conveys useful motor signal information, and these cortical reps are directly shaped by the same ascending motor signals.

The model uses a simple dopamine (DA) "critic" system implemented in code, in `mseq_env.go`.  It just adapts a `RewPred` prediction whenever a reward is processed, using a simple learning rate (`RewPredLRate=0.01`).  The `RewPredMin=0.1` is critical to maintain negative DA signals for failures, so these cannot be fully predicted away.  Partial credit is given for partially correct sequence outputs, which is important for the length-3 sequences but actually somewhat detrimental to the length-2 case.  Partial credit is computed as a *probability* of reward as a function of the number of correct actions: `p = NCorrect / SeqLen`, and the reward value if given is also equal to this `p` value.  If you just continuously give a constant `p` value partial credit, the model learns to expect this and fails to progress further.

These are ecologically reasonable properties under the assumption that reward is discounted by effort, and random action choices will sometimes result in good outcomes, but generally with additional unnecessary steps.

To simplify the use of a consistent motor sequence across the parallel-data processing used on GPU (NData > 1), we just keep the target sequence as 0,1,2 etc, because the model doesn't know any better, and the random initial weights have no bias either.

