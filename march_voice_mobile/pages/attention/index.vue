<template>
	<uni-list class="list-item">
		<attentionAndFansCell class="item" v-for="(item, index) in list" :key="index" @change="change(item.user.id)" @inToPageMine="inToPageMine(item.user.id)"
		 :id="item.user.id" :nickname="item.user.nickname" :isFollow="item.isFollow" :avatarPath="item.user.avatarPath">
			<!-- <span slot="afterNicknameText">关注了你</span> -->
			<span slot="middleText">
				<span class="inner-text">关注 {{item.followTotal}}</span>
				<span class="inner-text">粉丝 {{item.fansTotal}}</span>
				<span class="inner-text">文章 {{item.articleTotal}}</span>
			</span>
			<span slot="underText">{{item.signature}}</span>
		</attentionAndFansCell>
		<uniLoadMore></uniLoadMore>
	</uni-list>
</template>

<script>
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
				// list1: [{
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }, {
				// 	"fansTotal": "10",
				// 	"followTotal": "10",
				// 	"signature": "asdasdasdasdasd",
				// 	"articleTotal": "10",
				// 	"user": {
				// 		"id": 1,
				// 		"avatarPath": "../../static/img/1.jpg",
				// 		"nickname": "hhh"
				// 	}
				// }],
				current: 1,
				size: 10
			}
		},
		components: {
			attentionAndFansCell,
			uniLoadMore
		},
		// created() {
		// 	this.getList()
		// },
		// 进入时加载
		onLoad: function(options) {
			let that = this;
			this.list = [];
			// setTimeout(function() {
			// 	that.list.push.apply(that.list, that.list1)
			// }, 1000);
			uni.startPullDownRefresh();
		},
		// 下拉刷新
		onPullDownRefresh() {
			let that = this;
			that.list = [];
			setTimeout(function() {
				// console.log("ASDASDASDASASDASD")
				// that.list.push.apply(that.list, that.list1)
				that.getList()
				uni.stopPullDownRefresh();
			}, 1000);
		},
		// 触底加载
		onReachBottom() {
			this.getList()
		},
		methods: {
			//获取关注列表
			getList() {
				let params = {
					current: this.current,
					size: this.size
				}
				getAttentionList(params).then(res => {
					console.log(res)
					if (res.code === 0 && res.data.length > 0) {
						this.list.push.apply(this.list, res.data)
						this.current++;
					}
				})
			},
			// 跳转页面
			inToPageMine(id) {
				console.log("ASD")
				uni.switchTab({
					url: '/pages/personalCenter/index'
				});
			},
			// 按钮样式切换
			change(index) {
				
				switch (this.list[index].isFollow) {
					case 0:
						changeStatus()
						this.list[index].id = 1;
						break;
					case 1:
						this.list[index].isFollow = 0;
						break;
				}
			}
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
