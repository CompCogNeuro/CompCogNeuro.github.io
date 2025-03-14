+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.bib"
+++

The **kinase** learning algorithm is an abstraction of the detailed chemical pathways involved in [[synaptic-plasticity]], which are mediated by _kinases_ such as [[camkii|CaMKII]] and [[dapk1|DAPK1]] that play central roles in these learning processes. This algorithm accomplishes [[error-driven-learning]] via a [[temporal-derivative]] computed between the faster CaMKII-mediated pathway and the slower DAPK1-mediated pathway, which are known to be in competition with each other [@GoodellZaegelCoultrapEtAl17].

See the [[temporal-derivative]] page for a critical illustration of how a competitive interaction between fast and slow pathways can compute the _error gradient_ at the heart of error-driven learning. This approach to computing the error gradient was pioneered in the [[boltzmann-machine]] by [@AckleyHintonSejnowski85; @HintonSejnowski86], and the _GeneRec_ (generalized recirculation) algorithm [@OReilly96] shows how the [[error-backpropagagion]] algorithm can be computed via a temporal derivative.

There is initial direct evidence showing that the direction of synaptic plasticity in neurons recorded in the mouse hippocampal area is consistent with this temporal derivative hypothesis [@JiangEtAlIP].


