<template>
	<view class="wrap">


		<!-- 用户头像公共组件 -->
		<attentionAndFansCell :aid="messageInfo.userId" :nickname="messageInfo.nickname" :avatarPath="messageInfo.avatarPath">
			<template v-slot:afterNicknameText>
				<text class="slot">{{tip}}
				</text>
				<image v-if="messageInfo.status==3&&messageInfo.type==1" src="../../static/img/aixin.png" mode="" style="width: 44rpx;height: 30rpx; margin-left: 10rpx;"></image>
				<image v-if="messageInfo.status==2&&messageInfo.type==1" src="../../static/img/xiaoku.png" mode="" style="width: 34rpx;height: 30rpx; margin-left: 10rpx;"></image>
				<image v-if="messageInfo.status==4&&messageInfo.type==1" src="../../static/img/qingzhu.png" mode="" style="width: 34rpx;height: 30rpx; margin-left: 10rpx;"></image>

			</template>
			<template v-slot:underText>
				<text>{{createDate}}</text>
			</template>
		</attentionAndFansCell>
		<!-- 评论内容 -->
		<text v-if="messageInfo.comment" class="comment">{{messageInfo.comment}}</text>
		<!-- 文章/想法部分 -->
		<view class="thecontent">
			<articleTitle class="thewords" v-if="messageInfo.type==0" :articleTitle="messageInfo.title" :articleId="messageInfo.articleId" />
			<articleTitle class="thewords" v-if="messageInfo.type==1" :articleTitle="messageInfo.content" :articleId="messageInfo.articleId" />
			<image v-if="messageInfo.image" :src=messageInfo.image mode="scaleToFill" class="theimg"></image>
		</view>
		<!-- 间隔槽 -->
		<u-gap height="10" bg-color="#f5f5f5"></u-gap>
	</view>
</template>

<script>
	import articleTitle from "../../marchVoiceComponents/showArticle/childComponents/articleTitle.vue"

	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	export default {
		props: {
			messageInfo: {
				type: Object,
				default () {
					return {
						type: 0,
						status: 1,
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						articleId: 1,
						comment: "哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
						image: require('static/img/2.jpg'),
						createTime: "2020/12/12",
						userId: "1",
						avatarPath: require('static/img/1.jpg'),
						nickname: "张三",
					}
				}
			}
		},
		components: {
			attentionAndFansCell,
			articleTitle
		},
		data() {
			return {
				tip: "赞了你的文章",
			}
		},
		computed: {
			createDate: function() {
				let oDate = new Date(this.messageInfo.createTime);
				let oYear = oDate.getFullYear();
				let oMonth = oDate.getMonth() + 1;
				let oDay = oDate.getDate();
				let oHour = oDate.getHours();
				let oMin = oDate.getMinutes();
				let oSec = oDate.getSeconds();
				let oTime = oYear + '-' + getzf(oMonth) + '-' + getzf(oDay) + ' ' + getzf(oHour) + ':' + getzf(oMin) + ':' + getzf(
					oSec); //最后拼接时间
				return oTime
			}
		},
		created() {
			if (this.messageInfo.type == 0) {
				if (this.messageInfo.status == 1) {
					this.tip = "评论了你的文章";
				} else if (this.messageInfo.status == 2) {
					this.tip = "收藏了你的文章";
				} else if (this.messageInfo.status == 3) {
					this.tip = "转发了你的文章";
				} else if (this.messageInfo.status == 4) {
					this.tip = "赞了你的文章";
				}
			} else if (this.messageInfo.type == 1) {
				if (this.messageInfo.status == 1) {
					this.tip = "评论了你的想法";
				} else if (this.messageInfo.status == 0) {

				} else if (this.messageInfo.status == 2 || this.messageInfo.status == 3 || this.messageInfo.status == 4) {
					this.tip = "送了你一个";
				}
			}
		}
	}
</script>

<style scoped>
	/* 评论样式 */
	.comment {
		margin-left: 110rpx;
		font-size: 28rpx;
		color: #505050;
		margin-top: 10rpx;
		margin-bottom: 10rpx;
		width: 600rpx;
		overflow: hidden;
		-webkit-line-clamp: 1;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-box-orient: vertical;
	}

	/* 头部插槽动作 */
	.slot {
		margin-left: 15rpx;
		font-size: 28rpx;
		color: #505050;
		font-weight: 300;
	}

	/* 文章和想法的样式 */
	.thecontent {
		margin-left: 110rpx;
		width: 600rpx;
		height: 140rpx;
		background-color: #f1f1f1;
		border-radius: 8rpx;
		display: flex;
		padding: 20rpx 20rpx;
		margin-bottom: 30rpx;
		margin-top: 6rpx;
	}

	.thewords {
		width: 1200rpx;
	}

	.theimg {
		width: 300rpx;
		height: 100rpx;
		border-radius: 4rpx;
		margin-left: 10rpx;
	}

	/* 公共头像组件样式 */
	>>>.attention-cell .flex-item {
		border: none;
		padding: 10rpx;
		margin-top: 17rpx;
	}
</style>
