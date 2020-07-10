package main

import (
	"fmt"
)

var array [5]int

// 表示值类型
// 声明
var a1 int

// 直接赋值为int类型
// 赋值
var a = 15
var b = "hello"

// _ 赋值丢弃，34被丢弃掉
var _, b1 = 34, 35

const (
	// 第一个iota表示为0
	a2 = 0
	b2 = "0"
	c
)

// 声明成组
var (
	x int
	y bool
)

func Judge() {
	if x > 0 {
	} else {
	}
}

func Judge2() {
	i := 0
Here:
	i++
	fmt.Println(i)
	goto Here
}

func loop() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func loop2() {
	var sum int
	for sum < 45 {
		sum += 1
	}
	fmt.Println(sum)
}

func loop3() {
	list := []string{"a", "b", "c"}
	for k, v := range list {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func loop4(i int) {
	switch i {
	case 1:
		fallthrough
	case 2:
		fmt.Println(2)
	}
}

func main() {
	//loop()
	//loop2()
	//loop3()
	loop4(2)
	//Judge2()
	//b = "c"
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(a2)
	//fmt.Println(b2)
	//fmt.Println(c)
	//array := [5]int{10, 20, 30, 40, 50}
	//for i := 0; i < 5; i++ {
	//	fmt.Print(array[i])
	//}
	//fmt.Println(os.TempDir())
	//fmt.Println(log.Flags())
}
