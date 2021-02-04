<template>
	<view class="content">
		<view v-for="(item,index) in recommendList">
			<attentionAndFansCell :showDteial="false" @change="change(index)"></attentionAndFansCell>
			<recommend :articleInfo="item" />
		</view>
	</view>
</template>

<script>
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	export default {
		components: {
			recommend,
			attentionAndFansCell
		},
		data() {
			return {
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
				}],
			}
		},
		methods: {
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

<style scoped>
	::v-deep .attention-cell>.flex-item{
		border-bottom: none;
		border-top: 1rpx #d9d9d9 solid;
		padding: 20rpx 30rpx 0rpx 30rpx;
	}
</style>
