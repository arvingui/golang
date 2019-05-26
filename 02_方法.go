package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Println("SetInfoValue")
}

func (p Person) SetValuePointer() {
	fmt.Println("SetValuePointer")
}

func main() {
	p := Person{"mike", 'm', 18}
	p.SetValuePointer()
}
