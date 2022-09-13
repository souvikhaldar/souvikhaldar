---
title: Master Theorem for asymptotic analysis of algorithms
date: 2022-09-12
draft: false
---
Master method is a set of formulae for calculaing the upper bound on the running time of a problem  which can be expressed in terms of recursive calls.
It assumes all sub-problems are of equal size, like for merge sort algorithm, each sub-problem is half of the problem, since in each recursion the list is broken into two halves.  

![](/images/master1.png)  


So the Big O of T(n) can be calculated as follows:  

![](/images/master2.png)  

## References
1.  Divide and Conquer, Sorting and Searching, and Randomized Algorithms, Stanford University
