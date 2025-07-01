+++
Categories = ["Activation", "Neuroscience"]
bibfile = "ccnlab.json"
+++

Before finishing up the final step in the detection process (generating an output), we will need to use the concept of the **equilibrium membrane potential**, which is the value of $V_m$ that the neuron will settle into and stay at, *given a fixed set of excitatory and inhibitory input conductances* (if these aren't steady, then the $V_m$ will likely be constantly changing as they change). This equilibrium value is interesting because it tells us more clearly how the tug-of-war process inside the neuron actually balances out in the end. Also, we will see in the next section that it is useful mathematically.

To compute the equilibrium membrane potential ($V_m^{eq}$), we can use an important mathematical technique: set the change in membrane potential (according to the iterative $V_m$ updating equation from above) to 0, and then solve the equation for the value of $V_m$ under this condition. In other words, if we want to find out what the equilibrium state is, we simply compute what the numbers need to be such that $V_m$ is no longer changing (i.e., its rate of change is 0). Here are the mathematical steps that do this:

{id="eq_vm_updt" title="Vm update"}
$$
V_m(t) = V_m(t-1) + dt_{vm} \left[ g_e (E_e-V_m) + g_i (E_i-V_m) + g_l (E_l-V_m) \right]
$$

This is the part that is driving the changes (time constant omitted as we are looking for equilibrium):

{id="eq_delta_vm" title="Delta Vm"}
$$
\Delta V_m = g_e \left(E_e-V_m\right) + g_i (E_i-V_m) + g_l (E_l-V_m)
$$

which we set to zero to find when it stops changing:

$$
0 = g_e \left(E_e-V_m\right) + g_i (E_i-V_m) + g_l (E_l-V_m)
$$

and then do some algebra to solve for $V_m$:

{id="eq_vm_eq" title="Equilibrium Vm"}
$$
V_m = \frac{g_e}{g_e + g_i + g_l} E_e + \frac{g_i}{g_e + g_i + g_l} E_i + \frac{g_l}{g_e + g_i + g_l} E_l
$$

The detailed math is shown below in [[#Derivation]].

In words, this says that the excitatory drive $E_e$ contributes to the overall $V_m$ as a function of the proportion of the excitatory conductance $g_e$ relative to the sum of all the conductances ($g_e + g_i + g_l$). And the same for each of the others (inhibition, leak). This is just what we expect from the tug-of-war picture: if we ignore $g_l$, then the $V_m$ "flag" is positioned as a function of the relative balance between $g_e$ and $g_i$ --- if they are equal, then $g_e / (g_e + g_i)$ is .5 (e.g., just put a "1" in for each of the g's --- 1/2 = .5), which means that the $V_m$ flag is half-way between $E_i$ and $E_e$.

So, all this math just to rediscover what we knew already intuitively! (Actually, that is the best way to do math --- if you draw the right picture, it should tell you the answers before you do all the algebra). But we'll see that this math will come in handy next.

Here is a version with the conductance terms explicitly broken out into the "g-bar" constants and the time-varying "g(t)" parts:

{id="eq_vm_eq" title="Equilibrium Vm full"}
$$
V_m = \frac{\overline{g}_e g_e(t)}{\overline{g}_e g_e(t) + \overline{g}_i g_i(t) + \overline{g}_l} E_e  \frac{\overline{g}_i g_i(t)}{\overline{g}_e g_e(t) + \overline{g}_i g_i(t) + \overline{g}_l} E_i + \frac{\overline{g}_l}{\overline{g}_e g_e(t) + \overline{g}_i g_i(t) + \overline{g}_l} E_l
$$

For those who really like math, the equilibrium membrane potential equation can be shown to be a [[neuron bayesian|Bayesian optimal detector]].

## Derivation

This shows all the algebra to derive the equilibrium membrane potential from the update equation --- it will only be viewable on the PDF version.

iterative Vm update equation:

$$
V_m(t) = V_m(t-1) + dt_{vm} \left[ g_e (E_e-V_m) + g_i (E_i-V_m) + g_l (E_l-V_m) \right] $$

just the change part:

$$
\Delta Vm = g_e \left(E_e-V_m\right) + g_i (E_i-V_m) + g_l (E_l-V_m)
$$

set it to zero:

$$
0 = g_e \left(E_e-V_m\right) + g_i (E_i-V_m) + g_l (E_l-V_m)
$$

solve for Vm: (multiply all the g's through)

$$
0 = g_e E_e - g_e V_m + g_i E_i - g_i V_m + g_l E_l - g_l V_m
$$

solve for Vm: (gather Vm terms on other side)

$$
g_e V_m + g_i V_m + g_l V_m = g_e E_e + g_i E_i + g_l E_l
$$

solve for Vm: (get only one Vm guy, then divide each side by g's to get..)

$$
V_m (g_e + g_i + g_l ) = g_e E_e + g_i E_i + g_l E_l
$$

solution!

$$
V_m = \frac{g_e E_e + g_i E_i + g_l E_l}{g_e + g_i + g_l}
$$

Another way of writing this solution, which makes its meaning a bit clearer, is:

$$
V_m = \frac{g_e}{g_e + g_i + g_l} E_e + \frac{g_i}{g_e + g_i + g_l} E_i + \frac{g_l}{g_e + g_i + g_l} E_l
$$

In the Adaptive Exponential function, there is an adaptive factor $\omega$ (greek omega) that enters into the membrane update equation, which we can include in our equilibrium calculation:

iterative Vm update equation:

$$
V_m(t) = V_m(t-1) + dt_{vm} \left[ g_e (E_e-V_m) + g_i (E_i-V_m) + g_l (E_l-V_m) - \omega \right]
$$

just the change part:

$$
\Delta Vm = g_e \left(E_e-V_m\right) + g_i (E_i-V_m) + g_l (E_l-V_m) - \omega
$$

set it to zero:

$$
0 = g_e \left(E_e-V_m\right) + g_i (E_i-V_m) + g_l (E_l-V_m) - \omega
$$

solve for Vm: (multiply all the g's through)

$$
0 = g_e E_e - g_e V_m + g_i E_i - g_i V_m + g_l E_l - g_l V_m - \omega
$$

solve for Vm: (gather Vm terms on other side)

$$
g_e V_m + g_i V_m + g_l V_m = g_e E_e + g_i E_i + g_l E_l - \omega
$$

solve for Vm: (get only one Vm guy, then divide each side by g's to get..)

$$
V_m (g_e + g_i + g_l ) = g_e E_e + g_i E_i + g_l E_l - \omega
$$

solution:

$$
V_m = \frac{g_e E_e + g_i E_i + g_l E_l - \omega}{g_e + g_i + g_l}
$$

And here is the derivation of the equation for $g_e^{\Theta}$:

equilibrium Vm at threshold:

$$
\Theta = \frac{g_e^{\Theta} E_e + g_i E_i + g_l E_l}{g_e^{\Theta} + g_i + g_l}
$$

solve for g_e:

$$
\Theta (g_e^{\Theta} + g_i + g_l) = g_e^{\Theta} E_e + g_i E_i + g_l E_l
$$

(multiply both sides by g's), then solve for g_e:

$$ 
\Theta g_e^{\Theta} = g_e^{\Theta} E_e + g_i E_i + g_l E_l  - \Theta g_i - \Theta g_l
$$

(bring non-g_e's back to other side), then solve for g_e:

$$
g_e^{\Theta} \Theta - g_e^{\Theta} E_e  = g_i (E_i - \Theta) + g_l (E_l  - \Theta)
$$

(bring other g_e over and consolidate other side, then divide both sides to isolate g_e to get), solution:

$$
g_e^{\Theta} = \frac{g_i (E_i - \Theta) + g_l (E_l  - \Theta)}{\Theta - E_e}
$$

In the AdEx function, there is an adaptive factor $\omega$ (greek omega) that enters into the membrane update equation, which we can include in our equilibrium calculation as above:

equilibrium Vm at threshold:

$$
\Theta = \frac{g_e^{\Theta} E_e + g_i E_i + g_l E_l - \omega}{g_e^{\Theta} + g_i + g_l}
$$

solve for g_e:

$$
\Theta (g_e^{\Theta} + g_i + g_l) = g_e^{\Theta} E_e + g_i E_i + g_l E_l - \omega
$$

(multiply both sides by g's), then solve for g_e:

$$
\Theta g_e^{\Theta} = g_e^{\Theta} E_e + g_i E_i + g_l E_l  - \Theta g_i - \Theta g_l - \omega$$

(bring non-g_e's back to other side), then solve for g_e:

$$
g_e^{\Theta} \Theta - g_e^{\Theta} E_e  = g_i (E_i - \Theta) + g_l (E_l  - \Theta) - \omega
$$

(bring other g_e over and consolidate other side, then divide both sides to isolate g_e to get), solution:

$$
g_e^{\Theta} = \frac{g_i (E_i - \Theta) + g_l (E_l  - \Theta) - \omega}{\Theta - E_e}
$$

