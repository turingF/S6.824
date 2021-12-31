package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

/*
  @Description: 使用Runner包来监控程序
*/

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")

// 使用New命名，类似于构造方法
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1), /*有缓冲通道，保证os不会阻塞*/
		complete:  make(chan error), /*无缓冲通道，task线程在main接受error后，task线程安全退出*/
		timeout:   time.After(d), /*after 会在一段时间d后返回一个time.Time的数据*/
	}
}

func (r *Runner /*使用引用来操作调用对象本身*/) Add(task ...func(int) /*... 代表可变参数长度*/) {
	r.tasks = append(r.tasks, task...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}() // 函数后面要加个()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}

}

// 执行注册任务，返回异常
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// task 是个函数 func(int)
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
