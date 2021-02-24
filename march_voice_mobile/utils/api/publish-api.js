import Axios from '../request.js';
import baseUrl from '../env.js'
/**
 * 获取文章标签接口
 */
export function getTags () {
	return Axios.get(baseUrl + '/api/article-recom')
}

/**
 * 添加文章接口
 * @param {Object} params { 标题 文章内容 文章首图url 分类（1/文章 2/笔记） 文章标签 类型（1/转载 0原创） 文章状态 }
 */
export function publishArticle (params) {
	return Axios.post(baseUrl + '/api/article', {
	    params: params,
	})
}

/**
 * 修改文章接口 
 * @param {Object} params { 标题 内容 文章状态 标签}
 */
export function upDateArticle (params){
	return Axios.put(baseUrl + '/api/article',{
		params
	})
}

/**
 * 获取用户文章类表
 * @param {Object} params { 当前页 页大小  } 
 */
export function getArticleList (params){
	return Axios.get(baseUrl + '/api/article/user', {
	    params: params,
	})
}
