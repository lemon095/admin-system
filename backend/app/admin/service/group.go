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

type Group struct {
	service.Service
}

// GetPage 获取Group列表
func (e *Group) GetPage(c *dto.GroupGetPageReq, p *actions.DataPermission, list *[]models.Group, count *int64) error {
	var err error
	var data models.Group

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("GroupService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Group对象
func (e *Group) Get(d *dto.GroupGetReq, p *actions.DataPermission, model *models.Group) error {
	var data models.Group

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetGroup error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Group对象
func (e *Group) Insert(c *dto.GroupInsertReq) error {
	var err error
	var data models.Group
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("GroupService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Group对象
func (e *Group) Update(c *dto.GroupUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Group{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("GroupService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Group
func (e *Group) Remove(d *dto.GroupDeleteReq, p *actions.DataPermission) error {
	var data models.Group

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveGroup error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
