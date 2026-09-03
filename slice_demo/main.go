package main

import "fmt"

func main(){
	// 声明切片的四种方式
	// 仅声明,a == nil
	var a []string
	// 用make创建,长度为4
	f := make([]string,4)
	// 用字面量创建的空切片,不为零值
	b := []int{}
	// 用字面量创建的非空切片
	c:=[]int{1,2,3}
	fmt.Println(a == nil)
	fmt.Println(f)
	fmt.Println(b == nil)
	fmt.Println(c)
	// 添加元素进切片c
	c = append(c,4)
	fmt.Println(c)
	// 切片的切片
	d := c[0:2]
	fmt.Println(d)
	// 长度len()和容量cap()
	fmt.Printf("长度:%d 容量:%d\n",len(d),cap(d))
}