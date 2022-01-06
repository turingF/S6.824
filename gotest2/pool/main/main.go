package main

import (
	"github.com/code/gotest2/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
  @Description:
*/

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct{
	ID int32
}

// dbConnection 实现了io.Closer的接口方法，默认继承了io.Closer
func (dbConn *dbConnection) Close() error {
	log.Println("Close connection",dbConn.ID)
	return nil
}

var idCounter int32
func createConnection() (io.Closer,error){
	id := atomic.AddInt32(&idCounter,1)
	log.Println("Create: New Connection ",id)

	return &dbConnection{id},nil
}

func main(){
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 以函数名为参数 传递给New方法
	p,err := pool.New(createConnection,pooledResources)
	if err != nil {
		log.Print(err)
		return
	}

	for query := 0;query<maxGoroutines;query++ {
		go func(q int) {
			performQueries(q,p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("ShutDown Program")
	p.Close()
}

func performQueries(query int,p *pool.Pool){
	conn,err := p.Acquire()
	defer p.Release(conn)
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("--- Query:[%d] Connection:[%d] \n",query,conn.(*dbConnection).ID)
}