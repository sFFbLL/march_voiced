<template>
	<view class="">
		<uni-list class="list-item">
			<attentionAndFansCell class="item" v-for="(item, index) in list" :key="index" :aid="item.id"
			 :id="item.id" :nickname="item.nickname" :isFollow="item.isFollow" :avatarPath="item.avatarPath">
				<span slot="middleText">
					<span class="inner-text">关注 {{item.followTotal}}</span>
					<span class="inner-text">粉丝 {{item.fansTotal}}</span>
					<span class="inner-text">文章 {{item.articleTotal}}</span>
				</span>
				<span slot="underText">{{item.signature}}</span>
			</attentionAndFansCell>
		</uni-list>
		<view v-if="list.length < 1">
			<u-empty text="没有数据"
			 mode="search"
			 class="nodate"></u-empty>
		</view>
	</view>
	
</template>

<script>
	// import { getOpenId } from "../../utils/auth.js"
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	import {
		getAttentionList,
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
		onShow: function(options) {
			uni.startPullDownRefresh();
			// this.openid = getOpenId()
			// console.log(this.openid,"ASDASDASDDASDADSDASASDSDSD")
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
			this.getList();
		},
		methods: {
			//获取关注列表
			getList() {
				let params = {
					current: this.current,
					size: this.size
				}
				getAttentionList(params).then(res => {
					// console.log(res)
					if (res.data.follow != null) {
						this.list.push.apply(this.list, res.data.follow)
						this.current++;
					}else{
						uni.showModal({
						    title: '提示',
						    content: '您还没有关注对象哦',
						    success: function (res) {
						        if (res.confirm) {
						            console.log('用户点击确定');
						        } else if (res.cancel) {
						            console.log('用户点击取消');
						        }
						    }
						});
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
	/deep/.uni-list--border-bottom{
		height:0 !important;
	}
	.nodate {
		background-color: #fff;
		min-height: 800rpx;
	}
</style>
