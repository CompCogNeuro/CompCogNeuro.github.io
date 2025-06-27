+++
Categories = ["Rubicon", "Computation", "Learning"]
bibfile = "ccnlab.json"
+++

**Reinforcement learning (RL)** is the widely-used term for more abstract, machine-learning and [[abstract neural network]] models that learn based on _trial and error_ exploration and reward / punishment signals. The [[Rubicon]] framework provides a biologically-detailed implementation of RL-level functionality in the context of the [[Axon]] model.

{id="figure_rl-setup" style="height:20em"}
![Elements of the reinforcement learning framework: The Environment provides States to the Agent over time, which in turn takes Actions that influence the state evolution. Rewards (and punishments) are delivered by the Environment under specific conditions, and the Agent's presumed job is to learn to take Actions that optimize the receipt of reward (and minimize punishments) over time. Learning the direct mapping of State → Action according to a learned Policy is known as model-free RL. Most RL models largely ignore the internal state of the agent itself, despite its obvious importance in shaping animal behavior.](media/fig_rl_state_action_reward_rat.png)

The major elements of an RL model are shown in [[#figure_rl-setup]], in the context of a rat **agent** that takes **actions** in an **environment** that delivers specific **state** values (via sensory perception for the rat), and which evolves over time under its own dynamics, under the influence of the agent's actions. The environment also delivers **rewards** (and punishments), and the standard objective of RL is to learn how to take actions that optimize the receipt of rewards over time.

The strategy for selecting actions can be usefully categorized into two different approaches, which can be thought of in terms of poles along a continuum (where this continuum is actually a high-dimensional space!):

* **Model-free**: learns a **policy** that is a _direct mapping_ from states to actions, i.e., _stimulus_ $\rightarrow$ _response_ (S-R) learning.

* **Model-based**: uses internal _models_ of the environment in addition to direct sensory experience of the environment, to select actions according to a _plan_ that is based on the model ([[#figure_rl-plan]]).

From a psychological perspective, this can be thought of as a distinction between _habit-based_ responding vs. a more thoughtful, considered basis of responding, which has been a central focus of theorizing for many years. For example, Thorndike's _law of effect_ ([[@Thorndike1898]]), which is the central idea behind _instrumental conditioning_, holds that any action that leads to a positive outcome will be increased in probability (and vice-versa for punishments). This is basically model-free RL, and it has long been associated with the function of the [[basal ganglia]] and the role of [[dopamine]] for shaping learning in this brain area (but not perhaps entirely accurately).

{id="figure_rl-plan" style="height:20em"}
![Model-based RL adds internal Goals, Models, and Plans to the basic RL paradigm, so that action selection can be based on internal representations of the possible outcomes of actions, not just learned stimulus-response associations. The CS (conditioned stimulus) can trigger activation of internal Goals to obtain the associated US (unconditioned stimulus) outcome, using an internal Model of the environment to develop a Plan of action to do so.](media/fig_rl_state_action_reward_plan_rat.png)

By contrast, _goal-directed_ or _controlled_ processing ([[@Tolman48]]; [[@BalleineDickinson98]]; [[@ShiffrinSchneider77]]; [[@CohenDunbarMcClelland90]]; [[@MillerCohen01]]) characterizes the more considered mode of processing, where possible consequences of an action are actually evaluated, instead of just responding on the basis of stimulus-response learning. This mode of processing typically depends on the [[prefrontal cortex]], which has the ability to maintain and manipulate internal representations (of possible actions and their consequences) in a way that is not generally possible in posterior [[neocortex]]. This distinction was also popularized by [[@^Kahneman11]] in terms of _system 1_ (fast) vs. _system 2_ (slow).

The direct mapping of the RL model-free vs. model-based concepts onto these biological and psychological terms, as initially proposed by [[@^Doya99]] and [[@^DawNivDayan05]], has a number of issues that we will encounter in the context of the [[Rubicon]] model ([[@OReillyNairRussinEtAl20]]). Nevertheless, conceptually all of these different terminologies are generally capturing similar intuitive ideas, which we can all relate to in terms of our daily experiences.

A nice thing about the RL framework is that it provides precise definitions and equations, which we will summarize in the following sections. Interestingly, even though model-based RL should in principle be more powerful than model-free, the additional complexities involved actually make it rather difficult to demonstrate advantages in practice, and many of the state-of-the-art models are much closer to the model-free end of the spectrum ([[@MoerlandBroekensJonker21]]; [[@PlaatKostersPreuss23]]).

## Model-free RL

From an abstract mathematical perspective, RL can be formulated as a _markov decision process_ (MDP), which just means that there are _discrete_ states $S$ that evolve over time according to a unitary (normalized) transition matrix $T$ that determines the probability of any given future state, based on the current state and any action taken by the agent. See [[@SuttonBarto98]] for the definitive textbook treatment.

The **temporal differences (TD)** algorithm provides a mechanism for learning to optimize expected reward in this context, and is one of the most central abstract learning rules used in RL ([[@Samuel59]]; [[@SuttonBarto81]]; [[@Sutton88]]; [[@SuttonBarto98]]). Furthermore, it provides a remarkably accurate explanation for many of the firing properties of [[dopamine]] neurons in the ventral tegmental area (VTA) of the brainstem ([[@MontagueDayanSejnowski96]]; [[@SchultzDayanMontague97]]).

Specifically TD provides a way of computing the [[#reward prediction error]] (RPE) signal in terms of the difference between predicted and actual reward outcomes. In the **actor critic** architecture ([[@BartoSuttonAnderson83]]), this RPE signal is generated by the _critic_, which provides the training signal for improving both motor actions taken by the _actor_ to obtain reward, and the accuracy of the critic's reward predictions. This is consistent with considerable neuroscience data showing that dopamine modulates learning throughout the brain, especially in the [[basal ganglia]] which has both actor and critic functionality.

### Reward prediction error

We start with the simplest, highly influential model of reward prediction error, the **Rescorla-Wagner** conditioning model [[@RescorlaWagner72]], which is mathematically identical to the _delta rule_ (see [[error-backpropagation]]).

The RPE $\delta$ is just the difference between the actual ($r$) and expected ($\hat{r}$ reward:

{id="eq_rw" title="Rescorla-Wagner RPE"}
$$
\delta = r - \hat{r}
$$

In a neural network, this expected reward $\hat{r}$ can be computed by synaptic weights from sensory input units:

{id="eq_rw-net" title="Rescorla-Wagner RPE"}
$$
\delta = r - \sum x w
$$

The weights can learn to accurately predict the actual reward values, and this delta value specifies the direction in which the weights should change:

{id="eq_rw-dw" title="Rescorla-Wagner learning"}
$$
\Delta w = \delta x
$$

This is identical to the delta learning rule, including the important dependence on the stimulus activity x --- you only want to change the weights for stimuli that are actually present (i.e., non-zero x's).

When the reward prediction is correct, then the actual reward value is _canceled out_ by the prediction, as shown in the second panel in [[dopamine#figure_da-schultz]] in [[dopamine]]. This rule also accurately predicts the other cases shown in the figure too (positive and negative reward prediction errors).

However, what the Rescorla-Wagner model fails to capture is the firing of [[dopamine]] to the onset of the CS in the second panel in [[dopamine#figure_da-schultz]]. 

### TD

In this context, the TD algorithm captures CS-onset firing by introducing time into the equation (as the name suggests) ([[@SuttonBarto81]]; [[@SuttonBarto98]]). Relative to Rescorla-Wagner, TD just adds one additional term to the delta equation, _f_, representing the _future_ reward values that might come later in time:

{id="eq_td-f" title="TD future reward"}
$$
\delta = (r + f) - \hat{r}
$$

Now the reward expectation $\hat{r}=\sum x w$ has to try to anticipate both the current reward _r_ and this future reward _f_. In a simple conditioning task, where the CS reliably predicts a subsequent reward, the onset of the CS results in an increase in this _f_ value, because once the CS arrives, there is a high probability of reward in the near future. Furthermore, this _f_ itself is not predictable, because the onset of the CS is not predicted by any earlier cue (and if it was, then that earlier cue would be the real CS, and drive the dopamine burst). Therefore, the r-hat expectation cannot cancel out the f value, and a dopamine burst ensues.

Although this _f_ value explains CS-onset dopamine firing, it raises the question of how can the system know what kind of rewards are coming in the future? Like anything having to do with the future, you fundamentally just have to guess, using the past as your guide as best as possible. TD does this by trying to _enforce consistency in reward estimates over time_. In effect, the estimate at time _t_ is used to train the estimate at time _t+1_, and so on, to keep everything as consistent as possible across time, and consistent with the actual rewards that are received over time.

This can all be derived in a very satisfying way by specifying something known as a **value function, V(t)** that is a sum of all present and future rewards, with the future rewards **discounted** by a _gamma_ factor, which captures the intuitive notion that rewards further in the future are worth less than those that will occur sooner. As the Wimpy character says in Popeye, "I'll gladly pay you Tuesday for a hamburger today." Here is that value function, which is an infinite sum going into the future:

{id="eq_v" title="Value function"}
$$
V(t) = \left. r(t) + \gamma^1 r(t+1) + \gamma^2 r(t+2) ... \right.
$$

We can get rid of the infinity by writing this equation *recursively*:

{id="eq_v-r" title="Recursive value function"}
$$
V(t) = \left. r(t) + \gamma V(t+1) \right.
$$

And because we don't know anything for certain, all of these value terms are really estimates, denoted by the little "hats" above them:

{id="eq_v-hat" title="Estimated value function"}
$$
\hat{V}(t) = r(t) + \gamma \hat{V}(t+1)
$$

So this equation tells us what our estimate at the current time _t_ should be, in terms of the future estimate at time _t+1_. Next, we subtract V-hat from both sides, which gives us an expression that is another way of expressing the above equality --- that the difference between these terms should be equal to zero:

$$
0 = \left( r(t) + \gamma \hat{V}(t+1) \right) - \hat{V}(t)
$$

This is mathematically stating the point that TD tries to keep the estimates consistent over time --- their difference should be zero. But as we are learning our V-hat estimates, this difference will _not_ be zero, and in fact, the extent to which it is not zero is the extent to which there is a reward prediction error:

{id="eq_td-delta" title="TD delta RPE"}
$$
\delta = \left( r(t) + \gamma \hat{V}(t+1) \right) - \hat{V}(t)
$$

If you compare this to the equation with _f_ in it above, you can see that:

$$
f = \gamma \hat{V}(t+1)
$$

and otherwise everything else is the same, except we've clarified the time dependence of all the variables, and our reward expectation is now a "value expectation" instead (replacing the r-hat with a V-hat). Also, as with Rescorla-Wagner, the delta value here drives learning of the value expectations.

The TD learning rule can be used to explain a number of different conditioning phenomena, and its fit with the firing of dopamine neurons in the brain has led to a huge amount of research progress. It represents a real triumph of the computational modeling approach for understanding (and predicting) brain function.

From a neuroscience perspective there are actually a number of different brain areas and circuits that work together to drive the firing of VTA dopamine neurons, as summarized in the [[PVLV]] model ([[@OReillyFrankHazyEtAl07]]; [[@HazyFrankOReilly10]]; [[@MollickHazyKruegerEtAl20]]). These mechanisms produce some important differences in dopamine firing behavior relative to the predictions of the TD model, including critically the absence of a chain-like progression of firing at the time of the US to the time of the CS, among others ([[@MollickHazyKruegerEtAl20]]). Nevertheless, the TD model provides an elegant and computationally powerful learning framework that explains many critical aspects of the role of dopamine in learning across different brain areas.

### TD versus temporal derivative learning

Although they share some core conceptual similarities, and a very similar name, it is important to distinguish the TD framework from the [[temporal derivative]] based learning mechanisms used in [[Axon]] for [[error-driven learning]] in the [[neocortex]], via the [[kinase algorithm]]. Most importantly, from a biological perspective, TD translates a temporal difference signal into the firing of dopamine neurons, which then _explicitly_ represents the RPE signal. By contrast, the prediction error signal in the temporal derivative framework _remains implicit_ as a difference in neural activity over time, which propagates throughout the neocortex via bidirectional connectivity. Each synapse then adjusts its synaptic weights locally in a way that is sensitive to these temporal derivatives, as contrasted with the direct neuromodulatory role played by dopamine.

## Model-based RL

Many different _model based_ mechanisms can be used in addition or instead of model-free RL ([[@MoerlandBroekensJonker21]]; [[@PlaatKostersPreuss23]]). Many of these mechanisms have direct analogs in the cognitive and neuroscience domain, and are useful for providing a computational insights into these areas. As will become apparent, each of these approaches has clear intuitive advantages, but also significant computational challenges that often prevent these advantages from being realized in more realistic, complex task domains. Overcoming these challenges is a major goal of the [[Rubicon]] model, as outlined below.

* **State abstraction**: instead of directly using the raw state inputs, higher-level abstractions often  provide a better basis space for deciding which actions to take. The importance of this computational principle is discussed at length in [[categorization]], and many critical aspects of human intelligence depend on powerful abstractions performed in the posterior [[neocortex]]. In effect, our brains provide us with a rich and complex internal model of the world, at many different levels of abstraction.

    The deep layered architectures of [[abstract neural network]]s enable these kinds of abstractions to develop, and the marriage of such architectures with reinforcement learning has led to impressive performance on challenging tasks such as playing a wide variety of games, from Atari video games to classics such as chess and go (e.g., [[@SilverHuangMaddisonEtAl16]]). A paradigmatic example is the _Deep Q-learning Network_ (DQN) ([[@MnihBadiaMirzaEtAl16]]), which couples a deep convolutional network (DCN) with TD-style Q-learning as described above, using "end-to-end" [[error backpropagation]] to train representations across many network layers, which are thus specifically optimized for action selection and reward prediction.

* **World models**: To the extent that our internal mental models of the world include _temporal dynamics_ (i.e., can approximate the MDP transition matrix in formal terms), then these models can potentially be used for _mental simulation_ or _imagination_ of different possible courses of action. An early paradigmatic example of this is the _Dyna_ architecture of [[@^Sutton91a]], which uses a learned transition matrix function to generate synthetic training data to accelerate the learning process.

    A world model is also known as a _forward model_ of the environment [[@JordanRumelhart92]], and it is essentially what [[predictive learning]] does: learning to predict what will happen next. As we discussed extensively in that page, predictive learning is essentially "free" when you have sequential structure in an environment, and this is what enables [[large language models]] to train on so much textual data. Thus, we assume that a major component of the posterior neocortical learning is learning these kinds of dynamical, predictive, forward models of the world.

    Recent examples of such models are the various _Dreamer_ networks ([[@HafnerLillicrapNorouziEtAl22]]), which do multi-step _roll-outs_ of predicted sequences of events using learned world models in the context of various video-game based environments. These roll-outs can then anticipate the future consequences of possible action choices, potentially enabling better decision-making.

    However, a major issue with such models is that any inaccuracies will _compound_ exponentially over iterated time steps, and typically the environment is sufficiently stochastic that the _branching factor_ for different possible future "timelines" also grows exponentially (recent Marvel-universe movies capture this problem nicely).
    
* **Multiscale models**: The longstanding principle of hierarchical, multi-scale representations (e.g., [[@BartoMahadevan03]]) can potentially help with the above problems in world models, by performing predictions at more abstract, coarse-grained levels where the dimensionality is sufficiently low, and the stochasticity is encapsulated in the abstractions instead of being directly manifest in the dynamics. This is a key element of the [[Rubicon]] model.

    Hierarchical approaches to RL encompass the concept of _subgoals_: breaking a complex problem down into more manageable sub-tasks, as in the widely-discussed _options_ framework in RL ([[@SuttonPrecupSingh99]]), and the _chunking_ mechanism in the ACT-R framework ([[@AndersonLebiere98]]). The difficulty with these and related hierarchical RL approaches is that it is often hard to know when and how to organize relevant subtasks, and the possibilities for the number of ways to do so is again subject to exponential combinatorial explosion problems ([[@PateriaSubagdjaTanEtAl21]]). Furthermore, almost everywhere any kind of _rigid_ hierarchical system has been proposed, one quickly realizes that the real world is not actually so amenable to such systems, and requires more graded, flexible, _soft_ frameworks.

* **Planning mechanisms**: There is a large literature on planning algorithms, which converges in many cases with approaches used in the RL world ([[@RussellNorvig16]]; [[@MoerlandBroekensJonker21]]). Intuitive ideas include searching forward from the current state to try to find a desired outcome, versus searching backward from the desired outcome state to find a path that connects to the current state. In every case, the planning process requires something like a world model that allows the different possible actions to be simulated in order to search through the space.

    The human-beating AI approaches to the games of chess and go relied on these classical-style state-space planning algorithms, constructed using domain-specific knowledge about the rules and relevant structure of these games ([[@SilverHuangMaddisonEtAl16]]; [[@SilverSchrittwieserSimonyanEtAl17]]).

    Despite the success of these specific examples, they fail to generalize to more complex real-world tasks due to the now-familiear exponential combinatorial explosion problem. Even in well-defined search spaces associated with games like chess, searching through the state space quickly becomes impossible, and significant domain-specific heuristics are required to aggressively prune the space and focus on only the most promising potential avenues.

## The Rubicon approach

In the context of the above model-based RL solutions, and the consistent challenges in terms of exponential combinatorial explosion problems that they face, it is useful to situate the [[Rubicon]] approach and why it might have the potential to succeed where other approaches have not.

First, the [[Rubicon]] model builds on a posterior neocortex that learns multi-level abstract world models via [[predictive learning]], providing a foundation for higher-level planning and control mechanisms. Unlike most existing AI / RL models, however, the [[Axon]] framework incorporates pervasive [[bidirectional connectivity]], which critically allows higher-level goal and plan level information to pervasively shape and directly constrain the processing of these posterior cortical networks. In addition to shaping learning in a manner similar to end-to-end error-backpropagation, this _top-down processing_ can potentially sharply prune and focus processing on only the most relevant aspects of the current sensory inputs.

Thus, we hypothesize that these top-down constraints can significantly mitigate the combinatorial explosion problems inherent in deciding what information is most relevant to include in the current world model encoding, and which aspects of the temporal evolution of this model are most relevant for the task at hand.

Furthermore, this bidirectional connectivity supports a robust form of _multiple constraint satisfaction_ processing where many different bottom-up sensory cues can interact dynamically with top-down goal and task plan representations to converge on the specific affordances for goal-relevant actions present in the current environment ([[@HopfieldTank85]]). The ability of this form of processing to quickly and efficiently converge on reasonable solutions to complex optimization problems that would otherwise require exponential time to search in more discrete ways (e.g., the monte-carlo tree search used in the go-playing models; [[@SilverSchrittwieserSimonyanEtAl17]]) is a key potential advantage.

Finally, the defining feature of the [[Rubicon]] model is that the system makes an explicit commitment to a given goal and plan of action at any given point in time, which magnifies all of the above effects by virtue of the strong shared focus throughout the neocortical network on this one goal and plan. Furthermore, making such a commitment requires a goal selection process with neural mechanisms that support the essential "mental time travel" process that activates relevant potential outcomes and plans at an earlier "proactive" step in the process, which effectively combines backward and forward planning into one big constraint satisfaction process at the point of goal selection.

