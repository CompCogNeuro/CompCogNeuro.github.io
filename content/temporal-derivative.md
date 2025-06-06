+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.json"
+++

The **temporal derivative** is the _change_ over _time_ in the activity states of neurons, and it is the central mechanism for [[error-driven learning]] in [[Axon]], via the [[kinase algorithm]], building on the mathematical framework of [[GeneRec]]. In the [[predictive learning]] context, time can be discretized into two distinct phases (as originally developed in the [[Boltzmann machine]]):

{id="figure_minus-plus"  style="height:20em"}
![Temporal derivative learning is based on the change over time in neural activity states, e.g., as discreteized into an initial Minus phase where the network is making a prediction, and a subsequent Plus phase where the actual outcome is experienced. Synapses everywhere can learn on this temporal difference, without requiring a separate structural pathway to convey the error, and error signals can arise from anywhere in a bidirectionally-connected network, which are key advantages of using these differences over time.](media/fig_minus_plus_phase_err.png)

* The **Minus phase** is when the network is generating a **prediction**.
* The **Plus phase** is when the network is experiencing the actual **outcome** (what actually happened).

The temporal derivative is the _difference_ between these two phases, i.e., the **prediction error**, and every neuron in the network learns as a function of this difference:

* Learning $\propto$ (Plus - Minus)

In simple equation form, the change in synaptic weight $w$ is a function of the phase-wise difference in activity state $y$ between the plus phase ($y^+$) and minus phase ($y^-$):

{id="eq_td" title="Temporal difference learning"}
$$
\Delta w \propto y^+ - y^-
$$

By contrast, in standard [[error backpropagation]] learning, error signals are propagated via _sturcturally separate_ pathways, using distinct equations from those governing activation propagation, which is the primary source of biological implausibility.

A key advantage of a temporal derivative is that _time happens everywhere_ in a network, allowing an error signal to spread over time to all areas in the network of neurons in the brain. By contrast, derivatives computed between different anatomical pathways require these pathways to remain at least somewhat segregated and organized within the network, which typically would end up strongly constraining the kinds of error signals that can be computed.

Thus, a temporal derivative is a very robust, general-purpose mechanism of the sort that one might be particularly suited to the messy, organic world of biology. Indeed, initial empirical support for this mechanism is reported in [[Jiang et al 2025]]. Outside of the [[neocortex]], the [[TD]] (temporal differences) algorithm for [[reinforcement learning]] shares the same temporal prediction error framework but maps onto very different neural substrates, in the form of [[dopamine]].

## Local computation of the temporal derivative

Another appealing property of the temporal derivative is that it can be computed _locally_ at each neuron and synapse, through a _competition between two chemical processes with different [[time constant]]s_. Specifically, if you subtract a _slower_ process from a _faster_ one, then this automatically computes a temporal derivative. 

This is illustrated in the following simple simulation, which shows the response to a "driver" input that drives the fast and slow chemical processes (see [[time constant]] for a detailed explanation of the exponential updating used here). In the brain, this driver is neural activity in the form of pre and post-synaptic spiking, which is integrated by a series of chemical pathways driven mainly by the influx of _Calcium_ ions (see [[synaptic plasticity]] and the [[kinase algorithm]] for details).

The driver input changes over time in a manner consistent with [[predictive learning]]: there is an initial _prediction_ value, and then a subsequent _outcome_ value ([[#figure_minus-plus]]). The prediction is the local neural activity associated with the brain state present when generating a prediction of what will happen next, and the outcome is this local activity when experiencing the actual outcome, immediately after making the prediction.

The prediction error is represented by the difference between the fast and slow traces _at the end of the time window_ when the sequence of prediction-then-outcome has completed. If this difference is positive, that reflects a positive-valued error gradient, and synaptic weights should correspondingly increase (known as **LTP** in the [[synaptic plasticity]] literature). Likewise, if it is negative, the synaptic weights should decrease (**LTD**).

{id="sim_td" title="Temporal Derivative from Fast - Slow" collapsed="true"}
```Goal
fastTau := 10.0 // time constant for fast integration
slowTau := 20.0 // time constant for slow integration
pred := 50.0
out := 80.0
var diffStr, fastStr, slowStr, predStr, outStr string

##
totalTime := 100
driver := zeros(totalTime) // driver is what is driving the system
fast := zeros(totalTime) // fast is a fast integrator of driver
slow := zeros(totalTime) // slow is a slow integrator of driver
##

func td() {
    fastStr = fmt.Sprintf("Fast Tau: %g", fastTau)
    slowStr = fmt.Sprintf("Slow Tau: %g", slowTau)
    predStr = fmt.Sprintf("Prediction: %g", pred)
    outStr = fmt.Sprintf("Outcome: %g", out)
    ##
    d := array(pred) // current drive
    f := 0.0 // current fast
    s := 0.0 // current slow
    fTau := array(fastTau)
    sTau := array(slowTau)
    ##
    for t := range 100 {
        if t == 75 {
            # d = array(out)
        }
        ##
        f += (1.0 / fTau) * (d - f) // f moves toward d
        s += (1.0 / sTau) * (d - s) // s moves toward f
        driver[t] = d
        fast[t] = f
        slow[t] = s
        ##
    }
    ##
    diff := fast[-1] - slow[-1]
    ##
    diffStr = fmt.Sprintf("<b>Weight Change ΔW ≅ Predition - Outcome = Fast - Slow = %7.2g</b>", diff.Float1D(0))
}

td()

plotStyler := func(s *plot.Style) {
    s.Range.SetMax(100).SetMin(0)
    s.Plot.XAxis.Label = "Time"
    s.Plot.XAxis.Range.SetMax(100).SetMin(0)
	s.Plot.Legend.Position.Left = true
}
plot.SetStyler(driver, plotStyler) 

fig1, pw := lab.NewPlotWidget(b)
dl := plots.NewLine(fig1, driver)
fl := plots.NewLine(fig1, fast)
sl := plots.NewLine(fig1, slow)
fig1.Legend.Add("Driver", dl)
fig1.Legend.Add("Fast", fl)
fig1.Legend.Add("Slow", sl)


diffTx := core.NewText(b)
diffTx.Styler(func(s *styles.Style) {
    s.Min.X.Ch(80) // clean rendering with variable width content
})
core.Bind(&diffStr, diffTx)

func updt() {
    td()
    dl.SetData(driver)
    fl.SetData(fast)
    sl.SetData(slow)
    diffTx.UpdateRender()
    pw.NeedsRender()
}

func addSlider(label *string, val *float64, mxVal float32) {
    tx := core.NewText(b)
    tx.Styler(func(s *styles.Style) {
        s.Min.X.Ch(40)  // clean rendering with variable width content
    })
	core.Bind(label, tx)
	sld := core.NewSlider(b).SetMin(1).SetMax(mxVal).SetStep(1).SetEnforceStep(true)
	sld.SendChangeOnInput()
	sld.OnChange(func(e events.Event) {
		updt()
		tx.UpdateRender()
	})
	core.Bind(val, sld)
}

addSlider(&predStr, &pred, 100)
addSlider(&outStr, &out, 100)
addSlider(&fastStr, &fastTau, 50)
addSlider(&slowStr, &slowTau, 50)
```

The code for this simulation updates the fast and slow variables according to a simple running-average update equation, e.g., for the $fast$ variable: 

{id="eq_fast-slow"}
$$
\rm{fast} += \frac{1}{\rm{fastTau}} (\rm{driver} - \rm{fast})
$$

This equation causes the variable (fast or slow) to move toward the driver value at a rate determined by the "tau" [[time constant]] factor. For example, if the driver is larger than fast, then $driver - fast$ is positive, so fast will increase to approach the value of the driver. If $fastTau = 10$, then it moves a 10th of the way toward the driver at each update (and gets roughly 2/3 of the way to the driver in _fastTau_ steps). This very simple type of update equation is used throughout [[axon]] and is likewise very prevalent in biology.

Some things you can try:

* Increase or decrease both `Prediction` and `Outcome` by the same amount: observe that the weight change is only sensitive to the _relative_ differences, due to the competitive, subtraction logic between `Fast - Slow`.

* Set both `Prediction` and `Outcome` to the same value (e.g., 50), and observe that this results in _zero_ weight change, which is consistent with there being no error in the prediction relative to the outcome. This holds even when you significantly increase or decrease the raw values, e.g., both 20 or both 80. This is a critical point of contrast with [[hebbian]] forms of learning, which are typically driven by the overall levels of activity, such that you would expect (larger) weight increases with more activity.

* There are important constraints on the `Tau` factors too. For example, with `Prediction` and `Outcome` both at 50, increase `Slow Tau` up to 35. You can see that the weight change is positive now, even though there is no prediction error, just because the Slow factor is too slow to catch up at the end. This means that the local chemical rate constants that produce these `Tau` factors must be properly tuned for the actual temporal dynamics of the network-level error signals. Although this might be considered biologically implausible, in fact there is strong evidence of prominent [[oscillatory-rhythms]] in the brain at different characteristic frequencies, including the [[alpha cycle]] at roughly 10Hz and the [[theta cycle]] at roughly 5Hz. These rhythms have been shown to strongly influence learning, in a manner consistent with this simple model and the [[kinase algorithm]] more generally.

In summary, [[#sim_td]] based on the competition between two simple exponential integration equations ([[#eq_fast-slow]]) demonstrates that a locally computed temporal derivative can drive synaptic changes in a manner consistent with an error signal that emerges over time.

## When is the temporal derivative computed?

A critical issue with this temporal derivative framework is that the accurate computation of a prediction error signal must happen at a specific point it time relative to the onset of the actual outcome, which you can see in the above example in terms of the effects of the different time constants. The precise timing of the prediction signals is less critical, because any neural activity that precedes the outcome can be considered a prediction, and the cumulative effects of the learning will cause these prior activity states to become a prediction in any case.

The [[kinase algorithm]] provides an answer to this key question (TODO: summary here!).

