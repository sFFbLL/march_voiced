<template>
	<view>
		<view class="comment" v-for="(res, pindex) in commentList">
			<view class="left">
				<image :src="'http://linbolun.cn/api/file/download/' + res.idAvatar" mode="aspectFill"></image>
			</view>
			<view class="right">
				<view class="top">
					<view class="name">{{ res.createByName }}</view>
				</view>
				<text v-if="res.replyName" class="reply-name">@{{res.replyName}}</text>
				<view class="content" @click="addComment(res.createBy,res.id,pindex,res.createByName)">{{ res.content }}</view>
				<view class="bottom">
					
					{{ format(res.createTime)}}
					
				</view>

				<view class="reply-box" v-if="showKids">
					<!-- 子评论 -->
					<view v-for="(item, index) in res.ChildComments">
						<view class="item">
							<view class="left">
								<image :src="'http://linbolun.cn/api/file/download/' + res.idAvatar" mode="aspectFill"></image>
							</view>
							<view class="right" @click="addComment(item.replyId,res.id,pindex,item.replyName)">
								<view class="name">{{ item.createByName }}</view>
								<text class="reply-name">@{{item.replyName}}</text>
								<text class="kid-content">{{ item.content }}</text>
								<text class="bottom">
									{{ format(item.createTime)}}
								</text>
							</view>
						</view>
					</view>
					<!-- 获取初次传过来的子评论数量，固定只显示3条子评论，如果长度大于等于3展示更多回复，如果小于3不展示 -->

					<view v-if="res.ChildComments&&res.ChildComments.length>=3">
						<view class="all-reply" @tap="toAllReply(pindex)">
							{{text}}
						</view>
					</view>

				</view>

			</view>

		</view>
	</view>
</template>

<script>
	import moment from 'moment';
	export default {
		data() {
			return {
				text: "展开更多回复>"
			};
		},
		props: {
			id: { //文章id
				type: Number,
				default: null
			},
			commentList: { //评论的列表
				type: Array,
				default: null
			},
			type: { //是从那个页面调用的，想法详情是1，文章详情是0，更多评论详情是3
				type: Number,
				default: null
			},
			showKids: { //是否展示子评论列表
				type: Boolean,
				default: true
			}
		},

		methods: {
			// 把时间戳转换为正确格式
			format(dateTime) {
				let stamp = new Date(dateTime);
				let time = moment(stamp).format('YYYY-MM-DD HH:mm:ss');
				return time;
			},
			// 跳转到全部回复
			toAllReply(pindex) {
				let parentComment = encodeURIComponent(JSON.stringify(this.commentList[pindex]));
				uni.navigateTo({
					url: '../../pages/moreComments/moreComments?parentComment=' + parentComment + '&id=' + this.id+'&type='+this.type
				})
			},
			// 添加评论（回复人id，父评论id，数据的index，被回复人姓名
			addComment(replyId, follewId, index, replyName) {
				let data = {
					childComment: true,
					index: index,
					replyId: replyId,
					follewId: follewId,
					replyName: replyName,
				}
				this.$emit('childFn', data);
			}
		}
	};
</script>

<style lang="scss" scoped>
	.comment:last-child {
		border: none;
	}

	.commentq {
		position: absolute;
		left: 650rpx;
	}

	.chat {
		transform: rotateY(180deg);
	}

	.name {
		color: #808080;
		font-size: 26rpx;
		line-height: 150%;
	}

	.comment {
		display: flex;
		padding: 30rpx 0;
		border-bottom: 1rpx solid #f0f0f0;

		.left {
			image {
				width: 70rpx;
				height: 70rpx;
				border-radius: 50%;
			}
		}

		.right {
			flex: 1;
			padding-left: 20rpx;
			font-size: 30rpx;

			.top {
				display: flex;
				justify-content: space-between;
				align-items: center;

				.like {
					display: flex;
					align-items: center;
					color: #9a9a9a;
					font-size: 26rpx;
				}

				.highlight {
					color: #5677fc;
				}
			}

			.content,
			.kid-content {
				margin-bottom: 10rpx;
				color: #505050;
				font-size: 28rpx;
				line-height: 150%;
				text-align: left;
			}

			.reply-box {
				.item {
					display: flex;
					padding: 20rpx 20rpx 0 20rpx;
				}

				.all-reply {
					padding: 20rpx;
					color: #2a82e4;
					font-size: 28rpx;
					line-height: 150%;
					text-align: left;
					font-weight: bold;

					.more {
						margin-left: 6rpx;
					}
				}
			}

			.bottom {
				margin-top: 20rpx;
				display: flex;
				font-size: 24rpx;
				color: #9a9a9a;

				.reply {
					color: #5677fc;
					margin-left: 10rpx;
				}
			}
		}
	}

	.reply-name {
		color: #2a82e4;
		font-size: 28rpx;
		line-height: 150%;
	}

	.item .left {
		display: flex;
	}

	.item .left image {
		flex: 1;
		width: 50rpx;
		height: 50rpx;
		line-height: 80rpx;
	}
</style>
