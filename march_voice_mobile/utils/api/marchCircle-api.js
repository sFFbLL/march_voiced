import Axios from '../request.js';

// 是否是三月圈成员，三月圈成员数量，简介
export function getMarchCircleInfo(params) {
	return Axios.get('/api/march/msg', {
		params: params,
	})
}
// 三月圈文章列表获取
export function marchCircleList(params) {
	return Axios.get('/api/march', {
		params: params,
	})
}

// 用户发布按钮
export function publishIdea(params) {
	return Axios.post('/api/march', params)
}

// 申请加入三月圈
export function joinMarchCircle(params) {
	return Axios.post('/api/apply/march',
		params)
}

// 点赞表情
export function changeFavour(params) {
	return Axios.post('/api/favour/march', params)
}


// 点击微信进行分享，调用微信接口获取签名信息
// export function getScanCode(params){
// 	return Axios.get(baseUrl+'/api/apply/march',{
// 		params:params,
// 	})
// }


// 三月圈详情页调用
export function ideaDetail(params) {
	return Axios.get(`/api/march/detail/${params}`)
}

// 三月圈详情页发布评论
export function publishComment(params) {
	return Axios.post('/api/comment/marchsoft', params)
}

// 三月圈详情页查询评论
export function ideaCommentList(params) {
	return Axios.get('/api/comment/marchsoft', {
		// id:params.id,
		// current:params.current,
		// size:params.size,
		// childSize:params.childSize
		params: params,
	})
}


/**
 * 获取子评论的接口
 * @param {Object} params {id:id,current:页码,size:一页大小,childSize:子评论条数}
 */
export function ideaChildCommentList(params) {
	return Axios.get('/api/comment/children/marchsoft', {
		params: params
	})
}
