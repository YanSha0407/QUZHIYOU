package diary

import (
	"QUZHIYOU/app/response"
	"QUZHIYOU/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func HomeList(c *gin.Context)  {

	resGin := response.Gin{C: c}

	var diary services.DiaryHomeService

	if err:=c.ShouldBindJSON(&diary);err!=nil{
		fmt.Println(err)
		resGin.Error("获取失败","获取动态失败")
	}else {
		diarys:=diary.GetAllDiary()
		resGin.Success("ok",&diarys)
	}








}