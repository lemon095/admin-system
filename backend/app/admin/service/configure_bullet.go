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

type ConfigureBullet struct {
	service.Service
}

// GetPage 获取ConfigureBullet列表
func (e *ConfigureBullet) GetPage(c *dto.ConfigureBulletGetPageReq, p *actions.DataPermission, list *[]models.ConfigureBullet, count *int64) error {
	var err error
	var data models.ConfigureBullet

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ConfigureBulletService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ConfigureBullet对象
func (e *ConfigureBullet) Get(d *dto.ConfigureBulletGetReq, p *actions.DataPermission, model *models.ConfigureBullet) error {
	var data models.ConfigureBullet

	err := e.Orm.Model(&data).First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetConfigureBullet error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ConfigureBullet对象
func (e *ConfigureBullet) Insert(c *dto.ConfigureBulletInsertReq) error {
	var err error
	var data models.ConfigureBullet
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ConfigureBulletService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ConfigureBullet对象
func (e *ConfigureBullet) Update(c *dto.ConfigureBulletUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ConfigureBullet{}
	e.Orm.First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ConfigureBulletService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ConfigureBullet
func (e *ConfigureBullet) Remove(d *dto.ConfigureBulletDeleteReq, p *actions.DataPermission) error {
	var data models.ConfigureBullet

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveConfigureBullet error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
