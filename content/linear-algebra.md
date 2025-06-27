+++
Categories = ["Computation"]
bibfile = "ccnlab.json"
+++

The language of **linear algebra** provides an important set of concepts for understanding the behavior and learning in neural networks, and also important analysis techniques including linear regression and [[principal components analysis]]. See also [[information theory]] for a similarly useful framework.

For example, the state of activity in a network layer can be thought of as an n-dimensional _vector_ of numbers representing the activity of each of the _n_ neurons in the layer. We can write this vector using bold lower-case letters, with **x** being the input layer and **y** (of dimension _m_) being the hidden layer above that. The set of weights interconnecting the two layers, **W**, is thus an _n x m_ _matrix_ (confusingly, _n_ = number of columns, _m_ = number of rows, but indexes are written in the opposite order).

{id="figure_matrix-mult"}
![Matrix multiplication of n=3 vector of input-layer activations x, by n x m matrix of weights W, gives the linear activity of m=2 hidden-layer activations y. Matrix multiplication is tricky because it generally operates on the right and is not commutative, so the order matters. x and y are represented as column vectors to make everything work correctly. As coded by the colors, the first hidden unit in y is computed as the dot product of the first row of the weight matrix, with each column applying to a different input unit in turn. The index i is used conventionally for the input layer in neural networks, but typically in linear algebra this would be a j for the column-wise variable of a matrix.](media/fig_matrix_multiplication_nnet.png)

[[#figure_matrix-mult]] shows how the matrix multiplication of the input vector **x** times the weight matrix **W** gives the linear activations of the hidden layer neurons, as in the following equation:

{id="eq_act_wt" title="Matrix multiplication"}
$$
{\bf y} = {\bf W} {\bf x}
$$

Each row of the weight matrix _W_ contains a vector of weights for each receiving neuron in the hidden layer, and the operation of matrix multiplication is just the accumulation of the individual _dot product_ (sum of element-wise products) multiplications of these two vectors:

{id="eq_act_wt_dot" title="Dot product"}
$$
y_j = \sum_i^n x_i W_{ji}
$$

This dot product should be familiar from the [[neuron]] chapter section on [[neuron#Computing input conductances]], and biologically corresponds to the total conductance caused by AMPA excitatory channels opening in response to the rate of sending neuron spiking (and thus glutamate neurotransmitter release), with the weight value corresponding to the number and efficacy of these AMPA receptors at each synapse.

## Dot products are projections

{id="figure_face-dim-prjn"}
![How synaptic weights act to project input patterns along specific dimensions or bases, in this case projecting the inputs along the dimensions of Emotion and Gender. In the left panel, the very high-dimensional face inputs (256 dimensions for a 16x16 image) are projected along two random weight vectors, allowing us to visualize this high-dimensional input space in a 2D plot. In the right panel, the specific synaptic weights trained for discriminating along the emotion vs. gender dimensions have transformed or rotated the input space into a much more systematic and well-organized, low-dimensional space. This is fundamentally what neurons do: organize and transform input patterns along relevant dimensions, and that is another way of stating that neurons detect stimuli along these dimensions. ](media/fig_face_categ_dim_prjn.png)

The dot product operation (also known as the _inner product_ or _scalar product_) takes two vectors and turns them into a single scalar number representing the extent to which the two vectors "align" with each other. This is known as a **projection** computation in linear algebra terms: you are effectively seeing one vector "through the lens" of the other vector. Thus, the receiving neuron sees the input vector "through the lens" of its synaptic weights for that input layer, as explored in the [[detector simulation]]. Anywhere there are zero values in the weights, the input values are "filtered out" and are irrelevant, while strong weights make those inputs highly important for driving the receiver's activity.

[[#figure_face-dim-prjn]] illustrates this projection operation in the context of the [[faces simulation]], as discussed in [[categorization]]. This projection operation organizes and systematizes the inputs along dimensions of behavioral importance, for example projecting a face input along dimensions of emotion and gender in the case shown in the figure, which you can explore in the [[faces simulation]].

Another linear algebra framing of this operation is in terms of the synaptic weight matrix being composed of **basis vectors**, where each row of the matrix defines a different _basis_ or _axis_ of a new, _rotated_ version of the space defined by the the input activity vector. Thus, the hidden layer neurons are creating a _transformation_ of the input space into this new rotated space, which amplifies certain things while filtering out others. 

## Eigenvectors and singular value decomposition

The process of [[principal components analysis]] (PCA) provides a way of selecting a new set of basis vectors to project the input space into, based on the linear algebra process of **eigenvector** decomposition, where the eigenvector is a special vector that keeps pointing in the same direction when multiplied by the matrix that it is an eigenvector of:

{id="eq_eigen" title="Eigenvectors and values"}
$$
{\bf A}{\bf v} = \lambda {\bf v}
$$

where **A** is a matrix, and **v** is an eigenvector of that matrix because its multiplication produces the same vector again, scaled by a scalar value $\lambda$ which is the associated **eigenvalue** for that eigenvector **v**.

The eigenvectors with the largest eigenvalues provide the most "information" about matrix **A**, and thus a typical use of PCA is for _dimensionality reduction_ to represent a much higher dimensional space in terms of its strongest 2 eigenvectors, which can then be plotted in a 2D plot. Otherwise, visualizing large n-dimensional spaces can be hard.

Interestingly, [[Hebbian learning]] can be shown mathematically to be performing PCA, providing a powerful conceptual understanding of what it might be doing: learning new efficient ways of representing the input activity space.

## Least squares regression

One of the most basic problems in linear algebra is to solve the _least squares_ problem:

{id="eq_lsq" title="Least squares"}
$$
\rm{minimize} \; \|{\bf W} {\bf x} - {\bf t}\|^2
$$

which means finding the values of the matrix **W** that minimizes the differences between _target_ activation values **t** and the **y** activation values that the network computes (using [[#eq_act_wt]]), across all of the relevant input activation states of **x**. This is the problem that _linear regression_ solves, and is another way of expressing the basic goal of [[error-driven learning]], and [[error backpropagation]] provides an iterative solution to this same problem in a two-layer linear network of this form.

Thus, two major forms of neural network learning, Hebbian and error-driven, can be understood in simple, essential form using the tools of linear algebra.


