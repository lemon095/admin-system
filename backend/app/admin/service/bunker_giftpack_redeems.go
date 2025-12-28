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

type GiftpackRedeems struct {
	service.Service
}

// GetPage 获取GiftpackRedeems列表
func (e *GiftpackRedeems) GetPage(c *dto.GiftpackRedeemsGetPageReq, p *actions.DataPermission, list *[]dto.GiftpackRedeemsGetPageResp, count *int64) error {
	var err error
	var data models.GiftpackRedeems

	err = e.Orm.Model(&data).
		Select("bunker_giftpack_redeems.redeem_code AS redeem_code, bunker_users.nickname AS nickname, bunker_users.gender AS gender, bunker_users.avatar AS avatar").
		Joins("INNER JOIN bunker_users ON bunker_users.union_id=bunker_giftpack_redeems.redeemer_union_id").
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("GiftpackRedeemsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取GiftpackRedeems对象
func (e *GiftpackRedeems) Get(d *dto.GiftpackRedeemsGetReq, p *actions.DataPermission, model *models.GiftpackRedeems) error {
	var data models.GiftpackRedeems

	err := e.Orm.Model(&data).First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetGiftpackRedeems error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建GiftpackRedeems对象
func (e *GiftpackRedeems) Insert(c *dto.GiftpackRedeemsInsertReq) error {
	var err error
	var data models.GiftpackRedeems
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("GiftpackRedeemsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改GiftpackRedeems对象
func (e *GiftpackRedeems) Update(c *dto.GiftpackRedeemsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.GiftpackRedeems{}
	e.Orm.First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("GiftpackRedeemsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除GiftpackRedeems
func (e *GiftpackRedeems) Remove(d *dto.GiftpackRedeemsDeleteReq, p *actions.DataPermission) error {
	var data models.GiftpackRedeems

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveGiftpackRedeems error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
