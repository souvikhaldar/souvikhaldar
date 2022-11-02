---
title: Asymptotic analysis of algorithms and Master Method
date: 2022-09-20
draft: false
---
## What is Asymptotic Analysis of algorithms?
Asymtotic analysis is the way of expressing how fast or slow an algorithm is, especially for large inputs. It removes the dependency to compare algorithms on the basis of its underlying programming language, speed of CPU, etc, hence it makes comparison between different algorithms easier.

## Notations of Asymptotic Analysis
Following are the notations for asymptotically analysing algorithms, where there are $n$ number of inputs and $T(n)$ signifies the time required for processing $n$ number of inputs.  
1. O(n) or Big-Oh notation: It signifies the running time of an algorithm in the worst case. It can also be thought as the upper bound on the time taken, that at max it will take this amount of time. 
$T(n)\space is\space O(f(n))$ if and only if there exists $c$ and $n_0$ such that
$$ T(n) \leq c \times f(n) \space \space \forall  n>n_0, n_0 \geq 0$$
![](/images/bigOh.png)
2. $\Omega(n)$ or Omega notation: It signifies the running time of an algorithm is the best case. 
$T(n)=\Omega(f(n))$ if and only if there exists $c$ and $n_0$ such that $$T(n) \geq c \times f(n) \space \space \forall n \geq n_0, n_0 \geq 0$$
3. $\theta(n)$ or Theta notation: It signifies the average running time of an algoritm and if often used interchangably with $O(n)$ 
   $$T(n) = \theta(n) \space if\space and\space only\space if\space T(n) = O(n)\space and\space T(n) = \Omega(n)$$
## Calculating Asymptotic running time of algorithms
If the algorithm can be implemented in the recursive fashion and each of the sub-problem is of equal size, we can use **Master Method** for calculating the running time of the algorithm in a very simple way.
The running time of the algorithm is expressed as: $$T(n)\space \leq \space aT(n/b)+ O(n^d)$$
where,   
a = Number of recursive calls   
b = The factor by which the input size decreases in each recursive call. Eg. for merge sort it would be 2 since we divide the input array into two halves in every step   
d = exponent in the running time of the work done outside the recusive call, also called "combine step”  

then,
the upper bound in the running time i.e O(n) is: 
$T(n) = O(n^d log_n)$  if $a=b^d$  
$T(n) = O(n^d)$  if $a \lt b^d$   
$T(n) = O(n^{\log_b{a}})$ if $a \gt b^d$





## References
1.  Divide and Conquer, Sorting and Searching, and Randomized Algorithms, Stanford University

