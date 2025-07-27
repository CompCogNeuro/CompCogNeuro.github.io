+++
Categories = ["Rubicon", "Neuroscience"]
bibfile = "ccnlab.json"
+++

The **PVLV** (primary value, learned value) model provides a biologically-detailed account of how different brain systems drive phasic dopamine firing in response to CSs (conditioned stimuli) trained by USs (unconditioned stimuli) across a range of different classical conditioning paradigms ([[@OReillyFrankHazyEtAl07]]; [[@HazyFrankOReilly10]]; [[@MollickHazyKruegerEtAl20]]). It provides a bridge from purely computational algorithms such as the TD (temporal differences) algorithm discussed in detail in [[reinforcement learning]] to the neural substrates of the [[limbic system]] that are at the heart of the [[Rubicon]] goal-driven framework.

The [[PVLV simulation]] provides an in-depth exploration of the key mechanisms of the framework, while this page discusses the main neural data behind it, and how it fits in within the larger Rubicon model. The updated version of the model implemented within the [[Axon]] spiking framework provides a more biologically-based implementation of several aspects of the most recent published model from [[@^MollickHazyKruegerEtAl20]], especially with respect to the way that the orbitofrontal cortex (OFC) component is activated to maintain expectations of subsequent goal outcomes, as part of the Rubicon goal-engaged maintenance network in the [[prefrontal cortex]]. In addition, the contributions of the [[lateral habenula]] and the neuromodulator [[acetylcholine]] have expanded from the original model, as has the model of [[emotion|drives]] underlying USs. 

Overview of two pathways and contributions, etc.


## Sign vs. goal trackers

## TODO: New research

### VP

* [[@FujimotoHoriNagaiEtAl19]] -- VP shows activity earlier than rmCD (NAcc) inputs: VP 115+/- 17ms,PANs 190+/- 25ms! VP encodes reward size and satiation level. VP inactivation disrupts. OFC -> rmCD encodes value: 

> (Yin et al., 2005; Gremel and Costa, 2013; Gremel et al., 2016); we thereby confirmed the role of rmCD for mediating goal-directed behavior.  First, the coding latency of VP neurons was sig- nificantly shorter than that of rmCD pro- jection neurons (PANs). Second, if this were the case, then response polarity would be opposite between the two areas, considering that the projection from rmCD to VP is GABAergic (Haber et al., 1990). However, a majority of the neurons in both areas showed excitatory response to cue. 

> A remaining question is as follows: where is such rich and rapid incentive-value information derived from? One possible source is the basolateral amygdala (BLA), which has a reciprocal connection to VP (Mitrovic and Napier, 1998; Root et al., 2015) and is known to contain neurons reflecting incentive value of cue with short latency (Paton et al., 2006; Belova et al., 2008; Jenison et al., 2011). Recent studies demonstrated that amygdala lesion impaired reward-based learning more severely than VS lesion in monkeys (Averbeck et al., 2014; Costa et al., 2016), supporting the contribution of BLA-VP projection in incentive-value pro- cessing. Another candidate is the projection from the subtha- lamic nucleus (STN). VP has a reciprocal connection with the medial STN that receives projections from limbic cortical areas (Haynes and Haber, 2013), composing the limbic cortico- subthalamo-pallidal “hyperdirect” pathway (Nambu et al.,

In general, this is compatible with standard story: gets early amyg and STN inputs and starts to modulate, but key Go / No signal still comes from striatum.

* [[@AhrensMeyerFergusonEtAl16]] -- VP shows sign trackers strongly activating VP at time of CSs onset, while goal trackers do not. -- Long trials: 7-8s -- key for not bothering to maintain CS.  Both show same large responses to CS onset. fig_vp_sign_track_cs_ahrens_etal_16.png

* [[@SagaRichardSgambato-FaureEtAl16]] -- VP CS activity for appetitive and aversive in monkey:

> We found 2 populations of neurons that were preferentially activated by appetitive and aversive conditioned stimuli (CSs). In addition, VP showed appetitive and aversive outcome anticipatory activities. These activity patterns indicate that VP is involved in encoding and maintaining CS-induced aversive contextual information. Furthermore, the disturbance of VP activity by bicuculline injection increased the number of error trials in aversive trials. In particular, the subjects released the response bar prematurely, showed no response at all, or failed to avoid the aversive outcome. 

* [[@TachibanaHikosaka12]] -- appetitive CS in VP

* [[@FagetZellSouterEtAl18]] -- opponent coding in VP

* [[@OttenheimerBariSutliefEtAl20]] -- VP shows RPE signatures, not very consistent with TD. Sucrose more rewarding that maltodextrin (MD); RPE when getting MD; subset of neurons show sensitivity to prior history consistent with basic RW RPE model.  Similar to [[@FujimotoHoriNagaiEtAl19]], show better decoding in VP vs. NAc, but VP has high firing rates and integrates -- similar to Barter et al -- GPi is always going to have a better signal, but NAc is key for learning..

VP key for task engagement.


* [[@FujimotoHoriNagaiEtAl19]] -- VP
