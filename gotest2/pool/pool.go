package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

/*
  @Description: 管理多个线程间的共享和独立资源
*/

type Pool struct {
	m        sync.Mutex
	resource chan io.Closer
	factory  func() (io.Closer, error)
	closed   bool
}

var ErrPoolClosed = errors.New("Pool has been closed")

// 创建一个对应缓冲区大小的资源池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size is too small")
	}

	return &Pool{
		factory:  fn,
		resource: make(chan io.Closer, size),
	}, nil
}

// 获取一个资源池的资源，如果没有则new一个新资源返回
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resource:
		log.Println("Acquire", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	default:
		log.Println("Acquire", "New Resource")
		return p.factory()
	}
}

// 将一个资源放到池内
func (p *Pool) Release(r io.Closer) {
	// lock all func body
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}
	select {
	// put param to queue
	case p.resource <- r:
		log.Println("Release:", "In Queue")
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// 关闭资源池
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	// close channel
	close(p.resource)

	for r := range p.resource {
		r.Close()
	}
}
