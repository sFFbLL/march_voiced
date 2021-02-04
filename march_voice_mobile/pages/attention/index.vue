<template>
	<uni-list class="list-item">
		<attentionAndFansCell class="item"
		 v-for="(item, index) in list"
		 :key="item.id"
		 :id="item.id"
		 :showDteial="showDteial"
		 :isAttention="item.isAttention"
		 @change="change(index)"></attentionAndFansCell>
		<uniLoadMore></uniLoadMore>
	</uni-list>
</template>

<script>
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	export default {
		data() {
			return {
				showDteial: true, //是否展示粉丝数关注数
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
		// 进入时加载
		onLoad: function (options) {
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
			setTimeout(function () {
				that.list = [];
				that.list.push.apply(that.list, that.list1)
				uni.stopPullDownRefresh();
			}, 1000);
		},
		// 触底加载
		onReachBottom() {
			this.list.push.apply(this.list, this.list1)
		},
		methods: {
			// 跳转页面
			inToMinePage(id) {
				alert(id)
			},
			// 按钮样式切换
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

</style>
