package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"

	"gorm.io/datatypes"
)

type ConfigureArmorGetPageReq struct {
	dto.Pagination `search:"-"`
	ConfigureArmorOrder
}

type ConfigureArmorOrder struct {
	Id        string         `form:"idOrder"  search:"type:order;column:id;table:configure_armor"`
	CreatedAt string         `form:"createdAtOrder"  search:"type:order;column:created_at;table:configure_armor"`
	UpdatedAt string         `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:configure_armor"`
	DeletedAt string         `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:configure_armor"`
	Name      string         `form:"nameOrder"  search:"type:order;column:name;table:configure_armor"`
	Value     datatypes.JSON `form:"valueOrder"  search:"type:order;column:value;table:configure_armor"`
}

func (m *ConfigureArmorGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ConfigureArmorInsertReq struct {
	Id    int            `json:"-" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureArmorInsertReq) Generate(model *models.ConfigureArmor) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureArmorInsertReq) GetId() interface{} {
	return s.Id
}

type ConfigureArmorUpdateReq struct {
	Id    int            `uri:"id" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureArmorUpdateReq) Generate(model *models.ConfigureArmor) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureArmorUpdateReq) GetId() interface{} {
	return s.Id
}

// ConfigureArmorGetReq 功能获取请求参数
type ConfigureArmorGetReq struct {
	Id int `uri:"id"`
}

func (s *ConfigureArmorGetReq) GetId() interface{} {
	return s.Id
}

// ConfigureArmorDeleteReq 功能删除请求参数
type ConfigureArmorDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ConfigureArmorDeleteReq) GetId() interface{} {
	return s.Ids
}
