<template>
	<!-- 文章内容显示 -->
	<view class="article-content">
		<view class="article-text">{{articleContent}}</view>
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

		},
		mounted() {}
	}
</script>

<style>
	.article-content {
		margin-top: 18rpx;
		display: flex;
		font-size: 27rpx;
		font-family: PingFangSC;
		color: #404040;
	}

	.article-text {
		display: inline-block;
	}

	.article-img {
		display: inline-block;
		margin-left: 30rpx;
	}

	.article-img image {
		width: 170rpx;
		height: 130rpx;
		background-color: rgba(238, 238, 238, 1);
		border-radius: 4px;
	}
</style>
