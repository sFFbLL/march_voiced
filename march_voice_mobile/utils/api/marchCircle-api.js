import Axios from '../request.js';
import baseUrl from '../env.js'

// 是否是三月圈成员，三月圈成员数量，简介
export function searchContact(params) {
    return Axios.get(baseUrl + '/api/users/selectUserLike', {
        params: params,
    })
}


// 三月圈文章列表获取

// 用户发布按钮

// 点赞表情

// 点击微信进行分享，调用微信接口