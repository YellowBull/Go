package main

import (
	"fmt"
)

type human struct {
	Age int
}

type woman struct {
	human     // 组合
	Sex, Name string
}
type man struct {
	human
	Name, like string
	Age        int
}

/**
 * 内部构造
 */
type animal struct {
	Name      string
	Constract struct {
		Age int
	}
}

func main() {
	xx := &woman{Name: "xx", Sex: "woman", human: human{Age: 12}} //构造是值类型传递
	fmt.Println(xx)
	xy := man{}
	xy.Name = "xy"
	xy.like = "money"
	xy.human.Age = 13
	//xy.Age //如果未定义Age与上面等效，如果定义了则不能使用这种方式
	fmt.Println(xy)

	changeName(xx)
	fmt.Println(xx)

	/**
	 * 匿名构造
	 */
	stru := &struct {
		Name string
		Age  int
	}{
		Name: "ss",
		Age:  12,
	}
	fmt.Println(stru)

	cat := animal{Name: "cat"}
	cat.Constract.Age = 1
	fmt.Println(cat)
}

func changeName(xx *woman) { //指针传递
	xx.Name = "GG"
}
