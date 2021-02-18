<template>
	<view class="personal-center">
		<view class="header">
			<attentionAndFansCell :id="userInfo.user.id"
			 :nickname="userInfo.user.nickname"
			 :avatarPath="userInfo.user.avatarPath"
			 :isFollow="userInfo.user.isFollow">
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
				<button :plain="true">编辑个人信息</button>
			</view>
		</view>
		<view class="content">
			<view class="tab-nav">
				<tabCarb :tabs="tabs"
				 v-on:tabActive='tabActive'></tabCarb>
			</view>
			<view class="kind-article-list">
				<view v-for="(item,index) in articleList1">
					<recommend :articleInfo="item"
					 class="arcitle-item"></recommend>
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
	import attentionAndFansCell from "../../marchVoiceComponents/attentionAndFansCell.vue"
	import tabCarb from '../../marchVoiceComponents/tabCard.vue'
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	import uniLoadMore from '../../components/uni-load-more/uni-load-more.vue'
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
				userInfo: {
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
				articleList: [],
				articleList1: [{
					id: 1,
					title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
					"content": "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
					"image": require('static/img/2.jpg'),
					upDateTime: "更新时间",
					"favourTotal": 1,
					"collectTotal": 1,
					"commentTotal": 1
				}, {
					id: 1,
					title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
					"content": "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
					"image": require('static/img/2.jpg'),
					"favourTotal": 1,
					"collectTotal": 1,
					"commentTotal": 1
				}, {
					id: 1,
					title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
					"content": "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
					"image": require('static/img/2.jpg'),
					"favourTotal": 1,
					"collectTotal": 1,
					"commentTotal": 1
				}, {
					id: 1,
					title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
					"content": "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
					"image": require('static/img/2.jpg'),
					"favourTotal": 1,
					"collectTotal": 1,
					"commentTotal": 1
				}],
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
			uniLoadMore
		},
		methods: {
			/* 切换选项卡选项 */
			tabActive(tabIndex) {
				this.tabs.map((value, index) => {
					value.isActive = tabIndex == index ? true : false;
				})
				if (tabIndex === 0) {} else if (tabIndex === 1) {} else if (tabIndex === 2) {}
				this.tabIndex = tabIndex;
			},
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
	>>>.inner-img {
		width: 112rpx;
		height: 112rpx;
		margin-right: 18rpx;
		box-shadow: 0px 0px 4px 1px rgba(0, 0, 0, 0.5);
	}

	/* 修改昵称样式 */
	>>>.inner-text-name {
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
		height: 100rpx;
		background-color: #fff;
	}

	.arcitle-item {
		padding: 30rpx 0;
		margin-bottom: 0;
		border-bottom: 1rpx solid #f0f0f0;
	}

	>>>.article-title .article-text {
		font-size: 34rpx;
	}
</style>
