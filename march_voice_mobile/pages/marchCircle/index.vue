<template>
	<view>
		<!-- 三月圈头部 -->
		<view class="titlebg">
			<view class="title">
				<span class="marchcicrle">三月圈</span>
				<span class="member">成员：</span><span class="total">条数：</span>

				<p class="info">简介：</p>
			</view>
			<view>
				<image class="wxlogo" src="../../static/img/wxlogo.png"></image>
				<button class="join">加入</button>


			</view>


		</view>

		<!-- 三月圈内容 -->
		<view class="content">
			<!-- 添加表情上方提示框 -->
			<view class="allEmoji" :class="{'visible':visible}" :style="{marginLeft:emojiPosition01+'px'}">
				<image class="addxiaoku oneemoji" @click="addAEmoji('face')" src="../../static/img/xiaoku.png" mode=""></image>
				<image class="addaixin oneemoji" @click="addAEmoji('like')" src="../../static/img/aixin.png" mode=""></image>
				<image class="addqingzhu oneemoji" @click="addAEmoji('favour')" src="../../static/img/qingzhu.png" mode=""></image>
				<uni-icons class="arrow" type="arrowdown" size="15" style="color: #e6e6e6; "></uni-icons>
			</view>
			<!-- 各种表情 -->
			<view class="idea">
				<view class="emoji" :class="{'clickEmoji':clickFace,'isDisplay':faceDisplay}" @click="clickAni('face')">
					<image class=" xiaoku" src="../../static/img/xiaoku.png" mode=""></image>
					<span>{{faceTotal}}</span>
				</view>
				<view class="emoji" :class="{'clickEmoji':clickLike}" @click="clickAni('like')">
					<image class="aixin" src="../../static/img/aixin.png" mode=""></image>
					<span>{{likeTotal}}</span>
				</view>
				<view class="emoji" :class="{'clickEmoji':clickFavour,'isDisplay':favourDisplay}" @click="clickAni('favour')">
					<image class="qingzhu" src="../../static/img/qingzhu.png" mode=""></image>
					<span>{{favourTotal}}</span>
				</view>


				<!-- 添加表情 -->
				<view class="emoji" @click="addEmoji()">
					<image class="addemoji" src="../../static/img/emoji.png" mode=""></image>
					<image class="addemoji add" src="../../static/img/add.png" mode=""></image>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				clickFace: false,
				clickLike: false,
				clickFavour: false,
				visible: true,
				faceTotal: 0,
				likeTotal: 0,
				favourTotal: 0,
				commentTotal: 0,
				faceDisplay: true,
				favourDisplay: true,
				emojiPosition01: 30,
				emojiPosition02: false,
				emojiPosition03: false,

			}
		},
		methods: {
			addEmoji() {
				this.visible = !this.visible;
				// 当展示三个表情时
				if (this.faceTotal > 0 && this.favourTotal > 0) {
					this.emojiPosition01 = 140;
				}
				// 当展示两个表情时
				if (this.faceTotal > 0 || this.favourTotal > 0) {
					this.emojiPosition01 = 85;
				}
				// 当展示一个表情时
				if (this.faceTotal <= 0 && this.favourTotal <= 0) {
					this.emojiPosition01 = 30;
				}

			},
			clickAni(emoji) {
				switch (emoji) {
					case 'face':
						this.clickFace = !this.clickFace;
						if (this.clickFace) {
							this.faceTotal++;
						} else {
							this.faceTotal--;
							if (this.faceTotal == 0)
								this.faceDisplay = true;
						}
						this.clickFavour = false;
						this.clickLike = false;
						this.likeTotal > 0 ? this.likeTotal-- : 0;
						this.favourTotal > 0 ? this.favourTotal-- : 0;
						if (this.favourTotal == 0)
							this.favourDisplay = true;
						break

					case 'like':
						this.clickLike = !this.clickLike;
						if (this.clickLike) {
							this.likeTotal++;
						} else {
							this.likeTotal--;
						}
						this.clickFavour = false;
						this.clickFace = false;
						this.faceTotal > 0 ? this.faceTotal-- : 0;
						this.favourTotal > 0 ? this.favourTotal-- : 0;
						if (this.faceTotal == 0)
							this.faceDisplay = true;
						if (this.favourTotal == 0)
							this.favourDisplay = true;
						break
					case 'favour':
						this.clickFavour = !this.clickFavour;
						if (this.clickFavour) {
							this.favourTotal++;
						} else {
							this.favourTotal--;
							if (this.favourTotal == 0)
								this.favourDisplay = true;
						}
						this.clickLike = false;
						this.clickFace = false;
						this.faceTotal > 0 ? this.faceTotal-- : 0;
						this.likeTotal > 0 ? this.likeTotal-- : 0;
						if (this.faceTotal == 0)
							this.faceDisplay = true;
						break

					default:
				}

			},
			addAEmoji(emoji) {
				switch (emoji) {
					case 'face':
						if (this.faceTotal <= 0)
							this.faceDisplay = false;
						break
					case 'favour':
						if (this.favourTotal <= 0)
							this.favourDisplay = false;
						break

					default:
				}
				this.clickAni(emoji);

				this.visible = !this.visible;
			}
		}
	}
</script>

<style>
	.isDisplay {
		display: none;
	}

	.oneemoji {
		margin-left: 25rpx;
	}

	.addqingzhu {
		width: 55rpx;
		height: 50rpx;
		top: 8rpx;
	}

	.addaixin {
		width: 60rpx;
		height: 42rpx;
		top: 8rpx;
	}

	.addxiaoku {
		width: 55rpx;
		height: 50rpx;
		top: 8rpx;
	}

	.visible {
		visibility: hidden;
	}

	.arrow {
		margin-left: 110rpx;
	}

	.clickEmoji {
		background-color: #b8e3ff;
	}

	.allEmoji {
		width: 260rpx;
		height: 68rpx;
		/* margin-left: 60rpx; */
		border: 0.5rpx solid #e1e1e1;
		border-radius: 15rpx 15rpx 15rpx 15rpx;
		box-shadow: 1rpx 1rpx 10rpx 1rpx #e1e1e1;
		z-index: 5;
	}

	.qingzhu {
		width: 35rpx;
		height: 30rpx;
		top: 5rpx;
	}

	.aixin {
		width: 45rpx;
		height: 30rpx;
		top: 5rpx;
	}

	.xiaoku {
		width: 35rpx;
		height: 30rpx;
		top: 5rpx;
	}

	.addemoji {
		width: 30rpx;
		height: 30rpx;
		top: 5rpx;
	}

	.add {
		margin-left: 10rpx;
	}

	.emoji span {
		margin-left: 10rpx;
		top: 8rpx;
	}

	.emoji {
		width: 92rpx;
		height: 44rpx;
		line-height: 40rpx;
		border-radius: 10rpx 10rpx 10rpx 10rpx;
		color: #999999;
		font-size: 24rpx;
		text-align: center;
		font-family: Arial;
		margin-left: 20rpx;
		border: 2rpx solid #999999;
	}

	.content {
		background-color: #F5F5F5;
		height: 100vh;
	}

	.idea {
		margin-top: 20rpx;
		width: 100%;
		height: 175rpx;
		line-height: 20rpx;
		color: #101010;
		font-size: 14rpx;
		text-align: center;
		font-family: Arial;
		display: flex;
		border: 1rpx solid rgba(255, 255, 255, 100);
		background-color: #FFFFFF;
	}

	/* 头部样式 */

	.marchcicrle {
		font-size: 36rpx;
		width: 108rpx;
		height: 46px;
	}

	.member {
		font-size: 28rpx;
		margin: 0 20rpx;
	}

	.info {
		margin-top: 30rpx;
		font-size: 32rpx;
	}

	.total {
		font-size: 28rpx;
		margin: 80rpx 20rpx;
	}

	.titlebg {
		display: flex;
		width: 100%;
		height: 320rpx;
		background-image: url(../../static/img/marchcicrlebg.jpg);
		background-position: center;
		background-size: 100% 100%;
	}

	.title {
		margin: 72rpx 25rpx;
		width: 416rpx;
		color: #FFFFFF;
		font-weight: bold;

	}

	.wxlogo {
		width: 70rpx;
		height: 70rpx;
		left: 50rpx;
		top: 100rpx;
	}

	.join {
		left: 143rpx;
		top: 30rpx;
		width: 120rpx;
		height: 54rpx;

		line-height: 55rpx;
		border-radius: 28rpx 28rpx 28rpx 28rpx;
		background-color: #404040;
		color: #FFFFFF;
		font-size: 28rpx;
		text-align: center;
		font-family: Arial;
		border: 0.5rpx solid #101010 inset;
	}
</style>
