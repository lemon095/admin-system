package configure

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
)

type ArmorService struct{}

// CreateArmor 创建防具配置记录
// Author [yourname](https://github.com/yourname)
func (armorService *ArmorService) CreateArmor(ctx context.Context, armor *configure.Armor) (err error) {
	err = global.GVA_DB.Create(armor).Error
	return err
}

// DeleteArmor 删除防具配置记录
// Author [yourname](https://github.com/yourname)
func (armorService *ArmorService) DeleteArmor(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&configure.Armor{}, "id = ?", ID).Error
	return err
}

// DeleteArmorByIds 批量删除防具配置记录
// Author [yourname](https://github.com/yourname)
func (armorService *ArmorService) DeleteArmorByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]configure.Armor{}, "id in ?", IDs).Error
	return err
}

// UpdateArmor 更新防具配置记录
// Author [yourname](https://github.com/yourname)
func (armorService *ArmorService) UpdateArmor(ctx context.Context, armor configure.Armor) (err error) {
	err = global.GVA_DB.Model(&configure.Armor{}).Where("id = ?", armor.ID).Updates(&armor).Error
	return err
}

// GetArmor 根据ID获取防具配置记录
// Author [yourname](https://github.com/yourname)
func (armorService *ArmorService) GetArmor(ctx context.Context, ID string) (armor configure.Armor, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&armor).Error
	return
}

// GetArmorInfoList 分页获取防具配置记录
// Author [yourname](https://github.com/yourname)
func (armorService *ArmorService) GetArmorInfoList(ctx context.Context, info configureReq.ArmorSearch) (list []configure.Armor, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&configure.Armor{})
	var armors []configure.Armor
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&armors).Error
	return armors, total, err
}
func (armorService *ArmorService) GetArmorPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
