package main

import (
	"runtime"
	"sync"
)

/*
  @Description:
*/

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go addCount(1)
	go addCount(2)

	wg.Wait()
	println(counter)

}

func addCount(c int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// give up cur thread and put back to list
		runtime.Gosched()
		value++
		counter = value
	}
}
