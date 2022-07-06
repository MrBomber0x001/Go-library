
## Types and Struct

- Types and Struct
- Receivers Structs with Functions

### Struct

You can think of struct as a class which support composition but not inheritance

```go
type Person struct {
    FirstName, LastName string
    Age int
}

func main(){
    var p1 Person

    fmt.Println("Person 1:", p1)

    var p2 = Person{
        FirstName: "Karan",
        LastName: "Meska",
        Age: 22,
    }

    // without fields name, you must provide all the fields value
    var p3 = Person {
        "Bruce", "wayne", 22
    }

    // anonymous struct

    var a = struct {
        Name string
    } {"Golang"}

    // creating a pointer to structs as well

    var p = Person {
        FirstName: "Yousef",
        LastName: "Meska",
        Age: 23
    }

    ptr := &p
    fmt.Println((*ptr).FirstName) // we don't need to deference it
    fmt.Println(ptr.FirstName) 
    // we can also use `new` built-in function as we used it before
    p := new(Person)
    p.FirstName = "Ypusef"
    p.LastName = "meska"
    p.Age = 22
     fmt.Println("Person", p)
}
```

**note**:
As a side note, two structs are equal if all their corresponding fields are equal as well.

#### Exported Fields

as we've learnt before that capitalize the function or a variable make it exported
here the same, and for fields also

```go
type Student struct {
    FirwstName string // exported
    age string // not exported 
}
```

#### Compostion over inheritance

1. Emdedding

```go
type Person struct {
    FirstName string
    LastaName string
    Age int
}

type SuperHero struct {
    Person
    Alias string
}
func main(){
    s := SuperHero{}
    s.FirstName = "Bruce"
    s.Alias = "Batman"
}
```

2. Compostion

```go
type Person struct {
    FirstName string
    LastaName string
    Age int
}

type SuperHero struct {
    Person Person
    Alias string
}
func main(){
    p := Person{"Bruce", "Wayne", 40}
    s := SuperHero{p, "batman"}
}
```

#### struct tags

A struct tag is just a tag that allows us to attach metadata information to the field which can be used for custom behavior using the reflect package.

You will often find tags in encoding packages, such as XML, JSON, YAML, ORMs, and Configuration management.

```go
type Struct struct{
    Property string `key:value1`
    Peoperty int `key:value2`
}

type Animal struct {
    Name string `json:"name"`
    Age int `json:"age"`
}
```

#### value or reference ?

Struct are value types not passed by refernece e.g when we pass a struct to another function, the function gets its own copy of the struct.

```go
package main

import "fmt"

type Point struct {
    X,Y float64
}

func main(){
    p1 := Point {1,2}
    p2 := p1 //Copy of p1 is assigned to p2

    p2.x = 2
    fmt.Println(p1) // Output: {1 2}
 fmt.Println(p2) // Output: {2 2}
}

```

Empty structs occupies zero bytes of storage

```go
package main
import (
    "fmt" 
"unsafe"
)
func main(){
    var s struct {}
    fmt.Println(unsafe.Sizeof(s)) // output: 0
}
```

TODO: playaround with Recievers

```go
package main

import (
 "log"
 "time"
)

type User struct {
 FirstName string
 LastName  string
 Age       int
 BirthDate time.Time
}

type myStruct struct {
 FirstName string
 Age       int
}

// Receivers structs with functions
/// A Receivers: ties the function to the type of struct, thats means I've access to my information from that functions, and we need to make it 

func (m *myStruct) printFirstName() string {
 return m.FirstName
}
func (m *myStruct) printAge() int {
 return m.Age
}
func (m* User) changeFirstName() string {
    newValue := "omar meska"
    m.FirstName = newValue;
    return m.FirstName;
}
func main() {
 user := User{
  FirstName: "Travis",
  LastName:  "Meska",
 }
 log.Println(user, user.Age, user.FirstName, user.LastName)
 var myVar myStruct
 myVar.FirstName = "Hamada"
 myVar.Age = 12
 myVar2 := myStruct{
  FirstName: "Yousef",
 }
 log.Println(myVar, myVar2)
 log.Println(myVar.printFirstName())
 log.Println(myVar.printAge())
 log.Println(user.changeFirstName())
}

// functions and variables starting with a captial letter in go, are visiable outside the package it's declared on
// not accessible
func whatever() {

}

//accessible outside the current package
func Whatever() {

}

var special string // not accessible outside
var Special string // accessible outside
```

in js it's similar to

```js

function MyStruct(firstname, age){
    this.firstname = firstname;
    this.age = age;
}

MyStruct.prototype.printFirstName = function(){
    return this.firstname;
}

let myVar = new MyStruct("yousef", "22");
myVar.printFirstName()
```

we can've buisness logic inside this Receivers
