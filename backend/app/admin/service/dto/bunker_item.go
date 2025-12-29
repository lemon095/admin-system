package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/global"
	common "go-admin/common/models"
)

type ItemGetPageReq struct {
	dto.Pagination `search:"-"`
	ItemId         int64  `form:"itemId"  search:"type:exact;column:item_id;table:bunker_item" comment:"道具id"`
	Type           int64  `form:"type"  search:"type:exact;column:type;table:bunker_item" comment:"道具类型"`
	Name           string `form:"name"  search:"type:exact;column:name;table:bunker_item" comment:"道具名称"`
	ItemOrder
}

type GetOptionResp struct {
	Value    int             `json:"value"`
	Label    string          `json:"label"`
	Children []GetOptionResp `json:"children"`
}

type ItemOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:bunker_item"`
	ItemId    string `form:"itemIdOrder"  search:"type:order;column:item_id;table:bunker_item"`
	Type      string `form:"typeOrder"  search:"type:order;column:type;table:bunker_item"`
	Name      string `form:"nameOrder"  search:"type:order;column:name;table:bunker_item"`
	Desc      string `form:"descOrder"  search:"type:order;column:desc;table:bunker_item"`
	Icon      string `form:"iconOrder"  search:"type:order;column:icon;table:bunker_item"`
	Extend    string `form:"extendOrder"  search:"type:order;column:extend;table:bunker_item"`
	IsEnable  string `form:"isEnableOrder"  search:"type:order;column:is_enable;table:bunker_item"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:bunker_item"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:bunker_item"`
}

func (m *ItemGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ItemInsertReq struct {
	Type     int64         `json:"type" comment:"道具类型"`
	Name     string        `json:"name" comment:"道具名称"`
	Desc     string        `json:"desc" comment:"道具描述"`
	Icon     string        `json:"icon" comment:"图片地址"`
	Extend   *string       `json:"extend" comment:"扩展属性"`
	IsEnable global.Status `json:"isEnable" comment:"状态(0:启用 1:禁用)"`
	Operator string        `json:"operator" comment:""`
	common.ControlBy
}

func (s *ItemInsertReq) Generate(model *models.Item) {
	model.Type = s.Type
	model.Name = s.Name
	model.Desc = s.Desc
	model.Icon = s.Icon
	model.Extend = s.Extend
	model.IsEnable = s.IsEnable
	model.ControlBy = s.ControlBy
}

func (s *ItemInsertReq) GetId() interface{} {
	return 0
}

type ItemUpdateReq struct {
	Id       int           `uri:"id" comment:""` //
	ItemId   int           `json:"itemId" comment:"道具id"`
	Type     int64         `json:"type" comment:"道具类型"`
	Name     string        `json:"name" comment:"道具名称"`
	Desc     string        `json:"desc" comment:"道具描述"`
	Icon     string        `json:"icon" comment:"图片地址"`
	Extend   *string       `json:"extend" comment:"扩展属性"`
	IsEnable global.Status `json:"isEnable" comment:"状态(0:启用 1:禁用)"`
	Operator string        `json:"operator" comment:""`
	common.ControlBy
}

func (s *ItemUpdateReq) Generate(model *models.Item) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ItemId = s.ItemId
	model.Type = s.Type
	model.Name = s.Name
	model.Desc = s.Desc
	model.Icon = s.Icon
	model.Extend = s.Extend
	model.IsEnable = s.IsEnable
	model.ControlBy = s.ControlBy
}

func (s *ItemUpdateReq) GetId() interface{} {
	return s.Id
}

// ItemGetReq 功能获取请求参数
type ItemGetReq struct {
	Id int `uri:"id"`
}

func (s *ItemGetReq) GetId() interface{} {
	return s.Id
}

// ItemDeleteReq 功能删除请求参数
type ItemDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ItemDeleteReq) GetId() interface{} {
	return s.Ids
}
