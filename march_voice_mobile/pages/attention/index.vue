<template>
	<uni-list class="list-item">
<<<<<<< Updated upstream
		<attentionAndFansCell class="item" v-for="(item, index) in list" :key="index" @change="change(item.user.id)" @inToPageMine="inToPageMine(item.user.id)"
		 :id="item.user.id" :nickname="item.user.nickname" :isFollow="item.isFollow" :avatarPath="item.user.avatarPath">
			<!-- <span slot="afterNicknameText">关注了你</span> -->
=======
		<attentionAndFansCell class="item" v-for="(item, index) in list" :key="index" @change="change(index)" @inToPageMine="inToPageMine(item.id)"
		 :id="item.id" :nickname="item.nickname" :isFollow="item.isFollow" :avatarPath="item.avatarPath">
>>>>>>> Stashed changes
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
	import { getOpenId } from "../../utils/auth.js"
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
		onLoad: function(options) {
			uni.startPullDownRefresh();
			this.openid = getOpenId()
			console.log(this.openid,"ASDASDASDDASDADSDASASDSDSD")
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
			//获取关注列表
			getList() {
				let params = {
					current: this.current,
					size: this.size
				}
				getAttentionList(params).then(res => {
					console.log(res)
					if (res.data.Follow.length > 0) {
						this.list.push.apply(this.list, res.data.Follow)
						this.current++;
					}
				})
			},
			// 跳转页面
			inToPageMine(id) {
				 setTimeout(()=>{
					uni.$emit('getOthersId', {
						id: id
					})
				},1000)
				
				uni.switchTab({
					url: '/pages/personalCenter/index'
				});
			},
			// 按钮样式切换
			change(index) {
				
				let params = {
					id: this.list[index].id
				}

				changeStatus(params).then(res =>{
					// 手动更改状态，刷新后查询新数据
					switch (this.list[index].isFollow) {
						case 0:
							this.list[index].isFollow = 1
							break;
						case 1:
							this.list[index].isFollow = 0
							break;
					}
				})
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
	/deep/.uni-list--border-bottom{
		height:0 !important;
	}
</style>
