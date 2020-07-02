package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"wechat": "cxy", "blog": "www.cxy.com"})
	})
	r.GET("/xmlstruct", func(c *gin.Context) {
		c.XML(200, User{ID: 123, Name: "cxy", Age: 13})
	})
	r.GET("/xmlusers", func(c *gin.Context) {
		allUsers := []User{{ID: 123, Name: "zs", Age: 20}, {ID: 456, Name: "ls", Age: 25}}
		c.XML(200, gin.H{"users": allUsers})
	})

	r.Run(":8080")
}

type User struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}
