package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/JonneyYan/server/wxpay"
)

const (
	appID     = "wx2f3392237a4fdf1d"
	appkey    = "123123"
	appSecret = "88b54e67fdfe04b7f832f25a313f5b9f"
)

var wx Weixin

func main() {
	var (
		port = 8088
	)
	fmt.Println("port:", port)
	wx = NewWX(appID, appSecret)
	flag.Parse()
	// 静态文件 os 绝对路径
	wd, err := os.Getwd() // 当前路径
	if err != nil {
		log.Fatal(err)
	}

	wx = NewWX(appID, appSecret)
	http.HandleFunc("/", handleIndexHTMLFile)
	http.HandleFunc("/pay", handlePay)
	http.HandleFunc("/login", handleLogin)

	http.HandleFunc("/payCallback", handleIndexHTMLFile)

	http.HandleFunc("/api/statistics", handleIndexHTMLFile)
	http.HandleFunc("/api/withdraw", handleIndexHTMLFile)

	http.HandleFunc("/*", handleIndexHTMLFile)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func handleIndexHTMLFile(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	cookie, err := r.Cookie("userInfo")
	// 如果用户未登录跳转到登录页面
	if err != nil || cookie.Value == "" {
		encodeurl := url.QueryEscape("http://shop.cocoabox.cn")
		url := wx.WebAuthRedirectURL(encodeurl, "snsapi_userinfo", "")
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
	openID := cookie.Value
	token, _ := r.Cookie("token")
	user := User{openID: openID, token: token.Value}

	userInfo, err := user.getUserInfo()
	// 如果用户不是会员，在cookie中注入isMember
	if err != nil {
		cookie := http.Cookie{Name: "isMember", Value: "false", Path: "/"}
		http.SetCookie(w, &cookie)
	}
	cookieUserInfo := http.Cookie{Name: "userInfo", Value: userInfo.toJSON(), Path: "/"}
	http.SetCookie(w, &cookieUserInfo)

	htmlFile := "../client/dist/index.html"
	fl, err := os.Open(htmlFile)
	if err != nil {
		fmt.Println(htmlFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		fmt.Fprint(w, buf[:n])
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	code := r.Form.Get("code")
	userResp, err := wx.GetWebAccessToken(code)

	if err != nil {
		log.Printf("用户登录失败!:%s \n", err)
	}

	if userResp.Ok() {
		cookieOpenID := http.Cookie{Name: "openID", Value: userResp.OpenID, Path: "/"}
		cookieToken := http.Cookie{Name: "token", Value: userResp.AccessToken, Path: "/"}
		http.SetCookie(w, &cookieOpenID)
		http.SetCookie(w, &cookieToken)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		encodeurl := url.QueryEscape("http://shop.cocoabox.cn")
		url := wx.WebAuthRedirectURL(encodeurl, "snsapi_userinfo", "")
		http.Redirect(w, r, url, http.StatusFound)
	}

}
func handlePay(w http.ResponseWriter, r *http.Request) {
	wxpay := wxpay.NewAPI()
}
