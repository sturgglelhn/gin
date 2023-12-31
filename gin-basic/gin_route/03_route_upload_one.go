package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
上传单个文件
*/

func main() {
	r := gin.Default()
	// 限制上传最大的尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	r.Run()
}
