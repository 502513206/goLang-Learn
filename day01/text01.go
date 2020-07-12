package main

import "fmt"

func printHello() {
	fmt.Println("Hello World")
}

type Animals struct {
	name string
}

type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company
}

type phone interface {
	call()
}

type xiaomi struct {
	name string
}

type Duck interface {
	call()
	fly()
}

func (phone xiaomi) call() {
	fmt.Println("小米手机")
}
func mainqwe() {
	myCom := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}
	staffInfo := staff{
		name:     "张三",
		age:      18,
		gender:   "男",
		position: "go开发",
		company:  myCom,
	}
	fmt.Println("%s在%s工作\n", staffInfo.name, staffInfo.companyName)

}
