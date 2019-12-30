package wxLogin

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/medivhzhan/weapp/v2"
	"os"
)

type User struct {
	Code string `form:"code" json:"code"`
}

func (u *User) WxUserLogin() serializer.Response {

	var user models.TbUser


	//	获取code 查询数据库 没有需创建用户

	res, err := weapp.Login(os.Getenv("WXAPP_ID"), os.Getenv("WXSECRET"), u.Code)

	if err != nil {
		// 处理一般错误信息
		return serializer.Response{
			Code:  0,
			Data:  nil,
			Msg:   "2",
			Error: "",
		}
	}

	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		return serializer.Response{
			Code:  0,
			Data:  nil,
			Msg:   "1",
			Error: "",
		}
	}

	isHave := models.DB.Model(&user).Where("OPEN_ID=?", res.OpenID).First(&user).RecordNotFound()

	if isHave {
		//获取token
		token, _ := weapp.GetAccessToken(os.Getenv("WXAPP_ID"), os.Getenv("WXSECRET"))

		//	创建用户
		oneUser := models.TbUser{OpenId: res.OpenID, SessionKey: res.SessionKey,AccessToken:token.AccessToken}


		models.DB.Save(&oneUser)

		return serializer.Response{
			Msg:   "创建用户",
			Error: "",
			Data:  serializer.BuildUser(oneUser),
		}
	} else {
		return serializer.Response{
			Data: serializer.BuildUser(user),
		}
	}

}
