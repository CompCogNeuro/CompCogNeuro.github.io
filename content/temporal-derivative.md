+++
Categories = ["Axon"]
+++

The **temporal derivative** is the central mechanism for [[error-driven-learning]] in [[axon]]. It computes a _difference_ or _change_ (i.e., _derivative_) over _time_, instead of more standard differences computed between different state variables (e.g., top-down vs. bottom-up signals, conveyed by different anatomical pathways). A key advantage of a temporal derivative is that _time happens everywhere_ in a network, allowing an error signal to spread over time to all areas in the network of neurons in the brain. By contrast, derivatives computed between different anatomical pathways require these pathways to remain at least somewhat segregated and organized within the network, which typically would end up strongly constraining the kinds of error signals that can be computed.

Thus, a temporal derivative is a very robust, general-purpose mechanism of the sort that one might be particularly suited to the messy, organic world of biology.

Another appealing property of the temporal derivative is that it can be computed _locally_ at each neuron and synapse, through a _competition between two chemical processes with different rate constants_. Specifically, if you subtract a _slower_ process from a _faster_ one, then this automatically computes a temporal derivative. This is illustrated in the following simple simulation:

```Goal
fastTau := 10.0 // time constant for fast integration
slowTau := 10.0 // for slow, which compounds on top of fast
drive1 := 50.0
drive2 := 80.0
var fastStr, slowStr, drive1Str, drive2Str string
##
totalTime := 100
driver := zeros(totalTime) // driver is what is driving the system
fast := zeros(totalTime) // fast is a fast integrator of driver
slow := zeros(totalTime) // slow is a slow integrator of driver
##
func td() {
    fastStr = fmt.Sprintf("Fast Tau: %g", fastTau)
    slowStr = fmt.Sprintf("Slow Tau: %g", slowTau)
    drive1Str = fmt.Sprintf("Drive 1: %g", drive1)
    drive2Str = fmt.Sprintf("Drive 2: %g", drive2)
    ##
    d := tensor.NewFloat64Scalar(drive1) // current drive
    f := 0.0 // current fast
    s := 0.0 // current slow
    fTau := tensor.NewFloat64Scalar(fastTau)
    sTau := tensor.NewFloat64Scalar(slowTau)
    ##
    for t := range 100 {
        tt := tensor.NewFloat64Scalar(float64(t))
        if t == 75 {
            # d = tensor.NewFloat64Scalar(drive2)
        }
        ##
        f += (1.0 / fTau) * (d - f) // f moves toward d
        s += (1.0 / sTau) * (f - s) // s moves toward f
        driver[t] = d
        fast[t] = f
        slow[t] = s
        ##
    }
}

td()

styMax := func(s *plot.Style) {
    s.Range.SetMax(100)
}
plot.SetStyler(driver, styMax) 

fig1 := lab.NewPlot(b)
dl := plots.NewLine(fig1, driver)
fl := plots.NewLine(fig1, fast)
sl := plots.NewLine(fig1, slow)
fig1.Legend.Add("Driver", dl)
fig1.Legend.Add("Fast", fl)
fig1.Legend.Add("Slow", sl)

func updt() {
    td()
    dl.SetData(driver)
    fl.SetData(fast)
    sl.SetData(slow)
    b.Update()
}

core.Bind(&drive1Str, core.NewText(b))
core.Bind(&drive1, core.NewSlider(b)).SetMin(1).SetMax(100).
    SetStep(1).SetEnforceStep(true).OnChange(func(e events.Event) {
	updt()
})
core.Bind(&drive2Str, core.NewText(b))
core.Bind(&drive2, core.NewSlider(b)).SetMin(1).SetMax(100).SetStep(1).SetEnforceStep(true).OnChange(func(e events.Event) {
	updt()
})
core.Bind(&fastStr, core.NewText(b))
core.Bind(&fastTau, core.NewSlider(b)).SetMin(1).SetMax(100).SetStep(1).SetEnforceStep(true).OnChange(func(e events.Event) {
	updt()
})
core.Bind(&slowStr, core.NewText(b))
core.Bind(&slowTau, core.NewSlider(b)).SetMin(1).SetMax(100).SetStep(1).SetEnforceStep(true).OnChange(func(e events.Event) {
	updt()
})
```

