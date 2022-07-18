```go
package main

import (
 "errors"
 "fmt"
 "net/http"
 "text/template"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(res http.ResponseWriter, req *http.Request) {
 renderTemplate(res, "home.page.html")
}

// About is the about page handler
func About(res http.ResponseWriter, req *http.Request) {
 sum, _ := addValues(2, 0)
 _, _ = fmt.Fprintf(res, fmt.Sprintf("The output is %d", sum))
}
func addValues(x, y int) (int, error) {
 return x + y, nil
}
func Divide(res http.ResponseWriter, req *http.Request) {
 result, err := divide2number(1, 0)
 if err != nil {
  fmt.Fprintf(res, err.Error())
  return
 }
 _, _ = fmt.Fprintf(res, fmt.Sprintf("The result of division is %f", result))
}

func divide2number(x, y float32) (float32, error) {
 if y <= 0 {
  err := errors.New("Cannot divide over 0")
  return 0, err
 }
 result := x / y
 return result, nil
}

// main is the entry point
func main() {
 http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
  n, err := fmt.Fprintf(res, "Hello world")
  fmt.Println("Bytes Written:", n) //TODO: see if it works!
  fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
  if err != nil {
   fmt.Println(err)
  }
 })
 http.HandleFunc("/home", Home)
 http.HandleFunc("/about", About)
 http.HandleFunc("/divide", Divide)

 fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

 _ = http.ListenAndServe(portNumber, nil)
}

// request is a pointer
// Fprintf (requires a response write to write to) and it returns number of bytes written and error
// ====
// res.send("welcome")

// Sprintf() allow you to take different data types and return them as a string

// Parse Templates
func renderTemplate(res http.ResponseWriter, tmpl string) {
 parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
 err := parsedTemplate.Execute(res, nil)
 if err != nil {
  fmt.Println("Error parsing template: ", err)
  return
 }
}

```

## Rendering Template and Template Caching

```go
package render

import (
 "bytes"
 "fmt"
 "html/template"
 "log"
 "net/http"
 "path/filepath"
)

var functions = template.FuncMap{}

/**
* RenderTemplate renders template for fn handlers
* Procedure:
  - Read Template from the disk
  - Parse the template
  - return it
*
*/

func RenderTemplate(res http.ResponseWriter, tmpl string) {
 // Just to show Template caching is working
 // _, err := RenderTemplateTest(res)

 // if err != nil {
 //  fmt.Println("Error getting template cache", err)
 // }

 tc, err := CreateTemplateCache()
 if err != nil {
  log.Fatal(err)
 }

 // get the template cache from app config
 // ok is used (any second variable) with map to check if there's actually value in the map for
 // the corresponding key (return true) and (return false) in case of not found [ convernsion]
 t, ok := tc[tmpl]

 if !ok {
  log.Fatal(err)
 }

 buf := new(bytes.Buffer) // put the template into some bytes [buffer]

 _ = t.Execute(buf, nil) //take that t (template) and store its content into the buffer, nil (for not passign any data)
 _, err = buf.WriteTo(res)
 if err != nil {
  fmt.Println("Error Writing template to browser", err)
 }
 // parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
 // err = parsedTemplate.Execute(res, nil)
 // if err != nil {
 //  fmt.Println("error parsing template", err)
 //  return
 // }
}

/**
* CreateTemplateCache: Parse all of the templates and store them into a map of templates "cache"
*
**/
func CreateTemplateCache() (map[string]*template.Template, error) {
 myCache := map[string]*template.Template{} // creating a template cache that contains templates

 pages, err := filepath.Glob("./templates/*.page.tmpl") // get all the pages in that folder
 if err != nil {
  return myCache, err
 }

 for _, page := range pages {
  name := filepath.Base(page) // to extract just the name of the page, not the full path
  fmt.Println("Page is currently", page)
  ts, err := template.New(name).Funcs(functions).ParseFiles(page)

  if err != nil {
   return myCache, err
  }

  matches, err := filepath.Glob("./templates/*.layout.tmpl")
  if err != nil {
   return myCache, err
  }
  if len(matches) > 0 {
   // parse those templates
   ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
   if err != nil {
    return myCache, err
   }
  }
  myCache[name] = ts
 }
 return myCache, nil
}

// [about.page.tmpl] => full parsed and ready to use templates
// [home.page.tmpl] =>  full parsed and ready to use templates
```

In this kind of implementation, we have a major problem, which is every time we hit the server, the CreateTemplateCache() is going to be executed and travsering all the templates folder and this is not an effiecent way to do.

we need a way to fire this only once the server is set up.
in the main function
we'll be creating application wide configurations

```go
package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
 TemplateCache map[string]*template.Template
}


```

main

```go
package main

import (
 "fmt"
 "log"
 "net/http"

 "github.com/MrBomber0x001/sample/pkg/config"
 handler "github.com/MrBomber0x001/sample/pkg/handlers"
 "github.com/MrBomber0x001/sample/pkg/render"
)

const portnumber = ":8080"

func main() {
 var app config.AppConfig

 tc, err := render.CreateTemplateCache()
 if err != nil {
  log.Fatal("Cannot create template cache")
 }

 app.TemplateCache = tc
 http.HandleFunc("/", handler.Home)
 http.HandleFunc("/about", handler.About)
 fmt.Println(fmt.Sprintf("Firing on port %s", portnumber))

 _ = http.ListenAndServe(portnumber, nil)

}
```
