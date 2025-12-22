package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"

	"github.com/shopspring/decimal"
)

type WorkshopGetPageReq struct {
	dto.Pagination `search:"-"`
	WorkshopOrder
}

type WorkshopOrder struct {
	Id              string          `form:"idOrder"  search:"type:order;column:id;table:workshop"`
	Name            string          `form:"nameOrder"  search:"type:order;column:name;table:workshop"`
	MaxOutput       decimal.Decimal `form:"maxOutputOrder"  search:"type:order;column:max_output;table:workshop"`
	MaterialOutput  decimal.Decimal `form:"materialOutputOrder"  search:"type:order;column:material_output;table:workshop"`
	MaterialRequire decimal.Decimal `form:"materialRequireOrder"  search:"type:order;column:material_require;table:workshop"`
	Level           uint8           `form:"levelOrder"  search:"type:order;column:level;table:workshop"`
	CreatedAt       string          `form:"createdAtOrder"  search:"type:order;column:created_at;table:workshop"`
	UpdatedAt       string          `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:workshop"`
}

func (m *WorkshopGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type WorkshopInsertReq struct {
	Id              int             `json:"-" comment:""` //
	Name            string          `json:"name" comment:"工坊名字"`
	MaxOutput       decimal.Decimal `json:"maxOutput" comment:"每个工坊最大产量"`
	MaterialOutput  decimal.Decimal `json:"materialOutput" comment:"产出材料"`
	MaterialRequire decimal.Decimal `json:"materialRequire" comment:"升级所需材料"`
	Level           uint8           `json:"level" comment:"等级"`
	common.ControlBy
}

func (s *WorkshopInsertReq) Generate(model *models.Workshop) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.MaxOutput = s.MaxOutput
	model.MaterialOutput = s.MaterialOutput
	model.MaterialRequire = s.MaterialRequire
	model.Level = s.Level
}

func (s *WorkshopInsertReq) GetId() interface{} {
	return s.Id
}

type WorkshopUpdateReq struct {
	Id              int             `uri:"id" comment:""` //
	Name            string          `json:"name" comment:"工坊名字"`
	MaxOutput       decimal.Decimal `json:"maxOutput" comment:"每个工坊最大产量"`
	MaterialOutput  decimal.Decimal `json:"materialOutput" comment:"产出材料"`
	MaterialRequire decimal.Decimal `json:"materialRequire" comment:"升级所需材料"`
	Level           uint8           `json:"level" comment:"等级"`
	common.ControlBy
}

func (s *WorkshopUpdateReq) Generate(model *models.Workshop) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.MaxOutput = s.MaxOutput
	model.MaterialOutput = s.MaterialOutput
	model.MaterialRequire = s.MaterialRequire
	model.Level = s.Level
}

func (s *WorkshopUpdateReq) GetId() interface{} {
	return s.Id
}

// WorkshopGetReq 功能获取请求参数
type WorkshopGetReq struct {
	Id int `uri:"id"`
}

func (s *WorkshopGetReq) GetId() interface{} {
	return s.Id
}

// WorkshopDeleteReq 功能删除请求参数
type WorkshopDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *WorkshopDeleteReq) GetId() interface{} {
	return s.Ids
}
