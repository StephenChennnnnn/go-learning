package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	srcFile := "/Users/chenxinyuan/TTTT.txt"
	destFile := "TTTT.txt"
	total, err := copyFile1(srcFile, destFile)
	fmt.Println(err)
	fmt.Println(total)
}

func copyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	//拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("succeed")
			break
		} else if err != nil {
			fmt.Println("false")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil

}

func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2, file1)
}

func copyFile3(srcFile, destFile string) (int, error) {
	input, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		fmt.Println("false: ", destFile)
		fmt.Println(err)
		return 0, err
	}

	return len(input), nil
}
