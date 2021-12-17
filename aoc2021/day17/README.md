# part 1, calc explained
For a high shot, you're going to be aiming up. Then you need to notice that for any upwards shot, the trajectory will always pass through y = 0 later on with opposite vertical speed.

To get the max height you need to maximise your initial vertical direction, which means also maximising your downwards vertical speed when you come back through y = 0.

So ... if your initial vertical speed is any greater than the distance to the lower y bound of the target box, the next step downwards after y = 0 will overshoot the target and you'll never get it back.

This assumes you can find a horizontal speed that will stop in the target zone (i.e. you're no longer travelling horizontally when you pass through the target) but I believe the problem has always been formulated thus.

Now to get max height you do the sum of the arithmetic series starting with your max initial vertical velocity. Simples

Source: https://www.reddit.com/r/adventofcode/comments/ria9mm/comment/howd1wd/?utm_source=share&utm_medium=web2x&context=3
