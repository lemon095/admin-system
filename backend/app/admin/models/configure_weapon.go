package models

import (
	"go-admin/common/models"
	"gorm.io/datatypes"
)

type ConfigureWeapon struct {
	models.Model

	Name  string         `json:"name" gorm:"type:varchar(8);comment:Name"`
	Value datatypes.JSON `json:"value" gorm:"type:json;comment:Value"`
	models.ModelTime
	models.ControlBy
}

func (ConfigureWeapon) TableName() string {
	return "configure_weapon"
}

func (e *ConfigureWeapon) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ConfigureWeapon) GetId() interface{} {
	return e.Id
}
