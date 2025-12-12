package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArmorApi struct{}

// CreateArmor 创建防具配置
// @Tags Armor
// @Summary 创建防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Armor true "创建防具配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /armor/createArmor [post]
func (armorApi *ArmorApi) CreateArmor(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var armor configure.Armor
	err := c.ShouldBindJSON(&armor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = armorService.CreateArmor(ctx, &armor)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteArmor 删除防具配置
// @Tags Armor
// @Summary 删除防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Armor true "删除防具配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /armor/deleteArmor [delete]
func (armorApi *ArmorApi) DeleteArmor(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := armorService.DeleteArmor(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteArmorByIds 批量删除防具配置
// @Tags Armor
// @Summary 批量删除防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /armor/deleteArmorByIds [delete]
func (armorApi *ArmorApi) DeleteArmorByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := armorService.DeleteArmorByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateArmor 更新防具配置
// @Tags Armor
// @Summary 更新防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Armor true "更新防具配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /armor/updateArmor [put]
func (armorApi *ArmorApi) UpdateArmor(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var armor configure.Armor
	err := c.ShouldBindJSON(&armor)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = armorService.UpdateArmor(ctx, armor)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindArmor 用id查询防具配置
// @Tags Armor
// @Summary 用id查询防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询防具配置"
// @Success 200 {object} response.Response{data=configure.Armor,msg=string} "查询成功"
// @Router /armor/findArmor [get]
func (armorApi *ArmorApi) FindArmor(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rearmor, err := armorService.GetArmor(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rearmor, c)
}

// GetArmorList 分页获取防具配置列表
// @Tags Armor
// @Summary 分页获取防具配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.ArmorSearch true "分页获取防具配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /armor/getArmorList [get]
func (armorApi *ArmorApi) GetArmorList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo configureReq.ArmorSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := armorService.GetArmorInfoList(ctx, pageInfo)
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

// GetArmorPublic 不需要鉴权的防具配置接口
// @Tags Armor
// @Summary 不需要鉴权的防具配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /armor/getArmorPublic [get]
func (armorApi *ArmorApi) GetArmorPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	armorService.GetArmorPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的防具配置接口信息",
	}, "获取成功", c)
}
