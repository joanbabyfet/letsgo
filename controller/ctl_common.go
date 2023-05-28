package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joanbabyfet/letsgo/common"
	"github.com/joanbabyfet/letsgo/service"
)

// 获取客户端ip
func Ip(c *gin.Context) {
	data := map[string]interface{}{
		"ip": c.ClientIP(),
	}
	common.Success(c, data, "")
}

// ping检测用,可查看是否返回信息及时间戳
func Ping(c *gin.Context) {
	//common.Success(c, nil, "")
	c.String(200, "pong")
}

// 获取图形验证码
func Captcha(c *gin.Context) {
	status, key, img := service.GetCaptcha()
	if status < 0 {
		common.Error(c, "生成验证码错误", -1, nil)
		return
	}

	//组装数据
	captcha_data := make(map[string]string, 2)
	captcha_data["key"] = key
	captcha_data["img"] = img

	data := make(map[string]interface{}, 1)
	data["captcha"] = captcha_data

	common.Success(c, data, "")
}

// 重载图片验证码
func ReloadCaptcha(c *gin.Context) {
	status, key, img := service.GetCaptcha()
	if status < 0 {
		common.Error(c, "生成验证码错误", -1, nil)
		return
	}

	//组装数据
	captcha_data := make(map[string]string, 2)
	captcha_data["key"] = key
	captcha_data["img"] = img

	data := make(map[string]interface{}, 1)
	data["captcha"] = captcha_data

	common.Success(c, data, "")
}
