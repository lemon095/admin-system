import service from '@/utils/request'
// @Tags Monster2
// @Summary 创建monster2
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster2 true "创建monster2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /monster2/createMonster2 [post]
export const createMonster2 = (data) => {
  return service({
    url: '/monster2/createMonster2',
    method: 'post',
    data
  })
}

// @Tags Monster2
// @Summary 删除monster2
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster2 true "删除monster2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monster2/deleteMonster2 [delete]
export const deleteMonster2 = (params) => {
  return service({
    url: '/monster2/deleteMonster2',
    method: 'delete',
    params
  })
}

// @Tags Monster2
// @Summary 批量删除monster2
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除monster2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /monster2/deleteMonster2 [delete]
export const deleteMonster2ByIds = (params) => {
  return service({
    url: '/monster2/deleteMonster2ByIds',
    method: 'delete',
    params
  })
}

// @Tags Monster2
// @Summary 更新monster2
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Monster2 true "更新monster2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /monster2/updateMonster2 [put]
export const updateMonster2 = (data) => {
  return service({
    url: '/monster2/updateMonster2',
    method: 'put',
    data
  })
}

// @Tags Monster2
// @Summary 用id查询monster2
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Monster2 true "用id查询monster2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /monster2/findMonster2 [get]
export const findMonster2 = (params) => {
  return service({
    url: '/monster2/findMonster2',
    method: 'get',
    params
  })
}

// @Tags Monster2
// @Summary 分页获取monster2列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取monster2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /monster2/getMonster2List [get]
export const getMonster2List = (params) => {
  return service({
    url: '/monster2/getMonster2List',
    method: 'get',
    params
  })
}

// @Tags Monster2
// @Summary 不需要鉴权的monster2接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.Monster2Search true "分页获取monster2列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /monster2/getMonster2Public [get]
export const getMonster2Public = () => {
  return service({
    url: '/monster2/getMonster2Public',
    method: 'get',
  })
}
