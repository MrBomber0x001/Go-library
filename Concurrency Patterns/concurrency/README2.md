
## Table of contents

1. Intro go CSP Concurrency model in Go (routines, channels, mutex, callback)
2. CSP concurrency patterns
3. Mutex
4. Deadlocks
5. Concurrency Patterns

Go is a multi-paradigm programming language

## Hubmle Introduction to concurrency

- Advantages
- Concept of Concurrency
- Differentiating between Concurrency and Parallelism
- CSP versus actor based Concurrency

Go's Concurrency allows us to:

- Construct `Streaming data pipelines`
- Make efficient use of I/O and multiple CPUs
- Allows complex systems with multiple components
- `Routines` can start, run and complete simultaneously
- Raises the efficiency of the applications

## Goroutines

- Creating first Goroutines
- Launching anonymous functions
- Using WaitGroup

we use goroutine to execute code we need to run concurrently
In go we achieve concurrency by working with `goroutines`, in fact the `main` loop of go could be considered goroutine
goroutines are used in places where we might use actors

- they are not threads, we can create millions of goroutines
- incredibly cheap
- Have small growth of stack

## channels

In simpler words, they are means to sending information from part of your program to another part of your program

```go
// You can edit this code!
// Click here and start typing.
package main

import (
 "log"
 "math/rand"
 "time"
)

const numPool = 10

func CalculateValue(intChan chan int) {
 randomNumber := RandomNumber(numPool)
 intChan <- randomNumber
}

func main() {
 intChan := make(chan int) // a channel that can hold only int
 defer close(intChan)      // as good practice to close the channel after main finish executing, that's why we use defer

 go CalculateValue(intChan)

 num := <-intChan
 log.Println(num)
}

func RandomNumber(n int) int {
 rand.Seed(time.Now().UnixNano())
 value := rand.Intn(n)
 return value
}

```

| output |
| -- |
| 4 |
**A real world example**

## Progress

- [ ] concurrency
- [ ] Goroutines
- [ ] callbacks
- [ ] Mutexes
- [ ] Channels
- [ ] Using it all  - Concurrent Singleton
