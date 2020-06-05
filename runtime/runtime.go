package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("GOROOT --> ", runtime.GOROOT())
	fmt.Println("os/platform --> ", runtime.GOOS)

	fmt.Println("逻辑CPU核数目 --> ", runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	//go func() {
	//	for i := 0; i < 5; i++ {
	//		fmt.Println("goroutine...")
	//		time.Sleep(time.Second)
	//	}
	//}()
	//for i := 0; i < 4; i++ {
	//	runtime.Gosched()
	//	fmt.Println("main..")
	//	time.Sleep(time.Second)
	//}
	//time.Sleep(5 * time.Second)
	//
	//go func() {
	//	fmt.Println("goroutine start...")
	//	time.Sleep(time.Second)
	//	fun()
	//	fmt.Println("goroutine end...")
	//}()
	//time.Sleep(3 * time.Second)
	//
	//a := 1
	//go func() {
	//	a = 2
	//	fmt.Println("子goroutine。。", a)
	//}()
	//a = 3
	//time.Sleep(1)
	//fmt.Println("main goroutine。。", a)

	//wg.Add(4)
	//go saleTickets("seller1")
	//go saleTickets("seller2")
	//go saleTickets("seller3")
	//go saleTickets("seller4")
	//wg.Wait()

	wg.Add(4)
	go writeData(1)
	go readData(2)
	go readData(4)
	go writeData(3)

	wg.Wait()
}

func fun() {
	defer fmt.Println("defer...")
	runtime.Goexit()
	fmt.Println("func fun...")
}

var wg sync.WaitGroup

var ticket = 10 // 100 tickets
var mutex sync.Mutex

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, " sell: ", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "sell out...")
			break
		}
		mutex.Unlock()
	}
}

var rwMutex sync.RWMutex

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "write start...")
	rwMutex.Lock()
	fmt.Println(i, "writing...")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "write end...")
}
func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "read start...")
	rwMutex.RLock()
	fmt.Println(i, "reading...")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock()
	fmt.Println(i, "read end...")
}
