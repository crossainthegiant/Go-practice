package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//给定一些任务,要求在规定的timeout内跑完。
//如果时间截止之前任务未完成报告timeout
//如果操作系统interrupt中断信号，报一个错误

type Runner struct {
	interruptChan chan os.Signal
	completeChan chan error
	timeoutChan <-chan time.Time//用来计时
	tasks []func(int)//task的列表
}

var ErrTimeout = errors.New("this is a timeout")
var ErrInterrupt = errors.New("this is a interrupt")

func New(t time.Duration) *Runner  {
	return &Runner{
		interruptChan: make(chan os.Signal, 1),
		completeChan:  make(chan error),
		timeoutChan:   time.After(t),
		tasks:         make([]func(int), 0),
	}
}

func (r *Runner)AddTask(tasks ...func(int))  {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner)Run() error {
	for id, task := range r.tasks{
		select {
		case <-r.interruptChan:
			signal.Stop(r.interruptChan)
			return ErrInterrupt
		default:
			task(id)
		}
	}
	return nil
}

func (r *Runner)Start() error {
	// relay interrupt from os
	signal.Notify(r.interruptChan, os.Interrupt)

	//run the tasks
	go func() {
		r.completeChan <- r.Run()
	}()

	select {
	case err := <- r.completeChan:
		return err
	case <-r.timeoutChan:
		return ErrTimeout

	}


}