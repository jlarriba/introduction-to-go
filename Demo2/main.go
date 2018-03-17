package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello, world.\n")

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	if sum == 45 {
		fmt.Printf("success!\n")
	}

	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	for {
		fmt.Println("Forever!")
		time.Sleep(time.Second * 3)
	}
}
