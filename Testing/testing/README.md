tests.go

```go
package main

import (
 "errors"
 "log"
)

func main() {
 result, err := divide(100.0, 10.0)
 if err != nil {
  log.Println(err)
  return
 }
 log.Println("result of division is", result)
}

func divide(x, y float32) (float32, error) {
 var result float32

 if y == 0 {
  return result, errors.New("Cannot Divide By 0")
 }
 result = x / y
 return result, nil
}

```

tests_test.go

```go
package main

import "testing"

var tests = []struct {
 name     string
 dividend float32
 divisor  float32
 expected float32
 isErr    bool
}{
 {"valid-data", 100.0, 10.0, 10.0, false},
 {"invalid-data", 100.0, 0.0, 0, true},
}

// tt is naming convension when randing over tests
func TestDivison(t *testing.T) {
 for _, tt := range tests {
  got, err := divide(tt.dividend, tt.divisor)
  if tt.isErr {
   if err == nil {
    t.Error("expected an erro but did not get one")
   }
  } else {
   if err != nil {
    t.Errorf("did not expect an error but gone one.", err.Error())
   }
  }
  if got != tt.expected {
   t.Errorf("Expected %f but got %f", tt.expected, got)
  }
 }
}

func TestDivide(t *testing.T) {
 _, err := divide(10.0, 1.0)
 if err != nil {
  t.Error("Got an error we should not have")
 }
}

func TestBadDivide(t *testing.T) {
 _, err := divide(10.0, 0)
 if err == nil {
  t.Error("Did not get an eror when we should have")
 }
}

/// the func must begin with word Test

// go test -v or go test -v -cover
```
