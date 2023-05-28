package model

import (
	"github.com/spf13/viper"
)

var Cache_key_example string = "example:id:%s"

type Example struct {
	ID          string `gorm:"type:char(32);column:id;default:'';primary_key;comment:''" json:"id"`
	Cat_id      int    `gorm:"type:int;column:cat_id;default:0;comment:'分類id'" json:"cat_id"`
	Title       string `gorm:"type:varchar(100);column:title;default:'';comment:'標題';index" json:"title"`
	Content     string `gorm:"type:text;column:content;comment:'內容'" json:"content"`
	Img         string `gorm:"type:varchar(255);column:img;default:'';comment:'圖片'" json:"img"`
	File        string `gorm:"type:varchar(255);column:file;default:'';comment:'附件'" json:"file"`
	Is_hot      int    `gorm:"type:tinyint;column:is_hot;default:0;comment:'是否热门新聞：0=否 1=是'" json:"is_hot"`
	Sort        int    `gorm:"type:smallint;column:sort;default:0;comment:'排序：数字小的排前面'" json:"sort"`
	Status      int    `gorm:"type:tinyint;column:status;default:1;comment:'状态：0=禁用 1=启用'" json:"status"`
	Create_time int    `gorm:"type:int;column:create_time;default:0;comment:'創建時間'" json:"create_time"`
	Create_user string `gorm:"type:char(32);column:create_user;default:'0';comment:'創建人'" json:"create_user"`
	Update_time int    `gorm:"type:int;column:update_time;default:0;comment:'修改時間'" json:"update_time"`
	Update_user string `gorm:"type:char(32);column:update_user;default:'0';comment:'修改人'" json:"update_user"`
	Delete_time int    `gorm:"type:int;column:delete_time;default:0;comment:'刪除時間'" json:"delete_time"`
	Delete_user string `gorm:"type:char(32);column:delete_user;default:'0';comment:'刪除人'" json:"delete_user"`
}

// 设置表名
func (e *Example) TableName() string {
	prefix := viper.GetString("db.prefix")
	return prefix + "example"
}
