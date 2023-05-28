package model

import (
	"github.com/spf13/viper"
)

var Cache_key_example string = "example:id:%s"

type Example struct {
	ID          string `gorm:"type:char(32);column:id;primary_key" json:"id"`
	Cat_id      int    `gorm:"type:int;column:cat_id" json:"cat_id"`
	Title       string `gorm:"type:varchar(100);column:title" json:"title"`
	Content     string `gorm:"type:text;column:content" json:"content"`
	Img         string `gorm:"type:varchar(255);column:img" json:"img"`
	File        string `gorm:"type:varchar(255);column:file" json:"file"`
	Is_hot      int    `gorm:"type:tinyint;column:is_hot" json:"is_hot"`
	Sort        int    `gorm:"type:smallint;column:sort" json:"sort"`
	Status      int    `gorm:"type:tinyint;column:status" json:"status"`
	Create_time int    `gorm:"type:int;column:create_time" json:"create_time"`
	Create_user string `gorm:"type:char(32);column:create_user" json:"create_user"`
	Update_time int    `gorm:"type:int;column:update_time" json:"update_time"`
	Update_user string `gorm:"type:char(32);column:update_user" json:"update_user"`
	Delete_time int    `gorm:"type:int;column:delete_time" json:"delete_time"`
	Delete_user string `gorm:"type:char(32);column:delete_user" json:"delete_user"`
}

// 设置表名
func (e *Example) TableName() string {
	prefix := viper.GetString("db.prefix")
	return prefix + "example"
}
