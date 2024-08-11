package init_package

import "fmt"

func init() {
	fmt.Println("另一个类的初始化")
}

func Method() (a string) {
	fmt.Println("另一个类的普通方法")
	return "来了老弟"
}
