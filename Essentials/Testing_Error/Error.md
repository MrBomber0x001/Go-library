# Error and Error handling

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

## Sentinel Errors

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

## custom errors
