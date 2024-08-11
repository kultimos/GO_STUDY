package main

import "fmt"

func main() {
	/**
	声明一个指定前三个数据的可变长数组,此时在func进行参数传递时,实际上就变成了引用传递,所以我估计goLang语言对于普通类型默认是值传递,若需要引用传递
	需要用指针来标识,而对于可变长数组和对象,估计都是默认是引用传递
	*/
	// 第一种声明方式
	var slice1 = []int{1, 2, 3}

	// 第二种声明方式,只声明,未初始化,也并未给该数组在内存中申请空间
	var slice2 []int
	if slice2 == nil {
		fmt.Println("slice2目前为null")
	}
	slice2 = make([]int, 3) //直到下面这步,才在内存中为该数组申请了3个长度的内存空间,并完成了初始化

	var slice3 = make([]int, 5) //声明+初始化 也可以用 slice3 := make([]int ,5)的方式来实现声明+初始化

	for index, value := range slice3 {
		fmt.Println("index: ", index, ", value: ", value)
	}
	changeSliceArray(slice1)
	fmt.Println("==========changeValue后====================")
	for index, value := range slice2 {
		fmt.Println("index: ", index, ", value: ", value)
	}
}

func changeSliceArray(array []int) {
	array[0] = 857
}
