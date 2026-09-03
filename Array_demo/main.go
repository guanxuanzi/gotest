package mainn

import "fmt"


func main(){
	var numbers = [3]int{1,2,3}
	fmt.Println(numbers)
	fmt.Println(numbers[2])
	// 数组直接赋值,会把整个数组复制一份,所以arr1修改不会影响arr
	arr := [3]int{1,2,3}
	arr1 := arr
	arr1[0] = 325
	fmt.Println(arr,arr1)
	// 数组遍历
	scores := [3]int{95,94,93}
	// 方法1
	for i := 0;i<len(scores);i++{
		fmt.Printf("学生%d的成绩是:%d\n",i+1,scores[i])
	}
	// 方法2
	for index,score := range scores{
		fmt.Printf("学生%d的成绩是:%d\n",index+1,score)
	}
}
