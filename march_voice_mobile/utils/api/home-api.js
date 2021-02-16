import request from '@/utils/request.js'
/**
 * 获取推荐文章接口
 * @param {Object} params {页码 页大小}
 */
export function getRecommend (params) {
	return request({
		url: '/api/article-recom',
		method: 'get',
		params
	})
}

/**
 * 获取关注动态接口
 * @param {Object} params {页码 页大小}
 */
export function getFollow (params) {
	return request({
		url: '/api/message/follow',
		method: 'get',
		params
	})
}

/**
 * 搜索接口
 * @param {Object} params {搜索关键字 页码 页大小}
 */
export function search (params) {
	return request({
		url: '/api/base/search',
		method: 'get',
		params
	})
}
