package main

import (
	"fmt"
)

var array [5]int = [5]int{1, 2, 3, 4, 5}
var array2 = [...]int{}

// 	声明字符串切片
var strList []string

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

func mapFunc() {
	var sources map[string]int = map[string]int{"english": 80}
	sources2 := map[string]int{"english": 100}
	sources3 := make(map[string]int)
	if sources3 == nil {
		fmt.Println("进行初始化")
	}
	fmt.Println(sources)
	fmt.Println(sources2)
	fmt.Println(sources3)
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

func loop4(i int) int {
	switch i {
	case 1:
		fallthrough
	case 2:
		fmt.Println(2)
	}
	return 1
}

func testFunc() {
	// new(Type),初始化零值
	ptr := new(int)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println("执行")
	// *field:解析地址中的值
}

func annoFunc() (int, int) {
	return 100, 200
}

func stringFunc() {
	var a4 byte = 'A'
	var b4 byte = 65
	var c4 rune = 'B'
	fmt.Println(a4)
	fmt.Println(b4)
	fmt.Println(c4)
}

func exceFunc() {
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)

	t2 := i.(string)
	fmt.Println(t2)

	defer func() {
		recover()
	}()
}

func selectFunc() {
	cc1 := make(chan string, 1)
	cc2 := make(chan string, 1)
	cc2 <- "hello"
	select {
	case msg1 := <-cc1:
		fmt.Println("c1 received:", msg1)
	case msg2 := <-cc2:
		fmt.Println("c2 received:", msg2)
	default:
		fmt.Println("...")
	}
}

func main() {
	exceFunc()
	mapFunc()
	a3, _b3 := annoFunc()
	_a3, b3 := annoFunc()
	fmt.Println(a3, b3, _a3, _b3)
	// 等当前函数执行完在执行
	// 晚于return执行，可用于释放资源
	defer testFunc()
	name := "张三"
	// 复制地址
	cName := &name
	fmt.Println(name)
	name = "李四"
	fmt.Println(name)
	fmt.Println(*cName)
	//loop()
	//loop2()
	//loop3()
	//loop4(2)
	fmt.Println(c)
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
