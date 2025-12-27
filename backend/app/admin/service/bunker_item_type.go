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

type ItemType struct {
	service.Service
}

// GetPage 获取ItemType列表
func (e *ItemType) GetPage(c *dto.ItemTypeGetPageReq, p *actions.DataPermission, list *[]models.ItemType, count *int64) error {
	var err error
	var data models.ItemType

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ItemTypeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ItemType对象
func (e *ItemType) Get(d *dto.ItemTypeGetReq, p *actions.DataPermission, model *models.ItemType) error {
	var data models.ItemType

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetItemType error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ItemType对象
func (e *ItemType) Insert(c *dto.ItemTypeInsertReq) error {
	var err error
	var data models.ItemType
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ItemTypeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ItemType对象
func (e *ItemType) Update(c *dto.ItemTypeUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ItemType{}
	c.Generate(&data)

	db := e.Orm.Where("id=?", c.Id).Updates(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ItemTypeService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ItemType
func (e *ItemType) Remove(d *dto.ItemTypeDeleteReq, p *actions.DataPermission) error {
	var data models.ItemType

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveItemType error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
