import Axios from './request.js';

// 补全用户信息
export function creatNewUser(params) {
	return Axios.post('/api/user/creatNewUser', {
		username: params.username,
		password: params.password,
		openid: params.openid
	})
}

// 查询用户是否存在
export function isExist(params) {
	return Axios.get('/api/user/isExist', {
		params: params,
	})
}

// 用户登录
export function login(params) {
	return Axios.post('/api/user/login', {
		code: params.code,
	})
}

// pc生成二维码
export function pcLogin() {
	return Axios.get('/api/user/getTicket')
}

// pc扫码登录判断
export function pcIsLogin(params) {
	return Axios.get('/api/user/check', {
		params
	})
}

export function fixUserInfo(params) {
	return Axios.post('/api/user/creatNewUser', {
		params
	})
}
