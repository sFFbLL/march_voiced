import Axios from './request.js'
import baseUrl from './env.js'

// 图片上传接口
export function unloadImage(params) {
	return Axios.post(baseUrl + '/api/base/article-recom', {
		params: params,
	})
}
