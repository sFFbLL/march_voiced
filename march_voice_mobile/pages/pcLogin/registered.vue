<template>
	<view class="content-box">
		<view class="content">
			<form @submit="formSubmit" @reset="formReset">

				<uni-easyinput type="text" name="username" placeholder="请输入用户名" />


				<uni-easyinput type="text" name="nickName" placeholder="请输入昵称" />


				<uni-easyinput type="password" name="password" placeholder="请输入密码" />


				<radio-group name="gender">
					<label>
						<radio value="男" /><text>男</text>
					</label>
					<label style="padding-left: 15rpx;">
						<radio value="女" /><text>女</text>
					</label>
				</radio-group>



				<uni-easyinput type="text" name="phone" placeholder="请输入手机号" />

				<button form-type="submit">Submit</button>
				<!-- <button form-type="submit">Submit</button> -->
				<button form-type="reset">重置</button>
			</form>
		</view>
	</view>
</template>

<script>
	import { fixUserInfo, } from '../../utils/login.js'
	import { setToken } from '../../utils/auth.js'
	export default {
		data() {
			return {}
		},
		methods: {
			formSubmit: function(e) {
				// console.log('form发生了submit事件，携带数据为：' + JSON.stringify(e.detail.value))
				var formdata = e.detail.value
				console.log(formdata)
				fixUserInfo(formdata).then(res => {
					if(res.code === 200){
						setToken(res.data.token)
					}
				})
				
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
</style>
