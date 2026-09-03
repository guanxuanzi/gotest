package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串长度 len()
	var str = "this is str"
	fmt.Println(len(str))

	// 字符串拼接
	var str1 = "你好"
	var str2 = "golang"
	fmt.Println(str1 + "," + str2)
	fmt.Println(fmt.Sprintf("%s, %s", str1, str2)) 

	// 字符串分割
	var s = "123-456-789"
	arr := strings.Split(s,"-")
	fmt.Println(arr)

	// 字符串遍历
	str3 := "你好,世界"
	// 方法1:for range循环遍历
	for index,char := range str3{
		fmt.Printf("位置 %d: 字符 %c,Unicode码点 %U\n", index, char, char)
	}
	// 方法2:按字节遍历
	for i := 0;i<len(str3);i++{
		fmt.Printf("位置 %d: 字节 %d\n", i, str3[i])
	}
	// 方法3:将字符串转为[]rune切片后遍历
	runes := []rune(str3)
	for i, r := range runes {
		fmt.Printf("位置 %d:字符 %c\n", i, r)
	}
}