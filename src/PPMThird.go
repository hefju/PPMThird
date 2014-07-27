package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLTemplates("templates/*")
	r.GET("/", func(c *gin.Context) {
//			c.String(200, "Hello, I'm PPMThird. power by gin")
			obj := gin.H{"name": "Main website","xyz":"title for con","abc":"一株"}
			c.HTML(200,"index.tmpl", obj)
		})

	r.GET("/user/:name", func(c *gin.Context) {
			name := c.Params.ByName("name")
			message := "Hello " + name
			//c.String(200, message)
			//c.XML(200,message)
			c.JSON(200,message)
		})

	r.POST("/login", func(c *gin.Context) {
			var json LoginJSON

			// If EnsureBody returns false, it will write automatically the error
			// in the HTTP stream and return a 400 error. If you want custom error
			// handling you should use: c.ParseBody(interface{}) error
			if c.EnsureBody(&json) {
				if json.User=="manu" && json.Password=="123" {
					c.JSON(200, gin.H{"status": "you are logged in"})
				}else{
					c.JSON(401, gin.H{"status": "unauthorized"})
				}
			}
		})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
	
}

type LoginJSON struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
