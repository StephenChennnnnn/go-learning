package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	cfg, err := ini.Load("my.ini")
	if err != nil {
		fmt.Printf("fail to read file: %v", err)
		os.Exit(1)
	}

	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("app mode: ", cfg.Section("").Key("app_mode").String())
	fmt.Println("app mode: ", cfg.Section("path").Key("data").String())

	// 做一些候选值限制的操作
	fmt.Println("server protocal: ", cfg.Section("server").Key("protocal").In("http", []string{"http", "https"}))
	fmt.Println("email protocal: ", cfg.Section("server").Key("protocal").In("smtp", []string{"imap", "smtp"}))

	// 自动类型转换
	fmt.Printf("port number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("enforce domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 修改值进行保存
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.SaveTo("my.ini.local")
}
