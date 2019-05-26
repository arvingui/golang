package main

import (
	"fmt"
	"regexp" //正则表达式需要的包
)

func main() {
	buf := "abc azc a7c aac 888 a9c tac"

	//1)解释规则，它会解析正则表达式，如果成功返回解释器
	regcxp.MustCompile(`a.c`) //反引号在golang代表原始字符串，这里也可用双引号
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}

	//2)根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1) //-1代表提取所有符合条件的字符串，这里如果填1就展示1个
	fmt.Println("result1 = ", result1)
}
