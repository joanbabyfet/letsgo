package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joanbabyfet/letsgo/common"
	"github.com/joanbabyfet/letsgo/service"
)

// 登入
func Login(c *gin.Context) {
	//verify := false
	//captcha := strings.TrimSpace(c.PostForm("captcha")) //获取自验证码接口
	//key := strings.TrimSpace(c.PostForm("key"))         //用户输入验证码

	var input struct { //参数过滤
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		Captcha  string `form:"captcha" binding:"required"` //获取自验证码接口
		Key      string `form:"key" binding:"required"`     //用户输入验证码
	}
	if err := c.Bind(&input); err != nil {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	//检测图片验证码
	// verify = service.Store.Verify(captcha, key, true)
	// if !verify {
	// 	common.Error(c, "图片验证码错误", -2, nil)
	// 	return
	// }

	var token string
	if input.Username == "admin" && input.Password == "Bb123456" {
		token = service.GenToken(input.Username)
	}

	//组装数据
	data := make(map[string]string, 3)
	data["access_token"] = token
	data["token_type"] = "bearer"

	common.Success(c, data, "")
}
