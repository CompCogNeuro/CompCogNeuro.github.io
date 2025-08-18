+++
Name = "PBWM"
Categories = ["Computation", "Neuroscience"]
bibfile = "ccnlab.json"
+++

The **PBWM** (prefrontal-cortex, basal-ganglia working memory) model ([[@OReillyFrank06]]) shows how **gating** mechanisms like those in the [[LSTM]] computational model can arise from the disinhibitory (modulatory, multiplicative) influence of the [[basal ganglia]] over the excitatory loops between the [[thalamus]] and [[prefrontal cortex]]. This model also demonstrated how phasic [[dopamine]] signals could train this gating functionality in a biologically-plausible manner, in contrast to the implausible [[error backpropagation]] mechanisms used in LSTM.

{id="figure_pbwm" style="height:25em"}
![Architecture of the PBWM system, where PFC provides top-down biasing to control sensory-motor processing in the rest of the brain, while the basal ganglia (BG) drives dynamic gating of PFC task / goal representations. The BG in turn is trained via phasic dopamine signals based on successful task performance.](media/fig_pbwm_architecture_bio.png)

[[#figure_pbwm]] shows the architecture of the PBWM system, with each component providing a well-defined function that together yields a dynamic form of cognitive control over behavior, relative to learned stimulus-response associations.

The need for a dynamic gating system in the basal ganglia arises because a robust maintenance system like the PFC cannot simultaneously update to encode new information and robustly maintain existing information. If it is tuned for robust maintenance, then that means it will not be distracted by new information. If it is tuned for rapid updating to encode new information, then it will not robustly maintain existing information.

This is a fundamental tradeoff that must be adjudicated by a dynamic, modulatory control signal that can rapidly switch the tuning between rapid updating and robust maintenance.

## Limitations of PBWM

There are a number of limitations of the PBWM framework that are addressed in the new [[Axon]] and [[Rubicon]] models.

* Limited learning capabilities. The reliance on phasic dopamine to train BG gating signals via something like standard [[reinforcement learning]] mechanisms is very inefficient, especially as the dimensionality of the problem space increases (i.e., the [[curse of dimensionality]]). It essentially amounts to serial trial-and-error [[search]] over when to update vs. maintain, whereas only a parallel, gradient-based learning mechanism scales well as dimensionality increases.

* Fine-grained gating signals are implausible and impractical. An LSTM model typically has separate input and output gates for each individual memory unit, providing a very fine-grained level of control. PBWM hypothesized larger pools or stripes of PFC units all controlled by a common set of gates, which is more consistent with the biological parameters. However, the relevant thalamic connectivity into the PFC appears to be relatively broad and diffuse, which is inconsistent with a fine-grained level of control. Furthermore, finer-grained control signals exacerbate the curse of dimensionality learning problems.

These limitations are addressed in the current [[Axon]] and [[Rubicon]] models as follows:

* Goal-driven areas of PFC (i.e., most of rodent PFC) are all updated by a common goal-engagement gating signal driven by ventral and medial BG areas, which modulate the MD (mediodorsal) thalamus that provides broad connectivity across all of the distributed goal areas in PFC. This low-dimensional signal is learned in the context of a careful goal-selection process that leverages prior learning across multiple relevant dimensions, and bidirectional [[constraint satisfaction]] among all the goal PFC areas, to make optimized choices.

* Motor areas of PFC and frontal neocortex that are involved in executing motor actions under influence of an engaged goal can use BG modulated thalamic signals as parallel, graded learning signals, similar to how the pulvinar nucleus of the thalamus operates in [[predictive learning]]. This builds on the powerful graded, parallel BG learning mechanisms, rather than the serial gating-like functionality envisioned in PBWM.

