---
title: Asymptotic analysis of algorithms and Master Method
date: 2022-09-20
draft: false
---
Asymtotic analysis is the way of expressing how fast or slow an algorithm is, especially for large inputs. It removes the dependency to compare algorithms on the basis of its underlying programming language, speed of CPU, etc, hence it makes comparison between different algorithms easier.
Following are the notations for asymptotically analysing algorithms, where there are $n$ number of inputs and $T(n)$ signifies the time required for processing $n$ number of inputs.  
1. O(n) or Big-Oh notation: It signifies the running time of an algorithm in the worst case. It can also be thought as the upper bound on the time taken, that at max it will take this amount of time. 
$T(n) = O(f(n))$ if and only if there exists $c$ and $n_0$ such that
$$ T(n) \leq c \times f(n) \space \space \forall  n>n_0, n_0 \geq 0$$

2. $\Omega(n)$ or Omega notation: It signifies the running time of an algorithm is the best case. 
$T(n)=\Omega(f(n))$ if and only if there exists $c$ and $n_0$ such that $$T(n) \geq c \times f(n) \space \space \forall n \geq n_0, n_0 \geq 0$$
3. $\theta(n)$ or Theta notation: It signifies the average running time of an algoritm and if often used interchangably with $O(n)$ 
   $T(n) = \theta(n)$ if and only if $T(n) = O(n)$ and $T(n) = \Omega(n)$ 



## References
1.  Divide and Conquer, Sorting and Searching, and Randomized Algorithms, Stanford University
