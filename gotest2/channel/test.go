package channel

import (
	"fmt"
	"sync"
	"time"
)

/*
  @Description: 使用channel来进行线程通讯
*/

var(
	wg sync.WaitGroup
)

func main(){
	wg.Add(1)
	// 通道分为缓冲和无缓冲两个类型
	buffer := make(chan string,10)
	//unbuffer := make(chan string)

	// buffer 作为参数传递给协程
	go printBuffer(buffer)
	time.Sleep(2000 * time.Millisecond)
	buffer <- "hello"
	wg.Wait()
}

func printBuffer(b chan string){
	defer wg.Done()
	value := <- b
	fmt.Printf("Buffer: %s \n",value)
}