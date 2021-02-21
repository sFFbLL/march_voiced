<template>
	<view class="message">
		<view class="header">
			<tabs :type="message" class="tag-nav" :tabs='tablist' v-on:tabActive='tabActive' />
		</view>
		<view class="content">
			<!-- 互动消息 -->
			<view v-if="tabIndex==0" v-for="item in interactList.records">
				<!-- 单个消息组件 -->
				<singleMessage :messageInfo="item"></singleMessage>
			</view>
			<!-- 关注消息 -->
			<view v-if="tabIndex==1" v-for="item in attentionList.records">
				<attentionAndFansCell :nickname="item.nickname" :avatarPath="item.avatarPath" :isFollow="item.isFollow">
					<template v-slot:afterNicknameText>
						<text class="slot">关注了你
						</text>
					</template>
				</attentionAndFansCell>
			</view>
			<!-- 系统消息 -->
			<view v-if="tabIndex==2" v-for="item in otherList.records">
				<!-- 其他消息组件 -->
				<otherMessage :otherList="item"></otherMessage>
				<!-- 间隔槽 -->
				<u-gap height="2" bg-color="#f5f5f5"></u-gap>
			</view>
		</view>
		<view v-show="isLoadMore">
			<uni-load-more :status="loadStatus"></uni-load-more>
		</view>
	</view>
</template>

<script>
	import tabs from '../../marchVoiceComponents/tabCard.vue'
	import articleContent from "../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue"
	import singleMessage from "../../marchVoiceComponents/message/singleMessage.vue"
	import otherMessage from "../../marchVoiceComponents/message/otherMessage.vue"
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
			attentionAndFansCell,
			otherMessage
		},
		data() {
			return {
				interactCurrent: 1, //互动当前页数，
				attentionCurrent: 1, //关注当前页数
				otherCurrent: 1, //其他消息当前页数,
				size: 10,
				isLoadMore: false, //是否加载中
				loadStatus: 'loading',
				interactLoadStatus: 'loading',
				attentionLoadStatus: 'loading',
				otherLoadStatus: 'loading',
				isLoading: false, // 是否为加载中
				isDisabled: false, // 是否禁用按钮点击
				isAttention: false,
				tabIndex: 0,
				message: true,
				attentionRead: false,
				otherRead: false,
				interactRead: false,
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
				interactList: {
					pages: "22",
					total: "100",
					current: "1",
					size: "1",
					records: [{
							type: 0,
							status: 1,
							title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
							articleId: 1,
							comment: "哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",
							content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
							image: require('static/img/2.jpg'),
							createTime: "2020/12/12",
							userId: "1",
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",

						},
						{
							type: 0,
							status: 1,
							title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
							articleId: 1,
							comment: "哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",
							content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
							image: require('static/img/2.jpg'),
							createTime: "2020/12/12",
							userId: "1",
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",

						},
						{
							type: 0,
							status: 1,
							title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
							articleId: 1,
							comment: "哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",
							content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
							image: require('static/img/2.jpg'),
							createTime: "2020/12/12",
							userId: "1",
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",

						},
						{
							type: 0,
							status: 1,
							title: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",
							articleId: 1,
							comment: "哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",
							content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
							image: require('static/img/2.jpg'),
							createTime: "2020/12/12",
							userId: "1",
							avatarPath: require('static/img/1.jpg'),
							nickname: "张三",

						},
					]
				},


			//关注消息
			attentionList: {
					pages: "22",
					total: "100",
					current: "1",
					size: "1",
					records: [
						{
						avatarPath: require('static/img/1.jpg'),
						nickname: "张三",
						id: 1,
						isFollow: 0,
						createTime: ""
					}, {
						avatarPath: require('static/img/1.jpg'),
						nickname: "张三",
						id: 1,
						isFollow: 0,
						createTime: ""
					}, {
						avatarPath: require('static/img/1.jpg'),
						nickname: "张三",
						id: 1,
						isFollow: 0,
						createTime: ""
					}, {
						avatarPath: require('static/img/1.jpg'),
						nickname: "张三",
						id: 1,
						isFollow: 0,
						createTime: ""
					},
					]
				},


				//其他消息
				otherList:{
					pages: "22",
					total: "100",
					current: "1",
					size: "1",
					records:[{
					type:"0",
					status:"1",
					createTime:"2020-12-12"
					},{
					type:"1",
					status:"1",
					createTime:"2020-12-12"
					},{
					type:"1",
					status:"0",
					createTime:"2020-12-12"
					},{
					type:"0",
					status:"0",
					createTime:"2020-12-12"
					},]
				},

		}
	},
	//上拉触底函数
	onReachBottom() {
			if (!this.isLoadMore && this.tabIndex == 0) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true
				this.interactCurrent += 1
				this.interact();
			} else if (!this.isLoadMore && this.tabIndex == 1) {
				this.isLoadMore = true
				this.attentionCurrent += 1
				this.attention();
			} else if (!this.isLoadMore && this.tabIndex == 2) {
				this.isLoadMore = true
				this.otherCurrent += 1
				this.other();
			}
		},
		// 下拉刷新获取最新数据
		onPullDownRefresh() {
			// 查询是否有未读消息
			// unreadMessage(1).then(res=>{
			// 	if(res.data.count>0){
			// 		this.$store.commit('changeInteract',0);
			// 	}
			// })
			// unreadMessage(2).then(res=>{
			// 	if(res.data.count>0){
			// 		this.$store.commit('changeAttention',0);
			// 	}
			// })
			// unreadMessage(3).then(res=>{
			// 	if(res.data.count>0){
			// 		this.$store.commit('changeOther',0);
			// 	}
			// })
			if (this.tabIndex == 0) {
				// 获取最新消息数据
				// this.interact();
				// 消除当前tab红点
				this.$store.commit('changeInteract', 1);

			} else if (this.tabIndex == 1) {
				// this.attention();
				this.$store.commit('changeAttention', 1);

			} else if (this.tabIndex == 2) {
				// this.other();
				this.$store.commit('changeOther', 1);

			}
			setTimeout(function() {
				uni.stopPullDownRefresh();
			}, 1000);
		},
		onShow() {
			uni.hideTabBarRedDot({
				index: 3
			})
			// 所有消息已读，消除底部消息红点
			// if ( this.$store.state.changeInteract == true &&this.$store.state.attentionRead == true && this.$store.state.otherRead == true) {

			// }
		},
		created() {
			// 互动消息已读
			this.$store.commit('changeInteract', 1);
			// readMessage(1).then(res => {
			// 	
			// })

			// 查询所有互动消息
			this.interact()
			// 查询关注和其他是否有未读消息
			// unreadMessage(2).then(res=>{
			// 	if(res.data.count>0){
			// 增加红点
			// 		this.$store.commit('changeAttention',0);
			// 	}
			// })
			// unreadMessage(3).then(res=>{
			// 	if(res.data.count>0){
			// 增加红点
			// 		this.$store.commit('changeOther',0);
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
					this.loadStatus = this.interactLoadStatus;
					this.interact();
				} else if (tabIndex == 1) {
					this.loadStatus = this.attentionLoadStatus;
					this.attention();
				} else {
					this.loadStatus = this.otherLoadStatus;
					this.other();
				}
				this.tabIndex = tabIndex;
			},
			// 请求互动消息数据，并发送已读请求
			interact() {
				let _this = this;
				let total = "";
				// 查询所有数据
				// let params = {
				// current: this.interactCurrent,
				// 	size:this.size
				// }
				// interactList(params).then(res => {
				// _this.interactList = [..._this.interactList,...res.data];
				// if(res.data.length<=_this.size){
				// 	_this.loadStatus='nomore';
				// }
				// })
				// 如果长度大于总条数
				if (this.interactList.length > 16) {
					_this.loadStatus = "nomore";
					_this.interactLoadStatus = "nomore";
				} else if (this.interactCurrent === 1) {
					_this.isLoadMore = false;
				} else {
					setTimeout(function() {
						_this.isLoadMore = false;
					}, 2000);
				}
				// 数据已读
				// readMessage(1).then(res => {
				// 	
				// })
			},
			// 请求关注消息数据，并发送已读请求
			attention() {
				let _this = this;
				// let params = {
				// current: this.attentionCurrent,
				// 	size:this.size
				// }
				// attentionList(params).then(res => {
				// _this.attentionList = [..._this.attentionList,...res.data];
				// if(res.data.length<=_this.size){
				// 	_this.loadStatus='nomore';
				// }
				// })
				if (this.attentionList.length > 16) {
					_this.loadStatus = "nomore";
					_this.attentionLoadStatus = "nomore";
				} else if (this.attentionCurrent === 1) {
					_this.isLoadMore = false;
				} else {
					setTimeout(function() {
						_this.isLoadMore = false;
					}, 2000);
				}
				// 数据已读
				// readMessage(1).then(res => {
				// 	
				// })
			},
			// 请求系统消息数据，并发送已读请求
			other() {
				let _this = this;
				// let params = {
				// current: this.otherCurrent,
				// 	size:this.size
				// }
				// otherList(params).then(res => {
				// _this.otherList = [..._this.otherList,...res.data];
				// if(res.data.length<=_this.size){
				// 	_this.loadStatus='nomore';
				// }
				// })
				if (this.otherList.length > 16) {
					_this.loadStatus = "nomore";
					_this.otherLoadStatus = "nomore";
				} else if (this.otherCurrent === 1) {
					_this.isLoadMore = false;
				} else {
					setTimeout(function() {
						_this.isLoadMore = false;
					}, 2000);
				}
				// 数据已读
				// readMessage(1).then(res => {
				// 	
				// })
			}
		},


	}
</script>

<style scoped>
	.header {
		position: fixed;
		top: 0;
		background-color: white;
		z-index: 99;
		width: 100%;
		height: 80rpx;
		line-height: 80rpx;
	}

	/* 公共头像组件样式 */
	>>>.attention-cell .flex-item {
		padding: 10rpx;
		margin-top: 17rpx;

	}

	/* 关注插槽 */
	.slot {
		margin-left: 20rpx;
		font-weight: 400;
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
		margin-top: 80rpx;
	}
</style>
