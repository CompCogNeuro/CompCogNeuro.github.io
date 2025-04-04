+++
Categories = ["Axon"]
bibfile = "ccnlab.json"
+++

**Artificial intellignce** (AI) has been used to describe a very wide range of different computational frameworks over the years, including symbolic reasoning systems developed from the 1950s to the 1990s, and the current [[abstract neural network]] models including [[large language models]] (LLMs).

The question of what "counts" as "intelligent" is tricky, because mostly we use humans as the only available point of reference, while also not wanting to be too directly tied to the quirks of our own particular brand of intelligence. The attempt to define a clear set of criteria for what counts as intelligent is notoriously difficult, with the Turing test providing an unsatisfyingly pragmatic approach: if we can't tell the difference between it and a human, it must be intelligent.

One productive distinction that is relatively clear, at least in principle, is between "narrow" vs. "general" forms of intelligence, with most of the "new AI" models being of the more narrow variety. For example, the _Deep Blue_ chess playing system created by IBM in 1997 was clearly very narrow, being exclusively designed to play chess. Most other game playing models have likewise been optimized for specific games, even when they also use domain-general learning mechanisms. Likewise, most other deep neural networks are only good at specific tasks, such as image recognition.

Thus, the new goal in AI is to accomplish artificial _general_ intelligence (AGI), where the system achieves something approaching a human level of generalized intelligence across a wide range of domains. Some would argue that LLMs have accomplished AGI, or are at least very close. However, the counter-argument is that they really don't have much "true" intelligence of their own, and instead they have just absorbed a significant amount of human intelligence as represented in the collective output of human writing that they have been trained on. Given the tremendous scope of what these models have been trained on, it is difficult to tell the difference between a mere "reshuffling" of the training material versus something that we would consider to be a sign of truly intelligent cognition.

This line of thinking suggests that _out-of-domain generalization_ is a necessary hallmark of AGI (e.g., [[@Chollet19]]; [[@RussinJoOReillyEtAl19]]). Here, the evidence is less clear for LLMs, when strong steps are taken to eliminate potential contamination from the training set. For example, a recent study examined performance on recent math competitions performed after the collection of the model's training data, and performance is overall relatively poor ([matharena.ai](https://matharena.ai/)).

At a very basic level LLMs are fundamentally limited by the fact that they lack significant online learning abilities, and thus cannot shape entirely new representational systems in novel domains. By contrast, humans are not actually very good at _zero shot_ generalization to novel domains, but instead we excel at _learning_ and goal-driven problem solving in novel domains.

Another obvious failing of current LLMs is that they are almost entirely dependent on external input from humans to shape their behavior. They don't spend their idle moments daydreaming about interesting reflections based on everything they know, devising interesting new problems to think about or discover the answers to. Fundamentally, they don't know what they know and what they don't know, and thus have no ability to formulate problem solving strategies focused on figuring out what they need to know to solve a problem. They have been trained to imitate human reasoning and this improves performance to some extent, but ultimately, like so much with LLMs, it is just an imitation of true human problem solving and intelligence.

## Axon general intelligence goals

The [[Rubicon]] framework provides a clear set of hypotheses about how goal-driven problem solving abilities can emerge through interacting brain networks, based on the overall Axon model. Some key features of this approach are as follows:

* General intelligence emerges through the interaction of two main systems: a rich and flexible _semantic memory_ system that encodes all manner of "common sense" knowledge about the world at many different levels of abstraction (which is reasonably well captured by an LLM), and a goal-driven _executive control_ system that can use [[conscious awareness]] (mediated via [[bidirectional connectivity]]) to inspect and manipulate the state of this semantic memory system in order to reason and problem solve. These two systems depend on different brain areas organized by evolution and learning to optimize different functional properties, which have been studied by cognitive neuroscientists for many years. A particularly unimaginative but nevertheless influential terminology describes these as System 1 and System 2 ([[@Kahneman11]]). LLMs, and essentially any other ANN framework, are lacking a well-developed model of this executive control system, which is the focus of the [[Rubicon]] model.

* The goal-driven executive function system drives learning throughout the brain, including the semantic memory system, to learn new representations that can help solve problems in novel domains. This "grit" ability to persevere through multiple failed attempts, driving invaluable learning, is a critical element of most examples of true innovation and ultimate success in the human sphere. Such qualities are simply missing in existing AI models.

* Both the semantic memory and executive control networks depend critically on bidirectional [[constraint satisfaction]] processing to efficiently search through vast reservoirs of knowledge to find the most relevant information for the task at hand. An LLM can approximate this kind of processing, but only in a feedforward direction, making it impossible to simultaneously integrate top-down and bottom-up constraints to find solutions that make sense according to the sensory inputs _and_ according to the current goals.

In summary, as postulated by [John McCarthy](http://jmc.stanford.edu/artificial-intelligence/what-is-ai/index.html) (the self-described "father of AI"):

> Intelligence is the computational part of the ability to achieve goals in the world.

And yet, most existing AI systems lack any significant goal-driven executive control and learning functionality, so it is unclear how they could achieve "true" intelligence. While the Axon framework has yet to demonstrate true intelligence either, at least it has a clear sense of some of the necessary ingredients and can work toward understanding these aspects better.

