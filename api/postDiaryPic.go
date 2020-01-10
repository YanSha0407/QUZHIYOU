package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func PostDiaryPic(c *gin.Context)  {
	//获取文件路径
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {

		log.Println(file.Filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dir + "/static/"+file.Filename)
	}
	c.IndentedJSON(200,gin.H{
		"code":0,
		"data":len(files),
		"msg":"ok",

	})

}