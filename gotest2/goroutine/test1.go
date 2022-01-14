package main

import (
	"fmt"
	"os"
	"os/signal"
)

/*
  @Description: catch os signal to go channel
*/

func main() {
	// init channel with buffer size:1
	interrupt := make(chan os.Signal,1)

	// detect os signal to channel
	signal.Notify(interrupt,os.Interrupt)

	select {
	// if channel has a elem,catch it
	case sign := <-interrupt :
		fmt.Println("Catch signal :",sign)
	}
}
