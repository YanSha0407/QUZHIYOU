package controllers

import (
	"QUZHIYOU/models"
	"QUZHIYOU/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ActivityHomeListController struct {
	beego.Controller
}

var dbmain *gorm.DB

func init() {

	db, err := gorm.Open("mysql",
		"root:loveys1314@tcp(127.0.0.1:3306)/qzy_official_service?charset=utf8&parseTime=True&loc=Local")
	dbmain = db
	dbmain.DB().SetMaxIdleConns(10)
	dbmain.DB().SetMaxOpenConns(300)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection succedssed")
	}

	dbmain.SingularTable(true)
	dbmain.AutoMigrate(&models.TbActivity{}, &models.TbBanner{})
	dbmain.LogMode(true)
}

//获取首页列表

func (this *ActivityHomeListController) ActivityList() {

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

	var activitys []*models.TbActivity

	dbmain.Select("ACTIVITY_ID,ACTIVITY_NAME,SUB_NAME,IMAGE,ORIGINAL_PRICE,TOTAL_NUM,PRICE_TAG,PRICE,STATUS").Limit(intsize).Offset(start).Find(&activitys)

	utils.ReturnHTTPSuccess(&this.Controller, &activitys)

	this.ServeJSON()

}

type ActivityInfoJson struct {
	ActivityInfo models.TbActivity  `json:"activityInfo"`
	Banner       []*models.TbBanner `json:"banner"`
}

//获取活动详情信息
func (this *ActivityHomeListController) ActivityInfo() {

	activityId := this.GetString("activityId")
	i64, _ := strconv.ParseInt(activityId, 10, 64)

	activityInfo := models.TbActivity{ActivityId: i64}

	dbmain.First(&activityInfo)
		activityInfo.SignStartTime=activityInfo.SignStartTime[:10]
		activityInfo.SignEndTime=activityInfo.SignEndTime[:10]

	var banners []*models.TbBanner

	dbmain.Where("ACTIVITY_ID=?", i64).Find(&banners)

	utils.ReturnHTTPSuccess(&this.Controller, ActivityInfoJson{ActivityInfo: activityInfo, Banner: banners})

	this.ServeJSON()

}

//back

//func (this *ActivityHomeListController) ActivityInfoback() {
//
//	o:=orm.NewOrm()
//
//	activityId := this.GetString("activityId")
//
//	i64,_ := strconv.ParseInt(activityId,10,64)
//
//	info:=models.TbActivity{ActivityId:i64}
//
//
//
//	o.QueryTable("tb_activity").Filter("ActivityId",i64).One(&info)
//
//
//
//	o.LoadRelated(&info,"Welfares")
//
//	o.LoadRelated(&info,"Addfroms")
//	o.LoadRelated(&info,"Addtos")
//
//
//
//	info.SignStartTime=info.SignStartTime[:10]
//	info.SignEndTime=info.SignEndTime[:10]
//
//	fmt.Println(info.Addfroms)
//
//
//	var add []*models.TbAddress
//	var addto []*models.TbAddress
//
//	for _,v:=range info.Addfroms{
//		//1出发地 2目的地
//		if v.Type ==1{
//			add = append(add, v)
//		}else if v.Type==2{
//			addto=append(addto,v)
//		}
//	}
//
//	info.Addfroms=add
//	info.Addtos=addto
//
//
//	//插入banner
//	var banners []models.TbBanner
//	o.QueryTable("tb_banner").Filter("ACTIVITYID",i64).All(&banners)
//
//
//
//	utils.ReturnHTTPSuccess(&this.Controller,ActivityInfoJson{ActivityInfo:info,Banner:banners})
//
//	this.ServeJSON()
//
//
//}
