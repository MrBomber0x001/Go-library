
### Looping

```go
package main

import "log"

type User struct {
 FirstName string
 LastName  string
 Age       int
}

func main() {

 for i := 0; i <= 10; i++ {
  log.Println(i)
 }
 //randing over slices
 mySlice := []string{"dog", "cat", "snake", "fish"}
 for _, x := range mySlice {
  log.Println(x)
 }
 //randing over maps
 log.Println("==================")
 myMap := make(map[string]string)
 myMap["first"] = "dog"
 myMap["second"] = "fish"
 myMap["third"] = "snake"
 myMap["fourth"] = "cat"

 for i, y := range myMap {
  log.Println(i, y)
 }

 var myStructSlice []User

 u1 := User{
  FirstName: "Yousef",
  LastName:  "Meska",
  Age:       22,
 }
 u2 := User{
  FirstName: "Omar",
  LastName:  "Meska",
  Age:       22,
 }
 myStructSlice = append(myStructSlice, u1, u2)
 log.Println("==================")

 for i, x := range myStructSlice {
  log.Println(i, x.Age)
 }

}

// one has loop type {for loop} only

/*
The range keyword is used in for loop to iterate over items of an array, slice, channel or map. With array and slices, it returns the index of the item as integer. With maps, it returns the key of the next key-value pair
*/
```
