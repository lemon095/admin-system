import service from '@/utils/request'
// @Tags Armor
// @Summary 创建防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Armor true "创建防具配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /armor/createArmor [post]
export const createArmor = (data) => {
  return service({
    url: '/armor/createArmor',
    method: 'post',
    data
  })
}

// @Tags Armor
// @Summary 删除防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Armor true "删除防具配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /armor/deleteArmor [delete]
export const deleteArmor = (params) => {
  return service({
    url: '/armor/deleteArmor',
    method: 'delete',
    params
  })
}

// @Tags Armor
// @Summary 批量删除防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除防具配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /armor/deleteArmor [delete]
export const deleteArmorByIds = (params) => {
  return service({
    url: '/armor/deleteArmorByIds',
    method: 'delete',
    params
  })
}

// @Tags Armor
// @Summary 更新防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Armor true "更新防具配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /armor/updateArmor [put]
export const updateArmor = (data) => {
  return service({
    url: '/armor/updateArmor',
    method: 'put',
    data
  })
}

// @Tags Armor
// @Summary 用id查询防具配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Armor true "用id查询防具配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /armor/findArmor [get]
export const findArmor = (params) => {
  return service({
    url: '/armor/findArmor',
    method: 'get',
    params
  })
}

// @Tags Armor
// @Summary 分页获取防具配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取防具配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /armor/getArmorList [get]
export const getArmorList = (params) => {
  return service({
    url: '/armor/getArmorList',
    method: 'get',
    params
  })
}

// @Tags Armor
// @Summary 不需要鉴权的防具配置接口
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.ArmorSearch true "分页获取防具配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /armor/getArmorPublic [get]
export const getArmorPublic = () => {
  return service({
    url: '/armor/getArmorPublic',
    method: 'get',
  })
}
