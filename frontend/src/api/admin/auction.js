import request from '@/utils/request'

// 查询Auction列表
export function listAuction(query) {
    return request({
        url: '/api/v1/auction',
        method: 'get',
        params: query
    })
}

// 查询Auction详细
export function getAuction (id) {
    return request({
        url: '/api/v1/auction/' + id,
        method: 'get'
    })
}


// 新增Auction
export function addAuction(data) {
    return request({
        url: '/api/v1/auction',
        method: 'post',
        data: data
    })
}

// 修改Auction
export function updateAuction(data) {
    return request({
        url: '/api/v1/auction/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除Auction
export function delAuction(data) {
    return request({
        url: '/api/v1/auction',
        method: 'delete',
        data: data
    })
}

