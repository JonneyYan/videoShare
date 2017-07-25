package main

import (
	"encoding/json"
	"log"
)

type User struct {
	openID string
	token  string
}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo() (string, error) {
	userResp, err := wx.GetUserInfo(u.token, u.openID)

	if userResp.Ok() {
		b, err := json.Marshal(userResp)
		if err != nil {
			log.Println("userinfo json error")
			return "", err
		}
		return string(b), nil
	}

	return "", nil
}

// Add 增加用户
func (u *User) Add() error {

// }

// Freeze 冻结用户
func (u *User) Freeze() {

}
