<template>
	<uni-list class="list-item">
		<attentionAndFansCell class="item" v-for="(item, index) in list" 
		:key="item.id" 
		@change="change(index)"
		@inToPageMine="inToPageMine(index)"
		:id="item.id"
		:nickname="item.nickname"
		:isFollow="item.isFollow"
		:avatarPath="item.avatarPath"
		>
			 <span slot="afterNicknameText">关注了你</span>
			 <span slot="middleText">
				<span class="inner-text">关注 {{item.fansNumber}}</span>
				 <span class="inner-text">粉丝 {{item.followNumber}}</span>
				 <span class="inner-text">文章 {{item.articleNumber}}</span>
			 </span>
			 <span slot="underText">{{item.signature}}</span>
			 </attentionAndFansCell>
		<uniLoadMore></uniLoadMore>
	</uni-list>
</template>

<script>
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	export default {
		data() {
			return {
				list: [],
				list1: [{
					id: "0", // 用于事件
					avatarPath: '../../static/img/1.jpg', // 显示头像
					nickname: "吴胜科", // 显示昵称
					signature: "阿萨德",
					isFollow: 0, // 判断是否互相关注
					fansNumber: "1", // 显示粉丝数
					followNumber:"2", // 显示关注数
					articleNumber:"3", // 显是文章数
				}]
			}
		},
		components: {
			attentionAndFansCell,
			uniLoadMore
		},
		// 进入时加载
		onLoad: function(options) {
			let that = this;
			this.list = [];
			// setTimeout(function() {
			// 	that.list.push.apply(that.list, that.list1)
			// }, 1000);
			uni.startPullDownRefresh();
		},
		// 下拉刷新
		onPullDownRefresh() {
			let that = this;
			that.list = [];
			setTimeout(function() {
				that.list.push.apply(that.list, that.list1)
				uni.stopPullDownRefresh();
			}, 1000);
		},
		// 触底加载
		onReachBottom() {
			this.list.push.apply(this.list, this.list1)
		},
		methods: {
			// 跳转页面
			inToPageMine(index) {
				console.log(index)
			},
			// 按钮样式切换
			change(index) {
				
				switch(this.list[index].isFollow){
					case 0:
						this.list[index].isFollow = 1;
						break;
					case 1:
						this.list[index].isFollow = 0;
						break;
				}
			}
		}
	}
</script>

<style scoped>
/* 关注数粉丝数文章数去掉最后的 |  */
	.inner-text-message span:not(:last-child) {
		margin-right: 10rpx;
		padding-right: 10rpx;
		border-right: 1rpx solid #969696;
	}
</style>
