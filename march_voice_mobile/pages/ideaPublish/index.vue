<template>
	<view class="wrap">
		<view>
			<textarea class="input" placeholder="这一刻的想法" maxlength="-1" :value="ideaWords" @input="getWords" />
			</view>
		<view class="imgs">
			<button @click="upImg()" style="width: 100px;height: 100px;"><image src="../../static/img/add.png" style="width: 100rpx;height: 100rpx; margin-top: 50rpx;"></image></button>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				ideaWords: "value"
			}
		},
		methods: {
			getWords(event) {
				this.ideaWords = event.target.value;
				console.log(this.ideaWords)
			},
			upImg(){
				uni.chooseImage({
					success(res) {
						console.log(JSON.stringify(res.tempFilePaths))
					}
				})
			}
		},
		created() {
			console.log(this.ideaWords)
		}
	}
</script>

<style scoped>
	/* 图片部分 */
	.imgs{
		display: flex;
	}
	/* 文字部分 */
	.input {
		width: 100%;
		height: 200rpx;
		text-align: left;
		overflow: hidden;
		-webkit-line-clamp: 2;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		background-color: #bacde1;
		margin-bottom: 20rpx;
	}

	.wrap {
		padding: 20rpx;
	}
</style>
