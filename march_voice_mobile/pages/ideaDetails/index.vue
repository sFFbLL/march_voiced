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
		<!-- 评论列表 -->
		<view class="comment-view">
			<view class="top-title">
				<text>评论</text>
				<text>({{commentCount}})</text>
			</view>
			<view class="comment-list">
				<comment :commentList="commentList"></comment>
			</view>
		</view>
		<commentInput v-on:sendComment='sendComment' :type="type" />
		
		
	</view>
</template>

<script>
	import {
		ideaDetail,
		publishComment,
		ideaCommentList,
		ideaChildCommentList
	} from '../../utils/api/marchCircle-api.js'
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import emojiControl from '../../marchVoiceComponents/marchCircle/emojiControl.vue'
	import imageAdaptation from '../../marchVoiceComponents/marchCircle/imageAdaptation.vue'
	import articleContent from '../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue'
	import comment from '../../marchVoiceComponents/comment/index.vue'
	import commentInput from '../../marchVoiceComponents/comment/commentInput.vue'
	export default {
		components: {
			emojiControl,
			attentionAndFansCell,
			imageAdaptation,
			articleContent,
			commentInput,
			comment
		},
		data() {
			return {
				ideaInfoList: {},
				commentList:[],
				commentCount:3,
				current: 1,
				size: 5,
				childSize: 3,
				type:1
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

	
			// 获取评论列表
			let params = {
				id: id,
				current: this.current,
				size: this.size,
				childSize: this.childSize,
			}
			ideaCommentList(params).then(res => {
				if (res.code === 0) {
					this.commentList = res.data;
					this.commentCount = res.data.length
				}
			
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
