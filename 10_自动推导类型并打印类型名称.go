package main //一个go程序必须有一个main包

import ( //导入包含，必须要使用
	"fmt"
)

func main() {
	//自动推导类型
	c := 10
	fmt.Printf("c type is %T\n", c) //%T为打印数据类型
	s := "hello"
	fmt.Println("s=", s)
	fmt.Printf("s type is %T\n", s)
}
