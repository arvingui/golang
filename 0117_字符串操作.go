/*
go语言字符串操作常用的函数有：
Contains、Join、Index、Repeat、Replace、Split、Trim、Fields
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		func Contains(s, substr string) bool
		功能：字符串s中是否包含substr，返回bool值
	*/
	//Contains函数：查看"hellogo"中是否包含"hello"，包含返回true，不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello")) //结果返回true
	fmt.Println(strings.Contains("hellogo", "abc"))   //结果返回false

	/*
		func Join(a []string, sep string) string
		功能：字符串连接，把slice a通过sep连接起来
	*/
	//Join函数：将若干个字符串用指定的字符 组在一起
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "@") //打印出来的结果是abc@hello@mike@go
	fmt.Println("buf = ", buf)

	/*
		func Index(s, sep string) int
		功能：在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	*/
	//Index函数：计算出指定子字符串在另一个母字符串当中的位置是第几个，从0开始数,有就返回位置数，没有返回-1
	fmt.Println(strings.Index("abcdhello", "hello")) //从0开始数，hello在第第4个，执行结果返回4
	fmt.Println(strings.Index("abcdhello", "go"))    //"abcdhello"中没有go，执行结果返回-1

	/*
		func Repeat(s string, count int) string
		功能：重复s字符串count次，最后返回重复的字符串
	*/
	//Repeat函数：按指定次数重复字符串
	buf = strings.Repeat("go", 3) //指将"go" 字符串重复三次，赋值给buf
	fmt.Println("buf = ", buf)    //打印结果为：gogogo

	/*
		func Replace（s, old, new string, n int） string
		功能：在s字符串中,把old字符串替换为new字符串,n表示替换的次数，小于0表示全部替换
	*/
	//Replace函数：
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//运行结果：
	//oinky oinky oink
	//moo moo moo

	//Split函数：以指定的分隔符拆分字符串
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("s2 = ", s2) //执行结果是：[hello abc go mike]

	//Trim函数：去掉字符串两头字符
	buf = strings.Trim("     are  you  ok?       ", " ") //去掉这个字符串两头的空格
	fmt.Printf("buf = #%s#\n", buf)                      //执行结果是are you ok?

	//Fields函数：去掉空格，把元素放入切片中
	s3 := strings.Fields("      are you ok?      ")
	fmt.Println("s3 = ", s3) //执行结果：[are you ok]
	for i, data := range s3 {
		fmt.Println(i, ", ", data) //执行结果：将切片的三个元素are you ok?迭代打印出来
	}
}
