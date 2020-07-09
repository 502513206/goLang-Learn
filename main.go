package main

import (
	"fmt"
	"log"
	"os"
)

var array [5]int

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < 5; i++ {
		fmt.Print(array[i])
	}
	fmt.Print(os.TempDir())
	fmt.Print(log.Flags())
}
