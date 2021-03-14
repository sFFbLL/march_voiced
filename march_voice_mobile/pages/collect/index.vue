<template>
	<view class="home">
		<view class="content">
			<view v-if="!tabIndex">
				<view v-for="(item,index) in recommendList">
					<recommend :articleInfo="item" />
				</view>
			</view>
			<view v-if="tabIndex">
				<view v-for="(item,index) in attentionList"
				 :key="index">
					<follow :articleInfo="item" />
				</view>
			</view>
			<view v-show="isLoadMore">
				<uni-load-more class="loading"
				 :status="loadStatus"
				 iconType="circle"></uni-load-more>
			</view>
		</view>
	</view>
</template>

<script>
	import uniIcons from "@/components/uni-icons/uni-icons.vue"
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	import follow from '../../marchVoiceComponents/showArticle/follow.vue'
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	import {
		getRecommend
	} from '@/utils/api/home-api.js'

	export default {
		components: {
			uniIcons,
			recommend,
			follow,
			uniLoadMore
		},
		data() {
			return {
				recommendCurrent: 1, //推荐当前页数，
				attentionCurrent: 1, //关注当前页数
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
				attentionList: [],
				loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				isLoadMore: false, //是否加载中
			}
		},
		onLoad() {
			if (!this.tabIndex) {
				this.recommend();
			} else {
				this.attention();
			}
		},

		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore && !this.tabIndex) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true
				this.recommendCurrent += 1
				this.recommend();
			} else if (!this.isLoadMore && this.tabIndex) {
				this.isLoadMore = true
				this.attentionCurrent += 1
				this.attention();
			}
		},
		methods: {
			recommend() {
				let _this=this;
				let params = {
					current: this.current,
					size:this.size
				}
				getRecommend(params).then(res => {
					_this.recommendList = [..._this.recommendList,...res.data];
					if(res.data.length<=_this.size){
						_this.loadStatus=nomore;
					}
				})
				if (this.recommendList.length > 16) {
					_this.loadStatus = "nomore";
				} else {
					setTimeout(function () {
						_this.isLoadMore = false;
						_this.recommendList = [..._this.recommendList, ...recommendList];
					}, 2000);
				}
			},
			attention() {
			
				let _this = this;
				let params = {
					current: this.current,
					size:this.size
				}
				getRecommend(params).then(res => {
					_this.recommendList = [..._this.recommendList,...res.data];
					if(res.data.length<=_this.size){
						_this.loadStatus=nomore;
					}
				})
				if (this.attentionList.length > 16) {
					_this.loadStatus = "nomore";
				} else {
					setTimeout(function () {
						_this.isLoadMore = false;
						_this.attentionList = [..._this.attentionList, ...attentionList];
					}, 2000);
				}
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
</style>
style>
