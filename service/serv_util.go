package service

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
)

// 全局变量
var Store = base64Captcha.DefaultMemStore

// 获取图片验证码
func GetCaptcha() (int, string, string) {
	status := 1
	driver := base64Captcha.DefaultDriverDigit
	captcha := base64Captcha.NewCaptcha(driver, Store)
	key, img, err := captcha.Generate()
	if err != nil {
		status = -1
		fmt.Println(err)
	}
	return status, key, img
}
