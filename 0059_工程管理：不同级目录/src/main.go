//1、不同目录，包名不一样（package main）
//2、调用别的包里面的函数，格式：包名.函数名（）
//3、调用别的包的函数，这个函数名字首字母必须大写

package main

import (
	"calc"
	"fmt"
)

//init函数是每个包的内部自带函数，每次都会第一个执行
//执行结果this is calc init
//this is main init
func init() {
	fmt.Println("this is main init")
}

func main() {
	a := calc.Add(10, 20) //不同目录之间函数调用，不仅要写包名引用，而且函数名首字母必须大写
	fmt.Println("a = ", a)
}

//_操作其实是引入该包，而不直接使用包里面的函数，而是调用该包里面的init函数
/*import (
	_ "fmt"
)*/

//src：放源代码，如果有多个文件或多个包，需配置GOPATH环境变量，值为src文件绝对路径
//要想自动生成bin或pkg目录，需使用go install命令，并需配置GOBIN环境变量
//src：源代码  pkg:放平台相关的库   bin:放可执行程序
