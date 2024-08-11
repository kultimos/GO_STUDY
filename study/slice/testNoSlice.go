package main

import (
	"fmt"
)

func main() {
	/**
	声明一个固定长度的数组的两种方式,第二种是对数组的一部分进行赋初值;
	但是这种固定长度的数组有一个很大的弊端在于进行func调用时,是值传递,而非引用传递
	*/
	var arrayA = [10]int{}
	var arrayB = [10]int{1, 2, 3, 4, 5}

	for index, value := range arrayA {
		fmt.Println("index: ", index, ", value: ", value)
	}
	fmt.Println(" =========================")
	for index, value := range arrayB {
		fmt.Println("index: ", index, ", value ", value)
	}
	changeArray(arrayB)
	fmt.Println("=============changeValue以后=============")
	for index, value := range arrayB {
		fmt.Println("index: ", index, ", value ", value)
	}

}

func changeArray(array [10]int) {
	array[1] = 857
	array[2] = 369
	array[3] = 999
}
