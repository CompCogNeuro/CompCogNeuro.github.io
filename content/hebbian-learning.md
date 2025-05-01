+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.json"
+++

The famous Canadian psychologist Donald O. Hebb predicted the nature of the NMDA channel many years in advance of its discovery, just by thinking about how learning should work at a functional level. Here is a key quote:

> _Let us assume that the persistence or repetition of a reverberatory activity (or "trace") tends to induce lasting cellular changes that add to its stability.… When an axon of cell A is near enough to excite a cell B and repeatedly or persistently takes part in firing it, some growth process or metabolic change takes place in one or both cells such that A's efficiency, as one of the cells firing B, is increased._

This can be more concisely summarized as _cells that fire together, wire together._ The NMDA channel is essential for this process, because it requires both pre and postsynaptic activity to allow $Ca^{++}$ to enter and drive learning. It can detect the *coincidence* of neural firing. Interestingly, Hebb is reputed to have said something to the effect of "big deal, I knew it had to be that way already" when someone told him that his learning principle had been discovered in the form of the NMDA receptor.

Mathematically, we can summarize Hebbian learning as:

$$
\Delta w = x y
$$

where $\Delta w$ is the change in synaptic weight *w*, as a function of sending activity *x* and receiving activity *y*.

Anytime you see this kind of pre-post product in a learning rule, it tends to be described as a form of Hebbian learning. For a more detailed treatment of Hebbian learning and various popular variants of it, see the *Hebbian Learning* Appendix.

As we'll elaborate below, this most basic form of Hebbian learning is very limited, because weights will only go up (given that neural activities are rates of spiking and thus only positive quantities), and will do so without bound. Interestingly, Hebb himself only seemed to have contemplated LTP, not LTD, so perhaps this is fitting. But it won't do anything useful in a computational model. Before we get to the computational side of things, we cover one more important result in the biology.

### Spike Timing Dependent Plasticity

{id="figure_stdp"  style="height:20em"}
![Spike timing dependent plasticity demonstrated in terms of temporal offset of firing of pre and postsynaptic neurons. If post fires after pre ($\Delta t > 0$, right side), then the weights go up, and otherwise ($\Delta t < 0$, left side) they go down. This fits with a causal flow of information from pre to post. However, more complex sequences of spikes wash out such precise timing and result in more generic forms of Hebbian-like learning. Reproduced from Bi and Poo (1998).](media/fig_stdp_bipoo98.png)

[[#figure_stdp]] shows the results from an experiment [[@BiPoo98]] that captured the imagination of many a scientist, and has resulted in extensive computational modeling work. This experiment showed that the precise order of firing between a pre and postsynaptic neuron determined the sign of synaptic plasticity, with LTP resulting when the presynaptic neuron fired before the postsynaptic one, while LTD resulted otherwise. This **spike timing dependent plasticity (STDP)** was so exciting because it fits with the *causal* role of the presynaptic neuron in driving the postsynaptic one. If a given pre neuron actually played a role in driving the post neuron to fire, then it will necessarily have to have fired in advance of it, and according to the STDP results, its weights will increase in strength. Meanwhile, pre neurons that have no causal role in firing the postsynaptic cell will have their weights decreased. However, this STDP pattern does not generalize well to realistic spike trains, where neurons are constantly firing and interacting with each other over 100's of milliseconds [[@ShouvalWangWittenberg10]]. Nevertheless, the STDP data does provide a useful stringent test for computational models of synaptic plasticity. We base our learning equations on a detailed model using more basic, biologically-grounded synaptic plasticity mechanisms that does capture these STDP findings [[@UrakuboHondaFroemkeEtAl08]], but which nevertheless result in quite simple learning equations when considered at the level of firing rate.

