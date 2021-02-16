<template>
	<view class="message">
		<view class="header">
			<tabs :attentionRead="attentionRead" :interactRead="interactRead" :otherRead="otherRead" class="tag-nav" :tabs='tablist' v-on:tabActive='tabActive' />
		</view>
		<view class="content">
			<view v-if="tabIndex==0" v-for="item in interactList">
				<singleMessage :messageInfo="item"></singleMessage>
			</view>

			<view v-if="tabIndex==1" v-for="item in attentionList">
				<attentionAndFansCell :showDteial="false">
					<template v-slot:doSomeThing>
						<text class="slot">关注了你
						</text>
					</template>
				</attentionAndFansCell>
			</view>
			<view v-if="tabIndex==2">
				<view>
					<view>

						<span class="inner-text">你投稿的文章审核
							<span class="inner-text">已通过。</span>
							<span class="inner-text">未通过！</span>
						</span>

						<span>
							<span class="inner-text">您已加入三月圈</span>
							<span class="inner-text">您加入三月圈请求被驳回</span>
						</span>



					</view>

				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import tabs from '../../marchVoiceComponents/tabCard.vue'
	import articleContent from "../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue"
	import singleMessage from "../../marchVoiceComponents/message/singleMessage.vue"
	
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import {
		interactList,
		attentionList,
		otherList,
		unreadMessage,
		readMessage
	} from "../../utils/api/message-api.js"
	export default {
		components: {
			tabs,
			articleContent,
			singleMessage,
			attentionAndFansCell
		},
		data() {
			return {
				isLoading: false, // 是否为加载中
				isDisabled: false, // 是否禁用按钮点击
				isAttention: false,
				tabIndex: 0,
				attentionRead: false,
				otherRead: false,
				interactRead:false,
				tablist: [{
						index: 0,
						value: '互动消息',
						isActive: true
					},
					{
						index: 1,
						value: '关注消息',
						isActive: false
					},
					{
						index: 2,
						value: '系统消息',
						isActive: false
					}
				],
				// 互动消息
				interactList: [{
						type: 0,
						status: 1,
						articleId: 1,
						commentId: "哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",
						image: require('static/img/2.jpg'),
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
						createTime: "更新时间",
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
						}

					}, {
						type: 1,
						status: 3,
						articleId: 1,
						commentId: "",
						image: require('static/img/2.jpg'),
						content: "哈哈哈哈在法国夫是德国人头地方杠。一件是言情大神匪",
						createTime: "更新时间",
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
						}

					}, {
						type: 0,
						status: 2,
						articleId: 1,
						commentId: "",
						image: require('static/img/2.jpg'),
						content: "哈哈哈哈在法国夫是德国人头地方杠。一件是言情大神匪",
						createTime: "更新时间",
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
						}

					}, {
						type: 0,
						status: 3,
						articleId: 1,
						commentId: "",
						image: require('static/img/2.jpg'),
						content: "哈哈哈哈在法国夫是德国人头地方杠。一件是言情大神匪",
						createTime: "更新时间",
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
						}

					}, {
						type: 1,
						status: 4,
						articleId: 1,
						commentId: "",
						image: require('static/img/2.jpg'),
						content: "哈哈哈哈在法国夫是德国人头地方杠。一件是言情大神匪",
						createTime: "更新时间",
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
						}

					}, {
						type: 1,
						status: 2,
						articleId: 1,
						commentId: "",
						image: require('static/img/2.jpg'),
						content: "哈哈哈哈在法国夫是德国人头地方杠。一件是言情大神匪",
						createTime: "更新时间",
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
						}

					},
					{
						type: 1,
						status: 1,
						articleId: 1,
						commentId: "哈哈哈哈在法国夫是德国人头地方杠",
						image: require('static/img/2.jpg'),
						content: "哈哈哈哈在法国夫是德国人头地方杠。一件是言情大神匪",
						createTime: "更新时间",
						title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
						}

					},
				],


				//关注消息
				attentionList: [{
						createTime: "",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
							isFollow:0
						}
					},
					{
						createTime: "",
						user: {
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",
							id: 1,
							isFollow:1
						}
					}
				],

				//其他消息
				otherList: [{
					type: 1,
					status: 0,
					createTime: "时间时间"

				}],

			}
		},
		onPullDownRefresh() {
			if (this.tabIndex == 0) {
				// this.interact();
				// 查询关注和其他是否有未读消息
				// unreadMessage(1).then(res=>{
				// 	if(res.data.count>0){
						this.attentionRead=false;
				// 	}
				// })
				// unreadMessage(2).then(res=>{
				// 	if(res.data.count>0){
						this.otherRead=false;
				// 	}
				// })
			} else if (this.tabIndex == 1) {
				this.attention();
			} else if (this.tabIndex == 2) {
				this.other();
			}
			setTimeout(function () {
			            uni.stopPullDownRefresh();
			        }, 1000);
		},
		created() {
			uni.hideTabBarRedDot({
				index: 3
			})
			// 互动消息已读
			let type = this.tabIndex
			// readMessage(type).then(res => {

			// 	// 消除小红点
			// })
			// 查询所有互动消息
			// interactList().then(res => {
			// 	this.interactList = res.data
			// })
			// 查询关注和其他是否有未读消息
			// unreadMessage(1).then(res=>{
			// 	if(res.data.count>0){
			// 		this.attentionRead=false;
			// 	}
			// })
			// unreadMessage(2).then(res=>{
			// 	if(res.data.count>0){
			// 		this.otherRead=false;
			// 	}
			// })
		},
		methods: {
			/* 切换选项卡选项 */
			tabActive(tabIndex) {
				this.tablist.map((value, index) => {
					value.isActive = tabIndex == index ? true : false;
				})
				if (tabIndex == 0) {
					this.interact();
				} else if (tabIndex == 1) {
					this.attention();
				} else {
					this.other();
				}
				this.tabIndex = tabIndex;
			},

			interact() {
				let type = this.tabIndex
				// readMessage(type).then(res => {
				// 	// 消除小红点
				// })
				// interactList().then(res => {
				// 	this.interactList = res.data
				// })
			},
			attention() {
				let type = this.tabIndex
				// readMessage(type).then(res => {
				// 	// 消除小红点
				// })
				// attentionList().then(res => {
				// 	this.attentionList = res.data
				// })
			},
			other() {
				let type = this.tabIndex
				// readMessage(type).then(res => {
				// 	// 消除小红点
				// })
				// otherList().then(res => {
				// 	this.otherList = res.data
				// })
			}
		},


	}
</script>

<style scoped>
	/* 关注插槽 */
	.slot{
		margin-left: 20rpx;
		font-weight: 400;
		}
	.header {
		height: 80rpx;
		line-height: 80rpx;
	}

	>>>.tab-card .head-nav .head-nav-bottom {
		width: 70%;
	}

	.flex-item {
		display: flex;
		flex-wrap: wrap;
		box-sizing: border-box;
		border-bottom: 6rpx solid #f0f0f0;
		padding: 20rpx 30rpx 20rpx 30rpx;
		/* margin-bottom: 20rpx; */
	}

	/* 左侧图片盒子样式 */
	.left-img-box {
		display: flex;
		align-items: center;
		/* flex: 1; */
		margin-right: 18rpx;
		/* width: 112rpx; */
		/* height: 112rpx; */
		vertical-align: middle;
	}

	.inner-img {
		height: 78rpx;
		width: 78rpx;
		border-radius: 50%;
	}

	/* 中部文字盒子样式 */
	.middle-text-box {
		align-items: center;
		flex: 5;
		display: flex;
	}

	/* 灰色字体 */
	.greay-text {
		color: #969696;
	}

	/* 昵称 */
	.inner-text-name {
		font-size: 35rpx;
		font-weight: bold;
		margin-right: 16rpx;
	}

	/* 关注数粉丝数文章数 */
	.inner-text-message {
		padding-top: 4rpx;
		font-size: 30rpx;
	}

	.inner-text {
		font-size: 35rpx;
	}

	>>>.article-img image {
		width: 120rpx;
		height: 120rpx;
	}

	>>>.article-text {
		font-size: 32rpx;
		font-weight: bold;
		-webkit-line-clamp: 2;
		line-height: 180%;
	}

	>>>.article-content {
		background-color: rgba(241, 241, 241, 1);
		border-radius: 4px;
		padding: 10px 12px;
	}

	.comment-content {
		margin-left: 96rpx;
		margin-top: 10rpx;
	}

	.text-comment {
		font-size: 30rpx;
		color: rgba(80, 80, 80, 1);
		width: 580rpx;
		white-space: nowrap;
		text-overflow: ellipsis;
		overflow: hidden;
	}

	.content {
		padding-left: 10rpx;
	}
</style>
