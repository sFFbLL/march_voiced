import Axios from '@/utils/request.js'
/**
 * 获取文章详情的接口
 * @param {Object} params {文章id}
 */
export function getArtileDetails(params) {
	return Axios.get('/api/article/detail', {
		params: params
	})
}

/**
 * 获取文章评论的接口
 * @param {Object} params {id:文章id,current:页码,size:一页大小,childSize:子评论条数}
 */
export function getArticleCommentList(params) {
	return Axios.get('/api/comment/article', {
		params: params
	})
}

/**
 * 新增文章评论的接口
 * @param {Object} params {id:文章id,current:页码,size:一页大小,childSize:子评论条数}
 */
export function addArticleComment(params) {
	return Axios.post('/api/comment/article', {
		params: params
	})
}

/**
 * 获取子评论的接口
 * @param {Object} params {id:文章id,current:页码,size:一页大小,childSize:子评论条数}
 */
export function getChildCommentList(params) {
	return Axios.get('/api/comment/children/article', {
		params: params
	})
}

/**
 * 文章点赞的接口
 * @param {Object} params {id:文章id,current:页码,size:一页大小,childSize:子评论条数}
 */
export function favour(params) {
	return Axios.post('/api/favour/article', {
		params: params
	})
}

/**
 * 文章收藏的接口
 * @param {Object} params {id:文章id,current:页码,size:一页大小,childSize:子评论条数}
 */
export function collect(params) {
	return Axios.post('/api/collect/article', {
		params: params
	})
}

/**s
 * 文章收藏的接口
 * @param {Object} params {id:文章id,current:页码,size:一页大小,childSize:子评论条数}
 */
export function reprint(params) {
	return Axios.post('/api/article/reprint', {
		params: params
	})
}
