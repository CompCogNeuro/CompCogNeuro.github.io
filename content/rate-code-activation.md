+++
Categories = ["Activation", "Axon"]
bibfile = "ccnlab.json"
+++

Most [[abstract neural network]] (ANN) models use a floating-point **activation** value to represent something like the overall firing rate of a biological [[neuron]]. This was the case in the [[Leabra]] model, prior to the development of the discrete spiking [[Axon]] model, for example.

The _rate code approximation_ can be an important way of simplifying the complexity of a model, and it is consistent with the reliable finding that the overall rate of firing of biological neurons does correlate well with the behavior and inferred internal representations in a large number of electrophysiological recordings. However, there is no doubt that it is an _approximation_ to how the actual system behaves, and there are a number of potential costs to using such an approximation across various levels of analysis [[@Brette15]].

As emphasized in the [[neuron]] chapter, discrete spiking has advantages in representing graded, probabilistic information, while also supporting a fast initial propagation of new stimulus information via the first spike responses of neurons. By contrast, rate code neurons, at least with the additional biologically-motivated properties of the [[Leabra]] model, could either respond quickly but in a not very graded manner, or slowly and with graded responding.

{id-"figure_rate-code-approx}
![Quality of the rate code approximation (rate line) to actual spiking rate (Spike line), over a range of excitatory input levels (GBarE). The rate code approximation is based on the "gelin" (linear in Ge) model comparing $Ge$ to $g_e^{\Theta}$, using the Noisy XX1 sigmoidal function, and also including spike rate adaptation as included in the AdEx model.](media/fig_neuron_rate_code_approx.png) }

[[#Figure_rate-code-approx]] shows the match between the [[Leabra]] rate code approximation and the actual rate of spiking in the [[AdEx]] discrete spiking equations used in [[Axon]]. Conceptually, we can think of this spiking rate as reflecting the net output of a small population of roughly 100 neurons that all respond to similar information. The neocortex is organized anatomically with **microcolumns** of roughly this number of neurons, where all of the neurons do indeed code for similar information. Use of this rate code activation enables smaller-scale models that converge on a stable interpretation of the input patterns rapidly, with an overall savings in computational time and model complexity.


