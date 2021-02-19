import Axios from '@/utils/request.js'
import baseUrl from '../env.js'
/**
 * 获取个人主页用户信息接口
 * @param {Object} params {id}
 */
export function getUserInfo (params) {
	return Axios.get(baseUrl + '/api/user/ingo', {
		params: params,
	})
}

/**
 * 获取个人主页文章列表接口
 * @param {Object} params {id,current,size,kind}
 */
export function getUserArticleList (params) {
	return Axios.get(baseUrl + '/api/article/user', {
		params: params,
	})
}

/**
 * 获取个人主页文章列表接口
 * @param {Object} params {id,current,size}
 */
export function getUserIdeaList (params) {
	return Axios.get(baseUrl + '/api/march/user', {
		params: params,
	})
}
