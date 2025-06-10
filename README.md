# Spreading-pseudo-randomnis
a little thought of mine written in go:

Lets say we have a list of thing that can either be true or false and at the start all the things "are true" but have a certain chance of changing to false after some time or event has happened

as we already declared once a thing has become false, it cant become true again.
One way to approach this problem would be to make an array of bools and change them randomly after some time/event

another way would be to define a seed and threshold
setting the seed will make sure that we always get the same randomnis
for an array we should change the seed for each bool in the array i.e. by multiplying it by the index or using the index as seed, like i did here
(we need to make sure that the seed changes between the bools only)

we then use the seed to generate a float that is in the range of the threshold. In my example i used numbers between 0 and 1 for both the random value and threshold.
then we check what is larger, the threshold or the random value
