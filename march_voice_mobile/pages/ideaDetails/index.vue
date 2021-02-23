<template>
	<view>
		<!-- 间隔槽 -->
		<u-gap height="20"
		 bg-color="#f5f5f5"></u-gap>
		<view class="ideacontent">

			<!-- 用户头像公共组件 -->
			<attentionAndFansCell :nickname="ideaInfoList.nickname"
			 :avatarPath="ideaInfoList.avatarPath"
			 :isFollow="ideaInfoList.isFollow"></attentionAndFansCell>
			<!-- 想法的文字部分 -->
			<articleContent :articleContent="ideaInfoList.content"></articleContent>
			<!-- 想法的图片部分组件 -->
			<imageAdaptation :imgList="ideaInfoList.imgList"></imageAdaptation>

			<!-- 点赞表情组件 -->
			<emojiControl :isShow="false"
			 :faceTotals="ideaInfoList.faceTotal"
			 :likeTotals="ideaInfoList.likeTotal"
			 :favourTotals="ideaInfoList.favourTotal"
			 :commentTotals="ideaInfoList.commentTotal"
			 :id="ideaInfoList.id"></emojiControl>
		</view>
		<commentInput v-on:sendComment='sendComment' />
	</view>
</template>

<script>
	import {
		ideaDetail,
		publishComment
	} from '../../utils/api/marchCircle-api.js'
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import emojiControl from '../../marchVoiceComponents/marchCircle/emojiControl.vue'
	import imageAdaptation from '../../marchVoiceComponents/marchCircle/imageAdaptation.vue'
	import articleContent from '../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue'
	import commentInput from '../../marchVoiceComponents/comment/commentInput.vue'
	export default {
		components: {
			emojiControl,
			attentionAndFansCell,
			imageAdaptation,
			articleContent,
			commentInput
		},
		data() {
			return {
				ideaInfoList: {
					// content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// updateTime: '2020/12/12',
					// faceTotal: 2,
					// likeTotal: 3,
					// favourTotal: 0,
					// commentTotal: 0,
					// ideaId: 0,
					// imgList: [
					// 	'../../static/img/cat.jpg', '../../static/img/cat.jpg'
					// ],
					// create_by: 0,
					// nickname: "xianer",
					// avatarPath: '../../static/img/cat.jpg',
					// isFollow: 0

				}
			}
		},
		onLoad(option) {
			this.ideaId = option.id;
		},
		created() {
			let id = this.ideaId

			// 获取想法详细信息接口
			ideaDetail(id).then(res => {
				this.ideaInfoList = res.data;
			})


		},

		methods: {
			sendComment() {
				console.log("发布评论");
			}
		}
	}
</script>

<style>
	/* 公共头像组件样式 */
	>>>.attention-cell .flex-item {
		border: none;
		padding: 0;
		margin-top: 17rpx;
	}

	/* 想法大模块 */
	.ideacontent {
		background-color: #FFFFFF;
		padding: 26rpx;
		line-height: 1.5;

	}

	/* 想法文字部分样式 */
	.wordscontent {
		margin-top: 20rpx;
	}
</style>
