+++
Categories = ["Axon", "Learning"]
bibfile = "ccnlab.json"
+++

The **temporal derivative** is the central mechanism for [[error driven learning]] in [[axon]], via the [[kinase algorithm]]. It computes a _difference_ or _change_ (i.e., _derivative_) over _time_, instead of more standard differences computed between different state variables (e.g., top-down vs. bottom-up signals, conveyed by different anatomical pathways). A key advantage of a temporal derivative is that _time happens everywhere_ in a network, allowing an error signal to spread over time to all areas in the network of neurons in the brain. By contrast, derivatives computed between different anatomical pathways require these pathways to remain at least somewhat segregated and organized within the network, which typically would end up strongly constraining the kinds of error signals that can be computed.

Thus, a temporal derivative is a very robust, general-purpose mechanism of the sort that one might be particularly suited to the messy, organic world of biology.

Another appealing property of the temporal derivative is that it can be computed _locally_ at each neuron and synapse, through a _competition between two chemical processes with different rate constants_. Specifically, if you subtract a _slower_ process from a _faster_ one, then this automatically computes a temporal derivative. 

This is illustrated in the following simple simulation, which shows the response to a "driver" input that drives the fast and slow chemical processes. In the brain, this driver is neural activity in the form of pre and post-synaptic spiking, which is integrated by a series of chemical pathways driven mainly by the influx of _Calcium_ ions (see the [[kinase algorithm]] for details).

The driver input changes over time in a manner consistent with [[predictive learning]]: there is an initial _prediction_ value, and then a subsequent _outcome_ value. Think of the prediction as the local neural activity associated with the brain state present when generating a prediction of what will happen next, and the outcome as this local activity when experiencing the actual outcome, immediately after making the prediction.

The key result is the difference between the fast and slow traces _at the end of the time window_ when the sequence of prediction-then-outcome has completed. If this difference is positive, that reflects a positive-valued error gradient, and synaptic weights should correspondingly increase (known as **LTP** in the [[synaptic plasticity]] literature). Likewise, if it is negative, the synaptic weights should decrease (**LTD**).

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
    core.Bind(val, core.NewSlider(b)).SetMin(1).SetMax(mxVal).
        SetStep(1).SetEnforceStep(true).SetChangeOnSlide(true).OnChange(func(e events.Event) {
    	updt()
        tx.UpdateRender()
    })
}

addSlider(&predStr, &pred, 100)
addSlider(&outStr, &out, 100)
addSlider(&fastStr, &fastTau, 50)
addSlider(&slowStr, &slowTau, 50)
```

The code for this simulation updates the fast and slow variables according to a simple running-average update equation, e.g., for the $fast$ variable:

$$ fast += (1 / fastTau) * (driver - fast) $$

This equation causes the variable (fast or slow) to move toward the driver value at a rate determined by the "tau" factor. For example, if the driver is larger than fast, then $driver - fast$ is positive, so fast will increase to approach the value of the driver. If $fastTau = 10$, then it moves a 10th of the way toward the driver at each update. This very simple type of update equation is used throughout [[axon]] and is likewise very prevalent in biology.

Some things you can try:

* Increase or decrease both `Prediction` and `Outcome` by the same amount: observe that the weight change is only sensitive to the _relative_ differences, due to the competitive, subtraction logic between `Fast - Slow`.

* Set both `Prediction` and `Outcome` to the same value (e.g., 50), and observe that this results in _zero_ weight change, which is consistent with there being no error in the prediction relative to the outcome. This holds even when you significantly increase or decrease the raw values, e.g., both 20 or both 80. This is a critical point of contrast with [[hebbian]] forms of learning, which are typically driven by the overall levels of activity, such that you would expect (larger) weight increases with more activity.

* There are important constraints on the `Tau` factors too. For example, with `Prediction` and `Outcome` both at 50, increase `Slow Tau` up to 35. You can see that the weight change is positive now, even though there is no prediction error, just because the Slow factor is too slow to catch up at the end. This means that the local chemical rate constants that produce these `Tau` factors must be properly tuned for the actual temporal dynamics of the network-level error signals. Although this might be considered biologically implausible, in fact there is strong evidence of prominent [[oscillatory-rhythms]] in the brain at different characteristic frequencies, including the [[alpha cycle]] at roughly 10Hz and the [[theta cycle]] at roughly 5Hz. These rhythms have been shown to strongly influence learning, in a manner consistent with this simple model and the [[kinase algorithm]] more generally.

