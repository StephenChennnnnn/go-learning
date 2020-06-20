package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Info("first beego exp...")
	beego.Run("localhost:8080")
}
