import Axios from 'axios';
import {
	getToken,
	setToken
} from "./auth.js"
import baseUrl from './env.js'
import returnWxcode from './wxcode.js'
import QS from 'qs'; // 引入qs模块，用来序列化post类型的数据，某些请求会用得到

Axios.defaults.baseURL = baseUrl;
Axios.defaults.timeout = 8000;
/****** 创建axios实例 ******/
// const Axios = axios.create({
//     baseURL: baseUrl, // api 的 base_url
//     timeout: 8000  // 请求超时时间
// });

// 请求拦截器
Axios.interceptors.request.use(
	// 在发送请求前要做的事儿
	(config) => {
		console.log("来到了全局request中");
		if (!getToken()) {
			console.log("没有token");
			//没有token，没登陆过，获取wxcode
			let code = returnWxcode();
			
			console.log(returnWxcode())
			console.log("成功拿到code")
			let params = {
				code: code,
				status: 1
			}
			判断该用户是否注册
			login(params).then(res => {
				console.log(res, "注册")
				if (res.data.status == 1) {
					// 跳转注册页面
					console.log("未登录")
					uni.navigateTo({
						url: "../login/login"
					})
				} else {
					// 登陆成功
					console.log(res.data.token)
					setToken(res.data.token);
					config.headers['Authorization'] = getToken() //让每个请求携带自定义token
					setOpenId(res.data.openid)
				}
			}).catch(err => {
				console.log(err, "err login")
			})


    // setToken(
    //   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjE3NzM2NTIxLCJpc3MiOiJteS1wcm9qZWN0In0.l3-dvGSa41PUIybA_Dmq50ZtePo6qgwe5YVBRTs8K8Q"
    // )
    config.headers['Authorization'] = getToken() //让每个请求携带自定义token
    config.headers['Content-type'] = "application/json;charset=utf-8";
    config.data = JSON.stringify(config.data);
    console.log(config);
    return config;
		}else{
			config.headers['Authorization'] = getToken() //让每个请求携带自定义token
		}
		config.headers['Content-type'] = "application/json;charset=utf-8";
		config.data = JSON.stringify(config.data);
		console.log(config);
		return config;


	}, err => {
		// 在请求错误时要做的事儿
		return err;
	}
)
// 响应拦截器
Axios.interceptors.response.use(
	(response) => {
		console.log(response.data.message, "success");
		// uni.showToast({
		// 	title: response.data.message,
		// 	duration: 2000,
		// 	icon: "none"
		// });
		return response.data;
	},
	(err) => {
		console.log("error33333");
		if (err && err.response) {
			switch (err.response.status) {
				case 400:
					err.message = '参数错误（400）';
					break;
				case 401:
					err.message = '未授权访问（401）';
					break;
				case 403:
					err.message = '权限错误（403）';
					break;
				case 404:
					err.message = '访问资源错误（404）';
				case 500:
					err.message = '服务器错误（500）';

			}
		} else if (err.response.status < 200 || err.response.status > 300) {
			err.message = '请求失败';
		}
		uni.showToast({
			title: err.message,
			duration: 2000,
			icon: "none"
		});
		return Promise.reject(err);
	}
)

// Axios.defaults.adapter = function (config) {  
//     return new Promise((resolve, reject) => {  
//         console.log(config)  
//         var settle = require('axios/lib/core/settle');  
//         var buildURL = require('axios/lib/helpers/buildURL');  
//         uni.request({  
//             method: config.method.toUpperCase(),  
//             url: buildURL(config.url, config.params, config.paramsSerializer),  
//             header: config.headers,  
//             data: config.data,  
//             dataType: config.dataType,  
//             responseType: config.responseType,  
//             sslVerify: config.sslVerify,  
//             complete:function complete(response){  
//                 response = {  
//                   data: response.data,  
//                   status: response.statusCode,  
//                   errMsg: response.errMsg,  
//                   header: response.header,  
//                   config: config  
//                 };  

//             settle(resolve, reject, response);  
//             }  
//         })  
//     })  
// }  

export default Axios;
