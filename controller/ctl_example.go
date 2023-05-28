package controller

import (
	"strconv"
	"time"

	"github.com/joanbabyfet/letsgo/common"
	"github.com/joanbabyfet/letsgo/global"
	"github.com/joanbabyfet/letsgo/model"
	"github.com/joanbabyfet/letsgo/service"

	"github.com/gin-gonic/gin"
)

// 添加参数
type AddExampleInput struct {
	Title   string `form:"title" binding:"required"`   //标题必填
	Content string `form:"content" binding:"required"` //内容必填
}

// 修改参数
type EditExampleInput struct {
	Title   string `form:"title" binding:"required"`   //标题必填
	Content string `form:"content" binding:"required"` //内容必填
}

// 获取列表
func ExampleIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))    //第几页
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10")) //每页几条
	title := c.DefaultQuery("title", "")
	offset := (page - 1) * limit

	var examples []model.Example
	query := global.DB.Where("status = 1")
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	query.Limit(limit).Offset(offset).Order("create_time DESC").Find(&examples)

	common.Success(c, examples, "")
}

// 获取详情
func ExampleDetail(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	var example model.Example
	ret := global.DB.First(&example, "id = ?", id)
	if ret.Error != nil {
		common.Error(c, "该文章不存在", -2, nil)
		return
	}
	common.Success(c, example, "")
}

// 添加
func AddExample(c *gin.Context) {
	var input AddExampleInput //参数过滤
	if err := c.Bind(&input); err != nil {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	//组装数据
	data := model.Example{
		ID:          common.Random(c, "web", 32),
		Title:       input.Title,
		Content:     input.Content,
		Cat_id:      1,
		Status:      1,
		Create_time: int(time.Now().Unix()),
		Create_user: "1",
	}
	ret := global.DB.Create(&data)
	if ret.Error != nil {
		common.Error(c, "添加失败", -2, nil)
		return
	}
	//寫入日志
	service.AddAdminUserOplog(c, "文章添加", 0)
	common.Success(c, nil, "添加成功")
}

// 修改
func EditExample(c *gin.Context) {
	id := c.PostForm("id")
	var input EditExampleInput //参数过滤
	if err := c.Bind(&input); err != nil {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	var example model.Example
	ret := global.DB.First(&example, "id = ?", id)
	if ret.Error != nil {
		common.Error(c, "该文章不存在", -2, nil)
		return
	}

	//组装数据
	data := model.Example{
		Title:       input.Title,
		Content:     input.Content,
		Cat_id:      1,
		Status:      1,
		Update_time: int(time.Now().Unix()),
		Update_user: "1",
	}
	global.DB.Model(&example).Updates(&data)
	//寫入日志
	service.AddAdminUserOplog(c, "文章修改 "+id, 0)

	common.Success(c, nil, "更新成功")
}

// 删除
func DeleteExample(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	ret := global.DB.Delete(&model.Example{}, "id = ?", id)
	if ret.Error != nil {
		common.Error(c, "删除失败", -2, nil)
		return
	}
	//寫入日志
	service.AddAdminUserOplog(c, "文章刪除 "+id, 0)

	common.Success(c, nil, "删除成功")
}

// 启用
func EnableExample(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	var example model.Example
	ret := global.DB.First(&example, "id = ?", id)
	if ret.Error != nil {
		common.Error(c, "该文章不存在", -2, nil)
		return
	}

	//组装数据, 这里要用map才能更新status为0, 且不能用&data
	data := map[string]interface{}{
		"Status":      1,
		"Update_time": int(time.Now().Unix()),
		"Update_user": "1",
	}
	global.DB.Model(&example).Updates(data)
	//寫入日志
	service.AddAdminUserOplog(c, "文章启用 "+id, 0)

	common.Success(c, nil, "启用成功")
}

// 禁用
func DisableExample(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		common.Error(c, "参数错误", -1, nil)
		return
	}

	var example model.Example
	ret := global.DB.First(&example, "id = ?", id)
	if ret.Error != nil {
		common.Error(c, "该文章不存在", -2, nil)
		return
	}

	//组装数据, 这里要用map才能更新status为0, 且不能用&data
	data := map[string]interface{}{
		"Status":      0,
		"Update_time": int(time.Now().Unix()),
		"Update_user": "1",
	}
	global.DB.Model(&example).Updates(data)
	//寫入日志
	service.AddAdminUserOplog(c, "文章禁用 "+id, 0)

	common.Success(c, nil, "禁用成功")
}
