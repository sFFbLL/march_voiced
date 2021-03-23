import Axios from '../request.js';

// 查询所有互动消息
export function interactList(params) {
	return Axios.get('/api/message/user', {
		params: params,
	})
}

// 查询所有关注消息
export function attentionList(params) {
	return Axios.get('/api/message/follow-notice', {
		params: params,
	})
}

// 查询所有其他消息
export function otherList(params) {
	return Axios.get('/api/message/system', {
		params: params,
	})
}

// 查询是否有未读消息
export function unreadMessage(params) {
	return Axios.get('/api/message/user/unread-count', {
		type: params,
	})
}


// 未读消息已读
export function readMessage(params) {
	return Axios.put('/api/message/user', {
		type: params,
	})
}
