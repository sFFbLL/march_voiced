export default (function() {
    if (process.env.NODE_ENV === 'development') {
        console.log('开发环境');
        // return 'http://120.53.225.156:8089';
    } else if (process.env.NODE_ENV === 'prodution') {
        return ;
        console.log('生产环境');
    } else if (process.env.NODE_ENV === 'test') {
        return ;
        console.log('测试环境');
    }
})() 
