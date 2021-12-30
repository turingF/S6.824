package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
  @Description: 多线程修改变量时，可以使用atomic包来原子性地安全修改
*/

var (
	shutdown int64
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("hello")
	go doWork("world")

	time.Sleep(1 * time.Second)

	atomic.StoreInt64(&shutdown,1)
	wg.Wait()

}

func doWork(name string) {
	defer wg.Done()

	for  {
		fmt.Printf("Doing %s work\n",name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1{
			fmt.Printf("Shut down %s work\n",name)
			break
		}
	}
}