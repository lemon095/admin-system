import service from '@/utils/request'
// @Tags Weapon1
// @Summary 创建武器1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon1 true "创建武器1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /weapon1/createWeapon1 [post]
export const createWeapon1 = (data) => {
  return service({
    url: '/weapon1/createWeapon1',
    method: 'post',
    data
  })
}

// @Tags Weapon1
// @Summary 删除武器1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon1 true "删除武器1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon1/deleteWeapon1 [delete]
export const deleteWeapon1 = (params) => {
  return service({
    url: '/weapon1/deleteWeapon1',
    method: 'delete',
    params
  })
}

// @Tags Weapon1
// @Summary 批量删除武器1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除武器1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon1/deleteWeapon1 [delete]
export const deleteWeapon1ByIds = (params) => {
  return service({
    url: '/weapon1/deleteWeapon1ByIds',
    method: 'delete',
    params
  })
}

// @Tags Weapon1
// @Summary 更新武器1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon1 true "更新武器1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weapon1/updateWeapon1 [put]
export const updateWeapon1 = (data) => {
  return service({
    url: '/weapon1/updateWeapon1',
    method: 'put',
    data
  })
}

// @Tags Weapon1
// @Summary 用id查询武器1
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Weapon1 true "用id查询武器1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weapon1/findWeapon1 [get]
export const findWeapon1 = (params) => {
  return service({
    url: '/weapon1/findWeapon1',
    method: 'get',
    params
  })
}

// @Tags Weapon1
// @Summary 分页获取武器1列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取武器1列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weapon1/getWeapon1List [get]
export const getWeapon1List = (params) => {
  return service({
    url: '/weapon1/getWeapon1List',
    method: 'get',
    params
  })
}

// @Tags Weapon1
// @Summary 不需要鉴权的武器1接口
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.Weapon1Search true "分页获取武器1列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /weapon1/getWeapon1Public [get]
export const getWeapon1Public = () => {
  return service({
    url: '/weapon1/getWeapon1Public',
    method: 'get',
  })
}
