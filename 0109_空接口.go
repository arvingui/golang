/*
空接口不包含任何方法，正因为如此，所有的类型都实现了空接口，
因此可以存储任意类型的数值，它有点类似于C语言的void*类型。

var v1 interface{} = 1    //将int类型赋值给interface{}
var v2 interface{} = "abc"    //将string类型赋值给interface{}
var v3 interface{} = &v2      //将*interface()类型赋值给interface{}
var v4 interface{} = struct{X int}(1)
var v5 interface{} = &struct{X int}(1)

当函数可以接受任意的对象实例时，我们会将其声明为Interface{}，最典型的的例子是
标准库fmt中PrintXXX系列的函数，例如：
func printf(fmt string, args ...interface{})
func Println(args ...interface{})
*/
package main

import "fmt"

/*
比如一般这样用的：表示下面这个函数可以接收任意类型参数
func xxx(arg ...interface{}) {

}
*/

func main() {
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)
}
