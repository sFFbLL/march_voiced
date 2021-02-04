<template>
	<!-- 文章标题 -->
	<view class="article-title">{{articleTitle}}</view>
</template>

<script>
	export default {
		data() {
			return {

			};
		},
		props: {
			articleTitle: {
				type: String,
				default: ""
			}
		},
		components: {

		},
		methods: {

		},
		mounted() {

		}
	}
</script>

<style>
	.article-title {
		color: #404040;
		font-size: 32rpx;
		font-family: PingFangSC;
		text-align: left;
		font-weight: bold;
	}
</style>
