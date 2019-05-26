package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子，只需一次
	//rand.Seed(100)   //如果种子一样，每次运行程序产生的随机数就一样
	rand.Seed(time.Now().UnixNano()) //以当前时间作为种子参数，这样可以保证每次随机数不一样

	for i := 0; i < 5; i++ {
		//产生随机数
		//fmt.Println("rand = ", rand.Int()) //产生随机很大的数
		fmt.Println("rand = ", rand.Intn(100)) //限制产生的随机数在100以内
	}
}
