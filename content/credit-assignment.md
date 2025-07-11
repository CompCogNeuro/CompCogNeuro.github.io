+++
Categories = ["Learning", "Computation"]
bibfile = "ccnlab.json"
+++

**Credit assignment** is the most important process in any kind of [[learning]] mechanism, which assigns differential credit / blame to [[neuron]]s for the current learning event. In many cases, it is accomplished by multiplying by the [[activation]] value of the sending neuron, as in both [[Hebbian learning]] and [[error backpropagation]] learning. See [[error backpropagation#credit assignment]] in error backpropagation for mathematical details.

## Temporal credit assignment

In a standard error backpropagation context, the credit assignment process takes place effectively in parallel using the current activation states of neurons. However, in a real-world continuous-time context, the true sources of credit or blame for a behaviorally or motivationally relevant outcome are typically in the past, e.g., a prior action that was taken by the agent. This creates the _temporal_ version of the credit assignment problem, which is much more challenging than the immediate credit assignment process available in error backpropagation.

[[Reinforcement learning]] (RL) algorithms must solve this temporal credit assignment problem in one way or another. For example, the TD (temporal differences) algorithm performs temporal credit assignment by cascading a reward prediction error (RPE) backward through time.

By contrast, the biologically-based [[PVLV]] algorithm for RL performs temporal credit assignment via a coordinated set of biological mechanisms involving the [[amygdala]], ventral [[basal ganglia]] (BG), and the [[orbitofrontal cortex]] (OFC). The amygdala learns to recognize stimuli associated with outcomes, and the OFC can actively maintain an expectation of these future outcomes over time, bridging temporal gaps to provide temporal credit assignment. The BG learns to trigger updating of the OFC activity state, using inputs from the amygdala and other areas. This is a major component of the broader [[Rubicon]] framework.

In [[large language models]] using the [[transformer]] architecture, temporal credit assignment is accomplished via standard parallel error backpropagation by providing all of the relevant temporal context as direct inputs to the network. This is not ecologically, neurally, or cognitively plausible: people's memory capacity for text is on the order of 7 or so words. Thus, different mechanisms are required.

