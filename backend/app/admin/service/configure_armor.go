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

type ConfigureArmor struct {
	service.Service
}

// GetPage 获取ConfigureArmor列表
func (e *ConfigureArmor) GetPage(c *dto.ConfigureArmorGetPageReq, p *actions.DataPermission, list *[]models.ConfigureArmor, count *int64) error {
	var err error
	var data models.ConfigureArmor

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ConfigureArmorService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ConfigureArmor对象
func (e *ConfigureArmor) Get(d *dto.ConfigureArmorGetReq, p *actions.DataPermission, model *models.ConfigureArmor) error {
	var data models.ConfigureArmor

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetConfigureArmor error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ConfigureArmor对象
func (e *ConfigureArmor) Insert(c *dto.ConfigureArmorInsertReq) error {
	var err error
	var data models.ConfigureArmor
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ConfigureArmorService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ConfigureArmor对象
func (e *ConfigureArmor) Update(c *dto.ConfigureArmorUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ConfigureArmor{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ConfigureArmorService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ConfigureArmor
func (e *ConfigureArmor) Remove(d *dto.ConfigureArmorDeleteReq, p *actions.DataPermission) error {
	var data models.ConfigureArmor

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveConfigureArmor error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
