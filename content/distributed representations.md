+++
Categories = ["Activation", "Computation", "Cognition", "Neuroscience"]
bibfile = "ccnlab.json"
+++

**Distributed representations** are simply the population of many individual [[neuron detector]]s, each of which is detecting something different. The aggregate pattern of output activity ("detection alarms") across this population of detectors can encode a wealth of information in _parallel_, without having to decide on a specific individual feature that is most relevant at the present moment.

Thus, distributed representations are a key part of the overall strategy to combat the [[curse of dimensionality]] in the universal computational process of [[search]]. Also see the discussion of [[combinatorial vs conjunctive]] representations for additional dimensions along which distributed representations can vary, and the computational implications thereof.

In the context of mental [[categorization]] to provide more abstract and efficient ways of representing behaviorally-relevant information, distributed representations can capture the amorphousness of a mental category, because it isn't just one single discrete factor that goes into it. There are many factors, each of which plays a role.

In other words, categorization is highly **polymorphous**: any given input can be categorized in many different ways at the same time. There is no such thing as _the_ appropriate level of categorization for any given thing. A chair can also be _furniture,_ _art,_ _trash,_ _firewood,_ _doorstopper,_ _plastic_ and any number of other such things.

Chairs have seating surfaces, and sometimes have a backrest, and typically have a chair-like shape, but their shapes can also be highly variable and strange. They are often made of wood or plastic or metal, but can also be made of cardboard or even glass. All of these different factors can be captured by the whole distributed population of neurons firing away to encode these and many other features (e.g., including surrounding context, history of actions and activities involving the object in question).

The same goes for the polymorphous nature of categories. One set of neurons may be detecting chair-like aspects of a chair, while others are activating based on all the different things that it might represent (material, broader categories, appearance, style etc). All of these different possible meanings of the chair input can be active _simultaneously_, which is well captured by a distributed representation with neurons detecting all these different categories at the same time.

The specific type of distributed representations present in the [[neocortex]] are _sparse_, with roughly 15% of the neurons active at a time. This type of representation has numerous advantages over more "dense" distributed representations that have many neurons active at a time ([[@SimoncelliOlshausen01]]; [[@OlshausenField96]]; [[@Barlow61]]). [[Inhibition]] via inhibitory interneurons is critical for creating these sparse representations.

{id="figure_dist-graded" style="height:20em"}
![Graded response as a function of similarity. This is one aspect of distributed representations, shown here in a neuron in the visual cortex of a monkey --- this neuron responds in a graded fashion to different input stimuli, in proportion to how similar they are to the thing that it responds most actively to (as far as is known from presenting a wide sample of different input images). With such graded responses ubiquitous in cortex, it follows that any given input will activate many different neuron detectors. Reproduced from Tanaka (1996).](media/fig_dist_rep_vis_bio.png)

{id="figure_maps" style="height:30em"}
![Distributed representations of different shapes mapped across regions of inferotemporal (IT) cortex in the monkey. Each shape activates a large number of different neurons distributed across the IT cortex, and these neurons overlap partially in some places. Reproduced from Tanaka (2003).](media/fig_tanaka03_topo_maps.png)

{id="figure_topo" style="height:30em"}
![Schematic diagram of topographically organized shape representations in monkey IT cortex, from Tanaka (2003) --- each small area of IT responds optimally to a different stimulus shape, and neighboring areas tend to have similar but not identical representations.](media/fig_tanaka03_topo.png)

Some real-world data on distributed representations is shown in the above Figures. [[#figure_dist-graded]] shows that individual neurons respond in a **graded** fashion as a function of **similarity** to inputs relative to the optimal thing that activates them. We saw this same property in the [[detector simulation]], especially when the leak conductance is reduced to allow firing to multiple stimuli.

[[#figure_maps]] shows an overall summary map of the topology of shape representations in monkey inferotemporal (IT) cortex, where each area has a given optimal stimulus that activates it, while neighboring areas have similar but distinct optimal stimuli. Thus, any given shape input will be encoded as a distributed pattern across all of these areas to the extent that it has features that are sufficiently similar to activate the different detectors. A more continuous transformation of different stimuli is sometimes present as well, as shown in [[#figure_topo]].

{id="figure_haxbyetal01"}
![Maps of neural activity in the human brain in response to different visual input stimuli (as shown --- faces, houses, chairs, shoes), recorded using functional magnetic resonance imaging (fMRI). There is a high level of overlap in neural activity across these different stimuli, in addition to some level of specialization. This is the hallmark of a distributed representation. Reproduced from Haxby et al. (2001).](media/fig_haxbyetal01_obj_maps.jpg)

Another demonstration of distributed representations comes from a landmark study by [[@^HaxbyGobbiniFureyEtAl01]], using functional magnetic resonance imaging (fMRI) of the human brain, while viewing different visual stimuli ([[#figure_haxbyetal01]]). They showed that contrary to prior claims that the visual system was organized in a strictly modular fashion, with completely distinct areas for faces vs. other visual categories, for example, there is in fact a high level of overlap in activation over a wide region of the visual system for these different visual inputs. They showed that you can distinguish which object is being viewed by the person in the fMRI machine based on these distributed activity patterns, at a high level of accuracy.

Critically, this accuracy level does not go down appreciably when you exclude the area that exhibits the maximal response for that object. Prior "modularist" studies had only reported the existence of these maximally responding areas. But as we know from the monkey data, neurons will respond in a graded way even if the stimulus is not a perfect fit to their maximally activating input, and Haxby et al. showed that these graded responses convey a lot of information about the nature of the input stimulus.

## Coarse coding

{id="figure_coarse-coding"}
![Coarse coding, which is an instance of a distributed representation with neurons that respond in a graded fashion. This example is based on the coding of color in the eye, which uses only 3 different photoreceptors tuned to different frequencies of light (red, green blue) to cover the entire visible spectrum. This is a very efficient representation compared to having many more receptors tuned more narrowly and discretely to different frequencies along the spectrum.](media/fig_coarse_coding.png)

[[#figure_coarse-coding]] illustrates an important specific case of a distributed representation known as **coarse coding**. This is not actually different from what we've described above, but the particular example of how the eye uses only 3 photoreceptors to capture the entire visible spectrum of light is a particularly good example of the power of distributed representations. Each individual frequency of light is uniquely encoded in terms of the _relative balance_ of graded activity across the different detectors.

For example, a color between red and green (e.g., a particular shade of yellow) is encoded as partial activity of the red and green units, with the relative strength of red vs. green determining how much it looks more orange vs. chartreuse. In summary, coarse coding is very important for efficiently encoding information using relatively few neurons.

### Coarse coding in high-dimensional spaces

When coarse coding is applied to high-dimensional spaces instead of a single linear dimension such as the frequency of light, the result is that individual neurons exhibit a pattern of **mixed selectivity** across many different dimensions ([[@FusiMillerRigotti16]]). For example, in the context of visual inputs, a given neuron might show selectivity to some aspects of color, shape, shading, curvature, depth, and motion (and random subsets thereof).

Mixed selectivity is a natural consequence of the extensive interconnectivity among neurons in the brain, so that any given neuron receives direct and indirect inputs from other neurons that have some degree of response to the relevant dimensions, and given its synaptic weights, it is thus likely to respond in somewhat complex and hard-to-describe ways as stimuli vary across these different dimensions.

This kind of neural responding is also ubiquitous in [[abstract neural network]] models, which start out with randomized initial weights, and retain a surprising amount of this randomness over the course of learning.

For example, even though neurons in primary visual cortex are known to have strong responses to relatively simple dimensions of visual input, such as orientation of edges, you can nevertheless find neurons that respond to just about any other behaviorally-relevant factor in this area too.

One critical lesson from this is that:

> any functional analysis based on the response properties of individual neurons should be regarded with a great deal of suspicion!

Instead, the resounding message from the perspective of distributed representations is that only _population-level_ analysis that takes into account the responses across many neurons in a given area can provide an accurate picture of what that  area is really representing. For example, _representational similarity analysis_ (RSA) provides a useful framework for comparing how similar the overall pattern of neural activity is for different types of inputs ([[@KriegeskorteMurBandettini08]]).

Interestingly, the use of _decoder_ models to analyze neural population activity is actually subject to many of the concerns as the analysis of individual neurons. This is because the diversity of neural responding within a given area generally allows a decoder trained to detect a particular pattern to perform above chance, even when the overall similarity structure of neural coding in a given area is not strongly aligned with that pattern.

## Localist representations

The opposite of a distributed representation is a **localist** representation, where a single neuron is active to encode a given category of information. Although we do not think that localist representations are characteristic of the actual brain, they are nevertheless often convenient to use for computational models, especially for input and output patterns to present to a network. It is often difficult to construct a suitable distributed pattern of activity to realistically capture the similarities between different inputs, so we resort to a localist input pattern with a single input neuron active for each different type of input (also known as a "one hot" encoding), and just let the network develop its own distributed representations from there.

{id="figure_halle-berry-neuron"}
![The famous case of a Halle Berry neuron recorded from a person with epilepsy who had electrodes implanted in their brain. The neuron appears sensitive to many different presentations of Halle Berry (including just seeing her name in text), but not to otherwise potentially similar people. Although this would seem to suggest the presence of localist "grandmother cells", in fact there are many other distributed neurons activated by any given input such as this within the same area, and even this neuron does exhibit some level of firing to similar distractor cases. Reproduced from Quiroga et al. (2005).](media/fig_halle_berry_neuron.jpg)

[[#figure_halle-berry-neuron]] shows the famous case of a "Halle Berry" neuron, recorded from a person with epilepsy who had electrodes implanted in their brain ([[@QuirogaReddyKreimanEtAl05]]). This would appear to be evidence for an extreme form of localist representation, known as a **grandmother cell** (a term apparently coined by Jerry Lettvin in 1969), denoting a neuron so specific yet abstract that it only responds to one's grandmother, based on any kind of input, but not to any other people or things. People had long scoffed at the notion of such grandmother cells.

Even though the evidence for such neurons is fascinating (including other neurons for Bill Clinton and Jennifer Aniston), it does little to change our basic understanding of how the vast majority of neurons in the cortex respond. Clearly, when an image of Halle Berry is viewed, a huge number of neurons at all levels of the cortex will respond, so the overall representation is still highly distributed. But it does appear that, amongst all the different ways of categorizing such inputs, there are a few highly selective "grandmother" neurons! One other outstanding question is the extent to which these neurons actually do show graded responses to other inputs --- there is some indication of this in the figure, and more data would be required to really test this more extensively.


