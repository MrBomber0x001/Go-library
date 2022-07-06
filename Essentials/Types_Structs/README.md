## Table Of Contents

- Types and Struct
- Receivers Structs with Functions

## TODO: playaround with Recievers

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
