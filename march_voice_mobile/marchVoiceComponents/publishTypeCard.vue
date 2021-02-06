<template>
	<view class="type-list-item" @click="toPublishPage">
		<view class="item-text">
			<view class="item-text-title">
				<slot name="title"></slot>
			</view>
			<view class="item-text-describe">
				<slot name="describe"></slot>
			</view>
		</view>
		<view class="item-img-box">
			<u-image class="inner-img" :src="src" height="60rpx" width="60rpx" :fade="fade" shape="circle"></u-image>
		</view>
	</view>
</template>

<script>
	export default {
			props:{
				type:{
					type:String,
					default:"text"
				},
				
			},
			data() {
				return {
					src: require('../static/img/1.jpg'),
					fade:false,
				}
			},
			methods:{
				toPublishPage(){
					uni.navigateTo({
						url:"richTextEditor?type="+this.type,
					})
				}
			}
		}
</script>

<style scoped>
	.type-list-item {
		bottom: 100rpx;
		box-sizing: border-box;
		bottom: 50px;
		display: flex;
		padding-bottom: 20rpx;
		/* padding: 20rpx 46rpx 20rpx 46rpx; */
	}

	.item-text {
		display: flex;
		flex: 9;
		flex-direction: column;
	}

	.item-text-title,
	.item-text-describe {
		display: flex;
		align-items: center;
		flex: 1;
	}

	.item-text-title {
		letter-spacing: 2px;
		color: rgba(0, 92, 255, 1);
		font-size: 16px;
		line-height: 150%;
		font-weight: bold;
	}

	.item-text-describe {
		color: rgba(153, 153, 153, 1);
		font-size: 14px;
		line-height: 150%;
	}

	.item-img-box {
		flex: 2;
		display: flex;
		align-items: center;
	}
	
	::v-deep .inner-img{
		position: absolute;
		right: 46rpx;
	}
</style>

