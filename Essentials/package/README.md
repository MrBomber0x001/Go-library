anyfile.go

we need first to make a module

```sh
go mod init <module_name>
go mod init meska
```

```go
// package main

// import (
//  "log"
//  "meska/helpers"
// )

// func main() {

//  var myVar helpers.SomeType
//  myVar.TypeName = "yousef"
//  log.Println(myVar)
// }

// // it's conventional to name modules with the github repo name
// // go mod init <name>

```

helpers/helpers.go

```go
package helpers

type SomeType struct {
 TypeName   string
 TypeNumber int
}
```
