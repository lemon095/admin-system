package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WeaponRouter struct{}

// InitWeaponRouter 初始化 武器配置 路由信息
func (s *WeaponRouter) InitWeaponRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	weaponRouter := Router.Group("weapon").Use(middleware.OperationRecord())
	weaponRouterWithoutRecord := Router.Group("weapon")
	weaponRouterWithoutAuth := PublicRouter.Group("weapon")
	{
		weaponRouter.POST("createWeapon", weaponApi.CreateWeapon)             // 新建武器配置
		weaponRouter.DELETE("deleteWeapon", weaponApi.DeleteWeapon)           // 删除武器配置
		weaponRouter.DELETE("deleteWeaponByIds", weaponApi.DeleteWeaponByIds) // 批量删除武器配置
		weaponRouter.PUT("updateWeapon", weaponApi.UpdateWeapon)              // 更新武器配置
	}
	{
		weaponRouterWithoutRecord.GET("findWeapon", weaponApi.FindWeapon)       // 根据ID获取武器配置
		weaponRouterWithoutRecord.GET("getWeaponList", weaponApi.GetWeaponList) // 获取武器配置列表
	}
	{
		weaponRouterWithoutAuth.GET("getWeaponPublic", weaponApi.GetWeaponPublic) // 武器配置开放接口
	}
}
