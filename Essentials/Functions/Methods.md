
## Methods (function receivers)

Go has types. And, you can define methods on types.\
A method: a normal function with a special receiver argument
a method has a type and struct

```go
//general syntax
func (variable T) Name(params) (returnTypes){

}
```

```go
type Car struct {
    Name string
    Year int
}
func (c Car) IsLatest() bool{
    return c.Year >= 2017 // think of it as `this`
}
func main(){
    c := Car {"Tesla", 2021}
    fmt.Println("IsLatest", c.IsLatest())
}

```

### Methods with Pointer receivers

For now, all what we've seen is operating on values, but what if i need to **update** a value

```go
func (c Car) UpdateName(name string){
    c.Name = name
}
func main(){
    c := Car {"Tesla", 2021}
    c.UpdateName("Toyota")
    fmt.Println("Car:", c) // Car: {Tesla 2021}
}
```

the name hasn't changed, because the function takes only a copy of the struct not a reference to it

```go
func (c *Car) UpdateName(name string){
    c.Name = name
} // the name will be changed
```

### Properties

- we can omit the variable part of the receiver as well if we're not using it

```go
func (Car) UpdateName(...){}
```

- Methods are not limited to structs but also used with `non-struct` types as well

```go
package main

import "fmt"

type MyInt int


func (i MyInt) isGreater(value int) bool {
    return i > MyInt(value)
}

func main(){
    i := MyInt(10)

    fmt.Println(i.isGreater(5))
}
```

### why methods not functions?

One thing I can think of right now is that methods can help us avoid naming conflicts.

Since a method is tied to a particular type, we can have the same method names for multiple receivers.

And of course it's a matter of preference
