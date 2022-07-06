package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey GO URL SHORTNER!",
		})

	})
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web s3rv3r - ERR:%v0", err))
	}
}
