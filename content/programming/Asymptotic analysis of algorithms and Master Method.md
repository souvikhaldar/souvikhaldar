---
title: Asymptotic analysis of algorithms and Master Method
date: 2022-09-20
draft: false
---
Asymtotic analysis is the way of expressing how fast or slow an algorithm is, especially for large inputs. It removes the dependency to compare algorithms from its underlying programming language, speed of CPU, etc, hence it makes comparison between different algorithms easier.
Following are the notations for asumtotically analysing algorithms:
1. O(n) or Big-Oh notation: It signifies the running time of an algorithm in the worst case. 
$T(n) = O(n)$ if and only if there exists $c$ and $n_0$ such that
$$ T(n) \leq c \times O(n)  \forall  n>n_0$$

2. $\Omega(n)$ or Omega notation: It signifies the running time of an algorithm is the best case. 
$T(n)=\Omega(n)$ if and only if there exists $c$ and $n_0$ such that $$T(n) \geq c \times f(n) \forall n \geq n_0$$


Master method is a set of formulae for calculaing the upper bound on the running time of a problem, which can be expressed in terms of recursive calls.
It assumes all sub-problems are of equal size, like for merge sort algorithm, each sub-problem is half of the problem, since in each recursion the list is broken into two halves.


So the Big O of T(n) can be calculated as follows:

## References
1.  Divide and Conquer, Sorting and Searching, and Randomized Algorithms, Stanford University
