package main

import (
	"fmt"
)

type A struct {
	Name string
}

func (a A) Tx() {
	a.Name = "aa"
	fmt.Println("A tx")
}

func (a *A) print() { //指针传递可以改变属性值
	a.Name = "aa"
	fmt.Println("A print")
}

func (a *A) count(c int) int { //带参数和返回值的方法
	for i := 0; i <= 100; i++ {
		c += i
	}
	return c
}

func main() {
	a := A{}
	a.Tx()
	fmt.Println(a.Name) //
	a.print()
	fmt.Println(a.Name) // aa

	fmt.Println(a.count(0)) // 5050
}
