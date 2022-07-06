
### Slices

A slice is a segment of an array, slices build on arrays and provde more power

<img src="https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-II%2Farrays-and-slices%2Fslice.png&w=1920&q=75" />

A slice consists of three things:

- A `pointer` reference to an underlying array.
- the lenght of the segment of the array that slice contains
- capacity which the maximum size up to which the segment can grow

```go
package main

import "fmt"

func main() {
 a := [5]int{1, 2, 3, 4, 5}

 s := a[1:4]

 fmt.Printf("Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))
 fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))

}
```

we don't need to specify any length

```go
// general syntax
var s []T
```

unlike arrays, the zero value of the slice is not `zeros` but a `nil`

```go
var s []string
fmt.Println(s)// [] 
 fmt.Println(s == nil) //true
```

There are multiple ways to initialize our slice, One way is to use the built-in `make` function and the other way to define a `slice literal`.

```go
make([]T, len, cap) []T

func main()[
    var s = make([]string, 0,0)
]
```

```go
var s = []string {"Go", "Typescript"}
```

Another way is by creating a segment of array

```go
a[low:high]

func main(){
    var a = [4]string {
        "c++",
        "Go",
        "Typescript",
        "Java"
    }
    s1 := a[0:2] // Select from 0 to 2
 s2 := a[:3]  // Select first 3
 s3 := a[2:]  // Select last 2
}
```

**note**:

- we can create a slice from another slice as well.
- we can iterate over slices as we iterate over arrays exactly

**Built-in slice functions**

- Copy: return the number of elemented copied
- Append: return a new slice containing all the elements

```go
//general syntax
copy(dst, src []T) int

package main

import "fmt"

func main() {

 s1 := []string{"a", "b", "c", "d"}
 s2 := make([]string, len(s1))

 e := copy(s2, s1)

 fmt.Println("Src:", s1)
 fmt.Println("Dst:", s2)
 fmt.Println("Elements:", e)
}
//output
Src: [a b c d]
Dst: [a b c d]
Elements: 4

```

```go
append(slice []T, elems ...T) []T

func main(){
    s1 := []string {"a", "b", "c", "d"}
    s2 := append(a1, "e", "f")
fmt.Println("a1:", a1)
 fmt.Println("a2:", a2)
    
}
```

**note**:
But if the given slice doesn't have sufficient capacity for the new elements then a new underlying array is allocated with a bigger capacity.

All the elements from the underlying array of the existing slice are copied to this new array, and then the new elements are appended.

#### properties

1. Slices are reference types, unlike arrays
=> modifying the elements of a slice will modify the corresponding elements in the referenced array.

```go
package main

func main(){
    a := [7]string {"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    
    s := a[0:2]

    s[0] = "Sun"

     fmt.Println(a) // Output: [Sun Tue Wed Thu Fri Sat Sun]
 fmt.Println(s) // Output: [Sun Tue]

}
```

2. Slices can be used with variadic types

```go
package main

func main(){
    values := []int {1,3,4}
    sum := add(value...)
    fmt.Println(sum)

}

func add(values ...int) int {
    sum := 0
    for _, v :=range values {
        sum += v
    }
    return sum
}
```
