package main

import "context"

func main() {
	//创建空context
	ctx := context.Background() //新建一个空的context，它没有deadline和cancel，也不能被取消，kv为空

	todoCTX := context.TODO() //返回空的context
}
