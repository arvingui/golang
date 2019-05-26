/*
字符串转换常用的函数：
Append、Format、Parse
Append函数：将整数转换为字符串后，添加到现有的字节数组中
Format函数：把其他类型的转换为字符串
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)

	//第2个数为要追加的数，第3个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcgohello")
	fmt.Println("slice = ", string(slice)) //转换string后再打印，执行结果为:true1234"abcgohello"

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	//'f'指打印格式，以小数方式， -1指小数点位数（紧缩模式），64代表以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str = ", str) //打印结果为：3.14

	//整形转字符串常用方式
	str = strconv.Itoa(6666)
	fmt.Println("str = ", str)

	//字符串转其他类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("true") //这个函数有两个返回值
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	//把字符串转换为整形
	a, _ := strconv.Atoi("567")
	fmt.Println("a = ", a)
}
