package main

import (
	"fmt"
)

func main() {
	//defer关键字
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Print(i) // 3 3 3 这是一个闭包，i是引用的方式传值
		}()
	}

	for i := 0; i < 3; i++ {
		defer fmt.Print(i) // 2 1 0 defer传参在i被定义时就拷贝了
	}

	A()

	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i) // 3 2 1 0
		defer func() {
			fmt.Println("defer() i=", i) // 4 4 4 4
		}()
		fs[i] = func() {
			fmt.Println("defer[] i=", i) // 4 4 4 4
		}
	}
	for _, v := range fs {
		v()
	}
}

//recover关键
func A() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover") // recover只有在defer调用的函数中才有用，作用是恢复程序
		}
	}()
	//panic关键字
	panic("panic") // 相当于抛出一个异常然后终止程序
}
