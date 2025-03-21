+++
Categories = ["Leabra"]
+++

We recommend that you run the simulation exercises associated with this book on the web using the links in the list of sims below.

**IMPORTANT:** This currently only fully works on recent versions of Chrome on macOS and Windows, so if you are on Linux, please see the alternative instructions at the bottom of this page.

All bug reports relating to the simulations should be filed in the [issue tracker](https://github.com/CompCogNeuro/sims/issues).

If you want more background information and the underlying code, the simulations are implemented with the Go version of the [emergent](https://github.com/emer/emergent) framework in the [sims](https://github.com/CompCogNeuro/sims) repository using [Cogent Core](https://cogentcore.org/core). This website is also built with Cogent Core.

## Usage

Each simulation has a `README` button, which directs your browser to open the corresponding `README.md` file on github.  This contains full step-by-step instructions for running the model, and questions to answer for classroom usage of the models.  See your syllabus etc for more info.

You can use standard `Ctrl+` and `Ctrl-` key sequences to zoom the display to the desired scale, and the Cogent Core settings window has more display options such as light/dark mode (this can be accessed through right click / two finger click / control click).

The main actions for running are in the `Toolbar` at the top, while the parameters of most relevance to the model are in the `Control panel` on the left.  Different output displays are selectable in the `Tabbed views` on the right of the window.

The [Go Emergent Wiki](https://github.com/emer/emergent/wiki/Home) contains various help pages for using things like the `NetView` that displays the network.

You can always access more detailed parameters by clicking on the button to the right off `Net` in the control panel (also by clicking on the layer names in the NetView), and custom params for this model are set in the `Params` field.

## List of sims and exercise questions

Here's a full list of all the simulations and the textbook exercise questions associated with them.

## Chapter 2: Neuron

* [neuron](https://compcogneuro.org/sims/ch2/neuron): Integration, spiking and rate code activation. (Questions **2.1 -- 2.7**)

* [detector](https://compcogneuro.org/sims/ch2/detector): The neuron as a detector -- demonstrates the critical function of synaptic weights in determining what a neuron detects. (Questions **2.8 -- 2.10**)

## Chapter 3: Networks

* [faces](https://compcogneuro.org/sims/ch3/faces): Face categorization, including bottom-up and top-down processing (used for multiple explorations in Networks chapter) (Questions **3.1 -- 3.3**)

* [cats_dogs](https://compcogneuro.org/sims/ch3/cats_dogs): Constraint satisfaction in the Cats and Dogs model. (Question **3.4**)

* [necker_cube](https://compcogneuro.org/sims/ch3/necker_cube): Constraint satisfaction and the role of noise and accommodation in the Necker Cube model. (Question **3.5**)

* [inhib](https://compcogneuro.org/sims/ch3/inhib): Inhibitory interactions via inhibitory interneurons, and FFFB approximation. (Questions **3.6 -- 3.8**)

## Chapter 4: Learning

* [self_org](https://compcogneuro.org/sims/ch4/self_org): Self organizing learning using BCM-like dynamic of XCAL (Questions **4.1 -- 4.2**).

* [pat_assoc](https://compcogneuro.org/sims/ch4/pat_assoc): Basic two-layer network learning simple input/output mapping tasks (pattern associator) with Hebbian and Error-driven mechanisms (Questions **4.3 -- 4.6**).

* [err_driven_hidden](https://compcogneuro.org/sims/ch4/err_driven_hidden): Full error-driven learning with a hidden layer, can solve any input output mapping (Question **4.7**).

* [family_trees](https://compcogneuro.org/sims/ch4/family_trees): Learning in a deep (multi-hidden-layer) network, showing advantages of combination of self-organizing and error-driven learning (Questions **4.8 -- 4.9**).

* [hebberr_combo](https://compcogneuro.org/sims/ch4/hebberr_combo): Hebbian learning in combination with error-driven facilitates generalization (Questions **4.10 -- 4.12**).

Note: no sims for chapter 5

## Chapter 6: Perception and Attention

* [v1rf](https://compcogneuro.org/sims/ch6/v1rf): V1 receptive fields from Hebbian learning, with lateral topography. (Questions **6.1 -- 6.2**)

* [objrec](https://compcogneuro.org/sims/ch6/objrec): Invariant object recognition over hierarchical transforms. (Questions **6.3 -- 6.5**)

* [attn](https://compcogneuro.org/sims/ch6/attn): Spatial attention interacting with object recognition pathway, in a small-scale model. (Questions **6.6 -- 6.11**)

## Chapter 7: Learning and Memory

* [abac](https://compcogneuro.org/sims/ch7/abac): Paired associate AB-AC learning and catastrophic interference. (Questions **7.1 -- 7.3**)

* [hip](https://compcogneuro.org/sims/ch7/hip): Hippocampus model and overcoming interference. (Questions **7.4 -- 7.6**)

* [priming](https://compcogneuro.org/sims/ch7/priming): Weight and Activation-based priming. (Questions **7.7 -- 7.8**)

## Chapter 8: Motor Control and Reinforcement Learning

* [bg](https://compcogneuro.org/sims/ch8/bg): Action selection / gating and reinforcement learning in the basal ganglia. (Questions **8.1 -- 8.3**)

* [rl](https://compcogneuro.org/sims/ch8/rl): Pavlovian Conditioning using Temporal Differences Reinforcement Learning. (Questions **8.4 -- 8.5**)

* `pvlv`: Pavlovian Conditioning with the PVLV model (Questions **8.6 -- 8.8**)  **NOT YET AVAIL!**

* `cereb`: Cerebellum role in motor learning, learning from errors. (Questions **8.9 -- 8.10**) **NOT YET AVAIL!**

## Chapter 9: Executive Function

* [a_not_b](https://compcogneuro.org/sims/ch9/a_not_b): Development of PFC active maintenance and the A-not-B task (Questions **9.1 -- 9.3**)

* [stroop](https://compcogneuro.org/sims/ch9/stroop): The Stroop effect and PFC top-down biasing (Questions **9.4 -- 9.6**)

* [sir](https://compcogneuro.org/sims/ch9/sir): Store/Ignore/Recall Task - Updating and Maintenance in more complex PFC model (Questions **9.7 -- 9.8**)

## Chapter 10: Language

* [sem](https://compcogneuro.org/sims/ch10/sem): Semantic Representations from World Co-occurrences and Hebbian Learning. (Questions **10.1 -- 10.3**)

* [ss](https://compcogneuro.org/sims/ch10/ss): Orthography to Phonology mapping and regularity, frequency effects. (Questions **10.4 -- 10.5**)

* [dyslexia](https://compcogneuro.org/sims/ch10/dyslexia): Normal and disordered reading and the distributed lexicon. (Questions **10.6 -- 10.11**)

* [sg](https://compcogneuro.org/sims/ch10/sg):  The Sentence Gestalt model. (Question **10.12**)

## Alternative ways to run

### Build from source

To run the sims locally on your computer on any platform, you must first follow the [Cogent Core setup instructions](https://www.cogentcore.org/core/setup/install). Then, you can clone the sims repository and run sims using `core run`:

```sh
git clone https://github.com/CompCogNeuro/sims
cd sims/ch2/neuron # or any other sim
core run
```

You can also use `core run web` to run a sim on the web, which does not require running the `core setup` command in the setup instructions.

### Prebuilt executables

We will provide updated prebuilt versions of the sims soon. You can see old prebuilt versions in the [releases](https://github.com/CompCogNeuro/sims/releases), which are not recommended. The even older [C++ emergent (cemer)](https://github.com/emer/cemer) sims project files are available here: [cecn_8_5_2.zip](https://github.com/CompCogNeuro/sims/releases/download/v1.2.2/cecn_8_5_2.zip) (no longer updated or supported; recommend transitioning to new ones).
