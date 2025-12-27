package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GiftpackRedeemsGetPageReq struct {
	dto.Pagination  `search:"-"`
	RedeemerUnionId string `form:"redeemerUnionId"  search:"type:exact;column:redeemer_union_id;table:bunker_giftpack_redeems" comment:"兑换人微信union_id"`
	RedeemCode      string `form:"redeemCode"  search:"type:exact;column:redeem_code;table:bunker_giftpack_redeems" comment:"兑换码"`
	GiftpackRedeemsOrder
}

type GiftpackRedeemsOrder struct {
	Id              string `form:"idOrder"  search:"type:order;column:id;table:bunker_giftpack_redeems"`
	RedeemerUnionId string `form:"redeemerUnionIdOrder"  search:"type:order;column:redeemer_union_id;table:bunker_giftpack_redeems"`
	RedeemCode      string `form:"redeemCodeOrder"  search:"type:order;column:redeem_code;table:bunker_giftpack_redeems"`
	CreatedAt       string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bunker_giftpack_redeems"`
}

func (m *GiftpackRedeemsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GiftpackRedeemsInsertReq struct {
	Id              int    `json:"-" comment:""` //
	RedeemerUnionId string `json:"redeemerUnionId" comment:"兑换人微信union_id"`
	RedeemCode      string `json:"redeemCode" comment:"兑换码"`
	common.ControlBy
}

func (s *GiftpackRedeemsInsertReq) Generate(model *models.GiftpackRedeems) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RedeemerUnionId = s.RedeemerUnionId
	model.RedeemCode = s.RedeemCode
}

func (s *GiftpackRedeemsInsertReq) GetId() interface{} {
	return s.Id
}

type GiftpackRedeemsUpdateReq struct {
	Id              int    `uri:"id" comment:""` //
	RedeemerUnionId string `json:"redeemerUnionId" comment:"兑换人微信union_id"`
	RedeemCode      string `json:"redeemCode" comment:"兑换码"`
	common.ControlBy
}

func (s *GiftpackRedeemsUpdateReq) Generate(model *models.GiftpackRedeems) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RedeemerUnionId = s.RedeemerUnionId
	model.RedeemCode = s.RedeemCode
}

func (s *GiftpackRedeemsUpdateReq) GetId() interface{} {
	return s.Id
}

// GiftpackRedeemsGetReq 功能获取请求参数
type GiftpackRedeemsGetReq struct {
	Id int `uri:"id"`
}

func (s *GiftpackRedeemsGetReq) GetId() interface{} {
	return s.Id
}

// GiftpackRedeemsDeleteReq 功能删除请求参数
type GiftpackRedeemsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GiftpackRedeemsDeleteReq) GetId() interface{} {
	return s.Ids
}
