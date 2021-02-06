<template>
	<view class="content">
		<uni-list class="list-item" v-for="(item,index) in recommendList1">
				<recommend :articleInfo="item" />
		</uni-list>
		<uniLoadMore></uniLoadMore>
	</view>
</template>

<script>
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	export default {
		components: {
			recommend,
			attentionAndFansCell,
			uniLoadMore
		},
		data() {
			return {
				recommendList1:[],
				recommendList: [{
					article: {
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						articleImg: require('static/img/2.jpg'),
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 1,
						userName: "张三",
						userImage: "",
						isFollow: 0
					}
				}, {
					article: {
						id: 2,
						title: "所以监听用户的截图操作，提示用户进行分",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						articleImg: "",
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 2,
						userName: "李四",
						userImage: "",
						isFollow: 1
					}
				},
				{
					article: {
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						articleImg: require('static/img/2.jpg'),
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 1,
						userName: "张三",
						userImage: "",
						isFollow: 0
					}
				}, {
					article: {
						id: 2,
						title: "所以监听用户的截图操作，提示用户进行分",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						articleImg: "",
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 2,
						userName: "李四",
						userImage: "",
						isFollow: 1
					}
				},
				{
					article: {
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						articleImg: require('static/img/2.jpg'),
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 1,
						userName: "张三",
						userImage: "",
						isFollow: 0
					}
				}, {
					article: {
						id: 2,
						title: "所以监听用户的截图操作，提示用户进行分",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						articleImg: "",
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 2,
						userName: "李四",
						userImage: "",
						isFollow: 1
					}
				}],
			}
		},
		methods: {
			// 进入时加载
			onLoad: function (options) {
				let that = this;
				this.recommendList1 = [];
				setTimeout(function() {
					that.recommendList1.push.apply(that.recommendList1, that.recommendList)
				}, 1000);
				uni.startPullDownRefresh();
			},
			// 下拉刷新
			onPullDownRefresh() {
				let that = this;
				// this.recommendList1 = [];
				setTimeout(function () {
					that.recommendList1 = [];
					that.recommendList1.push.apply(that.recommendList1, that.recommendList)
					uni.stopPullDownRefresh();
				}, 1000);
			},
			// 触底加载
			onReachBottom() {
				this.recommendList1.push.apply(this.recommendList1, this.recommendList)
			},
			recommend() {
				// console.log(1);
			},
			attention() {
				// console.log(2);
			},
			change(index) {
				this.list[index].isAttention = !this.list[index].isAttention
			}
		}
	}
</script>

<style>
	.content {
		background-color: rgba(241, 241, 241, 1);
	}
	.header {
		height: 100rpx;
		line-height: 100rpx;
		background-color: #fff;
		margin-bottom: 15rpx;
	}
</style>
