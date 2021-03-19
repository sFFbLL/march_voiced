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
	import {
		getUserName,
		getAvatarPath
	} from "../../utils/auth.js"
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
				// 插入一条新的评论
				// 判断是对文章评论
				if (!this.addCommentArg.childComment) {
					// 把数据传给父组件显示到页面
					let newcomment = {
						createByName: getUserName(),
						idAvatar: getAvatarPath(),
						content: this.comment,
						createTime: new Date(),
						commentKids: []
					}
					this.$emit("addComment", newcomment);
				} else {
					// 判断是对评论评论
					let newcomment = {
						createByName: getUserName(),
						idAvatar: getAvatarPath(),
						content: this.comment,
						createTime: new Date(),
						replyName: this.addCommentArg.replyName,
						index: this.addCommentArg.index,
						commentKids: []
					}
					this.$emit('addChildComment', newcomment);
				}
				this.close();


			},

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
