package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MonsterRouter struct{}

// InitMonsterRouter 初始化 怪物配置 路由信息
func (s *MonsterRouter) InitMonsterRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	monsterRouter := Router.Group("monster").Use(middleware.OperationRecord())
	monsterRouterWithoutRecord := Router.Group("monster")
	monsterRouterWithoutAuth := PublicRouter.Group("monster")
	{
		monsterRouter.POST("createMonster", monsterApi.CreateMonster)             // 新建怪物配置
		monsterRouter.DELETE("deleteMonster", monsterApi.DeleteMonster)           // 删除怪物配置
		monsterRouter.DELETE("deleteMonsterByIds", monsterApi.DeleteMonsterByIds) // 批量删除怪物配置
		monsterRouter.PUT("updateMonster", monsterApi.UpdateMonster)              // 更新怪物配置
	}
	{
		monsterRouterWithoutRecord.GET("findMonster", monsterApi.FindMonster)       // 根据ID获取怪物配置
		monsterRouterWithoutRecord.GET("getMonsterList", monsterApi.GetMonsterList) // 获取怪物配置列表
	}
	{
		monsterRouterWithoutAuth.GET("getMonsterPublic", monsterApi.GetMonsterPublic) // 怪物配置开放接口
	}
}
