<template>
	<view class="attention-cell">
		<view class="flex-item">
			<!-- 左侧头像盒子 -->
			<a @click.stop="inToMine"
			 class="left-img-box inner-box">
				<image class="inner-img"
				 :src="getImgUrl(avatarPath)"
				 mode="aspectFill"></image>
			</a>
			<!-- 中部文字 -->
			<view class="middle-text-box inner-box"
			 @click.stop="inToMine">
				<view class="inner-middle-box inner-box">
					<!-- 昵称盒子 -->
					<view class="inner-text-name">
						{{nickname}}
						<!-- 昵称后跟随的文本内容 -->
						<slot name="afterNicknameText"></slot>
					</view>
					<!-- 中部添加文本（无长度限制，自动换行） -->
					<view class="inner-text-message greay-text">
						<slot name="middleText"></slot>
					</view>
					<!-- 底部添加文本（有长度限制，自动省略） -->
					<view class="inner-text-selfintroduce greay-text">
						<slot name="underText"></slot>
					</view>
				</view>
			</view>
			<!-- 按钮盒子 -->
			<view class="right-button-box">
				<!-- 无<view v-if="user.follow != null"> ================================================================== -->
				<view v-if="isFollow != null">
					<button v-show="isFollow === 0"
					 class="right-button"
					 type="default-green"
					 :disabled="isDisabled"
					 iconType="circle"
					 @click.stop="changeBtn">
						<span>关 注</span>
					 </button>
					<button v-show="isFollow === 1"
					 class="right-button"
					 type="default"
					 :disabled="isDisabled"
					 @click.stop="changeBtn"
					 iconType="circle">
						<span>已关注</span>
					</button>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import settings from '../settings.js'
	export default {
		props: {
			// 用于事件
			id: {
				type: Number,
				default: null
			},
			// 显示昵称
			nickname: {
				type: String,
				default: "昵称"
			},
			// 显示头像
			avatarPath: {
				type: String,
				default: null
			},
			// 是否关注
			isFollow: {
				type: Number,
				default: null
			},
		},
		data() {
			return {
				isDisabled: false ,// 是否禁用按钮点击
				openid:null,
			}
		},
		methods: {
			// 按钮切换显示
			changeBtn(id) {
				this.isDisabled = true
				let that = this;
				setTimeout(function () {
					that.isDisabled = false
					// 
					let params = {
						id: id
					}
					
					changeStatus(params).then(res =>{
						// 手动更改状态，刷新后查询新数据
						// switch (this.list[index].isFollow) {
						// 	case 0:
						// 		this.list[index].isFollow = 1
						// 		break;
						// 	case 1:
						// 		this.list[index].isFollow = 0
						// 		break;
						// }
						this.isFollow = !this.isFollow
					})
				}, 1000);
			},
			// 进入其他页面
			inToMine() {
				this.$emit('inToPageMine')
			},
			getImgUrl(src) {
				return settings.imgUrl + src;
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
		box-shadow: 0px 0px 1px 1px rgba(0, 0, 0, 0.2);
	}

	/* 中部文字盒子样式 */
	.middle-text-box {
		align-items: center;
		flex: 6;
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
		align-items: center
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
