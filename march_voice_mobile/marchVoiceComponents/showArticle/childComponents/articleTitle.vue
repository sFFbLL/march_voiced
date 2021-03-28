<template>
	<!-- 文章标题 -->
	<view class="article-title" @click="articleDetails">
		<view style="display: flex;">
			<text class="article-text" v-html="articleTitle"></text>
			<text v-if="articleStatus==2" style="width: 102rpx; color: gray; ">审核中</text>
		</view>
	</view>
</template>

<script scoped>
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
			articleStatus: {
				type: Number,
				default: null
			},
			articleTitle: {
				type: String,
				default: ""
			},
			commentTotal: {
				type: Number,
				default: 0
			},
			type:{
				type: String,
				default:null
			}

		},
		methods: {
			articleDetails() {
				if(this.type === "draft"){
					uni.navigateTo({
						url: '../../pages/publish/richTextEditor?type=draft&id=' + this.articleId + '&commentTotal=' + this.commentTotal,
					})
					return;
				}
				// 跳转到页面
				uni.navigateTo({
					url: '../articleDetails/index?id=' + this.articleId + '&commentTotal=' + this.commentTotal
				})

			}
		},
		mounted() {

		}
	}
</script>

<style>
	.article-title .article-text {
		color: #404040;
		font-size: 34rpx;
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

	.article-text {
		display: inline-block;
		width: 100%;
	}
</style>
