+++
Categories = ["Axon"]
bibfile = "ccnlab.json"
+++

**Rubicon** is a theory originally proposed by [[@^HeckhausenGollwitzer87]] that was independently developed and elaborated in [[@^OReillyHazyMollickEtAl14]] and [[@^OReilly20]] (see also [[@HerdKruegerNairEtAl21]]; [[@OReillyNairRussinEtAl20]]), which motivates a systems-neuroscience model of goal-driven learning and cognition in the [[Axon]] model.

The core idea that gives the theory its name is that a significant level of psychological / emotional investment accompanies the transition into the **goal engaged** state, which is analogous to Caesar crossing the _Rubicon_ river, thereby fully committing to a course of action, with recognition of the potentially grave consequences of doing so.

From a subjective, intuitive perspective, most of us have likely experienced significant decisions of this sort, in terms of major life choices with significant implications (e.g., where to go to college, whether to get married, have kids, etc). But the Rubicon theory takes a further step, arguing that _every_ decision to engage in a coordinated plan of action is associated with at least some level of psychological engagement.

One consequence of this engagement is that at least some level of _disappointment_, _frustration_, _sadness_, and/or  _anger_ would be experienced if the desired outcome of this plan is not accomplished. Indeed, the pervasive and universal nature of these emotional states, which are directly tied to goal failure, supports the reality of these goal-engaged states.

A further consequence of goal engagement is that it puts pressure on the **goal selection** process to be relatively conservative, because of the risk of experiencing these negative emotions upon goal failure. This can explain why people are often seemingly "irrationally reluctant" to start doing something, which is otherwise known as _procrastination_. You have likely experienced this phenomenon of dreading some effortful task like cleaning, writing a paper, or doing taxes, only to be pleasantly surprised at how easy it actually was relative to earlier expectations.

This discrepancy can be seen as a rational consequence of the commitment (and opportunity) costs associated with engaging in a given goal, and the nearly universal experience of procrastination provides further evidence of the pervasive nature of these goal-centric factors in shaping our daily lives.

From a computational perspective, the central premise of the Rubicon framework is that it is necessary for an agent to commit to a particular course of action in a prospective, outcome-directed manner, in order to accomplish relevant goals, and that there are two different value functions in effect for the goal selection vs. goal engaged states:

* The goal-engaged value function is weighted toward _accomplishing the currently-selected goal_, because that provides important benefits, the most obvious of which is that constantly switching "horses" mid-stream and generally dithering about in a non-committal fashion is unlikely to lead to the same kind of success that fully committing to a plan will. In particular, putting everything into a goal, and then failing (or succeeding), is the best way to learn. If you don't really try, and you fail, then there is a significant  [[credit assignment]] problem: did you fail because you didn't try, or because you need to adopt a different plan, or goal, next time?

* The goal-selection value function is weighted toward _conservatively evaluating the available goal options_, and picking the one that will produce the best benefit-to-cost ratio, taking into account all the relevant contextual factors (e.g., internal states such as hunger, and external states such as visible opportunities). As we often experience with procrastination, this means that you will tend to pick easier, short-term goals over longer, harder goals.

The proverb from the _Tao Te Ching_ summarizes the important issue of goal timescales:

> "A journey of a thousand miles begins with a single step"

The active goal selected at any point in time is more like the single step in a direction that is relevant for the long journey, rather than the entire journey itself, which is represented as a broader set of active _context_ information that informs and constrains individual goal selection steps. The timescale of goal engagement should be as brief as possible, while still resulting in a meaningful and measurable outcome, because this minimizes the costs of overcommitting. 

The widely-cited SMART goal-setting framework (specific, measurable, achievable, relevant, and time-bound ([[@Doran81]]) provides a good rule-of-thumb for what an active engaged goal should be, with specificity being critical for actually facilitating the transition to the active goal state, so there is an actual concrete action plan to guide behavior. This is consistent with the implementation intention if-then plans described by [[@^GollwitzerSheeran06]].

Thus, the core goal-driven system governs the _inner-loop_ of behavior, operating over relatively short time intervals of seconds to minutes, to actively guide behavior. Because this core goal system is evolutionarily ancient and supported by a number of subcortical brain areas working in concert with the neocortex, it tends to fly below the level of [[conscious awareness]]. The more conscious, deliberative aspects of goal selection constitute a more diffuse _outer-loop_ of context and constraints that also get activated and updated in the course of the inner-loop goal selection and engagement process, but are not directly guiding online behavioral choices in the way that the active goal does ([[@HerdKruegerNairEtAl21]]).

Thus, when you are sitting around thinking about what to do, you may engage in a number of these inner-loop goal-engaged steps in order to evaluate different possibilities. When you decide to take the first step of a larger journey, the active maintenance support associated with the transition to goal engagement also "lifts the boats" of the vaguer plans associated with these longer timescale plans, so they are carried along with the tide of inner-loop goal steps. [[@^Gollwitzer12]] terms these broader plans "wants" (and their even more vague predecessors, "wishes"), while the active inner-loop is "willing".

From a neuroscience perspective, these goal-related principles require neural mechanisms to somehow organize the overall dynamic of goal selection and engagement, and provide the goal selection process with good estimates of the cost / benefit tradeoffs for different options, etc. The emotional / motivational system is central, because what we experience subjectively as the emotional outcomes of goal success or failure are driven by brainstem circuits that have evolved to keep us oriented toward survival-relevant behaviors.

Based on extensive research, these brain systems involve the classical [[limbic system]] of areas, including the ventral and medial [[prefrontal cortex]], [[basal ganglia]], and [[amygdala]], and neuromodulatory systems including [[dopamine]].

The subsequent sections provide a more detailed overview of the relevant computational, cognitive, and neuroscience issues underlying the Rubicon framework. Further elaboration is then provided in the following pages:

* [[Limbic system]] provides a more detailed biological overview of the relevant brain areas, including the [[basal ganglia]] and [[prefrontal cortex]].
* [[PVLV]] is the biologically-based core of [[reinforcement learning]] mechanisms based on these brain areas, which accounts for a wide range of data on phasic [[dopamine]] firing and the animal conditioning literature.
* [[Arm maze simulation]] puts everything together in an integrated model of basic instrumental choice in a multi-arm decision making paradigm, with a simulated rodent-like agent.

## Computational overview

From the perspective of existing [[reinforcement learning]] models (this page is strongly recommended reading as a next step after this overview), the Rubicon framework implements various forms of [[reinforcement learning#model-based]] RL, including state abstraction via a predictive world model, and it can do planning in parallel using [[constraint satisfaction]] via [[bidirectional connectivity]], thereby mitigating the [[curse of dimensionality]] associated with serial planning dynamics.

The process of trial-and-error [[search]] through the space of actions in different environments, which defines the [[reinforcement learning]] problem, is plagued by this curse of dimensionality as the action and state spaces get more complex. Thus, as discussed extensively in the RL page, successful learning requires various ways of compressing the environmental state representations, and employing parallel search strategies at multiple levels, including constraint-satisfaction and [[optimized representations]] that are supported by the Axon mechanisms.

Furthermore, Rubicon represents an attempt to [[computational-cognitive-neuroscience#reverse engineer]] millions of years of evolution, by building in critical neural circuits that support the goal selection and maintenance process, based on learning that is focused on acquiring the information needed to make well-informed choices about which goals to pursue in a given context. This evolutionary foundation represents a significant departure from existing RL models, but we believe it is necessary to mitigate the curse of dimensionality problem as environments and actions become more complex.

{id="figure_rubicon-rl" style="height:20em"}
![Functional elements of the Rubicon computational model. Prior to having an engaged goal, exploration of possible goals occurs. A CS (conditioned stimulus) can signal the possibility of a desired outcome (US, unconditioned stimulus), and constraint satisfaction planning activates potential model and plan representations to achieve the outcome. If this overall plan / goal is above threshold, the entire distributed goal representation is gated and maintained in prefrontal cortex, to guide subsequent action selection and provide expectations for monitoring progress toward the goal.](media/fig_rl_state_action_reward_plan_rat.png)

The net effect of these goal-oriented mechanisms is to organize behavior around the following computational steps (illustrated in [[#figure_rubicon-rl]]):

1. In the absence of an existing engaged goal, explore the environment in a manner directed by efficient, abstract internal models, looking for cues suggesting potentially-useful outcomes that might be available (e.g., food or other resources).

2. For each such possible opportunity, use parallel constraint satisfaction processing to efficiently search through and synthetically activate learned representations of possible action plans required to obtain the desired outcome. Constraints involved in this processing include effort, risk, and uncertainty factors, which are weighed against the contextualized benefit of the outcome (e.g., influenced by level of current hunger). Different brain areas within the ventral and medial prefrontal cortex specialize on learning these different factors, and they all interact bidirectionally during this goal selection process.

3. The best available option is selected after an evaluative process that is modulated by urgency and other factors, which determine the amount of time and effort spent on the decision. This is when the system crosses the Rubicon into the goal-engaged state, and becomes focused on accomplishing the selected goal, rather than conservatively evaluating options. The estimates of benefits and costs, and the results of the planning process, are all _gated_ into a robustly-maintained active maintenance state in their respective prefrontal cortex areas, providing a distributed goal representation with critical information needed to guide the goal pursuit process.

4. The maintained action plan biases the selection of motor actions consistent with the plan ([[@MillerCohen01]]; [[@OReillyBraverCohen99]]), providing a dynamic policy guiding behavior, instead of a rigid, statically-learned one as in [[reinforcement learning#model-free]] RL. As action affordances arise, this plan can be updated via ongoing constraint satisfaction processing, leveraging parallel search through learned environment and action models to effectively narrow and focus the overt behavioral search on the most promising actions.

5. The maintained cost and benefit estimates provide active benchmarks against which ongoing progress is measured. If things are taking longer and becoming more uncertain than expected, with progress moving away instead of toward the anticipated outcome, this can trigger goal abandonment (i.e., the "give up" action). This comes with a critical learning signal that registers the failure to accomplish this goal, and updates the estimate representations accordingly. This "dopamine dip" is experienced as disappointment subjectively. 

    Instead of giving up in the face of unexpected difficulties, the system could expend additional effort to overcome these difficulties (i.e., _perseverance_, which can then turn into _perseveration_ if it goes on too long). The neuromodulator [[norepinepherine]] has been implicated in this switch between perseverance and giving up ([[@Aston-JonesCohen05]]). The decision about whether and when to give up, versus try harder, depends on relevant contextual factors, all informed by prior learning experiences. Considerable evidence suggests that the [[lateral habenula]] plays a critical role in making this decision.

6. When a desired outcome is achieved, then an appropriate positive reward signal (dopamine burst) is generated (internally, as a function of the outcome relative to the expectation), which updates estimates accordingly, and is experienced as happiness or satisfaction.

Thus, behavior is actively guided by the maintained distributed goal state, thereby providing a more coherent and yet dynamic mode of behavior, relative to the common practice of selecting actions using a softmax probabilistic model based on estimates of future value, as in the TD model. The consistency of action selection over time as guided by the active plan is critical for navigating high-dimensional spaces, where even a small noise probability can accumulate across multiple steps and effectively derail the ability to accomplish a goal requiring a sequence of goal-driven actions ([[@EcoffetHuizingaLehmanEtAl21]]).

The obvious cost of this approach is that committing to a specific goal may prevent the ability to opportunistically select another goal, thus putting a premium on the accuracy of the goal selection process. However, the key hypothesis is that the learning benefits of sticking with a given goal until it is accomplished (or abandoned) will be significant vs. constantly selecting actions according to local probabilistic estimates.

As discussed above, these commitment and opportunity costs can be mitigated if the system is biased toward engaging in the shortest timescale of goals that lead to a tangible outcome (i.e., something that can be clearly evaluated as having been accomplished or not). These short term goals can often function as subgoals toward larger, longer timescale goals, but the active goal selection and engagement dynamic specifically operates on this inner-loop of goal-driven behavior, where there is just one actively engaged goal state at a time.

If the timescale of the inner-loop shrinks down to just one action, it would appear that the system collapses to a standard RL framework. However, even if the inner-loop is relatively short, a key hypothesis is that the mechanism of conservative goal selection, and tracking of outcomes relative to expectations using active maintenance of the initial goal selection state, will serve to _bootstrap_ more strategic proactive behavior overall. As noted above, the longer timescale, outer-loop goals get carried along with the sequence of inner-loop steps, and are thus shaped by the proactive, conservative, outcome-oriented nature of these steps.

From a computational perspective, the ultimate challenge that this model must solve is to learn novel goal representations that are sufficiently [[combinatorial vs conjunctive|combinatorial]] as to support [[generalization|generative]], systematic behavior in novel environments. In short, we need an answer for how the system can synthesize a novel plan for something it has never done before, in advance of actually doing it. The working hypothesis is that by building in proactive, outcome-focused learning mechanisms that solve the [[credit assignment#temporal credit assignment]] problem, along with the optimized representational dynamics, the Rubicon framework will answer this central unsolved problem.

## Biological overview

{id="figure_goal-bio" style="height:20em"}
![Distributed goal representations across different prefrontal cortex (PFC) areas, each with associated areas of the striatum of the basal ganglia. The mediodorsal (MD) thalamus provides a common point of disinhibitory gating control, by which phasic updating of the PFC goal state can be controlled at the point of goal selection and after the goal outcome. OFC = orbitofrontal cortex; IL = infralimbic; ACC = anterior cingulate cortex; PL = prelimbic; dlPFC = dorsolateral PFC (isomorphic to ALM = anteriolateral motor cortex in rodents).](media/fig_rubicon_loops_spiral_goals.png)

The distributed representation of a _goal_ state in the ventral and medial [[prefrontal cortex]] (PFC) is shown in [[#figure_goal-bio]]. Each area has bidirectional connectivity with associated subcortical areas that directly encode specific goal-relevant parameters, and is interconnected with associated [[basal ganglia]] (BG) areas, as supported by a wide range of data ([[@AlexanderDeLongStrick86]]; [[@OngurPrice00]]; [[@SaddorisGallagherSchoenbaum05]]; [[@FrankClaus06]]; [[@RushworthBehrensRudebeckEtAl07]]; [[@SchoenbaumRoeschStalnakerEtAl09]]; [[@KouneiherCharronKoechlin09]]; [[@KennerleyBehrensWallis11]]; [[@PauliHazyOReilly12]]; [[@RudebeckMurray14]]; [[@RichWallis16]]; [[@HuntMalalasekeraBerkerEtAl17]]; [[@PezzuloRigoliFriston18]]).

* _Orbitofrontal cortex (OFC)_ encodes biologically and affectively salient outcomes, i.e., the _unconditioned stimulus (US)_ in conditioning theory. There is considerable evidence that the OFC maintains and tracks these US-like states ([[@RichWallis16]]), and OFC damage impairs the ability to adapt behavior to rapid changes in these US outcomes ([[@BalleineDickinson98]]). As in classical conditioning theory, these US representations ground the motivational system to pursue biologically-based [[emotion|needs]]. There is an extensive literature on _drives_ going back to [[@^Hull43]] and Maslow's hierarchy of needs ([[@Maslow43]]), which anchor the reward value of the USs, in part via subcortical projections that converge into the OFC [[@OngurPrice00]].

* _Infralimbic cortex (IL)_ provides a more abstract representation of US _value_ relative to the more detailed outcome representations in OFC.

* _Dorsolateral PFC (dlPFC)_ represents motor action plans at a high level, and strongly influences other brain areas in support of such plans ([[@MillerCohen01]]; [[@Desimone96]]; [[@OReillyBraverCohen99]]).

* _Anterior cingulate cortex (ACC)_ has extensive interconnectivity with corresponding regions of the dorsal PFC motor areas including the dlPFC, and encodes motor-triggered utility values such as the effort and uncertainty associated with different motor plans ([[@AlexanderBrown11]]; [[@ShenhavBotvinickCohen13]]).

* _Prelimbic cortex (PL)_ provides an integration of all of the above signals into an overall _utility_ value that encodes the outcome value (from IL) relative to the costs of the action plans required to obtain the outcome (from ACC).

These areas intercommunicate during the goal selection process via extensive bidirectional connectivity, supporting the [[constraint satisfaction]] process that performs efficient parallel [[search]] through the high-dimensional space of possible outcomes, plans, and their overall utility value, to converge on a reasonable distributed goal state that satisfies the constraints imposed by each of these areas, along with the relevant environment and internal state inputs ([[@HerdKruegerNairEtAl21]]).

The corresponding basal ganglia areas receive inputs from these frontal areas and learn to evaluate the extent to which the current active state represents a good vs. bad bet, in the goal selection process, driving disinhibition of downstream thalamic areas if the overall goal representation is above threshold. In particular, the mediodorsal (MD) nucleus of the thalamus has diverse projections into all of these goal-related frontal areas, and can provide a coordinated _gating_ signal that locks in robust active maintenance of the distributed goal state across these areas ([[@HerdKruegerNairEtAl21]]; [[@OReillyFrank06]])

TODO: parallel BG motor learning with distributed action representations and graded control signals, to do efficient parallel search.

<!--- TODO: svloboda cites -->

TODO: Amygdala, VS, LHb systems, PVLV, DA, drives, etc.

## Cognitive overview

In many animals, this inner loop of goal engagement may be related to the universal phenomenon of bouts and pauses in behavior ([[@KramerMcLaughlin01]]; [[@Shull11]]; [[@FalligantHagopianNewland24]]), which is evident by observing squirrels in a park or your back yard, for example. Interestingly, pigeons do not show the signatures of this behavior in instrumental conditioning tasks, but rodents robustly do ([[@FalligantHagopianNewland24]]), suggesting a possible role of more developed goal-related brain systems. This same characteristic is evident in people as well: we may not quite realize how often we lapse in and out of focus while performing a task, spending chunks of time mind wandering in between whatever else we are doing ([[@ShinagawaYamada25]]).


## Rubicon pages


