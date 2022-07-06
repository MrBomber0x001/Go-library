
## Working with JSON

1. consistent structure and pre knowing one
2. inconsisent structure or you don't know the structure

```go
package main

import (
 "encoding/json"
 "fmt"
 "log"
)

type Person struct {
 FirstName string `json:"first_name"`
 LastName  string `json:"last_name"`
 HasDog    bool   `json:"has_dog"`
}

func main() {
 myJson := `
 [
  {
   "first_name": "Yousf",
   "last_name": "meska",
   "has_dog": true
  },
  {
   "first_name": "Ahmed",
   "last_name": "Meska",
   "has_dog": false
  }
 ]
 `
 // write json to a struct
 var unmarshalled []Person
 // Unmarshall([]bytes date, ptr *v)
 //[]bytes(myJson) to convert string to bytes
 err := json.Unmarshal([]byte(myJson), &unmarshalled)
 if err != nil {
  log.Printf("error unmarshalling")
 }
 log.Printf("unmarshalled: %v", unmarshalled)

 // write json from a struct
 var mySlice []Person

 var marshall1 Person
 marshall1.FirstName = "Wally"
 marshall1.LastName = "werst"
 marshall1.HasDog = true

 mySlice = append(mySlice, marshall1)

 newJson, err := json.MarshalIndent(mySlice, "", "     ") // gives us a slice of bytes
 if err != nil {
  log.Println("error marshalling", err)
 }
 fmt.Println(string(newJson))
}


// For development purpose use MarsharIndex, for production use Marshal

// this topic was about writing and reading json with known format, but what if i don't know
// there's another way to deal with such situation
//TODO: search for it
```
