package main

import "fmt"

/**
  第一行的iota的默认值是0
  注意iota的+1是以行为单位
  iota只能出现在const的定义的()之中
*/

const (
	a int = 100
	b     = iota
	c     = iota
	d     = iota
)

const (
	x, y = iota + 1, iota + 2
	j, k = iota * 2, iota * 3
)

func main() {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
}
