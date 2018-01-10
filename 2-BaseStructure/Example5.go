package main

import (
	"fmt"
)

func main() {
	a := 2
	if a, _ := 1, 2; a > 0 {
		fmt.Println(a) //a=1
	}
	fmt.Println(a) //a=2
}
