package main
import "fmt"

func safeDivide(a,b int) {
	// recover只在defer中有效
	defer func(){
		if r := recover(); r != nil {
			fmt.Println("捕获到panic:",r)
		}
	}()
	// 当b为0时会触发panic
	fmt.Println(a/b)
}

func main() {
	safeDivide(10,2)
	safeDivide(10,0)
	fmt.Println("程序继续进行")
}
