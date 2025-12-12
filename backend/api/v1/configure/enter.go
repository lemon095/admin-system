package configure

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ MonsterApi }

var monsterService = service.ServiceGroupApp.ConfigureServiceGroup.MonsterService
