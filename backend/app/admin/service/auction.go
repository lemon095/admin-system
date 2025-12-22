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

type Auction struct {
	service.Service
}

// GetPage 获取Auction列表
func (e *Auction) GetPage(c *dto.AuctionGetPageReq, p *actions.DataPermission, list *[]models.Auction, count *int64) error {
	var err error
	var data models.Auction

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("AuctionService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Auction对象
func (e *Auction) Get(d *dto.AuctionGetReq, p *actions.DataPermission, model *models.Auction) error {
	var data models.Auction

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetAuction error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Auction对象
func (e *Auction) Insert(c *dto.AuctionInsertReq) error {
	var err error
	var data models.Auction
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("AuctionService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Auction对象
func (e *Auction) Update(c *dto.AuctionUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Auction{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("AuctionService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Auction
func (e *Auction) Remove(d *dto.AuctionDeleteReq, p *actions.DataPermission) error {
	var data models.Auction

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveAuction error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
