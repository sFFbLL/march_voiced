<script>
	import {
		returnWxcode,
		getWxCode
	} from "./utils/wxcode.js"
	import {
		getToken,
		setToken
	} from "./utils/auth.js"
	import {
		login,
		creatNewUser
	} from "./utils/login.js"
	export default {
		onLaunch: function () {

			console.log('App Launch')

		},
		onShow: function () {
			console.log('App Show')
		},
		onHide: function () {
			console.log('App Hide')
		},
		 beforeCreate() {
			console.log("beforeCreate");
			if (!getToken()) {
				console.log("没有token");
				//没有token，没登陆过，获取wxcode
				let code = returnWxcode();
				console.log("app.vue")
				let params = {
					code: code,
					status: 1
				}
				// 判断该用户是否注册
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
						setOpenId(res.data.openid)
					}
				}).catch(err => {
					console.log(err, "err login")
				})

			} else {
			
				console.log("tologin")
			}

		},


	}
</script>


<style lang="scss">
	/* 注意要写在第一行，同时给style标签加入lang="scss"属性 */
	@import "uview-ui/index.scss";
	@import "static/css/iconfont.css";
</style>
