package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	/*
		seek(offset,whence),设置指针光标的位置
		第一个参数：偏移量
		第二个参数：如何设置
			0：seekStart表示相对于文件开始，
			1：seekCurrent表示相对于当前偏移量，
			2：seek end表示相对于结束。


		const (
			SeekStart   = 0 // seek relative to the origin of the file
			SeekCurrent = 1 // seek relative to the current offset
			SeekEnd     = 2 // seek relative to the end
		)

		随机读取文件：
			可以设置指针光标的位置
	*/

	//file, _ := os.OpenFile("/Users/chenxinyuan/abcd.txt", os.O_RDWR, 0)
	//defer file.Close()
	//bs := []byte{0}
	//
	//file.Read(bs)
	//fmt.Println(string(bs))
	//
	//file.Seek(4, io.SeekStart)
	//file.Read(bs)
	//fmt.Println(string(bs))
	//file.Seek(2, io.SeekStart)
	//file.Read(bs)
	//fmt.Println(string(bs))
	//
	//file.Seek(3, io.SeekCurrent)
	//file.Read(bs)
	//fmt.Println(string(bs))
	//
	//file.Seek(0, io.SeekEnd)
	//file.WriteString("ABC")

	// 断点续传
	srcFile := "/Users/chenxinyuan/abcd.txt"
	destFile := "/Users/chenxinyuan/abc.txt"
	tempFile := destFile + "_temp"
	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm)

	defer file1.Close()
	defer file2.Close()
	// 1.读临时文件中读数据，根据 seek
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100)
	n1, err := file3.Read(bs)
	fmt.Println(n1)
	countStr := string(bs[:n1])
	fmt.Println(countStr)
	count, _ := strconv.ParseInt(countStr, 10, 64)
	fmt.Println(count)

	// 2.设置读写读偏移量
	file1.Seek(count, 0)
	file2.Seek(count, 0)
	data := make([]byte, 2)
	n2 := -1 // 读的偏移量
	n3 := -1 // 写的偏移量
	total := int(count)

	for {
		// 3.读数据
		n2, err = file1.Read(data)
		if err == io.EOF {
			fmt.Println("succeed")
			file3.Close()
			//os.Remove(tempFile)
			break
		}
		// 将数据写到目标文件
		n3, _ = file2.Write(data[:n2])
		total += n3
		// 将复制总量，存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		// 假装断电
		if total > 10 {
			panic("duan dian ")
		}
	}
}
