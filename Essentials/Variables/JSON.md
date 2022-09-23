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
 HairColor string `json:"hair_color"`
 HasDog    bool   `json:"has_dog"`
}

func main() {
 myJson := `
 [
  {
   "first_name": "yousef",
   "last_name": "meska",
   "hair_color": "black",
   "has_dog": true
  },
  {
   "first_name": "omar",
   "last_name": "meska",
   "hair_color": "whirte",
   "has_dog": false
  }
 ]
 `

 // write json into struct
 var unmarshalled []Person
 err := json.Unmarshal([]byte(myJson), &unmarshalled)
 if err != nil {
  log.Println("Error unmarshalling json", err)
 }

 log.Printf("unmarshalled: %v", unmarshalled)

 // write json from a struct
 var mySlice []Person
 var m1 Person
 m1.FirstName = "Wally"
 m1.LastName = "West"
 m1.HairColor = "red"
 m1.HasDog = false
 mySlice = append(mySlice, m1)

 var m2 Person
 m2.FirstName = "yousef"
 m2.LastName = "meska"
 m2.HairColor = "yellow"
 m2.HasDog = true
 mySlice = append(mySlice, m2)

 newJson, err := json.MarshalIndent(mySlice, "", " ")
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
