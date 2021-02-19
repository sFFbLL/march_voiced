<template>
	<view class="search">
		<!-- 搜索页面头部 -->
		<view class="search-header">
			<view class="search-input">
				<u-search :show-action="false"
				 placeholder="搜索感兴趣的内容"
				 v-model="searchText"
				 @search="search()"
				 @clear="clear()"></u-search>
			</view>
		</view>
		<view class="content">
			<view v-if="!userList.length && !articleList.length">
				<u-empty text="没有数据"
				 mode="search"
				 class="nodate"></u-empty>
			</view>
			<!-- 用户列表-->
			<view class="search-user"
			 v-if="userList.length">
				<u-section title="相关用户"
				 :show-line="false"
				 :bold="false"
				 color="#969696"
				 sub-color="#969696"
				 @click="moreUser()"></u-section>
				<view class="user-list">
					<view class="user-item"
					 v-for="(item,index) in userList">
						<image :src=item.avatarPath
						 mode="aspectFill"
						 class="user-img"></image>
						<view class="username">{{item.nickname}}
						</view>
					</view>
				</view>
			</view>
			<view class="divide"
			 v-if="userList.length && articleList.length"></view>
			<!-- 文章列表 -->
			<view class="search-article"
			 v-if="articleList.length">
				<view v-for="(item,index) in articleList"
				 :key="index">
					<searchArticleItem :articleInfo="item" />
				</view>
			</view>
		</view>
		<!-- 下拉加载更多 -->
		<view v-show="isLoadMore">
			<uni-load-more class="loading"
			 :status="loadStatus"
			 iconType="circle"></uni-load-more>
		</view>
	</view>
	</view>
</template>

<script>
	import {
		searchArticle,
		searchUser
	} from '@/utils/api/home-api.js'
	import searchArticleItem from "@/marchVoiceComponents/showArticle/searchArticle.vue"
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	export default {
		data() {
			return {
				searchText: "",
				articleList: [],
				userList: [],
				loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				isLoadMore: false, //是否加载中
				articleCurrent: 1, //关注当前页数
				articleSize: 10,
				userSize: 4,
				articleList1: [{
						id: 1,
						title: "所以监听用户的截图操作，提示<font color=red>用户</font>进行分，我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。<font color=red>在写作</font>一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",
						createTime: "2021-02-16",
						favourTotal: 0,
						commentTotal: 0,
						nickname: "张三"
					},
					{
						id: 2,
						title: "个大学生啊，我该怎么学编程？我还是个大学生啊我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",
						createTime: "2021-02-16",
						favourTotal: 0,
						commentTotal: 0,
						nickname: "张三"
					},
					{
						id: 3,
						title: "所以监听用户的截图操作，提示用户进生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",
						createTime: "2021-02-16",
						favourTotal: 0,
						commentTotal: 0,
						nickname: "张三"
					},
					{
						id: 4,
						title: "所以监听用户的截图操怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",
						createTime: "2021-02-16",
						favourTotal: 0,
						commentTotal: 0,
						nickname: "张三"
					},
					{
						id: 5,
						title: "所以监听用户的截图操怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",
						createTime: "2021-02-16",
						favourTotal: 0,
						commentTotal: 0,
						nickname: "张三"
					}
				],
				userList1: [{
						id: 1,
						nickname: "张三",
						avatarPath: require('static/img/2.jpg')
					},
					{
						id: 2,
						nickname: "张四",
						avatarPath: require('static/img/1.jpg')
					},
					{
						id: 3,
						nickname: "张五",
						avatarPath: require('static/img/2.jpg')
					},
					{
						id: 4,
						nickname: "张六",
						avatarPath: require('static/img/2.jpg')
					}
				]
			}
		},
		components: {
			searchArticleItem,
			uniLoadMore
		},
		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore) {
				this.isLoadMore = true
				this.articleCurrent += 1
				this.getList();
			}
		},
		methods: {
			search() {
				// let _this = this;
				if (this.searchText) {
					// this.getArticleList();
					// this.getUserList();
					this.userList = this.userList1;
					this.articleList = [...this.articleList, ...this.articleList1];
				}
			},
			clear() {
				this.searchText = "";
			},
			getList() {
				if (this.searchText) {
					let _this = this;
					if (this.articleList.length > 16) {
						_this.loadStatus = "nomore";
					} else {
						setTimeout(function () {
							_this.isLoadMore = false;
							// this.getArticleList();
							_this.articleList = [..._this.articleList, ..._this.articleList1];
						}, 2000);
					}
				}
			},
			moreUser() {
				uni.navigateTo({
					url: '../search/searchUser'
				})
			},
			getArticleList() {
				let _this = this;
				let params = {
					searchText: this.searchText,
					current: this.current,
					size: this.articleSize
				}
				searchArticle().then(res => {
					_this.articleList = [..._this.articleList, ...res.data.articleList];
				})
			},
			getUserList() {
				let _this = this;
				let params = {
					searchText: this.searchText,
					size: this.userSize
				}
				searchUser().then(res => {
					_this.userList = res.data.userList;
				})
			}

		}
	}
</script>

<style scoped>
	.search {
		background-color: #f7f7f7;
	}

	.nodate,
	.search-user,
	.search-article,
	.search-header {
		background-color: #fff;
	}

	.search-header {
		position: fixed;
		height: 110rpx;
		width: 100%;
		line-height: 100rpx;
		background-color: #fff;
		border-bottom: 1rpx solid #f0f0f0;
		z-index: 10;
	}

	.search-input {
		width: 90%;
		margin: auto;
	}

	.content {
		padding-top: 110rpx;
	}

	.nodate {
		min-height: 800rpx;
	}

	.search-user {
		width: 100%;
		padding: 30rpx;
	}

	.user-list {
		margin-top: 30rpx;
	}

	.user-item {
		display: inline-block;
		width: 25%;
	}

	.user-img {
		display: flex;
		width: 70rpx;
		height: 70rpx;
		margin: auto;
		border-radius: 50%;
	}

	.divide {
		height: 30rpx;
		border-top: 1rpx solid #f0f0f0;
		border-bottom: 1rpx solid #f0f0f0;
	}

	.search-article {
		padding: 20rpx 30rpx 0 30rpx;
	}

	.user-item .username {
		width: 60%;
		margin: auto;
		margin-top: 10rpx;
		text-align: center;
		font-size: 22rpx;
		color: #303133;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	>>>.uni-load-more .uni-load-more__img {
		height: 30rpx !important;
		width: 30rpx !important;
	}
</style>
