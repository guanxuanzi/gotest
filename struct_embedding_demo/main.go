package main
import "fmt"

type Person struct {
	Name string
	Age int
}
// 嵌入Person的指针
type Student struct {
	*Person
	School string
}
// 嵌入Person
type Teacher struct {
	Person
	School string
}

func main() {
	stu := Student{
		Person:&Person{Name:"John",Age:20},
		School:"MIT",
	}
	fmt.Println(stu) // {0x3586197ce0c0 MIT}
	// 嵌入字段可以直接通过结构体名访问
	stu.Person.Age=21
	fmt.Println(stu.Person.Age)

	tea := Teacher{
		Person: Person{Name:"Amy",Age:30},
		School:"Harvard",
	}
	fmt.Println(tea) // {{Amy 30} Harvard}
	tea.Person.Age=25
	fmt.Println(tea.Person.Age)

	// 嵌入指针,大家共享一个Person,一个修改全部一起改
	stu1 := stu
	stu1.Person.Age = 22
	fmt.Println(stu1.Person.Age) // 22
	fmt.Println(stu.Person.Age) // 22

	// 嵌入值,每个有自己的Person,修改不会影响到其他
	tea1 := tea
	tea1.Person.Age = 40
	fmt.Println(tea1.Person.Age) // 40
	fmt.Println(tea.Person.Age) // 25
}