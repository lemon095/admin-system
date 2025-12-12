package configure

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	MonsterRouter
	WeaponRouter
	BulletRouter
	ArmorRouter
}

var (
	monsterApi = api.ApiGroupApp.ConfigureApiGroup.MonsterApi
	weaponApi  = api.ApiGroupApp.ConfigureApiGroup.WeaponApi
	bulletApi  = api.ApiGroupApp.ConfigureApiGroup.BulletApi
	armorApi   = api.ApiGroupApp.ConfigureApiGroup.ArmorApi
)
