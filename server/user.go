package main

import "fmt"

// User 用户
type User struct {
	openID string
	token  string
	info   string
}

// Add 获取用户信息
func (u *User) getUserInfo() (UserInfoResponse, error) {
	userInfo, err := wx.GetUserInfo(u.token, u.openID)
	if !userInfo.Ok() || err != nil {
		return userInfo, fmt.Errorf("获取用户信息失败")
	}
	return userInfo, nil
}

// // Add 增加用户
// func (u *User) add() error {

// }

// // Login 登录
// func (u *User) login() {

// }
