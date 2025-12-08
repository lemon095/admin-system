import service from '@/utils/request'
// @Tags Weapon
// @Summary 创建武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon true "创建武器配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /weapon/createWeapon [post]
export const createWeapon = (data) => {
  return service({
    url: '/weapon/createWeapon',
    method: 'post',
    data
  })
}

// @Tags Weapon
// @Summary 删除武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon true "删除武器配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon/deleteWeapon [delete]
export const deleteWeapon = (params) => {
  return service({
    url: '/weapon/deleteWeapon',
    method: 'delete',
    params
  })
}

// @Tags Weapon
// @Summary 批量删除武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除武器配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon/deleteWeapon [delete]
export const deleteWeaponByIds = (params) => {
  return service({
    url: '/weapon/deleteWeaponByIds',
    method: 'delete',
    params
  })
}

// @Tags Weapon
// @Summary 更新武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon true "更新武器配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weapon/updateWeapon [put]
export const updateWeapon = (data) => {
  return service({
    url: '/weapon/updateWeapon',
    method: 'put',
    data
  })
}

// @Tags Weapon
// @Summary 用id查询武器配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Weapon true "用id查询武器配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weapon/findWeapon [get]
export const findWeapon = (params) => {
  return service({
    url: '/weapon/findWeapon',
    method: 'get',
    params
  })
}

// @Tags Weapon
// @Summary 分页获取武器配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取武器配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weapon/getWeaponList [get]
export const getWeaponList = (params) => {
  return service({
    url: '/weapon/getWeaponList',
    method: 'get',
    params
  })
}

// @Tags Weapon
// @Summary 不需要鉴权的武器配置接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.WeaponSearch true "分页获取武器配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /weapon/getWeaponPublic [get]
export const getWeaponPublic = () => {
  return service({
    url: '/weapon/getWeaponPublic',
    method: 'get',
  })
}
