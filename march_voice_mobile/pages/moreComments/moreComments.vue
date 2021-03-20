<template>
	<view>
		<view class="comment-list">
			<!-- 父评论 -->
			<comment :showKids="showKids" :type="type" :commentList="commentList" @childFn="comment"></comment>
			<!-- 所有子评论 -->
			<comment :showKids="showKids" :type="type" :commentList="ChildCommentsList" @childFn="comment"></comment>
			<!-- 评论输入组件 -->
			<commentInput :show="showAddComment" v-if="isComment" @addComment="addComment" @addChildComment="addChildComment"
			 @childFn="parentFn" :type="type" :addCommentArg="addCommentArg" />
		</view>
	</view>
</template>

<script>
	import comment from '../../marchVoiceComponents/comment/index.vue'
	import {
		getChildCommentList,
		addArticleComment
	} from '../../utils/api/articleDetails-api.js'
	export default {
		data() {
			return {
				type: 3,
				commentList: [],
				showKids: false,
				id: 0,
				ChildCommentsList: [],
				artilceId: 0,
				showAddComment:true,
				isComment:true
			}
		},
		onLoad(option) {
			let data = JSON.stringify(option);

			console.log(data);
			this.artilceId = data.id;
			this.commentList = data.commentList;
			this.id = this.commentList.id;
			this.getMore();
			this.addCommentArg = {
				id: this.artilceId,
				replyId: 0,
				follewId: 0,
				childComment: false,
			}
		},
		methods: {
			// 查询这个评论的所有子评论
			getMore() {
				let params = {
					id: this.id
				}
				getChildCommentList(params).then(res => {
					this.ChildCommentsList = res.data;
				})
			},
			// 调用新增评论接口
			newComment(params) {
				// 文章评论发布接口
				addArticleComment(params).then(res => {});
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
					// id: newid
				}
				this.commentList.unshift(newcomment);
				this.isComment = false;
				this.commentCount++;
				// 把参数传给接口
				let params = {
					id: this.addCommentArg.id,
					content: payload,
					replyId: this.addCommentArg.replyId,
					followId: this.addCommentArg.follewId,
				}
				console.log(params)
				this.newComment(params);
			},
		}
	}
</script>

<style>

</style>
