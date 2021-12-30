package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
  @Description:
*/

func main() {
	// alloc a processor to scheduler
	runtime.GOMAXPROCS(runtime.NumCPU())

	// same as wordcount
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	go func (){
		defer wg.Done()
		for count :=0; count <3 ; count++ {
			for char:= 'A';char < 'A' + 26; char++ {
				fmt.Printf("%c ", char);
				time.Sleep(100 * time.Millisecond)
			}
			println()
		}
	}()

	go func (){
		defer wg.Done()
		for count :=0; count <3 ; count++ {
			for char:= 'a';char < 'a' + 26; char++ {
				fmt.Printf("%c ", char);
				time.Sleep(100 * time.Millisecond)
			}
			println()
		}
	}()

	fmt.Println("Waiting to Fininsh")
	wg.Wait()

	fmt.Println("Ending the program!")

}
