<template>
	<!-- 推荐页面文章展示组件 -->
	<view class="recommend-component">
		<!-- 文章标题 -->
		<articleTitle :articleTitle="articleInfo.title" :articleId="articleInfo.id" :commentTotal="articleInfo.commentTotal" />
		<!-- 作者信息 -->
		<attentionAndFansCell v-if="notshow" :aid="articleInfo.create_by" :nickname="articleInfo.nickname"
		 :avatarPath="articleInfo.avatarPath" :isFollow="articleInfo.isFollow"></attentionAndFansCell>
		<!-- 文章内容 -->
		<articleContent :articleContent="articleInfo.describe" :articleImg="articleInfo.image" :commentTotal="articleInfo.commentTotal"
		 :id="articleInfo.id" />
		<!-- 文章互动数量-->
		<articleInteract :favourTotal="articleInfo.favourTotal" :commentTotal="articleInfo.commentTotal" :collectTotal="articleInfo.collectTotal"
		 v-if="isArticleInteract"></articleInteract>
	</view>
</template>

<script scoped>
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
			},
			isArticleInteract: {
				type: Boolean,
				default: true
			},
			notshow:{
				type: Boolean,
				default: false
			}
		},
		components: {
			articleTitle,
			articleContent,
			articleInteract,
			attentionAndFansCell
		},
		methods: {

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
