import Axios from 'axios';
import getToken from "./auth.js"
import baseUrl from './env.js'

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
		// uni.showLoading({
		//     title: '加载中'
		// });
      config.headers['Authorization'] = getToken() //让每个请求携带自定义token
        config.headers['Content-type'] = 'application/json';
        return config;
    }, err => {
		   // 在请求错误时要做的事儿
        return err;
    }
)
// 响应拦截器
Axios.interceptors.response.use(
    (response) => {
       // uni.hideLoading();
		uni.showToast({
		    title: '来到了response拦截success中',
		    duration: 2000,
			icon:success
		});
        return response.data;
    },
    (err) => {
        console.log(err);
        if(err && err.response) {
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
                    err.message = '服务器错误（500）'
            }
			
        } else if (err.response.status < 200 || err.response.status > 300 ) {
            err.message = '请求失败';
        }
		uni.showToast({
		    title:  err.message,
		    duration: 2000,
			icon:none
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
