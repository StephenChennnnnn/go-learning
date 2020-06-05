package main

import (
	"fmt"
	"time"
)

func main() {
	//var ch1 chan bool
	//fmt.Println(ch1)
	//fmt.Printf("%T\n", ch1)
	//ch1 = make(chan bool)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println("in goroutine, i: ", i)
	//	}
	//	ch1 <- true
	//	fmt.Println("end...")
	//}()
	//
	//data := <-ch1
	//fmt.Println("data --> ", data)
	//fmt.Println("over")

	ch1 := make(chan int)
	done := make(chan bool)
	go func() {
		fmt.Println("in goroutine...")
		time.Sleep(3 * time.Second)
		data := <-ch1
		fmt.Println("data: ", data)
		done <- true
	}()

	time.Sleep(5 * time.Second)
	ch1 <- 100

	<-done
	fmt.Println("over")
}
