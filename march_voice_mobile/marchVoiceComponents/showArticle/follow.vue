<template>
	<!-- 推荐页面文章展示组件 -->
	<view class="recommend-component">
		<!-- 作者信息 -->
		<attentionAndFansCell :aid="articleInfo.create_by"
		 :nickname="articleInfo.nickname"
		 :avatarPath="articleInfo.avatarPath">
			<view slot="underText"
			 class="followMessage">
				<text v-if="articleInfo.status===0">
					发布
				</text>
				<text v-if="articleInfo.status===1">
					评论
				</text>
				<text v-if="articleInfo.status===2">
					收藏
				</text>
				<text v-if="articleInfo.status===3">
					转载
				</text>
				<text v-if="articleInfo.status===4">
					赞
				</text>
				<text>了该内容</text>
			</view>
		</attentionAndFansCell>
		<!-- 文章标题 -->
		<articleTitle :articleTitle="articleInfo.title"
		 :articleId="articleInfo.articleId" />
		<!-- 文章内容 -->
		<articleContent :articleContent="articleInfo.describe"
		 :articleImg="articleInfo.image"
		 :id="articleInfo.articleId"/>
		<!-- 文章互动数量-->
		<articleInteract :favourTotal="articleInfo.favourTotal"
		 :commentTotal="articleInfo.commentTotal"
		 :dateTime="articleInfo.createTime"></articleInteract>
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

<style scoped>
	.recommend-component {
		background-color: #fff;
		margin: auto;
		margin-bottom: 15rpx;
		padding: 30rpx 30rpx;
	}

	>>>.attention-cell .flex-item {
		border: none;
		padding: 0;
	}

	>>>.article-title {
		margin-top: 17rpx;
	}

	.followMessage {
		font-size: 26rpx;
		width: 420rpx;
		color: #969696;
	}
</style>
