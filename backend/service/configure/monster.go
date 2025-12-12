package configure

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
)

type MonsterService struct{}

// CreateMonster 创建怪物配置记录
// Author [yourname](https://github.com/yourname)
func (monsterService *MonsterService) CreateMonster(ctx context.Context, monster *configure.Monster) (err error) {
	err = global.GVA_DB.Create(monster).Error
	return err
}

// DeleteMonster 删除怪物配置记录
// Author [yourname](https://github.com/yourname)
func (monsterService *MonsterService) DeleteMonster(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&configure.Monster{}, "id = ?", ID).Error
	return err
}

// DeleteMonsterByIds 批量删除怪物配置记录
// Author [yourname](https://github.com/yourname)
func (monsterService *MonsterService) DeleteMonsterByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]configure.Monster{}, "id in ?", IDs).Error
	return err
}

// UpdateMonster 更新怪物配置记录
// Author [yourname](https://github.com/yourname)
func (monsterService *MonsterService) UpdateMonster(ctx context.Context, monster configure.Monster) (err error) {
	err = global.GVA_DB.Model(&configure.Monster{}).Where("id = ?", monster.ID).Updates(&monster).Error
	return err
}

// GetMonster 根据ID获取怪物配置记录
// Author [yourname](https://github.com/yourname)
func (monsterService *MonsterService) GetMonster(ctx context.Context, ID string) (monster configure.Monster, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&monster).Error
	return
}

// GetMonsterInfoList 分页获取怪物配置记录
// Author [yourname](https://github.com/yourname)
func (monsterService *MonsterService) GetMonsterInfoList(ctx context.Context, info configureReq.MonsterSearch) (list []configure.Monster, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&configure.Monster{})
	var monsters []configure.Monster
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

	err = db.Find(&monsters).Error
	return monsters, total, err
}
func (monsterService *MonsterService) GetMonsterPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
