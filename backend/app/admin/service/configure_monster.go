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

type ConfigureMonster struct {
	service.Service
}

// GetPage 获取ConfigureMonster列表
func (e *ConfigureMonster) GetPage(c *dto.ConfigureMonsterGetPageReq, p *actions.DataPermission, list *[]models.ConfigureMonster, count *int64) error {
	var err error
	var data models.ConfigureMonster

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ConfigureMonsterService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ConfigureMonster对象
func (e *ConfigureMonster) Get(d *dto.ConfigureMonsterGetReq, p *actions.DataPermission, model *models.ConfigureMonster) error {
	var data models.ConfigureMonster

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetConfigureMonster error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ConfigureMonster对象
func (e *ConfigureMonster) Insert(c *dto.ConfigureMonsterInsertReq) error {
	var err error
	var data models.ConfigureMonster
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ConfigureMonsterService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ConfigureMonster对象
func (e *ConfigureMonster) Update(c *dto.ConfigureMonsterUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ConfigureMonster{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ConfigureMonsterService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ConfigureMonster
func (e *ConfigureMonster) Remove(d *dto.ConfigureMonsterDeleteReq, p *actions.DataPermission) error {
	var data models.ConfigureMonster

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveConfigureMonster error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
