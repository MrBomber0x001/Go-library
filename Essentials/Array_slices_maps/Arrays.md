
### Arrays

Fixed-size collection of the same type data

```go
// general syntax 
var a[n] T
var a[n]T = [n]T{V1, V2, V3, ... , Vn}
```

```go
func main(){
    var arr = [4]int {1,2,3,4}
    arr := [4] int {1,2,3,4} // shorthand
    // looping: 1. len() 2. range 
    for i := 0; i < len(arr); i++ {
        fmt.Prinf("Index: %d, Element: %d\n", i, arr[i])
    }
    for i, e:= range arr {
        fmt.Prinf("index %d, Element: %d\n", i, e)
    }

    for range arr {

    }

    // multi-dimensional

    arr := [2][4]int {
        {1,2,3,4},
        {5,6,7,8}
    }
    for i, e := range arr {
        fmt.Prinf("Index: %d, Element: %d\n", i, e)
    }

    // let the compiler infer the length of the array and by using [...] instead

    arr := [...][4]int {
        {1,2,3,4},
        {4,5,6,7}
    }
}
```

#### Properties

Those Properties are really important and to consider when dealing with arrays

- Arrays's length is part of its type

Arrays with different sizes are not equal, this mean we can resize an array because we are producing a new type

```go
package main

func main(){
    var a = [4] int {1,2,3,4}
    var b[2] int = a // Error, can't use a (type [4]int) as type [2]int in assignment
}
```

- arrays are value-types not passed by reference

```go
func main(){
    var a = [7] string {"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    var b = a // Copy of a is assigned to b
    b[0] = "Monday"

    fmt.Println(a) // Output: [Mon Tue Wed Thu Fri Sat Sun]
 fmt.Println(b) // Output: [Monday Tue Wed Thu Fri Sat Sun]
}
```
