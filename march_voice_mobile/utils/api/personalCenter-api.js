import Axios from '@/utils/request.js'
/**
 * 获取个人主页用户信息接口
 * @param {Object} params {id}
 */
export function getUserInfo(params) {
	return Axios.get("/api/user/ingo/?id="+params.id)
}

/**
 * 获取个人主页文章列表接口
 * @param {Object} params {id,current,size,kind}
 */
export function getUserArticleList(params) {
	return Axios.get('/api/article/user', {
		id: params.id,
		current: params.current,
		size: params.size,
		kind: params.kind
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
