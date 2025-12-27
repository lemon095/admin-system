import request from '@/utils/request'

// 查询GiftpackRedeems列表
export function listGiftpackRedeems(query) {
    return request({
        url: '/api/v1/giftpack-redeems',
        method: 'get',
        params: query
    })
}

// 查询GiftpackRedeems详细
export function getGiftpackRedeems (id) {
    return request({
        url: '/api/v1/giftpack-redeems/' + id,
        method: 'get'
    })
}


// 新增GiftpackRedeems
export function addGiftpackRedeems(data) {
    return request({
        url: '/api/v1/giftpack-redeems',
        method: 'post',
        data: data
    })
}

// 修改GiftpackRedeems
export function updateGiftpackRedeems(data) {
    return request({
        url: '/api/v1/giftpack-redeems/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除GiftpackRedeems
export function delGiftpackRedeems(data) {
    return request({
        url: '/api/v1/giftpack-redeems',
        method: 'delete',
        data: data
    })
}

