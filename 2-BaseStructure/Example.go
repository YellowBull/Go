/**
 * 程序包名
 * (只有package为main的包才能使用main函数)
 *（main包里必须有main函数）
 */
package main

/**
 * 导入的包名
 * （如果你导入了不相关的包，编译会报错，强迫你删掉）
 * 	如果你看包名不爽可以起别名格式：
 * 	import ( 
 * 		myName "fmt"
 * 	)
 * 	或者
 * 	import ( 
 * 		. "fmt"
 * 	)
 * 	这种方式调用方法不需要指定包名（不建议）
 * 	例如Println()可直接使用
 */
import (
	"fmt"
)

//常量定义
const (
	PI=3.14
)

//全局变量定义
var (
	name = "xxx"
)

//一般类型声明
type number int

//结构声明
type myStruct struct{}

//接口声明
type myInterface interface{}

//主函数（函数首字母大小写区分是否可以被外部调用public/private）
func main() {
	fmt.Println("Hello word 你好世界")
}
