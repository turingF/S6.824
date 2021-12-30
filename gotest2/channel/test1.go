package channel

import (
	"fmt"
	"math/rand"
	"time"
)

/*
  @Description: 使用无缓冲通道来模拟两个线程互传数据（打球）
*/

//var wg sync.WaitGroup

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	wg.Add(2)

	go player("XY",court)
	go player("ABC",court)

	// start
	court <- 1
	wg.Wait()
}

func player(name string,channel chan int){
	defer wg.Done()

	// for 默认是while true
	for  {
		ball,ok := <- channel
		if !ok {
			fmt.Printf("Player %s won\n",name)
			return
		}

		p := rand.Intn(100)
		if p%13 == 0 {
			fmt.Printf("Player %s Missing ball\n",name)
			close(channel)
			return
		}

		fmt.Printf("Player %s catched ball, hit %d\n",name,ball)
		ball++
		time.Sleep(500 * time.Millisecond)
		channel <- ball

	}

}