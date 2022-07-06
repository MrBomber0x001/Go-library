
## Modules

modules: collection of go packages stored in a file tree with `go.mod` file as its root
`GOPATH` variable that defines the root of your workspace and it contains the following folders:

- `src`: Go source code
- `pkg`: compiled package code
- `bin`: compiled binaries and executables

<img src="https://www.karanpratapsingh.com/_next/image?url=%2Fstatic%2Fcourses%2Fgo%2Fchapter-I%2Fmodules%2Fgopath.png&w=3840&q=75" />

```sh
go mod init example
```

a module can correspond to `Github` repo if you plan to publish it

```sh
go mod init github.example.com/user/pkg/v1
```

-- install dependecies

```sh
go get github.com/rs/zerolog # or
go install github.com/rs/zerolog
```

`go.sum` file contains expected hashes of the content of the new modules
-- list the dependecies and update them

```sh
go list -m -u all
```

-- removing dependecies that are not used, or install the used ones

```sh
go mod tidy
```

#### Vendoring

Vendoring is the act of making your own copy of the 3rd party packages your project is using. Those copies are traditionally placed inside each project and then saved in the project repository.

```go
package main
import "github.com/rs/zerolog/log"

func main(){
    log.Info().Msg("Hello")
}
```

there will be folder newly created titled "vendor"

### Workspaces (multi-module)

workspaces allow use to work with multiple modules simulatensouly without having to edit go.mod for each module, each module within a workspace is treated as a root module when resolving dependecies.

```sh
mkdir workspaces && cd workspaces
mkdir hello && cd hello
go mod init hello
```

```bash
$ go get golang.org/x/example
go: downloading golang.org/x/example v0.0.0-20220412213650-2e68773dfca0
go: added golang.org/x/example v0.0.0-20220412213650-2e68773dfca0
```

```go
package main

import (
 "fmt"

 "golang.org/x/example/stringutil"
)

func main() {
 result := stringutil.Reverse("Hello Workspace")
 fmt.Println(result)
}
```

```bash
$ go run main.go
ecapskroW olleH
```

But if we want to modify `stringutil` module that our code depends on ?
Until now, we had to do it using the `replace` directive in the go.mod file, but now let's see how we can use workspaces here.

```bash
cd ..
pwd # workspaces
$ go work init
$ cat go.work
    go 1.18
$ go work use ./hello
```

now let's download `stringutil` and update the `Reverse` function

```sh
$ git clone https://go.googlesource.com/example
Cloning into 'example'...
remote: Total 204 (delta 39), reused 204 (delta 39)
Receiving objects: 100% (204/204), 467.53 KiB | 363.00 KiB/s, done.
Resolving deltas: 100% (39/39), done.
```

head over to `example/stringutil/reverse.go`

```go
func Reverse(s string) string {
 return fmt.Sprintf("I can do whatever!! %s", s)
}
```

now head over to the parent directory `workspaces`

```sh
$ go work use ./example
$ cat go.work
go 1.18

use (
 ./example
 ./hello
)
```

now if we head over to `hello` module and run `main.go`

```sh
$ go run main.go
I can do whatever!! Hello Workspace
```

### Packages

A package is nothing but a directory containing one or more Go source files, or other Go packages.

This means every Go source file must belong to a package and package declaration is done at top of every source file as follows.

So far, we've done everything inside of package main. By convention, executable programs (by that I mean the ones with the main package) are called Commands, others are simply called Packages.

let's work from a clean spot

```sh
mkdir hello && cd hello
touch main.go
go mod init hello
mkdir custom
cd custom && touch code.go
```

```go
// hello/custom/code.go
package custom
var value int = 10 // will not be exported
var Value int = 20 // will be exported 
--------------------------------
//hello/main.go
package main 
    import (
        "fmt"
        myNewAlis "hello/custom"
    )
func main(){
   fmt.Println(myNewAlias.Value)
}
```

### Versioning

go.mod structure

```go
module <name>
go <version>
require <module-path> <semanitc-version>
replace <module-path> <semanitc-version>
exclude <module-path> <semanitc-version>
```

--> Updgrading and downgrading

- we can specify a commit hash
- also a semanitc-version

```sh
go get <url>@commit-hash
```

--> Updating

```sh
go get -u <url>
```

---> Versioning
there are two main ways of versioning

1. using module-path and v2 folders

```sh
parent directory:
    - go.mod
    - ...
    - v2:
        - go.mod
        - ...
```

2. using branching

```sh
main branch:
    - go.mod
    - ...
v1/branch:
    - go.mod
    - ....
v2/branch:
    - go.mod
    - ...
```

Lastly, I will add that, Go doesn't have a particular "folder structure" convention, always try to organize your packages in a simple and intuitive way.
