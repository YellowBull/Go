package main

import (
	"fmt"
)

func main() {
	//数组的定义方式一：仅指定数组长度的定义方式
	var arr [3]int
	fmt.Println(arr)
	//数组的定义方式二：指定第一二个值的定义方式
	var arr2 = [3]int{1, 3}
	fmt.Println(arr2)
	//数组的定义方式三：指定第二个值为2的定义方式
	var arr3 = [3]int{0: 1, 1: 3}
	fmt.Println(arr3)
	//数组的定义方式四：不指定数组长度的定义方式
	arr4 := [...]int{1, 2, 3}
	fmt.Println(arr4)
	//数组的定义方式五：不指定数组长度的定义方式但指定第5个位置长度为10的定义方式
	arr5 := [...]int{4: 10}
	fmt.Println(arr5) //[0 0 0 0 10]
	//数组的定义方式六：给定长度的二维数组定义
	arr6 := [2][3]int{}
	fmt.Println(arr6)
	//数组的定义方式七：不给定长度的二维数组定义
	arr7 := [...][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr7) //[[1 2 3] [4 5 6]]

	/**
	 * 数组比较
	 */
	fmt.Println(arr2 == arr3) //true

	/**
	 * 冒泡
	 */
	arr8 := []int{4, 5, 2, 6, 1, 9, 7, 8}
	fmt.Println(arr8) //[4 5 2 6 1 9 7 8]
	num := len(arr8)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arr8[i] > arr8[j] {
				arr8[i], arr8[j] = arr8[j], arr8[i] //交换俩值
			}
		}
	}
	fmt.Println(arr8) //[1 2 4 5 6 7 8 9]
}
