+++
Categories = ["Computation", "Cognition", "Learning"]
bibfile = "ccnlab.json"
+++

**Generalization** is one of the most important concepts in both [[computation]] and [[cognition]], referring to the ability to apply prior learning to a novel situation. It is also known as **transfer** of learning in the psychological literature, and is closely related to **generativity** (the ability to generate novel behaviors or outputs) and **systematicity** (the ability to behave in a rule-like manner, e.g., to apply the rules of grammar to nonwords or sentences that don't otherwise make sense), both of which are discussed here.

Without the ability to generalize, animals in the real world would be much less able to survive, because the world presents us with a constantly varying set of situations and challenges. Humans are particularly good at some forms of generalization, and the "generalization gap" between us and extant [[artificial intelligence]] models motivates the quest for [[artificial intelligence#artificial general intelligence]] (AGI) (e.g., [[@Chollet19]]). As of this writing, humans are still the only universally recognized exemplar of an AGI system.

## Types of generalization

{id="figure_generalizations" style="height:15em"}
![Domains of generalization, visualized in geographical terms where the training data lives in a convex shaped subspace, e.g., an island. Near transfer can occur via interpolation, where there is sufficient data surrounding the test items that weighted averages of nearby training items sufficies. This is known as independent and identically distributed (iid) test data -- it comes from the same probability distribution that created the training data. Geographically, it is like a small lake on the island, surrounded by training data. Far trasnsfer, represented by points completely outside the domain (convex hull) of the training data (i.e., in the "ocean" beyond the island), requries some kind of extrapolation, and is much more difficult. A small cove or bay on the edge of the island (on the edge of the training data) represents an intermediate "edge" case. Concept of the figure due to Alex Petrov.](media/fig_generalization_iid_ood_lake_ocean.png)

[[#figure_generalizations]] shows three different types of generalization relative to a given set of training data, which is assumed to lie within some kind of overall "perimeter" within a larger high-dimensional space. Relatively easy forms of generalization involve test cases that lie solidly within this perimeter, which are characterized in the statistics literature as _i.i.d._ = _independent and identically-distributed_ test items. These can be solved by interpolating across close training exemplars, for example.

The much more challenging and interesting forms of generalization involve the _out of domain_ (_o.o.d_) test cases, that lie out in the "ocean" beyond the "island" of the training data. This is known as **far transfer** in the psychological literature, and it has been the focus of numerous studies on the generalization of learning and training techniques. In general, the results here have been disappointing in many domains, particularly in the domain of [[executive function]], where far transfer is highly variable across studies and the subject of considerable controversy ([[@KarbachKray21]]).

In general, AI models have also shown very limited abilities to perform far transfer, and often require massive amounts of training to even produce reasonable levels of near transfer (e.g., [[@AlbrechtFettermanFogelmanEtAl22]]; [[@Chollet19]]). Assessing the extent of generalization in [[large language models]] is complicated significantly by the vast datasets that these models are trained on, and attempts to find generalization tests that are truly outside of the training set suggest that their far transfer abilities are more limited than their ability to flexibly combine elements from within their extensive training set (e.g., [matharena.ai](https://matharena.ai/)).

In addition to these domain-based distinctions in types of generalization, there is an important dimension that has to do with _generativity_: does the generalization test involve a novel response or not? If you have to make up a new song or poem, that is very different from just recognizing a strange-looking dog as being a member of the dog [[categorization|category]], which you already have a concept for.

In general, generativity requires forming novel combinations of existing representational elements, i.e., a _combinatorial_ code, as discussed in [[combinatorial vs conjunctive]] representations.

## Basis for generalization

What makes networks of neurons able to generalize their learning? The following are some well-established mechanisms:

* **Interpolation** works by taking a weighted average of the "votes" from nearby representations learned from the training set, which will generally work well if the testing items are sufficiently densely surrounded by training items (i.e., "lakes", not "oceans" in [[#figure_generalization]]).

* **Abstraction**: As discussed at length in the [[categorization]] page, the formation of a hierarchy of increasingly abstract detectors can enable a novel input to be recognized as a member of an existing category. Once categorized, any existing associations between that category and appropriate behaviors or other mental representations can then be activated.

* **Combinatorial codes**: [[Distributed representations]] can encode different _separable_ aspects using independent representational features, such that novel inputs are encoded using novel _combinations_ of these existing representations. For example, the visual system has neurons that separately represent color vs. shape features, so that it can automatically generalize to novel combinations of shape and color. This is possible because these dimensions are readily extracted in distinct ways, and the natural world provides significant examples of the relative independence of these dimensions. However, combinatorial codes suffer from the [[binding problem]], which creates other challenges as discussed in [[combinatorial vs conjunctive]].

* **Relational** generalization depends on abstraction of relationships among entities, so that these relationships can then be applied to novel domains, also known as _systematicity_ ([[@OReillyRanganathRussin22]]; [[@WebbFranklandAltabaaEtAl24]]). One widely-studied example is [[analogical reasoning]], where the relationships in a "source" domain are applied to those in a novel "target" domain ([[@Gentner83]]; [[@Holyoak12]]). For example, the source domain could be planets _orbiting_ a star, applied to the target domain of electrons _orbiting_ the nucleus in an atom ([[#figure_structural]]). These kinds of abstract relational representations are also known as _structure sensitive_ representations or processes, and are a central feature of symbolic AI models, in the form of _propositional_ representations ([[@FodorPylyshyn88]]; [[@AndersonLebiere98]]). These representations form the basis of _rule-based_ generalization, where a generalizable rule captures necessary relationships among entities in abstract terms. Relational representations are known to exist in the [[parietal]] lobe ([[@OReillyRanganathRussin22]]; [[@SummerfieldLuyckxSheahan20]]).

{id="figure_structural" style="height:20em"}
![Far transfer to o.o.d. cases can happen via structural extrapolation based on a "bridging" concepts that allow familiar relationships to be applied to enties well outside the training set.  In this illustration, the analogical structure of orbiting allows us to use more familiar concepts to extrapolate to the invisible atomic world of electrons orbiting the nucleus.](media/fig_generalization_bridge_ood.png)

As illustrated conceptually in [[#figure_structural]], successful far transfer requires a lot of cognitive work. You first have to build the relevant structural representations (the "relational bridge") within the scope of your existing knowledge, and then you have to perform the _structure mapping_ process ([[@Gentner83]]) to traverse this bridge out into the wild unknown "ocean" beyond the safety of your existing knowledge island. Then you have to figure out if it is a bridge to nowhere, or whether it really provides a solid portal into this new realm. In the case of the atomic system, the simple orbiting model only partially applies, and many strange and new concepts from quantum mechanics are instead required to accurately understand how things work in this mysterious, invisible microscopic realm.

## Math supports generalization

In the case of science, a huge amount of effort has gone into the development of _mathematics_ as the primary foundational bridging structure that allows us to understand all manner of systems great and small. Math has all of the above properties that support generalization:

* Numbers are the ultimate abstraction, preserving only the abstract property of _quantity_ and discarding everything else.

* Mathematical operations can be combined in endless ways to tackle novel problems.

* Many aspects of math provide precise and general ways of capturing relationships (greater than, less than, equality, set theory, etc).

Cognitive neuroscience research has shown that areas of the parietal lobe support many of these representations in the brains of humans and other animals. For example a kind of "mental number line" has been characterized in the parietal lobe of multiple species from humans to primates to rodents ([[@DehaeneDehaene-LambertzCohen98]]; [[@NiederMiller04]]; [[@NiederDehaene09]]). More generally, mathematical reasoning is thought to leverage the extensive spatial representations that develop in the parietal lobe ([[@UngerleiderMishkin82]]; [[@OReilly10]]).

<!--- TODO: point to models of spatial learning here -->

## Programming supports generalization

Computer programs are another important source domain for generalization, and provide an external tool that directly enables systematic, generalizable behavior to be implemented. Indeed computational simulations of brains using computer programs is the entire foundation of the effort undertaken here. These simulations would be impossible for any individual to perform on paper, much less in their head.

All programming languages provide strong relational abstractions, in the form of types, functions, and structures, that then support strongly systematic generalized functionality. For example, functions specify arguments that can take on any value, subject to constraints imposed by the type system:

```Go
Add := func(a, b int) int {
    sum := a + b
    fmt.Printf("%d + %d = %d\n", a, b, sum)
    return sum
}

Add(2, 3)
Add(5, 7)
```

Thus, this one function can operate on any two arguments, providing massive levels of systematicity and generalization. The types can be further qualified and specified to more strongly constrain the relationships among variables, etc.

Consistent with these properties, we argue that the ability to be [[self-programmable]] at least to some extent is an essential property for supporting systematic out-of-domain generalization abilities in the human brain. We can use our [[language]] capacities to talk ourselves through extended sequences of actions, and remember key intermediate values, enabling us to solve complex novel problems that otherwise would be impossible in a single parallel step of standard neural processing.


