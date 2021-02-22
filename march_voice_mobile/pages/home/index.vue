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
					<recommend :articleInfo="item" />
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
		getFollow
	} from '@/utils/api/home-api.js'
	import {
		returnWxcode,
		getWxCode
	} from "../../utils/wxcode.js"
	import {
		getToken,
		setToken
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
				recommendCurrent: 1, //推荐当前页数，
				followCurrent: 1, //关注当前页数
				size: 10,
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
			}
		},
		beforeCreate() {
			// 判断是否有token
			if (!getToken()) {
				//没有token，没登陆过，获取wxcode
				let code = returnWxcode();
				console.log(code)
				login(code).then(res => {
					if (res.data.status == 1) {
						// 跳转注册页面
						uni.navigateTo({
							url: "../login/login"
						})

					} else {
						// 登陆成功
						setToken(res.data.token);
						setOpenId(res.data.openid)
					}
				})

			}

		},
		onShow() {
			check()
		},
		onLoad() {
			this.recommend();
			this.follow();
		},

		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore && !this.tabIndex) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true
				this.recommendCurrent += 1
				this.recommend();
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
				let recommendList = [{
					id: 1,
					title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
					content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
					image: require('static/img/2.jpg'),
					upDateTime: "更新时间",
					favourTotal: 1,
					collectTotal: 1,
					commentTotal: 1,
					user: {
						id: 1,
						nickname: "张三",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 0
					}
				}, {
					id: 2,
					title: "所以监听用户的截图操作，提示用户进行分，我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",
					content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",
					image: "",
					upDateTime: "更新时间",
					favourTotal: 1,
					collectTotal: 1,
					commentTotal: 1,
					user: {
						id: 2,
						nickname: "李四",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 1
					}
				}, {
					id: 1,
					title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
					content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
					image: require('static/img/2.jpg'),
					upDateTime: "更新时间",
					favourTotal: 1,
					collectTotal: 1,
					commentTotal: 1,
					user: {
						id: 1,
						nickname: "张三",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 0
					}
				}, {
					id: 1,
					title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
					content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
					image: require('static/img/2.jpg'),
					upDateTime: "更新时间",
					favourTotal: 1,
					collectTotal: 1,
					commentTotal: 1,
					user: {
						id: 1,
						nickname: "张三",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 0
					}
				}];
				let _this = this;
				/* let params = {
					current: this.current,
					size:this.size
				}
				getRecommend(params).then(res => {
					_this.recommendList = [..._this.recommend,...res.data];
					if(res.data.length<=_this.size){
						_this.loadStatus=nomore;
					}
				}) */
				if (this.recommendList.length > 16) {
					_this.loadStatus = "nomore";
					_this.recommendLoadStatus = "nomore";
				} else if (this.recommendCurrent === 1) {
					_this.isLoadMore = false;
					_this.recommendList = [..._this.recommendList, ...recommendList];
				} else {
					setTimeout(function() {
						_this.isLoadMore = false;
						_this.recommendList = [..._this.recommendList, ...recommendList];
					}, 2000);
				}
			},
			// 关注
			follow() {
				// console.log(2);
				let followList = [{
						articleId: 2,
						title: "所以监听用户的截图操作，提示用户进行分，我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						image: "",
						createTime: "2020-08-24",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1,
						status: 0,
						user: {
							id: 2,
							nickname: "李四",
							avatarPath: require('../../static/img/1.jpg')
						}
					}, {
						articleId: 2,
						title: "所以监听用户的截图操作，提示用户进行分，我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						image: require('static/img/2.jpg'),
						createTime: "2020-08-24",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1,
						status: 1,
						user: {
							id: 2,
							nickname: "李四",
							avatarPath: require('../../static/img/1.jpg')
						}
					},
					{
						articleId: 2,
						title: "所以监听用户的截图操作，提示用户进行分，我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						image: "",
						createTime: "2020-08-24",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1,
						status: 2,
						user: {
							id: 2,
							nickname: "李四",
							avatarPath: require('../../static/img/1.jpg')
						}
					}, {
						articleId: 2,
						title: "所以监听用户的截图操作，提示用户进行分，我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						image: require('static/img/2.jpg'),
						createTime: "2020-08-24",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1,
						status: 3,
						user: {
							id: 2,
							nickname: "李四",
							avatarPath: require('../../static/img/1.jpg')
						}
					}
				];
				let _this = this;
				/* let params = {
					current: this.current,
					size:this.size
				}
				getFollow(params).then(res => {
					_this.recommendList = [..._this.recommend,...res.data];
					if(res.data.length<=_this.size){
						_this.loadStatus=nomore;
					}
				}) */
				if (this.followList.length > 16) {
					_this.loadStatus = "nomore";
				} else if (this.followCurrent === 1) {
					_this.isLoadMore = false;
					_this.followList = [..._this.followList, ...followList];
				} else {
					setTimeout(function() {
						_this.isLoadMore = false;
						_this.follLoadStatus = "nomore";
						_this.followList = [..._this.followList, ...followList];
					}, 2000);
				}
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
