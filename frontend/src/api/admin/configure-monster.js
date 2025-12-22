import request from '@/utils/request'

// 查询ConfigureMonster列表
export function listConfigureMonster(query) {
  return request({
    url: '/api/v1/configure-monster',
    method: 'get',
    params: query
  })
}

// 查询ConfigureMonster详细
export function getConfigureMonster(id) {
  return request({
    url: '/api/v1/configure-monster/' + id,
    method: 'get'
  })
}

// 新增ConfigureMonster
export function addConfigureMonster(data) {
  return request({
    url: '/api/v1/configure-monster',
    method: 'post',
    data: data
  })
}

// 修改ConfigureMonster
export function updateConfigureMonster(data) {
  return request({
    url: '/api/v1/configure-monster/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除ConfigureMonster
export function delConfigureMonster(data) {
  return request({
    url: '/api/v1/configure-monster',
    method: 'delete',
    data: data
  })
}

