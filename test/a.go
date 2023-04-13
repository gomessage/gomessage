package main

import (
	"fmt"
	"gomessage/test/mock"
	"gomessage/test/real2"
)

// Retriever 接口
type Retriever interface {
	Get(curl string) string
}

// 函数
func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {

	var aaa Retriever

	aaa = mock.Retrievers{Context: "zhangsan"}
	fmt.Println(download(aaa))

	aaa = real2.Retrievers{UserAgent: "aaa", TimeOut: 3}
	fmt.Println(download(aaa))
}
