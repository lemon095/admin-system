package models

import (
	"go-admin/common/models"

	"github.com/shopspring/decimal"
)

type Workshop struct {
	models.Model

	Name            string          `json:"name" gorm:"type:varchar(8);comment:工坊名字"`
	MaxOutput       decimal.Decimal `json:"maxOutput" gorm:"type:decimal(10,2);comment:每个工坊最大产量"`
	MaterialOutput  decimal.Decimal `json:"materialOutput" gorm:"type:decimal(10,2);comment:产出材料"`
	MaterialRequire decimal.Decimal `json:"materialRequire" gorm:"type:decimal(10,2);comment:升级所需材料"`
	Level           uint8           `json:"level" gorm:"type:tinyint;comment:等级"`
	models.ModelTime
	models.ControlBy
}

func (Workshop) TableName() string {
	return "workshop"
}

func (e *Workshop) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Workshop) GetId() interface{} {
	return e.Id
}
