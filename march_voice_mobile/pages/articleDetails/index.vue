<template>
	<view class="article-details">
		<view class="view-content">
			<view class="article-title">{{articleInfo1.title}}</view>
			<attentionAndFansCell :nickname="articleInfo1.nickname"
			 :avatarPath="articleInfo1.avatarPath"
			 :isFollow="articleInfo1.isFollow"
			 class="user-info"></attentionAndFansCell>
			<view class="article-message">
				<text class="word-num">
					<text>字数</text>
					<text class="word-num-num">123</text>
				</text>
				<text class="time">
					<text>{{articleInfo1.updateTime}}</text>
				</text>
			</view>
			<view class="article-content">
				<jinEdit :html="articleInfo1.content"
				 :readOnly="true"
				 uploadFileUrl="http://www.kuntong.site/api/file/uploadImage"></jinEdit>
			</view>
		</view>
		<view class="bottom-nav">
			<view class="bottom-nav-centre">
				<view class="attention-icon comment">
					<uni-icons type="chat"></uni-icons>
				</view>
				<view class="attention-icon like">
					<u-icon name="heart"></u-icon>
				</view>
				<view class="attention-icon">
					<uni-icons type="star"></uni-icons>
				</view>
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
	export default {
		data() {
			return {
				id: "",
				articleInfo: {

				},
				articleInfo1: {
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
				}
			};
		},
		components: {
			articleTitle,
			attentionAndFansCell,
			jinEdit
		},
		onLoad: function (option) { //option为object类型，会序列化上个页面传递的参数
			this.id = option //打印出上个页面传递的参数。
		},
		methods: {

		},
		mounted() {

		}
	}
</script>

<style scoped>
	.view-content {
		margin: 0 30rpx;
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
		height: 100rpx;
		border: 1rpx solid #f0f0f0;
		border-width: 1rpx 0;
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
</style>
