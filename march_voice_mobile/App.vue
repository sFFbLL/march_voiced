<script>
	import {returnWxcode,getWxCode} from "./utils/wxcode.js"
	import {getToken,setToken} from "./utils/auth.js"
	import {login,creatNewUser} from "./utils/login.js"
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		},
		beforeCreate() {
		// // 判断是否有token
		// if(!getToken()){
		// 	//没有token，没登陆过，获取wxcode
		// 	let code = returnWxcode();
		// 	console.log(code)
		// 	login(code).then(res=>{
		// 		if(res.data.status==1){
		// 			// 跳转注册页面
		// 			uni.navigateTo({
		// 				url:"../login/login"
		// 			})
					
		// 		}else{
		// 			// 登陆成功
		// 			setToken(res.data.token);
		// 			setOpenId(res.data.openid)
		// 		}
		// 	})
			
		// }
			
		},
		
		
	}
</script>

<style lang="scss">
	/* 注意要写在第一行，同时给style标签加入lang="scss"属性 */
	@import "uview-ui/index.scss";
</style>
