package main
import "fmt"

// 闭包----工厂函数
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 闭包----计数器
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	double := makeMultiplier(2)
	fmt.Println(double(5)) // Output: 10

	counter := makeCounter()
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
}