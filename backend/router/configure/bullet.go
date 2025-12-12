package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BulletRouter struct{}

// InitBulletRouter 初始化 子弹配置 路由信息
func (s *BulletRouter) InitBulletRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	bulletRouter := Router.Group("bullet").Use(middleware.OperationRecord())
	bulletRouterWithoutRecord := Router.Group("bullet")
	bulletRouterWithoutAuth := PublicRouter.Group("bullet")
	{
		bulletRouter.POST("createBullet", bulletApi.CreateBullet)             // 新建子弹配置
		bulletRouter.DELETE("deleteBullet", bulletApi.DeleteBullet)           // 删除子弹配置
		bulletRouter.DELETE("deleteBulletByIds", bulletApi.DeleteBulletByIds) // 批量删除子弹配置
		bulletRouter.PUT("updateBullet", bulletApi.UpdateBullet)              // 更新子弹配置
	}
	{
		bulletRouterWithoutRecord.GET("findBullet", bulletApi.FindBullet)       // 根据ID获取子弹配置
		bulletRouterWithoutRecord.GET("getBulletList", bulletApi.GetBulletList) // 获取子弹配置列表
	}
	{
		bulletRouterWithoutAuth.GET("getBulletPublic", bulletApi.GetBulletPublic) // 子弹配置开放接口
	}
}
