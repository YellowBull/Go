package main

import (
	"fmt"
)

const ( //常量定义默认约定都大写，如果是定义内部常量而不想被包外内容访问可以用'_'开头
	E = 1
	F //会赋上面的值
	G
	A, B, C, D = 1, "a", 'c', A //常量定义时'c'会转为ascii对应的int值
	H, I, J, K                  //会赋上面对应的值
	L          = iota           //标记l在当前组中的位置，是个枚举类从0开始，但是不会因为多值复制而变此时L=5
	_ONE       = 1              //只有内部能访问的常量因为原因是'_'不是大写字母
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)
	fmt.Println(K)
	fmt.Println(L)
	fmt.Println(_ONE)
}
