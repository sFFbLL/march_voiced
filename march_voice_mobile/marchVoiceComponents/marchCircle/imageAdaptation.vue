<template>
	<view>
		<view class="allImage">
			<view class="images" v-for="(item,index) in imgList" :key="index">
				<image class="oneimg" :src="item" mode="aspectFill" :style="{width:imgWidth+'rpx',height:imgHeight+'rpx'}"></image>
			</view>
		</view>
	</view>
</template>

<script>
	export default{
		props: {
			imgList:{
				type:Array
			}
		},
		data() {
			return {
				imgWidth: 702,
				imgHeight: 300,
				imgDisplay: 'flex',
			}
		},
		mounted() {
			this.judgeImg();
		},
		methods: {
			judgeImg() {
				if (this.imgList == null) {
					this.imgDisplay = 'none';
				} else if (this.imgList.length == 1) {
					this.imgHeight = 300;
					this.imgWidth = 702;
				} else if (this.imgList.length == 2 || this.imgList.length == 4) {
					this.imgWidth = 321;
					this.imgHeight = 280;
				} else {
					this.imgWidth = 220;
					this.imgHeight = 220;
				}
			}
		}
	}
</script>

<style>
	/* 想法图片排列样式 */
	
	.allImage {
		display: flex;
		margin-top: 10rpx;
		flex-wrap: wrap;
		justify-content: flex-start;
	}
	
	.images:not(:nth-child(3n)) {
		/* margin-right: 10rpx; */
	}
	
	.images {
		margin-right: 10rpx;
	}
	
	
</style>
