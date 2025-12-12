package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BulletApi struct{}

// CreateBullet 创建子弹配置
// @Tags Bullet
// @Summary 创建子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Bullet true "创建子弹配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bullet/createBullet [post]
func (bulletApi *BulletApi) CreateBullet(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var bullet configure.Bullet
	err := c.ShouldBindJSON(&bullet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bulletService.CreateBullet(ctx, &bullet)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBullet 删除子弹配置
// @Tags Bullet
// @Summary 删除子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Bullet true "删除子弹配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bullet/deleteBullet [delete]
func (bulletApi *BulletApi) DeleteBullet(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := bulletService.DeleteBullet(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBulletByIds 批量删除子弹配置
// @Tags Bullet
// @Summary 批量删除子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bullet/deleteBulletByIds [delete]
func (bulletApi *BulletApi) DeleteBulletByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := bulletService.DeleteBulletByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBullet 更新子弹配置
// @Tags Bullet
// @Summary 更新子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Bullet true "更新子弹配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bullet/updateBullet [put]
func (bulletApi *BulletApi) UpdateBullet(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var bullet configure.Bullet
	err := c.ShouldBindJSON(&bullet)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = bulletService.UpdateBullet(ctx, bullet)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBullet 用id查询子弹配置
// @Tags Bullet
// @Summary 用id查询子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询子弹配置"
// @Success 200 {object} response.Response{data=configure.Bullet,msg=string} "查询成功"
// @Router /bullet/findBullet [get]
func (bulletApi *BulletApi) FindBullet(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rebullet, err := bulletService.GetBullet(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebullet, c)
}

// GetBulletList 分页获取子弹配置列表
// @Tags Bullet
// @Summary 分页获取子弹配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.BulletSearch true "分页获取子弹配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bullet/getBulletList [get]
func (bulletApi *BulletApi) GetBulletList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo configureReq.BulletSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bulletService.GetBulletInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetBulletPublic 不需要鉴权的子弹配置接口
// @Tags Bullet
// @Summary 不需要鉴权的子弹配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bullet/getBulletPublic [get]
func (bulletApi *BulletApi) GetBulletPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	bulletService.GetBulletPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的子弹配置接口信息",
	}, "获取成功", c)
}
