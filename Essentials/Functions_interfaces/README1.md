### Interfaces

```go
package main

import "log"

type Animal interface {
 Says() string
 NumberOfLegs() int
}

type Dog struct {
 Name  string
 Breed string
}

type Gorilla struct {
 Name          string
 Color         string
 NumberOfTeeth int
}

func (d Dog) Says() string {
 return "Woof!"
}

func (d Dog) NumberOfLegs() int {
 return 4
}
func main() {
 dog := Dog{
  Name:  "Samson",
  Breed: "German Shepherd",
 }

 //even thought this function is expecting to get type of "Animal" and I'm passign it a type of "Dog"
 //But there's no error, because I made that "dog" satify the requirement of the "Animal" Interface by Adding the two method on it. on the "Dog" struct
 PrintInfo(dog)

 // gorilla := Gorilla{
 //  Name:          "King Kong",
 //  Color:         "Black",
 //  NumberOfTeeth: 32,
 // }

 //PrintInfo(gorilla) // here's an error

}

func PrintInfo(a Animal) {
 log.Println("This Animal Says", a.Says(), "and has ", a.NumberOfLegs(), "Legs")
}

// interfaces here are just like the interface in Typescript
// anything can implement interface by just implementing the required!
// there will be an error if we use Reciever as pointer
func (d *Dog) Say() string {
 return "woof"
}

```
