+++
Categories = ["Rubicon", "Neuroscience"]
bibfile = "ccnlab.json"
+++

The **basal ganglia** (BG) are a collection of nuclei located at the central "core" of the brain ([[#figure_bg-anatomy]]), forming the functional bridge between the highest levels of processing in the [[neocortex]], and the motor control networks in the brainstem (midbrain and hindbrain, progressing down to the spinal cord). The primary outputs of the BG are downward to these motor networks, and, via extensive connections into the thalamus, back up to the frontal areas of the neocortex and [[prefrontal cortex]]. The **striatum** is the input nucleus of the BG, anatomically composed of the _caudate_ and _putamen_, which receives extensive projections from the neocortex and other brain areas.

{id="figure_bg-anatomy" style="height:25em"}
![Major areas of the basal ganglia (BG) and associated brain areas. The striatum (composed of the caudate and putamen as large-scale anatomical features) is the input layer of the BG, receiving a wide range of inputs from all over the brain, especially the neocortex. The nucleus accumbens (NAc), i.e,. the ventral striatum, is the input area for the goal-driven, affective circuits. The globus pallidus externus (GPe) is the next stage of processing, providing the central "core" integration for the BG, with the GP internus (GPi) sending BG outputs subcortically and, via extensive connections into the thalamus, back up to the frontal neocortex. The substantia nigra pars reticulata (SNr) is also a major output pathway, while the pars compacta (SNc) provides dopamine innervation into the entire BG. The subthalamic nucleus (STN) also receives extensive cortical input from frontal areas, and interacts bidirectionally with the GPe. The amygdala is not a part of the BG but interacts extensively with it and is located at the extreme end of the tail, and the lateral habenula (LHb) plays an essential role in driving dips in dopamine, based on extensive inputs from the ventral striatum among many other areas.](media/fig_bg_anatomy_human.png)

The SPNs (spiny projection neurons; also called medium spiny neurons, MSNs) in the striatum receive the strongest [[dopamine]] inputs, and have the greatest ability to rapidly clear dopamine from the synapse after it has been released, of any area in the brain. This allows these neurons to learn from rapid _phasic_ changes in dopamine, i.d., _bursts_ and _dips_ relative to the baseline _tonic_ firing level, in a manner consistent with the core principles of [[reinforcement learning]] (RL; [[@MarkowitzGillisJayEtAl23]]; [[@HowardLiGeddesEtAl17]]; [[@NairGutierrez-ArenasErikssonEtAl15]]; [[@ShenFlajoletGreengardEtAl08]]; [[@Frank05]]).

There is extensive evidence showing that the BG encodes detailed aspects of motor control, including impressive decoding of complex naturalistic motor behaviors in freely-behaving rats ([[@MarkowitzGillisBeronEtAl18]]; [[@KlausMartinsPaixaoEtAl17]]), in addition to simpler conditioned learning tasks ([[@YttriDudman18]]). Furthermore, studies of animals with complete lesions of the neocortex demonstrate that the BG can drive extensive fine-grained motor control through its descending motor pathway projections ([[@GrillnerRobertsonKotaleski20]]; [[@ParkCoddingtonDudman20]]; [[@ArberCosta22]]). In evolutionarily more ancient vertibrates without much cortex, such as the lamprey (which has essentially the same BG structures as the rodent and primate; [[@GrillnerRobertson16]]), the BG is the primary driver of learning and behavior.

Thus, the striatum is widely believed to be the neural substrate for basic model-free RL, where reward prediction error (RPE) signals from the dopamine _critic_ system train the BG _actor_ to improve motor actions to increase overall reward intake (see [[reinforcement learning#figure_actor-critic]] in the RL page). 

However, in animals with more neocortex, especially in primates and humans, the BG also works intimately in concert with the cortex to organize behavior, by way of its extensive connections into the thalamus. These thalamic projections can provide two functions: direct modulation of cortical activity, which occurs in a very broad manner (which we term **gating**), and more fine-grained training of cortical learning through the same type of thalamic-driven [[predictive learning]] signals used in the standard [[Axon]] [[error-driven learning]] model. Thalamic gating is a central mechanism in the [[Rubicon]] framework, for locking in the active maintenance of distributed goal representations across various regions of the prefrontal cortex.

Thus, these two roles of the BG make it one of the most central brain structures for understanding both basic motor control and more advanced goal-driven planning and higher-level cognition. The circuits that enable it to play these critical roles are explored first in the case of basic motor control below, and then elaborated in terms of the relationships with the [[prefrontal cortex]] in that page.

One overarching way of understanding the role of the BG across all these levels, is in terms of a _bidirectional_ and _modulatory_ control system, that uses learning across two _opponent pathways_ to _disinhibit_ pathways through other neural circuits, in order to select and control the flow of activity through these other circuits. Thus, it is a kind of "puppet master" pulling the strings of the brain, to get it to do the "right thing" in order to optimize overall reward, and accomplish desired goals.

## Opponent pathways: D1 vs D2 / Go vs No

{id="figure_bg-rat" style="height:25em"}
![The basal ganglia in a rodent brain, showing the two major pathways: direct from the striatum to the GPi (globus pallidus internus) and SNr (substantia nigra pars reticulata) output pathways, and indirect that makes an additional hop through the GPe (globus pallidus externus). Figure from Gerfen & Surmeier, 2011.](media/fig_bg_anatomy_rat.png)

One of the most salient features of the BG circuitry, illustrated in the context of the anatomy of the BG within the rodent brain in [[#figure_bg-rat]], is that it has two distinct pathways from the striatum to the output areas of the _substantia nigra pars reticulata_ (SNr) and globus pallidus internus segment (GPi) (these two output areas are functionally identical for the present purposes). Because all of the major neurons in the BG are inhibitory (unlike the cortex, where the principal neurons are excitatory), these two pathways end up having opposing effects, a fact which has driven much of the theorizing about what the BG actually does.

Specifically, distinct subsets of SPN neurons in the striatum project either _direct_ to the GPi/SNr, or go through the _indirect_ pathway via _globus pallidus externus_ (GPe). Because the GPi/SNr output neurons are themselves inhibitory, the net effect of the direct pathway is to _disinhibit_ the areas that the BG projects to. Therefore, the direct pathway can be thought of as a "Go" pathway: it allows behavior to proceed. By contrast, the indirect pathway adds an additional minus sign in the chain, making it net _inhibitory_ on BG outputs, and thus a "No" pathway that tends to prevent behavior from proceeding.

You can compute these effects just by adding up the number of $-$ signs in the pathway: if it is even, then the net effect is positive / disinhibitory, and if it is odd, then the net effect is negative / inhibitory.

The simple _Go_ vs. _No_ logic of these two BG pathways aligns with the apparent deficits in people with Parkinson's disease, that can be characterized as a specific problem in _initiating_ motor actions (which was nicely illustrated in the movie _Awakenings_, where a patient could keep walking once they got started, but otherwise could be stuck for hours). This convergence led several authors to suggest that the primary function of the BG is in _action selection_: the decision of what action to perform ([[@ChevalierDeniau90]]; [[@AlexanderCrutcher90a]]; [[@GurneyPrescottRedgrave01]]; [[@FrankLoughryOReilly01]]), and that once this selection has been made, it can thus proceed without further input from the BG, explaining the selective initiation deficits in Parkinson's patients.

{id="figure_bg-act-sel" style="height:20em"}
![Conceptual model of the BG performing action selection, allowing only one selected action to proceed (via the direct Go pathway), while inhibiting the others via the indirect No pathway.](media/fig_bg_action_sel_dam.png)

This action selection model posits that the selected action gets a _Go_ disinhibitory signal from the BG, while all the other unselected actions get a _No_ inhibitory signal, as illustrated in [[#figure_bg-act-sel]]. The computational models of this function involved inhibitory connectivity among the SPN neurons to perform a classic _winner take all_ process ([[@GurneyPrescottRedgrave01]]; [[@FrankLoughryOReilly01]]). This model also has the advantage of distinguishing the contribution of the BG from that of the cerebellum, which is widely thought to be important for rapid online adjustments to motor control signals, and not for selection or initiation.

However, many experiments with animals suggested that the main function of the BG in motor control is actually the opposite of this action selection model: it is all about the online control of the motor action as it unfolds, and actually _not_ about the initiation or selection process (see [[@ParkCoddingtonDudman20]] for an extensive discussion). Another challenge to the simple action selection hypothesis is that there is only sparse, weak inhibitory connections among SPN neurons in the striatum, which seems unlikely to support a robust inhibitory competition and selection dynamic ([[@TunstallOorschotKeanEtAl02]]). Furthermore, the direct and indirect pathway neurons associated with a given action were shown to _both_ be strongly activated, instead of having fully opposite patterns of activity ([[@CuiJunJinEtAl13]]).

The following considerations address these issues and provide a coherent picture of BG function:

* The simple direct vs. indirect pathway story is inconsistent with further anatomical results, which instead point to a central role for the GPe within the BG circuit for weighing the balance of direct and indirect pathway inputs.
* Different areas of BG have different dynamics (dorsal vs ventral)
* Different output pathways (motor vs. thalamocoritcal) have different functions
* Different species have different balances of BG vs. cortical influence on motor control

{id="figure_pcore-bio" style="height:20em"}
![PCore model of the basal ganglia.](media/fig_pcore_v2_pr_ak.png)



