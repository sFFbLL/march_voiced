<template>
	<view>
		<view class="wrap">
			<view class="inputcontent">
				<view class="name">
					<uni-icons type="person" size="30" color="white"></uni-icons>
					<input placeholder="请设置用户名" @input="nameInput" placeholder-style="color:white" />
				</view>
				<view class="password">
					<uni-icons type="locked" size="30" color="white"></uni-icons>
					<input placeholder="请设置密码" @input="passwordInput" password="true" placeholder-style="color:white" />
				</view>
				<button :loading="btnloading" class="submitbtn" @click="toSubmit">提交</button>
			</view>
		</view>
	</view>
</template>

<script>
	import {
		isExist,
		creatNewUser,
		login
	} from "../../utils/login.js"
	import {
		returnWxcode,
		getWxCode
	} from "../../utils/wxcode.js"
	import {
		getToken,
		setToken,
		setOpenId,
		getOpenId,
		setUserName,
		setAvatarPath,
	} from "../../utils/auth.js"
	export default {
		data() {
			return {
				btnloading: false,
				nameValue: "",
				passwordValue: "",
			}
		},
		methods: {
			// 获取用户名和密码
			nameInput(event) {
				this.nameValue = event.target.value;
			},
			passwordInput(event) {
				this.passwordValue = event.target.value;
			},
			// 提交
			toSubmit() {
				if (this.nameValue == "" || this.passwordValue == "") {
					uni.showToast({
						title: '请完善信息',
						icon: "none",
						position: "top",
						duration: 2000
					});
				} else {
					this.btnloading = true;
					let isExistOr = true;
					let openid = getOpenId()
					let _this = this;
					let params = {
						username: _this.nameValue,
						password: _this.passwordValue,
						openid: openid,
					}
					// 调用补全信息的接口获得用户的token
					creatNewUser(params).then(res => {
						setToken(res.data.token);
						information().then(res => {
							setUserName(res.data.nickname);
							setAvatarPath(res.data.avatarPath);
							
						})
						// 跳转到首页
						uni.switchTab({
							url: '../home/index'
						});
					})


				}
			}


		}
	}
</script>

<style scoped>
	/* 提交按钮 */
	.submitbtn {
		width: 300rpx;
		margin-top: 250rpx;
		border-radius: 40rpx;
		color: #FFFFFF;
		background-image: linear-gradient(140deg, #7ce1e6 0%, #5b97e5 100%);
	}

	.inputcontent {
		position: relative;
		top: 400rpx;
		margin-left: 80rpx;
		width: 600rpx;
		color: #FFFFFF;
	}

	.inputcontent input {
		margin-left: 30rpx;
		height: 3.5em;
	}

	/* 输入密码 */
	.password {
		border-bottom: rgba(255, 255, 255, 0.5) 3rpx solid;
		display: flex;
		line-height: 2;
		height: 100rpx;
	}

	/* 输入用户名 */
	.name {
		height: 100rpx;
		border-bottom: rgba(255, 255, 255, 0.5) 3rpx solid;
		display: flex;
		line-height: 2;
		margin-bottom: 40rpx;
	}

	/* 整个页面 */
	.wrap {
		height: 100vh;
		width: 100vh;
		background-image: url(../../static/img/loginbg1.jpeg);
		/* background-repeat: no-repeat;
		background-position: 100% 100%; */
	}
</style>
