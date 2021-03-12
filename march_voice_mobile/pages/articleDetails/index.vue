<template>
	<view class="article-details">
		<view class="view-content">
			<view class="article-title">{{articleInfo.title}}</view>
			<attentionAndFansCell :nickname="articleInfo.nickname" :avatarPath="articleInfo.avatarPath" :isFollow="articleInfo.isFollow"
			 class="user-info"></attentionAndFansCell>
			<view class="article-message">
				<text class="word-num">
					<text>字数</text>
					<text class="word-num-num">{{articleInfo.word_count}}</text>
				</text>
				<text class="time">
					<text>{{articleInfo.update_time}}</text>
				</text>
			</view>
			<view class="article-content">
				<jinEdit :html="articleInfo.content" :readOnly="true"></jinEdit>
			</view>
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

		<!-- 底部导航 -->
		<view class="bottom-nav" v-if="!isComment">
			<view class="bottom-nav-centre">
				<!-- 评论图标 -->
				<view class="attention-icon comment">
					<view class="attention-item" @click="comment()">
						<u-icon name="pinglun" custom-prefix="custom-icon"></u-icon>
						<text class="attention-num">{{articleInfo.commentTotal}}</text>
					</view>
				</view>
				<!-- 喜欢图标 -->
				<view class="attention-icon like">
					<view class="attention-item">
						<u-icon name="heart"></u-icon>
						<text class="attention-num">{{articleInfo.favourTotal}}</text>
					</view>
				</view>
				<!-- 收藏图标 -->
				<view class="attention-icon collect">
					<view class="attention-item">
						<u-icon name="shoucang1" custom-prefix="custom-icon"></u-icon>
						<text class="attention-num">{{articleInfo.collectTotal}}</text>
					</view>
				</view>
				<!-- 转发 -->
				<view class="attention-icon">
					<u-icon name="zhuanfa"></u-icon>
				</view>
			</view>
		</view>
		<commentInput v-on:sendComment='sendComment' v-if="isComment" :type="type" />
	</view>
</template>

<script>
	import articleTitle from "../../marchVoiceComponents/showArticle/childComponents/articleTitle.vue"
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import jinEdit from '../../components/jin-edit/jin-edit.vue';
	import comment from '../../marchVoiceComponents/comment/index.vue'
	import commentInput from '../../marchVoiceComponents/comment/commentInput.vue'
	import {
		getArtileDetails,
		getArticleCommentList,
		addArticleComment,
		getChildCommentList,
		favour,
		collect,
		reprint
	} from '../../utils/api/articleDetails-api.js'
	export default {
		data() {
			return {
				id: "",
				current: 1,
				size: 5,
				childSize: 3,
				articleInfo: {},
				commentCount: 0,
				commentList: [],
				isComment: false,
				type:0
			}
		},
		components: {
			articleTitle,
			attentionAndFansCell,
			jinEdit,
			comment,
			commentInput
		},
		onLoad: function(option) { //option为object类型，会序列化上个页面传递的参数
			this.id = option //打印出上个页面传递的参数。
			// 获取文章详情内容
			getArtileDetails(option).then(res => {
				if (res.code === 0) {
					this.articleInfo = res.data;
				}
			}).catch(err => {
				console.log(err, "err login")
			})

			// 获取评论列表
			let params = {
				id: option,
				current: this.current,
				size: this.size,
				childSize: this.childSize,
			}
			getArticleCommentList(params).then(res => {
				if (res.code === 0) {
					this.commentList = res.data;
					this.commentCount = res.data.length
				}

			})
		},
		created() {

		},
		methods: {
			// 评论
			comment() {
				this.isComment = true;
			},
			// 发布评论
			sendComment() {
				console.log("发布评论");
				this.isComment = false;
			}
		},
		mounted() {
			let commentCount = 0;
			this.commentList.forEach((comment) => {
				commentCount++;
				comment.commentKids.forEach(count => {
					commentCount++;
				})
			})
			this.commentCount = commentCount;
		}
	}
</script>

<style scoped>
	.article-details {
		background-color: #f7f7f7;
	}

	.view-content {
		padding: 0 30rpx;
		background-color: #fff;
	}

	.article-title {
		color: #404040;
		font-size: 36rpx;
		line-height: 150%;
		font-weight: bold;
	}

	>>>.attention-cell .flex-item {
		padding: 20rpx 0;
		border: none;
	}

	>>>.attention-cell .left-img-box {
		margin-right: 0;
	}

	>>>.attention-cell .right-button {
		background-color: #fff;
	}

	.article-message {
		color: #999797;
		font-size: 24rpx;
	}

	.word-num-num {
		margin-left: 4rpx;
	}

	.time {
		float: right;
	}

	.article-content {
		padding: 0 10rpx;
	}

	>>>.article-content .ql-container {
		margin: 0;
		width: 100%;
		color: #404040;
		font-size: 32rpx;
		line-height: 150%;
	}

	>>>.article-content .container {
		padding-top: 20rpx;
	}

	>>>.article-content .ql-container .ql-editor {
		text-align: justify;
	}

	>>>.article-content .tool-view {
		display: none;
	}

	/* 底部导航:点赞，评论等 */
	.bottom-nav {
		position: fixed;
		bottom: 0;
		width: 100%;
		height: 120rpx;
		border: 1rpx solid #f0f0f0;
		border-width: 1rpx 0;
		background-color: #fff;
	}

	.bottom-nav-centre {
		display: flex;
		margin: 0 50rpx;
		height: 100%;
	}

	.attention-icon::before {
		content: '';
		display: inline-block;
		vertical-align: middle;
		height: 100%;
	}

	.attention-icon {
		flex: 1;
		text-align: center;
	}

	>>>.attention-icon .u-icon {
		font-size: 50rpx !important;
		color: #404040 !important;
	}


	.attention-item {
		position: relative;
		display: inline-block;
	}

	.attention-icon .attention-num {
		position: absolute;
		top: -8rpx;
	}

	/* 评论样式 */
	.top-title {
		height: 88rpx;
		color: #505050;
		font-size: 34rpx;
		line-height: 150%;
		font-weight: bold;
	}

	.top-title text:first-child {
		display: inline-block;
		margin: 0 10rpx 0 30rpx;
		line-height: 88rpx;
	}

	.comment-view {
		margin-bottom: 120rpx;
	}

	.comment-list {
		padding: 0 30rpx;
		background-color: #fff;
	}
</style>
