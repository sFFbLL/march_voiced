<template>
	<view>
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

				<span>{{emojiList.faceTotal}}</span>
			</view>
			<view class="emoji" :class="{'clickEmoji':clickLike}" @click="clickAni('like')">
				<image class="aixin" src="../../static/img/aixin.png" mode=""></image>
				<span>{{emojiList.likeTotal}}</span>
			</view>
			<view class="emoji" :class="{'clickEmoji':clickFavour,'isDisplay':favourDisplay}" @click="clickAni('favour')">
				<image class="qingzhu" src="../../static/img/qingzhu.png" mode=""></image>
				<span>{{emojiList.favourTotal}}</span>
			</view>
			<!-- 添加表情 -->
			<view class="emoji" @click="addEmoji()">
				<image class="addemoji" src="../../static/img/emoji.png" mode=""></image>
				<image class="addemoji add" src="../../static/img/add.png" mode=""></image>
			</view>
			<!-- 评论+评论数量 -->
			<view class="comment" v-if="isShow">
				<u-icon name="chat" color="#999999" size="40" class="chat"></u-icon>
				<span class="commentTotal">{{emojiList.commentTotal}}</span>

			</view>
		</view>
	</view>
</template>

<script>
	export default {
		props: {
			emojiList: {

			},
			isShow:{
				type:Boolean,
				default:true
			}
		},
		data() {
			return {
				clickFace: false,
				clickLike: false,
				clickFavour: false,
				visible: true,
				faceDisplay: true,
				favourDisplay: true,
				emojiPosition01: 30,
			}
		},
		created() {
			if (this.emojiList.faceTotal > 0) {
				this.faceDisplay = false;
			}
			if (this.emojiList.favourTotal > 0) {
				this.favourDisplay = false;
			}
			// 当展示3个表情时
			if (this.emojiList.faceTotal > 0 && this.emojiList.favourTotal > 0) {
				this.emojiPosition01 = 140;
			}
			// 当展示两个表情时
			else if (this.emojiList.faceTotal > 0 || this.emojiList.favourTotal > 0) {
				this.emojiPosition01 = 85;
			}
		},
		methods: {
			addEmoji() {
				this.visible = !this.visible;
				// 当展示三个表情时
				if (this.emojiList.faceTotal > 0 && this.emojiList.favourTotal > 0) {
					this.emojiPosition01 = 140;
				}
				// 当展示两个表情时
				else if (this.emojiList.faceTotal > 0 || this.emojiList.favourTotal > 0) {
					this.emojiPosition01 = 85;
				}
				// 当展示一个表情时
				else if (this.emojiList.faceTotal <= 0 && this.emojiList.favourTotal <= 0) {
					this.emojiPosition01 = 30;
				}

			},
			clickAni(emoji) {
				switch (emoji) {
					case 'face':
						if (this.clickFace) {
							this.emojiList.faceTotal--;
							this.clickFace = !this.clickFace;
							if (this.emojiList.faceTotal == 0)
								this.faceDisplay = true;
						} else {
							this.clickFace = !this.clickFace;
							if (!this.clickLike && !this.clickFavour) {
								this.emojiList.faceTotal++;
							} else if (this.clickLike) {
								this.emojiList.likeTotal--;
								this.emojiList.faceTotal++;
								this.clickLike = !this.clickLike;
							} else if (this.clickFavour) {
								this.emojiList.favourTotal > 1 ? this.emojiList.favourTotal-- : this.favourDisplay = true;
								this.emojiList.faceTotal++;
								this.clickFavour = !this.clickFavour;
							}
						}

						break

					case 'like':
						if (this.clickLike) {
							this.emojiList.likeTotal--;
							this.clickLike = !this.clickLike;
						} else {
								this.clickLike = !this.clickLike;
							if (!this.clickFace && !this.clickFavour) {
								this.emojiList.likeTotal++;
							} else if (this.clickFace) {
								this.emojiList.likeTotal++;
								this.emojiList.faceTotal > 1 ? this.emojiList.faceTotal-- : this.faceDisplay = true;
								this.clickFace = !this.clickFace;
							} else if (this.clickFavour) {
								this.emojiList.favourTotal > 1 ? this.emojiList.favourTotal-- : this.favourDisplay = true;
								this.emojiList.likeTotal++;
								this.clickFavour = !this.clickFavour;
							}
						}


						break

					case 'favour':
						if (this.clickFavour) {
							this.emojiList.favourTotal--;
							this.clickFavour = !this.clickFavour;
							if (this.emojiList.favourTotal == 0)
								this.favourDisplay = true;

						} else {
							this.clickFavour = !this.clickFavour;
							if (!this.clickLike && !this.clickFace) {
								this.emojiList.favourTotal++;
							} else if (this.clickLike) {
								this.emojiList.likeTotal--;
								this.emojiList.favourTotal++;
								this.clickLike = !this.clickLike;
							} else if (this.clickFace) {
								this.emojiList.faceTotal > 1 ? this.emojiList.faceTotal-- : this.faceDisplay = true;
								this.emojiList.favourTotal++;
								this.clickFace = !this.clickFace;
							}
						}

						break

					default:
				}
			},
			addAEmoji(emoji) {
				// switch (emoji) {
				// 	case 'face':
				// 		if (this.emojiList.faceTotal >= 0)
				// 			this.faceDisplay = false;
				// 		break
				// 	case 'favour':
				// 		if (this.emojiList.favourTotal >= 0)
				// 			this.favourDisplay = false;
				// 		break

				// 	default:
				// }
				this.clickAni(emoji);

				this.visible = !this.visible;
			}
		}
	}
</script>

<style>
	/* 评论图标样式 */
	.comment {
		position: absolute;
		left: 650rpx;
	}

	.commentTotal {
		font-size: 24rpx;
		color: #999999;
		margin-left: 10rpx;
	}

	.chat {
		transform: rotateY(180deg);
	}

	/* 表情的样式 */
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

	.ideacontent {
		background-color: #FFFFFF;
		padding: 26rpx;

	}

	.idea {
		/* margin-top: 20rpx; */
		width: 100%;
		/* height: 75rpx; */
		/* line-height: 20rpx; */
		color: #101010;
		font-size: 14rpx;
		text-align: center;
		font-family: Arial;
		display: flex;
		border: 1rpx solid rgba(255, 255, 255, 100);
	}

	
</style>
