import Axios from '../request.js';
import baseUrl from '../env.js'

// 查询所有互动消息
export function interactList(params) {
	return Axios.get(baseUrl + '/api/message/user', {
		params: params,
	})
}

// 查询所有关注消息
export function attentionList(params) {
	return Axios.get(baseUrl + '/api/message/follow-notice', {
		params: params,
	})
}

// 查询所有其他消息
export function otherList(params) {
	return Axios.get(baseUrl + '/api/message/other', {
		params: params,
	})
}

// 查询是否有未读消息，每次路由跳转查
export function unreadMessage(params) {
	return Axios.get(baseUrl + '/api/message/user/unread-count', {
		params: params,
	})
}


// 未读消息已读
export function readMessage(params) {
	return Axios.post(baseUrl + '/api/message/user', {
		type: params.type,
	})
}
