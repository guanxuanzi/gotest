package main

import "fmt"

// iota枚举,用来定义一组连续的常量
// 从0开始,每多一行 +1
const (
	StatusPending = iota
	StatusRunning
	StatusDone
)

func main() {
	// 定义常数(不可修改)
	const Pi = 3.14159
	const AppName = "MyApp"
	fmt.Println(Pi,AppName)
	fmt.Println(StatusPending)
	fmt.Println(StatusRunning)
	fmt.Println(StatusDone)
}



