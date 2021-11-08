package main

import (
	"fmt"
	"reflect"
)

type User struct {
	name string
	age  int
}

func (u User) PrintName() {
	fmt.Println(u.name)
}

func (u User) PrintAge() {
	fmt.Println(u.age)
}

type Aggregator func(int, int) int

var (
	user = User{
		name: "lilei",
		age:  24,
	}

	add Aggregator = func(a int, b int) int {
		return a + b
	}

	sub Aggregator = func(a int, b int) int {
		return a - b
	}
)

func inspect(variable interface{}) {
	t := reflect.TypeOf(variable)
	v := reflect.ValueOf(variable)

	if t.Kind() == reflect.Struct {
		fmt.Printf("--------------------------fields %d ----------------------\n", t.NumField())
		for idx := 0; idx < t.NumField(); idx++ {
			fieldType := t.Field(idx)
			fieldVal := v.Field(idx)
			fmt.Printf("\t %v = %v\n", fieldType, fieldVal)
		}

		//literate its methods
		fmt.Printf("--------------------------methods %d ----------------------\n", t.NumMethod())
		for idx := 0; idx < t.NumMethod(); idx++ {
			methodType := t.Method(idx)
			fmt.Printf("\t input_num=%d, output_num=%d\n", methodType.Type.NumIn(), methodType.Type.NumOut())
		}
	} else if t.Kind() == reflect.Func {
		fmt.Printf("\t this function has  %d inputs and %d outputs", t.NumIn(), t.NumOut())
	}
	fmt.Printf("\n\n")
}

func main() {
	inspect(user)
	inspect(add)
	inspect(sub)
}
