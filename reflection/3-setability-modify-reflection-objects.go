package main

import (
	"fmt"
	"reflect"
)

func main() {
	fVar := 3.14
	v := reflect.ValueOf(fVar)
	//v.CanAddr()
	fmt.Printf("is float canSet: %v,canAddr:%v\n", v.CanSet(), v.CanAddr())
	vP := reflect.ValueOf(&fVar)
	fmt.Printf("is float canSet: %v,canAddr:%v\n", vP.Elem().CanSet(), vP.Elem().CanAddr())
	//在使用set之前一定要先用canSet和canAddr检测下是否符合条件
	vP.Elem().SetFloat(2.333)
	println(fVar)
}
