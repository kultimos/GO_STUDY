package main

import (
	_ "GO_STUDY/study/import_init/init_package" // 匿名导包,这样导进来的包我们无法去使用其内部的方法,但是会在我们的main方法执行时调用其init方法,奇怪的需求
	//"GO_STUDY/study/import_init/init_package"
	// MyMethod "GO_STUDY/study/import_init/init_package" //相当于给导进来的这个包起了一个别名,这种方式感觉是最好的,因为有的包的包名可能非常长和复杂
	//. "GO_STUDY/study/import_init/init_package" //这样导进来的包里的方法可以直接用
	"fmt"
)

func main() {
	//a := init_package.Method()
	//a := MyMethod.Method() //通过别名去调用其方法
	//a := Method() //通过.的方式导入的包,可以直接调用其方法,但我觉得这种写法的可读性很差
	//fmt.Println("调用其他人的方法返回的结果是: ", a)
}

func init() {
	fmt.Println("我自己的init初始化")
}

/**
  可以看到大概得调用顺序,是先调用了init_package中forInit中的init方法,然后调用自己的init方法
  在其他包中而非main包中定义的方法的大小写是有讲究的,简单说就是大写开头的方法是public的,而小写的是private的
*/
