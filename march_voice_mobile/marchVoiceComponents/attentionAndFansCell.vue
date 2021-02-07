<template>
	<view class="attention-cell">
		<view class="flex-item">
			<!-- 左侧头像盒子 -->
			<a @click="inToPageMine" class="left-img-box inner-box">
				<image class="inner-img" :src=user.imgUrl mode="aspectFill"></image>
			</a>
			<!-- 中部文字 -->
			<view class="middle-text-box inner-box" @click="inToPageMine">
				<view class="inner-middle-box inner-box">
					<!-- 昵称盒子 -->
					<view class="inner-text-name">
						{{user.name}}<slot name="doSomeThing"></slot>
					</view>
					<!-- 关注数粉丝数文章数盒子 -->
					<view v-if="showDteial" class="inner-text-message greay-text">
						<span class="inner-text">关注 {{user.count}}</span><span class="inner-text">粉丝 {{user.count}}</span><span class="inner-text">文章
							{{user.count}}</span>
					</view>
					<!-- 个人介绍盒子 (单行文本)-->
					<view v-if="showIntrouduce" class="inner-text-selfintroduce greay-text">
						{{user.selfIntrouduce}}
					</view>
					<view v-if="showDate" class="inner-text-selfintroduce greay-text">
						{{user.date}}
					</view>
				</view>
			</view>
			<!-- 按钮盒子 -->
			<view class="right-button-box">
				<button v-show="user.follow === false" class="right-button" type="default-green" :disabled="isDisabled"
				 iconType="circle" @click.stop="changeBtn"><span>关 注</span></button>
				<button v-show="user.follow === true" class="right-button" type="default" :disabled="isDisabled"
				 @click.stop="changeBtn" iconType="circle">
					<slot name="hasAttention"></slot>
				</button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		props: {
			user:{
				type:Object,
				default(){
					return{
						id: "",// 用于事件
						imgUrl: require('../static/img/1.jpg'),// 显示头像
						name:"昵称",// 显示昵称
						count:"0",// 关注数粉丝数文章数
						selfIntrouduce:"",// 个人简介
						follow:true,// 判断是否互相关注
						date:"2021-02-07 11:59"
					}
				}
			},
			// 显示个人简介
			showIntrouduce: {
				type: Boolean,
				default: true
			},
			// 是否展示关注粉丝文章数
			showDteial: {
				type: Boolean,
				default: true
			},
			showDate:{
				type: Boolean,
				default: true
			}
		},
		data() {
			return {
				isLoading: false, // 是否为加载中
				isDisabled: false // 是否禁用按钮点击
			}
		},
		methods: {
			// 按钮切换显示
			changeBtn() {
				this.isDisabled = true
				this.isLoading = true
				let that = this;
				setTimeout(function() {
					that.isDisabled = false
					that.isLoading = false
					that.$emit('change');
				}, 1000);
			},
			// 进入其他页面
			inToPageMine() {
				console.log(this.id)
			}
		}
	}
</script>

<style scoped>
	/* 大盒子样式 */
	.flex-item {
		display: flex;
		box-sizing: border-box;
		border-bottom: 6rpx solid #f0f0f0;
		padding: 20rpx 30rpx 20rpx 30rpx;
		/* margin-bottom: 20rpx; */
	}

	/* 左侧图片盒子样式 */
	.left-img-box {
		display: flex;
		align-items: center;
		flex: 1;
		margin-right: 10rpx;
		/* width: 112rpx; */
		/* height: 112rpx; */
		vertical-align: middle;
	}

	/* 内部盒子样式 */

	/* .inner-box {
		display: inline-block;
	} */

	/* .inner-box {
    display: inline-block;
  } */

	/* 图片样式 */
	.inner-img {
		height: 70rpx;
		width: 70rpx;
		border-radius: 50%;
	}

	/* 中部文字盒子样式 */
	.middle-text-box {
		align-items: center;
		flex: 5;
		display: flex;
	}

	/* 灰色字体 */
	.greay-text {
		color: #969696;
	}

	/* 昵称 */
	.inner-text-name {
		font-size: 30rpx;
		font-weight: 700;
		vertical-align: middle;
	}

	/* 关注数粉丝数文章数 */
	.inner-text-message {
		padding-top: 4rpx;
		font-size: 24rpx;
	}

	/* 关注数粉丝数文章数去掉最后的 |  */
	.inner-text-message span:not(:last-child) {
		margin-right: 10rpx;
		padding-right: 10rpx;
		border-right: 1rpx solid #969696;
	}

	/* 自我介绍 */
	.inner-text-selfintroduce {
		padding-top: 4rpx;
		font-size: 24rpx;
		width: 420rpx;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	/* 右侧按钮盒子 */
	.right-button-box {
		display: flex;
		flex: 1;
		align-items:center
	}

	/* 右侧按钮 */
	.right-button {
		border-radius: 58rpx;
		width: 130rpx;
		height: 56rpx;
		font-size: 27rpx;
		font-weight: 400;
		line-height: 56rpx;
		color: #969696;
	}

	/* 修改按钮样式 */
	uni-button[type=default] {
		border: 1rpx #969696 solid;
		box-sizing: border-box;
	}

	/* 修改按钮样式 */
	uni-button[type=default-green] {
		color: #FFF;
		background-color: rgba(64, 64, 64, 1);
	}

	/* 修改点击按钮后的样式 */
	.button-hover[type=default-green] {
		background-color: rgba(26, 26, 26, 1.0);
	}

	/* 修改按钮样式 */
	uni-button {
		margin: 0;
		padding: 0;
	}
</style>
