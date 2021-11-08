package main

import (
	"fmt"
	"reflect"
)

func printMeta(obj interface{}) {
	t := reflect.TypeOf(obj) //Type
	n := t.Name()
	v := reflect.ValueOf(obj) //Value
	k := t.Kind()
	fmt.Printf("type:%s, name:%s,value:%s,kind:%s\n", t, n, v, k)
}

type handler func(int, int) int

func main() {
	//Go中，每一个接口变量的实现都对应一个pair：<value,type> 记录着实际变量的值和它的类型，reflection就是用来检测类型和值的
	//提供了两种方法，reflect.valueof, reflect.typeof
	var intVar int64 = 10
	stringVar := "hello"
	type book struct {
		name  string
		pages int
	}
	sum := func(a, b int) int {
		return a + b
	}
	var sub handler = func(a int, b int) int {
		return a - b
	}
	sl := make([]int, 3)
	printMeta(intVar)
	printMeta(stringVar)
	printMeta(book{
		name:  "harry potter",
		pages: 500,
	})
	printMeta(sum)
	printMeta(sub)
	printMeta(sl)
	//type:int64, name:int64,value:%!s(int64=10),kind:int64
	//type:string, name:string,value:hello,kind:string
	//type:main.book, name:book,value:{harry potter %!s(int=500)},kind:struct
	//type:func(int, int) int, name:,value:%!s(func(int, int) int=0x4c96a0),kind:func
	//type:main.handler, name:handler,value:%!s(main.handler=0x4c96c0),kind:func
	//type:[]int, name:,value:[%!s(int=0) %!s(int=0) %!s(int=0)],kind:slice

}
