package main

import "fmt"

func main() {
	A() //函数调用方式一

	//函数调用方式二
	b := B
	b("b", 1, 2, 3, 4)

	//函数调用方式三
	str, coun := C("b", 1, 2, 3, 4)
	fmt.Println(str)
	fmt.Println(coun)

	/**
	 * 匿名函数
	 */
	//函数调用方式四，匿名函数调用
	a := 1
	arr := []int{1, 2, 3}
	func(a int, arr []int) { //可直接赋值给指定参数例如 a := func(){}
		a = 3
		arr[1] = 0
	}(a, arr)
	fmt.Println(a)   // 1 基础类型传参是值传递 但是可以通过传递指针的方式在函数中改变值
	fmt.Println(arr) // [1 0 3] 数组是传递了地址

	//函数调用方式五
	arg0 := 10
	clo := closure(arg0)
	fmt.Println(clo(1)) // 11
	fmt.Println(clo(2)) // 12
}

/**
 * 函数定义
 */
//无参数无返回值函数
func A() {
	fmt.Println("none")
}

//带参无返回值函数
func B(a string, b, c, d int, e ...interface{}) {
	fmt.Println(a)
	fmt.Println(b + c + d)
	fmt.Println(e)
}

//带参有返回值的函数
func C(a string, b, c, d int, e ...interface{}) (str string, coun int) {
	fmt.Println("all")
	return a, b + c + d
}

//闭包函数
func closure(arg0 int) func(int) int {
	return func(arg1 int) int {
		return arg0 + arg1
	}
}
