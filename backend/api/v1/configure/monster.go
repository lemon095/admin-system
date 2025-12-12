package configure

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/configure"
	configureReq "github.com/flipped-aurora/gin-vue-admin/server/model/configure/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MonsterApi struct{}

// CreateMonster 创建怪物配置
// @Tags Monster
// @Summary 创建怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Monster true "创建怪物配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /monster/createMonster [post]
func (monsterApi *MonsterApi) CreateMonster(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var monster configure.Monster
	err := c.ShouldBindJSON(&monster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = monsterService.CreateMonster(ctx, &monster)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMonster 删除怪物配置
// @Tags Monster
// @Summary 删除怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Monster true "删除怪物配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /monster/deleteMonster [delete]
func (monsterApi *MonsterApi) DeleteMonster(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := monsterService.DeleteMonster(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMonsterByIds 批量删除怪物配置
// @Tags Monster
// @Summary 批量删除怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /monster/deleteMonsterByIds [delete]
func (monsterApi *MonsterApi) DeleteMonsterByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := monsterService.DeleteMonsterByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMonster 更新怪物配置
// @Tags Monster
// @Summary 更新怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body configure.Monster true "更新怪物配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /monster/updateMonster [put]
func (monsterApi *MonsterApi) UpdateMonster(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var monster configure.Monster
	err := c.ShouldBindJSON(&monster)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = monsterService.UpdateMonster(ctx, monster)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMonster 用id查询怪物配置
// @Tags Monster
// @Summary 用id查询怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询怪物配置"
// @Success 200 {object} response.Response{data=configure.Monster,msg=string} "查询成功"
// @Router /monster/findMonster [get]
func (monsterApi *MonsterApi) FindMonster(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	remonster, err := monsterService.GetMonster(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remonster, c)
}

// GetMonsterList 分页获取怪物配置列表
// @Tags Monster
// @Summary 分页获取怪物配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.MonsterSearch true "分页获取怪物配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /monster/getMonsterList [get]
func (monsterApi *MonsterApi) GetMonsterList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo configureReq.MonsterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := monsterService.GetMonsterInfoList(ctx, pageInfo)
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

// GetMonsterPublic 不需要鉴权的怪物配置接口
// @Tags Monster
// @Summary 不需要鉴权的怪物配置接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /monster/getMonsterPublic [get]
func (monsterApi *MonsterApi) GetMonsterPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	monsterService.GetMonsterPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的怪物配置接口信息",
	}, "获取成功", c)
}
