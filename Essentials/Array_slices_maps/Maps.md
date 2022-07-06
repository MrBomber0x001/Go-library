# Maps

Maps are un-orderd key-value pair data structure
keys are unique, but values are not, used for fast lookup, deletion

```go
// general syntax
var m map[K]V
k=> key type
V => v type


func main(){
    var m map[string]int
}
```

the zero value of map is `nil`.
A nilmap has no keys. Moreover, any attempt to add keys to a nilmap will result in a runtime error

## Initialization

1. Make functions
2. map literal

```go
func main(){
    var m = make(map[string]int )


    var m = map[string]int {
        "a": 0,
        "b": 1, // note the last trailing comma is neccessary
    }
}
```

we can make our custom types as well

```go
type User struct {
    Name string
}

func main(){
    var m = map[string]User{
        "a": User{"Peter"}, // we can remove "User" keyword
        "b": User{"Seth"},
    }
    fmt.Println(m)
}

```

## Operations on Maps

1. Adding
2. Retrieving
3. Updating
4. Deleting
5. Iteration

**Adding**

```go
m[key] = v
m["c"] = User{"Steve"}

```

**Retrieving**

```go

c, ok := m["c"] // return a bool value indicate if the key exists also or not

// if you used a key that is not presen in the map?
// you will get the zero value of that map's type.
```

**Updating**
by re-assigning

```go
m["a"] = "roger"
```

**deleting**
we can delete using the built-in `delete` function

```go
delete(map, key_to_be_deleted) # doesn't return any value
```

**Iteration**
because maps are un orderd collection, it's not guaranteed that the result will be always the same every time we iterate

```go
func main(){
    var m = map[string]User{
        "a": {"Peter"},
        "b": {"meska"},
    }
    m["c"] = User{"Steve"}

    for k,v := range m {
        fmt.Println("Key: %s, Value: %v", key, value)
    }
}
```

## Properties

- maps are `reference` type

```go
package main

import "fmt"

type User struct {
 Name string
}

func main() {
 var m1 = map[string]User{
  "a": {"Peter"},
  "b": {"Seth"},
 }

 m2 := m1
 m2["c"] = User{"Steve"}

 fmt.Println(m1) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
 fmt.Println(m2) // Output: map[a:{Peter} b:{Seth} c:{Steve}]
}
```
