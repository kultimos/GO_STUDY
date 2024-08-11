package main

import "fmt"

func main() {
	a := 5
	changeValue(a)
	fmt.Println("a = ", a)
	changePointValue(&a)
	fmt.Println("a = ", a)
}

func changeValue(p int) {
	p = 100
}

func changePointValue(p *int) {
	*p = 999
}

/**
  这里的点在于,通常我们的写法中实际上是一个值传递,传递进去的不是对象的地址,所以我们无论在方法中如何修改,不影响外部穿进去的参数
  而通过 *和& 明确了传递的是地址以后,自然对内存地址上的数据改变,会影响到入参本身
  所以golang语言本身是一个值传递的语言?通过指针的方式实现引用传递？
*/
