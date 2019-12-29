package qrcode

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"github.com/medivhzhan/weapp/v2"
)


var token string

func init() {
	res, err := weapp.GetAccessToken("wx27f763c834a6e9aa", "2ac10ca85c83589fde470f990d6ad02e")
	if err != nil {

		return
	}
	if err := res.GetResponseError(); err != nil {

		return
	}

	token = res.AccessToken
}

func Getqrcode(c *gin.Context) {
	getter := weapp.UnlimitedQRCode{
		Scene:     "id=1",
		Page:      "homeSub/pages/homeDetail/homeDetail",
		Width:     430,
		AutoColor: true,
		LineColor: weapp.Color{"255", "255", "255"},
		IsHyaline: false,
	}

	resp, res, err := getter.Get(token)
	if err != nil {
		return
	}
	if err := res.GetResponseError(); err != nil {
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	f,_:=os.Create("code.png")

	defer f.Close()

	f.Write(content)

}