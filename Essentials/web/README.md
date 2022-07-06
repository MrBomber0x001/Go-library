TODO: why request is a pointer?

## http request and error handling

- put your handlers in one places
handlers:
  - home.go
  - about.go

```go
package main

import (
 "errors"
 "fmt"
 "net/http"
)

// CONFIGURATIONS

const portNumber = ":8080"

func addValue(x, y int) int {
 return x + y
}

func divideValues(x, y float32) (float32, error) {
 if y <= 0 {
  err := errors.New("Cannot divide by zero")
  return 0, err
 }
 result := x / y

 return result, nil
}
func main() {
 // if you are about to see a new method or interface look for it's documentation
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("welcome"))
  n, err := fmt.Fprintf(w, "hello world")
  if err != nil {
   fmt.Println(err)
  }
  fmt.Println(fmt.Sprintf("Number of bytes written %d", n))
 })
 http.HandleFunc("/home", Home)
 http.HandleFunc("/about", About)
 fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
 err := http.ListenAndServe(portNumber, nil)
 if err != nil {
  fmt.Println("Error starting the server!")
 }
}

// ROUTERS

func Home(w http.ResponseWriter, r *http.Request) {
 fmt.Println("Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
 sum := addValue(2, 2)
 _, _ = fmt.Fprintf(w, fmt.Sprintf("this is the about page and 2 + 2 is %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
 result, err := divideValues(100.0, 10.0)
 if err != nil {
  fmt.Fprintf(w, err.Error())
 }
 fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, result))
}

```

## Rendering Templates

using a built-in function called `template`

```go
func renderTemplate(w http.ResponseWriter, tmpl string) {
    parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
    err := parsedTemplate.Execute(w, nil) // "nil" for passing no data to the tempalate
    if err != nil {
        fmt.Printf("Error parsing template", err)
        return
    }
}

func Home(w http.ResponseWriter, r *http.Request){
    renderTemplate(w, "home.html")
}

```

creating `template cache`

## Session Management

## Configuration management
