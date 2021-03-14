<template>
	<uni-list class="list-item">
		<attentionAndFansCell class="item" v-for="(item, index) in list" :key="index" :aid="item.id"
		 :nickname="item.nickname" :isFollow="item.isFollow" :avatarPath="item.avatarPath">
		</attentionAndFansCell>
		<uniLoadMore></uniLoadMore>
	</uni-list>
</template>

<script>
import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	import {
		getFansnList,
		changeStatus
	} from '../../utils/api/attention-and-fans-api.js'
	export default {
		data() {
			return {
				list: [],
				current: 1,
				size: 10
			}
		},
		components: {
			attentionAndFansCell,
			uniLoadMore
		},
		// 进入时加载
		onLoad: function(options) {
			// let that = this;
			// this.list = [];
			// setTimeout(function() {
			// 	that.list.push.apply(that.list, that.list1)
			// }, 1000);
			uni.startPullDownRefresh();
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.current = 1;
			let that = this;
			that.list = [];
			setTimeout(function() {
				that.getList()
				uni.stopPullDownRefresh();
			}, 1000);
		},
		// 触底加载
		onReachBottom() {
			this.getList()
		},
		methods: {
			//获取粉丝列表
			getList() {
				let params = {
					current: this.current,
					size: this.size
				}
				getFansnList(params).then(res => {
					console.log(res)
					if (res.data.follow.length > 0) {
						this.list.push.apply(this.list, res.data.follow)
						this.current++;
					}
				})
			},
		}
	}
</script>

<style scoped>
	/* 关注数粉丝数文章数去掉最后的 |  */
	.inner-text-message span:not(:last-child) {
		margin-right: 10rpx;
		padding-right: 10rpx;
		border-right: 1rpx solid #969696;
	}
</style>
