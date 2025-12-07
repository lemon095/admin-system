package model

import (
	"admin-system/database"
	"time"

	"gorm.io/datatypes"
)

func tableName(category string) string {
	table := "system_"
	switch category {
	case "monster":
		table = table + "monster"
	case "armor":
		table = table + "armor"
	case "bullet":
		table = table + "bullet"
	default:
		table = table + "weapon"
	}
	return table
}

type System struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"type:varchar(16);unique"`
	Value     datatypes.JSON `json:"value" gorm:"type:json"` // 存防具数值配置
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func SystemList(category string) ([]System, error) {
	var results []System
	if err := database.GormDB.Table(tableName(category)).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func SystemGetByID(category string, id int) (System, error) {
	var results System
	if err := database.GormDB.Table(tableName(category)).First(&results, id).Error; err != nil {
		return results, err
	}
	return results, nil
}

func SystemUpdateByID(category string, id int, value any) error {
	return database.GormDB.Table(tableName(category)).Where("id=?", id).Update("value", value).Error
}

func SystemCreate(category string, data *System) error {
	return database.GormDB.Table(tableName(category)).Create(data).Error
}

func SystemDeleteByID(category string, id int) error {
	return database.GormDB.Table(tableName(category)).Delete(&System{}, id).Error
}
