<template>
	<!-- 文章互动 -->
	<view class="article-interact">
		<view class="interact-item">
			<slot name="before"></slot>
			{{favourTotal}}<text>赞</text><text class="point">·</text>
			{{commentTotal}}<text>评论</text>
			<text v-if="collectTotal != null"
			 class="collectTotal">
				<text class="point">·</text>
				{{collectTotal}}<text>收藏</text>
			</text>
			<text v-if="dateTime != null"
			 class="dateTime">
				<text class="point">·</text>
				{{format(dateTime)}}
			</text>
			<slot name="after"></slot>
		</view>
	</view>
</template>

<script scoped>
	import moment from 'moment';
	export default {
		data() {
			return {

			};
		},
		props: {
			favourTotal: {
				type: Number,
				default: 0
			},
			commentTotal: {
				type: Number,
				default: 0
			},
			collectTotal: {
				type: Number,
				default: null
			},
			dateTime: {
				type: Number,
				default: null
			}
		},
		components: {

		},
		methods: {
			format(dateTime) {
				let stamp = new Date(dateTime);
				let time = moment(stamp).format('YYYY-MM-DD HH:mm:ss');
				return time;
			}
		},
		mounted() {}
	}
</script>

<style>
	.article-interact {
		margin-top: 19rpx;
		color: rgba(153, 153, 153, 100);
		font-size: 22rpx;
		font-family: PingFangSC;
	}

	.interact-item uni-text {
		margin-left: 8rpx;
	}

	.point {
		margin: 0 8rpx;
	}

	.upDateTime {
		display: inline;
	}
</style>
