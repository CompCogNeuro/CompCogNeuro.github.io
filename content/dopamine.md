+++
Categories = ["Rubicon"]
bibfile = "ccnlab.json"
+++

**Dopamine** is a [[neuromodulator]] that modulates learning throughout the brain. It is secreted primarily by two adjacent brain areas in the brainstem, the **ventral tegmental area (VTA)** and **substantia nigra pars compacta (SNc)**. From a computational perspective, activity of these dopamine neurons closely matches the behavior of the TD (_temporal differences_) [[reinforcement learning]] algorithm, which computes the _reward prediction error_ (RPE; the difference between predicted and actual reward received).

From a biological perspective many different neural pathways converge on the VTA and SNc to drive the firing of dopamine neurons. The [[PVLV]] model provides a well-validated framework for understanding what each of these different pathways contribute, and is a core element of the broader [[Rubicon]] framework for goal-driven, motivated behavior.

{id="figure_da-schultz" style="height:30em"}
![Characteristic patterns of neural firing of the dopaminergic neurons in the ventral tegmental area (VTA) and substantia nigra pars compacta (SNc), in a simple conditioning task (Schultz et al, 1997). Prior to conditioning, when a reward is delivered, the dopamine neurons fire a burst of activity (top panel --- histogram on top shows sum of neural spikes across the repeated recording traces shown below, with each row being a different recording trial). After the animal has learned to associate a conditioned stimulus (CS) (e.g., a tone) with the reward, the dopamine neurons now fire to the onset of the CS, and not to the reward itself. If a reward is withheld after the CS, there is a dip or pause in dopamine firing, indicating that there was some kind of prediction of the reward, and when it failed to arrive, there was a negative prediction error. This overall pattern of firing across conditions is highly consistent with reinforcement learning models based on reward prediction error. Reproduced from Schultz et al, 1997](media/fig_schultz97_vta_td.png)

