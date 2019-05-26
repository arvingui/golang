//1、不同目录，包名不一样（package main）

package calc //不同目录，包名不一样（package main）

//init函数是每个包的内部自带函数，每次都会第一个执行
func init() {
	fmt.Println("this is calc init")
}

func Add(a, b int) int {
	return a + b
}
