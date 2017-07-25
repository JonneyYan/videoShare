package api

import (
	"log"
	"net/http"
	"net/url"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
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
