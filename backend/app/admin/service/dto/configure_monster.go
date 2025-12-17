package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"

	"gorm.io/datatypes"
)

type ConfigureMonsterGetPageReq struct {
	dto.Pagination `search:"-"`
	ConfigureMonsterOrder
}

type ConfigureMonsterOrder struct {
	Id        string         `form:"idOrder"  search:"type:order;column:id;table:configure_monster"`
	CreatedAt string         `form:"createdAtOrder"  search:"type:order;column:created_at;table:configure_monster"`
	UpdatedAt string         `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:configure_monster"`
	DeletedAt string         `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:configure_monster"`
	Name      string         `form:"nameOrder"  search:"type:order;column:name;table:configure_monster"`
	Value     datatypes.JSON `form:"valueOrder"  search:"type:order;column:value;table:configure_monster"`
}

func (m *ConfigureMonsterGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ConfigureMonsterInsertReq struct {
	Id    int            `json:"-" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureMonsterInsertReq) Generate(model *models.ConfigureMonster) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureMonsterInsertReq) GetId() interface{} {
	return s.Id
}

type ConfigureMonsterUpdateReq struct {
	Id    int            `uri:"id" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureMonsterUpdateReq) Generate(model *models.ConfigureMonster) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureMonsterUpdateReq) GetId() interface{} {
	return s.Id
}

// ConfigureMonsterGetReq 功能获取请求参数
type ConfigureMonsterGetReq struct {
	Id int `uri:"id"`
}

func (s *ConfigureMonsterGetReq) GetId() interface{} {
	return s.Id
}

// ConfigureMonsterDeleteReq 功能删除请求参数
type ConfigureMonsterDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ConfigureMonsterDeleteReq) GetId() interface{} {
	return s.Ids
}
