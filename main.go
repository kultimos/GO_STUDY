package main

// 导包方式也跟java有些区别
import (
	"fmt"
	"time"
)

func main() { //左花括号必须跟方法声明在同一行
	fmt.Println("我先睡一秒")
	// 方法末尾不需要也不建议使用 ;
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, World!")
}
