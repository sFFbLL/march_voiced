<template>
	<!-- 推荐页面文章展示组件 -->
	<!-- 
	 upDateTime: "更新时间",
	 favourTotal: 1,
	 collectTotal: 1,
	 commentTotal: 1
	 -->
	<view class="recommend-component">
		<!-- 文章标题 -->
		<articleTitle :articleTitle="article.title" />
		<attentionAndFansCell :showDteial="false"></attentionAndFansCell>
		<!-- 文章内容 -->
		<articleContent :articleContent="article.content"
		 :articleImg="article.articleImg" />
		<!-- 文章互动数量-->
		<articleInteract :favourTotal="article.favourTotal"
		 :commentTotal="article.commentTotal"
		 :collectTotal="article.collectTotal"></articleInteract>
	</view>
</template>

<script>
	import articleTitle from "./childComponents/articleTitle.vue"
	import articleContent from "./childComponents/artilceContent.vue"
	import articleInteract from "./childComponents/articleInteract.vue"
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	export default {
		data() {
			return {
				article: {},
				userInfo: {}
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
			articleInteract,
			attentionAndFansCell
		},
		methods: {

		},
		mounted() {
			this.article = this.articleInfo.article;
			this.user = this.articleInfo.userInfo;
		}
	}
</script>

<style>
	.recommend-component {
		background-color: #fff;
		margin: auto;
		margin-bottom: 15rpx;
		padding: 30rpx 30rpx;
	}

	>>>.attention-cell .flex-item {
		border: none;
		padding: 0;
		margin-top: 17rpx;
	}
</style>
