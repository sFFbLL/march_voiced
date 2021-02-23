<template>
	<view class="search-article-item">
		<articleTitle :articleTitle="articleInfo.title"
		 :articleId="articleInfo.articleId"></articleTitle>

		<articleContent :articleContent="articleInfo.content"
		 :id="articleInfo.id"></articleContent>

		<!-- 文章互动数量-->
		<articleInteract :favourTotal="articleInfo.favourTotal"
		 :commentTotal="articleInfo.commentTotal">
			<text class="nickname"
			 slot="before">
				{{articleInfo.nickname}}
			</text>
			<text class="time"
			 slot="after">
				{{articleInfo.createTime}}
			</text>
		</articleInteract>
	</view>
</template>

<script>
	import articleTitle from "./childComponents/articleTitle.vue"
	import articleContent from "./childComponents/artilceContent.vue"
	import articleInteract from "./childComponents/articleInteract.vue"
	export default {
		data() {
			return {

			};
		},
		props: {
			articleInfo: {
				type: Object,
				default: {}
			}
		},
		components: {
			articleTitle,
			articleContent,
			articleInteract
		},
		methods: {

		},
		mounted() {

		}
	}
</script>

<style scoped>
	.search-article-item {
		padding: 20rpx 0;
		overflow: hidden;
		border-bottom: 1px solid #f0f0f0;
	}

	>>>.article-title {
		line-height: 100%;
	}

	>>>.article-title .article-text {
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	>>>.article-content {
		color: #999;
		font-size: 24rpx;
		margin-top: 10rpx;
	}

	>>>.article-interact {
		font-size: 20rpx;
		margin-top: 10rpx;
	}

	>>>.interact-item uni-text {
		margin-left: 4rpx;
	}

	>>>.point {
		visibility: hidden;
	}

	.nickname {
		display: inline-block;
		padding-right: 20rpx;
		max-width: 200rpx;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		line-height: 100%;
	}
	
	.time{
		float: right;
		margin-right: 20rpx;
	}
</style>
