package controllers

import (
	"QUZHIYOU/models"
	"QUZHIYOU/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type ActivityHomeListController struct {
	beego.Controller
}

func updateJsonKeysTopic(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "ActivityId":
				delete(val, k)
				val["activityId"] = v
			case "ActivityName":
				delete(val, k)
				val["activityName"] = v
			case "SubName":
				delete(val, k)
				val["subName"] = v
			case "Tags":
				delete(val, k)
				val["tags"] = v
			case "TotalNum":
				delete(val, k)
				val["totalNum"] = v
			case "Image":
				delete(val, k)
				val["image"] = v
			case "OriginalPrice":
				delete(val, k)
				val["originalPrice"] = v
			case "SignStartTime":
				delete(val, k)
				val["signStartTime"] = v
			case "SignEndTime":
				delete(val, k)
				val["signEndTime"] = v
			case "Price":
				delete(val, k)
				val["price"] = v
			case "Status":
				delete(val, k)
				val["status"] = v
			}
		}
	}
}

//获取首页列表

 func (this *ActivityHomeListController) ActivityList() {

	o := orm.NewOrm()

	page := this.GetString("page")
	size := this.GetString("size")

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

	var activitys []orm.Params

	o.QueryTable("tb_activity").
		OrderBy("-ActivityId").
		Limit(intsize, start).
		Values(&activitys, "ActivityId", "ActivityName", "SubName", "Tags", "TotalNum", "Image", "OriginalPrice", "SignStartTime", "SignEndTime", "Price", "Status")


	updateJsonKeysTopic(activitys)



	utils.ReturnHTTPSuccess(&this.Controller, &activitys)

	this.ServeJSON()

}


type ActivityInfoJson struct {

	ActivityInfo  models.TbActivity 	`json:"activityInfo"`
	Banner        []models.TbBanner    `json:"banner"`
}




//获取活动详情信息
func (this *ActivityHomeListController) ActivityInfo() {

	o:=orm.NewOrm()

	activityId := this.GetString("activityId")

	i64,_ := strconv.ParseInt(activityId,10,64)

	info:=models.TbActivity{ActivityId:i64}



	o.QueryTable("tb_activity").Filter("ActivityId",i64).One(&info)




	o.LoadRelated(&info,"Welfares")

	o.LoadRelated(&info,"Addfroms")



	info.SignStartTime=info.SignStartTime[:10]
	info.SignEndTime=info.SignEndTime[:10]


	for k,v:=range info.Welfares{
		fmt.Println(k,v)
	}




	//插入banner
	var banners []models.TbBanner
	o.QueryTable("tb_banner").Filter("ACTIVITYID",i64).All(&banners)











	utils.ReturnHTTPSuccess(&this.Controller,ActivityInfoJson{ActivityInfo:info,Banner:banners})

	this.ServeJSON()


}
