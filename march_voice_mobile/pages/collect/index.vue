<template>
	<view class="home">
		<view class="content">
			<view v-for="(item,index) in correntList">
				<recommend :articleInfo="item" />
			</view>
			<!-- 间隔槽 -->
			<!-- <u-gap height="5" bg-color="#f5f5f5"></u-gap> -->
			<view v-if="isLoadMore">
				<uni-load-more class="loading" :status="loadStatus" iconType="circle"></uni-load-more>
			</view>
			
		</view>
		<view v-if="correntList.length < 1">
			<u-empty text="没有数据"
			 mode="search"
			 class="nodate"></u-empty>
		</view>
	</view>
</template>

<script>
	import uniIcons from "@/components/uni-icons/uni-icons.vue"
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	import {
		collectArticle
	} from "../../utils/api/personalCenter-api.js"
	export default {
		components: {
			uniIcons,
			recommend,
			uniLoadMore
		},
		data() {
			return {
				current: 1, //推荐当前页数，
				size: 10,
				correntList: [],
				loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				isLoadMore: false, //是否加载中
			}
		},
		onLoad() {
			this.getCollentList();
		},
		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true;
				this.current += 1;
				this.getCollentList();

			}
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.current = 1;
			this.correntList = [];
			this.getCollentList();
			setTimeout(function() {
				uni.stopPullDownRefresh();
			}, 1000);
		},
		methods: {
			getCollentList() {
				let params = {
					current: this.current,
					size: this.size
				}
				collectArticle(params).then(res => {
					if (res.data.records) {
						if (this.current === 1) {
							this.isLoadMore = false;
						} else {
							setTimeout(function() {
								this.isLoadMore = false;
							}, 3000);
						}
						this.correntList = [...this.correntList, ...res.data.records];

					} else {
						this.loadStatus = "loading";
						this.isLoadMore = false;
					}
				})
			}
		}
	}
</script>

<style>
	.home,
	.home .content {
		background-color: rgba(241, 241, 241, 1);
		margin-bottom: 30rpx;
	}

	.header {
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
	
	.nodate {
		background-color: #fff;
		min-height: 800rpx;
	}
</style>
