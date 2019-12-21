package diary

import (
	"QUZHIYOU/app/response"
	"QUZHIYOU/services"
	"github.com/gin-gonic/gin"
)

func HomeList(c *gin.Context)  {

	resGin := response.Gin{C: c}

	var diary services.DiaryHomeService

	if err:=c.ShouldBind(&diary);err!=nil{
		resGin.Error("获取失败","获取动态失败")
	}else {
		diarys:=diary.GetAllDiary()
		resGin.Success("ok",&diarys)
	}








}