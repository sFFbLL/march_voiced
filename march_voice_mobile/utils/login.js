import Axios from './request.js';
import baseUrl from './env.js'

// 补全用户信息
export function creatNewUser(params) {
	return Axios.post(baseUrl + '/api/user/creatNewUser', {
		username: params.username,
		password: params.password
	})
}

// 查询用户是否存在
export function isExist(params) {
	return Axios.get(baseUrl + '/api/user/isExist', {
		username: params,
	})
}

// 用户登录
export function login(params) {
	console.log(baseUrl)
	return Axios.post(baseUrl + '/api/user/login', {
		code: params.code,
	})
}

// pc生成二维码
export function pcLogin(){
	return Axios.get(baseUrl + '/api/user/getTicket')
}

// pc扫码登录判断
export function pcIsLogin(params){
	return Axios.get(baseUrl + '/api/user/check',{
		params
	})
}

export function fixUserInfo(params){
	return Axios.post(baseUrl + '/api/user/creatNewUser',{
		params
	})
}