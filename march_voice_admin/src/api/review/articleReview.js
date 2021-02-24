import request from '@/utils/request'
export function articleAdopt(data) {
  return request({
    url: '/api/apply/article',
    method: 'put',
    data
  })
}
export function marchAdopt(data) {
  return request({
    url: '/api/apply/march',
    method: 'put',
    data
  })
}
export function recommendAdopt(data) {
  return request({
    url: '/api/article/recommend/' + data,
    method: 'put'
  })
}
export function getInfoById(url, id) {
  return request({
    url: url + '/' + id,
    method: 'get'
  })
}
export function sensitiveAdopt(id) {
  return request({
    url: '/api/article/word',
    method: 'get',
    params: { id: id }
  })
}
export function sensitive(url, id) {
  return request({
    url: url,
    method: 'get',
    params: { id: id }
  })
}
