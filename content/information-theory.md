+++
Categories = ["Computation"]
bibfile = "ccnlab.json"
+++

The framework of **information theory**, like that of [[linear algebra]], provides a number of central concepts for understanding what neural networks do.

The basic mathematical definition of _information_ was given by Claude Shannon ([[@Shannon48]]), in terms of the concept of _entropy_, which is equivalent to _uncertainty_ in a probabilistic context. Intuitively, a situation with maximum uncertainty is also the one in which you can _obtain the most information_ by learning anything more about that situation. Conversely, if you already know everything about a situation, then you can't get any more information about it.

Thus, the amount of information contained in a message is a function of the level of uncertainty reduction imparted by the receipt of that message.

This can be applied to binary codes, where numbers are represented by 1's and 0's. For example, the number 5 is 101, because the first binary digit 1 represents the number 4, the second digit (0) represents the number 2, and the last digit represents the number 1, so 1*4 + 0*2 + 1*1 = 5. If you see any one of these digits out of a 3-digit message, it reduces the uncertainty by a factor of 2, because there are 8 different numbers that could possibly be represented by 3 binary digits (0 = 000 to 7 = 111), and seeing one of the digits is like dividing that number by 2. If you saw that the first digit was a 0, then that would restrict the remaining uncertainty to the values between 0 and 3, for example.

If you introduced a fourth binary digit that also represented the value 4, then this redundant digit would not add any new information, in the sense that if you already knew all of the other 3 digits, learning the value of this 4th digit would not reduce any uncertainty; you would already know its value.

Similarly, if each value of the binary digit was not equally likely to occur across a set of messages (e.g., it was very rare to get a value above 3, i.e., a 1 in the 4-value digit), then the total information in such a message would actually be reduced, because of the reduction in uncertainty associated with this digit. Every time you got a message where the value of that digit was a 0, you would just say "yeah I knew that already", and only on the rare occasion when it was a 1 would you be surprised.

Even though it might seem like the surprise associated with getting that rare 1 would outweigh all the other boring expected cases, it turns out that the mathematics of entropy show that they don't: the maximum information is always when each option is equally likely.

Mathematically, the expression for _entropy_ explains all of the above:

{id="eq_entropy" title="Entropy"}
$$
H = - \sum_i p_i \log_2(p_i)
$$

where the index $i$ is over each discrete possible outcome out of an enumerable set of such possibilities. For example, the expression for a single binary value that can either be a 0 or a 1, with probability $p$ of being a 1, is:

{id="eq_entropy-bit" title="Entropy of a bit"}
$$
H = - p \log_2(p) - (1-p) \log_2(1-p)
$$

It turns out that the maximum value of this expression occurs when $p = 0.5$, i.e., it is equally likely to be 0 or 1, and it falls off symmetrically on either side. The numerical value of this expression at $p = 0.5$ is 1, which means there is 1 _bit_ of information, because we're using the $\log_2$ base 2 (binary) logarithm. Thus, only when a bit has equally-probable 0 and 1 states does it truly contain 1 bit of information!

The entropy measure can be usefully applied to quantifying the information content of a population of neurons in a layer. The overall information content is maximized when the probability of each neuron becoming active is at 0.5, and there are no correlations between any of the neurons. This is not at all what neural activity looks like in the [[neocortex]]: most neurons are essentially inactive most of the time, with a roughly 15% chance of being significantly activated. Furthermore, there are often strong correlations across different neurons. This tells us that these neurons are not maximizing the amount of information that can be represented, and therefore there must be other important criteria driving the nature of these representations.

For example, correlations help convey information about relationships, which is useful for making inferences, and sparse codes with low overall activity levels are more energy efficient, and better match the statistics of the natural world, where many things occur with lower probabilities.

Nevertheless, if a form of [[principal components analysis]] is operating in the brain, it will tend to reduce the correlations among neurons and increase the variance in their firing, so perhaps there are multiple conflicting forces at work, and what we observe is a balance of those forces.

## Mutual information

The concept of _mutual information_ is very useful for quantifying the relationship between two signals, in terms of the amount of uncertainty left over about one signal after receiving the other. This has often been used to analyze the relationship between neural representations in different layers, or among neurons within one layer. For example, the computation of _independent components analysis (ICA)_ ([[@JuttenHerault91]]) involves minimizing mutual information, which should make each component maximally independently informative.

Mathematically, mutual information is defined as:

{id="eq_mi" title="Mutual information"}
$$
I(X;Y) = \sum_{x,y} p(x,y) \log \frac{p(x,y)}{p(x) p(y)}
$$

where $p(x,y)$ is the _joint probability_ of two events occuring together, which is being normalized by the _independent probabilities_ of each event on its own ($p(x) p(y)$). Intuitively, if the joint probability is the same as the independent probabilities, then there is no special mutual information between those events. But if the joint probability is much higher than the independent probability, then they have high mutual information.

Recall that taking the logarithm of a fraction is equivalent to _subtracting_ the logs of the numerator minus the denominator, because the exponent on the denominator is -1, so the right-side factor with the ratio will end up being 0 if both values are the same.

## KL divergence (information gain)

Another central concept in information theory that has been widely applied in [[abstract neural network]] models is the Kullback-Leibler divergence, which is a way of comparing two different probability distributions $p$ and $q$:

{id="eq_kl" title="KL divergence"}
$$
KL(p,q) = \sum_x p(x) \log \frac{p(x)}{q(x)}
$$

This is of a similar form as the mutual information, and can be understood in the same terms: if $p$ and $q$ are basically the same across values of x in the distribution, then the values in the log will cancel out and result in a 0 difference.

Error-driven learning has been formulated in terms of the KL divergence, where $p$ is the target, correct probability distribution and $q$ is the current one produced by the network, and learning is trying to reduce the KL divergence between the two.

