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

In general, AI models have also shown very limited abilities to perform far transfer, and often require massive amounts of training to even produce reasonable levels of near transfer (e.g., [[@AlbrechtFettermanFogelmanEtAl22]]). The extent of generalization in [[large language models]] is complicated significantly by the vast training sets that these models are trained on, and attempts to disentangle this suggest that their far transfer abilities are much more limited than their ability to flexibly combine elements from within their extensive training set (e.g., [matharena.ai](https://matharena.ai/)).

In addition to these domain-based distinctions in types of generalization, there is an important dimension that has to do with _generativity_: does the generalization test involve a novel response or not? If you have to make up a new song or poem, that is very different from just recognizing a strange-looking dog as being a member of the dog [[categorization|category]], which you already have a concept for.

In general, generativity requires forming novel combinations of existing representational elements, i.e., a _combinatorial_ code, as discussed in [[combinatorial vs conjunctive]] representations.

## Basis for generalization

What makes networks of neurons able to generalize their learning? The following are some well-established mechanisms:

* **Interpolation** works by taking a weighted average of the "votes" from nearby representations learned from the training set, which will generally work well if the testing items are sufficiently densely surrounded by training items (i.e., "lakes", not "oceans" in [[#figure_generalization]]).

* **Abstraction**: As discussed at length in the [[categorization]] page, the formation of a hierarchy of increasingly abstract detectors can enable a novel input to be recognized as a member of an existing category. Once categorized, any existing associations between that category and appropriate behaviors or other mental representations can then be activated.

* **Combinatorial codes**: [[Distributed]] representations can encode different _separable_ aspects using independent representational features, such that novel inputs are encoded using novel _combinations_ of these existing representations. For example, the visual system has neurons that separately represent color vs. shape features, so that it can automatically generalize to novel combinations of shape and color. This is possible because these dimensions are readily extracted in distinct ways, and the natural world provides significant examples of the relative indpendence of these dimensions.

* **Rules**: This is the most complex case, requiring extraction and formulation of a rule-like representation, and the ability to generically apply such a rule to novel inputs. One instance of this kind of operation is [[analogical reasoning]], where _relationships_ in one "source" domain can be applied to those in a novel "target" domain. Thus, rule-based generalization depends critically on extracting _relational_ representations ([[@OReillyRanganathRussin22]]; [[@WebbFranklandAltabaaEtAl24]]).

<!--- TODO: figure showing building a canatlevered bridge out into the ocean, and example of using orbiting analogy to understand invisible atomic realm -->



