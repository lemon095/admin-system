import request from '@/utils/request'

// 查询Workshop列表
export function listWorkshop(query) {
    return request({
        url: '/api/v1/workshop',
        method: 'get',
        params: query
    })
}

// 查询Workshop详细
export function getWorkshop (id) {
    return request({
        url: '/api/v1/workshop/' + id,
        method: 'get'
    })
}


// 新增Workshop
export function addWorkshop(data) {
    return request({
        url: '/api/v1/workshop',
        method: 'post',
        data: data
    })
}

// 修改Workshop
export function updateWorkshop(data) {
    return request({
        url: '/api/v1/workshop/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除Workshop
export function delWorkshop(data) {
    return request({
        url: '/api/v1/workshop',
        method: 'delete',
        data: data
    })
}
