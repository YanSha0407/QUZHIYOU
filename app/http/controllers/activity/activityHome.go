package activity

import (
	"QUZHIYOU/app/models"
	"QUZHIYOU/app/response"
	"QUZHIYOU/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//获取首页列表

func ActivityList(this *gin.Context) {

	page := this.Query("page")
	size := this.Query("size")

	var intpage, intsize int
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

	models.DB.
		Debug().
		Preload("Welfares").
		Select("ACTIVITY_ID,ACTIVITY_NAME,SUB_NAME,IMAGE,ORIGINAL_PRICE,TOTAL_NUM,PRICE_TAG,PRICE,STATUS,TAGS").
		Limit(intsize).
		Offset(start).
		Find(&activitys)

	// 响应
	resGin := response.Gin{C: this}

	resGin.Success("ok", &activitys)

}

type ActivityInfoJson struct {
	ActivityInfo models.TbActivity  `json:"activityInfo"`
	Banner       []*models.TbBanner `json:"banner"`
}

//获取活动详情信息
func ActivityInfo(this *gin.Context) {

	activityId := this.Query("activityId")
	i64, _ := strconv.ParseInt(activityId, 10, 64)

	activityInfo := models.TbActivity{
		ActivityId: i64,
	}

	models.DB.
		Debug().
		Preload("Welfares").
		Preload("AddressFrom").
		Preload("AddressTo").
		First(&activityInfo)
	activityInfo.SignStartTime = activityInfo.SignStartTime[:10]
	activityInfo.SignEndTime = activityInfo.SignEndTime[:10]

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

	models.DB.Where("ACTIVITY_ID=?", i64).Find(&banners)

	this.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data":    &ActivityInfoJson{ActivityInfo: activityInfo, Banner: banners},
	})

}
