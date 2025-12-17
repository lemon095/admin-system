package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type ConfigureWeapon struct {
	service.Service
}

// GetPage 获取ConfigureWeapon列表
func (e *ConfigureWeapon) GetPage(c *dto.ConfigureWeaponGetPageReq, p *actions.DataPermission, list *[]models.ConfigureWeapon, count *int64) error {
	var err error
	var data models.ConfigureWeapon

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ConfigureWeaponService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ConfigureWeapon对象
func (e *ConfigureWeapon) Get(d *dto.ConfigureWeaponGetReq, p *actions.DataPermission, model *models.ConfigureWeapon) error {
	var data models.ConfigureWeapon

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetConfigureWeapon error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ConfigureWeapon对象
func (e *ConfigureWeapon) Insert(c *dto.ConfigureWeaponInsertReq) error {
	var err error
	var data models.ConfigureWeapon
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ConfigureWeaponService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ConfigureWeapon对象
func (e *ConfigureWeapon) Update(c *dto.ConfigureWeaponUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ConfigureWeapon{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ConfigureWeaponService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ConfigureWeapon
func (e *ConfigureWeapon) Remove(d *dto.ConfigureWeaponDeleteReq, p *actions.DataPermission) error {
	var data models.ConfigureWeapon

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveConfigureWeapon error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
