package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//
	var a int = 12
	var b float64 = 12
	var c interface{} = a
	d := 12 // will be an int
	fmt.Printf("a=(%T,%v)\n", a, a)
	fmt.Printf("b=(%T,%v)\n", b, b)
	fmt.Printf("c=(%T,%v)\n", c, c)
	fmt.Printf("d=(%T,%v)\n", d, d)
	useInt(12)
	useFloat(12)
	fmt.Println("a==12:", a == 12) // true
	fmt.Println("b==12:", b == 12) // true
	fmt.Println("c==12:", c == 12) // true
	fmt.Println("a==c:", a == c)   // true
	fmt.Println("b==c:", b == c)   // false

	//
	var e float64 = 12
	var f interface{} = e
	fmt.Println("f==12:", f == 12)
	fmt.Printf("f=(%T,%v)\n", f, f)
	fmt.Printf("hard-coded=(%T,%v)\n", 12, 12)

	//
	var i *int = nil
	var cc interface{}
	if i == nil {
		cc = nil
	}
	fmt.Printf("cc=(%T,%v)\n", cc, cc)
	fmt.Println("cc==nil: ", cc == nil)
	bbm := BBMgr{}
	t, y := bbm.QueryAA(1)
	fmt.Println(y)
	fmt.Printf("t=(%T,%v)\n", t, t)
	fmt.Println("t==nil: ", t == nil)
	t, y = bbm.QueryAA(3)
	fmt.Println(y)
	fmt.Printf("t=(%T,%v)\n", t, t)
	fmt.Println("t==nil: ", t == nil)

	//
	tty, _ := os.OpenFile("E:\\Task.txt", os.O_RDWR, 0)
	var r io.Reader
	r = tty
	var w io.Writer
	w = r.(io.Writer)
	var v interface{} = tty
	fmt.Printf("tty=(%T,%v)\n", tty, tty)
	fmt.Printf("r=(%T,%v)\n", r, r)
	fmt.Printf("w=(%T,%v)\n", w, w)
	fmt.Printf("v=(%T,%v)\n", v, v)
}

func useInt(n int) {
	fmt.Printf("useInt=(%T,%v)\n", n, n)
}

func useFloat(n float64) {
	fmt.Printf("useFloat=(%T,%v)\n", n, n)
}

type AAer interface {
	That() int
}
type BB struct {
}
type BBMgr struct {
	bbMap map[int]AAer
}

func (this *BB) That() int {
	return 0
}
func (this *BBMgr) QueryAA(id int) (AAer, bool) {
	aa, succeed := this.bbMap[id]
	//这里如果没有找到数据 返回的接口去做 nil比较的话
	//发现 AAer永远不为nil，但是执行AAer的函数就报错
	// 实际上这里应该写
	// aa, ok :=this.bbMap[id]
	// if !ok {
	// return nil
	// }
	return aa, succeed
}
