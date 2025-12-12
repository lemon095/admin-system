import service from '@/utils/request'
// @Tags Bullet
// @Summary 创建子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Bullet true "创建子弹配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /bullet/createBullet [post]
export const createBullet = (data) => {
  return service({
    url: '/bullet/createBullet',
    method: 'post',
    data
  })
}

// @Tags Bullet
// @Summary 删除子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Bullet true "删除子弹配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bullet/deleteBullet [delete]
export const deleteBullet = (params) => {
  return service({
    url: '/bullet/deleteBullet',
    method: 'delete',
    params
  })
}

// @Tags Bullet
// @Summary 批量删除子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除子弹配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bullet/deleteBullet [delete]
export const deleteBulletByIds = (params) => {
  return service({
    url: '/bullet/deleteBulletByIds',
    method: 'delete',
    params
  })
}

// @Tags Bullet
// @Summary 更新子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Bullet true "更新子弹配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bullet/updateBullet [put]
export const updateBullet = (data) => {
  return service({
    url: '/bullet/updateBullet',
    method: 'put',
    data
  })
}

// @Tags Bullet
// @Summary 用id查询子弹配置
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Bullet true "用id查询子弹配置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bullet/findBullet [get]
export const findBullet = (params) => {
  return service({
    url: '/bullet/findBullet',
    method: 'get',
    params
  })
}

// @Tags Bullet
// @Summary 分页获取子弹配置列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取子弹配置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bullet/getBulletList [get]
export const getBulletList = (params) => {
  return service({
    url: '/bullet/getBulletList',
    method: 'get',
    params
  })
}

// @Tags Bullet
// @Summary 不需要鉴权的子弹配置接口
// @Accept application/json
// @Produce application/json
// @Param data query configureReq.BulletSearch true "分页获取子弹配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bullet/getBulletPublic [get]
export const getBulletPublic = () => {
  return service({
    url: '/bullet/getBulletPublic',
    method: 'get',
  })
}
