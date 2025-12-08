import service from '@/utils/request'
// @Tags Weapon3
// @Summary 创建Weapon3
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon3 true "创建Weapon3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /weapon33/createWeapon3 [post]
export const createWeapon3 = (data) => {
  return service({
    url: '/weapon33/createWeapon3',
    method: 'post',
    data
  })
}

// @Tags Weapon3
// @Summary 删除Weapon3
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon3 true "删除Weapon3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon33/deleteWeapon3 [delete]
export const deleteWeapon3 = (params) => {
  return service({
    url: '/weapon33/deleteWeapon3',
    method: 'delete',
    params
  })
}

// @Tags Weapon3
// @Summary 批量删除Weapon3
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Weapon3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weapon33/deleteWeapon3 [delete]
export const deleteWeapon3ByIds = (params) => {
  return service({
    url: '/weapon33/deleteWeapon3ByIds',
    method: 'delete',
    params
  })
}

// @Tags Weapon3
// @Summary 更新Weapon3
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Weapon3 true "更新Weapon3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weapon33/updateWeapon3 [put]
export const updateWeapon3 = (data) => {
  return service({
    url: '/weapon33/updateWeapon3',
    method: 'put',
    data
  })
}

// @Tags Weapon3
// @Summary 用id查询Weapon3
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Weapon3 true "用id查询Weapon3"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weapon33/findWeapon3 [get]
export const findWeapon3 = (params) => {
  return service({
    url: '/weapon33/findWeapon3',
    method: 'get',
    params
  })
}

// @Tags Weapon3
// @Summary 分页获取Weapon3列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Weapon3列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weapon33/getWeapon3List [get]
export const getWeapon3List = (params) => {
  return service({
    url: '/weapon33/getWeapon3List',
    method: 'get',
    params
  })
}

// @Tags Weapon3
// @Summary 不需要鉴权的Weapon3接口
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.Weapon3Search true "分页获取Weapon3列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /weapon33/getWeapon3Public [get]
export const getWeapon3Public = () => {
  return service({
    url: '/weapon33/getWeapon3Public',
    method: 'get',
  })
}
