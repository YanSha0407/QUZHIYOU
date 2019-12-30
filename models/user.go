package models



//用户
type TbUser struct {
	UserId   int64  ` gorm:"primary_key;column:USER_ID" json:"userId" `
	OpenId string  ` gorm:"column:OPEN_ID" json:"openId" `
	SessionKey       string `  gorm:"column:SESSION_KEY" json:"sessionKey"`
	UnionId       string `  gorm:"column:UNION_ID" json:"unionId"`
	UserName       string `  gorm:"column:USER_NAME" json:"userName"`
	UserFrom      string `  gorm:"column:USER_FROM" json:"userFrom"`
	AccessToken      string `  gorm:"column:ACCESS_TOKEN" json:"accessToken"`
}