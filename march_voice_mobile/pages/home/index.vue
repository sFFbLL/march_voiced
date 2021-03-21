<template>
	<!-- 首页 -->
	<view class="home">
		<view class="header">
			<!-- 头部选项卡 -->
			<tabs class="tag-nav" :tabs='tablist' v-on:tabActive='tabActive' />
			<!-- 搜索图标 -->
			<view class="search" @click="search()">
				<view class="search-icon">
					<uni-icons type="search" size="25" color="#999"></uni-icons>
				</view>
			</view>
		</view>
		<view class="content">
			<!-- 推荐 -->
			<view v-if="!tabIndex">
				<view v-for="(item,index) in recommendList">
					<recommend v-if="recommendList.length" :notshow="notshow" :articleInfo="item" />
				</view>
			</view>
			<!-- 关注 -->
			<view v-if="tabIndex">
				<view v-for="(item,index) in followList" :key="index">
					<follow :articleInfo="item" />
				</view>
			</view>
			<!-- 下拉加载更多 -->
			<view v-show="isLoadMore">
				<uni-load-more class="loading" :status="loadStatus" iconType="circle"></uni-load-more>
			</view>
		</view>
	</view>
</template>

<script>
	import uniIcons from "@/components/uni-icons/uni-icons.vue"
	import tabs from '../../marchVoiceComponents/tabCard.vue'
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	import follow from '../../marchVoiceComponents/showArticle/follow.vue'
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	import {
		check
	} from '../../utils/checkUnRead.js'
	import {
		getRecommend,
		getFollow,
		getArticleList
	} from '@/utils/api/home-api.js'
	import {
		forLogin
	} from '../../utils/silentLogin.js'
	import {
		returnWxcode,
		getWxCode,
		parseCode
	} from "../../utils/wxcode.js"
	import {
		getToken,
		setToken,
		setOpenId,
		getOpenId
	} from "../../utils/auth.js"
	import {
		login,
		creatNewUser
	} from "../../utils/login.js"
	export default {
		components: {
			uniIcons,
			tabs,
			recommend,
			follow,
			uniLoadMore
		},
		data() {
			return {
				notshow: true,
				recommendCurrent: 1, //推荐当前页数，
				articleCurrent: 1, //普通文章页数
				followCurrent: 1, //关注当前页数
				size: 5,
				tabIndex: '',
				tablist: [{
						index: 0,
						value: '推荐',
						isActive: true
					},
					{
						index: 1,
						value: '关注',
						isActive: false
					}
				],
				recommendList: [],
				followList: [],
				loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				recommendLoadStatus: 'loading',
				followLoadStatus: 'loading',
				isLoadMore: false, //是否加载中
				isRecommend: true, //是否是推荐文章
			}
		},
		// mounted() {
		// 	this.recommend();
		// 	this.follow();
		// },
		created() {
			if (!getToken()) {
				forLogin();
				this.recommend();
				this.follow();
			} else {
				this.recommend();
				this.follow();
			}
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.recommendCurrent = 1; //推荐当前页数，
			this.articleCurrent = 1; //普通文章页数
			this.followCurrent = 1; //关注当前页数
			this.recommendList = [];
			this.followList = [];
			this.recommend();
			this.follow();
			setTimeout(function() {
				uni.stopPullDownRefresh();
			}, 2000);
		},
		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore && !this.tabIndex) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true
				if (this.isRecommend) {
					this.recommendCurrent += 1;
					this.recommend();
				} else {
					this.articleCurrent += 1;
					this.getArticleList();
				}

			} else if (!this.isLoadMore && this.tabIndex) {
				this.isLoadMore = true
				this.followCurrent += 1
				this.follow();
			}
		},
		methods: {
			/* 切换选项卡选项 */
			tabActive(tabIndex) {
				this.tablist.map((value, index) => {
					value.isActive = tabIndex == index ? true : false;
				})
				if (!tabIndex) {
					this.loadStatus = this.recommendLoadStatus;
					this.isLoadMore = false;
				} else if (tabIndex) {
					this.loadStatus = this.followLoadStatus;
					this.isLoadMore = false;
				}
				this.tabIndex = tabIndex;
			},
			// 推荐
			recommend() {
				let _this = this;
				let params = {
					current: this.recommendCurrent,
					size: this.size
				}
				getRecommend(params).then(res => {
					if (res.data) {
						if (this.recommendCurrent === 1) {
							_this.isLoadMore = false;
						} else {
							setTimeout(function() {
								_this.isLoadMore = false;
							}, 3000);
						}
						if (res.data.length <= this.size) {
							_this.getArticleList();
						}
						_this.recommendList = [..._this.recommendList, ...res.data];
					} else {
						this.loadStatus = 'nomore';
						_this.getArticleList();
					}

				})
			},
			getArticleList() {
				let _this = this;
				let params = {
					current: this.articleCurrent,
					size: this.size
				}
				getArticleList(params).then(res => {
					if (res.data) {
						if (this.articleCurrent === 1) {
							_this.isLoadMore = false;

						} else {
							setTimeout(function() {
								_this.isLoadMore = false;
							}, 3000);
						}
						_this.recommendList = [..._this.recommendList, ...res.data];
					} else {
						this.loadStatus = 'nomore';
					}

				})
			},
			// 关注
			follow() {
				let _this = this;
				let params = {
					current: this.followCurrent,
					size: this.size
				}
				getFollow(params).then(res => {
					if (res.data) {
						if (this.followCurrent === 1) {
							_this.isLoadMore = false;
						} else {
							setTimeout(function() {
								_this.isLoadMore = false;
							}, 3000);
						}
						_this.followList = [..._this.followList, ...res.data.records];
					} else {
						this.loadStatus = 'nomore';
					}


				})
			},
			search() {
				uni.navigateTo({
					url: '../search/index'
				})
			}
		}
	}
</script>

<style scoped>
	.home,
	.home .content {
		background-color: #f7f7f7;
		margin-bottom: 30rpx;
	}

	.header {
		position: fixed;
		z-index: 10;
		width: 100%;
		height: 100rpx;
		line-height: 100rpx;
		background-color: #fff;
		margin-bottom: 15rpx;
	}

	.tag-nav {
		display: inline-block;
		height: 100%;
		width: 90%;
	}

	.search {
		display: inline-block;
		width: 8%;
	}

	.search .search-icon {
		position: absolute;
		top: 5rpx;
	}

	.content {
		background-color: #fff;
		padding-top: 115rpx;
	}

	.loading {
		background-color: #fff;
	}

	>>>.uni-load-more .uni-load-more__text {
		font-size: 28rpx;
	}

	>>>.uni-load-more .uni-load-more__img {
		height: 30rpx !important;
		width: 30rpx !important;
	}
</style>
