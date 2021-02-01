if (process.env.NODE_ENV === 'development') {
    // process.env.BASE_URL = 'http://120.53.225.156:8089/';
    console.log('开发环境', process.env.BASE_URL)
} else if (process.env.NODE_ENV === 'prodution') {
    console.log('生产环境');
} else if (process.env.NODE_ENV === 'test') {
    console.log('测试环境');
}

module.exports = {
    publicPath: '/front/',
    outputDir: 'dist',
    devServer: {
        port: 8080,
        proxy: {
            '/api': {
                target: process.env.BASE_URL,
                ws: false,
                secure: false,
                changeOrigin: true
            }
        }
    }
}