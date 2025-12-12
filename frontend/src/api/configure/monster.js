import service from '@/utils/request'
// @Tags Monster
// @Summary 创建怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster true "创建怪物配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /monster/createMonster [post]
export const createMonster = (data) => {
  return service({
    url: '/monster/createMonster',
    method: 'post',
    data
  })
}

// @Tags Monster
// @Summary 删除怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster true "删除怪物配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monster/deleteMonster [delete]
export const deleteMonster = (params) => {
  return service({
    url: '/monster/deleteMonster',
    method: 'delete',
    params
  })
}

// @Tags Monster
// @Summary 批量删除怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除怪物配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monster/deleteMonster [delete]
export const deleteMonsterByIds = (params) => {
  return service({
    url: '/monster/deleteMonsterByIds',
    method: 'delete',
    params
  })
}

// @Tags Monster
// @Summary 更新怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster true "更新怪物配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /monster/updateMonster [put]
export const updateMonster = (data) => {
  return service({
    url: '/monster/updateMonster',
    method: 'put',
    data
  })
}

// @Tags Monster
// @Summary 用id查询怪物配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Monster true "用id查询怪物配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /monster/findMonster [get]
export const findMonster = (params) => {
  return service({
    url: '/monster/findMonster',
    method: 'get',
    params
  })
}

// @Tags Monster
// @Summary 分页获取怪物配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取怪物配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monster/getMonsterList [get]
export const getMonsterList = (params) => {
  return service({
    url: '/monster/getMonsterList',
    method: 'get',
    params
  })
}

// @Tags Monster
// @Summary 不需要鉴权的怪物配置接口
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.MonsterSearch true "分页获取怪物配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /monster/getMonsterPublic [get]
export const getMonsterPublic = () => {
  return service({
    url: '/monster/getMonsterPublic',
    method: 'get',
  })
}
