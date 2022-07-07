# Testing

test files are presuffixed with _test, and testing is built-in Go

```sh
.
├── go.mod
├── main.go
└── math
    ├── add.go
    └── add_test.go

```

```go
pacakge math
func Add(a,b int) int {
    return a + b
}
```

```go
package main
import (
    "example/math"
    "fmt"
)

func main(){
    result := math.Add(2,2)
    fmt.Println(result)
}
```

 we can write our test in the same package if we wanted, but I personally think doing this in a separate package helps us write tests in a more decoupled way.

```go
packge math_test
import "testing"
func TestAdd(t *testing.T) {
    got := math.Add(1,1)
    expected := 2

    if got != expected {
        t.Fail()
    }
}
```

```sh
$ go test ./math
ok      example/math 0.429s

$ go test ./.. # to test all the packages
?       example [no test files]
ok      example/math 0.34

$ go clean -testcache # to avoid caching testing result problems and run the test after that 
```

## Table driven tests

```go
package math_test

import (
    "example/math"
    "testing"
)

type addTestCase struct {
    a,b, expected int
}
// slice of structs
var testCases = []addTestCase {
    {1,1,3},
    {25,25,50},
    {2,1,3},
    {1, 10, 11}
}

func TestAdd(t *testing.T) {
    for _, tc := range testcases {
        got := math.Add(tc.a, tc.b)
        if got != tc.expected {
            t.Errorf("Expected %d but got %d", tc.expected, got)
        }
    }
}
```

## code coverage

How much of your code got covered

```sh
$ go test ./math -coverprofile=coverage.out
ok      example/math    0.385s  coverage: 100.0% of statements
$ go tool cover -html=coverage.out
```

## Fuzz testing

in 1.18v, a type of automated testing that continuously manipulates inputs to a program to find bugs
Go fuzzing uses coverage guidance to intelligently walk through the code being fuzzed to find and report failures to the user.
Since it can reach edge cases that humans often miss, fuzz testing can be particularly valuable for finding bugs and security exploits.

```go
func FuzzTestAdd(f *testing.F) {
    f.Fuzz(func(t *testing.T, a, b int){
        math.Add(a,b)
    })
}
```

```sh
$ go test -fuzz FuzzTestAdd example/math
fuzz: elapsed: 0s, gathering baseline coverage: 0/192 completed
fuzz: elapsed: 0s, gathering baseline coverage: 192/192 completed, now fuzzing with 8 workers
fuzz: elapsed: 3s, execs: 325017 (108336/sec), new interesting: 11 (total: 202)
fuzz: elapsed: 6s, execs: 680218 (118402/sec), new interesting: 12 (total: 203)
fuzz: elapsed: 9s, execs: 1039901 (119895/sec), new interesting: 19 (total: 210)
fuzz: elapsed: 12s, execs: 1386684 (115594/sec), new interesting: 21 (total: 212)
PASS
ok      foo 12.692s
```

if we add an edge case

```go
func Add(a,b) int {
    if a > b + 10 {
        panic("B must be greateer than A")
    }
    return a + b
}
```

```sh
$ go test -fuzz FuzzTestAdd example/math
warning: starting with empty corpus
fuzz: elapsed: 0s, execs: 0 (0/sec), new interesting: 0 (total: 0)
fuzz: elapsed: 0s, execs: 1 (25/sec), new interesting: 0 (total: 0)
--- FAIL: FuzzTestAdd (0.04s)
    --- FAIL: FuzzTestAdd (0.00s)
        testing.go:1349: panic: B is greater than A
```

**Trivial Example**

```go
package main


func main(){
    result, err := Divide(100.0, 10.0) 
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("result of division is", result)
}

func Divide(x, y float32) (float32, error) {
    var result float32

    if y <= 0 {
        return result, errors.New("cannot divide by 0")
    }
    result = x / y
    return result, nil
}
```

```go
package main
import "testing"

var tests = [] struct {
    name string
    divident float32
    divisor float32
    expected float32
    isErr bool
}{
    {"valid-date", 100.0, 10.0, 10.0, false},
     {"invalid-date", 100.0, 0.0, 0.0, true},
     {"expect-fraction": -1.0, -777.0, 0.0012834, false}
}
func TestDivide(t *testing.T) {
    _, err := Divide(10.0, 1.0)
    if err != nil {
        t.Error("Got an error when we should not have")
    }
}

func TestBadDivide(t *testing.T) {
    _, err := Divide(10.0, 0)
    if err == nil {
        t.Error("Did not get an error when we should have")
    }
}

func TestDivision(t *testing.T){
    for _, tt := range tests {
        got, err := Divide(tt.dividen, tt.divisor)
            if tt.isErr {
        if err == nil {
            t.Error("expected an error but did n't get one!")
        }
    } else {
        if err != nil {
            t.Error("did n't expect an error but got one", err.Error())
        }
    } 

    if got != tt.expected {
        t.Error("expected %f but got %f", tt.expected, got)
    }
    }

}
```
