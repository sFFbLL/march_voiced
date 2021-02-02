<template>
	<view class="">
		<view class="flex-item" @click="inToPageMine">
			<a class="left-img-box inner-box">
				<image class="inner-img" :src=imgUrl mode="heightFix"></image>
			</a>
			<view class="middle-text-box inner-box">
				<view class="inner-middle-box inner-box">
					<view class="inner-text-name">
						{{title}}
					</view>
					<view class="inner-text-message greay-text">
						<span class="inner-text">关注 {{message}}</span><span class="inner-text">粉丝 {{message}}</span><span class="inner-text">文章
							{{message}}</span>
					</view>
					<view class="inner-text-selfintroduce greay-text">
						{{selfIntrouduce}}
					</view>
				</view>
			</view>
			<view class="right-button-box">
				<button v-show="isAttention === false" class="right-button" type="default-green" :loading="isLoading" :disabled="isDisabled"
				 @click.stop="changeBtn">关 注</button>
				<button v-show="isAttention === true" class="right-button" type="default" :loading="isLoading" :disabled="isDisabled"
				 @click.stop="changeBtn">√ 已关注</button>
			</view>
		</view>
	</view>

</template>

<script>
	export default {
		props: {
			id: {
				type: String,
				default: ""
			},
			imgUrl: {
				type: String,
				default: require('../static/img/my1.png')
			},
			title: {
				type: String,
				default: "此处为昵称"
			},
			message: {
				type: String,
				default: "0"
			},
			selfIntrouduce: {
				type: String,
				default: "此处为个人奥术大师大大所多所多所大所简介"
			},
			// isAttention: {
			// 	type:Boolean,
			// 	default:false
			// }
		},
		data() {
			return {
				isAttention: false,
				isLoading: false,
				isDisabled: false
			}
		},
		methods: {
			changeBtn() {
				this.isDisabled = true
				this.isLoading = true
				let that = this;
				setTimeout(function() {
					console.log(that.isDisabled)
					that.isDisabled = false
					that.isLoading = false
					console.log(that.isDisabled)
					that.isAttention = !that.isAttention
				}, 2000);
				console.log("奥!")
			},
			// cancleAttenton() {
			// 	this.isDisabled = true
			// 	this.isLoading = true
			// 	let that = this;
			// 	setTimeout(function() {
			// 		console.log(that.isDisabled)
			// 		that.isDisabled = false
			// 		that.isLoading = false
			// 		console.log(that.isDisabled)
			// 		that.isAttention = !that.isAttention
			// 	}, 2000);
			// 	console.log("取消了奥!")
			// },
			inToPageMine() {
				console.log()(this.id)
			}
		}
	}
</script>

<style scoped>
	.flex-item {
		display: flex;
		box-sizing: border-box;
		border-bottom: 2rpx solid #f0f0f0;
		padding: 0 48rpx 20rpx 48rpx;
		margin-bottom: 20rpx;
	}

	.left-img-box {
		flex: 1;
		margin-right: 16rpx;
		/* width: 112rpx; */
		/* height: 112rpx; */
		vertical-align: middle;
	}

	.inner-box {
		display: inline-block;
	}

	.inner-img {
		height: 100%;
		width: 100%;
	}

	.middle-text-box {
		flex: 5;
		display: flex;
	}

	.greay-text {
		color: #969696;
	}

	.inner-text-name {
		font-size: 30rpx;
		font-weight: 700;
	}

	.inner-text-message {
		padding-top: 4rpx;
		font-size: 24rpx;
	}

	.inner-text-message span:not(:last-child) {
		margin-right: 10rpx;
		padding-right: 10rpx;
		border-right: 1rpx solid #969696;
	}

	.inner-text-selfintroduce {
		padding-top: 4rpx;
		font-size: 24rpx;
		width: 360rpx;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	.right-button-box {
		display: flex;
		flex: 2;
		align-items: center;
	}

	.right-button {
		border-radius: 112rpx;
		width: 160rpx;
		height: 50rpx;
		font-size: 23rpx;
		font-weight: 400;
		line-height: 50rpx;
		color: #969696;
	}

	uni-button[type=default] {
		border: 1rpx #969696 solid;
		box-sizing: border-box;
	}

	uni-button[type=default-green] {
		color: #FFF;
		background-color: #42c02e;
	}

	.button-hover[type=default-green] {
		background-color: #3dac29;
	}
</style>
