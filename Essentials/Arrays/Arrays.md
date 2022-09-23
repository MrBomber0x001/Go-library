
### Arrays

Fixed-size collection of the same type data,
once array is declared, neither the type nor the length can be changed, to add more element you need to make a new array with the length you need and copy from an old array, all the array elements are initialized with their respective zero-value

```go
// general syntax 
var array[n] T
var a = [n]T{V1, V2, V3, ... , Vn}
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

### Array with pointers

Copying an array of pointers copies the pointer values and not the values that the pointers are pointing to

```go
var array1 [3]*string
 array2 := [3]*string{new(string), new(string), new(string)}
 *array2[0] = "Red"
 *array2[1] = "Green"
 *array2[2] = "Black"

 array1 = array2
 *array1[0] = "Yellow"
 fmt.Println(`array1`, array1)
 fmt.Println(`array2`, array2)
 for i := 0; i < len(array1); i++ {
  fmt.Println(*array1[i])
  fmt.Println(*array2[i])
 }
```

### Arrays with functions

```go
//* passing an array to a function is very expensive, because the array elements gets copied and passed to the function
 var array[1e6] int
 foo(array)


 func foo(array [1e6]int){
  // each time foo is called, 8MB of memeory has to be allocated on the stack and then the value
  // of the array has to be copied into that allocation
 }

 // we can handle this situation better by passing a pointer to the array
 foo(&array);
 func foo(array *[1e6] int){
  // instead we copied 8bytes.
  // each time foo is called, only 8 bytes are required of memeory to be allocated on the stack for the pointer variable
 }
```
