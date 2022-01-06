package main

import (
	"github.com/code/gotest2/runner"
	"log"
	"os"
	"time"
)

/*
  @Description:
*/

const timeout = 3 * time.Second

func main() {
	log.Println("Start working")

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	// if func define var
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Stop cause [timeout]")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Stop cause [interrupt]")
			os.Exit(2)
		}
		// 正常情况下会以错误码0退出
	}

	log.Println("Main Process End !")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Start start task [%d]", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
