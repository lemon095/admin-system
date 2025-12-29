package service

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type GiftPack struct {
	service.Service
}

// GetPage 获取GiftPack列表
func (e *GiftPack) GetPage(c *dto.GiftPackGetPageReq, p *actions.DataPermission, list *[]models.GiftPack, count *int64) error {
	var err error
	var data models.GiftPack
	db := e.Orm.Model(&data)
	if c.IsEnable != nil {
		db = db.Where("is_enable = ?", c.IsEnable)
	}
	if c.GiftName != nil {
		db = db.Where("gift_name = ?", c.GiftName)
	}
	if c.RedeemCode != nil {
		db = db.Where("redeem_code = ?", c.RedeemCode)
	}

	err = db.Scopes(
		cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
	).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("GiftPackService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取GiftPack对象
func (e *GiftPack) Get(d *dto.GiftPackGetReq, p *actions.DataPermission, model *models.GiftPack) error {
	var data models.GiftPack

	err := e.Orm.Model(&data).First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetGiftPack error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建GiftPack对象
func (e *GiftPack) Insert(c *dto.GiftPackInsertReq) error {
	var err error
	var data models.GiftPack
	c.Generate(&data)
	data.RedeemCode = Random12StringSecure()
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("GiftPackService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改GiftPack对象
func (e *GiftPack) Update(c *dto.GiftPackUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.GiftPack{}
	e.Orm.First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Where("id=?", data.Id).Updates(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("GiftPackService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

func (e *GiftPack) UpdateStatus(c *dto.GiftPackUpdateReq, p *actions.DataPermission) error {
	var err error
	db := e.Orm.Exec("UPDATE bunker_gift_pack SET is_enable = NOT is_enable WHERE id = ?", c.Id)
	if err = db.Error; err != nil {
		e.Log.Errorf("GiftPackService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除GiftPack
func (e *GiftPack) Remove(d *dto.GiftPackDeleteReq, p *actions.DataPermission) error {
	var data models.GiftPack

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveGiftPack error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

func Random12StringSecure() string {
	firstChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	otherChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, 12)

	// 首字母大写
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(firstChars))))
	result[0] = firstChars[n.Int64()]

	for i := 1; i < 12; i++ {
		n, _ = rand.Int(rand.Reader, big.NewInt(int64(len(otherChars))))
		result[i] = otherChars[n.Int64()]
	}

	return string(result)
}
