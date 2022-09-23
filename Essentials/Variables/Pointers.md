
## Pointers

pointers are very useful in many cases, it helps us change a local scoped variables from outside it's body function

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

// A pointer points to a specific location "Memory Address", this memory location contains the value
```

## Pointers

it can be used like this:

```go
// General syntax 
var x *T // T is any type ya3ny 
```

```go
var x int = 10 // memory address (0x001)
var p *int = &x // memory address (0x002) which contain the value (0x001)

 fmt.Println(x, p)
 fmt.Println(&x, &p)
 fmt.Println(*p) // Defrerencing

 *p = 20 (0x001 => 20)
  fmt.Println(x, *p)

```

output:

```sh
10 0xc00001a0b8
0xc00001a0b8 0xc000006028
10
20 20
```

**note**: printing a pointer that doesn't contain a value would result in a `nil` value
So `nil` is a predeclared identifier in Go that represents zero value for pointers, interfaces, channels, maps, and slices.

**another way to declare pointers**
using `new` keyword

```go
func main(){
    p := new(int)
    *p = 100
    fmt.Println("value", *p)
 fmt.Println("address", p)
}
```

**Pointers as function argument**

```go
myFunction(&a)

func myFunction(ptr *int){
    ... your code
}
```

**pointer to a pointer**

```go
func main(){
    p := new(int)
    *p = 100

    p1 := &p
    fmt.Println("P value", *p, " address", p)
 fmt.Println("P1 value", *p1, " address", p)

  fmt.Println("Dereferenced value", **p1)

}
```

```sh
$ go run main.go
P value 100  address 0xc0000be000
P1 value 0xc0000be000  address 0xc0000be000
Dereferenced value 100
```

**note**:
Go doesn't support pointer arithemtic like c/c++, however we can compare two pointers `===`

```go
p := &a
p1 := &a

fmt.Println(p == p1)
```

Pointers help us pass a huge amount of data without any effort, and change those data effiecnlty, it will help us in many situation we need to modify data from anothe pacakge or module

one of the situation it helps us to manage very effiecnlty, when we need to process large arrays
imagine you've an array of 1 million integers and you need to pass this array to a function, you know that array is passed by values, this means Go will copies allocation about 8MB of memory and then fill the stack with the million variable, this operation is costly is terms of performance and memory management
we can handle this situation using pointers

```go
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

Slices take care of this situation as it used pointers behind the scene
