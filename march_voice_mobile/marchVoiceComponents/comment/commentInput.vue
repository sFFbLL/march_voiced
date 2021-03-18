<template>
	<view>
		<!-- 评论吸底框 -->
		<view class="commentwarp" v-if="!show">
			<!-- 想法评论显示 -->
			<view class="comment">
				<view class="comment-view">
					<textarea auto-height maxlength="-1" class="commentinput" @input="onkeyinput" placeholder="请输入评论"></textarea>
				</view>
				<button class="sendbtn" size="mini" :disabled="sendDisabled" @click="sendComment()">发布</button>
			</view>
		</view>
		<!-- 文章评论显示 -->
		<u-popup v-model="show" mode="bottom" @close="close" border-radius="10">
			<view class="comment">
				<view class="comment-view">
					<textarea auto-height maxlength="-1" class="commentinput" @input="onkeyinput" placeholder="请输入评论"></textarea>
				</view>
				<button class="sendbtn" size="mini" :disabled="sendDisabled" @click="sendComment()">发布</button>
			</view>
		</u-popup>
	</view>
</template>

<script>
	import {
		publishIdea
	} from '../../utils/api/marchCircle-api.js'
	import {
		addArticleComment
	} from '../../utils/api/articleDetails-api.js'
	import {
		information
	} from '../../utils/api/personInfo-api.js'
	export default {
		data() {
			return {
				sendDisabled: true,
				comment: null,
			}
		},

		props: {
			type: {
				type: Number,
				default: null
			},
			addCommentArg: {
				type: Object
			},
			show: {
				type: Boolean
			}

		},
		methods: {
			// 获取输入内容
			onkeyinput(event) {
				this.comment = event.target.value;
				this.sendDisabled = false;
			},
			// 关闭评论
			close() {
				this.$emit('childFn');
			},
			// 新增评论
			sendComment() {
				let that = this;
				let params = {
					id: this.addCommentArg.id,
					content: this.comment,
					replyId: this.addCommentArg.replyId,
					followId: this.addCommentArg.follewId,
				}
				if (this.type === 1) {
					// 想法发送评论接口
					publishIdea(params).then(res => {
						this.close();
						return res.code
					}).then(res => {
						// 后端添加数据成功
						// 调用查询用户接口
						let user;
						information().then(res => {
							user = res.data;
						})
						// 插入一条新的评论
						// 判断是对文章评论
						if (!this.addCommentArg.childComment) {
							// 把数据传给父组件显示到页面
							let newcomment = {
								createByName: user.nickname,
								idAvatar: user.avatarPath,
								content: this.comment,
								createTime: new Date(),
								commentKids: []
							}
							this.$emit('addComment', newcomment);
						} else {
							// 判断是对评论评论
							let newcomment = {
								createByName: user.nickname,
								idAvatar: user.avatarPath,
								content: this.comment,
								createTime: new Date(),
								replyName: this.addCommentArg.replyName,
								index: this.addCommentArg.index,
								commentKids: []
							}
							that.$emit('addChildComment', newcomment);
						}
					})

				} else if (this.type === 0) {
					let _this = this;
					// 文章评论发布接口
					addArticleComment(params).then(res => {
						this.close();
						return res.code
					}).then(res => {
						// 后端添加数据成功
						// 调用查询用户接口
						information().then(res => {
							return res.data
						}).then(res => {
							_this.addComment(res);
							console.log(127127)
						})

					})
				}


			},
			// 插入新的评论在页面
			addComment(res) {
				let that =this;
				// 插入一条新的评论
				// 判断是对文章评论
				if (!this.addCommentArg.childComment) {
					console.log(131)
					// 把数据传给父组件显示到页面
					let newcomment = {
						createByName: res.nickname,
						idAvatar: res.avatarPath,
						content: this.comment,
						createTime: new Date(),
						commentKids: []
					}
					console.log(newcomment)
					that.$emit('addComment', newcomment);
					console.log(152)
				} else {
					// 判断是对评论评论
					let newcomment = {
						createByName: user.nickname,
						idAvatar: user.avatarPath,
						content: this.comment,
						createTime: new Date(),
						replyName: this.addCommentArg.replyName,
						index: this.addCommentArg.index,
						commentKids: []
					}
					console.log(newcomment)
					this.$emit('addChildComment', newcomment);
				}
			}
		}
	}
</script>

<style scoped>
	.comment {
		display: flex;
		margin-top: 30rpx;
	}

	/* 评论 */
	.commentinput {
		line-height: 62rpx;
		width: 550rpx;
	}

	.sendbtn {
		height: 68rpx;
		width: 140rpx;
		background-color: #FFFFFF;
		color: #007AFF;
		font-size: 28rpx;
		margin-bottom: 30rpx;
	}

	>>>.sendbtn::after {
		border: none;
	}

	.comment-view {
		max-height: 126rpx;
		overflow: scroll;

	}

	.commentwarp {
		display: flex;
		align-items: center;
		width: 100%;
		font-size: 28rpx;
		padding: 15rpx 28rpx;
		text-align: left;
		position: fixed;
		bottom: 0;
		color: #303133;
		border-top: 2rpx solid #efefef;
		border-radius: 20rpx;
		background-color: #fff;
	}
</style>
