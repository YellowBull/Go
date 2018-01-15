package main

import (
	"fmt"
)

type Usb interface {
	Name() string
	Connect
}

type Connect interface {
	Connect()
}
type PhoneConnect struct {
	name string
}

func (pc PhoneConnect) Name() string { // 不能写入指针
	return pc.name
}

func (pc PhoneConnect) Connect() {
	fmt.Println("PhoneConnect : ", pc.name)
}

type TVConnect struct {
	name string
}

func (tv TVConnect) Connect() {
	fmt.Println("PhoneConnect : ", tv.name)
}
func main() {
	// 调用方式一
	// var usb Usb
	// pc := PhoneConnect{"PhoneConnect"}
	// usb = Usb(pc)
	// usb.Connect()

	// 调用方式二
	// pc := PhoneConnect{"PhoneConnect"}
	// pc.Connect()
	// Disconnect(pc)

	// 调用方式三
	// pc := PhoneConnect{"PhoneConnect"}
	// var conn Connect
	// conn = Connect(pc) //强制类型转换
	// conn.Connect()

	// 错误的调用方式
	var usb Usb
	tv := TVConnect{"TVConnect"}
	usb = Usb(tv)
	usb.Connect() // TVConnect does not implement Usb (missing Name method)
}

func Disconnect(usb interface{}) { //usb Usb//证明已实现接口（receiver）
	// fmt.Println("Disconnect : " + usb.Name())//方式一

	// 方式二
	// if pc, ok := usb.(PhoneConnect); ok { //如果是PhoneConnect实现了Usb接口则返回ture
	// 	fmt.Println("Disconnect : " + pc.Name())
	// }

	// 方式三
	switch usb.(type) {
	case PhoneConnect:
		fmt.Println("Disconnect : " + usb.(PhoneConnect).name)
	default:
		fmt.Println("未定义类型")
	}
}
