package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArmorRouter struct{}

// InitArmorRouter 初始化 防具配置 路由信息
func (s *ArmorRouter) InitArmorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	armorRouter := Router.Group("armor").Use(middleware.OperationRecord())
	armorRouterWithoutRecord := Router.Group("armor")
	armorRouterWithoutAuth := PublicRouter.Group("armor")
	{
		armorRouter.POST("createArmor", armorApi.CreateArmor)             // 新建防具配置
		armorRouter.DELETE("deleteArmor", armorApi.DeleteArmor)           // 删除防具配置
		armorRouter.DELETE("deleteArmorByIds", armorApi.DeleteArmorByIds) // 批量删除防具配置
		armorRouter.PUT("updateArmor", armorApi.UpdateArmor)              // 更新防具配置
	}
	{
		armorRouterWithoutRecord.GET("findArmor", armorApi.FindArmor)       // 根据ID获取防具配置
		armorRouterWithoutRecord.GET("getArmorList", armorApi.GetArmorList) // 获取防具配置列表
	}
	{
		armorRouterWithoutAuth.GET("getArmorPublic", armorApi.GetArmorPublic) // 防具配置开放接口
	}
}
