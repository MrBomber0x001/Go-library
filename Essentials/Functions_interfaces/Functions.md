
## Functions

```go

// 1. simple declartion
// 2. Returning value, multiple values
// 3. named returns
func main(){
    s, i:= myFunction("helo")
}
func myFunction(p1 string) (string, int) {
    msg := fmt.Sprinf("%s function", p1)
    return msg, 10
}

func myOtherFunction(p1 string) (s string, i int){
    s = fmt.Sprintf("%s function", p1)
    i = 10 
    return // naked return
}


// Functions are first class in go


func foo(){
    fn := func(){
        fmt.Println("inside fn")
    }

    fn();
}

func bar(){
    func(){
        fmt.Println("inside fn")
    }() // anonymous function
}
```

#### Closures

A simple definition can be that a closure is a function value that references variables from outside its body.
In other words, a closure is an inner function that has access to the variables in the scope in which it was created
This applies even when the outer function finishes execution and the scope gets destroyed.

```go
func main(){

    add := myFunction()

    add(5)

    fmt.Println(add(10))
}
func myFunction() func(int) int {
    sum := 0

    return func(v int) int {
        sum += v
        return sum
    }
}

```

```go
// Variadic function
func main() {
 sum := add(1, 2, 3, 5)
 fmt.Println(sum)
}

func add(values ...int) int {
 sum := 0

 for _, v := range values {
  sum += v
 }

 return sum
}
```

```go
/// init function
package main

import "fmt"

func init() {
 fmt.Println("Before main!")
}
func init() {
 fmt.Println("Hello again?")
}
func main() {
 fmt.Println("Running main!")
}

/// unlinke main, there can be more than one init function in one or multiple files
/// if `init` is declared in mulitple files, are processed according to the lexicographic filename order
/// it may be very useful for establishing a database connection, fetching configuration files, setting up environment variables, etc.
```

```go
// defer and defer stack
//defer statements are stacked and executed in a last in first out manner
// defer is incredibly useful for doing cleanup or error handling
func main(){
    defer fmt.Prinln("I am a yousef")
    defer fmt.Println("Are you")
    fmt.Println("Doing some work")
}
```

```bash
$ go run main.go
Doing some work...
Are you
I am a yousef
```
