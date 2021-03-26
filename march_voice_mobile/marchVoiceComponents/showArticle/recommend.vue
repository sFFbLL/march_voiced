<template>
	<!-- 推荐页面文章展示组件 -->
	<view class="recommend-component">
		<!-- 文章标题 -->
		<articleTitle :articleTitle="articleInfo.title" :articleStatus="articleInfo.status" :articleId="articleInfo.id"
		 :commentTotal="articleInfo.commentTotal" />
		<!-- 作者信息 -->
		<attentionAndFansCell v-if="notshow" :aid="articleInfo.create_by" :nickname="articleInfo.nickname" :avatarPath="articleInfo.avatarPath"
		 :isFollow="articleInfo.isFollow"></attentionAndFansCell>
		<!-- 文章内容 -->
		<articleContent :articleContent="articleInfo.describe" :articleImg="articleInfo.image" :commentTotal="articleInfo.commentTotal"
		 :id="articleInfo.id" />
		<!-- 文章互动数量-->
		<articleInteract :favourTotal="articleInfo.favourTotal" :commentTotal="articleInfo.commentTotal" :collectTotal="articleInfo.collectTotal"
		 v-if="isArticleInteract"></articleInteract>
		<view v-if="draftsCreatime">{{format(draftsCreatime)}}</view>
	</view>
</template>

<script scoped>
	import articleTitle from "./childComponents/articleTitle.vue"
	import articleContent from "./childComponents/artilceContent.vue"
	import articleInteract from "./childComponents/articleInteract.vue"
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'

	import moment from 'moment';
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
			notshow: {
				type: Boolean,
				default: false
			},
			draftsCreatime: {
				type: Number,

			}
		},
		
		components: {
			articleTitle,
			articleContent,
			articleInteract,
			attentionAndFansCell
		},
		methods: {
			// 把时间戳转换为正确格式
			format(dateTime) {
				let stamp = new Date(dateTime);
				let time = moment(stamp).format('YYYY-MM-DD HH:mm:ss');
				return time;
			}
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
