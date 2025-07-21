+++
Categories = ["Rubicon", "Neuroscience"]
Title = "STN"
Name = "STN"
bibfile = "ccnlab.json"
+++

The **subthalamic nucleus** (STN) is a part of the [[basal ganglia]] circuit, which receives direct input from the [[neocortex]], and is bidirectionally connected with the globus pallidus externus (GPe) and projects excitation to the output nuclei (SNr, GPi). In the PCore model, the STN is responsible for driving an initial brake on BG disinhibition by exciting these output nuclei, and then it receives a rebound inhibitory signal from the GPePr (prototypical GPe) neurons, which establishes a long-lasting pause in activity that opens up a brief time window for the BG to perform its decision-making function. This provides a mechanism behind the phasic firing of BG neurons.

This page provides a more detailed discussion of neural data on the STN to better understand its functionality, given that it plays such a critical and somewhat controversial role in the circuit.

As reviewed in [[basal ganglia]], much of the available data on BG firing is consistent with a brief, phasic window of activity for individual neurons, with otherwise remarkably silent striatal neurons outside of these brief windows. This is overall compatible with the contributions of the STN as hypothesized.


{id="figure_fujimoto-kita93" style="height:20em"}
![Neural recordings from STN neurons, showing two initial bursts of activity followed by a long period of inhibition, which then recovers. The top panel shows a finer time-scale zoom-in of the two initial bursts, while the bottom panel shows a longer time-scale window of inhibition, of roughly 150 ms. From Fujimoto & Kita, 1993](media/fig_stn_pause_fujimoto_kita_93_fig3.png)

{id="figure_magill04" style="height:20em"}
![Neural recordings from STN neurons, with different response profiles, with panels A and B showing two initial bursts of activity followed by a long period of inhibition, which then recovers. The recovery window is not specifically shown but is described in the paper. From Magill et al., 2004](media/fig_stn_pause_magill_etal_04.png)

Furthermore, direct recordings of neural activity in the STN clearly demonstrate the pausing behavior, as shown in [[#figure_fujimoto-kita93]] and [[#figure_magill04]] from [[@^FujimotoKita93]] and [[@^MagillSharottBevanEtAl04]] respectively. There is a roughly 150 ms window during which the STN neurons experience sustained inhibition after the initial bursts.

The cause of this inhibition was unknown to the authors of the above papers, but there is some evidence from further electrophysiological studies and computational models that it could be based on inactivation of Na+ channels due to a high sustained plateau potential ([[@KassMintz06]]; [[@BeurrierBioulacAudinEtAl01]])

## Hold-your-horses vs pause models

TODO, discuss here..
