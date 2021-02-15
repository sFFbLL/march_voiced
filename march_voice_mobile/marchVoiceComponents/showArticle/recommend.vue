<template>
	<!-- 推荐页面文章展示组件 -->
	<view class="recommend-component">
		<!-- 文章标题 -->
		<articleTitle :articleTitle="articleInfo.title"
		 :articleId="articleInfo.articleId" />
		<!-- 作者信息 -->
		<attentionAndFansCell :userInfo="articleInfo.user"
		 :showDteial="false"
		 :showIntrouduce="false"
		 :showDate="false"
		 :user="articleInfo.user"><span slot="hasAttention">已关注</span></attentionAndFansCell>
		<!-- 文章内容 -->
		<articleContent :articleContent="articleInfo.content"
		 :articleImg="articleInfo.articleImg"
		 :articleId="articleInfo.articleId" />
		<!-- 文章互动数量-->
		<articleInteract :favourTotal="articleInfo.favourTotal"
		 :commentTotal="articleInfo.commentTotal"
		 :collectTotal="articleInfo.collectTotal"></articleInteract>
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
