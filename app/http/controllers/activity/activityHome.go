package activity

import (
	"QUZHIYOU/app/models"
	"QUZHIYOU/app/response"
	"QUZHIYOU/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取首页列表

func ActivityList(this *gin.Context) {

	page := this.Query("page")
	size := this.Query("size")

	var (
		intpage, intsize int

	)
	if page == "" {
		intpage = 1
	} else {
		intpage = utils.String2Int(page)
	}

	if size == "" {
		intsize = 10
	} else {
		intsize = utils.String2Int(size)
	}

	start := (intpage - 1) * intsize

	var activitys []*models.TbActivity


	showRows:=[]string{"ACTIVITY_ID","ACTIVITY_NAME","SUB_NAME","IMAGE","ORIGINAL_PRICE","TOTAL_NUM","PRICE_TAG","PRICE","STATUS","TAGS"}

	    models.DB.
		Select(showRows).
		Limit(intsize).
		Offset(start).
		Order("ACTIVITY_ID desc").
		Find(&activitys)


	// 响应
	resGin := response.Gin{C: this}
	resGin.Success("ok", &activitys)

}


//活动详情data
type ActivityInfoJson struct {
	ActivityInfo models.TbActivity  `json:"activityInfo"`
	Banner       []*models.TbBanner `json:"banner"`
}

//获取活动详情信息
func ActivityInfo(this *gin.Context) {

	//获取id转为int64
	activityId := utils.String2Int64(this.Query("activityId"))

	activityInfo := models.TbActivity{
		ActivityId: activityId,
	}

	models.DB.
		Debug().
		Preload("Welfares").
		Preload("AddressFrom").
		Preload("AddressTo").
		First(&activityInfo)


	activityInfo.FormatTime(&activityInfo)



	var add []*models.TbAddress
	var addto []*models.TbAddress

	for _, v := range activityInfo.AddressFrom {
		//1出发地 2目的地
		if v.Type == 1 {
			add = append(add, v)
		} else if v.Type == 2 {
			addto = append(addto, v)
		}
	}

	activityInfo.AddressFrom = add
	activityInfo.AddressTo = addto

	var banners []*models.TbBanner

	models.DB.Find(&banners,"ACTIVITY_ID=?",activityId)

	this.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data":    &ActivityInfoJson{ActivityInfo: activityInfo, Banner: banners},
	})

}
