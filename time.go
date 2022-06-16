package main

import (
	"fmt"
	"time"
)

func test1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Trứng rán cần mỡ")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Yêu không cần cớ")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Bắp cần bơ")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Cần cậu cơ")
}

// func test2() {
// 	time.Sleep(150 * time.Millisecond)
// 	fmt.Println("Bắp cần bơ")
// 	time.Sleep(150 * time.Millisecond)
// 	fmt.Println("Cần cậu cơ")
// }

func main() {
	go test1()
	// go test2()
	time.Sleep(5 * time.Second)
	fmt.Println("Good bye")

	// c1 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c1 <- "result 1"
	// }()
	// select {
	// case res := <-c1:
	// 	fmt.Println(res)
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("timeout 1")
	// }
	// c2 := make(chan string, 1)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c2 <- "result 2"
	// }()
	// select {
	// case res := <-c2:
	// 	fmt.Println(res)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("timeout 2")
	// }
}
