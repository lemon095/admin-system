package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GroupGetPageReq struct {
	dto.Pagination `search:"-"`
	GroupOrder
}

type GroupOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:group"`
	Name        string `form:"nameOrder"  search:"type:order;column:name;table:group"`
	Icon        string `form:"iconOrder"  search:"type:order;column:icon;table:group"`
	Number      string `form:"numberOrder"  search:"type:order;column:number;table:group"`
	LastRanking string `form:"lastRankingOrder"  search:"type:order;column:last_ranking;table:group"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:group"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:group"`
}

func (m *GroupGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GroupReq struct {
	Id          int    `json:"-" comment:"战团id"` // 战团id
	Name        string `json:"name" comment:"战团名称"`
	Icon        string `json:"icon" comment:"战团图标"`
	Number      int64  `json:"number" comment:"战团人数"`
	LastRanking int64  `json:"lastRanking" comment:"上次boss战战团排名"`
}

type GroupInsertReq struct {
	GroupReq
	common.ControlBy
}

func (s *GroupInsertReq) Generate(model *models.Group) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Icon = s.Icon
	model.Number = s.Number
	model.LastRanking = s.LastRanking
}

func (s *GroupInsertReq) GetId() interface{} {
	return s.Id
}

type GroupUpdateReq struct {
	GroupReq
	common.ControlBy
}

func (s *GroupUpdateReq) Generate(model *models.Group) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Icon = s.Icon
	model.Number = s.Number
	model.LastRanking = s.LastRanking
}

func (s *GroupUpdateReq) GetId() interface{} {
	return s.Id
}

// GroupGetReq 功能获取请求参数
type GroupGetReq struct {
	Id int `uri:"id"`
}

func (s *GroupGetReq) GetId() interface{} {
	return s.Id
}

// GroupDeleteReq 功能删除请求参数
type GroupDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GroupDeleteReq) GetId() interface{} {
	return s.Ids
}
