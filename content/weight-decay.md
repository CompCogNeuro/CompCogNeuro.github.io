+++
Categories = ["Learning", "Mechanisms"]
bibfile = "ccnlab.json"
+++

In many [[abstract neural network]] models, **weight decay** is often used as a _regularization_ factor that imposes _a-priori_ constraints or biases on [[error backpropagation learning]]. In the context of the [[bias-variance tradeoff]], the weight decay bias reduces the variance across different learning outcomes, favoring solutions where the weight values are generally as small as possible while still being sufficient to solve the problem.

Weight decay has also long been used in statistical regression, which is mathematically equivalent to a 2-layer neural network with a linear activation function. Specifically, [lasso](https://en.wikipedia.org/wiki/Lasso_(statistics)) regression uses weight decay based on the L1 norm (absolute value of the weight), while [ridge](https://en.wikipedia.org/wiki/Ridge_regression) regression uses weight decay based on the L2 norm (squared value of the weights) (wikipedia links).

In the [[Leabra]] model, [[Hebbian learning]] is used as a regularizing bias, similar to weight decay.

