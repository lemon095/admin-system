import service from '@/utils/request'
// @Tags Weapon2
// @Summary 创建表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon2 true "创建表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /weapon2/createWeapon2 [post]
export const createWeapon2 = (data) => {
  return service({
    url: '/weapon2/createWeapon2',
    method: 'post',
    data
  })
}

// @Tags Weapon2
// @Summary 删除表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon2 true "删除表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon2/deleteWeapon2 [delete]
export const deleteWeapon2 = (params) => {
  return service({
    url: '/weapon2/deleteWeapon2',
    method: 'delete',
    params
  })
}

// @Tags Weapon2
// @Summary 批量删除表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon2/deleteWeapon2 [delete]
export const deleteWeapon2ByIds = (params) => {
  return service({
    url: '/weapon2/deleteWeapon2ByIds',
    method: 'delete',
    params
  })
}

// @Tags Weapon2
// @Summary 更新表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon2 true "更新表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weapon2/updateWeapon2 [put]
export const updateWeapon2 = (data) => {
  return service({
    url: '/weapon2/updateWeapon2',
    method: 'put',
    data
  })
}

// @Tags Weapon2
// @Summary 用id查询表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Weapon2 true "用id查询表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weapon2/findWeapon2 [get]
export const findWeapon2 = (params) => {
  return service({
    url: '/weapon2/findWeapon2',
    method: 'get',
    params
  })
}

// @Tags Weapon2
// @Summary 分页获取表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weapon2/getWeapon2List [get]
export const getWeapon2List = (params) => {
  return service({
    url: '/weapon2/getWeapon2List',
    method: 'get',
    params
  })
}

// @Tags Weapon2
// @Summary 不需要鉴权的表接口
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.Weapon2Search true "分页获取表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /weapon2/getWeapon2Public [get]
export const getWeapon2Public = () => {
  return service({
    url: '/weapon2/getWeapon2Public',
    method: 'get',
  })
}
