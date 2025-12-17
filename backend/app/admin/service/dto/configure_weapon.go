package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"

	"gorm.io/datatypes"
)

type ConfigureWeaponGetPageReq struct {
	dto.Pagination `search:"-"`
	ConfigureWeaponOrder
}

type ConfigureWeaponOrder struct {
	Id        string         `form:"idOrder"  search:"type:order;column:id;table:configure_weapon"`
	CreatedAt string         `form:"createdAtOrder"  search:"type:order;column:created_at;table:configure_weapon"`
	UpdatedAt string         `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:configure_weapon"`
	DeletedAt string         `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:configure_weapon"`
	Name      string         `form:"nameOrder"  search:"type:order;column:name;table:configure_weapon"`
	Value     datatypes.JSON `form:"valueOrder"  search:"type:order;column:value;table:configure_weapon"`
}

func (m *ConfigureWeaponGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ConfigureWeaponInsertReq struct {
	Id    int            `json:"-" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureWeaponInsertReq) Generate(model *models.ConfigureWeapon) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureWeaponInsertReq) GetId() interface{} {
	return s.Id
}

type ConfigureWeaponUpdateReq struct {
	Id    int            `uri:"id" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureWeaponUpdateReq) Generate(model *models.ConfigureWeapon) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureWeaponUpdateReq) GetId() interface{} {
	return s.Id
}

// ConfigureWeaponGetReq 功能获取请求参数
type ConfigureWeaponGetReq struct {
	Id int `uri:"id"`
}

func (s *ConfigureWeaponGetReq) GetId() interface{} {
	return s.Id
}

// ConfigureWeaponDeleteReq 功能删除请求参数
type ConfigureWeaponDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ConfigureWeaponDeleteReq) GetId() interface{} {
	return s.Ids
}
