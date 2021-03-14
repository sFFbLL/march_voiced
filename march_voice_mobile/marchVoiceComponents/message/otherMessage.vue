<template>
	<view class="wrap">
		<image v-if="otherList.status==0" style="width:30rpx; height: 30rpx;margin: 0 10rpx;" src="../../static/img/success.png"
		 mode="scaleToFill"></image>
		<image v-if="otherList.status==1" style="width:30rpx; height: 30rpx;margin: 0 10rpx;" src="../../static/img/fail.png"
		 mode="scaleToFill"></image>
		<text v-if="otherList.type==0">您的文章</text>
		<navigator :url="url" v-if="otherList.type==0" style="color: #007AFF;display: inline-block;">《{{otherList.title}}》</navigator>
		<text class="tipWords">{{tipWords}}</text>


		<view class="time"> <text>{{createDate}}</text></view>

	</view>
</template>

<script>
	export default {
		props: {
			otherList: {
				type: Object
			}
		},
		data() {
			return {
				tipWords: "",
				url: "../../pages/articleDetails/index?id="
			}
		},
		computed: {
			createDate: function() {
				let date = new Date();

				date.setTime(this.otherList.createTime * 1000);

				console.log(date.toLocaleDateString());
				return date.toLocaleDateString()
			}
		},
		created() {
			if (this.otherList.type == 0) {
				this.url = "../../pages/articleDetails/index?id=" + this.otherList.articleId;
				if (this.otherList.status == 0) {
					this.tipWords = "通过审核，已为您自动发布。";

				} else if (this.otherList.status == 1) {
					this.tipWords = "含有敏感信息，请进行修正后重新发布。"
				}
			} else if (this.otherList.type == 1) {
				if (this.otherList.status == 0) {
					this.tipWords = "恭喜你成功加入三月圈，快去看看吧。"
				} else if (this.otherList.status == 1) {
					this.tipWords = "由于权限不足，未能成功加入三月圈。"
				}
			}
		}

	}
</script>

<style scoped>
	.wrap {

		padding: 50rpx 10rpx;

	}

	.tipWords {}

	.time {
		margin-top: 10rpx;
		font-size: 20rpx;
		color: #616367;
		margin-left: 10rpx;
	}
</style>
