+++
Categories = ["Axon"]
+++

The **temporal derivative** is the central mechanism for [[error-driven-learning]] in [[axon]]. It computes a _difference_ or _change_ (i.e., _derivative_) over _time_, instead of more standard differences computed between different state variables (e.g., top-down vs. bottom-up signals, conveyed by different anatomical pathways). A key advantage of a temporal derivative is that _time happens everywhere_ in a network, allowing an error signal to spread over time to all areas in the network of neurons in the brain. By contrast, derivatives computed between different anatomical pathways require these pathways to remain at least somewhat segregated and organized within the network, which typically would end up strongly constraining the kinds of error signals that can be computed.

Thus, a temporal derivative is a very robust, general-purpose mechanism of the sort that one might be particularly suited to the messy, organic world of biology.

Another appealing property of the temporal derivative is that it can be computed _locally_ at each neuron and synapse, through a _competition between two chemical processes with different rate constants_. Specifically, if you subtract a _slower_ process from a _faster_ one, then this automatically computes a temporal derivative. This is illustrated in the following simple simulation:

```Goal
##
totalTime := 100
driver := zeros(totalTime) // driver is what is driving the system
fast := zeros(totalTime) // fast is a fast integrator of driver
slow := zeros(totalTime) // slow is a slow integrator of driver
d := 0 // current drive
f := 0 // current fast
s := 0 // current slow
fTau := 40 // time constant for fast integration: 40 msec
sTau := 40 // for slow, which compounds on top of fast
##

for i, t := range time.Values {
    ##
    d = t / totalTime
    f += (1 / fTau) * (d - f) // f moves toward d
    s += (1 / sTau) * (f - s) // s moves toward f
    driver[i] = d
    fast[i] = f
    slow[i] = s
}
fig1 := lab.NewPlot(b)
plots.NewLine(fig1, driver)
plots.NewLine(fig1, fast)
plots.NewLine(fig1, slow)
```

