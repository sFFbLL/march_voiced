<template>
	<view>
		<view>
			<view v-for="(item,index) in draftListCopy">
				<recommend :draftsCreatime="item.create_time" :articleInfo="item" :isArticleInteract="false" class="arcitle-item item"></recommend>
			
			</view>
			<view v-if="draftListCopy.length <= 1">
				<u-empty text="没有数据" mode="search" class="nodate"></u-empty>
			</view>
		</view>
		<!-- 下拉加载更多 -->
		<view v-if="isLoadMore">
			<uni-load-more class="loading" :status="loadStatus" iconType="circle"></uni-load-more>
		</view>
	</view>

</template>

<script>
	import {
		getUserArticleList,
	} from '@/utils/api/personalCenter-api.js'
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'
	export default {
		components: {
			recommend
		},
		data() {
			return {
				draftList: [],
				draftListCopy: [],
				draftCurrent: 1, //草稿当前页数
				loadStatus: 'loading',
				draftLoadStatus: 'loading',
				isLoadMore: false, //是否加载中，
				size: 10,
				id: 0,
			}
		},
		onLoad(option) {
			this.id = Number(option.id);
			this.getDraftList();
			console.log(this.id)
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.draftCurrent = 1;
			this.draftList = [];
			setTimeout(function() {
			
				uni.stopPullDownRefresh();
			}, 1500);
			this.getDraftList();
		},
		// 触底加载
		onReachBottom() {
			if (!this.isLoadMore) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true;
				this.draftCurrent += 1;
				this.getDraftList();
			}
		},
		methods: {
			// 获取草稿箱列表
			getDraftList() {
				let params = {
					id: this.id, //用户id
					current: this.draftCurrent, //当前页
					size: this.size, //页大小
					status: 0 //草稿
				}
				getUserArticleList(params).then(res => {
					if (res.data) {
						if (this.draftCurrent === 1) {
							this.isLoadMore = false;
						} else {
							setTimeout(function() {
								this.loadStatus = "nomore";

							}, 2000);
							this.isLoadMore = false;
						}
						this.draftList = [...this.draftList, ...res.data];
						this.draftListCopy=this.draftList;
						this.isNextPage = true
					} else {
						this.loadStatus = "loading";
						this.draftLoadStatus = "loading";
						this.isLoadMore = false;
						this.isNextPage = false;
					}
				})
			},
		}
	}
</script>

<style>
	.nodate {
		background-color: #fff;
		min-height: 800rpx;
	}

	.arcitle-item {
		margin-bottom: 0;
	}

	.item {
		border-bottom: 1rpx solid #f0f0f0;
	}
</style>
