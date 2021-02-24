import Axios from '../request.js';
import baseUrl from '../env.js'

// 是否是三月圈成员，三月圈成员数量，简介
export function getMarchCircleInfo(params) {
    return Axios.get(baseUrl + '/api/march/msg', {
        params: params,
    })
}


// 三月圈文章列表获取
export function marchCircleList(params){
	return Axios.get(baseUrl+'/api/march',{
		params:params,
	})
}

// 用户发布按钮
export function publishIdea(params){
	return Axios.post(baseUrl+'/api/march',{
		content:params.content,
		imageList:params.imageList,
		image:params.image
	})
}

// 申请加入三月圈
export function joinMarchCircle(params){
	return Axios.post(baseUrl+'/api/apply/march',{
	params:params,
	})
}

// 点赞表情
export function changeFavour(params){
	return Axios.post(baseUrl+'/api/favour/march',{
		id:params.id,
		type:params.type,
	})
}


// 点击微信进行分享，调用微信接口获取签名信息
// export function getScanCode(params){
// 	return Axios.get(baseUrl+'/api/apply/march',{
// 		params:params,
// 	})
// }


// 三月圈详情页调用
export function ideaDetail(params){
	return  Axios.get(baseUrl+'/api/march/detail/id',{
		id:params.id,
	})
}

// 三月圈详情页发布评论
export function publishComment(params){
	return  Axios.post(baseUrl+'/api/march/comment',{
		params:params,
	})
}
