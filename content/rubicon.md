+++
Categories = ["Axon"]
bibfile = "ccnlab.json"
+++

**Rubicon** is a theory originally proposed by [[@^HeckhausenGollwitzer87]] that was independently developed and elaborated in [[@^OReillyHazyMollickEtAl14]] and [[@^OReilly20]], which motivates a systems-neuroscience model of goal-driven learning and cognition in the [[Axon]] model.

The core idea that gives the theory its name is that a significant level of psychological / emotional investment accompanies the transition into the **goal engaged** state, which is analogous to Caesar crossing the _Rubicon_ river, thereby fully committing to a course of action, with recognition of the potentially grave consequences of doing so.

From a subjective, intuitive perspective, most of us have likely experienced significant decisions of this sort, in terms of major life choices with significant implications (e.g., where to go to college, whether to get married, have kids, etc). But the Rubicon theory takes a further step, arguing that _every_ decision to engage in a coordinated plan of action is associated with at least some level of psychological engagement.

One consequence of this engagement is that at least some level of _disappointment_, _frustration_, _sadness_, and/or  _anger_ would be experienced if the desired outcome of this plan is not accomplished. Indeed, the very pervasive and universal nature of these emotional states, which are directly tied to goal failure, supports the reality of these goal-engaged states.

A further consequence of goal engagement is that it puts more pressure on the **goal selection** process to be relatively conservative, because of the risk of experiencing these negative emotions upon goal failure. This can explain why people are often seemingly "irrationally reluctant" to start doing something, which is otherwise known as _procrastination_. You have likely experienced this phenomenon of dreading some effortful task like cleaning, writing a paper, or doing taxes, only to be pleasantly surprised at how easy it actually was relative to earlier expectations.

This discrepancy can be seen as a rational consequence of the commitment (and opportunity) costs associated with engaging in a given goal, and the nearly universal experience of procrastination provides further evidence of the pervasive nature of these goal-centric factors in shaping our daily lives.

From a computational perspective, the central premise of the Rubicon framework is that it is necessary for an agent to commit to a particular course of action over some non-trivial time interval, in order to accomplish relevant goals, and that there are two different value functions in effect for the goal selection vs. goal engaged states:

* The goal-engaged value function is weighted toward accomplishing the currently-selected goal, because that provides important benefits beyond the basic intuitive idea that constantly switching "horses" mid-stream and generally dithering about in a non-committal fashion is unlikely to lead to the same kind of success that fully committing to a plan will. In particular, putting everything into a goal, and then failing (or succeeding), is the best way to learn. If you don't really try, and you fail, then there is a significant  [[credit assignment]] problem: did you fail because you didn't try, or because you need to adopt a different plan, or goal, next time?

* The goal-selection value function is weighted toward conservatively evaluating the available goal options, and picking the one that will produce the best benefit-to-cost ratio, taking into account all the relevant contextual factors (e.g., internal states such as hunger, and external states such as visible opportunities). As we often experience with procrastination, this means that you will tend to pick easier, short-term goals over longer, harder goals.

From a neuroscience perspective, these goal-related principles require neural mechanisms to somehow organize the overall dynamic of goal selection and engagement, and provide the goal selection process with good estimates of the cost / benefit tradeoffs for different options, etc. The emotional / motivational system is central, because what we experience subjectively as the emotional outcomes of goal success or failure are driven by brainstem circuits that have evolved to keep us oriented toward survival-relevant behaviors.

Based on extensive research, these brain systems involve the classical [[limbic system]] of areas, including the ventral and medial [[prefrontal cortex]], [[basal ganglia]], and [[amygdala]], and neuromodulatory systems including [[dopamine]].

The following sections provide a more detailed overview of the relevant computational, cognitive, and neuroscience issues underlying the Rubicon framework. The core learning and goal selection, maintenance mechanisms are first integrated in the context of the [[PVLV]] model of phasic dopamine firing (i.e., classical conditioning), and then developed further in the [[arm maze simulation]] of basic instrumental choice.

## Computational issues

From the perspective of existing [[reinforcement learning]] models, the Rubicon framework implements various forms of [[reinforcement learning#model-based]] RL, including state abstraction via a predictive world model, and it can do planning in parallel using [[constraint satisfaction]] via [[bidirectional connectivity]], thereby mitigating the [[curse of dimensionality]] associated with serial planning dynamics.

The process of trial-and-error [[search]] through the space of actions in different environments, which defines the [[reinforcement learning]] problem, is plagued by this curse of dimensionality as the action and state spaces get more complex. Thus, as discussed extensively in the RL page, successful learning requires various ways of compressing the environmental state representations, and employing parallel search strategies at multiple levels, including constraint-satisfaction and [[optimized representations]] that are supported by the Axon mechanisms.

Furthermore, Rubicon represents an attempt to reverse engineer millions of years of evolution, by building in critical neural circuits that support the goal selection and maintenance process, based on learning that is focused on acquiring the information needed to make well-informed choices about which goals to pursue in a given context. This evolutionary foundation represents a significant departure from existing RL models, but we believe it is necessary to mitigate the curse of dimensionality problem as environments and actions become more complex.

The net effect of these goal-oriented mechanisms is to organize the model's behavior around the following computational steps:

1. In the absence of an existing engaged goal, explore the environment in a manner directed by efficient, abstract internal models, looking for any indication of potentially-useful outcomes that might be available (e.g., food or other resources).

2. For each such possible opportunity, use parallel constraint satisfaction processing to efficiently search through and synthetically activate learned representations of possible action plans required to obtain the desired outcome. Constraints involved in this processing include effort, risk, and uncertainty factors, which are weighed against the contextualized benefit of the outcome (e.g., influenced by level of current hunger). Different brain areas within the ventral and medial prefrontal cortex specialize on learning these different factors, and they all interact bidirectionally during this goal selection process.

3. The best available option is selected after an evaluative process that is modulated by urgency and other factors, which determine the amount of time and effort spent on the decision. This is when the system crosses the Rubicon into the goal-engaged state, and becomes focused on accomplishing the selected goal, rather than conservatively evaluating options. The estimates of benefits and costs, and the results of the planning process, are all _gated_ into a robustly-maintained active maintenance state in their respective prefrontal cortex areas, providing a distributed goal representation with critical information needed to guide the goal pursuit process.

4. The maintained action plan biases the selection of motor actions consistent with the plan, providing a dynamic policy guiding behavior, instead of a rigid, statically-learned one. As action affordances arise, this plan can be updated via ongoing constraint satisfaction processing, leveraging parallel search through learned environment and action models to effectively narrow and focus the overt behavioral search on the most promising actions.

5. The maintained cost and benefit estimates provide active benchmarks against which ongoing progress is measured. If things are taking longer and becoming more uncertain than expected, with progress moving away instead of toward the anticipated outcome, this can trigger goal abandonment (i.e., the "give up" action). This comes with a critical learning signal that registers the failure to accomplish this goal, and updates the estimate representations accordingly. This "dopamine dip" is experienced as disappointment subjectively. However, unexpected difficulties can also trigger the expenditure of additional effort to overcome these difficulties, depending on relevant contextual factors, all informed by prior learning experiences.

6. When a desired outcome is achieved, then an appropriate positive reward signal (dopamine burst) is generated (internally, in consideration of the outcome relative to the expectation), which updates estimates accordingly, and is experienced as happiness or satisfaction.

<!--- TODO: figure! -->

Thus, unlike may existing RL models, behavior is actively guided by the maintained distributed goal state, thereby providing a more coherent and yet dynamic mode of behavior, relative to the common practice of selecting actions using a softmax probabilistic model based on estimates of future value, as in the TD model. The consistency of action selection over time as guided by the active plan is critical for navigating high-dimensional spaces, where even a small noise probability can accumulate across multiple steps and effectively derail the ability to accomplish a goal requiring a sequence of goal-driven actions.

The obvious cost of this approach is that committing to a specific goal may prevent the ability to opportunistically select another goal, thus putting a premium on the accuracy of the goal selection process. However, the key hypothesis is that the learning benefits of sticking with a given goal until it is accomplished (or abandoned) will be significant vs. constantly selecting actions according to local probabilistic estimates.

TODO:

* short goals, maintained context handles longer-term goals by biasing shorter term selections.
* 


The central premise of Rubicon is that _accomplishing desired outcomes (goals) requires a sustained focus of cognition and behavior over time._

This may sound obvious, but it has major implications, and is not widely deployed in current AI / machine learning [[reinforcement learning]] models, because it requires solving some of the most pressing unsolved problems in [[Computational Cognitive Neuroscience]] (in the opinion of the author at least).

One of the main such problems is:

* How can a system learn novel goal representations? How can you possibly synthesize a novel plan for something you've never done before? If you can't do that, then you can never do anything new!

There have been various proposals for solving this problem, which involve various forms of _chunking_ (combining basic action elements into new chunked elements, and so on; e.g., [[@SuttonPrecupSingh99]]; [[@AndersonLebiere98]]), and other heuristic, clustering-based approaches ([[@CollinsFrank13]]; [[@SchapiroRogersCordovaEtAl13]]). 



Finally, the defining feature of the [[Rubicon]] model is that the system makes an explicit commitment to a given goal and plan of action at any given point in time, which magnifies all of the above effects by virtue of the strong shared focus throughout the neocortical network on this one goal and plan. Furthermore, making such a commitment requires a goal selection process with neural mechanisms that support the essential "mental time travel" process that activates relevant potential outcomes and plans at an earlier "proactive" step in the process, which effectively combines backward and forward planning into one big constraint satisfaction process at the point of goal selection.


## Rubicon pages


