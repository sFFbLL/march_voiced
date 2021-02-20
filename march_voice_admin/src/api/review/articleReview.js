import request from '@/utils/request'
export function articleAdopt(data) {
    return request({
        url: '/api/apply/article',
        method: 'put',
        data
    })
}
