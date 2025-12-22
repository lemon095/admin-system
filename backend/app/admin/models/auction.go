package models

import (
	"go-admin/common/models"
)

type Auction struct {
	models.Model

	PriceMin int64  `json:"priceMin" gorm:"type:decimal(10,2);comment:价格下限"`
	PriceMax int64  `json:"priceMax" gorm:"type:decimal(10,2);comment:价格上限"`
	Desc     string `json:"desc" gorm:"type:varchar(255);comment:文本"`
	models.ModelTime
	models.ControlBy
}

func (Auction) TableName() string {
	return "auction"
}

func (e *Auction) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Auction) GetId() interface{} {
	return e.Id
}
