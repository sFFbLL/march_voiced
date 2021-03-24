<template>
	<view>
		<!-- 间隔槽 -->
		<u-gap height="20" bg-color="#f5f5f5"></u-gap>
		<view class="ideacontent" v-if="ideaInfoList.id">
			<!-- 用户头像公共组件 -->
			<attentionAndFansCell :aid="ideaInfoList.id" :nickname="ideaInfoList.nickname" :avatarPath="ideaInfoList.avatarPath"
			 :isFollow="ideaInfoList.isFollow"></attentionAndFansCell>
			<!-- 想法的文字部分 -->
			<jinEdit :html="ideaInfoList.content" :readOnly="true"></jinEdit>
			<!-- <articleContent :articleContent="ideaInfoList.content"></articleContent> -->
			<!-- 想法的图片部分组件 -->
			<imageAdaptation :imgList="ideaInfoList.imageList"></imageAdaptation>

			<!-- 点赞表情组件 -->
			<emojiControl :isShow="false" :faceTotals="ideaInfoList.faceTotal" :likeTotals="ideaInfoList.likeTotal"
			 :favourTotals="ideaInfoList.favourTotal" :type="ideaInfoList.type" :commentTotals="ideaInfoList.commentTotal" :id="ideaInfoList.id"></emojiControl>
		</view>
		<!-- 评论列表 -->
		<view class="comment-view">
			<view class="top-title">
				<text>评论</text>
				<text>({{commentCount}})</text>
			</view>
			<view v-if="commentCount === 0 || commentCount === 'NaN'">
				<u-empty text="没有数据"
				 mode="search"
				 class="nodate"></u-empty>
			</view>
			<view class="comment-list">
				<comment :id="ideaId" :type="type" @getMore="getMore" :commentList="commentList" @childFn="comment"></comment>
			</view>
		</view>

		<commentInput :show="showAddComment" v-if="isComment" @addComment="addComment" @addChildComment="addChildComment"
		 @childFn="parentFn" :type="type" :addCommentArg="addCommentArg" />
		<!-- 下拉加载更多 -->
		<view v-if="isLoadMore">
			<uni-load-more class="loading" :status="loadStatus" iconType="circle"></uni-load-more>
		</view>

	</view>
</template>

<script>
	import {
		ideaDetail,
		publishComment,
		ideaCommentList,
		ideaChildCommentList
	} from '../../utils/api/marchCircle-api.js'
	import jinEdit from '../../components/jin-edit/jin-edit.vue';
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import emojiControl from '../../marchVoiceComponents/marchCircle/emojiControl.vue'
	import imageAdaptation from '../../marchVoiceComponents/marchCircle/imageAdaptation.vue'
	import articleContent from '../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue'
	import comment from '../../marchVoiceComponents/comment/index.vue'
	import commentInput from '../../marchVoiceComponents/comment/commentInput.vue'
	import {
		getUserName,
		getAvatarPath
	} from "../../utils/auth.js"
	export default {
		components: {
			emojiControl,
			attentionAndFansCell,
			imageAdaptation,
			articleContent,
			commentInput,
			comment,
			jinEdit
		},
		data() {
			return {
				isLoadMore: false, //是否加载中
				loadStatus: 'loading',
				ideaInfoList: {},
				commentList: [],
				commentCount: 0,
				current: 1,
				size: 5,
				childSize: 3,
				type: 1,
				addCommentArg: {},
				showAddComment: false,
				isComment: true,
				ideaId: 0,
			}
		},
		onLoad(option) {
			this.ideaId = Number(option.id);
			this.commentCount = Number(option.commentTotal);
		},
		created() {
			let id = this.ideaId;
			this.addCommentArg = {
				id: id,
				replyId: 0,
				follewId: 0,
				childComment: false,
			}
			this.getIdea();
			this.getComments();
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.current = 1;
			this.commentList = [];
			this.getIdea();
			this.getComments();
			setTimeout(function() {
				uni.stopPullDownRefresh();
			}, 2000);
		},
		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true
				this.current += 1;
				this.getComments()
			}
		},
		methods: {
			// 获取想法详情
			getIdea() {
				let id = this.ideaId;
				let _this = this;
				// 获取想法详细信息接口
				ideaDetail(id).then(res => {
					if (res.code === 0) {
						_this.ideaInfoList = res.data;
					}
				}).catch(err => {
					console.log(err, "err login")
				})
			},
			// 获取评论列表
			getComments() {
				// 获取评论列表
				let params = {
					id: this.ideaId,
					current: this.current,
					size: this.size,
					childSize: this.childSize,
				}
				ideaCommentList(params).then(res => {
					if (res.data.CommentSum) {
						if (this.current === 1) {
							this.loadStatus = "nomore";
							this.isLoadMore = false;
						} else {
							setTimeout(function() {
								this.loadStatus = "nomore";
								this.isLoadMore = false;
							}, 1000);
						}
						this.commentList = [...this.commentList, ...res.data.CommentSum];
					} else {
						this.loadStatus = 'nomore';
					}

				})
			},
			// 控制评论弹出框的显示开
			comment(payload) {
				let id = this.ideaId;
				this.showAddComment = true;
				// 如果是子组件点击进入评论，传入这条评论的参数
				if (payload) {
					this.addCommentArg = {
						childComment: payload.childComment,
						id: id,
						index: payload.index,
						replyId: payload.replyId,
						follewId: payload.follewId,
						replyName: payload.replyName,
					}
				}
			},
			// 控制评论弹出框的显示关
			parentFn(payload) {
				this.showAddComment = false;
			},

			// 调用新增评论接口
			newComment(params) {
				// 想法评论发布接口
				publishComment(params).then(res => {});
			},
			// 添加一条评论
			addComment(payload) {
				// 把数据添加到本地数组里
				let newcomment = {
					createByName: getUserName(),
					idAvatar: getAvatarPath(),
					content: payload,
					createTime: new Date(),
					ChildComments: [],
				}
				this.commentList.unshift(newcomment);
				this.showAddComment = false;
				this.commentCount++;
				let params = {
					id: this.addCommentArg.id,
					content: payload,
					replyId: this.addCommentArg.replyId,
					followId: this.addCommentArg.follewId,
				}
				this.newComment(params);
			},
			// 添加一条子评论
			addChildComment(payload) {
				let childs = this.commentList[payload.index];
				// 判断子评论是否为空
				if (!childs.ChildComments) {
					this.commentList[payload.index].ChildComments = [];
				}
				// 把数据加到子评论
				this.commentList[payload.index].ChildComments.push(payload);
				// 调用增加评论接口
				let params = {
					id: this.addCommentArg.id,
					content: payload.content,
					replyId: this.addCommentArg.replyId,
					followId: this.addCommentArg.follewId,
				}
				this.newComment(params);
				this.commentCount++;
				let id = this.ideaId;
				this.addCommentArg = {
					id: id,
					replyId: 0,
					follewId: 0,
					childComment: false,
				}
			},

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

	>>>.ql-container{
		min-height: 40rpx;
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

	.comment-list {
		margin-bottom: 60px;
	}
	
	.nodate {
		background-color: #fff;
		min-height: 800rpx;
	}
</style>
