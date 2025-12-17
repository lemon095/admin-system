import request from '@/utils/request'

// 查询ConfigureBullet列表
export function listConfigureBullet(query) {
    return request({
        url: '/api/v1/configure-bullet',
        method: 'get',
        params: query
    })
}

// 查询ConfigureBullet详细
export function getConfigureBullet (id) {
    return request({
        url: '/api/v1/configure-bullet/' + id,
        method: 'get'
    })
}


// 新增ConfigureBullet
export function addConfigureBullet(data) {
    return request({
        url: '/api/v1/configure-bullet',
        method: 'post',
        data: data
    })
}

// 修改ConfigureBullet
export function updateConfigureBullet(data) {
    return request({
        url: '/api/v1/configure-bullet/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除ConfigureBullet
export function delConfigureBullet(data) {
    return request({
        url: '/api/v1/configure-bullet',
        method: 'delete',
        data: data
    })
}

