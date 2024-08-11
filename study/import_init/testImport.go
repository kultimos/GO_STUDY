package main

import (
	"GO_STUDY/study/import_init/init_package"
	"fmt"
)

func main() {
	a := init_package.Method()
	fmt.Println("调用其他人的方法返回的结果是: ", a)
}

func init() {
	fmt.Println("我自己的init初始化")
}

/**
  可以看到大概得调用顺序,是先调用了init_package中forInit中的init方法,然后调用自己的init方法;
  在其他包中而非main包中定义的方法的大小写是有讲究的,简单说就是大写开头的方法是public的,而小写的是private的
*/
