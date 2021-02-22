import Axios from './request.js'
import baseUrl from './env.js'



// 上传图片
export function unloadImage(params) {
	uni.uploadFile({
		url: baseUrl + '/api/base/article-recom',
		files: params,
		header: {
			'Content-Type': 'multipart/form-data'
		},
		success: (res) => {
			console.log(res.data.full_path)
			return res.data.full_path
		},
		fail: (res) => {
			return res
		}
	})
}
