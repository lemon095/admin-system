import request from '@/utils/request'

// 查询ConfigureArmor列表
export function listConfigureArmor(query) {
  return request({
    url: '/api/v1/configure-armor',
    method: 'get',
    params: query
  })
}

// 查询ConfigureArmor详细
export function getConfigureArmor(id) {
  return request({
    url: '/api/v1/configure-armor/' + id,
    method: 'get'
  })
}

// 新增ConfigureArmor
export function addConfigureArmor(data) {
  return request({
    url: '/api/v1/configure-armor',
    method: 'post',
    data: data
  })
}

// 修改ConfigureArmor
export function updateConfigureArmor(data) {
  return request({
    url: '/api/v1/configure-armor/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除ConfigureArmor
export function delConfigureArmor(data) {
  return request({
    url: '/api/v1/configure-armor',
    method: 'delete',
    data: data
  })
}

