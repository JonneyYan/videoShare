package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Weixin struct {
	AppID     string
	AppSecret string
}
type WeiXinResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (resp *WeiXinResponse) Ok() bool {
	if resp.Errcode == 0 {
		return true
	}
	return false
}

type WebAccessWeiXinResponse struct {
	WeiXinResponse
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openID"`
	Scope        string `json:"scope"`
	UnionID      string `json:"unionid"`
}

type AccessWeiXinResponse struct {
	WeiXinResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type UserInfoResponse struct {
	WeiXinResponse
	OpenID     string   `json:"openID"`
	Nickname   string   `json:"nickname"`
	Sex        int64    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"unionid"`
}

func (u *UserInfoResponse) toJSON() string {
	b, err := json.Marshal(u)
	if err != nil {
		return fmt.Sprintf("json err:%v", err)
	}
	return fmt.Sprintf(string(b))
}

type JSSDKTicketResponse struct {
	WeiXinResponse
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

const (
	webAuthRedirectURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	getAccessToken     = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	getWebAccessToken  = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	getUserInfo        = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openID=%s&lang=zh_CN"
	getJSSDKTicket     = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

func NewWX(appID string, appSecret string) Weixin {
	return Weixin{
		AppID:     appID,
		AppSecret: appSecret,
	}
}

func (wx *Weixin) WebAuthRedirectURL(redirectURI string, scope string, state string) string {
	redirectURIEscaped := url.QueryEscape(redirectURI)
	return fmt.Sprintf(webAuthRedirectURL, wx.AppID, redirectURIEscaped, scope, state)
}

func requestGet(url string) ([]byte, error) {
	log.Println("get response on url: ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("failed to get url")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response body")
		return nil, err
	}
	log.Println("response body is ", string(body))

	return body, nil
}

func (wx *Weixin) GetJSSDKTicket(accessToken string) (JSSDKTicketResponse, error) {
	var response JSSDKTicketResponse
	url := fmt.Sprintf(getJSSDKTicket, accessToken)
	log.Println("get access token request url: ", url)
	body, err := requestGet(url)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to parse body to json")
		return response, err
	}
	log.Printf("body json response is %v\n", response)
	return response, nil
}

func (wx *Weixin) GetAccessToken() (AccessWeiXinResponse, error) {
	var response AccessWeiXinResponse
	url := fmt.Sprintf(getAccessToken, wx.AppID, wx.AppSecret)
	log.Println("get access token request url: ", url)
	body, err := requestGet(url)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to parse body to json")
		return response, err
	}
	log.Printf("body json response is %v\n", response)
	return response, nil
}

func (wx *Weixin) GetWebAccessToken(code string) (WebAccessWeiXinResponse, error) {
	var response WebAccessWeiXinResponse
	url := fmt.Sprintf(getWebAccessToken, wx.AppID, wx.AppSecret, code)
	log.Println("get web access token request url: ", url)
	body, err := requestGet(url)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to parse body to json")
		return response, err
	}
	log.Printf("body json response is %v\n", response)
	return response, nil
}

func (wx *Weixin) GetUserInfo(accessToken string, openID string) (UserInfoResponse, error) {
	var response UserInfoResponse
	url := fmt.Sprintf(getUserInfo, accessToken, openID)
	log.Println("get user info request url: ", url)
	body, err := requestGet(url)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("failed to parse body to json")
		return response, err
	}
	log.Printf("body json response is %v\n", response)
	return response, nil
}

func (wx *Weixin) JSSDKSignature(jssdkTicket string, noncestr string, timestamp int64, url string) string {
	string1 := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", jssdkTicket, noncestr, timestamp, url)
	return fmt.Sprintf("%x", sha1.Sum([]byte(string1)))
}
