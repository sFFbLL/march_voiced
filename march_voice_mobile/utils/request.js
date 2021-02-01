import Axios from 'axios';

Axios.defaults.timeout = 8000;
Axios.interceptors.request.use(
    (config) => {
        config.headers['Authorization'] = 'Bearer eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiJiNGE1NmEzMmJiYTg0NGQwOGE4NzRiN2M0NWZlMTFiZSIsImF1dGgiOiJhZG1pbiIsInN1YiI6IjEifQ.gYKXlkel36RWtAUwyUnpAj7d6TeSwoj0cofpGX1sANhqgVmglEB0fa0qW8auNVZ239iERVh1xw3vPXAmrTZXtA';
        config.headers['Content-type'] = 'application/json';
        return config;
    }, err => {
        return err;
    }
)

Axios.interceptors.response.use(
    (response) => {
        console.log('来到了response拦截success中');
        return response.data;
    },
    (err) => {
        console.log('来到了response拦截failure中');
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
        return Promise.reject(err);
    }
)

Axios.defaults.adapter = function (config) {  
    return new Promise((resolve, reject) => {  
        console.log(config)  
        var settle = require('axios/lib/core/settle');  
        var buildURL = require('axios/lib/helpers/buildURL');  
        uni.request({  
            method: config.method.toUpperCase(),  
            url: buildURL(config.url, config.params, config.paramsSerializer),  
            header: config.headers,  
            data: config.data,  
            dataType: config.dataType,  
            responseType: config.responseType,  
            sslVerify: config.sslVerify,  
            complete:function complete(response){  
                response = {  
                  data: response.data,  
                  status: response.statusCode,  
                  errMsg: response.errMsg,  
                  header: response.header,  
                  config: config  
                };  

            settle(resolve, reject, response);  
            }  
        })  
    })  
}  

export default Axios;
