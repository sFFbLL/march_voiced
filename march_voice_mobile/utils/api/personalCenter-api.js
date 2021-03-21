import Axios from '@/utils/request.js'
/**
 * 获取个人主页用户信息接口
 * @param {Object} params {id}
 */
export function getUserInfo(params) {
	return Axios.get(`/api/user/info/${params.id}`)
}

/**
 * 获取个人主页文章列表接口
 * @param {Object} params {id,current,size,kind}
 */
export function getUserArticleList(params) {
	return Axios.get('/api/article/user', {
		params: params
	})
}

/**
 * 获取个人主页文章列表接口
 * @param {Object} params {id,current,size}
 */
export function getUserIdeaList(params) {
	return Axios.get('/api/march/user', {
		params: params,
	})
}
// 获取收藏列表
export function collectArticle(params) {
	return Axios.get('/api/collect/article', {
		params: params,
	})
}
