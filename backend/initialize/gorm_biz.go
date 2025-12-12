package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(configure.Monster{}, configure.Bullet{}, configure.Armor{}, configure.Weapon{})
	if err != nil {
		return err
	}
	return nil
}
