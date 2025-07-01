+++
Categories = ["Activation", "Neuroscience"]
bibfile = "ccnlab.json"
+++

This page provides a detailed treatment of the electrophysiology of the [[neuron]], explaining how differential concentrations of individual ions lead to the electrical dynamics of the neuron.

First, some basic facts of electricity.  Electrons and protons, which together make up atoms (along with neutrons), have electrical charge (the electron is negative, and the proton is positive).  An **ion** is an atom where these positive and negative charges are out of balance, so that it carries a **net charge**.  Because the brain carries its own salt-water ocean around with it, the primary ions of interest are:

* **sodium ($Na^+$)** which has a net positive charge.
* **chloride ($Cl^-$)** which has a net negative charge.
* **potassium ($K^+$)** which has a net positive charge.
* **calcium ($Ca^{++}$)** which has _two_ net positive charges.

{id="figure_electricity"}
![Basic principles of electricity: when there is an imbalance of positive and negative charged ions, these ions will flow so as to cancel out this imbalance.  The flow of ions is called a current (I), driven by the potential (level of imbalance) V with the conductance G (e.g., size of the opening between the two chambers) determining how quickly the ions can flow.](media/fig_electricity.png)

As we noted in the main chapter, these ions tend to flow under the influence of an electrical potential (voltage), driven by the basic principle that **opposite charges attract and like charges repel**.  If there is an area with more positive charge than negative charge (i.e., an **electrical potential**), then any negative charges nearby will be drawn into this area (creating an electrical **current**), thus nullifying that imbalance, and restoring everything to a neutral potential. [[#figure_electricity]] shows a simple diagram of this dynamic.  The **conductance** is effectively how wide the opening or path is between the imbalanced charges, which determines how quickly the current can flow.

**Ohm's law** formalizes the situation mathematically:

{id="eq_ohms" title="Ohm's law"}
$$
I = G V
$$

(i.e., current = conductance times potential).

{id="figure_diffusion"}
![Diffusion is the other major force at work in neurons --- it causes each ion individually to balance out its concentration uniformly across space (i.e., on both sides of the chamber).  Concentration imbalances can then cause ions to flow, creating a current, just like electrical potential forces.](media/fig_diffusion.png)

The other major force at work in the neuron is **diffusion**, which causes individual ions to move around until they are uniformly distributed across space ([[#figure_diffusion]]).  Interestingly, the diffusion force originates from random movements of the ions driven by heat --- ions are constantly bumping around through space, with a mean velocity proportional to the temperature of the environment they're in. This constant motion creates the diffusion force as a result of the inevitable increase in **entropy** of a system --- the maximum entropy state is where each ion is uniformly distributed, and this is in effect what the diffusion force represents. The key difference between the diffusion and electrical force is:

* Diffusion operates individually on each ion, regardless of its charge compared to other ions etc --- each ion is driven by the diffusion force to spread itself uniformly around. In contrast, electrical forces ignore the identity of the ion, and only care about the net electrical charge.  From electricity's perspective, $Na^+$ and $K^+$ are effectively equivalent.

It is this critical difference between diffusion and electrical forces that causes different ions to have different driving potentials, and thus exert different influences over the neuron.

{id="figure_ions"}
![Major ions and their relative concentrations inside and outside the neuron (indicated by the size of the circles).  These relative concentration differences give rise to the different driving potentials for different ions, and thus determine their net effect on the neuron (whether they pull it "up" for excitation or "down" for inhibition).](media/fig_ions.png)

[[#figure_ions]] shows the situation inside and outside the neuron for the major ion types.  The concentration imbalances all stem from a steady **sodium pump** that pumps $Na^+$ ions out of the cell.  This creates an imbalance in electrical charge, such that the inside of the neuron is more negative (missing all those $Na^+$ ions) and the outside is more positive (has an excess of these $Na^+$ ions).

This negative net charge (i.e., **negative resting potential**) of about -70mV pushes the negative $Cl^-$ ions outside the cell as well (equivalently, they are drawn to the positive charge outside the cell), creating a concentration imbalance in chloride as well.  Similarly, the $K^+$ ions are drawn ''into'' the cell by the extra negative charge within, creating an opposite concentration imbalance for the potassium ions.

All of these concentration imbalances create a strong diffusion force, where these ions are trying to distribute themselves more uniformly.  But this diffusion force is counteracted by the electrical force, and when the neuron is at rest, it achieves an **equilibrium** state where the electrical and diffusion forces exactly balance and cancel each other out.   Another name for the diving potential for an ion (i.e., which direction it pulls the cell's membrane potential) is the **equilibrium potential** --- the electrical potential at which the diffusion and electrical forces exactly balance.

As shown in [[#figure_ions]], the $Cl^-$ and $K^+$ ions have driving potentials that are essentially equivalent to the resting potential, -70mV.  This means that when the cell's membrane potential is at this -70mV, there is no net current across the membrane for these ions --- everything will basically stay put.

Mathematically, we can capture this phenomenon using the same equation we derived from the tug-of-war analogy:

{id="eq_ohm_I" title="Current"}
$$
I = G (E-V)
$$

Notice that this is just a simple modification of Ohm's law --- the E value (the driving potential) "corrects" Ohm's law to take into account any concentration imbalances and the diffusion forces that they engender.  If there are no concentration imbalances, then E = 0, and you get Ohm's law (modulo a minus sign that we'll deal with later).

If we plug an E value of -70mV into this equation, then we see that the current is 0 when V = -70mV.  This is the definition of an equilibrium state.  No net current.

Now consider the $Na^+$ ion.  Both the negative potential inside the neuron, and the concentration imbalance, drive this ion to want to move into the cell.  Thus, at the resting potential of -70mV, the current for this ion will be quite high if it is allowed to flow into the cell.  Indeed, it will not stop coming into the cell until the membrane potential gets all the way up to $+55mV$ or so.  This equilibrium or driving potential for $Na^+$ is positive, because it would take a significant positive potential to force the $Na^+$ ions back out against their concentration difference.

The bottom line of all this is that synaptic channels that allow $Na^+$ ions to flow will cause $Na^+$ to flow *into* the neuron, and thereby excite the receiving neuron.  In effect, the sodium pump "winds up" the neuron by creating these concentration imbalances, and thus the potential for excitation to come into the cell against a default background of the negative resting potential.

Finally, when excitatory inputs do cause the membrane potential to increase, this has the effect of drawing more $Cl^-$ ions back into the cell, creating an inhibitory pull back to the -70mV resting value, and similarly it pushes $K^+$ ions out of the cell, which also makes the inside of the cell more negative, and has a net inhibitory effect.  The $Cl^-$ ions only flow when inhibitory GABA channels are open, and the $K^+$ ions flow all the time through the always-open leak channels.


