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
			}
		},
		components: {

		},
		methods: {
			toDetails() {
				if (this.isIdea) {
					// 跳转详情页面
					uni.navigateTo({
						url: '../ideaDetails/index?id=' + this.id
					})
				} else {
					// 跳转到编辑页面
					uni.navigateTo({
						url: '../articleDetails/index?id=' + this.id
					})
				}
			}
		},
		created() {
			this.textContent = this.articleContent;
			console.log(this.textContent)
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
