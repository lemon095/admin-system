package dto

import (
	"go-admin/common/global"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GiftPackGetPageReq struct {
	dto.Pagination `search:"-"`
	GiftPackOrder

	IsEnable   *global.Status `form:"isEnable"`
	GiftName   *string        `form:"giftName"`
	RedeemCode *string        `form:"redeemCode"`
}

type GiftPackOrder struct {
	Id            string        `form:"idOrder"  search:"type:order;column:id;table:bunker_gift_pack"`
	GiftName      string        `form:"giftNameOrder"  search:"type:order;column:gift_name;table:bunker_gift_pack"`
	GiftNum       string        `form:"giftNumOrder"  search:"type:order;column:gift_num;table:bunker_gift_pack"`
	GiftLimit     string        `form:"giftLimitOrder"  search:"type:order;column:gift_limit;table:bunker_gift_pack"`
	RedeemCode    string        `form:"redeemCodeOrder"  search:"type:order;column:redeem_code;table:bunker_gift_pack"`
	GiftDesc      string        `form:"giftDescOrder"  search:"type:order;column:gift_desc;table:bunker_gift_pack"`
	GoldCoinNum   string        `form:"goldCoinNumOrder"  search:"type:order;column:gold_coin_num;table:bunker_gift_pack"`
	SilverCoinNum string        `form:"silverCoinNumOrder"  search:"type:order;column:silver_coin_num;table:bunker_gift_pack"`
	Item          string        `form:"itemOrder"  search:"type:order;column:item;table:bunker_gift_pack"`
	StartAt       string        `form:"startAtOrder"  search:"type:order;column:start_at;table:bunker_gift_pack"`
	EndAt         string        `form:"endAtOrder"  search:"type:order;column:end_at;table:bunker_gift_pack"`
	IsEnable      global.Status `form:"isEnableOrder"  search:"type:order;column:is_enable;table:bunker_item"`
	CreatedAt     string        `form:"createdAtOrder"  search:"type:order;column:created_at;table:bunker_gift_pack"`
	UpdatedAt     string        `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bunker_gift_pack"`
}

func (m *GiftPackGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GiftPackInsertReq struct {
	Id            int           `json:"-" comment:""` //
	GiftName      string        `json:"giftName" comment:"礼包名称"`
	GiftNum       uint8         `json:"giftNum" comment:"礼包码配置数量"`
	GiftLimit     uint8         `json:"giftLimit" comment:"单用户兑换次数"`
	RedeemCode    string        `json:"redeemCode" comment:"兑换码"`
	GiftDesc      string        `json:"giftDesc" comment:"礼包描述"`
	GoldCoinNum   uint8         `json:"goldCoinNum" comment:"金币数量"`
	SilverCoinNum uint8         `json:"silverCoinNum" comment:"银币数量"`
	Item          string        `json:"item" comment:"道具"`
	CreatedBy     string        `json:"createdBy" comment:"道具"`
	StartAt       time.Time     `json:"startAt" comment:"有效时间起始"`
	EndAt         time.Time     `json:"endAt" comment:"有效时间截止"`
	IsEnable      global.Status `json:"isEnable" comment:"状态(0:启用 1:禁用)"`
	Operator      string        `json:"operator" comment:""`
	common.ControlBy
}

func (s *GiftPackInsertReq) Generate(model *models.GiftPack) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GiftName = s.GiftName
	model.GiftNum = s.GiftNum
	model.GiftLimit = s.GiftLimit
	model.RedeemCode = s.RedeemCode
	model.GiftDesc = s.GiftDesc
	model.GoldCoinNum = s.GoldCoinNum
	model.SilverCoinNum = s.SilverCoinNum
	model.Item = s.Item
	model.StartAt = s.StartAt
	model.EndAt = s.EndAt
	model.CreatedBy = s.CreatedBy
	model.IsEnable = global.STATUS_DISABLE
}

func (s *GiftPackInsertReq) GetId() interface{} {
	return s.Id
}

type GiftPackUpdateReq struct {
	Id            int           `uri:"id" comment:""` //
	GiftName      string        `json:"giftName" comment:"礼包名称"`
	GiftNum       uint8         `json:"giftNum" comment:"礼包码配置数量"`
	GiftLimit     uint8         `json:"giftLimit" comment:"单用户兑换次数"`
	RedeemCode    string        `json:"redeemCode" comment:"兑换码"`
	GiftDesc      string        `json:"giftDesc" comment:"礼包描述"`
	GoldCoinNum   uint8         `json:"goldCoinNum" comment:"金币数量"`
	SilverCoinNum uint8         `json:"silverCoinNum" comment:"银币数量"`
	Item          string        `json:"item" comment:"道具"`
	StartAt       time.Time     `json:"startAt" comment:"有效时间起始"`
	EndAt         time.Time     `json:"endAt" comment:"有效时间截止"`
	IsEnable      global.Status `json:"isEnable" comment:"状态(0:启用 1:禁用)"`
	Operator      string        `json:"operator" comment:""`
	common.ControlBy
}

func (s *GiftPackUpdateReq) Generate(model *models.GiftPack) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GiftName = s.GiftName
	model.GiftNum = s.GiftNum
	model.GiftLimit = s.GiftLimit
	model.RedeemCode = s.RedeemCode
	model.GiftDesc = s.GiftDesc
	model.GoldCoinNum = s.GoldCoinNum
	model.SilverCoinNum = s.SilverCoinNum
	model.Item = s.Item
	model.StartAt = s.StartAt
	model.EndAt = s.EndAt
}

func (s *GiftPackUpdateReq) GetId() interface{} {
	return s.Id
}

// GiftPackGetReq 功能获取请求参数
type GiftPackGetReq struct {
	Id int `uri:"id"`
}

func (s *GiftPackGetReq) GetId() interface{} {
	return s.Id
}

// GiftPackDeleteReq 功能删除请求参数
type GiftPackDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GiftPackDeleteReq) GetId() interface{} {
	return s.Ids
}
