+++
Categories = ["Rubicon", "Neuroscience", "Cognition"]
bibfile = "ccnlab.json"
+++

[This page covers needs, drives, motivation, and affect in relation to emotion]

**Emotion** is something everyone experiences but it remains somewhat difficult to define precisely. From the computational and neuroscience perspective taken here, we can provide a more precise definition:

> Emotion reflects the activity of midbrain and brainstem systems that have been shaped by evolution to guide the behavior of the organism in adaptive ways to satisfy its **needs** in relation to its external and internal state. These brain systems are anchored by the hypothalamus and associated nuclei that directly sense the internal body state (interoception), many of which are conveyed by the vagus nerve via the medulla oblongota. The [[amygdala]] makes connections between sensory inputs processed by the cortex (visual, auditory) and underlying body states. Important changes in body state are signalled by neuromodulators such as [[dopamine]] and [[serotonin]] that are driven by these systems, and have widespread effects on the entire brain. There are many layers of control, integration, and modulation of these systems, including the ventral and medial [[basal ganglia]] via the ventral pallidum, the [[lateral habenula]], and ventral and medial [[prefrontal cortex]] areas.

In other words, emotion by this definition is the domain of the [[limbic system]], and computationally is related to [[reinforcement learning]] (RL). However, RL typically only deals with a very limited scope of "emotion", as captured in a single scalar reward value. The [[Rubicon]] framework encompasses a broader range of states and the relationship between needs / drives and the current internal state, including goals and **motivational** state.

The term **affect** typically encompasses a broader scope than emotion, but we use them essentially interchangably here, under the above definition.

{id="figure_valence-arousal" style="height:25em"}
![Valence (positive vs. negative) vs. arousal (high vs. low activation) in the 2D _circumplex_ model of emotion.](media/fig_emotion_valence_arousal.png)

{id="figure_plutchik" style="height:40em"}
![Plutchik's (2001) wheel of emotions, with arousal (intensity) represented as distance from the center along any of 8 different categories of opponent emotions.](media/fig_emotion_plutchik.png)

From a subjective, psychological perspective, emotions can be categorized most broadly in terms of **valence** (positive vs. negative) and **arousal** (high vs. low) ([[#figure_valence-arousal]]). Progressively more differentiated systems have been developed, with Plutchik's wheel of emotions providing a nice comprehensive set ([[#figure_plutchik]]; [[@Plutchik01]]).

{id="figure_maslow" style="height:35em"}
![Maslow's hierarchy of needs, progressing from most essential to those that gain salience once the lower-level needs are satisfied.](media/fig_maslow_hierarchy_needs.png)

What about more basic feelings, like _hunger_, _thirst_, etc? [[@^Maslow43]] developed a hierarchical schema to organize human needs, with the most basic ones providing the foundation, and higher-level ones only entertained once the basic needs are satisfied ([[#figure_maslow]]).

{id="figure_needs" style="height:40em"}
![Emotional states as providing guidance toward satisfying needs, with a lack of need satisfaction generally driving negative emotional states, while need satisfaction drives positive ones. The list of needs is ordered  with the most essential needs at the bottom, as in Maslow's hierarchy. We don't usually think of things like hunger and thirst as emotions, but according to this system, they play the same overall role, and share many neural substrates in common. Resources refers to any kind of material thing needed to survive in the current physical environment (money, territory, nesting material, tools, building materials, etc). The social factors include S = self perspective and O = other perspective. Most of these needs apply across all species of mammals and other vertibrates to variable extents. The social needs are particularly important for shaping flexible, open-ended human cognition.](media/fig_emotion_needs.png)

We can build on this schema to define a comprehensive organization of emotions organized around the things that humans need to do from an evolutionary perspective, as summarized in [[#figure_needs]]. This table provides a place for most of the terms in Plutchik's wheel of emotions, and covers most of the needs in Maslow's hierarchy, providing a reasonable basis for seeing how emotions help guide us toward the things we need to survive and thrive.

Interestingly, most of these needs apply across all species of mammals and even more broadly to all animals. This suggests that emotion per se is widely shared across the animal kingdom, not something that is uniquely human, which is consistent with emotion being rooted in the deepest, most evolutionarily ancient portions of the brainstem. What the human brain adds is a much enlarged [[neocortex]] and an associated ability to gain [[conscious awareness]] of all this emotional stuff going on, largely via the extensive connectivity from these brainstem areas into the medial and ventral regions of the [[prefrontal cortex]].

{id="figure_drives" style="height:5em"}
![Drive reduction theory according to Hull, 1943. Basic needs create drives when those needs are not satisfied, and behavior is then recruited to satisfy those drives.](media/fig_drive_reduction_hull.png)

From a computational perspective, we can define each of these needs as having a current **drive** level ([[@Hull43]]), which reflects the current lack of satisfaction of the need ([[#figure_drives]]). This is the model incorporated in the [[Rubicon]] model, with a variable number of need factors that each have associated drives and USs that satisfy them.

## Social needs

Like many other species, humans are strongly social animals, and we depend on others to survive and thrive. There is a strong correlation between brain size and size of social networks in primates ([[@Dunbar92]]; [[@Dunbar16]]), suggesting that we owe our big brains to being highly social animals: it takes careful thought and planning to navigate the complexities of the social world. Furthermore, [[@^Tomasello01]] argues that humans have a unique drive to share that is not evident in even our closest primate relatives (e.g., chimpanzees). This sharing instinct is what drives the pervasive nature of our cultural evolution, where we acquire much of our knowledge from the accumulated wisdom of those who have come before us.

<!--- ## Universality of emotions -->
<!--- TODO: faces -->
