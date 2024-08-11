package main

import "fmt"

func main() {
	// 声明myMap1是一个key为string,value为int类型的map,第一行是声明,第二行是实例化
	var myMap1 map[string]int
	myMap1 = make(map[string]int, 2)
	myMap1["java"] = 1
	myMap1["vue"] = 2
	myMap1["go"] = 3
	fmt.Println(myMap1)

	myMap2 := make(map[string]string)
	myMap2["30"] = "Curry"
	myMap2["23"] = "James"
	myMap2["34"] = "Giannis"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"php":  "mateng",
		"java": "wysgyau",
		"nil":  "mengxin",
	}
	fmt.Println(myMap3)
}
