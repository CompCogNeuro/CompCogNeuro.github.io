+++
Categories = ["Rubicon", "Neuroscience"]
bibfile = "ccnlab.json"
+++

The **basal ganglia** (BG) are a collection of nuclei located at the central "core" of the brain ([[#figure_bg-anatomy]]), forming the functional bridge between the highest levels of processing in the [[neocortex]], and the motor control networks in the brainstem (midbrain and hindbrain, progressing down to the spinal cord). The primary outputs of the BG are downward to these motor networks, and, via extensive connections into the thalamus, back up to the frontal areas of the neocortex and [[prefrontal cortex]]. The **striatum** is the input nucleus of the BG, anatomically composed of the _caudate_ and _putamen_, which receives extensive projections from the neocortex and other brain areas.

{id="figure_bg-anatomy" style="height:25em"}
![Major areas of the basal ganglia (BG) and associated brain areas. The striatum (composed of the caudate and putamen as large-scale anatomical features) is the input layer of the BG, receiving a wide range of inputs from all over the brain, especially the neocortex. The nucleus accumbens (NAc), i.e,. the ventral striatum, is the input area for the goal-driven, affective circuits. The globus pallidus externus (GPe) is the next stage of processing, providing the central "core" integration for the BG, with the GP internus (GPi) sending BG outputs subcortically and, via extensive connections into the thalamus, back up to the frontal neocortex. The substantia nigra pars reticulata (SNr) is also a major output pathway, while the pars compacta (SNc) provides dopamine innervation into the entire BG. The subthalamic nucleus (STN) also receives extensive cortical input from frontal areas, and interacts bidirectionally with the GPe. The amygdala is not a part of the BG but interacts extensively with it and is located at the extreme end of the tail, and the lateral habenula (LHb) plays an essential role in driving dips in dopamine, based on extensive inputs from the ventral striatum among many other areas.](media/fig_bg_anatomy_human.png)

The SPNs (spiny projection neurons; also called medium spiny neurons, MSNs) in the striatum receive the strongest [[dopamine]] inputs, and have the greatest ability to rapidly clear dopamine from the synapse after it has been released, of any area in the brain. This allows these neurons to learn from rapid _phasic_ changes in dopamine, i.d., _bursts_ and _dips_ relative to the baseline _tonic_ firing level, in a manner consistent with the core principles of [[reinforcement learning]] (RL; [[@MarkowitzGillisJayEtAl23]]; [[@HowardLiGeddesEtAl17]]; [[@NairGutierrez-ArenasErikssonEtAl15]]; [[@ShenFlajoletGreengardEtAl08]]; [[@Frank05]]).

There is extensive evidence showing that the BG encodes detailed aspects of motor control, including impressive decoding of complex naturalistic motor behaviors in freely-behaving rats ([[@MarkowitzGillisBeronEtAl18]]; [[@KlausMartinsPaixaoEtAl17]]), in addition to simpler conditioned learning tasks ([[@YttriDudman18]]). Furthermore, studies of animals with complete lesions of the neocortex demonstrate that the BG can drive remarkably intact fine-grained motor control through its descending motor pathway projections ([[@GrillnerRobertsonKotaleski20]]; [[@ParkCoddingtonDudman20]]; [[@ArberCosta22]]). In evolutionarily more ancient vertibrates without much cortex, such as the lamprey (which has essentially the same BG structures as the rodent and primate; [[@GrillnerRobertson16]]), the BG is the primary driver of learning and behavior.

Thus, the striatum is widely believed to be the neural substrate for the _actor_ component of the brain's reinforcement learning system, where reward prediction error (RPE) signals from the dopamine _critic_ system train the BG actor to improve motor actions to increase overall reward (see [[reinforcement learning#figure_actor-critic]] in the RL page). 

However, in animals with more neocortex, especially in primates and humans, the BG also works intimately in concert with the cortex to organize behavior, by way of its extensive connections into the thalamus. These thalamic projections can provide two functions: direct modulation of cortical activity, which occurs in a very broad manner (which we term **gating**), and more fine-grained training of cortical learning through the same type of thalamic-driven [[predictive learning]] signals used in the [[Axon]] [[error-driven learning]] model. Thalamic gating is a central mechanism in the [[Rubicon]] framework, for locking in the active maintenance of distributed goal representations across various regions of the prefrontal cortex.

Thus, these two roles of the BG make it one of the most central brain structures for understanding both basic motor control and more advanced goal-driven planning and higher-level cognition. The circuits that enable it to play these critical roles are explored first in the case of basic motor control below, and then elaborated in terms of the relationships with the [[prefrontal cortex]] in that page.

One overarching way of understanding the role of the BG across all these levels, is in terms of a _bidirectional_ and _modulatory_ control system, that uses learning across two _opponent pathways_ to _disinhibit_ pathways through other neural circuits, in order to select and control the flow of activity through these other circuits. Thus, it is a kind of "puppet master" pulling the strings of the brain, to get it to do the "right thing" in order to optimize overall reward, and accomplish desired goals.

## Opponent pathways: D1 vs D2 / Go vs No

{id="figure_bg-rat" style="height:25em"}
![The basal ganglia in a rodent brain, showing the two major pathways: direct from the striatum to the GPi (globus pallidus internus) and SNr (substantia nigra pars reticulata) output pathways, and indirect that makes an additional hop through the GPe (globus pallidus externus). Figure from Gerfen & Surmeier, 2011.](media/fig_bg_anatomy_rat.png)

One of the most salient features of the BG circuitry, illustrated in the context of the anatomy of the BG within the rodent brain in [[#figure_bg-rat]], is that it has two distinct pathways from the striatum to the output areas of the _substantia nigra pars reticulata_ (SNr) and globus pallidus internus segment (GPi) (these two output areas are functionally identical for the present purposes). Because all of the major neurons in the BG are inhibitory (unlike the cortex, where the principal neurons are excitatory), these two pathways end up having opposing effects, a fact which has driven much of the theorizing about what the BG actually does.

Specifically, distinct subsets of SPN neurons in the striatum project either _direct_ to the GPi/SNr, or go through the _indirect_ pathway via _globus pallidus externus_ (GPe). Because the GPi/SNr output neurons are themselves inhibitory, the net effect of the direct pathway is to _disinhibit_ the areas that the BG projects to. Therefore, the direct pathway can be thought of as a "Go" pathway: it allows behavior to proceed. By contrast, the indirect pathway adds an additional minus sign in the chain, making it net _inhibitory_ on BG outputs, and thus a "No" pathway that tends to prevent behavior from proceeding.

You can compute these effects just by adding up the number of $-$ signs in the pathway: if it is even, then the net effect is positive / disinhibitory, and if it is odd, then the net effect is negative / inhibitory.

The simple _Go_ vs. _No_ logic of these two BG pathways aligns with impairments in people with Parkinson's disease, which can be characterized as a specific problem in _initiating_ motor actions (which was nicely illustrated in the movie _Awakenings_, where a patient could keep walking once they got started, but otherwise could be stuck for hours unable to start). This convergence led several authors to suggest that the primary function of the BG is in _action selection_: the decision of what action to perform ([[@ChevalierDeniau90]]; [[@AlexanderCrutcher90a]]; [[@GurneyPrescottRedgrave01]]; [[@FrankLoughryOReilly01]]). Once this selection has been made, it can thus proceed without further input from the BG, explaining the selective initiation deficits in Parkinson's patients.

{id="figure_bg-act-sel" style="height:20em"}
![Conceptual model of the BG performing action selection, allowing only one selected action to proceed (via the direct Go pathway), while inhibiting the others via the indirect No pathway.](media/fig_bg_action_sel_dam.png)

This action selection model posits that the selected action gets a _Go_ disinhibitory signal from the BG, while all the other unselected actions get a _No_ inhibitory signal, as illustrated in [[#figure_bg-act-sel]]. The computational models of this function involved inhibitory connectivity among the SPN neurons to perform a classic _winner take all_ process ([[@GurneyPrescottRedgrave01]]; [[@FrankLoughryOReilly01]]). This model also has the advantage of distinguishing the contribution of the BG from that of the cerebellum, which is widely thought to be important for rapid online adjustments to motor control signals, and not for selection or initiation.

However, many experiments with animals suggested that the main function of the BG in motor control is actually the opposite of this action selection model: it is all about the online control of the motor action as it unfolds, and actually _not_ about the initiation or selection process (see [[@ParkCoddingtonDudman20]] for an extensive discussion). Another challenge to the simple action selection hypothesis is that there is only sparse, weak inhibitory connectivity among SPN neurons in the striatum, which seems unlikely to support a robust inhibitory competition and selection dynamic ([[@TunstallOorschotKeanEtAl02]]). Furthermore, the direct and indirect pathway neurons associated with a given action were shown to _both_ be strongly activated, instead of having fully opposite patterns of activity ([[@CuiJunJinEtAl13]]).

The following considerations address these issues to provide a coherent picture of BG function:

* The simple direct vs. indirect pathway story is inconsistent with more recent anatomical data, which instead point to a central role for the GPe within the BG circuit for weighing the balance of direct and indirect pathway inputs. Both the direct and indirect pathways project strongly into the GPe, which also has multiple distinct neuron subtypes. When all of this is included in a model, it provides a different way of understanding how the system operates, and what it is best suited to doing.

* Different areas of BG with distinct functional roles, e.g., dorsal lateral areas for motor control, vs medial and ventral areas for goal selection, have different patterns of connectivity and neural dynamics. The result is that dorsal lateral areas can play a more online control function, while ventral and medial areas can drive a more phasic goal selection process that leads to the initiation of actions.

* Different output pathways (motor vs. thalamocortical) likewise have different connectivity and functionality, with descending pathways having more focal projections to different brainstem motor areas, while ascending pathways through the thalamus have broader connectivity that is suggestive of a more modulatory role.

* Different species have different balances of BG vs. cortical influence on motor control, with primates and humans having a much stronger degree of cortical influence over motor control, while rodents and other species with less cortical development are more driven by the BG.

## The Pallidal-core (PCore) model

Although there were earlier indications regarding the inaccuracy of the "classical" direct vs. indirect pathway model, relatively recent molecular labeling techniques have now provided definitive evidence for a new anatomical model, which puts the GPe in a more central role in shaping the dynamics of the BG ([[@CourtneyPamukcuChan23]] (review); [[@MalletMicklemHennyEtAl12]]; [[@SaundersMacoskoWysokerEtAl18]]; [[@CuiDuChangEtAl21]]; [[@GuilhemsangMallet24]]).

{id="figure_pcore-bio" style="height:25em"}
![PCore model of the basal ganglia, which is centered around the multiple projections into and out of the GPe neurons that affect every other part of the BG circuitry, putting the GPe Pallidum at the core of its function. The GPeAk (arkypallidal) neurons receive from the direct pathway striatal neurons (dSPN), while the prototypical (GPePr) neurons receive from the indirect pathway, as in the classical model (hence the name). Because GPeAk projects inhibition back up to the striatum, it must be inhibited in order to disinhibit the SPN neurons, which is accomplished by the direct pathway inputs. The iSPN neurons can also get some relief by inhibiting the GPeAk in cases where they are more active, and are directly inhibiting the dSPNs (but not the other way around). The hyperdirect pathway into the STN drives initial "brakes" on the system preventing premature responding.](media/fig_pcore_v2_pr_ak.png)

Our implementation of this new circuitry is summarized in [[#figure_pcore-bio]], showing two of the main  subtypes of GPe neurons: _GPeAk_ are the _arkypallidal_ GPe neurons, which express the molecular markers NPAS1 and FOXP2, while the GPePr are the _prototypical_ GPe neurons that have a connectivity pattern similar to the GPe in the classical model. As shown in the figure, roughly 45% of the GPe neurons are prototypical, while 18% are arkypallidal, with another 12% projecting to the SNc dopamine area (similar to [[#striosome]] neurons in the striatum, which we discuss below). The remaining neurons constitute a more heterogenous group, which we ignore for the time being.

The functional division in the striatum remains largely as in the classical model, with the direct-pathway SPN neurons (dSPN) that predominantly express the D1 dopamine receptor, and the indirect pathway iSPN neurons with predominantly D2 receptors. These two types of dopamine receptor are synergistic with the functional opposition of the Go (dSPN) vs. No (iSPN) pathways, with D1 having excitatory and LTP-promoting effects with bursts of dopamine, while D2 is inhibitory and promotes LTD with these dopamine bursts ([[@ShenFlajoletGreengardEtAl08]]; [[@Frank05]]). The opposite pattern holds for dopamine dips, directly implementing both sides of Thorndike's _law of effect_ for instrumental conditioning: do more of things that give you dopamine bursts, and less of things that give you dips in dopamine.

The one new wrinkle in the striatum is that the lateral inhibition among the SPN neurons is strongly asymmetric, with iSPN inhibiting dSPN but not the other way around ([[@TavernaIlijicSurmeier08]]), which is compatible with the remainder of the dynamics from the GPe.

Unlike the classical model, the GPe receives significant input from the direct pathway dSPN neurons, which strongly favors the GPeAk arkypallidal neurons. Meanwhile, the prototypical GPe neurons are so-named because they almost exclusively receive input from the iSPN neurons, as in the classical model. The GPeAk neurons also diverge from the classical model by projecting back up to the striatum, which provides a key insight into their function. Meanwhile, the prototypical GPe neurons inhibit themselves and the GPeAk neurons, while also sending the classically-described inhibitory projection to the output nuclei (SNr and GPi).

Finally, the STN (subthalamic nucleus) plays a critical role in the BG circuit (and is a major target of therapeutic treatments in Parkinson's disease), supporting the _hyperdirect_ projection from the cortex and uniquely sending excitatory glutamatergic projections to the GPe and SNr / GPi neurons. This set of connections has long been recognized as important for providing an initial "brake" on any disinhibitory effects of the BG, by driving an initial burst of excitation to the SNr and GPi outputs that are inhibiting the downstream targets of the BG.

In the PCore model, the STN projections into the GPePr are critical for driving a reciprocal inhibitory reflection back from GPePr to STN, which opens up a window where the SNr / GPi outputs can actually become net inhibited, because the STN neurons go into an extended pause in firing after their initial burst of activity ([[@FujimotoKita93]]; [[@MagillSharottBevanEtAl04]]).

{id="figure_pcore-dyn" style="height:30em"}
![Dynamics of the PCore model for a Go > No case (left panel) vs a No > Go case (right panel), illustrating the central role of the GPeAk neurons. Each layer shows a raster plot of spikes across the 25 neurons per layer, with time going back in depth for each layer. Thus, you can see the initial burst of activity in the STN driven by the hyperdirect pathway inputs, which puts on the "brake" at the start, allowing the rest of the dynamics to unfold, as the brake is released by GPePr inhbiting the STN. See text for further explanation of each of the steps highlighted. The striatum neurons are labeled as Mtx, representing the matrix (vs. striosome) subset, with dSPN = Go and iSPN = No.](media/fig_pcore_v2_vs_sim_go_no.png)

Figure [[#figure_pcore-dyn]] shows the PCore model in action ([[BGventral simulation]]) in cases where it ends up being net disinhibitory ("Go") and net inhibitory ("No"). For the Go case, the key steps are:

0. The STN hyperdirect activity provides a burst of activation to the SNr, GPePr, and GPeAk neurons, effectively preventing the SNr output from being inhibited. This is the initial "brake" on the system, which is released when the GPePr inhibits the STN in turn, and the [[neuron channels#SKCa]] calcium-gated K channels produce a longer-lasting inhibitory pause, providing a window for the BG to control the output pathways (i.e., a _gating window_).

1. Stronger learned weights from the "ACC" inputs (which are just clamped input layers in this model) to the dSPN cause these neurons to respond more vigorously than the opponent iSPN neurons.

2. This dSPN activity directly inhibits the GPeAk, initiating a "positive" feedback loop that facilitates further overall "Go" mode activity.

3. The reduced GPeAk activity disinhibits the striatum neurons (both dSPN and iSPN), which allows the dSPN pathway to ramp up further, although this is also tempered by increased iSPN pathway firing which directly inhibits the dSPN cells.

4. The increased striatal dSPN activity inhibits the SNr neurons...

5. Thereby accomplishing the classical Go pathway dynamic of disinhibiting the downstream targets of the BG, in this case the thalamus.

The case illustrated on the right shows what happens when the iSPN _No_ pathway gets more excited by the inputs.

1. Now the iSPN neurons are more active initially.

2. Which inhibits the GPePr, thereby disinhibiting the GPeAk.

3. Thus the continued GPeAk activity inhibits the striatum, preventing the dSPN Go pathway from getting more active.

4. Therefore, the SNr remains active...

5. And the thalamus remains inhibited by the SNr.

{id="figure_ketzef-etal" style="height:30em"}
![Results of optogenetic selective activation of iSPN (a) or STN (b) neurons. Activating ISPNs inhibited GPePr ("Proto") neurons, while disinhibiting GPeAk ("Arky") neurons. Activating the STN much more strongly activated GPePr vs GPeAk neurons, consistent with a stronger projection to GPePr.](media/fig_ketzef_etal_21_ispn_stn_stim.png)

These dynamics are consistent with various recorded patterns of neural activity, for example in a recent optogenetic stimulation experiment ([[@KetzefSilberberg21]]) that selectively activated iSPN neurons or STN neurons ([[#figure_ketzef-etal]]).

### Functional benefits of PCore

At a broad-brushstroke level of analysis, the PCore model exhibits similar behavior to the classical BG model, with a strong opponent dynamic between a net Go vs. a net No pathway. However, there are key properties of this circuit that make it considerably more robust than the classical model. 

Perhaps the most important feature is that that there are multiple ways in which these opposing pathways also have offsetting effects, which keep the overall dynamics in a more evenly balanced regime overall. For example, the iSPN inhibition of dSPNs balances the disinhibitory positive feedback loop from GPeAk to the dSPNs. Likewise, the self-inhibtion by the GPePr neurons, and their inhibition of the GPeAk neurons, provide offsetting effects.

Overall, we find in our model that these dynamics allow the system to apply a ratio-based decision threshold across a wide range of raw input activation strengths, consistently exhibiting disinhibitory Go dynamics when the relative strength of initial Go vs. No pathway activation favors Go. In addition, as the Go vs. No balance gets closer, the system takes longer to make a decision, allowing more time for additional input signals to be integrated and improve the overall quality of the resulting decision.

TODO: show figs from BGventral model. Fix urgency!

Also, the GPeAk disinhibition of striatum strongly predicts that there will be some correlation in dSPN and iSPN firing, even though they are in an overall oppositional relationship, consistent with observed data ([[@CuiJunJinEtAl13]]). However, the oppositional relationship is still evident overall, consistent with the properties of dSPN and iSPN firing in naturalistic behavioral contexts ([[@MarkowitzGillisBeronEtAl18]]; [[@KlausMartinsPaixaoEtAl17]]).

Overall, several features of the PCore circuit produce behavior consistent with an an action-initiation role, including the brake-then-release dynamics in the STN, and the progressive positive feedback loops between dSPN and GPeAk. Nevertheless, the balanced dynamics should also support generating more graded, balanced combinations of net disinhibitory tone, consistent with an ability to provide ongoing modulatory control over motor actions as they unfold.

## Functional contributions of the BG

Thus, we return to the big picture question of what the BG is actually doing, now that we understand a bit more about how it seems to function at a circuit dynamics level.

* massive compression of inputs (reference numbers of neurons): still enough degrees of freedom to provide detailed control over different motor pathways and muscle groups -- 13K still way more than number of distinct muscle groups!

* major problem with action selection: how is the "menu" represented and how is this mapped through with parallel pathways for each possible action, which is necessary for the standard action selection logic.

* multiple different parallel loops through BG wrt to cortex and also downstream motor areas: different dynamics in different loops?

* Broader feedback from FSi etc: dorsal bg.

## Functional organization of the thalamus

Rovo et al.

connectivity of different thalmic areas.

broader connectivity of BG loops etc.

