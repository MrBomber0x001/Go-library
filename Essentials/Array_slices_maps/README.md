# table of content

### Intro

```go
package main

import (
 "log"
 "sort"
)

type User struct {
 FirstName string
 LastName  string
}

func main() {
 myMap := make(map[string]string)
 myOtherMap := make(map[string]int)
 myMap["person"] = "yousef"
 myMap["another_person"] = "ahmed"
 myOtherMap["first"] = 1
 myOtherMap["second"] = 2
 log.Println(myMap["person"], myMap["another_person"])
 myStructMap := make(map[string]User)
 me := User{
  FirstName: "Yousef",
  LastName:  "meska",
 }
 myStructMap["me"] = me
 log.Println(myStructMap["me"].FirstName)

 // if you don't know what type you will put in the map, do it
 // myMapp := make(map[string]interface{})

 //** Slices
 var myStringSlices []string
 var myIntSlices []int
 myStringSlices = append(myStringSlices, "Trevor", "Meska")
 myStringSlices = append(myStringSlices, "Mary")
 myIntSlices = append(myIntSlices, 2, 3, 4, 5, 1)
 sort.Ints(myIntSlices)
 log.Println(myIntSlices)
 log.Println(myStringSlices)

 numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
 log.Println(numbers)

 log.Println(numbers[0:2])

 names := []string{"one", "seven", "three"}
 log.Println(names)
}

// map[index_to_lookup_the_map] type_value_to_store

// values in a map are randomized, if you put them in order, they are not actually stored in order

// slices are ordered in the wy you put the value

//TODO: there some notes about maps, yet!
```
