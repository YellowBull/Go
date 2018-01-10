package main

import (
	"fmt"
)

func main() {
LABLE1:
	for { //这是个无限循环
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABLE1 //跳出LABLE1同级处的循环语句，即这里的无限循环
			}
		}
	}
	fmt.Println("Jump Break")

	for { //这是个无限循环
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABLE2 //跳到指定的LABLE2的位置达到跳出循环的效果
			}
		}
	}
LABLE2:
	fmt.Println("Jump Goto")

LABLE3:
	for i := 0; i < 10; i++ {
		for { //这是个无限循环
			continue LABLE3 //跳出此次循环且调到LABLE3的指定位置继续执行
		}
	}
	fmt.Println("Jump Continue")
}
