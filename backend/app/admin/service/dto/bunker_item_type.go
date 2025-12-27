package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ItemTypeGetPageReq struct {
	dto.Pagination `search:"-"`
	ItemTypeOrder
}

type ItemTypeOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:bunker_item_type"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bunker_item_type"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bunker_item_type"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:bunker_item_type"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:bunker_item_type"`
}

func (m *ItemTypeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ItemTypeInsertReq struct {
	Id       int    `json:"-" comment:""` //
	Name     string `json:"name" comment:"道具类型"`
	Operator string `json:"operator" comment:""`
	common.ControlBy
}

func (s *ItemTypeInsertReq) Generate(model *models.ItemType) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
}

func (s *ItemTypeInsertReq) GetId() interface{} {
	return s.Id
}

type ItemTypeUpdateReq struct {
	Id       int    `json:"id" comment:""` //
	Name     string `json:"name" comment:"道具类型"`
	Operator string `json:"operator" comment:""`
	common.ControlBy
}

func (s *ItemTypeUpdateReq) Generate(model *models.ItemType) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Operator = s.Operator
}

func (s *ItemTypeUpdateReq) GetId() interface{} {
	return s.Id
}

// ItemTypeGetReq 功能获取请求参数
type ItemTypeGetReq struct {
	Id int `uri:"id"`
}

func (s *ItemTypeGetReq) GetId() interface{} {
	return s.Id
}

// ItemTypeDeleteReq 功能删除请求参数
type ItemTypeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ItemTypeDeleteReq) GetId() interface{} {
	return s.Ids
}
