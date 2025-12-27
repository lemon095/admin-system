package models

import (
	"go-admin/common/models"
)

type ItemType struct {
	models.Model

	Name     string `json:"name" gorm:"type:varchar(8);comment:道具类型"`
	Operator string `gorm:"column:operator;type:varchar(8);comment:操作人" json:"operator"`
	models.ModelTime
	models.ControlBy
}

func (ItemType) TableName() string {
	return "bunker_item_type"
}

func (e *ItemType) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ItemType) GetId() interface{} {
	return e.Id
}
