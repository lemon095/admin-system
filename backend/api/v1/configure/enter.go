package configure

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	MonsterApi
	WeaponApi
	BulletApi
	ArmorApi
}

var (
	monsterService = service.ServiceGroupApp.ConfigureServiceGroup.MonsterService
	weaponService  = service.ServiceGroupApp.ConfigureServiceGroup.WeaponService
	bulletService  = service.ServiceGroupApp.ConfigureServiceGroup.BulletService
	armorService   = service.ServiceGroupApp.ConfigureServiceGroup.ArmorService
)
