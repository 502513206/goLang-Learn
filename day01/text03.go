package main

import (
	"fmt"
	"sync"
	"time"
)

var a chan int
var b [10]chan int

func sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
		fmt.Println("In goroutine %d\n", v)
		time.Sleep(10 * time.Millisecond)

	}
	fmt.Println(sum)
	return sum
}

// 放入另一个协程中
func hello(pipline chan int) {
	fmt.Println(<-pipline)
}

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d: %d\n", x, i)
	}
}

func add(count *int, group *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count++
		lock.Unlock()
	}
	group.Done()
}
func main() {
	var lock *sync.Mutex
	lock = new(sync.Mutex)
	count := 0
	var a1 sync.WaitGroup
	a1.Add(3)
	go add(&count, &a1, lock)
	go add(&count, &a1, lock)
	go add(&count, &a1, lock)
	//a1.Add(2)
	//go worker(1, &a1)
	//go worker(2, &a1)
	a1.Wait()
	fmt.Println(count)
	//go sum(1, 2, 3, 4, 5)
	//go sum(6, 7, 8, 9, 10)
	//fmt.Println("执行顺序")
	//time.Sleep(time.Second)
	pipline := make(chan int)
	// 死锁
	// 使接受者代码在发送者之前
	// 使用缓存信道
	//mydata := <-pipline
	go hello(pipline)
	pipline <- 200
	//fmt.Println(mydata)
	fmt.Println("信道可缓冲%d 个数据\n", cap(pipline))
	fmt.Println("信道当前有%d 个数据\n", len(pipline))
	//x, ok := <-pipline
	//close(pipline)
}

func init() {
	//fmt.Println("Hello init")
}
