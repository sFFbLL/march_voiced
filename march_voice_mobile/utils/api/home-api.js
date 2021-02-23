import Axios from '@/utils/request.js'
import baseUrl from '../env.js'
/**
 * 获取推荐文章接口
 * @param {Object} params {页码 页大小}
 */
export function getRecommend (params) {
	return Axios.get(baseUrl + '/api/article/top', {
		params: params
	})
}

/**
 * 获取关注动态接口
 * @param {Object} params {页码 页大小}
 */
export function getFollow (params) {
	return Axios.get(baseUrl + '/api/base/follow', {
		params: params
	})
}

/**
 * 搜索文章接口
 * @param {Object} params {搜索关键字 页码 页大小}
 */
export function searchArticle (params) {
	return Axios.get(baseUrl + '/api/base/searchArticle', {
		params: params
	})
}

/**
 * 搜索用户接口
 * @param {Object} params {搜索关键字 页码 页大小}
 */
export function searchUser (params) {
	return Axios.get(baseUrl + '/api/base/searchUser', {
		params: params
	})
}
