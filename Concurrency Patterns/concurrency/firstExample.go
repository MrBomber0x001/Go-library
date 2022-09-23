// This program show how the go runtime schedulers work by displaying the alphabets 3 time concurrenctly
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // notice the difference in result when changing the number of logical processor
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Starting Goroutines")

	go func() {
		defer wg.Done()
		for counter := 0; counter < 3; counter++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for counter := 0; counter < 3; counter++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println(" ")
	}()

	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTermainting Program")

}
