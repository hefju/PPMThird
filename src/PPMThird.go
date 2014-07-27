package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
			c.String(200, "Hello, I'm PPMThird.")

		})

	r.GET("/user/:name", func(c *gin.Context) {
			name := c.Params.ByName("name")
			message := "Hello " + name
			//c.String(200, message)
//			c.XML(200,message)
			c.JSON(200,message)

		})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}
