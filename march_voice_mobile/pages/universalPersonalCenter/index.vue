<template>
	<view class="personal-center">
		<!-- 头部用户信息 -->
		<view class="header">
			<attentionAndFansCell :aid="userInfo.id" :nickname="userInfo.nickname" :avatarPath="userInfo.avatarPath" :isFollow="userInfo.isFollow"
			 :isMine="true" class="top-user-info">
				<view slot="underText" class="user-signature">{{userInfo.signature}}</view>
			</attentionAndFansCell>
			<view class="total">
				<view class="fans-total" @click="toFans">
					<text class="number">{{userInfo.fansTotal}}</text>
					<text class="text">粉丝</text>
				</view>
				<view class="follow-total" @click="toAttention">
					<text class="number">{{userInfo.followTotal}}</text>
					<text class="text">关注</text>
				</view>
				<view class="collect-total" @click="toCollect" v-if="userInfo.isMe===1">
					<text class="number">{{userInfo.collectTotal}}</text>
					<text class="text">收藏</text>
				</view>
			</view>
			<view class="edit-user" v-if="userInfo.isMe===1">
				<button :plain="true" @click="goToEdit()">编辑个人信息</button>
			</view>
		</view>

		<!-- 用户文章想法列表 -->
		<view class="content">
			<view class="tab-nav">
				<tabCarb :tabs="tabs" v-on:tabActive='tabActive'></tabCarb>
			</view>
			<view class="kind-article-list">
				<!-- 文章列表 -->
				<view v-if="!tabIndex">
					<view v-for="(item,index) in articleList">
						<recommend :articleInfo="item" class="arcitle-item item"></recommend>
					</view>
					<view v-if="articleList.length < 1">
						<u-empty text="没有数据"
						 mode="search"
						 class="nodate"></u-empty>
					</view>
				<view v-for="(item,index) in articleListCopy" v-if="!tabIndex">
					<recommend :articleInfo="item" class="arcitle-item item"></recommend>

				</view>
				
				<!-- 想法列表 -->
				<view v-if="tabIndex === 1">
					<view v-for="(item,index) in ideaList">
						<view class="ideacontent item">
							<!-- 用户头像公共组件 -->
							<attentionAndFansCell :notTap="notTap" :nickname="userInfo.nickname" :avatarPath="userInfo.avatarPath" :isFollow="userInfo.isFollow">
								<view slot="underText">{{item.updateTime}}</view>
							</attentionAndFansCell>
							<!-- 想法的文字部分 -->
							<articleContent :articleContent="item.content" :isIdea="true" :id="item.id"></articleContent>
							<!-- 想法的图片部分组件 -->
							<imageAdaptation :imgList="item.imgList"></imageAdaptation>
							<!-- 点赞表情组件+评论 -->
							<emojiControl class="emoji-control" :faceTotals="item.faceTotal" :likeTotals="item.likeTotal" :favourTotals="item.favourTotal"
							 :commentTotals="item.commentTotal" :id="item.id"></emojiControl>
						</view>
					</view>
					<view v-if="ideaList.length < 1">
						<u-empty text="没有数据"
						 mode="search"
						 class="nodate"></u-empty>
					</view>
				</view>

				<!-- 下拉加载更多 -->
				<view v-show="isLoadMore">
					<uni-load-more class="loading" :status="loadStatus" iconType="circle"></uni-load-more>
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
				notTap: true,
				loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				articleLoadStatus: 'loading',
				ideaLoadStatus: 'loading',
				draftLoadStatus: 'loading',
				isLoadMore: false, //是否加载中
				tabIndex: '',
				userInfo: {},
				userId: 0,
				articleList: [],
				ideaList: [],
				articleListCopy: [],
				ideaListCopy: [],
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
		onLoad(option) {
			this.userId = Number(option.id);
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.articleCurrent = 1; //文章当前页数，
			this.ideaCurrent = 1; //想法当前页数
			let that = this;
			that.articleList = [];
			that.ideaList = [];
			this.getArticleList();
			this.getIdeaList();
			setTimeout(function() {
				uni.stopPullDownRefresh();
			}, 2000);
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
			}
		},
		methods: {
			// 点击粉丝数事件
			toFans() {
				uni.navigateTo({
					url: '../fans/index'
				})
			},
			// 点击关注数事件
			toAttention() {
				uni.navigateTo({
					url: '../attention/index'
				})
			},
			// 点击收藏数事件
			toCollect() {
				uni.navigateTo({
					url: '../collect/index'
				})
			},

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
				}
				this.tabIndex = tabIndex;
			},
			getUserInfo() {
				let _this = this;
				let params = {
					id: _this.userId
				}
				getUserInfo(params).then(res => {
					_this.userInfo = res.data;
				})
			},
			// 获取想法列表
			getIdeaList() {
				let params = {
					id: this.userId,
					current: this.ideaCurrent,
					size: this.size
				}
				getUserIdeaList(params).then(res => {
					if (res.data) {
						if (this.ideaCurrent === 1) {
							this.isLoadMore = false;
						} else {
							setTimeout(function() {
								this.loadStatus = "nomore";

							}, 2000);
							this.isLoadMore = false;
						}

						this.ideaList = [...this.ideaList, ...res.data];
					} else {
						this.loadStatus = 'nomore';
					}


				})

			},
			// 获取文章列表
			getArticleList() {
				let params = {
					id: this.userId,
					current: this.articleCurrent,

					size: this.size,
					kind: 2
				}
				getUserArticleList(params).then(res => {
					if (res.data) {
						if (this.articleCurrent === 1) {
							this.isLoadMore = false;
						} else {
							setTimeout(function() {
								this.loadStatus = "nomore";

							}, 2000);
							this.isLoadMore = false;
						}
						this.articleList = [...this.articleList, ...res.data];
						this.articleListCopy = this.articleList;
					} else {
						this.loadStatus = 'nomore';
					}
				})

			},
			goToEdit() {
				uni.navigateTo({
					url: '../personalInfo/index?id=' + this.userInfo.id
				})
			}
		},
		created() {
			this.getUserInfo();
			this.getArticleList();
			this.getIdeaList();
		},
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

	.nodate {
		background-color: #fff;
		min-height: 800rpx;
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
	
	.nodate {
		background-color: #fff;
		min-height: 800rpx;
	}
</style>
