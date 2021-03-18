import Axios from './request.js'



// 上传图片
// export function unloadImage(params) {
// 	uni.uploadFile({
// 		url: baseUrl + '/api/file/uploadImage',
// 		files: params,
// 		header: {
// 			'Content-Type': 'multipart/form-data'
// 		},
// 		success: (res) => {
// 			console.log(res.data.full_path)
// 			return res.data.full_path
// 		},
// 		fail: (res) => {
// 			return res
// 		}
// 	})
// }

export function unloadImage(params) {
	return Axios.post('/api/file/uploadImage', {
		file: params.file,
	})
}
