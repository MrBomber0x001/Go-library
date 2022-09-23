# Generics

parameterized types, where the type could be specified later

Take for example a sum function, if we need to sum different type, it would be a disaster because we need to define function for each type

```go
package main

import "fmt"

func sumInt(a, b int) int {
 return a + b
}

func sumFloat(a, b float64) float64 {
 return a + b
}

func sumString(a, b string) string {
 return a + b
}

func main() {
 fmt.Println(sumInt(1, 2))
 fmt.Println(sumFloat(4.0, 2.0))
 fmt.Println(sumString("a", "b"))
}
```

but with generics it becomes alot easier

```go
// general syntax
func fnName[T constraint](){

}
```

`T` type parameter
`constraint` will be interface that allows any type implementing the interface

```go
func sum[T interface{}] (a, b T) T {
    fmt.Println(a, b)
}

```

starting from 1.18 we can use `any` and we use type inference

```go
func sum[T any](a, b T) T {
    fmt.Println(a,b)
}
```

```go
func main(){
    sum(1,2) // expliciy type argument
    sum[float64](4.0, 2.0)
    sum[string]("a", "b")
    // type inference
    sum(1, 2)
sum(4.0, 2.0)
sum("a", "b")

}
func sum[T interface{}](a, b T) T {
        fmt.Println(a,b)
}
```

```go
func sum[T any](a, b T) T {
    return a + b;
}

func main(){
    fmt.Println(sum(1, 2))
fmt.Println(sum(4.0, 2.0))
fmt.Println(sum("a", "b"))

}
```

we will get an error

```sh
$ go run main.go
./main.go:6:9: invalid operation: operator + not defined on a (variable of type T constrained by any)
```

While constraint of type any generally works it does not support operators.

out interface should define a `type set` containing `int`, `float` and `string`
<img src="https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-III%2Fgenerics%2Ftypeset.png&w=3840&q=75" />

```go
type SumConstraint interface {
    int | float64 | string
}


func sum[T SumConstraint](a, b T) T {
    return a + b
}
func main(){
    fmt.Println(sum(1, 2))
 fmt.Println(sum(4.0, 2.0))
 fmt.Println(sum("a", "b"))
}
```

we can also use the `constraints` package which defines a set of useful constraints to be used with type parameters

![alt](https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-III%2Fgenerics%2Fconstraints-package.png&w=3840&q=75)

```sh
go get golang.org/x/exp/constraints
```

```go
import (
    "fmt"
    "golang.org/x/exp/constraints"
)

func sum[t constraints.Ordered](a, b T) T {
    return a + b
}
func main(){
    fmt.Println(sum(1, 2))
    fmt.Println(sum(4.0, 2.0))
    fmt.Println(sum("a", "b"))
}
```

Here's the implementatio of constrains.Ordered

```go
type Ordered interface {
    Integer | Float | ~string
}
```

`~`:  is a new token added to Go and the expression ~string means the set of all types whose underlying type is string

## when to use Generics

- Functions that operate on arrays, slices, maps, and channels.
- General purpose data structures like stack or linked list.
- To reduce code duplication.
