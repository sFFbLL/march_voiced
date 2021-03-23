<template>
	<!-- 文章内容显示 -->
	<view class="article-content" @click="toDetails">
		<view>
			<text class="article-text" v-html="textContent" ref="articleContent">{{textContent}}</text>

		</view>
		<view v-if="articleImg != ''" class="article-img">
			<image :src="articleImg" mode="aspectFill"></image>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				textContent: ""
			};
		},
		props: {
			id: {
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
			},
			isIdea: {
				type: Boolean,
				default: false
			},
			commentTotal: {
				type: Number,
				default: 0
			}
		},
		components: {

		},
		methods: {
			toDetails() {
				if (this.isIdea) {
					// 跳转详情页面
					uni.navigateTo({
						url: '../ideaDetails/index?id=' + this.id + '&commentTotal=' + this.commentTotal
					})
				} else {

					// 跳转到页面
					uni.navigateTo({
						url: '../articleDetails/index?id=' + this.id + '&commentTotal=' + this.commentTotal
					})
				}
			}
		},
		created() {
			this.textContent = this.articleContent;
		},
		mounted() {
			this.textContent = this.$refs.articleContent.$el.innerText.replace(/\ +/g, "").replace(/[\r\n]/g, "");
		},
	}
</script>

<style scoped>
	.article-content {
		margin-top: 18rpx;
		display: flex;
		justify-content: space-between;
		font-size: 28rpx;
		font-family: PingFangSC;
		color: #404040;
		min-height: 140rpx;
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
		/* margin-right: 30rpx; */

	}

	.article-img image {
		width: 180rpx;
		height: 140rpx;
		background-color: rgba(238, 238, 238, 1);
		border-radius: 4px;
	}
</style>
