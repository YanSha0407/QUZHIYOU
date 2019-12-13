package models

import (
	_ "github.com/go-sql-driver/mysql"
)



type TbActivity struct {
	ActivityId    int64     ` gorm:"primary_key;column:ACTIVITY_ID" json:"activityId"`     // 活动ID
	ActivityName  string    ` json:"activityName,omitempty" gorm:"column:ACTIVITY_NAME"` // 活动名称
	SubName       string    `json:"subName,omitempty" gorm:"column:SUB_NAME"`
	Tags          string    ` json:"tags,omitempty" gorm:"column:TAGS"`                     // 首页标志：#免费 #旅游 #香山 #吃喝玩乐 #一日游
	TagsInfo      string    ` json:"tagsInfo,omitempty" gorm:"column:TAGS_INFO"`           // 详情页标志：'河南同乡','免费一日游','吃喝全包','客车接送'
	Price         int64     ` json:"price" gorm:"column:PRICE"`                   // 现价
	PriceTag      string    ` json:"priceTag,omitempty" gorm:"column:PRICE_TAG"`           // 首届
	OriginalPrice int64     ` json:"originalPrice,omitempty" gorm:"column:ORIGINAL_PRICE"` // 原价
	Image         string    ` json:"image,omitempty" gorm:"column:IMAGE"`
	MemNum        int64     ` json:"memNum,omitempty" gorm:"column:MEM_NUM"`                 // 已报名人数
	TotalNum      int64     ` json:"totalNum,omitempty" gorm:"column:TOTAL_NUM"`             // 计划总人数
	CollectionNum int64     ` json:"collectionNum,omitempty" gorm:"column:COLLECTION_NUM"`   // 收藏数
	Status        string    ` json:"status,omitempty" gorm:"column:STATUS"`                   // 活动状态：未开始/进行中/已结束
	Author        string    ` json:"author,omitempty" gorm:"column:AUTHOR"`                   // 创建者
	DateAdd       string ` json:"dateAdd,omitempty" gorm:"column:DATE_ADD"`               // 活动创建时间
	DateUpdate    string ` json:"dateUpdate,omitempty" gorm:"column:DATE_UPDATE"`         // 活动更新时间
	SignStartTime  string   ` json:"signStartTime,omitempty" gorm:"column:SIGN_START_TIME"` // 报名开始时间
	SignEndTime   string    ` json:"signEndTime,omitempty" gorm:"column:SIGN_END_TIME"`     // 报名结束时间
	ActiveTime    string ` json:"active_time,omitempty" gorm:"column:ACTIVE_TIME"`         // 活动时间
	Views         int64     ` json:"views,omitempty" gorm:"column:VIEWS"`                     // 浏览量
	HtmlCon       string    ` json:"htmlCon,omitempty" gorm:"column:HTML_CON"`

//	关联字段
    Welfares      []*TbWelfare  `json:"welfareList,omitempty" gorm:"FOREIGNKEY:ActivityId;ASSOCIATION_FOREIGNKEY:ActivityId" `

	AddressFrom   []*TbAddress  `json:"gatherAddList,omitempty" gorm:"FOREIGNKEY:ActivityId;ASSOCIATION_FOREIGNKEY:ActivityId"`

	AddressTo     []*TbAddress   `json:"destinationList,omitempty"  gorm:"FOREIGNKEY:ActivityId;ASSOCIATION_FOREIGNKEY:ActivityId"`

}

type TbWelfare struct {
	WelfareID  int64  `gorm:"primary_key;column:WELFARE_ID" json:"welfareId"`   // 福利ID
	Tag        string `gorm:"column:TAG" json:"tag"`                 // 福利类型
	Des       string ` gorm:"column:DES"json:"des"`                 // 福利详细描述
	ActivityId int64  `gorm:"INDEX;column:ACTIVITY_ID" json:"-"` // 活动ID

}

type TbBanner struct {
	BannerID   int64  ` gorm:"primary_key;column:BANNER_ID" json:"banner_id"`
	ActivityId int64  ` gorm:"column:ACTIVITY_ID" json:"activityId"` // 活动ID
	URL       string `  gorm:"column:URL" json:"url"`                 // 图片URL
}


type TbAddress struct {
	AddressId   int64   `gorm:"primary_key;column:ADDRESS_ID" json:"addressId"`     // 地址ID
	AddressName string  `gorm:"column:ADDRESS_NAME" json:"addressName"` // 地址
	Address     string  `gorm:"column:ADDRESS" json:"address"`           // 地址详情
	Type        int64   `gorm:"column:TYPE" json:"type"`                 // 地址类型：1-集合地/2-目的地
	Lat        float64  `gorm:"column:LAT" json:"lat"`                   // 纬度
	Lng         float64 `gorm:"column:LNG" json:"lng"`                   // 经度
	ActivityId  int64   `gorm:"INDEX;column:ACTIVITY_ID" json:"-"`   // 活动ID

}






