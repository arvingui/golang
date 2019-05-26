package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "yoyo", 3: "go"}

	//第一个返回值为key，第二个返回值为value，遍历结果是无序的，每次执行顺序可能都不一样
	for key, value := range m1 {
		fmt.Printf("%d==========> %s\n", key, value)
	}

	//如何判断一个key值是否存在
	//第一个返回值为key所对应的value，第二个返回值是key是否存在的结果，存在为true，否则为false
	value, ok := m1[1]
	if ok == true {
		fmt.Println("m1[1] = ", value)
	} else {
		fmt.Println("key不存在")
	}
}
