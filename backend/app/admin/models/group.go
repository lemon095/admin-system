package models

import (
	"go-admin/common/models"
)

type Group struct {
	models.Model

	Name        string `json:"name" gorm:"type:varchar(8);comment:战团名称"`
	Icon        string `json:"icon" gorm:"type:varchar(255);comment:战团图标"`
	Number      int64  `json:"number" gorm:"type:tinyint;comment:战团人数"`
	LastRanking int64  `json:"lastRanking" gorm:"type:tinyint;comment:上次boss战战团排名"`
	models.ModelTime
	models.ControlBy
}

func (Group) TableName() string {
	return "bunker_group"
}

func (e *Group) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Group) GetId() interface{} {
	return e.Id
}
