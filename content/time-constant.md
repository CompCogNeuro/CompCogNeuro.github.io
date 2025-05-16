+++
Categories = ["Axon"]
bibfile = "ccnlab.json"
+++

{id="sim_tau" title="Tau exponential integration time constant" collapsed="true"}
```Goal
tau := 10.0 // time constant
drive := 50.0
init := 0.0
val := init
var tauStr, driveStr, initStr, valAtTauStr string

##
totalTime := 100
driver := zeros(totalTime) // driver is what is driving the value
vals := zeros(totalTime) // values
exps := zeros(totalTime) // exponential function values
##

func valUpdate() {
    tauStr = fmt.Sprintf("Tau: %g", tau)
    driveStr = fmt.Sprintf("Driver: %g", drive)
    initStr = fmt.Sprintf("Init: %g", init)
    ##
    d := array(drive) // current drive
    iv := array(init)
    dv := d - iv
    v := array(init) // current
    ttau := array(tau)
    ##
    for t := range 100 {
        ##
        tv := array(t)
        driver[t] = d
        vals[t] = v
        v += (1.0 / ttau) * (d - v) // v moves toward d
        exps[t] = iv + dv * (1 - exp(-tv / ttau))
        ##
    }
    tint := int(tau)
    tval := vals.Float(tint)
    dist := drive - init
    pdist := 100.0
    if dist != 0 {
        pdist = 100 * (tval-init) / dist
    }
    valAtTauStr = fmt.Sprintf("<b>Value at Tau (%d) = %7.3g = %7.3g %% of way to Drive, %7.3g %% left to go</b>", tint, tval, pdist, 100 - pdist)
}

valUpdate()

plotStyler := func(s *plot.Style) {
    s.Range.SetMax(100).SetMin(0)
    s.Plot.XAxis.Label = "Time"
    s.Plot.XAxis.Range.SetMax(100).SetMin(0)
	s.Plot.Legend.Position.Left = true
}
plot.SetStyler(driver, plotStyler) 

fig1, pw := lab.NewPlotWidget(b)
dl := plots.NewLine(fig1, driver)
vl := plots.NewLine(fig1, vals)
el := plots.NewLine(fig1, exps)
fig1.Legend.Add("Driver", dl)
fig1.Legend.Add("Value", vl)
fig1.Legend.Add("Exp", el)

valAtTauTx := core.NewText(b)
valAtTauTx.Styler(func(s *styles.Style) {
    s.Min.X.Ch(80) // clean rendering with variable width content
})
core.Bind(&valAtTauStr, valAtTauTx)

func updt() {
    valUpdate()
    dl.SetData(driver)
    vl.SetData(vals)
    el.SetData(exps)
    valAtTauTx.UpdateRender()
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

addSlider(&tauStr, &tau, 50)
addSlider(&initStr, &init, 100)
addSlider(&driveStr, &drive, 100)
```

[[#sim_tau]] provides an illustration of the process of **exponential updating** governed by a **time constant** `tau` ($\tau$), where a given value is updated at each time step to move some fraction of the way toward a "driving" value, where this fraction is $\frac{1}{\tau}$:

{id="eq_tau" title="Exponential updating"}
$$
v(t+1) = v(t) + \frac{1}{\tau} (d - v(t))
$$

This type of equation is used almost everywhere in [[Axon]], and is a basic consequence of many different types of physical processes, including Ohm's law for the influence of current on electrical potential, which is the core equation driving our model of the [[neuron]].

Chemical reactions also obey the same kind of function, governed by a tau factor that represents the speed of the chemical reaction. For example, two initial "ingredients" _A_, _B_ can be combined into a _compound_ _AB_ by a chemical reaction with a _forward_ rate constant of $K_f$:

```
      Kf
A + B --> AB
```
As the reaction proceeds, there is less of the raw ingredients available, so it naturally slows down as a function of the amount of these ingredients available, resulting in the same dynamic as [[#eq_tau]].

## Value at Tau

As you can see in [[#eq_tau]] as you experiment with different `Tau`, `Init`, and `Driver` values, it is always about 64% of the way to the driver after Tau time steps of updating. This can be computed directly from the exponential function that can compute the value directly at any given time step (which is plotted as `Exp` in [[#sim_tau]]):

{id="eq_exp" title="Exponential function"}
$$
v(t) = v(0) + (d - v(0)) \left(1 - e^{-t/\tau} \right)
$$

where $v(0)$ is the `Init` initial value at time step 0. If we set this to 0, and evaluate at $t=\tau$, this becomes:

{id="eq_exp-tau" title="Value at t = tau"}
$$
v(\tau) = d \left(1 - e^{-1} \right) \approx 1 - 1/2.72 \approx 0.632..
$$

The small discrepancy in reported values is due to the numerical differences from iterating the update equation ([[#eq_tau]]) at discrete time intervals, versus the fully continuous update computed by the exponential function (due to calculus and all that).

In the simplest terms, you can think of $\tau$ as the amount of time it takes to get roughly 2/3 of the way to the driving value.

## Half-life

Another value you have probably heard of is the _half-life_ of an exponential process (radioactive decay also follows this same function), which is when the value is 50% of the way to the driving value. This happens when $t = ln(2) \tau$ or roughly $t = 0.7 \tau$, as you can see by setting $v(t) = \frac{1}{2}d$ and solving for _t_:

{id="eq_hl-tau" title="Half-life time"}
$$
\frac{1}{2} d = d \left(1 - e^{-t/\tau} \right)
$$

$$
\left(1 - e^{-t/\tau} \right) = \frac{1}{2}
$$

take the natural log of both sides:

$$
\frac{-t}{\tau} = ln(1) - ln(2)
$$

$$
t = ln(2) \tau
$$

See this [Wikipedia page](https://en.wikipedia.org/wiki/Time_constant) for more information.
