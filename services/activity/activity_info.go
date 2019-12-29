package activity

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"QUZHIYOU/utils"
)

//活动详情data
type ActivityInfo struct {


	ActivityId string `form:"activityId" json:"activityId"`

}



func (activityInfo *ActivityInfo) GetActivityInfo(ActivityId string) serializer.Response {



	actiInfo := models.TbActivity{
		ActivityId: utils.String2Int64(ActivityId),
	}

	var banners []*models.TbBanner


	models.DB.Find(&banners,"ACTIVITY_ID=?",ActivityId)


	err :=models.DB.
		Debug().
		Preload("Welfares").
		Preload("AddressFrom").
		Preload("AddressTo").
		First(&actiInfo).Error
	if err != nil {
		return serializer.Response{
			Code: 404,
			Msg:    "活动不存在",
			Error:  err.Error(),
		}
	}


	return serializer.ActivityInfoResponse(actiInfo,&banners)

}