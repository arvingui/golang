package main //就是一个结论，超集可以转换为子集，但子集不可以转换为超集

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type Personer interface { //超集
	Humaner //匿名字段，继承了sayhi()这个方法
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

//Student实现了sayhi()这个方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student) sing(lrc string) {
	fmt.Println("Student在唱着：", lrc)
}

func main() {
	//超集可以转换为子集，但子集不可以转换为超集
	var iPro Personer //超集
	iPro = &Student{"mike", 666}

	var i Humaner //子集

	//iPro = i //error
	i = iPro //可以，超集可以转换为子集
	i.sayhi()

}
