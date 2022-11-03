---
title: Golang Garbage Collector Comes With A New Feature
date: 2022-11-03
draft: false
---
## Introduction
The Garbage Collector (GC) is responsible for clearing up memory in the heap section of the main memory which becomes inaccessible, like the pointer which was pointing to that memory in the heap is now set to `nil`.
This is great! We don't need to manually `free` up the heap like we are required to in C programming language.
But every benefit comes with its own shortcomings.
The GC takes up the CPU time that could have otherwise been used to run the application. It can be really expensive on the CPU at times!

To deal with this scenario, Golang GC uses a mechanism to limit the GC, primarily by using an environment variable called `GOGC`. It is a percentage of **New Heap** to **Live Heap**, till which GC should not run. Let me explain it with an example, if the **New Heap** (i.e. the memory accumulated in the heap section that is not currently being used in the running application) is `GOGC` % of the **Live Heap** (i.e. the amount of heap memory currently being used by the program), then the GC will be executed to clean up the heap. So if `GOGC` is set to 100 (its default value), then GC will run when the New Heap memory becomes double of Live heap memory. 

## New Feature
But with the latest version of GC, comes a new feature to set the memory limit directly upon which GC should run, rather than using `GOGC` as a ratio. So if let's say, live memory is typically 20 MB, you can set the memory limit (you can use `runtime.SetMemoryLimit`) to 45 MB to trigger whenever heap memory size reaches it. So now you can set `GOGC=off` to never let GC run, but trigger it if and when we reach a limit! That's great, but other languages had this feature for a very long time 😂 
The reason for the delay is that, although it seems trivial to set a memory limit, it can essentially make the GC run very frequently if the limit is not set properly, leading to very poor performance of the application! The GC handles this situation too, so if the GC occupies 50% of the CPU for more than 2 secs, GC will be stopped. This mechanism is still under discussion but will evolve over time.

## Conclusion
So, to conclude, `GOGC` should be set according to our constraint of memory and CPU, which one is more precious to us! Setting `GOGC` to a lower value with trigger GC often, leading to lower memory usage but higher CPU usage. On the other hand, setting to a higher percentage will lead to higher memory usage but lower CPU usage. So the goal should be, to test various scenarios and find the sweet spot, and oh, now you can turn it off as well!  
## References
1. https://tip.golang.org/doc/gc-guide
2. Go day 2022, which happened on 3rd November 2022
