package main

import (
	_ "encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "cxy"})
	})

	r.GET("/users/123", func(c *gin.Context) {
		c.JSON(200, user{ID: 123, Name: "cxy", Age: 25})
	})
	r.GET("/users", func(c *gin.Context) {
		allUsers := []user{{ID: 123, Name: "zs", Age: 12}, {ID: 456, Name: "ls", Age: 13}}
		c.IndentedJSON(200, allUsers)
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "<b>hello</b>"})
	})
	r.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{"message": "<b>hello</b>"})
	})
	r.GET("/asciiJson", func(c *gin.Context) {
		c.AsciiJSON(200, gin.H{"message": "hello 哈啦嘿"})
	})

	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{"wechat": "cxy"})
	})

	r.Run(":8080")
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
