package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joanbabyfet/letsgo/common"
	"github.com/joanbabyfet/letsgo/service"
)

// 添加参数
type LoginInput struct {
	Captcha string `form:"captcha" binding:"required"`
	Key     string `form:"key" binding:"required"`
}

// 登入
func Login(c *gin.Context) {
	verify := false
	captcha := strings.TrimSpace(c.PostForm("captcha")) //获取自验证码接口
	key := strings.TrimSpace(c.PostForm("key"))         //用户输入验证码
	if captcha == "" || key == "" {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	//检测图片验证码
	verify = service.Store.Verify(captcha, key, true)
	if !verify {
		common.Error(c, "图片验证码错误", -2, nil)
		return
	}

	common.Success(c, nil, "")
}
