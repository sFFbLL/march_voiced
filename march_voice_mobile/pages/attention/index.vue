<template>
	<uni-list class="list-item">
		<attentionAndFansCell class="item"
		 v-for="(item, index) in list"
		 :key="item.id"
		 :id="item.id"
		 :showDteial="showDteial"
		 :showDateAndNumber="showDateAndNumber"
		 :isAttention="item.isAttention"
		 @change="change(index)"><span slot="hasAttention">已关注</span></attentionAndFansCell>
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
				showDateAndNumber:true,
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
				this.list[index].isAttention = !this.list[index].isAttention
			}
		}
	}
</script>

<style scoped>

</style>
