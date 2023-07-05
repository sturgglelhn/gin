package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
1.模拟实现权限验证中间件
	·有2个路由，login和home
	·login用于设置cookie
	·home是访问查看信息的请求
	·在请求home之前，先跑中间件代码，检验是否存在cookies
2.访问home，会显示错误，因为权限校验未通过
3.Cookie的缺点
	·不安全，明文
	·增减带宽消耗
	·可以被禁用
	·cookie有上限
*/

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}

		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不在调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	// 1.创建路由
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8000")
}
