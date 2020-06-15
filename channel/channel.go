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

	//ch1 := make(chan int)
	//done := make(chan bool)
	//go func() {
	//	fmt.Println("in goroutine...")
	//	time.Sleep(3 * time.Second)
	//	data := <-ch1
	//	fmt.Println("data: ", data)
	//	done <- true
	//}()
	//
	//time.Sleep(5 * time.Second)
	//ch1 <- 100
	//
	//<-done
	//fmt.Println("over")

	/*
		select 类似于 switch 语句，
		但 select 会随机执行一个可运行的 case。
		如果没有 case 可运行，它将阻塞，直到有 case 可运行。
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1: ", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2: ", num2)
		} else {
			fmt.Println("ch2 close...")
		}
	}
}
