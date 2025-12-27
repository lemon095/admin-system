import request from '@/utils/request'

// 查询Item列表
export function listItem(query) {
    return request({
        url: '/api/v1/item',
        method: 'get',
        params: query
    })
}

// 查询Item详细
export function getItem (id) {
    return request({
        url: '/api/v1/item/' + id,
        method: 'get'
    })
}


// 新增Item
export function addItem(data) {
    return request({
        url: '/api/v1/item',
        method: 'post',
        data: data
    })
}

// 修改Item
export function updateItem(data) {
    return request({
        url: '/api/v1/item/'+data.id,
        method: 'put',
        data: data
    })
}

export function updateItemStatus(id) {
    return request({
        url: '/api/v1/item/status/'+id,
        method: 'put'
    })
}

// 删除Item
export function delItem(data) {
    return request({
        url: '/api/v1/item',
        method: 'delete',
        data: data
    })
}

