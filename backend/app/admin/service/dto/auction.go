package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"

	"github.com/shopspring/decimal"
)

type AuctionGetPageReq struct {
	dto.Pagination `search:"-"`
	AuctionOrder
}

type AuctionOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:auction"`
	PriceMin  string `form:"priceMinOrder"  search:"type:order;column:price_min;table:auction"`
	PriceMax  string `form:"priceMaxOrder"  search:"type:order;column:price_max;table:auction"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:auction"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:auction"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:auction"`
}

func (m *AuctionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type AuctionInsertReq struct {
	Id       int             `json:"-" comment:""` //
	PriceMin decimal.Decimal `json:"priceMin" comment:"价格下限"`
	PriceMax decimal.Decimal `json:"priceMax" comment:"价格上限"`
	Desc     string          `json:"desc" comment:"文本"`
	Operator string          `json:"operator" comment:""`
	common.ControlBy
}

func (s *AuctionInsertReq) Generate(model *models.Auction) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PriceMin = s.PriceMin
	model.PriceMax = s.PriceMax
	model.Desc = s.Desc
}

func (s *AuctionInsertReq) GetId() interface{} {
	return s.Id
}

type AuctionUpdateReq struct {
	Id       int             `uri:"id" comment:""` //
	PriceMin decimal.Decimal `json:"priceMin" comment:"价格下限"`
	PriceMax decimal.Decimal `json:"priceMax" comment:"价格上限"`
	Desc     string          `json:"desc" comment:"文本"`
	Operator string          `json:"operator" comment:""`
	common.ControlBy
}

func (s *AuctionUpdateReq) Generate(model *models.Auction) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PriceMin = s.PriceMin
	model.PriceMax = s.PriceMax
	model.Desc = s.Desc
}

func (s *AuctionUpdateReq) GetId() interface{} {
	return s.Id
}

// AuctionGetReq 功能获取请求参数
type AuctionGetReq struct {
	Id int `uri:"id"`
}

func (s *AuctionGetReq) GetId() interface{} {
	return s.Id
}

// AuctionDeleteReq 功能删除请求参数
type AuctionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *AuctionDeleteReq) GetId() interface{} {
	return s.Ids
}
