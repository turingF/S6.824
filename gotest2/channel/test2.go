package main

import (
	"fmt"
	"sync"
	"time"
)

/*
  @Description: 通过无缓冲通道来模拟多个线程线性传递数据（接力棒）
*/

var(
	wg sync.WaitGroup
)

func main(){
	wg.Add(1)
	bang := make(chan int)
	go run(bang)

	bang <- 1
	wg.Wait()
}

func run(channel chan int){

	var newRound int
	round := <- channel

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Player %d start running...\n",round)
	time.Sleep(1000 * time.Millisecond)

	if round != 4{
		newRound = round + 1
		fmt.Printf("Player %d ready \n",newRound)
		go run(channel)
	}

	if round == 4 {
		fmt.Printf("Player %d touch the goal\n",round)
		wg.Done()
		return
	}

	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Player exchange %d to %d\n",round,newRound)
	channel <- newRound
	time.Sleep(200 * time.Millisecond)

	fmt.Printf("Player %d Stop !\n",round)

}