+++
Name = "LSTM"
Categories = ["Computation"]
bibfile = "ccnlab.json"
+++

The **LSTM** (long short-term memory) model of [[@^HochreiterSchmidhuber97]] uses multiple forms of multiplicative **gating** to implement a particularly robust yet flexible form of activation-based working memory in an [[abstract neural network]] trained with [[error backpropagation]].

{id="figure_lstm" style="height:10em"}
![The LSTM memory cell (rectangle) with constant error carousel (CEC; circle with diagonal chord). Maintenance gating g (left side) multiplicatively determines when to update and encode new information, while output gating (right side) determines when to read information out of the memory cell to drive behavior. From Hochreiter & Schmidhuber, 1997, Figure 1.](media/fig_hochreiter_schmidhuber_97_lstm.png)

The key insight that motivated the model was that the ability to maintain a trace of information robustly across time in the form of neural activity signals requires a precisely-balanced form of recurrent neural activity that avoids exponential increases or decay across time. This was termed the _constant error carousel_ (CEC) ([[#figure_lstm]]).

Once you have such a robust maintenance mechanism, it is then definitionally devoid of any internal dynamics that could determine when new information is updated vs. existing information is maintained. Thus, they introduced an _input gate_ that determines when to update new information into the CEC memory, and an _output gate_ to determine when to read information out of the memory, to drive responding or other processes.

In equation form, the input gating to update the memory state is:

{id="eq_in" title="Input gating"}
$$
s_{c_j}(t) = s_{c_j}(t-1) + g(net_{c_j})(t) y^{in_j}(t)  
$$

where $s_{c_j}(t)$ is the CEC's activity state at timestep $t$; $g(net_{c_j}(t))$ is a nonlinear, squashing activation function producing values in the range 0..1, and $y^{in_j}(t)$ is the activation of the input gate function $in_j$.

The output from the memory cell is:

{id="eq_out" title="Output gating"}
$$
y^{c_j}(t) = y^{out_j}(t) h(s_{c_j})(t)
$$

where $y^{c_j}(t)$ is the memory cell's output at each timestep; $y^{out_j}(t)$ is the activity of the output gate unit $out_j$; and $h(s_{c_j}(t))$ is a nonlinear function of the CEC's current state value, $s_{c_j}$.

Other forms of gating were introduced in [[@^GersSchmidhuberCummins00]] and [[@^SchmidhuberGersEck02]] (clear and maintenance gates).

The [[PBWM]] (prefrontal-cortex, basal-ganglia working memory) model ([[@OReillyFrank06]]) shows how these gating mechanisms can arise from the disinhibitory (modulatory, multiplicative) influence of the [[basal ganglia]] over the excitatory loops between the [[thalamus]] and [[prefrontal cortex]].

