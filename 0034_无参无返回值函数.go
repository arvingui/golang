/*
	Go语言函数定义格式如下：

	func FuncName (参数列表) (o1 type1, o2 type2 返回类型) {

		函数体

		return v1, v2     返回多个值
	}
*/

/*
	注意：FuncName: 函数名称，根据约定，函数名首字母小写为private，大写为public
		参数列表：函数可以有0个或多个参数，参数格式为：变量名 类型，多个参数用逗号分隔

*/

package main

import "fmt"

//无参无返回值函数的定义
//注意：调用函数的定义放在main函数前后都行，不像c语言放在后面还要在前面声明
func MyFunc() {
	a := 666
	fmt.Println("a = ", a)
}

func main() {
	//无参无返回值函数的调用：函数名()
	MyFunc()
}
