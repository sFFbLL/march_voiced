<template>
	<uni-list class="cell-box">
		<attentionAndFansCell v-for="(item, index) in list" :key="item.id" :id="item.id"></attentionAndFansCell>
		<uniLoadMore></uniLoadMore>
	</uni-list>
</template>

<script>
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue';
	import uniLoadMore from "@/components/uni-load-more/uni-load-more.vue"
	export default {
		data() {
			return {
				list: [],
				list1: [{
						id: '1'
					},
					{
						id: '2'
					},
					{
						id: '3'
					},
					{
						id: '4'
					},
					{
						id: '5'
					},
					{
						id: '6'
					},
					{
						id: '7'
					},
					{
						id: '8'
					}, 
					{
						id: '9'
					},
					{
						id: '10'
					},
					{
						id: '11'
					},
					{
						id: '12'
					},
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
			setTimeout(function() {
				// that.list.push.apply(that.list,that.list1)
				that.list.push.apply(that.list, that.list1)
			}, 1000);
			uni.startPullDownRefresh();
		},
		onPullDownRefresh() {
			console.log('refresh');
			setTimeout(function() {
				uni.stopPullDownRefresh();
			}, 1000);
		},
		onReachBottom() {
			this.list.push.apply(this.list, this.list1)
		},
		methods: {
			inToMinePage(id) {
				alert(id)
			}
		}
	}
</script>

<style scoped>
	.cell-box {}
</style>
