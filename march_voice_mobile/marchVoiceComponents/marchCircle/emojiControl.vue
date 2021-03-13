<template>
	<view>

		<!-- 各种表情 -->
		<view class="idea">
			<view class="emoji" :class="{'clickEmoji':clickFace,'isDisplay':faceDisplay}" @click="clickAni('face')">
				<image class=" xiaoku" src="../../static/img/xiaoku.png" mode=""></image>
				<span>{{faceTotal}}</span>
			</view>
			<view class="emoji"
			 :class="{'clickEmoji':clickLike}"
			 @click="clickAni('like')">
				<image class="aixin"
				 src="../../static/img/aixin.png"
				 mode=""></image>
				<span>{{likeTotal}}</span>
			</view>
			<view class="emoji"
			 :class="{'clickEmoji':clickFavour,'isDisplay':favourDisplay}"
			 @click="clickAni('favour')">
				<image class="qingzhu"
				 src="../../static/img/qingzhu.png"
				 mode=""></image>
				<span>{{favourTotal}}</span>
			</view>
			<!-- 添加表情 -->
			<view class="emoji" @click="addEmoji()">
				<image class="addemoji" src="../../static/img/emoji.png" mode=""></image>
				<image class="addemoji add" src="../../static/img/add.png" mode=""></image>
				<!-- 添加表情上方提示框 -->
				<view class="allEmoji" :class="{'visible':visible}">
					<image class="addxiaoku oneemoji" @click="addAEmoji('face')" src="../../static/img/xiaoku.png" mode=""></image>
					<image class="addaixin oneemoji" @click="addAEmoji('like')" src="../../static/img/aixin.png" mode=""></image>
					<image class="addqingzhu oneemoji" @click="addAEmoji('favour')" src="../../static/img/qingzhu.png" mode=""></image>
					<u-icon class="arrow" name="arrow-down-fill" color="#f5f5f5" size="30"></u-icon>
				</view>
			</view>
			<!-- 评论+评论数量 -->

			<view class="comment"
			 v-if="isShow" >
				<u-icon name="chat"
				 color="#999999"
				 size="40"
				 class="chat"></u-icon>
				<span class="commentTotal">{{commentTotal}}</span>


			</view>
		</view>
	</view>
</template>

<script>
	import {changeFavour} from '../../utils/api/marchCircle-api.js'
	export default {
		props: {
			id: {
				type: Number
			},
			faceTotals: {
				type: Number
			},
			likeTotals: {
				type: Number
			},
			favourTotals: {
				type: Number
			},
			commentTotals: {
				type: Number
			},
			// 只有详情页的时候才传false

			isShow: {
				type: Boolean,
				default: true
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
				faceTotal: this.faceTotals,
				likeTotal: this.likeTotals,
				favourTotal: this.favourTotals,
				commentTotal: this.commentTotals
			}
		},

		created() {
			if (this.faceTotal > 0) {
				this.faceDisplay = false;
			}
			if (this.favourTotal > 0) {
				this.favourDisplay = false;
			}
			// 当展示3个表情时
			if (this.faceTotal > 0 && this.favourTotal > 0) {
				this.emojiPosition01 = 140;
			}
			// 当展示两个表情时
			else if (this.faceTotal > 0 || this.favourTotal > 0) {
				this.emojiPosition01 = 85;
			}
		},
		methods: {
			// 控制展示三个表情的位置
			addEmoji() {
				this.visible = !this.visible;
				// 当展示三个表情时
				if (this.faceTotal > 0 && this.favourTotal > 0) {
					this.emojiPosition01 = 140;
				}
				// 当展示两个表情时
				else if (this.faceTotal > 0 || this.favourTotal > 0) {
					this.emojiPosition01 = 85;
				}
				// 当展示一个表情时
				else if (this.faceTotal <= 0 && this.favourTotal <= 0) {
					this.emojiPosition01 = 30;
				}

			},


			// 调用接口对点赞状态改变
			emjoiInterface(params) {
				changeFavour(params).then(res => {
					console.log("表情点赞状态改变")
				}).catch(err=>{
					console.log("err")
				})
			},
			// 点击表情事件处理
			clickAni(emoji) {
				let params = {
					id: this.id,
					type: 0
				}
				console.log(this.clickFace)
				switch (emoji) {

					case 'face':

						// 如果face表情已经处于点击状态
						if (this.clickFace) {
							this.faceTotal--;
							this.clickFace = !this.clickFace;
							if (this.faceTotal == 0)
								this.faceDisplay = true;
						} else { //如果face表情处于未点击状态
							this.clickFace = !this.clickFace;
							this.faceDisplay = false;
							// 1.并且另外两个表情也处于未点击状态
							if (!this.clickLike && !this.clickFavour) {
								this.faceTotal++;
							} else if (this.clickLike) { //2.like表情正处于点击状态,
								this.likeTotal--;
								this.faceTotal++;
								this.clickLike = !this.clickLike;
							} else if (this.clickFavour) {
								this.favourTotal--;
								if (!this.favourTotal)
									this.favourDisplay = true;
								this.faceTotal++;
								this.clickFavour = !this.clickFavour;
							}
						}
						this.emjoiInterface(params)
						break

					case 'like':

						if (this.clickLike) {
							this.likeTotal--;
							this.clickLike = !this.clickLike;
						} else {
							this.clickLike = !this.clickLike;
							if (!this.clickFace && !this.clickFavour) {
								this.likeTotal++;
							} else if (this.clickFace) {
								this.likeTotal++;
								this.faceTotal--
								if (!this.faceTotal)
									this.faceDisplay = true;
								this.clickFace = !this.clickFace;
							} else if (this.clickFavour) {
								this.favourTotal--
								if (!this.favourTotal)
									this.favourDisplay = true;
								this.likeTotal++;
								this.clickFavour = !this.clickFavour;
							}
						}
						params.type = 1;
						this.emjoiInterface(params)
						break

					case 'favour':

						if (this.clickFavour) {
							this.favourTotal--;
							this.clickFavour = !this.clickFavour;
							if (this.favourTotal == 0)
								this.favourDisplay = true;

						} else {
							this.clickFavour = !this.clickFavour;

							this.favourDisplay = false;
							if (!this.clickLike && !this.clickFace) {
								this.favourTotal++;
							} else if (this.clickLike) {
								this.likeTotal--;
								this.favourTotal++;
								this.clickLike = !this.clickLike;
							} else if (this.clickFace) {
								this.faceTotal--
								if (!this.faceTotal)
									this.faceDisplay = true;
								this.favourTotal++;
								this.clickFace = !this.clickFace;
							}
						}
						params.type = 2;
						this.emjoiInterface(params)

						break

					default:
				}
			},

			// 添加表情事件处理
			addAEmoji(emoji) {
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
		display: none;
	}

	.arrow {
		position: absolute;
		top: 55rpx;
		left: 115rpx;
	}

	.clickEmoji {
		background-color: #b8e3ff;
	}

	.allEmoji {
		position: absolute;
		bottom: 64rpx;
		left: -80rpx;
		width: 260rpx;
		height: 68rpx;
		background-color: #f5f5f5;
		border: 0.5rpx solid #e1e1e1;
		border-radius: 15rpx 15rpx 15rpx 15rpx;
		box-shadow: 1rpx 1rpx 10rpx 1rpx #e1e1e1;
		z-index: 10;
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
		position: relative;
		width: 92rpx;
		height: 44rpx;
		line-height: 40rpx;
		border-radius: 10rpx 10rpx 10rpx 10rpx;
		color: #999999;
		font-size: 24rpx;
		text-align: center;
		font-family: Arial;
		margin-right: 20rpx;
		border: 2rpx solid #999999;
	}

	.ideacontent {
		background-color: #FFFFFF;
		padding: 26rpx;

	}

	.idea {
		margin-top: 20rpx;
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
