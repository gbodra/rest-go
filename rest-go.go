package main

import (
	"fmt"
	"time"
)

func main() {
	initTime := time.Now().Nanosecond()

	fmt.Println("Hello, world.")

	endTime := time.Now().Nanosecond()
	fmt.Println("Time (nanosec): ", (endTime - initTime))
}
