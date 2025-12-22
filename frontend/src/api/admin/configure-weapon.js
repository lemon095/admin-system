import request from '@/utils/request'

// 查询ConfigureWeapon列表
export function listConfigureWeapon(query) {
  return request({
    url: '/api/v1/configure-weapon',
    method: 'get',
    params: query
  })
}

// 查询ConfigureWeapon详细
export function getConfigureWeapon(id) {
  return request({
    url: '/api/v1/configure-weapon/' + id,
    method: 'get'
  })
}

// 新增ConfigureWeapon
export function addConfigureWeapon(data) {
  return request({
    url: '/api/v1/configure-weapon',
    method: 'post',
    data: data
  })
}

// 修改ConfigureWeapon
export function updateConfigureWeapon(data) {
  return request({
    url: '/api/v1/configure-weapon/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除ConfigureWeapon
export function delConfigureWeapon(data) {
  return request({
    url: '/api/v1/configure-weapon',
    method: 'delete',
    data: data
  })
}

