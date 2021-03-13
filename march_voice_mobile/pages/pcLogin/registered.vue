<template>
	<view class="content-box">
		<view class="content">
			<form @submit="formSubmit" @reset="formReset" class="form-box">
				<p class="lable">用户名</p>
				<uni-easyinput type="text" name="username" placeholder="请输入用户名" />
				<!-- <p class="lable">昵称</p> -->
				<!-- <uni-easyinput type="text" name="nickName" placeholder="请输入昵称" /> -->
				<p class="lable">密码</p>
				<uni-easyinput type="password" name="password" placeholder="请输入密码" />
				<!-- <p class="lable">性别</p>
				<radio-group name="gender" style="padding:10rpx 20rpx 10rpx">
					<label>
						<radio value="男" checked="true"/><text>男</text>
					</label>
					<label style="padding-left: 15rpx;">
						<radio value="女" /><text>女</text>
					</label>
				</radio-group> -->
				<!-- <p class="lable">手机号</p> -->
				<!-- <uni-easyinput type="text" name="phone" placeholder="请输入手机号" /> -->
				
				<uni-easyinput type="text" name="openid" disabled="true" style="display: none;"/>
				<button form-type="submit" class="btn"><span>注册</span></button>
				<!-- <button form-type="submit">Submit</button> -->
				<button form-type="reset" class="btn"><span>重置</span></button>
			</form>
		</view>
		<uni-popup ref="popup" type="center"><uni-popup-message type="error" message="注册失败!" :duration="2000"></uni-popup-message></uni-popup>
	</view>
</template>

<script>
	import { fixUserInfo } from '../../utils/login.js'
	import { setToken,getOpenId } from '../../utils/auth.js'
	export default {
		data() {
			return {}
		},
		methods: {
			formSubmit: function(e) {
				// console.log('form发生了submit事件，携带数据为：' + JSON.stringify(e.detail.value))
				var formdata = e.detail.value
				
				// var myreg=/^[1][3,4,5,7,8][0-9]{9}$/;
				// if(!myreg.test(formdata.phone) | formdata.username === undefined | formdata.nickName === undefined | formdata.password === undefined | formdata.gender === ""){
				// 	this.$refs.popup.open()
				// 	return
				// }
				var openid = getOpenId();
				formdata.openid = "openid";
				if(formdata.username === undefined | formdata.password === undefined | formdata.openid === undefined){
					this.$refs.popup.open()
					return
				}
				// if(myreg.test(formdata.phone) && formdata.username && formdata.nickName && formdata.password && formdata.gender){
				if(formdata.username && formdata.password && formdata.openid){
					fixUserInfo(formdata).then(res => {
						if(res.code === 200){
							setToken(res.data.token)
						}
					})
				}
				
				
				// uni.showModal({
				// 	content: '表单数据内容：' + JSON.stringify(formdata),
				// 	showCancel: false
				// });
			},
			formReset: function(e) {
				console.log('清空数据')
			}
		}
	}
</script>

<style scoped>
	page {
		height: 90%;
	}

	page .content-box {
		background-color: rgba(166, 165, 143, 0.8);
		height: 100%;
		display: flex;
		justify-content: center;
		/* 水平居中 */
		align-items: center;
	}

	page .content {
		/* display: flex; */
		background-color: #fff;

	}

	/deep/.uni-forms-item__content {
		display: flex;
		align-items: center;
	}

	.row {
		display: flex;
		align-items: center;
		position: relative;
		padding: 0 30upx;
		height: 110upx;
		background: #fff;

		.tit {
			flex-shrink: 0;
			width: 120upx;
			font-size: 30upx;

		}

		.input {
			flex: 1;
			font-size: 30upx;

		}

		.icon-shouhuodizhi {
			font-size: 36upx;

		}
	}

	.add-btn {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 690upx;
		height: 80upx;
		margin: 60upx auto;
		background-color: rgb(28, 42, 134);
		color: #fff;
		border-radius: 10upx;
		box-shadow: 1px 2px 5px rgba(28, 42, 134, 0.4);
	}
	
	.content{
		width: 600rpx;
		padding: 20rpx 0;
		border-radius: 10rpx;
	}
	.lable{
		padding: 10rpx 20rpx;
	}
	.uni-easyinput{
		padding: 0 20rpx;
	}
	.btn{
		margin-top: 20rpx;
		width: 200rpx;
		height: 60rpx;
		display: flex;
		align-items: center;
		text-align: center;
	}
	.btn span {
		width: inherit;
	}
</style>
