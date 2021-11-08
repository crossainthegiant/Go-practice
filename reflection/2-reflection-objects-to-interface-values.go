package main

import (
	"fmt"
	"reflect"
)

func main() {
	//first example showing convert reflect.value to float
	floatVar := 3.14
	v := reflect.ValueOf(floatVar)

	newFloat := v.Interface().(float64)
	fmt.Println(newFloat + 1.2)

	//second example showing convert reflect.value to slice
	sliceVar := make([]int, 5)
	v = reflect.ValueOf(sliceVar)
	v = reflect.Append(v, reflect.ValueOf(2))
	newSlice := v.Interface().([]int)
	newSlice = append(newSlice, 4)
	fmt.Println(newSlice)

}
