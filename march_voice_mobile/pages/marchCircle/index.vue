<template>
	<view>
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
		<view class="content">
			<view class="idea"></view>
			<view class="idea">
				<view class="emoji">
					<image class=" xiaoku"  src="../../static/img/xiaoku.png" mode="" ></image>
					<span>296</span>
				</view>
				<view class="emoji">
					<image class="aixin" src="../../static/img/aixin.png" mode="" ></image>
					<span>29</span>
				</view>
				<view class="emoji">
					<image class="qingzhu" src="../../static/img/qingzhu.png" mode="" ></image>
					<span>29</span>
				</view>
				<view class="emoji allEmoji" @click="addEmoji()">
					<image class="addemoji" src="../../static/img/emoji.png" mode="" ></image>
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
				
			}
		},
		methods: {
			addEmoji(){
				
			}
		}
	}
</script>

<style>
	.allEmoji{
		
	}
	.qingzhu{
		width: 35rpx;
		height: 30rpx;
			top:8rpx;
	}
	.aixin{
		width: 45rpx;
		height: 30rpx;
			top:8rpx;
	}
	.xiaoku{
		width: 35rpx;
		height: 30rpx;
			top:8rpx;
	}
	.addemoji{
		width: 30rpx;
		height: 30rpx;
		top:8rpx;
	}
	.add{
		margin-left: 10rpx;
	}
	.emoji span{
		margin-left: 10rpx;
		top: 8rpx;
	}
	.emoji{
		width: 92rpx;
		height: 44rpx;
		line-height: 40rpx;
		border-radius: 10rpx 10rpx 10rpx 10rpx;
		background-color: rgba(255, 255, 255, 100);
		color: #999999;
		font-size: 24rpx;
		text-align: center;
		font-family: Arial;
		margin-left: 20rpx;
		border: 2rpx solid rgba(153, 153, 153, 100);
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
