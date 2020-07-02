package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/html", func(c *gin.Context) {
		c.Status(200)
		const templateText = `wechat: {{printf "%s" .}}`
		tmpl, err := template.New("htmlTest").Parse(templateText)
		if err != nil {
			log.Fatal("parsing: %s", err)
		}
		tmpl.Execute(c.Writer, "cxy")
	})

	r.SetFuncMap(template.FuncMap{
		"md5": MD5,
	})
	r.LoadHTMLFiles("html/index.html")
	r.GET("/htmldir", func(c *gin.Context) {
		c.HTML(200, "index.html", "cxy")
	})

	r.Run(":8080")
}

func MD5(in string) (string, error) {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:]), nil
}
