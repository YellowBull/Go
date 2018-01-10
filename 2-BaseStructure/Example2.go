package main

import "fmt"

var a int //int类型默认值为0

func main() {
	b, _, d, e := 1, 2, 3, 4 //_下划线表示忽略，同时go可以实现多值赋值
	var f [1]int             //int类型数组默认元素值为0
	var g bool               //bool类型默认值为false
	//var h string = 65
	//var i int = (h)//不能够将一个string类型的值强转为int类型
	j := 100.1
	k := int(j)    //可以将一个float类型的值强转为int类型
	l := string(k) //将一个int类型的值强转为string类型时会使用对应的ascii码对应的字符
	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	//fmt.Println(h)
	//fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
