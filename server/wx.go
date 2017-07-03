package main

import (
	"fmt"

	_ "github.com/bitly/go-simplejson"
)

type WXAPI struct {
	appKey    string
	appSecret string
}

// WX 微信
var WX WXAPI

func NewWXAPI(appKey, appSecret string) WXAPI {
	WX = WXAPI{appKey, appSecret}
	return WX
}

// GetAccessToken 通过Code 获取 access_token
func (wx *WXAPI) GetAccessToken(code string) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", wx.appKey, wx.appSecret, code)
}

func (wx *WXAPI) GetUserInfo(at, openID string) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN ", at, openID)

}

func (wx *WXAPI) RefreshToken(token string) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s ", wx.appKey, token)
	js, err := NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
		}
	}`))

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
}
