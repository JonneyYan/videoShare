package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/JonneyYan/server/api"
)

const (
	appID     = "123123"
	appkey    = "123123"
	appSecret = "123123"
	mchID     = "123"
)

var wx Weixin

func main() {
	var (
		port = flag.Int("port", 80, "请输入服务器端口")
	)

	flag.Parse()
	// 静态文件 os 绝对路径
	wd, err := os.Getwd() // 当前路径
	if err != nil {
		log.Fatal(err)
	}

	wx = NewWX(appID, appSecret)
	http.HandleFunc("/", handleIndexHTMLFile)
	http.HandleFunc("/pay", api.HandlePay)
	http.HandleFunc("/login", api.HandleLogin)

	http.HandleFunc("/payCallback", handleIndexHTMLFile)

	http.HandleFunc("/statistics", handleIndexHTMLFile)
	http.HandleFunc("/withdraw", handleIndexHTMLFile)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func handleIndexHTMLFile(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	cookie, err := r.Cookie("openID")
	if err != nil || cookie.Value == "" {
		encodeurl := url.QueryEscape("http://shop.cocoabox.cn")
		url := wx.WebAuthRedirectURL(encodeurl, "snsapi_userinfo", "")
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
	openID := cookie.Value
	token, _ := r.Cookie("token")
	user := User{openID: openID, token: token.Value}

	userInfo, err := user.GetUserInfo()
	if err != nil {
		cookie := http.Cookie{Name: "isMember", Value: "false", Path: "/"}
		http.SetCookie(w, &cookie)
	}
	cookieUserInfo := http.Cookie{Name: "userInfo", Value: userInfo, Path: "/"}
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
