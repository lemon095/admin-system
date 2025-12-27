import request from '@/utils/request'

// 查询ItemType列表
export function listItemType(query) {
    return request({
        url: '/api/v1/item-type',
        method: 'get',
        params: query
    })
}

// 查询ItemType详细
export function getItemType (id) {
    return request({
        url: '/api/v1/item-type/' + id,
        method: 'get'
    })
}


// 新增ItemType
export function addItemType(data) {
    return request({
        url: '/api/v1/item-type',
        method: 'post',
        data: data
    })
}

// 修改ItemType
export function updateItemType(data) {
    return request({
        url: '/api/v1/item-type/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除ItemType
export function delItemType(data) {
    return request({
        url: '/api/v1/item-type',
        method: 'delete',
        data: data
    })
}

