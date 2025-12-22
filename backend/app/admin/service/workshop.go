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

type Workshop struct {
	service.Service
}

// GetPage 获取Workshop列表
func (e *Workshop) GetPage(c *dto.WorkshopGetPageReq, p *actions.DataPermission, list *[]models.Workshop, count *int64) error {
	var err error
	var data models.Workshop

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("WorkshopService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Workshop对象
func (e *Workshop) Get(d *dto.WorkshopGetReq, p *actions.DataPermission, model *models.Workshop) error {
	var data models.Workshop

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetWorkshop error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Workshop对象
func (e *Workshop) Insert(c *dto.WorkshopInsertReq) error {
	var err error
	var data models.Workshop
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("WorkshopService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Workshop对象
func (e *Workshop) Update(c *dto.WorkshopUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Workshop{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("WorkshopService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Workshop
func (e *Workshop) Remove(d *dto.WorkshopDeleteReq, p *actions.DataPermission) error {
	var data models.Workshop

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveWorkshop error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
