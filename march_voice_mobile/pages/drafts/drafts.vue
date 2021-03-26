<template>
	<view>
		draft
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
	export default {
		data() {
			return {
				draftList: [],
				draftCurrent: 1, //草稿当前页数
				loadStatus: 'loading',
				draftLoadStatus: 'loading',
				
				isLoadMore: false, //是否加载中
			}
		},
		methods: {
			// 获取草稿箱列表
			getDraftList() {
				let params = {
					id: 0,
					current: this.draftCurrent,
					size: this.size,
					status: 0
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

</style>
