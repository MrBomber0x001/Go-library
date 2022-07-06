
## Trivial Example

```go
type mobile struct {
 brand string
}

type laptop struct {
 cpu string
}

type toaster struct {
 amount int
}

type kettle struct {
 quantity string
}
type socket struct{}

```

```go
func (m mobile) Draw(power int){
    fmt.Printf("%T --> brand: %s, power: %d", m, m.brand, power)
}

```

```go
func (socket) Plug(device mobile, power int){

}
