package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir, err := ioutil.TempDir("/Users/chenxinyuan/TTTT", "Test")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(dir) // 用完删除
	fmt.Printf("%s\n", dir)

	// 创建临时文件
	f, err := ioutil.TempFile(dir, "Test")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(f.Name()) // 用完删除
	fmt.Printf("%s\n", f.Name())

	dirname := "/Users/chenxinyuan/TTTT"
	listFiles(dirname, 0)
}

func listFiles(dirName string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|   " + s
	}

	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirName + "/" + fi.Name()
		fmt.Println("%s%s\n", s, filename)
		if fi.IsDir() {
			listFiles(filename, level+1)
		}
	}
}
