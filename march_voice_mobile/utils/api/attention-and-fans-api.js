import Axios from '@/utils/request.js'
import baseUrl from '../env.js'

//获取关注列表
export function getAttentionList(params){
	return Axios.get(baseUrl + '/api/follow', {
		params: params,
	})
}

// 获取粉丝列表
export function getFansnList(params){
	return Axios.get(baseUrl + '/api/follow/fans', {
		params: params,
	})
}

//修改关注状态
export function changeStatus(params){
	return Axios.post(baseUrl + '/api/follow', {
		id:params.id,
	})
}