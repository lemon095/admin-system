package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"

	"gorm.io/datatypes"
)

type ConfigureBulletGetPageReq struct {
	dto.Pagination `search:"-"`
	ConfigureBulletOrder
}

type ConfigureBulletOrder struct {
	Id        string         `form:"idOrder"  search:"type:order;column:id;table:configure_bullet"`
	CreatedAt string         `form:"createdAtOrder"  search:"type:order;column:created_at;table:configure_bullet"`
	UpdatedAt string         `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:configure_bullet"`
	DeletedAt string         `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:configure_bullet"`
	Name      string         `form:"nameOrder"  search:"type:order;column:name;table:configure_bullet"`
	Value     datatypes.JSON `form:"valueOrder"  search:"type:order;column:value;table:configure_bullet"`
}

func (m *ConfigureBulletGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ConfigureBulletInsertReq struct {
	Id    int            `json:"-" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureBulletInsertReq) Generate(model *models.ConfigureBullet) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureBulletInsertReq) GetId() interface{} {
	return s.Id
}

type ConfigureBulletUpdateReq struct {
	Id    int            `uri:"id" comment:""` //
	Name  string         `json:"name" comment:""`
	Value datatypes.JSON `json:"value" comment:""`
	common.ControlBy
}

func (s *ConfigureBulletUpdateReq) Generate(model *models.ConfigureBullet) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Name = s.Name
	model.Value = s.Value
}

func (s *ConfigureBulletUpdateReq) GetId() interface{} {
	return s.Id
}

// ConfigureBulletGetReq 功能获取请求参数
type ConfigureBulletGetReq struct {
	Id int `uri:"id"`
}

func (s *ConfigureBulletGetReq) GetId() interface{} {
	return s.Id
}

// ConfigureBulletDeleteReq 功能删除请求参数
type ConfigureBulletDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ConfigureBulletDeleteReq) GetId() interface{} {
	return s.Ids
}
