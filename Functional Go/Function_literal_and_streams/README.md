## function literal and streams

### function literal

unnamed function (anonymous) that is later bound into a variable when used

- can be used inside other functions
- has the advatages of using the parameters of the functions it is nested in
- comfortable and easy one-time usage (can't be used amon packages)

### stream

just like programming languges, streams are used for dealing with streams of data, modifiny them or generating new data based on a specific rule or conditions

- Not supported by default in Go TODO: check it for updates, in that moment of time (before 1.18) there was no Generics and streams were depending on them
- Popular: map, reduce, filter
- easy to do computations over slices or array

```go
// for example a Filter function which takes a slice and return only segement of data based on a condition
Filter(weighs, func(w int) bool {
    return w > 80
}
```
