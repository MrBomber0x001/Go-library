
## Variables

### Table of Contents

- Variable declarations

```go
// single declartion 
var variable_name variable_type = value

// multiple declarations
var var_1, var_2 variable_type = val1, val2
// or
var (
    var_1 variable_type1 = val1
    var_2 variable_type2 = val2
)

// type omission (type will be inferred)
var var_name = val1

// declarations + assigent (shorthand sytax)
var_name := val // shorthand works only inside `function` bodies`

```

- Constants

```go
const constant = value

// only constant can be assigned to other constants
const a = 10
const b = a; // works

var a = 10
const b = 10 // error
```

- Data types

```go
// Strings
var name string = "hello"
var anotherName = `hello 
                    It's me`
var value bool = false

// Numbers (signed)
var i int = 404                     // Platform dependent
var i8 int8 = 127                   // -128 to 127
var i16 int16 = 32767               // -2^15 to 2^15 - 1
var i32 int32 = -2147483647         // -2^31 to 2^31 - 1
var i64 int64 = 9223372036854775807 // -2^63 to 2^63 - 1

// Number (unsigned)
var ui uint = 404                     // Platform dependent
var ui8 uint8 = 255                   // 0 to 255
var ui16 uint16 = 65535               // 0 to 2^16
var ui32 uint32 = 2147483647          // 0 to 2^32
var ui64 uint64 = 9223372036854775807 // 0 to 2^64
var uiptr uintptr                     // Integer representation of a memory address

// Go has two additional integer types called `byte` and `rune` which are alises for `uint8` and `int32`

type byte = uint8
type rune = int32

//A rune represents a unicode code point.
var b byte = 'a'
var r rune = '🍕'


/// Floating point, go has two Floating point types
var f32 float32 = 1.7812 // IEEE-754 32-bit
var f64 float64 = 3.1415 // IEEE-754 64-bit


```

```go
package main

import "fmt"

// type alias
type myAlias = string

// Defined types (doesn't use an equal sign)
type MyDefined string

func main() {
 var b byte = 'a'
 var r rune = '🍕'

 fmt.Println(b)
 fmt.Println(r)

 // Complex numbers
 var c1 complex128 = complex(10, 1)
 var c2 complex128 = 12 + 4i
 // var c4 complex64 = complex(12, 2)
 var c3 complex128 = c1 + c2
 fmt.Println(c1, c2)
 fmt.Println(c3)

 // any variable declared without initialization is zero-valued
 var i int
 var f float64
 var bo bool
 var s string

 fmt.Printf("%v %v %v %q\n", i, f, bo, s)

 // Type conversion
 ii := 42
 ff := float64(ii)
 uu := uint(ff)
 // T for type
 fmt.Printf("%T %T", ff, uu)

 // Aliases
 var str myAlias = "I am an alias"
 fmt.Println("%T - %s", str, str) // output: string - I am an alias

 var myStr MyDefined = "I am defined"
 fmt.Println("%T - %s", myStr, myStr) // output: main.MyDefined - I am defined

 //TODO: what's the difference then between defined types and type alias
 //(Defined Type): it cannot be used interchangeably with the underlying type like alias types.
 var alias myAlias
 // var def MyDefined

 var copy1 string = alias

 //var copy2 string = def // an error

 fmt.Println("hello %s", copy1) // ''
}
```

```go
/* Formatting
1. Printf (enable formatting )
2. Println and Print
3. Sprintf, Sprint, Sprintln 
(return a string instead of printing it) */

func main(){
    percent := (3/5) * 100
    fmt.Printf("%.2f %%", percent);

    s := fmt.Sprintf("hex:%x bin:%b", 10, 10)
    fmt.Println(s)
}
```

```go
// Conditional
1. Simple If/Else
2. Compact If/Else
func main(){
    // simple if
    x := 10
    if x > 5 {
        ...
    } else if x > 4 {
        ...
    }
     else {
        ...
     }

     // compact if (this is quite common pattern)
     if x := 10; x > 5 {
        ...
     }


     // Switch
     1. `break` is automatically added at the end of each case

     switch x {
        case 10: 
            ..
        case 12:
            ..
        case 13:
        case 14:
            ...
        default:
            ...
     }

     2. shorthand syntax
     switch day := "monday"; day {
        case "monday":
            ...
        default:
            ...
     }
     
     3. `fallThrough` keyword to transfer control to the next case even though the currenct case might have matched

     switch day := "monday"; day {
        case "monday":
            fmt.Println("time to workd")
            fallthrough
        case "friday":
            fmt.Println("ley's party")
        default:
            ...
     } // output : time to workd, let's party
     4. using switch without any `condition` == switch true
     x := 40
     switch {
        case x > 4:
            ...
        default:
            ....
     }
}
```

```go
// loops: only `for` loop
// go support break and continue
func main(){
    for i:=0; i < 10; i++ {
        if i < 2 {
            continue
        }
        if i > 5 {
            break 
        }
    }
    i := 0
    for ;i<10; {
        i += 1
    }

    for {
        // forever
    }
}
```

- Variable Shadowing

the entire go application is compiled to a single binary file

### Variables

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

### Pointers

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
