package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//定义map方式一
	var myMap map[int]string = make(map[int]string)
	fmt.Println(myMap)

	//定义map方式二
	myMap2 := make(map[int]string)
	fmt.Println(myMap2)

	//给定义的map赋值
	myMap2[1] = "OK"
	fmt.Println(myMap2[1]) // OK

	//定义一个map数组
	mapSlice := make([]map[int]string, 5)
	fmt.Println(mapSlice) // [map[] map[] map[] map[] map[]]

	//定义一个map中的map
	doubleMap := make(map[int]map[int]string)
	fmt.Println(doubleMap)
	doubleMap[1] = make(map[int]string)
	doubleMap[1][1] = "my double map"
	fmt.Println(doubleMap) // map[1:map[1:my double map]]

	//判定map是否创建(内部map并没有创建)val为空
	val, ok := doubleMap[2][1]
	if !ok {
		//未创建
		doubleMap[2] = make(map[int]string)
		doubleMap[2][1] = "Good"
	}
	val, ok = doubleMap[2][1]
	fmt.Println(val, ok) // Good true

	/**
	 * 给map数组赋值
	 * @type {[type]}
	 */
	for i, _ := range mapSlice {
		fmt.Print(mapSlice[i]) // map[]map[]map[]map[]map[]
		mapSlice[i] = make(map[int]string)
		mapSlice[i][i] = strconv.Itoa(i)
	}
	fmt.Println()
	fmt.Println(mapSlice) // [map[0:0] map[1:1] map[2:2] map[3:3] map[4:4]]

	/**
	 * 遍历map
	 */
	//赋值方式一
	myMap[1] = "a"
	myMap[2] = "b"
	myMap[3] = "c"
	myMap[4] = "d"
	//赋值方式二
	// myMap = map[int]string{5: "e", 6: "f"}
	for k, v := range myMap {
		fmt.Print(k)
		fmt.Print(v) //2b 3c 4d 1a无序
	}
	fmt.Println()
	//排序map(利用有序的数组间接使得map有序)
	sortSlice := make([]int, len(myMap))
	i := 0
	for k, _ := range myMap {
		sortSlice[i] = k
		i++
	}
	fmt.Println(sortSlice) // [3 4 1 2 0]

	sort.Ints(sortSlice)   // 给int类型数据排序
	fmt.Println(sortSlice) // [1 2 3 4] 这样就可以按顺序取到对应的map值
	for _, v := range sortSlice {
		fmt.Print(myMap[v]) //abcd
	}
	fmt.Println()

	/**
	 * 改变map中的值
	 */
	//基础方式
	fmt.Println(myMap[1])
	myMap[1] = "change"
	fmt.Println(myMap[1])
	//for循环中改变
	fmt.Println(myMap) // map[1:change 2:b 3:c 4:d]
	for _, v := range myMap {
		fmt.Print(v)
		v = "allChange"
		fmt.Print(v)
	} // c allChange d allChange change allChange b allChange
	fmt.Println()
	fmt.Println(myMap) // map[1:change 2:b 3:c 4:d] for循环中只改变值不会影响原本的map

	for k, _ := range myMap {
		myMap[k] = "allChange"
	}
	// map[1:allChange 2:allChange 3:allChange 4:allChange]
	// for循环中根据key值可以改变原本的map的值
	fmt.Println(myMap)
}
