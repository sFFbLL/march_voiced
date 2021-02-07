import getScanCode from "./api/marchCircle-api.js"

// jwx.js
//#ifdef H5
const jweixin = require('jweixin-module')
//#endif
export function configWeiXin() {
	//查询后端返回的公众号参数信息
	getScanCode().then(res => {
		console.log("进入到请求qianming的回调函数中==========入参是：", res);
		// 通过config接口注入权限验证配置

		wx.config({
			debug: false, // 开启调试模式
			appId: res.data.appId, // 必填，公众号的唯一标识
			timestamp: res.data.timeStamp, // 必填，生成签名的时间戳
			nonceStr: res.data.nonceStr, // 必填，生成签名的随机串
			signature: res.data.signNature, // 必填，签名，见附录1
			jsApiList: [
				'updateAppMessageShareData', // 获取“分享给朋友”按钮点击状态及自定义分享内容接口
			] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
		});



		// 通过ready接口处理成功验证
		wx.ready(function() { //需在用户可能点击分享按钮前就先调用

			wx.updateAppMessageShareData({
				title:'三月之声-三月圈',
				desc:'你的好友在三月圈发布了新的想法，戳我围观哦',
				link:window.location.href,
				imgUrl:'',
				success:function(){}
			});
		});
		// config信息验证失败会执行error函数，
		// 如签名过期导致验证失败，具体错误信
		// 息可以打开config的debug模式查看，也可以在返回的res参数中查看，对于SPA可以在这里更新签名。
		wx.error(function(res) {
			console.log("error---------" + res);
		})

		wx.checkJsApi({
			jsApiList: ['updateAppMessageShareData'], // 需要检测的JS接口列表，所有JS接口列表见附录2,
			success: function(res) {
				console.log("checkJsApi-------" + res);
				console.log("checkJsApi-------" + res.checkResult.updateAppMessageShareData);

				// 以键值对的形式返回，可用的api值true，不可用为false
				// 如：{"checkResult":{"chooseImage":true},"errMsg":"checkJsApi:ok"}
			}
		});
	});
}
