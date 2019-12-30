package activity


import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

// List 活动列表
func (service *ListVideoService) List() serializer.Response {


	activitys := []models.TbActivity{}

	total := 0

	if service.Size == 0 {
		service.Size = 1
	}
	if service.Page == 0 {
		service.Page = 1
	}

	start := (service.Page - 1) * service.Size

	if err := models.DB.Model(models.TbActivity{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	showRows:=[]string{"ACTIVITY_ID","ACTIVITY_NAME","SUB_NAME","IMAGE","ORIGINAL_PRICE","TOTAL_NUM","PRICE_TAG","PRICE","STATUS","TAGS"}
	if err := models.DB.Limit(service.Size).Offset(start).Find(&activitys).Select(showRows).Error; err != nil {
		return serializer.Response{
			Code: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildActivitys(activitys), uint(total))
}
