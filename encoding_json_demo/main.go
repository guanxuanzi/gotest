package main
import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
}
// 有tag版本
type Person2 struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	// 字面量初始化
	person := Person{Name:"John",Age:30}
	// 返回类型 []byte----将go结构转为json
	jsonData,err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:",err)
		return
	}
	fmt.Printf("jsonData type:%T\n",jsonData) // 返回jsonData的类型 []uint8(也就是[]byte)
	fmt.Println(string(jsonData)) 			  // {"Name":"John","Age":30}
	// 使用反引号方便编写
	jsonStr := `{"Name":"John","Age":20}`
	// 准备一个由构造体Person搭建的空容器person2
	var person2 Person
	// 将json转化为go结构----将json的输出传入&person2容器中 &用来取地址
	err = json.Unmarshal([]byte(jsonStr),&person2)
	if err != nil {
		fmt.Println("Error marshalling to JSON:",err)
		return
	}
	// {John 30}
	fmt.Println(person2)
	// 用字面量进行序列化
	person3 := Person2 {Name:"Amy",Age:30}
	jsonData,err = json.Marshal(person3)
	if err!= nil {
		fmt.Println("Error marshalling to JSON:",err)
		return 
	}
	// {"name":"Amy","age":30}
	fmt.Println(string(jsonData))

	jsonStr = `{"name":"Amy","age":30}`
	var person4 Person2
	err=json.Unmarshal([]byte(jsonStr),&person4)
	if err != nil {
		fmt.Println("Error unmarshalling form JSON:",err)
		return
	}
	// {Amy 30}
	fmt.Println(person4)
}
