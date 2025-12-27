package models

import (
	"go-admin/common/global"
	"go-admin/common/models"
)

type Item struct {
	models.Model

	ItemId   int64         `json:"itemId" gorm:"type:bigint;comment:道具id"`
	Type     int64         `json:"type" gorm:"type:bigint;comment:道具类型"`
	Name     string        `json:"name" gorm:"type:varchar(8);comment:道具名称"`
	Desc     string        `json:"desc" gorm:"type:varchar(255);comment:道具描述"`
	Icon     string        `json:"icon" gorm:"type:varchar(255);comment:图片地址"`
	Extend   *string       `json:"extend" gorm:"type:json;comment:扩展属性"`
	IsEnable global.Status `json:"isEnable" gorm:"type:tinyint(1);comment:状态(0:启用 1:禁用)"`
	models.ModelTime
	models.ControlBy
}

func (Item) TableName() string {
	return "bunker_item"
}

func (e *Item) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Item) GetId() interface{} {
	return e.Id
}
