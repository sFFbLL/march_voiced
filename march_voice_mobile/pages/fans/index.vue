<template>
	<uni-list class="cell-box">
		<attentionAndFansCell v-for="(item, index) in list" :key="item.id" :id="item.id" :showDteial="showDteial"
		 :isAttention="item.isAttention" @change="change(index)"><span slot="hasAttention">互相关注</span></attentionAndFansCell>
		<uniLoadMore></uniLoadMore>
	</uni-list>
</template>

<script>
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	export default {
		data() {
			return {
				showDteial: false,
				list: [],
				list1: [{
						id: '1',
						isAttention: true
					},
					{
						id: '2',
						isAttention: true
					},
					{
						id: '3',
						isAttention: true
					},
					{
						id: '4',
						isAttention: true,

					},
					{
						id: '5',
						isAttention: true
					},
					{
						id: '6',
						isAttention: true
					},
					{
						id: '7',
						isAttention: true
					},
					{
						id: '8',
						isAttention: true
					},
					{
						id: '9',
						isAttention: true
					},
					{
						id: '10',
						isAttention: true
					},
					{
						id: '11',
						isAttention: true
					},
					{
						id: '12',
						isAttention: true
					},
					{
						id: '13',
						isAttention: true
					},
					{
						id: '14',
						isAttention: true
					},
					{
						id: '15',
						isAttention: true
					},
					{
						id: '16',
						isAttention: true
					}
				]
			}
		},
		components: {
			attentionAndFansCell,
			uniLoadMore
		},
		onLoad: function(options) {
			let that = this;
			this.list = [];
			// setTimeout(function() {
			// 	that.list.push.apply(that.list, that.list1)
			// }, 1000);
			uni.startPullDownRefresh();
		},
		onPullDownRefresh() {
			let that = this;
			setTimeout(function() {
				that.list = [];
				that.list.push.apply(that.list, that.list1)
				uni.stopPullDownRefresh();
			}, 1000);
		},
		onReachBottom() {
			this.list.push.apply(this.list, this.list1)
		},
		methods: {
			inToMinePage(id) {
				alert(id)
			},
			change(index) {
				// console.log(index)
				// console.log(this.list[index])
				this.list[index].isAttention = !this.list[index].isAttention
				// console.log(this.list[index].isAttention)
			}
		}
	}
</script>

<style scoped>
	::v-deep .right-button {
		width: 160rpx;
	}
</style>
