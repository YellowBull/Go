package main

import (
	"fmt"
)

func main() {
	a := 1
	switch a { //switch用法一
	case 0:
		fmt.Println("a = 0")
	case 1:
		fmt.Println("a = 1")
	default:
		fmt.Println("None")
	}

	switch { //switch用法二
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("None")
	}
	//switch用法三(定义局部变量进行判定,且作为局部变量定义的时候定义完成后用';'隔开
	//局部变量的好处时即使与全局变量重名也不影响局部变量的使用)
	switch a := 1; {
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("None")
	}
}
