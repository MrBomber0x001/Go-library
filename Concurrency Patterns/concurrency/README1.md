# Table of Content

## Concurrency vs Parallelism

Using concurrency, we can achieve the same results in lesser time, thus increasing the overall performance and efficiency of our programs.

![alt](https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-IV%2Fconcurrency%2Fconcurrency-vs-parallelism.png&w=2048&q=75)

As Rob pike once said:
> Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once

There are two types of Concurrency models

1. Actor model
2. CSP (Communication Sequential Processing) model

Go relises on CSP mode

## CSP

CSP describe interfactions between concurrent  and gives us a mental model of thinking about concurrency

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

A go routin follow the the fork-join model of concurrency where a child process splits from it's parent process and run concurrently with it,
and after the child process finishes execution, it merges back into it's parent in point called 'join point'

<put an image>

```go
package main

func sayHi(message string){
    fmt.Println(message);
}

func main(){
    go sayHi("hi");
    time.Sleep(1 * time.Second); // notice this line
}
```

if we didn't put this time.Sleep, the main go routine func will finishe execution and will not wait for the sayHi() output, this line is for delaying it

## Channels

```go
var ch chan T // <nil>
ch := make(chan T) // will be initialized

```

T is the data type of the value we want to send and receive via channel

### Sharing data between channels

```go
package main

func speak(arg string, ch chan string){
    ch <- arg // send
}
func main(){
    ch := make(chan string)

    go speak("Hello world", ch)

    data := <- chan // receive
    fmt.Println(data)
}
```

### Buffered channels

```go
func main(){

    ch := make(chan string, 3)

    go speak("Hello world 1", ch)
    go speak("Hello world 2", ch)
    go speak("Hello world 3", ch)
    go speak("Hello world 4", ch)

    data1 := <- ch
    fmt.Println(data1)

    data2 := <- ch
    fmt.Println(data2)

    data3 := <- ch
    fmt.Println(data3)

    data4 := <- ch
    fmt.Println(data4)
}
```

No guarantee of the order

By default, a channel is unbuffered and has a capacity of 0, hence, we omit the second argument to the make function.

### Directional channel

here the channel can only send, and will panic if we try to receive value

```go
func speak(message string, ch chan <- string){
    ch <- arg // send only
}

```

### Closing a channel

using close(), after finishing using channels, close it

```go
// your code

close(ch)
```

we can optionally test whether a channel has been closed or not by assigning a second parameter to receivers

```go
// your code

data1, ok := ch
fmt.Println(data1, ok)
close(ch)
```

if `ok` is `false` that's means there are no more values to receive and the channel is closed

this is similar to check if a key exists or not in a map

### Properties

1. a send to nil channel blocks forever (deadlock) or receiving from it

```go
var c chan string
c <- "Hello, World!" // Panic: all goroutines are asleep - deadlock!
fmt.Println(<-c) // Panic: all goroutines are asleep - deadlock!

```

2. a send to closed channel panics

```go
var c = make(chan string, 1)
c <- "hello world"
close(c)
c <- "heloo world" // Panic: send on closed channel
```

3. A receive from a closed channel returns the zero value immediately

```go
var c = make(chan int, 2)
c <- 5
c <- 4
close(c)

for i:=0; i<4;i++ {
    fmt.Println("%d", <-c) // output: 5 4 0 0 
}
```

4. we can range over channels using `for` and `range`

```go
 ch := make(chan string, 2)

 ch <- "Hello"
 ch <- "World"

 close(ch)

 for data := range ch {
  fmt.Println(data)
 }

```

## Select

TODO: <https://www.youtube.com/watch?v=tG7gII0Ax0Q>
select allows us to listen for multiple channels, and which one is ready first and read from it
and then going to listen for both of them again and which one is ready first also read from it

The select statement blocks the code and waits for multiple channel operations simultaneously.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

Select like switch has a default statement, which will be exectuted if no channels are ready

```go
func main() {
 one := make(chan string)
 two := make(chan string)

 for x := 0; x < 10; x++ {
  go func() {
   time.Sleep(time.Second * 2)
   one <- "One"
  }()

  go func() {
   time.Sleep(time.Second * 1)
   two <- "Two"
  }()
 }

 for x := 0; x < 10; x++ {
  select {
  case result := <-one:
   fmt.Println("Received:", result)
  case result := <-two:
   fmt.Println("Received:", result)
  default:
   fmt.Println("Default...")
   time.Sleep(200 * time.Millisecond)
  }
 }

 close(one)
 close(two)
}

```

It's important to note, an empty select {} will block forever

```go
// one go routine is sending a message every second and other is sending a message every two seconds
func main(){
    chans := []chan int{ 
        make(chan int),
        make(chan int)
    }

    for i := range chans {
        go func(i int, ch chan <- int){
            for {
                time.Sleep(time.Duration(i) * time.Second)
                ch <- i
            }
        }(i+1, chans[i])
    }

    for i := 0; i < 12; i++ {
        i1 := <- chans[0] // a problem
        i1 := <- chans[1] // a problem

        select {
            case m0 := <- chans[0]:
                log.Println("received", m0)
            case m1 := <- chans[1]:
                log.Println("received", m1)
        }
    }
}

```

**Another practical example**

```go
package main
type result struct {
    url string
    err error
    latency time.Duration
}

func get(url string, ch chan <- result ){
    start := time.Now()

    if resp, err := http.Get(url); err != nil {
        ch <- result {url, err, 0}
    } else {
        t := time.Since(start).Round(time.Millisecond)
        ch <- result {url, nil, t}
        resp.Body.Close()
    }
}

func main(){
    results := make(chan result)
    list := []string{
        "http://amazon.com",
        "http://google.com",
        "http://facebook.com",
    }

    for _, url := range list {
        go get(url, results)
    }

    for range list{
        r := <-results
        if r.err != nil {
            log.Printf("%-20s %s\n", r.url, r.err)
        } else {
            log.Printf("%-20s %s\n", r.url, r.latency)
        }
    }
}
## Sync Package
