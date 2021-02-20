import request from '@/utils/request.js'

/**
 * 获取文章标签接口
 */
export function getTags () {
	return request({
		url: '/api/article-recom',
		method: 'get',
	})
}

/**
 * 添加文章接口
 * @param {Object} params { 标题 文章内容 文章首图url 分类（1/文章 2/笔记） 文章标签 类型（1/转载 0原创） 文章状态 }
 */
export function publishArticle (params) {
	return request({
		url: '/api/article',
		method: 'post',
		params
	})
}