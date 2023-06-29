package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*表单参数*/

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	// 默认是8080
	r.Run()
}
