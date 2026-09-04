package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 访问对应链接,将结果+错误分别填入resp和err中
	resp, err := http.Get("http://example.com/")
	// 如果有错误就输出并且返回结束
	if err != nil {
		fmt.Println(err)
		return
	}
	// 程序结束时关闭网页Body
	defer resp.Body.Close()
	// 打印出结果
	fmt.Println(resp)

}