# Table of Contents

- Variable declarations
- Variable Shadowing


the entire go application is compiled to a single binary file
## Variables

```go
package main

import "log"

// Variable shadowing
var s = "string" // package level variable
// every main package should've a main function or every go application must have main function 
// main funcion does not take any parameters
func main() {
 var s2 = "xxx"
 log.Println(`s`, s)
 log.Println(`s2`, s2)
 saySomething("xxx")
 saySomethinElse("hola")

 s := "another string" // s here is not s in the package level, we use := to implicitly tell the compiler to inheir from the value without the need to specifiy the need for telling the type
 log.Println(s)
}

// overwritting
func saySomething(s string) (string, string) {
 log.Println(`s from saySomething func is`, s) // xxx
 return s, "World"
}

func saySomethinElse(s3 string) string {
 log.Println(`s from saySomethingElse is `, s) // 'string'
 return s3
}

```

- Variable Scope

```go
package main

import (
 "fmt"
 "log"
)

// every go file should've a package
// every go app should've only one main func

func main() {
 var whatToSay string
 var saySomethingElse string
 var i int
 fmt.Println("Hello world")
 whatToSay = saySomething("Learning Go")
 fmt.Println(whatToSay)
 log.Println(whatToSay)
 saySomethingElse = saySomething("Something Else")
 log.Println(saySomethingElse)
 log.Println(saySomething("Finally"))
 i = 7
 log.Println(i) // 7, without 7 it would be 0
 var world string
 whatToSay, world = returnTwoValues("Hello")
 log.Println("world + whatToSay", whatToSay, world)
 // _ is a convesion for ignoring
 saySomethingElse, _ = returnTwoValues("Hello")
 log.Println("finally", saySomethingElse)

}

func saySomething(s string) string {
 return s
}

// functions can return more than a value
func returnTwoValues(s string) (string, string) {
 return s, "World"
}

// if you create a variable, you should use it, otherwise the compiler will refuse to compile
// main function doesn't take any parameters at all.
// main is the entry point

// go has first class functions

// _ (underscore) means ignore
```

- Pointers

```go
package main

import "log"

func main() {
 // myStirng is only scoped to the main func
 var myString string
 myString = "Good"
 log.Println("First", myString)
 changeUsingPointer(&myString)
 log.Println("Second", myString)
}

// Pass by Reference
func changeUsingPointer(s *string) {
 log.Println("s is set to", s, "but the actual value contained in that address is", *s, "whai is &s")
 newValue := "red"
 *s = newValue
}

// A pointer points to a specific location "Memory Address"
```

### Decision Structure

```go
package main

import "log"

func main() {
 isTrue := true
 var isFalse bool
 isFalse = false

 if isTrue {
  log.Println("isTrue is", isTrue)
 } else {
  log.Println("isFalse is", isFalse)
 }

 myNum := 100
 isNotFalse := true

 if myNum > 99 && isNotFalse {
  log.Println("myNum is greater than 99")
 }
 myVar := "cat"

 switch myVar {
 case "dog":
  log.Println("myVar is set to", myVar)

 case "cat":
  log.Println("myVar is set to", myVar)

 default:
  log.Println("cat is something else")
 }

}
```

## Working with JSON

```go
package main

import (
 "encoding/json"
 "fmt"
 "log"
)

type Person struct {
 FirstName string `json:"first_name"`
 LastName  string `json:"last_name"`
 HairColor string `json:"hair_color"`
 HasDog    bool   `json:"has_dog"`
}

func main() {
 myJson := `
 [
  {
   "first_name": "yousef",
   "last_name": "meska",
   "hair_color": "black",
   "has_dog": true
  },
  {
   "first_name": "omar",
   "last_name": "meska",
   "hair_color": "whirte",
   "has_dog": false
  }
 ]
 `

 // write json into struct
 var unmarshalled []Person
 err := json.Unmarshal([]byte(myJson), &unmarshalled)
 if err != nil {
  log.Println("Error unmarshalling json", err)
 }

 log.Printf("unmarshalled: %v", unmarshalled)

 // write json from a struct
 var mySlice []Person
 var m1 Person
 m1.FirstName = "Wally"
 m1.LastName = "West"
 m1.HairColor = "red"
 m1.HasDog = false
 mySlice = append(mySlice, m1)

 var m2 Person
 m2.FirstName = "yousef"
 m2.LastName = "meska"
 m2.HairColor = "yellow"
 m2.HasDog = true
 mySlice = append(mySlice, m2)

 newJson, err := json.MarshalIndent(mySlice, "", " ")
 if err != nil {
  log.Println("error marshalling", err)
 }

 fmt.Println(string(newJson))

}

// For development purpose use MarsharIndex, for production use Marshal

// this topic was about writing and reading json with known format, but what if i don't know
// there's another way to deal with such situation
//TODO: search for it
```
