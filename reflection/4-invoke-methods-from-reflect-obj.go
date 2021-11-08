package main

import (
	"fmt"
	"reflect"
)

//使用reflection调用一个Obj的方法

type Student struct {
	name string
}

func (s *Student) DoHomework(number int) {
	fmt.Printf("%s is doing homework\n", s.name, number)
}

func main() {
	//user reflection to invoke the DoHomework of a student
	s := Student{name: "lilei"}
	v := reflect.ValueOf(&s)
	methodv := v.MethodByName("DoHomework")
	if methodv.IsValid() {
		methodv.Call([]reflect.Value{reflect.ValueOf(3)})
	}
}
