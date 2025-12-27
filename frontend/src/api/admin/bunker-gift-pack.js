import request from '@/utils/request'

// 查询GiftPack列表
export function listGiftPack(query) {
    return request({
        url: '/api/v1/gift-pack',
        method: 'get',
        params: query
    })
}

// 查询GiftPack详细
export function getGiftPack (id) {
    return request({
        url: '/api/v1/gift-pack/' + id,
        method: 'get'
    })
}


// 新增GiftPack
export function addGiftPack(data) {
    return request({
        url: '/api/v1/gift-pack',
        method: 'post',
        data: data
    })
}

// 修改GiftPack
export function updateGiftPack(data) {
    return request({
        url: '/api/v1/gift-pack/'+data.id,
        method: 'put',
        data: data
    })
}

// 修改GiftPack状态
export function updateGiftPackStatus(id) {
    return request({
        url: '/api/v1/gift-pack/status/'+id,
        method: 'put'
    })
}

// 删除GiftPack
export function delGiftPack(data) {
    return request({
        url: '/api/v1/gift-pack',
        method: 'delete',
        data: data
    })
}

