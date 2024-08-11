package main

import "fmt"

/*
在声明全局变量时,方式4与其他三种方式是有区别的
即不能使用 := 的方式声明全局变量,该方式只能在函数体内声明
*/
var gA int
var gB int = 200
var gC = 25.123

//gD := "sad"

func main() {
	/**
	  四种变量的声明方式
	*/
	var a int      // 1.声明一个变量,值为默认值
	var b int = 10 // 2.声明一个变量.并提供一个初始值
	var c = 22.2   // 3.初始化时省略数据类型,通过赋予的值自动匹配当前的变量的数据类型
	d := "yoo"     // 4.最常用的一种方式, :=表示声明+初始化

	fmt.Println("a的默认值是", a)
	fmt.Println("b的初始值是", b)

	fmt.Println("c的初始值是", c)
	fmt.Printf("c的类型是: %T\n", c)

	fmt.Printf("d = %s, d的类型是: %T\n", d, d)

	fmt.Println("======================================")

	fmt.Println("gA = ", gA)
	fmt.Println("gB =", gB)

	fmt.Println("======================================")

	/**
	声明多个变量
	*/
	var x, y int = 100, 200
	fmt.Println("x =", x, ", y = ", y)

	var kk, ll = 25.1, "hah"
	fmt.Println("kk= ", kk, ", ll = ", ll)

	var (
		vv int  = 100
		tt bool = true
	)
	fmt.Println("vv = ", vv, ", tt = ", tt)
}
