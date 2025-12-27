package models

import (
	"go-admin/common/models"

	"gorm.io/datatypes"
)

type ConfigureMonster struct {
	models.Model

	Name  string         `json:"name" gorm:"type:varchar(8);comment:Name"`
	Value datatypes.JSON `json:"value" gorm:"type:json;comment:Value"`
	models.ModelTime
	models.ControlBy
}

func (ConfigureMonster) TableName() string {
	return "bunker_monster"
}

func (e *ConfigureMonster) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ConfigureMonster) GetId() interface{} {
	return e.Id
}
