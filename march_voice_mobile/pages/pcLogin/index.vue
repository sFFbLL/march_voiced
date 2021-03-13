<template>
	<!-- <view class="background"> -->
		<image class="code-img" :src="src" mode=""></image>
<!-- 	</view> -->
</template>

<script>
	import { pcLogin,pcIsLogin } from '../../utils/login.js'
	import { setToken,setOpenId } from '../../utils/auth.js'
	export default {
		data() {
			return {
				ticket:null,
				strData:null,
				baseSrc:"https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=",
				src:''
			}
		},
		created() {
			this.login()
			this.interval = setInterval(() => {
				this.isLogin(this.strData)
			}, 2000)
		},
		beforeDestroy() {
			clearInterval(this.interval)
		},
		methods: {
			login(){
				pcLogin().then(res => {
						console.log("res",res)
						this.strData = res.data.strdata
						this.src = this.baseSrc + res.data.Ticket
				}).catch(err=>{
					console.log("errlogin")
				})
			},
			isLogin(strdata){
				let params = {
					strData:strdata
				}
				pcIsLogin(params).then(res => {
						switch(res.code){
							case 0://用户未扫码
								break;
							case 1://用户扫码未注册
								setOpenId(res.data.openid);
								uni.redirectTo({
									url:'../pcPublish/registered',// 注册页面
									// res.openid
								})
								break;
							case 2://用户获取token成功
								setToken(res.token)
								console.log("ASDADADA")
								uni.redirectTo({
									url:'../pcPublish/pcPublish'
								})
								break;
							}
				})
				
			}
		}
	}
</script>

<style scoped>
.uni-page-head[uni-page-head-type=default]~uni-page-wrapper {
	background: black !important;
}
.background{
	background-color: #262626;
	box-sizing: content-box;
}
.code-img{
	position: absolute;
	left: calc(50% - 250rpx);
	top: calc(50% - 250rpx);
}
uni-image{
	width: 500rpx;
	height: 500rpx;
}
</style>
