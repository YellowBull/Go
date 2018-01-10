package main

import (
	"fmt"
)

func main() {
	a := 1
	for { //for循环的第一种形式等同于while(true){}
		a++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}
	fmt.Println("Over")
	for a < 6 { //for循环的第二张形式等同于while(条件){}
		a++
		fmt.Println(a)
	}
	fmt.Println("Over2")

	for i := 0; i < 3; i++ { //经典for循环
		fmt.Println(i)
	}
	fmt.Println("Over3")
}
