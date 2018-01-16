package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	User
	Color string
}

type User struct {
	Name string
	Age  int
	Sex  string
}

func (user User) Eat(food string) {
	fmt.Println("My name is ", user.Name, "I like eating ", food)
}

func (user User) Like() {
	fmt.Println("I like you")
}

func main() {

	us := User{Name: "张三", Age: 22, Sex: "男"}
	Info(us)        // 查看具体信息
	fmt.Println(us) // {张三 22 男}
	Set(&us)        // 修改name的值
	fmt.Println(us) // {李四 22 男}

	cla := reflect.ValueOf(us)
	meth := cla.MethodByName("Eat")                // 根据方法名获取方法
	args := []reflect.Value{reflect.ValueOf("苹果")} // 参数集合slice
	meth.Call(args)                                // 使用获取的方法并传入参数 My name is  李四 I like eating  苹果

	per := Person{User: User{Name: "张三", Age: 22, Sex: "男"}, Color: "Yellow"}
	cal := reflect.TypeOf(per)
	color := cal.Field(1)
	name := cal.FieldByIndex([]int{0, 0})
	user := cal.Field(0)
	fmt.Println(user)  // {User  main.User  0 [0] true} 属性名称 属性类型 （未知。。。） 属性位置 是否为匿名属性
	fmt.Println(name)  // {Name  string  0 [0] false}
	fmt.Println(color) // {Color  string  40 [1] false}
	fmt.Println(cal)   // main.Person

	x := 123
	v := reflect.ValueOf(&x) // 拿到x的地址 0xc04204a130
	v.Elem().SetInt(999)     // 设定值
	fmt.Println(x)           // 999

}

func Info(in interface{}) {
	cla := reflect.TypeOf(in) // 拿到接口类型 main.User
	fmt.Println("TypeOf :", cla)

	vals := reflect.ValueOf(in) // 拿到接口的所有值
	fmt.Println("ValueOf", vals)

	if k := cla.Kind(); k != reflect.Struct {
		fmt.Println("传入了错误类型")
		return
	}

	for i := 0; i < cla.NumField(); i++ { // 根据属性长度遍历所有属性
		attr := cla.Field(i)             // 拿到第i个属性
		val := vals.Field(i).Interface() // 拿到第i个属性的值
		fmt.Println(attr.Name, attr.Type, val)
	}

	for i := 0; i < cla.NumMethod(); i++ { // 根据接口中方法个数遍历
		method := cla.Method(i)               // 拿到第i个方法
		fmt.Println(method.Name, method.Type) // Eat func(main.User, string) Like func(main.User)
	}
}

func Set(in interface{}) {
	cla := reflect.ValueOf(in)
	if cla.Kind() == reflect.Ptr && !cla.Elem().CanSet() {
		fmt.Println("必须为指针而且不能被设置")
		return
	}
	cla = cla.Elem() // 可修改对象

	name := cla.FieldByName("Name") // 拿到指定名称的属性
	if !name.IsValid() {            // 如果名称不存在
		return
	}
	if name.Kind() == reflect.String { // 如果类型为string
		name.SetString("李四")
	}
}
