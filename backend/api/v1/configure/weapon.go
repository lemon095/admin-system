package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WeaponApi struct{}

// CreateWeapon 创建武器配置
// @Tags Weapon
// @Summary 创建武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Weapon true "创建武器配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /weapon/createWeapon [post]
func (weaponApi *WeaponApi) CreateWeapon(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var weapon configure.Weapon
	err := c.ShouldBindJSON(&weapon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = weaponService.CreateWeapon(ctx, &weapon)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWeapon 删除武器配置
// @Tags Weapon
// @Summary 删除武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Weapon true "删除武器配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /weapon/deleteWeapon [delete]
func (weaponApi *WeaponApi) DeleteWeapon(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := weaponService.DeleteWeapon(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteWeaponByIds 批量删除武器配置
// @Tags Weapon
// @Summary 批量删除武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /weapon/deleteWeaponByIds [delete]
func (weaponApi *WeaponApi) DeleteWeaponByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := weaponService.DeleteWeaponByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateWeapon 更新武器配置
// @Tags Weapon
// @Summary 更新武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Weapon true "更新武器配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /weapon/updateWeapon [put]
func (weaponApi *WeaponApi) UpdateWeapon(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var weapon configure.Weapon
	err := c.ShouldBindJSON(&weapon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = weaponService.UpdateWeapon(ctx, weapon)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindWeapon 用id查询武器配置
// @Tags Weapon
// @Summary 用id查询武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询武器配置"
// @Success 200 {object} response.Response{data=configure.Weapon,msg=string} "查询成功"
// @Router /weapon/findWeapon [get]
func (weaponApi *WeaponApi) FindWeapon(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reweapon, err := weaponService.GetWeapon(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reweapon, c)
}

// GetWeaponList 分页获取武器配置列表
// @Tags Weapon
// @Summary 分页获取武器配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.WeaponSearch true "分页获取武器配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /weapon/getWeaponList [get]
func (weaponApi *WeaponApi) GetWeaponList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo configureReq.WeaponSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := weaponService.GetWeaponInfoList(ctx, pageInfo)
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

// GetWeaponPublic 不需要鉴权的武器配置接口
// @Tags Weapon
// @Summary 不需要鉴权的武器配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /weapon/getWeaponPublic [get]
func (weaponApi *WeaponApi) GetWeaponPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	weaponService.GetWeaponPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的武器配置接口信息",
	}, "获取成功", c)
}
