package main
import "fmt"

// 定义一个接口Animal 满足Speak()方法的都可以通过接口来调用
type Animal interface{
	Speak() string
}
// 定义一个函数,用于调用Speak方法
func Speak(a Animal) {
	fmt.Println(a.Speak())
}
// 定义一个结构体Dog
type Dog struct {
	Name string
}
// 实现Animal接口的Speak函数,使用类型Dog定义变量
func (d Dog) Speak() string {
	return "Woof!"
}
// 定义一个结构体Cat
type Cat struct {
	Name string
}
// 实现Animal接口的Speak函数,使用类型Cat定义变量
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	dog := Dog{Name:"Rex"}
	cat := Cat{Name:"Whiskers"}
	// 满足接口要求的结构体都可以传入
	Speak(dog)
	Speak(cat)
}