<template>
	<view class="article-details">
		<view class="view-content">
			<view class="article-title">{{articleInfo.title}}</view>
			<attentionAndFansCell :nickname="articleInfo.nickname"
			 :avatarPath="articleInfo.avatarPath"
			 :isFollow="articleInfo.isFollow"
			 class="user-info"></attentionAndFansCell>
			<view class="article-message">
				<text class="word-num">
					<text>字数</text>
					<text class="word-num-num">123</text>
				</text>
				<text class="time">
					<text>{{articleInfo.updateTime}}</text>
				</text>
			</view>
			<view class="article-content">
				<jinEdit :html="articleInfo.content"
				 :readOnly="true"></jinEdit>
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
		<view class="bottom-nav">
			<view class="bottom-nav-centre">
				<!-- 评论图标 -->
				<view class="attention-icon comment">
					<view class="attention-item"
					 @click="comment()">
						<u-icon name="pinglun"
						 custom-prefix="custom-icon"></u-icon>
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
						<u-icon name="shoucang1"
						 custom-prefix="custom-icon"></u-icon>
						<text class="attention-num">{{articleInfo.collectTotal}}</text>
					</view>
				</view>
				<!-- 转发 -->
				<view class="attention-icon">
					<u-icon name="zhuanfa"></u-icon>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import articleTitle from "../../marchVoiceComponents/showArticle/childComponents/articleTitle.vue"
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import jinEdit from '../../components/jin-edit/jin-edit.vue';
	import comment from '../../marchVoiceComponents/comment/index.vue'
	export default {
		data() {
			return {
				id: "",
				articleInfo: {
					"title": "我们始终没有认清自己，不知道内心真正想要的，更不懂坚持自律的意义是什么",
					"content": "<span><p>现在很多APP在用户截图时会自动提示用户是否要进行分享。这个时机确实抓的很棒，一般来说，用户截图大多数时候都是为了分享给他人，少部分是为了留底备份。</p><p>用户分享内容到社交媒体或好友，不应该是一种粗暴的强制行为，我们应该在保证产品本身内容有吸引力的核心前提下，仔细揣摩用户心理，结合产品本身的特色，在不同情境下提供给用户最合适的分享平台及方式，让用户分享成为一种水到渠成的自然行为，甚至在某些时候还能给用户带来一些小的惊喜就更棒了。</p><p>尤其是刚上线的产品，很难通过产品的内部体系来实现快速的用户增长，所以会更加依赖于分享来达到广泛的传播，获取目标用户。</p><p>在常用的社交分享组件中，微信...</p></span>",
					"kind": 1,
					"tag": 1,
					"type": 1,
					"updateTime": "更新时间",
					"favourTotal": 1,
					"collectTotal": 1,
					"commentTotal": 1,
					"user_id": 1,
					"nickname": "郝庆会",
					"avatarPath": require('static/img/2.jpg'),
					"isFollow": 1
				},
				commentCount: 0,
				commentList: [{
					"id": 1,
					"content": "这里没有评论",
					"replyName": null,
					"replyId": 0,
					"idAvatar": require('static/img/2.jpg'),
					"createBy": 1,
					"createTime": "1.24 23:01",
					"createByName": "张三",
					"commentKids": [{
						"id": 2,
						"replyId": 0,
						"replyName": null,
						"createBy": 2,
						"createByName": "李四",
						"idAvatar": require('static/img/1.jpg'),
						"content": "针不戳",
						"createTime": "创建时间",
						"commentKids": []
					}, {
						"id": 3,
						"replyId": 2,
						"replyName": "李四",
						"createBy": 1,
						"createByName": "张三",
						"idAvatar": require('static/img/2.jpg'),
						"content": "不戳啥不戳",
						"createTime": "创建时间",
						"commentKids": []
					}]
				}, {
					"id": 4,
					"content": "看文章，别说话",
					"replyName": null,
					"replyId": 0,
					"idAvatar": require('static/img/图片 8.jpg'),
					"createBy": 4,
					"createTime": "1.24 23:01",
					"createByName": "王麻子",
					"commentKids": [{
							"id": 5,
							"replyId": 0,
							"replyName": null,
							"createBy": 2,
							"createByName": "李四",
							"idAvatar": require('static/img/1.jpg'),
							"content": "王麻子nb",
							"createTime": "创建时间",
							"commentKids": []
						}, {
							"id": 6,
							"replyId": 0,
							"replyName": null,
							"createBy": 1,
							"createByName": "张三",
							"idAvatar": require('static/img/2.jpg'),
							"content": "王麻子nb +10086",
							"createTime": "创建时间",
							"commentKids": []
						}, {
							"id": 7,
							"replyId": 1,
							"replyName": "张三",
							"createBy": 4,
							"createByName": "王麻子",
							"idAvatar": require('static/img/图片 8.jpg'),
							"content": "你可消停会吧！！！",
							"createTime": "1.24 23:01",
							"commentKids": []
						},
						{
							"id": 8,
							"replyId": 4,
							"replyName": "王麻子",
							"createBy": 1,
							"createByName": "张三",
							"idAvatar": require('static/img/2.jpg'),
							"content": "我就不",
							"createTime": "1.24 23:01",
							"commentKids": []
						}
					]
				}]
			};
		},
		components: {
			articleTitle,
			attentionAndFansCell,
			jinEdit,
			comment
		},
		onLoad: function (option) { //option为object类型，会序列化上个页面传递的参数
			this.id = option //打印出上个页面传递的参数。
		},
		methods: {
			comment() {

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
