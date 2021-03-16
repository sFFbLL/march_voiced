import Axios from '../request.js';


// 查询用户信息
export function information(params) {
	return Axios.get('/api/user/information', {
		id: params,
	})
}

// 修改用户信息
export function modInformation(params) {
	return Axios.put('/api/user/modInformation', {
		nickname: nickname,
		sex: sex,
		signature: signature,
		avatarPath: avatarPath,
	})
}
