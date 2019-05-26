package main

import (
	"fmt"
)

func test() {
	a := 10
	fmt.Println("a = ", a)
}

func main() {
	//定义在{}里面的变量就是局部变量，只能在{}里面才有效
	//执行到定义变量那句话，才开始给变量分配内存，离开作用域自动释放内存

	//a = 111
	{
		i := 10
		fmt.Println("i = ", i)
	}

	//上面那个错误不经常犯，下面这个错误经常犯
	if flag := 3; flag == 3 {
		fmt.Println("flag = ", flag)
	}
	flag = 4 //报错flag未定义，上面的flag := 3只是在if的{}里有效
}
