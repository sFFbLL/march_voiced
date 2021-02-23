import Axios from '../request.js';
import baseUrl from '../env.js'


// 查询用户信息
export function information(params) {
	return Axios.get(baseUrl + '/api/user/information', {
		params: params,
	})
}

// 修改用户信息
export function modInformation(params) {
	return Axios.put(baseUrl + '/api/user/modInformation', {
		nickname: params.nickname,
		sex: params.sex,
		signature: params.signature,
		avatarPath: params.avatarPath,
	})
}