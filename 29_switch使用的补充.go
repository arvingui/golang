package main

import (
	"fmt"
)

func main() {
	switch num := 6; num { //支持一个初始化语句，用分号分隔
	case 1:
		fmt.Println("按下的是1楼")
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	//一个case可以支持多种情况
	case 4, 5, 6:
		fmt.Println("按下的是4或5或6楼")
	default:
		fmt.Println("按下的是xxx楼")

		score := 85
		switch score { //可以没有条件
		case score > 90: //case后面可以放条件语句，这是跟c语言不一样的
			fmt.Println("优秀")
		case score > 80:
			fmt.Println("良好")
		case score > 70:
			fmt.Println("中等")
		case score > 60:
			fmt.Println("及格")
		default:
			fmt.Println("不及格")

		}

	}
}
