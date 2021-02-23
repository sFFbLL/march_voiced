<template>
	<!-- 评论吸底框 -->
	<view class="commentwarp">
		<view class="comment-view">
			<textarea auto-height
			 maxlength="-1"
			 class="commentinput"
			 @input="onkeyinput"
			 placeholder="请输入评论"></textarea>
		</view>
		<button class="sendbtn"
		 size="mini"
		 :disabled="sendDisabled"
		 @click="sendComment()">发布</button>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				sendDisabled: true,
				comment: null,
			}
		},
		methods: {
			onkeyinput(event) {
				this.comment = event.target.value;
				this.sendDisabled = false;
			},
			sendComment() {
				// 发送评论接口
				// publishComment(comment).then(res => {
				// console.log("success send comment");
				this.$emit('sendComment');
				// })
				// 重新获取评论？
			}
		}
	}
</script>

<style scoped>
	.commentwarp{
		background-color: #fff;
	}
	/* 评论 */
	.commentinput {
		line-height: 42rpx;
	}

	.sendbtn::before {}

	.sendbtn {
		height: 68rpx;
		width: 140rpx;
		background-color: #FFFFFF;
		color: #007AFF;
		font-size: 28rpx;
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
		/*定义body的元素垂直居中*/
		width: 100%;
		font-size: 28rpx;
		padding: 15rpx 28rpx;
		text-align: left;
		position: fixed;
		bottom: 0;
		border-top: 1rpx solid #d1d6de;
		color: #303133;
	}
</style>
