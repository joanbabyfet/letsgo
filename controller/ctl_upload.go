package controller

import (
	"path"

	"github.com/gin-gonic/gin"
	"github.com/joanbabyfet/letsgo/common"
	"github.com/spf13/viper"
)

// 普通上传
func Upload(c *gin.Context) {
	dir := c.DefaultPostForm("dir", "image")
	file, err := c.FormFile("file")
	if err != nil {
		common.Error(c, "参数错误", -1, nil)
		return
	}
	upload_dir := "uploads/" + dir              //文件上传目录
	dst := path.Join(upload_dir, file.Filename) //组装文件路径, /uploads/image/xxx.jpg
	// 上传文件到指定的目录
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		common.Error(c, "上传失败", -2, nil)
		return
	}

	//组装数据
	filelink := viper.GetString("app.file_link")
	data := make(map[string]string, 3)
	data["realname"] = file.Filename
	data["filename"] = file.Filename
	data["filelink"] = filelink + "/" + dir + "/" + file.Filename

	common.Success(c, data, "")
}
