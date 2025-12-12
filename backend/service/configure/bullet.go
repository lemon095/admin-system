package configure

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
)

type BulletService struct{}

// CreateBullet 创建子弹配置记录
// Author [yourname](https://github.com/yourname)
func (bulletService *BulletService) CreateBullet(ctx context.Context, bullet *configure.Bullet) (err error) {
	err = global.GVA_DB.Create(bullet).Error
	return err
}

// DeleteBullet 删除子弹配置记录
// Author [yourname](https://github.com/yourname)
func (bulletService *BulletService) DeleteBullet(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&configure.Bullet{}, "id = ?", ID).Error
	return err
}

// DeleteBulletByIds 批量删除子弹配置记录
// Author [yourname](https://github.com/yourname)
func (bulletService *BulletService) DeleteBulletByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]configure.Bullet{}, "id in ?", IDs).Error
	return err
}

// UpdateBullet 更新子弹配置记录
// Author [yourname](https://github.com/yourname)
func (bulletService *BulletService) UpdateBullet(ctx context.Context, bullet configure.Bullet) (err error) {
	err = global.GVA_DB.Model(&configure.Bullet{}).Where("id = ?", bullet.ID).Updates(&bullet).Error
	return err
}

// GetBullet 根据ID获取子弹配置记录
// Author [yourname](https://github.com/yourname)
func (bulletService *BulletService) GetBullet(ctx context.Context, ID string) (bullet configure.Bullet, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bullet).Error
	return
}

// GetBulletInfoList 分页获取子弹配置记录
// Author [yourname](https://github.com/yourname)
func (bulletService *BulletService) GetBulletInfoList(ctx context.Context, info configureReq.BulletSearch) (list []configure.Bullet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&configure.Bullet{})
	var bullets []configure.Bullet
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

	err = db.Find(&bullets).Error
	return bullets, total, err
}
func (bulletService *BulletService) GetBulletPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
