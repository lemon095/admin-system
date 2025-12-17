package models

import (
	"go-admin/common/models"

	"gorm.io/datatypes"
)

type ConfigureArmor struct {
	models.Model

	Name  string         `json:"name" gorm:"type:varchar(8);comment:Name"`
	Value datatypes.JSON `json:"value" gorm:"type:json;comment:Value"`
	models.ModelTime
	models.ControlBy
}

func (ConfigureArmor) TableName() string {
	return "configure_armor"
}

func (e *ConfigureArmor) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ConfigureArmor) GetId() interface{} {
	return e.Id
}
