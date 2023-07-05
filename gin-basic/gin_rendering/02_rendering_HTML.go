package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin支持加载HTML模版，然后根据模版参数进行配置并返回相应的数据，本质上就是字符串替换
LoadHTMLGlob()方法可以加载模版文件
*/
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("gin_rendering/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "html/index.html", gin.H{"title": "我是测试", "ce": "123456"})
	})
	r.Run()
}
