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

// =========================================
// 空接口
func print(a interface{}) {
	fmt.Println(a)
}

// 空接口用于验证类型断言
var x interface{} = "Hello,Go!"

func main() {
	dog := Dog{Name:"Rex"}
	cat := Cat{Name:"Whiskers"}
	// 满足接口要求的结构体都可以传入
	Speak(dog)
	Speak(cat)

	print("=========================================")
// =========================================
// 空接口类型的切片:可以放入任意类型
	a := []interface{}{"nihao",2,true}
	print(a)
	// key为string,valuer味空接口的Map
	b := map[string]interface{}{"name":"张三","age":20,"gender":"男"}
	print(b)

	// 结构体字段使用空接口
	c := struct {Name interface{}}{Name:"张三"}
	print(c)


	// 用于练习类型断言x.(string)
	if value, ok := x.(string);ok {
		fmt.Println("x是string:",value)
	} else {
		fmt.Println("x不是string")
	}
}