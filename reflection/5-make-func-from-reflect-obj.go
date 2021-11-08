package main

import (
	"fmt"
	"reflect"
	"time"
)

//使用反射创建函数
func makeTimeFunc(f interface{}) interface{} {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)

	if tf.Kind() != reflect.Func {
		panic("expect a function")
	}

	wrapper := reflect.MakeFunc(tf, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		results = vf.Call(args)
		end := time.Now()
		fmt.Printf("the function takes %v\n", end.Sub(start))
		return results
	})
	return wrapper.Interface()
}

func TimeMe() {
	time.Sleep(1 * time.Second)
}

func main() {
	timedFunc := makeTimeFunc(TimeMe).(func())
	timedFunc()
}
