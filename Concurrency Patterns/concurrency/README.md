# Table of Content

## Concurrency vs Parallelism

Using concurrency, we can achieve the same results in lesser time, thus increasing the overall performance and efficiency of our programs.

![alt](https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-IV%2Fconcurrency%2Fconcurrency-vs-parallelism.png&w=2048&q=75)

As Rob pike once said:
> Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once

There are two types of Concurrency models

1. Actor model
2. CSP (Communication Sequential Processing) model

## CSP

CSP describe interfactions between concurrent processes

independept concurrent processes communicate with each other sharing channels between them

<img src="https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-IV%2Fconcurrency%2Fcsp.png&w=2048&q=75">

### Basic Concepts

- Data race: two processes or more trying to access the same resoruces concurrently
maybe one was reading while the other was writing
- Race conditions: the timing or order of events affects the correctness of piece of code
- Deadlock: all processesa are blocked while waiting for each other and the program cannot procesed further
- Livelocks: processes that perform concurrent operations, but do nothing to move the state of the state of the program forward
- Starvation: when a process is deprived of necessary resources and unable to complete its functions, might happen because of `deadlock` or `insuffienct scheudling algorithms`

#### Deadlocks

There are four 4 conditions, known as `Coffman conditions` which when are satisfied, then a deadlock occurs

1. mutual execution
2. Hold and wait
3. No Preemption
4. Circular wait

1. Mutual Execution
A concurrent process holds at least one resource at any one time making it non-sharable.

![alt](https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-IV%2Fconcurrency%2Fmutual-exclusion.png&w=1920&q=75)

2. Hold and wait
A concurrent process holds a resource and is waiting for an additional resource.

Process 2 is allocating (holding) resource 2 and 3 and waiting for resources 1 which is locked by process 1.
![alt](https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-IV%2Fconcurrency%2Fhold-and-wait.png&w=1920&q=75)

3. No preemption
A resource held by a concurrent process cannot be taken away by the system. It can be freed by the process holding it.

In the diagram below, Process 2 cannot preempt Resource 1 from Process 1. It will only be released when Process 1 relinquishes it voluntarily after its execution is complete.
![alt](https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-IV%2Fconcurrency%2Fno-preemption.png&w=1920&q=75)

4. Circular wait
A process is waiting for the resource held by the second process, which is waiting for the resource held by the third process, and so on, till the last process is waiting for a resource held by the first process. Hence, forming a circular chain.
(All if waiting)

![alt](https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-IV%2Fconcurrency%2Fcircular-wait.png&w=1920&q=75)

## gorountines

Rob pike legenedary qoute:
> Don't communicate by sharing memory, share memory by communications

A _goroutine_ is a lightweight `thread of execution` that is manages by the Go runtime. not an actual thread, you could've million of goroutines
the main function run as a goroutine.
essentially let us write asynchronous code in a synchronous manner

There's a `runtime scheduler` by go which uses cooperative scheduling.
A single thread may run thousands of goroutines in them

we can turn any function into a goroutine by simply using the `go` keyword

```go
go fn(x, y, z)
```

### Fork-join model
