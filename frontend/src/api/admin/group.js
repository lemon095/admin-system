import request from '@/utils/request'

// 查询Group列表
export function listGroup(query) {
    return request({
        url: '/api/v1/group',
        method: 'get',
        params: query
    })
}

// 查询Group详细
export function getGroup (id) {
    return request({
        url: '/api/v1/group/' + id,
        method: 'get'
    })
}


// 新增Group
export function addGroup(data) {
    return request({
        url: '/api/v1/group',
        method: 'post',
        data: data
    })
}

// 修改Group
export function updateGroup(data) {
    return request({
        url: '/api/v1/group/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除Group
export function delGroup(data) {
    return request({
        url: '/api/v1/group',
        method: 'delete',
        data: data
    })
}

