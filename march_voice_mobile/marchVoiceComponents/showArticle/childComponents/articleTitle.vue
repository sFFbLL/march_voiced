<template>
	<!-- 文章标题 -->
	<view class="article-title"
	 @click="articleDetails">{{articleTitle}}</view>
</template>

<script>
	export default {
		data() {
			return {

			};
		},
		props: {
			articleId: {
				type: Number,
				default: null
			},
			articleTitle: {
				type: String,
				default: ""
			}
		},
		components: {

		},
		methods: {
			articleDetails() {
				// console.log(1)
				// 跳转到编辑页面
				uni.navigateTo({
					url: '../articleDetails/index?id='+this.articleId
				})
			}
		},
		mounted() {

		}
	}
</script>

<style>
	.article-title {
		color: #404040;
		font-size: 32rpx;
		font-family: "微软雅黑";
		text-align: left;
		font-weight: bold;
		overflow: hidden;
		-webkit-line-clamp: 2;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		line-height: 150%;
	}
</style>
