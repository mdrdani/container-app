package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	hostname, _ := os.Hostname()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "api server",
			"server":  hostname,
			"author": "M Dani Ramanda",
		})
	})
	r.Run("0.0.0.0:8144")
}
