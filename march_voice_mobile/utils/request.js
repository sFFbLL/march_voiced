import Axios from 'axios';
import {
	getToken,
	setToken
} from "./auth.js"
import {
	returnWxcode,
	parseCode
} from './wxcode.js'
import {
	login,
	creatNewUser
} from "./login.js"
import QS from 'qs'; // 引入qs模块，用来序列化post类型的数据，某些请求会用得到

Axios.defaults.baseURL = "http://linbolun.cn";
Axios.defaults.timeout = 8000;

// 请求拦截器
Axios.interceptors.request.use(
	// 在发送请求前要做的事儿
	(config) => {
		console.log("来到了全局request中");
		// setToken("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNjE1OTk0NjA5LCJpc3MiOiJteS1wcm9qZWN0In0.R-IviSlw1nw2U6ijLnHyJfrLi7kSBNyaIf6sZKueIL0")
		if (getToken()) {
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
		console.log(response.data, "success");
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



export default Axios;
