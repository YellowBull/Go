package main

import (
	"fmt"
)

func main() {
	//数组切片定义方式一：
	var slice1 = []int{}
	fmt.Println(slice1)

	//数组切片定义方式二：
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := arr[1:5]  //从第二个到第五个值都归入selice中（包左不包右）
	fmt.Println(slice2) //[2 3 4 5]

	//数组切片定义方式三：
	slice3 := make([]int, 3, 6) //arg0 类型 arg1 长度 arg3 预存长度
	fmt.Println(slice3)

	//从数组中获得的片长度超过预存长度可以增长
	//从数组切片中获得切片长度超过预存长度会报错
	slice3 = arr[1:9]
	slice4 := make([]int, 2, 3)
	// slice4 = slice3[0:5]//报错
	slice4 = slice3[0:2]
	fmt.Println(slice3)
	fmt.Println(slice4)

	//copy 方法
	slice5 := make([]int, 3, 5)
	copy(slice5, slice2)
	fmt.Println(slice5) //[2 3 4]
	copy(slice5, slice2[5:8])
	fmt.Println(slice5) //[7 8 9]

	//append 方法
	copy(slice5, slice2[1:4])
	fmt.Println(slice5) //[3 4 5]
	slice5 = append(slice5, 1, 2, 3, 4, 5, 6)
	//slice5 = append(slice5, slice3[1:6])//cannot use slice3[1:6] (type []int) as type int in append
	fmt.Println(slice5) //[3 4 5 1 2 3 4 5 6]
}
