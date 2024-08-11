package main

import "fmt"

func main() {
	map1 := map[string]string{
		"China":  "Beijing",
		"USA":    "MIL",
		"FRANCE": "PARIS",
	}
	printMap(map1)
	delete(map1, "China")
	printMap(map1)
	map1["USA"] = "LA"
	printMap(map1)
}

func printMap(theMap map[string]string) {
	for key, value := range theMap {
		fmt.Println("key: ", key, ", value:", value)
	}
	fmt.Println("======================")
}
