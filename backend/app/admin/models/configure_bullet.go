package models

import (
	"go-admin/common/models"

	"gorm.io/datatypes"
)

type ConfigureBullet struct {
	models.Model

	Name  string         `json:"name" gorm:"type:varchar(8);comment:Name"`
	Value datatypes.JSON `json:"value" gorm:"type:json;comment:Value"`
	models.ModelTime
	models.ControlBy
}

func (ConfigureBullet) TableName() string {
	return "configure_bullet"
}

func (e *ConfigureBullet) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ConfigureBullet) GetId() interface{} {
	return e.Id
}
