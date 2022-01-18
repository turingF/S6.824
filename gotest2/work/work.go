package work

import "sync"

/*
  @Description: worker 并发模式
*/

// Worker offer
type Worker interface {
	Task()
}

// Pool offers a pool
type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

func New(maxCnts int) *Pool{
	p := Pool {
		work: make(chan Worker),
	}
	p.wg.Add(maxCnts)

	for i := 0; i< maxCnts ; i++ {
		go func(){
			// range will block until channel has a value
			for w:= range p.work{
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) Run(w Worker){
	p.work <- w
}

func (p *Pool) ShutDown() {
	close(p.work)
	p.wg.Wait()
}