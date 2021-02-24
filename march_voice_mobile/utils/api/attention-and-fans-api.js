import Axios from '@/utils/request.js'
import baseUrl from '../env.js'

//获取关注列表
export function getAttentionList(params){
	return Axios.get(baseUrl + '/api/follow', {
		params: params,
	})
}
//修改关注状态
export function changeStatus(params){
	return Axios.get(baseUrl + '/api/follow/verify', {
		params: params,
	})
}