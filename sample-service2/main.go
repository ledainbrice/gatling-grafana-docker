package main

import "github.com/gin-gonic/gin"
import "fmt"

var counter int = 0

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		counter = counter + 1
		msg := fmt.Sprintf("hello world! %d", counter)
		fmt.Printf(msg)
		c.JSON(200, gin.H{
			"message": msg,
		})
	})
	r.Run(":4000")
}