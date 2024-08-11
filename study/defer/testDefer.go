package main

import "fmt"

/**
  defer虽然写在func方法体内,但实际上defer的运行是在func的右括号结束后才去执行的,所以可以从total()中看出,实际上defer修饰的内容是在return之
后才运行的
  并且多个defer修饰的内容类似于入栈出栈的关系,defer1最先进栈,所以defer1反而是最后执行的;而最后入栈的defer3却会先被执行;
*/

func main() {
	defer defer1()
	defer defer2()
	defer defer3()
	fmt.Println("虽然我在defer后面写")
	fmt.Println("但是我在defer前面执行")

	fmt.Println("----------------------------------------")
	total()
}

func total() (a int) {
	defer defer1()
	return defer2()
}

func defer1() (a int) {
	fmt.Println("我是defer1")
	return 1
}

func defer2() (a int) {
	fmt.Println("我是defer2")
	return 2
}

func defer3() {
	fmt.Println("我是defer3")
}
