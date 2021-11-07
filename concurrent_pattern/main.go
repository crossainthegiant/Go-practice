package main

import (
	"fmt"
	"github.com/crossainthegiant/learningGo/concurrent_pattern/runner"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Println("task finished", id)
	}
}

func main() {

	r := runner.New(5 * time.Second)
	r.AddTask(createTask(), createTask(), createTask())
	err := r.Start()
	switch err {
	case runner.ErrInterrupt:
		fmt.Println("interrupt error")
	case runner.ErrTimeout:
		fmt.Println("Time out!")
	default:
		fmt.Println("all tasks completed")

	}

}
