package main

import (
	"fmt"
	"time"
)

func main() {
	//timer := time.NewTimer(3 * time.Second)
	//fmt.Printf("%T\n", timer)
	//fmt.Println(time.Now())
	//time.Sleep(4 * time.Second)
	//fmt.Println(time.Now())
	//ch2 := timer.C
	//fmt.Println(<-ch2)

	//fmt.Println("--------------------")
	//timer2 := time.NewTimer(5 * time.Second)
	//go func() {
	//	<-timer2.C
	//	fmt.Println("timer2 end...")
	//}()
	//time.Sleep(3 * time.Second)
	//stop := timer2.Stop()
	//if stop {
	//	fmt.Println("timer2 stop...")
	//}

	fmt.Println("--------------------")
	ch1 := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch1)
	fmt.Println(time.Now())
	time.Sleep(4 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-ch1)
}
