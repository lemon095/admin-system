import service from '@/utils/request'
// @Tags Monster1
// @Summary 创建怪物1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster1 true "创建怪物1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /monster1/createMonster1 [post]
export const createMonster1 = (data) => {
  return service({
    url: '/monster1/createMonster1',
    method: 'post',
    data
  })
}

// @Tags Monster1
// @Summary 删除怪物1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster1 true "删除怪物1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monster1/deleteMonster1 [delete]
export const deleteMonster1 = (params) => {
  return service({
    url: '/monster1/deleteMonster1',
    method: 'delete',
    params
  })
}

// @Tags Monster1
// @Summary 批量删除怪物1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除怪物1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monster1/deleteMonster1 [delete]
export const deleteMonster1ByIds = (params) => {
  return service({
    url: '/monster1/deleteMonster1ByIds',
    method: 'delete',
    params
  })
}

// @Tags Monster1
// @Summary 更新怪物1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster1 true "更新怪物1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /monster1/updateMonster1 [put]
export const updateMonster1 = (data) => {
  return service({
    url: '/monster1/updateMonster1',
    method: 'put',
    data
  })
}

// @Tags Monster1
// @Summary 用id查询怪物1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Monster1 true "用id查询怪物1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /monster1/findMonster1 [get]
export const findMonster1 = (params) => {
  return service({
    url: '/monster1/findMonster1',
    method: 'get',
    params
  })
}

// @Tags Monster1
// @Summary 分页获取怪物1列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取怪物1列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monster1/getMonster1List [get]
export const getMonster1List = (params) => {
  return service({
    url: '/monster1/getMonster1List',
    method: 'get',
    params
  })
}

// @Tags Monster1
// @Summary 不需要鉴权的怪物1接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.Monster1Search true "分页获取怪物1列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /monster1/getMonster1Public [get]
export const getMonster1Public = () => {
  return service({
    url: '/monster1/getMonster1Public',
    method: 'get',
  })
}
