package main

import (
	"runtime"
	"sync"
)

/*
  @Description: 使用互斥锁来保证并发正确性
*/

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go count()
	go count()

	wg.Wait()
	println(counter)
}

func count() {
	defer wg.Done()

	for i := 0; i < 1000; i++ {

		//value := counter
		//value++
		//counter = value

		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}
