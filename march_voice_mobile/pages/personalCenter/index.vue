<template>
	<view class="personal-center">
		<!-- 头部用户信息 -->
		<view class="header">
			<attentionAndFansCell :id="userInfo.user.id"
			 :nickname="userInfo.user.nickname"
			 :avatarPath="userInfo.user.avatarPath"
			 :isFollow="userInfo.user.isFollow"
			 class="top-user-info">
				<view slot="underText"
				 class="user-signature">{{userInfo.signature}}</view>
			</attentionAndFansCell>
			<view class="total">
				<view class="fans-total">
					<text class="number">{{userInfo.fansTotal}}</text>
					<text class="text">粉丝</text>
				</view>
				<view class="follow-total">
					<text class="number">{{userInfo.followTotal}}</text>
					<text class="text">关注</text>
				</view>
				<view class="collect-total">
					<text class="number">{{userInfo.collectTotal}}</text>
					<text class="text">收藏</text>
				</view>
			</view>
			<view class="edit-user">
				<button :plain="true"
				 @click="goToEdit()">编辑个人信息</button>
			</view>
		</view>

		<!-- 用户文章想法列表 -->
		<view class="content">
			<view class="tab-nav">
				<tabCarb :tabs="tabs"
				 v-on:tabActive='tabActive'></tabCarb>
			</view>
			<view class="kind-article-list">
				<!-- 文章列表 -->
				<view v-for="(item,index) in articleList"
				 v-if="!tabIndex">
					<recommend :articleInfo="item"
					 class="arcitle-item item"></recommend>
				</view>
				<!-- 想法列表 -->
				<view v-for="(item,index) in ideaList"
				 v-if="tabIndex === 1">
					<view class="ideacontent item">
						<!-- 用户头像公共组件 -->
						<attentionAndFansCell :nickname="userInfo.user.nickname"
						 :avatarPath="userInfo.user.avatarPath"
						 :isFollow="userInfo.user.isFollow">
							<view slot="underText">{{item.updateTime}}</view>
						</attentionAndFansCell>
						<!-- 想法的文字部分 -->
						<articleContent :articleContent="item.content"
						 :isIdea="true"
						 :id="item.id"></articleContent>
						<!-- 想法的图片部分组件 -->
						<imageAdaptation :imgList="item.imgList"></imageAdaptation>
						<!-- 点赞表情组件 -->
						<emojiControl :emojiList="emojiList"
						 class="emoji-control"></emojiControl>
					</view>
				</view>
				<view v-for="(item,index) in draftList"
				 v-if="tabIndex === 2">
					<recommend :articleInfo="item"
					 :isArticleInteract="false"
					 class="arcitle-item item"></recommend>
				</view>
				<!-- 下拉加载更多 -->
				<view v-show="isLoadMore">
					<uni-load-more class="loading"
					 :status="loadStatus"
					 iconType="circle"></uni-load-more>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import {
		getUserInfo,
		getUserArticleList,
		getUserIdeaList
	} from '@/utils/api/personalCenter-api.js'
	import attentionAndFansCell from "../../marchVoiceComponents/attentionAndFansCell.vue"
	import tabCarb from '../../marchVoiceComponents/tabCard.vue'
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	import uniLoadMore from '../../components/uni-load-more/uni-load-more.vue'
	import emojiControl from '../../marchVoiceComponents/marchCircle/emojiControl.vue'
	import imageAdaptation from '../../marchVoiceComponents/marchCircle/imageAdaptation.vue'
	import articleContent from '../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue'
	export default {
		data() {
			return {
				articleCurrent: 1, //文章当前页数，
				ideaCurrent: 1, //想法当前页数
				draftCurrent: 1, //草稿当前页数
				size: 10,
				loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				articleLoadStatus: 'loading',
				ideaLoadStatus: 'loading',
				draftLoadStatus: 'loading',
				isLoadMore: false, //是否加载中
				tabIndex: '',
				userInfo: {},
				userInfo1: {
					signature: "向往的生活：面朝大海，春暖花开。",
					followTotal: 12,
					fansTotal: 345,
					collectTotal: 54,
					user: {
						id: 1,
						avatarPath: require('../../static/img/图片 8.jpg'),
						nickname: "吴胜科",
						isFollow: 0
					}
				},
				articleList: [{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg'),
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg'),
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg'),
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg'),
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
				],
				ideaList: [{
					id: 1,
					imgList: [
						'../../static/img/cat.jpg'
					],
					content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹",
					updateTime: "暂无时间",
					faceTotal: 0,
					likeTotal: 0,
					favourTotal: 0,
					commentTotal: 0
				}, {
					id: 1,
					imgList: [
						'../../static/img/cat.jpg', '../../static/img/cat.jpg',
						'../../static/img/cat.jpg', '../../static/img/cat.jpg',
						'../../static/img/cat.jpg',
					],
					content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹",
					updateTime: "暂无时间",
					faceTotal: 0,
					likeTotal: 0,
					favourTotal: 0,
					commentTotal: 0
				}, {
					id: 1,
					imgList: [
						'../../static/img/cat.jpg', '../../static/img/cat.jpg',
						'../../static/img/cat.jpg', '../../static/img/cat.jpg',
						'../../static/img/cat.jpg', '../../static/img/cat.jpg',
						'../../static/img/cat.jpg', '../../static/img/cat.jpg',
						'../../static/img/cat.jpg'
					],
					content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹",
					updateTime: "暂无时间",
					faceTotal: 0,
					likeTotal: 0,
					favourTotal: 0,
					commentTotal: 0
				}],
				draftList: [{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg')
					},
					{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg')
					},
					{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg')
					},
					{
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						image: require('static/img/2.jpg')
					},
				],
				emojiList: {
					faceTotal: 0,
					likeTotal: 6,
					favourTotal: 0,
					commentTotal: 0,
				},
				tabs: [{
						index: 0,
						value: '文章',
						isActive: true
					},
					{
						index: 1,
						value: '想法',
						isActive: false
					},
					{
						index: 2,
						value: '草稿箱',
						isActive: false
					}
				],
			}
		},
		components: {
			attentionAndFansCell,
			tabCarb,
			recommend,
			uniLoadMore,
			emojiControl,
			imageAdaptation,
			articleContent
		},
		onLoad() {
			this.getArticleList();
			this.getIdeaList();
			this.getDraftList();
		},

		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore && !this.tabIndex) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true;
				this.articleCurrent += 1;
				this.getArticleList();
			} else if (!this.isLoadMore && this.tabIndex === 1) {
				this.isLoadMore = true;
				this.ideaCurrent += 1;
				this.getIdeaList();
			} else if (!this.isLoadMore && this.tabIndex === 2) {
				this.isLoadMore = true
				this.draftCurrent += 1
				this.getDraftList();
			}
		},
		methods: {
			/* 切换选项卡选项 */
			tabActive(tabIndex) {
				this.tabs.map((value, index) => {
					value.isActive = tabIndex == index ? true : false;
				})
				if (tabIndex === 0) {
					this.isLoadMore = false;
					this.loadStatus = this.articleLoadStatus;
				} else if (tabIndex === 1) {
					this.isLoadMore = false;
					this.loadStatus = this.ideaLoadStatus;
				} else if (tabIndex === 2) {
					this.isLoadMore = false;
					this.loadStatus = this.draftLoadStatus;
				}
				this.tabIndex = tabIndex;
			},
			getUserInfo() {
				/* let _this = this;
				let params = {
					id: 1
				}
				getUserInfo(params).then(res => {
					_this.userInfo = res.data;
				}) */
				this.userInfo = this.userInfo1;
			},
			getIdeaList() {
				let _this = this;
				/* let params = {
					id: 1,
					current: _this.ideaCurrent,
					size: _this.size
				}
				getUserIdeaList(params).then(res => {
					_this.ideaList = [..._this.ideaList, ...res.data];
				}) */
				if (this.ideaList.length > 7) {
					_this.loadStatus = "nomore";
					_this.ideaLoadStatus = "nomore";
				} else if (this.ideaCurrent === 1) {
					_this.isLoadMore = false;
					_this.ideaList = [..._this.ideaList, ..._this.ideaList];
				} else {
					setTimeout(function () {
						_this.isLoadMore = false;
						_this.ideaList = [..._this.ideaList, ..._this.ideaList];
					}, 2000);
				}
			},
			getDraftList() {
				let _this = this;
				/* let params = {
					id: 1,
					current: _this.ideaCurrent,
					size: _this.size,
					kind: 1
				}
				getUserArticleList(params).then(res => {
					_this.draftList = [..._this.draftList, ...res.data];
				}) */
				if (this.draftList.length > 10) {
					_this.loadStatus = "nomore";
					_this.draftLoadStatus = "nomore";
				} else if (this.draftCurrent === 1) {
					_this.isLoadMore = false;
					_this.draftList = [..._this.draftList, ..._this.draftList];
				} else {
					setTimeout(function () {
						_this.isLoadMore = false;
						_this.draftList = [..._this.draftList, ..._this.draftList];
					}, 2000);
				}
			},
			getArticleList() {
				let _this = this;
				/* let params = {
					id: 1,
					current: _this.ideaCurrent,
					size: _this.size,
					kind: 2
				}
				getUserArticleList(params).then(res => {
					_this.articleList = [..._this.articleList, ...res.data];
				}) */
				if (this.articleList.length > 10) {
					_this.loadStatus = "nomore";
					_this.articleLoadStatus = "nomore";
				} else if (this.articleCurrent === 1) {
					_this.isLoadMore = false;
					_this.articleList = [..._this.articleList, ..._this.articleList];
				} else {
					setTimeout(function () {
						_this.isLoadMore = false;
						_this.articleList = [..._this.articleList, ..._this.articleList];
					}, 2000);
				}
			},
			goToEdit() {
				uni.navigateTo({
					url: '../personalInfo/index?id=' + this.userInfo.user.id
				})
			}
		},
		created() {
			this.getUserInfo();
		}
	}
</script>

<style scoped>
	.personal-center {
		background-color: #f7f7f7;
	}

	.header {
		padding: 20rpx 0;
		background-color: #fff;
	}

	.content {
		margin-top: 20rpx;
	}

	>>>.attention-cell .flex-item {
		border: none;
	}

	/* 修改头像样式 */
	>>>.top-user-info .inner-img {
		width: 112rpx;
		height: 112rpx;
		margin-right: 18rpx;
		box-shadow: 0px 0px 4px 1px rgba(0, 0, 0, 0.5);
	}

	/* 修改昵称样式 */
	>>>.top-user-info .inner-text-name {
		font-size: 36rpx;
	}

	.user-signature {
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		width: 96%;
		font-size: 28rpx;
	}

	/* 展示粉丝关注数的整体框样式 */
	.total {
		display: flex;
		margin-left: 30rpx;
		margin-top: 10rpx;
	}

	.total view {
		margin-right: 54rpx;
	}

	.total .number {
		margin-right: 5rpx;
		color: #2A82E4;
		font-size: 32rpx;
		line-height: 150%;
		font-weight: bold;
	}

	.total .text {
		color: #999;
		font-size: 28rpx;
		line-height: 150%;
	}

	.edit-user button {
		margin: 28rpx 0 0 30rpx;
		width: 594rpx;
		height: 60rpx;
		border-radius: 30rpx;
		border: #999 solid 1px;
		color: rgba(80, 80, 80, 1);
		font-size: 32rpx;
		font-weight: bold;
		line-height: 170%;
	}

	/* 修改选项卡样式 */
	.tab-nav {
		padding: 0 30rpx;
		height: 90rpx;
		width: 100%;
		line-height: 90rpx;
		background-color: #fff;
		margin-bottom: 4rpx;
	}

	>>>.tab-card .center {
		font-size: 32rpx;
	}

	.kind-article-list {
		padding: 0 30rpx;
		background-color: #fff;
	}

	.arcitle-item {
		padding: 30rpx 0;
		margin-bottom: 0;
	}

	>>>.article-title .article-text {
		font-size: 34rpx;
	}

	>>>.uni-load-more .uni-load-more__text {
		font-size: 28rpx;
	}

	>>>.uni-load-more .uni-load-more__img {
		height: 30rpx !important;
		width: 30rpx !important;
	}

	.item {
		border-bottom: 1rpx solid #f0f0f0;
	}

	/* 想法用户头像 */
	>>>.ideacontent .attention-cell .flex-item {
		padding: 30rpx 0 0 0;
	}

	>>>.allImage {
		margin-top: 20rpx;
	}

	.emoji-control {
		margin-bottom: 30rpx;
	}
</style>
