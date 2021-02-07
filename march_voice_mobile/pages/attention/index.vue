<template>
	<uni-list class="list-item">
		<attentionAndFansCell class="item" v-for="(item, index) in list" :key="item.id" :showDteial="showDteial"
		 :showSelfIntrouduce="showSelfIntrouduce" :showDate="showDate" :user="list1[index]" @change="change(index)"><span
			 slot="doSomeThing">关注了你</span><span slot="hasAttention">已关注</span></attentionAndFansCell>
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
				showSelfIntrouduce: true,
				showDate: true,
				list: [],
				list1: [{
					id: "", // 用于事件
					imgUrl: require('../../static/img/1.jpg'), // 显示头像
					name: "昵称", // 显示昵称
					count: "0", //关注数等
					selfIntrouduce: "阿萨德",
					follow: true, // 判断是否互相关注
					count: "1",
					date: "2020-20-01"
				}]
			}
		},
		components: {
			attentionAndFansCell,
			uniLoadMore
		},
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
			setTimeout(function() {
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
				this.list[index].follow = !this.list[index].follow
			}
		}
	}
</script>

<style scoped>

</style>
