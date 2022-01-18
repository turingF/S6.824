package main

import (
	"github.com/code/gotest2/work"
	"log"
	"sync"
	"time"
)

/*
  @Description:
*/

var names = []string {
	"foo",
	"bar",
	"hello",
	"world",
	"xy",
}

type namePrinter struct {
	name string
}


func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main(){
	p := work.New(20)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i:=0; i< 100 ; i++ {
		for _,name := range names{
			np := namePrinter{
				name : name,
			}
			go func(){
				// namePrinter complete Task(), so it implements [Worker] interface
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.ShutDown()
}