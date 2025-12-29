package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Item struct {
	service.Service
}

// GetPage 获取Item列表
func (e *Item) GetPage(c *dto.ItemGetPageReq, p *actions.DataPermission, list *[]models.Item, count *int64) error {
	var err error
	var data models.Item

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ItemService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *Item) GetOption() (res []dto.GetOptionResp, err error) {
	var itemTypes []models.ItemType
	if err = e.Orm.Model(&models.ItemType{}).Select("id,name").Find(&itemTypes).Error; err != nil {
		e.Log.Errorf("ItemService itemTypes error:%s \r\n", err)
		return nil, err
	}

	for _, item := range itemTypes {
		res = append(res, dto.GetOptionResp{
			Value: item.Id,
			Label: item.Name,
		})
	}

	var items []models.Item
	err = e.Orm.Model(&models.Item{}).Select("id,name,item_id").Find(&items).Error
	if err != nil {
		e.Log.Errorf("ItemService GetOption error:%s \r\n", err)
		return nil, err
	}

	for k, v := range res {
		for _, i := range items {
			if !strings.HasPrefix(cast.ToString(i.ItemId), cast.ToString(v.Value)) {
				continue
			}
			res[k].Children = append(res[k].Children, dto.GetOptionResp{
				Value: i.ItemId,
				Label: i.Name,
			})
		}
	}
	return res, nil
}

// Get 获取Item对象
func (e *Item) Get(d *dto.ItemGetReq, p *actions.DataPermission, model *models.Item) error {
	var data models.Item

	err := e.Orm.Model(&data).First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetItem error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Item对象
func (e *Item) Insert(c *dto.ItemInsertReq) error {
	var err error
	var data models.Item
	c.Generate(&data)
	err = e.Orm.Omit("item_id").Create(&data).Error
	if err != nil {
		e.Log.Errorf("ItemService Insert error:%s \r\n", err)
		return err
	}

	itemID := cast.ToInt64(fmt.Sprintf("%d%06d", data.Type, data.GetId()))
	err = e.Orm.Model(data).Where("id=?", data.Id).Update("item_id", itemID).Error
	if err != nil {
		e.Log.Errorf("ItemService Insert ItemID error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Item对象
func (e *Item) Update(c *dto.ItemUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Item{}
	e.Orm.First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ItemService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

func (e *Item) UpdateStatus(c *dto.ItemUpdateReq, p *actions.DataPermission) error {
	var err error
	db := e.Orm.Exec("UPDATE bunker_item SET is_enable = NOT is_enable, operator = ? WHERE id = ?", c.Operator, c.Id)
	if err = db.Error; err != nil {
		e.Log.Errorf("ItemService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Item
func (e *Item) Remove(d *dto.ItemDeleteReq, p *actions.DataPermission) error {
	var data models.Item

	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveItem error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
