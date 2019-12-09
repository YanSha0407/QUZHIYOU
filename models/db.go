package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)



type Tb_activity struct {
	ActivityId    int64     `orm:"pK;column(ACTIVITY_ID)" json:"activityId"`     // 活动ID
	ActivityName  string    ` json:"activityName"` // 活动名称
	SubName       string    `json:"subName"`
	Tags          string    ` json:"tags"`                     // 首页标志：#免费 #旅游 #香山 #吃喝玩乐 #一日游
	TagsInfo      string    ` json:"tagsInfo"`           // 详情页标志：'河南同乡','免费一日游','吃喝全包','客车接送'
	Price         int64     ` json:"price"`                   // 现价
	PriceTag      string    ` json:"priceTag"`           // 首届
	OriginalPrice int64     ` json:"originalPrice"` // 原价
	Image         string    ` json:"image"`
	MemNum        int64     ` json:"memNum"`                 // 已报名人数
	TotalNum      int64     ` json:"totalNum"`             // 计划总人数
	CollectionNum int64     ` json:"collectionNum"`   // 收藏数
	Status        string    ` json:"status"`                   // 活动状态：未开始/进行中/已结束
	Author        string    ` json:"author"`                   // 创建者
	DateAdd       time.Time ` json:"dateAdd"`               // 活动创建时间
	DateUpdate    time.Time ` json:"dateUpdate"`         // 活动更新时间
	SignStartTime  string ` json:"signStartTime"` // 报名开始时间
	SignEndTime   string ` json:"signEndTime"`     // 报名结束时间
	ActiveTime    time.Time ` json:"active_time"`         // 活动时间
	Views         int64     ` json:"views"`                     // 浏览量
	HtmlCon       string    ` json:"htmlCon"`

	Welfares      []*Tb_welfare    `orm:"reverse(many)" json:"welfareList"`  //多个福利

	Addfroms      []*Tb_address    `orm:"reverse(many)"  json:"addfroms"`


}

type Tb_welfare struct {
	WelfareID  int64  `orm:"pK;column(WELFARE_ID)" json:"welfareId"`   // 福利ID
	Tag        string ` json:"tag"`                 // 福利类型
	Des       string ` json:"des"`                 // 福利详细描述
	ActivityId int64  ` orm:"column(ACTIVITY_ID)" json:"activityId"` // 活动ID

	Activity  *Tb_activity  `orm:"rel(fk)" json:"activity"`
}

type Tb_address struct {
	AddressId   int64   `orm:"pK;column(ADDRESS_ID)" json:"addressId"`     // 地址ID
	AddressName string  `orm:"column(ADDRESS_NAME)" json:"addressName"` // 地址
	Address     string  `orm:"column(ADDRESS)" json:"address"`           // 地址详情
	Type        int64   `orm:"column(TYPE)" json:"type"`                 // 地址类型：1-集合地/2-目的地
	Lat        float64  `orm:"column(LAT)" json:"lat"`                   // 纬度
	Lng         float64 `orm:"column(LNG)" json:"lng"`                   // 经度
	ActivityId  int64   `orm:"column(ACTIVITY_ID)" json:"activityId"`   // 活动ID

	Activity    *Tb_activity  `orm:"rel(fk)" orm:"-" json:"activity"`


}

type Tb_banner struct {
	BannerID   int64  `orm:"pK;column(BANNER_ID)" json:"banner_id"`
	ActivityId int64  `orm:"column(ACTIVITY_ID)" json:"activity_id"` // 活动ID
	Url        string `orm:"column(URL)" json:"url"`                 // 图片URL
}







type tb_enroll struct {
	ENROLLID       int64  `orm:"enroll_id" json:"enroll_id"`               // 报名表ID
	ACTIVITYID     int64  `orm:"activity_id" json:"activity_id"`           // 活动ID
	USERID         int64  `orm:"user_id" json:"user_id"`                   // 用户ID
	CONTACTSIDCARD string `orm:"contacts_id_card" json:"contacts_id_card"` // 报名用户身份证号
	PHONE          string `orm:"phone" json:"phone"`                       // 联系人手机号
}

type tb_evaluate struct {
	EVALUATEID int64     `orm:"evaluate_id" json:"evaluate_id"` // 评价ID
	USERID     int64     `orm:"user_id" json:"user_id"`         // 用户ID
	ACTIVITYID int64     `orm:"activity_id" json:"activity_id"` // 活动ID
	AVA        string    `orm:"ava" json:"ava"`
	CONTENT    string    `orm:"content" json:"content"`   // 评价内容
	DATEADD    time.Time `orm:"date_add" json:"date_add"` // 评价时间
}



type tb_collection struct {
	COLLECTIONID int64 `orm:"collection_id" json:"collection_id"` // 收藏表ID
	USERID       int64 `orm:"user_id" json:"user_id"`             // 用户ID
	ACTIVITYID   int64 `orm:"activity_id" json:"activity_id"`     // 活动ID
}

type tb_contacts struct {
	CONTACTSID     int64  `orm:"contacts_id" json:"contacts_id"`           // 联系人表ID
	CONTACTSNAME   string `orm:"contacts_name" json:"contacts_name"`       // 联系人姓名
	CONTACTSIDCARD string `orm:"contacts_id_card" json:"contacts_id_card"` // 联系人身份证号
	CONTACTSSEX    int64  `orm:"contacts_sex" json:"contacts_sex"`         // 联系人性别：1-男/2-女
	USERID         int64  `orm:"user_id" json:"user_id"`                   // 用户ID
}

type tb_user struct {
	USERID     int64  `orm:"user_id" json:"user_id"`
	OPENID     string `orm:"open_id" json:"open_id"`         // 用户唯一标识
	SESSIONKEY string `orm:"session_key" json:"session_key"` // 会话密钥
	UNIONID    string `orm:"union_id" json:"union_id"`       // 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回
	USERNAME   string `orm:"user_name" json:"user_name"`     // 用户姓名
	USERFROM   string `orm:"user_from" json:"user_from"`     // 用户来自哪里
}






func init() {

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123qaz!@#@tcp(39.97.230.148:3306)/qzy_official_service?charset=utf8", 30)

	fmt.Println("连接成功")

	// register model
	orm.RegisterModel(new(Tb_activity))
	orm.RegisterModel(new(Tb_welfare))
	orm.RegisterModel(new(Tb_banner))
	orm.RegisterModel(new(Tb_address))


}
