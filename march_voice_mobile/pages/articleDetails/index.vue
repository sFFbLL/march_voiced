<template>
	<view class="article-details">
		<view class="view-content">
			<view class="article-title">{{articleInfo.title}}</view>
			<attentionAndFansCell :aid="articleInfo.create_by" :nickname="articleInfo.nickname" :avatarPath="articleInfo.avatarPath"
			 :isFollow="articleInfo.isFollow" class="user-info"></attentionAndFansCell>
			<view class="article-message">
				<text class="word-num">
					<text>字数</text>
					<text class="word-num-num">{{articleInfo.word_count}}</text>
				</text>
				<text class="time">
					<text>{{format(articleInfo.update_time)}}</text>
				</text>
			</view>
			<view class="article-content">
				<text v-html="articleInfo.content"></text>
			</view>
		</view>

		<!-- 评论列表 -->
		<view class="comment-view">
			<view class="top-title">
				<text>评论</text>
				<text>({{commentCount}})</text>
			</view>
			<view class="comment-list">
				<comment :kidsCommentCount="kidsCommentCount" :commentList="commentList" @childFn="comment"></comment>
			</view>
		</view>

		<!-- 底部导航 -->
		<view class="bottom-nav">
			<view class="bottom-nav-centre">
				<!-- 评论图标 -->
				<view class="attention-icon comment">
					<view class="attention-item" @click="comment()">
						<u-icon name="pinglun" custom-prefix="custom-icon"></u-icon>
						<!-- <text class="attention-num">{{articleInfo.commentTotal}}</text> -->
					</view>
				</view>
				<!-- 喜欢图标 -->
				<view class="attention-icon like" @click="like()">
					<view class="attention-item">
						<u-icon :name="likeIcon" :color="likeColor"></u-icon>
						<text class="attention-num">{{articleInfo.favourTotal}}</text>
					</view>
				</view>
				<!-- 收藏图标 -->
				<view class="attention-icon collect" @click="collent()">
					<view class="attention-item">
						<u-icon :name="collentIcon" custom-prefix="custom-icon"></u-icon>
						<text class="attention-num">{{articleInfo.collectTotal}}</text>
					</view>
				</view>
				<!-- 转发 -->
				<view class="attention-icon" @click="transpond()">
					<u-icon name="zhuanfa"></u-icon>
				</view>
			</view>
		</view>
		<commentInput :show="showAddComment" v-if="isComment" @addComment="addComment" @addChildComment="addChildComment"
		 @childFn="parentFn" :type="type" :addCommentArg="addCommentArg" />
		<!-- 下拉加载更多 -->
		<view v-show="isLoadMore">
			<uni-load-more class="loading" :status="loadStatus" iconType="circle"></uni-load-more>
		</view>
	</view>
</template>

<script>
	import articleTitle from "../../marchVoiceComponents/showArticle/childComponents/articleTitle.vue"
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import jinEdit from '../../components/jin-edit/jin-edit.vue';
	import comment from '../../marchVoiceComponents/comment/index.vue'
	import moment from 'moment';
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
				isLoadMore: false, //是否加载中
				loadStatus: 'loading',
				showAddComment: true,
				id: 0,
				current: 1,
				size: 5,
				childSize: 3,
				articleInfo: {},
				commentCount: 0,
				commentList: [],
				isComment: false,
				type: 0,
				addCommentArg: {},
				clickChild: false,
				likeIcon: "heart",
				collentIcon: "shoucang1",
				likeColor: "black",
				clickLike: false,
				clickCollect: false,
				kidsCommentCount: [],
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
			this.id = Number(option.id) //打印出上个页面传递的参数。进行数据类型转换
			let id = this.id;
			// 获取文章详情内容
			getArtileDetails(id).then(res => {
				if (res.code === 0) {
					this.articleInfo = res.data;
				}
			}).catch(err => {
				console.log(err, "err login")
			})
			this.getComments()

		},
		created() {
			let id = this.id
			this.addCommentArg = {
				id: this.id,
				replyId: 0,
				follewId: 0,
				childComment: false,
			}
		},
		methods: {
			// 点击喜欢按钮
			like() {
				// 如果是第一次点击喜欢，增加数据
				if (!this.clickLike) {

					let id = this.id;
					this.likeIcon = "heart-fill";
					this.likeColor = "#d81e06";
					this.articleInfo.favourTotal++;
					// 调用喜欢接口
					favour(id).then(res => {

					})
				} else {
					this.articleInfo.favourTotal--;
					this.likeIcon = "heart";
					this.likeColor = "#080100";
				}
				this.clickLike = !this.clickLike;
			},
			// 点击收藏按钮
			collent() {
				if (!this.clickCollect) {
					let id = this.id;
					this.collentIcon = "shoucang2";
					this.articleInfo.collectTotal++;
					// 调用收藏接口
					collect(id).then(res => {

					})
				} else {
					this.collentIcon = "shoucang1";
					this.articleInfo.collectTotal--;
				}
				this.clickCollect = !this.clickCollect
			},
			transpond() {
				let id = this.id;
				reprint(id).then(res => {})
				// 提示转发成功
				uni.showLoading({
					title: '加载中'
				});
				setTimeout(function() {
					uni.hideLoading();
				}, 2000);
				setTimeout(function() {
					uni.showToast({
						title: '转发成功',
						duration: 2000
					});
				}, 2000);

			},
			// 控制评论弹出框的显示开
			comment(payload) {
				let id = this.id;
				this.isComment = true;
				// 如果是子组件点击进入评论，传入这条评论的参数
				if (payload) {
					this.addCommentArg = {
						childComment: payload.childComment,
						id: this.id,
						index: payload.index,
						replyId: payload.replyId,
						follewId: payload.follewId,
						replyName: payload.replyName,
					}
				}
			},
			// 控制评论弹出框的显示关
			parentFn(payload) {
				this.isComment = false;
			},
			// 添加一条评论
			addComment(payload) {
				console.log(payload)
				console.log(239239239)
				this.commentList.unshift(payload);
				this.isComment = false;
				this.commentCount++;
			},
			// 添加一条子评论
			addChildComment(payload) {
				this.commentList[payload.index].commentKids.push(payload);
				this.isComment = false;
				this.commentCount++;
			},
			// 把时间戳转换为正确格式
			format(dateTime) {
				let stamp = new Date(dateTime);
				let time = moment(stamp).format('YYYY-MM-DD HH:mm:ss');
				return time;
			},
			// 获取评论列表
			getComments() {
				// 获取评论列表
				let params = {
					id: this.id,
					current: this.current,
					size: this.size,
					childSize: this.childSize,
				}
				getArticleCommentList(params).then(res => {
					if (res.code === 0) {
						this.commentList = res.data.CommentSum;
						this.commentCount = res.data.CommentSum.length
						let comments = this.commentList
						// 获取子评论传来的的数量
						for (let kid of comments) {
							if (kid.ChildComments) {
								this.kidsCommentCount.push(kid.ChildComments.length)
							} else {
								this.kidsCommentCount.push(0)
							}
						}


					}

				})
			}
		},
		mounted() {
			let commentCount = this.articleInfo.commentTotal;
			this.commentList.forEach((comment) => {
				commentCount++;
				comment.commentKids.forEach(count => {
					commentCount++;
				})
			})
			this.commentCount = commentCount;
		},
		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore ) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true
				this.current += 1;
				this.getComments()

			}
		},
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
		padding: 20rpx 10rpx;

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
