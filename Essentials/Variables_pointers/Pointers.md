
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
