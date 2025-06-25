+++
Categories = ["Neuroscience", "Rubicon"]
bibfile = "ccnlab.json"
+++

A **Neuromodulator** is a type of neurotransmitter that has _modulatory_ effects on the learning and activity of neurons. Neuromodulators typically have very broad effects across the brain, and thus are critical for conveying the most important, survival-relevant signals, and also for modulating overall levels of arousal across the wake-sleep cycle. The [[Rubicon]] framework provides biologically-based implementations of some of these neuromodulators, which play a central role in goal-driven motivated behavior.

There are a handful of such neuromodulators which all have overlapping yet distinguishable effects across the brain, including:

* [[Dopamine]] (DA), which is particularly important for modulating _learning_ by encoding a [[TD#reward prediction error]] (RPE; the difference between predicted and actual rewards). Although popular culture has appropriated dopamine to mean "pleasure", it is really much more of a learning signal, because the RPE is a _difference_ measure, not the raw reward signal itself.

* [[Acetylcholine]] (ACh), which is also important for learning but also significantly modulates the excitability of neurons in the [[basal ganglia]] and other areas. In the [[Rubicon]] model, ACh is a clear _salience_ signal, reflecting both reward and reward prediction events, and the onset of novel stimuli. Unlike dopamine, ACh is not an error (RPE-like) signal, but rather represents the "raw" reward or reward prediction. This allows it to modulate the excitability of neurons involved in engaging new goal states, as contrasted with the learning-specific role of dopamine, where the RPE difference property is critical.


