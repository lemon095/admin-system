package models

import (
	"go-admin/common/models"
)

type GiftpackRedeems struct {
	models.Model

	RedeemerUnionId string `json:"redeemerUnionId" gorm:"type:varchar(64);comment:兑换人微信union_id"`
	RedeemCode      string `json:"redeemCode" gorm:"type:char(12);comment:兑换码"`
	models.ModelTime
	models.ControlBy
}

func (GiftpackRedeems) TableName() string {
	return "bunker_giftpack_redeems"
}

func (e *GiftpackRedeems) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *GiftpackRedeems) GetId() interface{} {
	return e.Id
}
