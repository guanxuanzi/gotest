package main
import "fmt"

type Animal interface{
	Speak() string
}

func Speak(a Animal) {
	fmt.Println(a.Speak())
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	dog := Dog{Name:"Rex"}
	cat := Cat{Name:"Whiskers"}

	Speak(dog)
	Speak(cat)
}