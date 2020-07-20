package main

import "github.com/gin-gonic/gin"
import "fmt"
import "os"

var counter int = 0

func main() {
	r := gin.Default()
	errlogfile, _ := os.Create("error.log")
	gin.DefaultErrorWriter = errlogfile
	r.GET("/ping", func(c *gin.Context) {
		errLogger := gin.DefaultErrorWriter
        errLogger.Write([]byte("[DIY ERROR] HELLO ERROR!\n"))
		counter = counter + 1
		msg := fmt.Sprintf("hello world! %d", counter)
		fmt.Printf(msg)
		c.JSON(200, gin.H{
			"message": msg,
		})
	})
	r.Run(":4000")
}