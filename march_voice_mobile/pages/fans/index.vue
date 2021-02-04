<template>
	<uni-list class="cell-box">
		<attentionAndFansCell v-for="(item, index) in list" :key="item.id" :id="item.id" :showDteial="showDteial"
		 :isAttention="item.isAttention" @change="change(index)"></attentionAndFansCell>
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
						isAttention: false
					},
					{
						id: '2',
						isAttention: false
					},
					{
						id: '3',
						isAttention: false
					},
					{
						id: '4',
						isAttention: false,

					},
					{
						id: '5',
						isAttention: false
					},
					{
						id: '6',
						isAttention: false
					},
					{
						id: '7',
						isAttention: false
					},
					{
						id: '8',
						isAttention: false
					},
					{
						id: '9',
						isAttention: false
					},
					{
						id: '10',
						isAttention: false
					},
					{
						id: '11',
						isAttention: false
					},
					{
						id: '12',
						isAttention: false
					},
					{
						id: '13',
						isAttention: false
					},
					{
						id: '14',
						isAttention: false
					},
					{
						id: '15',
						isAttention: false
					},
					{
						id: '16',
						isAttention: false
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
	.cell-box {
		/* padding: 20rpx; */
	}
</style>
