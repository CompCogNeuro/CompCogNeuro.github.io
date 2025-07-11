+++
Categories = ["Computation"]
bibfile = "ccnlab.json"
+++

The **curse of dimensionality** ([[@Bellman57]]) arises in many aspects of neural computation, and provides an essential tool for understanding the relative strengths and weaknesses of computational algorithms. Human intuitions are generally (wildly) inaccurate about these issues, which effectively prevent any serial algorithm from succeeding in even a modestly complex problem space.

Thus, the main conclusion here is that only **parallel** algorithms that scale approximately linearly as the dimensionality increases are viable for real-world problems, as elaborated in the discussion of [[search]] as a unifying computational framework. This is why neural networks are much more useful than symbolic approaches to [[artificial intelligence]], for example. [[Reinforcement learning]], i.e., trial-and-error search, is severely impacted by the curse of dimensionality.

## Combinatorial explosion

The curse arises from the [combinatorial explosion](https://en.wikipedia.org/wiki/Combinatorial_explosion) (wikipedia link) that arises remarkably quickly as the dimensionality of a space increases. This combinatorial explosion is already present in the very lowest dimensions, e.g., in the number of pixels in a 2D monitor as you increase the pixel density. The impressive-sounding _megapixel_ increases touted by manufacturers actually correspond to relatively modest increases in the 1D linear density (e.g., 1,280 x 720 is almost 1 megapixels, but 1,920 x 1,080 is already 2 megapixels).

Moving up to the 3D volume level, the increases in mass associated with increases in linear body size of animals has major real-world consequences for limiting how big things can become, and how easy it is to fly (e.g., a bee scaled up simply can't fly).

In a simple binary space with two values per dimension, the total number of unique combinations is:

{id="eq_exp" title="Exponential explosion"}
$$
N = 2^d
$$

This innocent-looking expression can be deceptive, and people have a hard time understanding how fast it grows. A legendary story involves a chessboard maker requesting to be paid by simply doubling the number of grains of wheat on each square of a [chessboard](https://en.wikipedia.org/wiki/Wheat_and_chessboard_problem). This innocent-sounding request results in 18,446,744,073,709,551,615 grains total! 

Another fun real-world demonstration is the difficulty in folding paper in half progressively more than 8 times, because the thickness grows exponentially. The world record is apparently 13, which required [54,000 feet of toilet paper](https://www.npr.org/sections/thetwo-way/2011/12/05/143150449/record-folders-54-000-feet-of-paper-13-folds-one-new-standard) (NPR link).

In the domain of computation, the size of the problem space for even relatively simple games grows exponentially, limiting the ability to apply simple [[search]] through this problem space to find good strategies. For example, it is only possible to fully search the problem space in the game of chess with 7 pieces in play, and even this took a massive amount of computation and storage.

