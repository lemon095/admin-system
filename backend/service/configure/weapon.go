package configure

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
)

type WeaponService struct{}

// CreateWeapon 创建武器配置记录
// Author [yourname](https://github.com/yourname)
func (weaponService *WeaponService) CreateWeapon(ctx context.Context, weapon *configure.Weapon) (err error) {
	err = global.GVA_DB.Create(weapon).Error
	return err
}

// DeleteWeapon 删除武器配置记录
// Author [yourname](https://github.com/yourname)
func (weaponService *WeaponService) DeleteWeapon(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&configure.Weapon{}, "id = ?", ID).Error
	return err
}

// DeleteWeaponByIds 批量删除武器配置记录
// Author [yourname](https://github.com/yourname)
func (weaponService *WeaponService) DeleteWeaponByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]configure.Weapon{}, "id in ?", IDs).Error
	return err
}

// UpdateWeapon 更新武器配置记录
// Author [yourname](https://github.com/yourname)
func (weaponService *WeaponService) UpdateWeapon(ctx context.Context, weapon configure.Weapon) (err error) {
	err = global.GVA_DB.Model(&configure.Weapon{}).Where("id = ?", weapon.ID).Updates(&weapon).Error
	return err
}

// GetWeapon 根据ID获取武器配置记录
// Author [yourname](https://github.com/yourname)
func (weaponService *WeaponService) GetWeapon(ctx context.Context, ID string) (weapon configure.Weapon, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&weapon).Error
	return
}

// GetWeaponInfoList 分页获取武器配置记录
// Author [yourname](https://github.com/yourname)
func (weaponService *WeaponService) GetWeaponInfoList(ctx context.Context, info configureReq.WeaponSearch) (list []configure.Weapon, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&configure.Weapon{})
	var weapons []configure.Weapon
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

	err = db.Find(&weapons).Error
	return weapons, total, err
}
func (weaponService *WeaponService) GetWeaponPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
