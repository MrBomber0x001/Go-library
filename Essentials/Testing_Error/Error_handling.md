# Table of Contents

- Error handling
- Panic and recover

## Error and Error handling

There's no execptional handling in Go, instead we use the built-in `error` type which is an interface type

There are multiple ways to construct an error, but there're two most common ones

1. using `new`

2. `fmt.Errorf` returns an error

```go
type error interface {
    Error() string
}

// # An example
package main
import "errors"
func main(){
    result, err := Divide(4,0)

    if err != nil {
        fmt.Println(err)
        // Do something with error
        return
    }
    fmt.Println(result)

}
func Divide(a, b int) (int, error) {
    if b <= 0 {
        return 0, errors.New("cannot divide by zero")// a convesion that error message must begin with lower case

        // or
        return 0, fmt.Errorf("cannot divide %d by zero", a)

    }
 return a/b, nil
}

```

### Sentinel Errors

 defining expected Errors so they can be checked explicitly in other parts of the code

 ```go
 package main

import (
 "errors"
 "fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func main() {...}

func Divide(a, b int) (int, error) {
 if b == 0 {
  return 0, ErrDivideByZero
 }

 return a/b, nil
}
```

### custom errors

## Panic and Recover

there are some situations where the program cannot continue.
A built-in function that stops the normal execution of the current `goroutine`.
 When a function calls panic, the normal execution of the function stops immediately and the control is returned back to the caller. This is repeated until the program exits with the panic message and stack trace.

```go
func panic(interface{})
```

```go
package main

func main(){
    WillPanic()
}

func WillPanic(){
    panic("Woah")
}
```

it's possible to regain control of a panicking program using the built-in `recover` function along with `defer`

```go
func recover() interface{}
```

```go
package main 
func main(){
    WillPanic()
}

func handlePanic(){
    data := recover()
    fmt.Println("Recovered:", data)
}

func WillPanic(){
    defer handlePanic()
    panic("Woah")
}
```

```sh
go run main.go
Recovered: Woah
```

we can consider `panic/recover` as similar to `try/catch`
🔺we should use important factor is that we should avoid panic and recover and use errors when possible

### when to use

- An recoverable error:
Which can be a situation where the program cannot simply continue its execution.
ex: For example, reading a configuration file which is important to start the program, as there is nothing else to do if the file read itself fails.

- Developer error
This is the most common situation. For example, dereferencing a pointer when the value is nil will cause a panic.
