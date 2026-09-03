package main

import "fmt"

func main(){
	// 两种创建map的方法
	scoresEmpty1 := map[string]int{}
	scoresEmpty2 := make(map[string]int)
	fmt.Println(scoresEmpty1,scoresEmpty2)

	scores := map[string]int{
		"张三": 90,
		"李四": 80,
	}
	fmt.Println("学生成绩:", scores)

	// 添加修改元素
	scores["张三"] = 100
	scores["王五"] = 70
	fmt.Println("学生成绩:", scores)

	// 获取元素
	zhang := scores["张三"]
	fmt.Println("张三的成绩:", zhang)

	// 使用ok判断键是否存在
	score, ok := scores["赵六"]
	if ok {
		fmt.Println("赵六的成绩:", score)
	}else{
		fmt.Println("赵六的成绩不存在")
	}

	// 删除元素
	delete(scores,"李四")
	fmt.Println("删除李四后:", scores)

	// 遍历元素
	for k,v := range scores {
		fmt.Printf("%s:成绩=%d\n", k, v)
	}

	// 获取map长度
	fmt.Println("map长度:", len(scores))
}
