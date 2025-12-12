package configure

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ MonsterRouter }

var monsterApi = api.ApiGroupApp.ConfigureApiGroup.MonsterApi
