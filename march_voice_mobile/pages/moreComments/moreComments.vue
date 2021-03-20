<template>
	<view>
		<view class="comment-list">
			<!-- 父评论 -->
			<comment :showKids="showKids" :type="type" :commentList="commentList" @childFn="comment"></comment>
			<!-- 间隔槽 -->
			<u-gap height="30" bg-color="#f5f5f5"></u-gap>
			<!-- 所有子评论 -->
			<comment :showKids="showKids" :type="type" :commentList="ChildCommentsList" @childFn="comment"></comment>
			<!-- 评论输入组件 -->
			<commentInput :show="showAddComment" v-if="isComment" @addComment="addComment" @addChildComment="addChildComment"
			 @childFn="parentFn" :type="type" :addCommentArg="addCommentArg" />
		</view>
		<!-- 下拉加载更多 -->
		<view v-show="isLoadMore">
			<uni-load-more class="loading" :status="loadStatus" iconType="circle"></uni-load-more>
		</view>
	</view>
</template>

<script>
	import comment from '../../marchVoiceComponents/comment/index.vue'
	import commentInput from '../../marchVoiceComponents/comment/commentInput.vue'
	import {
		ideaChildCommentList,
		publishComment
	} from '../../utils/api/marchCircle-api.js'
	import {
		getChildCommentList,
		addArticleComment
	} from '../../utils/api/articleDetails-api.js'
	export default {
		components: {

			commentInput,
			comment
		},
		data() {
			return {
				type: 3,
				isLoadMore: false, //是否加载中
				commentList: [],
				showKids: false,
				loadStatus: 'loading',
				id: 0,
				ChildCommentsList: [],
				artilceId: 0,
				showAddComment: false,
				isComment: true,
				addCommentArg: null,
				current: 1,
				kidsCount: 0,
				size: 50,
				parentComment: {},
				fromType: 0,
			}
		},
		onLoad(option) {
			// 获取文章id和父评论
			this.fromType = Number(option.type);
			this.artilceId = Number(option.id);
			this.parentComment = JSON.parse(decodeURIComponent(option.parentComment))

			this.kidsCount = this.parentComment.ChildComments ? this.parentComment.ChildComments.length : 0;

			this.id = this.parentComment.id;
			this.commentList.push(this.parentComment);

			this.getMore();
			this.addCommentArg = {
				id: this.artilceId,
				replyId: this.parentComment.createBy,
				follewId: 0,
				childComment: true,
				replyName: this.parentComment.createByName
			}
		},
		methods: {
			// 查询这个评论的所有子评论
			getMore() {
				let params = {
					id: this.id,
					current: this.current,
					size: this.size,
				}
				if (this.fromType == 0) {
					getChildCommentList(params).then(res => {
						if (res.data.Comment) {
							this.ChildCommentsList = [...this.ChildCommentsList, ...res.data.Comment];
							if (this.current * this.size >= this.kidsCount) {
								this.loadStatus = "nomore";
							} else
							if (this.current === 1) {
								this.isLoadMore = false;
								this.ChildCommentsList = [...this.ChildCommentsList, ...res.data.Comment];
							} else {
								setTimeout(function() {
									this.isLoadMore = false;
									this.ChildCommentsList = [...this.ChildCommentsList, ...res.data.Comment];
								}, 3000);
							}
						}
					})
				} else if (this.fromType == 1) {
					ideaChildCommentList(params).then(res => {
						if (res.data.Comment) {
							this.ChildCommentsList = [...this.ChildCommentsList, ...res.data.Comment];
							if (this.current * this.size >= this.kidsCount) {
								this.loadStatus = "nomore";
							} else
							if (this.current === 1) {
								this.isLoadMore = false;
								this.ChildCommentsList = [...this.ChildCommentsList, ...res.data.Comment];
							} else {
								setTimeout(function() {
									this.isLoadMore = false;
									this.ChildCommentsList = [...this.ChildCommentsList, ...res.data.Comment];
								}, 3000);
							}
						}
					})
				}

			},
			// 控制评论弹出框的显示开
			comment(payload) {
				let id = this.artilceId;
				this.showAddComment = true;
				// 如果是子组件点击进入评论，传入这条评论的参数
				if (payload) {
					this.addCommentArg = {
						childComment: payload.childComment,
						id: id,
						index: payload.index,
						replyId: payload.replyId,
						follewId: this.parentComment.id,
						replyName: payload.createByName,
					}
				}
			},
			// 控制评论弹出框的显示关
			parentFn(payload) {
				this.showAddComment = false;
			},
			// 调用新增评论接口
			newComment(params) {
				if(this.fromType==1){
					// 想法评论发布接口
					publishComment(params).then(res => {});
				}else if(this.fromType==0){
						// 文章评论发布接口
					addArticleComment(params).then(res => {});
				}
				
			},
			// 添加一条子评论
			addChildComment(payload) {
				// 把数据加到子评论
				console.log(payload)
				this.ChildCommentsList.unshift(payload);
				// 调用增加评论接口
				let params = {
					id: this.addCommentArg.id,
					content: payload.content,
					replyId: this.addCommentArg.replyId,
					followId: this.addCommentArg.follewId,
				}
				console.log(params)
				this.newComment(params);
				let id = this.ideaId;
			},
			onReachBottom() { //上拉触底函数
				if (!this.isLoadMore) { //此处判断，上锁，防止重复请求
					this.isLoadMore = true
					this.current += 1;
					this.getMore()
				}
			},


		}
	}
</script>

<style>

</style>
