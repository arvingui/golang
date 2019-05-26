package main

import (
	"fmt"
)

func main() {
	num := 1

	switch num {
	case 1:
		fmt.Println("按下的是1楼")
		break //go语言保留了break关键字，不写也默认包含，所以写不写一样
	case 2:
		fmt.Println("按下的是2楼")
		break
	case 3:
		fmt.Println("按下的是3楼")
		break
	default:
		fmt.Println("按下的是10楼")
	}

	switch num {
	case 1:
		fmt.Println("按下的是1楼")
		fallthrough //fallthrough就相当于c语言不加break，会列出第一个符合case加下面所有case
	case 2:
		fmt.Println("按下的是2楼")
		fallthrough
	case 3:
		fmt.Println("按下的是3楼")
		fallthrough
	default:
		fmt.Println("按下的是10楼")
	}
}
