+++
Categories = ["Rubicon", "Simulations"]
bibfile = "ccnlab.json"
+++

{id="sim_inhib" collapsed="true"}
```Goal
// see https://github.com/emer/axon/tree/main/sims/armmaze for source code
armmaze.EmbedSim(b)
```

<div>

This model explores how the [[Rubicon]] model exhibits goal-driven decision-making in the case of a simulated rat in a multi-arm maze, where each arm has different cost vs. benefit tradeoffs that must be learned and then used to make better decisions.

[[@FriedmanHommaGibbEtAl15]] shows that PL projections to dorsomedial striatum (ALM / dlPFC motor planning area) activate inhibitory interneurons in striosomes, which then inhibit striosomes, during high-conflict cases where cost-benefit ratio is high. Inhibiting this PL projection led to an increase in high-reward choices specifically in the high-cost, high-reward case. Activating PL had opposite effects. Activating ACC lead to increase in higher-reward preference across the board.

