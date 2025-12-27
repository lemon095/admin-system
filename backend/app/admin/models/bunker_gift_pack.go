package models

import (
	"go-admin/common/global"
	"time"

	"go-admin/common/models"
)

type GiftPack struct {
	models.Model

	GiftName      string        `json:"giftName" gorm:"type:varchar(8);comment:礼包名称"`
	GiftNum       uint8         `json:"giftNum" gorm:"type:tinyint unsigned;comment:礼包码配置数量"`
	RedeemedNum   uint8         `json:"redeemedNum" gorm:"type:tinyint unsigned;comment:已兑换数量"`
	GiftLimit     uint8         `json:"giftLimit" gorm:"type:tinyint unsigned;comment:单用户兑换次数"`
	RedeemCode    string        `json:"redeemCode" gorm:"type:char(12);comment:兑换码"`
	GiftDesc      string        `json:"giftDesc" gorm:"type:varchar(255);comment:礼包描述"`
	GoldCoinNum   uint8         `json:"goldCoinNum" gorm:"type:tinyint unsigned;comment:金币数量"`
	SilverCoinNum uint8         `json:"silverCoinNum" gorm:"type:tinyint unsigned;comment:银币数量"`
	Item          string        `json:"item" gorm:"type:json;comment:道具"`
	StartAt       time.Time     `json:"startAt" gorm:"type:datetime;comment:有效时间起始"`
	EndAt         time.Time     `json:"endAt" gorm:"type:datetime;comment:有效时间截止"`
	IsEnable      global.Status `json:"isEnable" gorm:"type:tinyint(1);comment:状态(0:启用 1:禁用)"`
	CreatedBy     string        `json:"created_by" gorm:"type:varchar(8);comment:创建人"`
	Operator      string        `gorm:"column:operator;type:varchar(8);comment:操作人" json:"operator"`
	models.ModelTime
	models.ControlBy
}

func (GiftPack) TableName() string {
	return "bunker_gift_pack"
}

func (e *GiftPack) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *GiftPack) GetId() interface{} {
	return e.Id
}
