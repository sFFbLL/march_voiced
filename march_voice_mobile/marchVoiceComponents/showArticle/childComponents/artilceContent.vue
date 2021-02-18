<template>
	<!-- 文章内容显示 -->
	<view class="article-content"
	 @click="articleDetails">
		<view>
			<text class="article-text"
			 v-html="articleContent"></text>
		</view>
		<view v-if="articleImg != ''"
		 class="article-img">
			<image :src="articleImg"
			 mode="aspectFill"></image>
		</view>
	</view>
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
			articleContent: {
				type: String,
				default: ""
			},
			articleImg: {
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
					url: '../articleDetails/index?id=' + this.articleId
				})
			}
		},
		mounted() {}
	}
</script>

<style scoped>
	.article-content {
		margin-top: 18rpx;
		display: flex;
		font-size: 28rpx;
		font-family: PingFangSC;
		color: #404040;
	}

	.article-text {
		overflow: hidden;
		-webkit-line-clamp: 3;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		line-height: 150%;
	}

	.article-img {
		display: inline-block;
		margin-left: 30rpx;
	}

	.article-img image {
		width: 170rpx;
		height: 100%;
		background-color: rgba(238, 238, 238, 1);
		border-radius: 4px;
	}
</style>
