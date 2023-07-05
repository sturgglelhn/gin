package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

/*
goroutine机智可以方便地实现异步处理
另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
*/

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(),Recovery()
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {
		// 只需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})

	// 2.同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})
	r.Run()
}
