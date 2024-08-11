package main

import "fmt"

func main() {
	c := method1(25, 32)
	fmt.Println("c = ", c)

	fmt.Println("=========================================")

	c, d := method2("hah", 13)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)

	fmt.Println("=========================================")

	c, d = method3(1, 2)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
}

func method1(a, b int) (c int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 100
}

func method2(a string, b int) (c, d int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 13, 16
}

func method3(a, b int) (c, d int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c = 99
	d = 76
	return
}
