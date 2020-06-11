package main

import (
	beego "github.com/astaxie"
	"github.com/astaxie/logs"
)

func main() {
	logs.Info("first beego exp...")
	beego.Run("localhost:8080")
}
