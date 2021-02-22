<template>
	<view class="content">
		<!---->
		<view class="base-info-tabbar">
			<text>基本信息</text>
		</view>

		<view class="base-info-content">
			<view class="base-info-head-img" style="height: 130rpx;">
				<text clas="base-info-area">头像</text>
				<view class="base-info-change-img">
					<image style="width: 100rpx; height: 100rpx; border-radius: 50rpx;" mode="aspectFill" :src="info.avatarPath"
					 @click="changeHeadImg()"></image>
					<uni-icons style="margin-left: 10rpx;" type="arrowright" size="15"></uni-icons>
				</view>

			</view>

			<view class="base-info-name">
				<text clas="base-info-area">昵称</text>
				<view class="base-info-change-name">
					<view class="limit-lengh">
						<text class="title" @click="changeNameOpen()">{{info.nickname}}</text>
					</view>
					<!-- 修改昵称弹出层 -->
					<uni-popup ref="npopup" type="dialog">
						<uni-popup-dialog type="info" mode="input" title="请修改昵称" placeholder="请输入2-15个字符" :duration="2000" :before-close="true"
						 @close="close" @confirm="changeNameConfirm"></uni-popup-dialog>
					</uni-popup>
					<!-- 警告弹出层 -->
					<uni-popup ref="wnpopup" type="dialog">
						<uni-popup-dialog mode="base" type="error" title="警告" content="请输入2-15个字符" :duration="2000" :before-close="true"
						 @close="close" @confirm="close"></uni-popup-dialog>
					</uni-popup>

					<uni-icons style="margin-left: 10rpx;" type="arrowright" size="15"></uni-icons>
				</view>
			</view>
			<view class="base-info-gender">
				<text clas="base-info-area">性别</text>
				<view class="base-info-gender-change">
					<picker @change="bindPickerChange" :value="info.sex" :range="array">
						<text class="uni-input">{{array[info.sex]}}</text>
					</picker>
					<uni-icons style="margin-left: 10rpx;" type="arrowright" size="15"></uni-icons>
				</view>
			</view>

			<view class="base-info-person-info">
				<text clas="base-info-area">个人简介</text>
				<view class="base-info-change-info">
					<view class="limit-lengh">
						<!-- -->
						<text class="title" @click="changeInfoOpen()">{{info.signature}}</text>
					</view>
					<!-- 修改个人简介弹出框 -->
					<uni-popup ref="popup1" type="dialog">
						<uni-popup-dialog mode="input" title="请修改个人简介" placeholder="用一句话介绍一下自己吧(50字以内)" :duration="2000" :before-close="true"
						 @close="close" @confirm="changeInfoConfirm"></uni-popup-dialog>
					</uni-popup>

					<uni-popup ref="popup1Up" type="dialog">
						<uni-popup-dialog mode="base" type="error" title="警告" content="请输入2-15个字符" :duration="2000" :before-close="true"
						 @close="close" @confirm="close"></uni-popup-dialog>
					</uni-popup>
					<uni-icons style="margin-left: 10rpx;" type="arrowright" size="15" @click="changeInfoOpen()"></uni-icons>
				</view>
			</view>
			<u-top-tips ref="sTips" :navbar-height="0"></u-top-tips>
		</view>
	</view>
</template>

<script>
	import {
		information,
		modInformation
	} from "../../utils/api/personInfo-api.js"
	import unloadImage from "../../utils/api.js"
	export default {
		data() {
			return {
				array: ['男', '女'],
				id: 0,
				info: {
					nickname: "xianer",
					sex: "0",
					signature: "这个人很懒，什么也没有写",
					avatarPath: "../../static/img/my1.png"

				},

			}
		},
		// 获取我的id
		onLoad(option) {
			this.id = option.id;
		},
		created() {
			// 调用查询用户信息的接口
			let id = this.id;
			let _this = this;
			// information(id).then(res=>{
			// 	_this.info=res.data[0];
			// })
		},

		methods: {
			toptip() {
				this.$refs.sTips.show({
					title: '设置成功',
					type: 'success',
					duration: '2300',
				});
			},
			bindPickerChange(e) {
				// console.log('picker发送选择改变，携带值为', e.target.value)
				this.info.sex = e.target.value;
				let sex = this.info.sex;
				let _this = this;
				modInformation(sex).then(res => {
					_this.toptip();
				})

			},
			// 更换头像
			changeHeadImg: function(e) {
				let _this = this;
				uni.chooseImage({
					count: 1, //默认9
					sizeType: ['compressed'], //可以指定是原图还是压缩图，默认二者都有
					sourceType: ['album', 'camera'], //从相册选择
					success: function(res) {
						_this.info.avatarPath = res.tempFilePaths[0];
						// 调用接口转化imgurl
						let file = res.tempFilePaths
						let avatarPath = unloadImage(file);
						// 调用修改信息接口
						let that = _this
						modInformation(avatarPath).then(res => {
							that.toptip()
						})
					}
				});
			},
			// 控制修改昵称弹出层打开
			changeNameOpen() {
				this.$refs.npopup.open()
			},
			/**
			 * 点击取消按钮触发
			 * @param {Object} done
			 */
			close(done) {
				// TODO 做一些其他的事情，before-close 为true的情况下，手动执行 done 才会关闭对话框
				// ...

				done();
			},

			/**
			 * 点击确认按钮触发
			 * @param {Object} done
			 * @param {Object} value
			 */
			changeNameConfirm(done, value) {
				// 输入框的值2-15个字符
				if (/^.{2,15}$/.test(value)) {
					//true
					this.info.nickname = value;
					let nickname = this.info.nickname
					let _this = this;
					modInformation(nickname).then(res => {
						_this.toptip()
					})

				} else {
					this.$refs.wnpopup.open();
					return;
				}

				// TODO 做一些其他的事情，手动执行 done 才会关闭对话框
				// ...
				done()
			},

			changeInfoOpen() {
				this.$refs.popup1.open()
			},

			changeInfoConfirm(done, value) {
				// 输入框的值0-50个字符
				if (/^.{2,50}$/.test(value)) {
					//true
					this.info.signature = value;
					let signature = this.info.signature
					let _this = this;
					modInformation(signature).then(res => {
						_this.toptip()
					})
				} else {
					this.$refs.popup1Up.open();
					return;
				}
				// TODO 做一些其他的事情，手动执行 done 才会关闭对话框
				// ...
				done()
			},
		}
	}
</script>

<style>
	.limit-lengh {
		width: 300rpx;
		text-align: right;
		white-space: nowrap;
		text-overflow: ellipsis;
		overflow: hidden;
	}

	.base-info-tabbar {
		font-size: 30rpx;
		font-weight: bold;
		display: flex;
		flex-direction: column;
		border-bottom: 1px solid #F1F1F1;
		padding: 0rpx 28rpx;
	}

	.base-info-content {
		color: #404040;
		padding: 0rpx 28rpx;
		font-size: 32rpx;
		display: flex;
		flex-direction: column;
	}

	.content>.base-info-tabbar {
		height: 86rpx;
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

	.base-info-gender-change,
	.base-info-change-img,
	.base-info-change-name,
	.base-info-change-info {
		display: flex;
		align-items: center;
	}

	.base-info-content>view {
		height: 86rpx;
		line-height: 86rpx;
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
	}

	.base-info-content>* {
		border-bottom: 1px dashed #F1F1F1;
	}
</style>
