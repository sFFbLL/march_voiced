<template>
	<view>
		<!-- 间隔槽 -->
		<u-gap height="20" bg-color="#f5f5f5"></u-gap>
		<view class="ideacontent">

			<!-- 用户头像公共组件 -->
			<attentionAndFansCell :nickname="ideaInfoList.nickname" :avatarPath="ideaInfoList.avatarPath" :isFollow="ideaInfoList.isFollow"></attentionAndFansCell>
			<!-- 想法的文字部分 -->
			<articleContent :articleContent="ideaInfoList.content"></articleContent>
			<!-- 想法的图片部分组件 -->
			<imageAdaptation :imgList="ideaInfoList.imgList"></imageAdaptation>

			<!-- 点赞表情组件 -->
			<emojiControl :isShow="false" :faceTotals="ideaInfoList.faceTotal" :likeTotals="ideaInfoList.likeTotal" :favourTotals="ideaInfoList.favourTotal"
					 :commentTotals="ideaInfoList.commentTotal" :id="ideaInfoList.id" ></emojiControl>
		</view>
		<!-- 评论吸底框 -->
		<view class="commentwarp">
			<input class="commentinput" confirm-type="send" @input="onkeyinput" placeholder="请输入评论" type="text" maxlength="-1" />
			<button class="sendbtn" size="mini" :disabled="sendDisabled" @click="sendComment()">发布</button>
		</view>
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
	export default {
		components: {
			emojiControl,
			attentionAndFansCell,
			imageAdaptation,
			articleContent
		},
		data() {
			return {
				sendDisabled: true,
				comment: null,
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
			ideaDetail(id).then(res=>{
				this.ideaInfoList=res.data;
			})


		},

		methods: {
			onkeyinput(event) {
				this.comment = event.target.value;
				this.sendDisabled = false;
			},
			sendComment() {
				// 发送评论接口
				// publishComment(comment).then(res => {
				console.log("success send comment");
				// })
				// 重新获取评论？
			}
		}
	}
</script>

<style>
	/* 评论 */
	.commentinput {
		width: 81%;
	}

	.sendbtn {
		background-color: #FFFFFF;
		color: #007AFF;
		font-size: 28rpx;
	}

	>>>.sendbtn::after {
		border: none;
	}

	.commentwarp {
		display: flex;
		width: 100%;
		font-size: 28rpx;
		padding: 15rpx 28rpx;
		text-align: left;
		position: fixed;
		bottom: 0;
		border-top: 1rpx solid #d1d6de;
		color: #303133;
	}

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
