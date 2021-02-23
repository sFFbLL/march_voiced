<template>
	<view>
		<!-- 三月圈头部 -->
		<view class="titlebg">
			<view class="title">
				<span class="marchcicrle">三月圈</span>
				<span class="member">成员:{{marchCircleInfo.people}}</span><span class="total">条数:{{marchCircleInfo.articles}}</span>

				<p class="info">简介: {{marchCircleInfo.brief}}</p>
			</view>
			<view>
				<image @click="share()"
				 class="wxlogo"
				 src="../../static/img/wxlogo.png"></image>
				<button v-if="sanyueMumber"
				 class="join"
				 :disabled="disabledJoin"
				 :style="{fontSize:fontSize+'rpx'}"
				 @click="join()">{{isjoin}}</button>
				<u-toast ref="uToast" />

			</view>


		</view>

		<!-- 三月圈内容 -->
		<view class="wrap">
			<!-- 间隔槽 -->

			<view v-for="item in reverseIdeaList">
				<u-gap height="30"
				 bg-color="#f5f5f5"></u-gap>
				<view class="ideacontent">
					<!-- 用户头像公共组件 -->
					<attentionAndFansCell :nickname="item.nickname"
					 :avatarPath="item.avatarPath"
					 :isFollow="item.isFollow"></attentionAndFansCell>
					<!-- 想法的文字部分 -->
					<articleContent :articleContent="item.content"
					 :isIdea="true"
					 :id="item.id"></articleContent>
					<!-- 想法的图片部分组件 -->
					<imageAdaptation :imgList="item.imgList"></imageAdaptation>
					<!-- 点赞表情组件+评论 -->
					<emojiControl :faceTotals="item.faceTotal"
					 :likeTotals="item.likeTotal"
					 :favourTotals="item.favourTotal"
					 :commentTotals="item.commentTotal"
					 :id="item.id"></emojiControl>
				</view>
			</view>

		</view>
		<view v-show="isLoadMore">
			<uni-load-more :status="loadStatus"></uni-load-more>
		</view>
		<!-- 微信分享遮罩层 -->
		<view class="mcover"
		 @click="isshow()"
		 :style="{display:mcoverDisplay}">
			<image src="https://oscimg.oschina.net/oscnet/fd2170a448e37826ae9f4d7088f287b8f24.jpg" />
		</view>
		<!-- 发布三月圈悬浮按钮 -->
		<view v-if="sanyueMumber"
		 @click="publish()"
		 class="publishbtn">
			<uni-icons class="addicon"
			 type="plusempty"
			 size="43"
			 color="white"></uni-icons>
		</view>

	</view>
</template>

<script>
	import {
		getMarchCircleInfo,
		marchCircleList,
		joinMarchCircle,
		changeFavour
	} from '../../utils/api/marchCircle-api.js'
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import emojiControl from '../../marchVoiceComponents/marchCircle/emojiControl.vue'
	import imageAdaptation from '../../marchVoiceComponents/marchCircle/imageAdaptation.vue'
	import * as jwx from '../../utils/jws.js'
	import articleContent from '../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue'
	import {
		check
	} from '../../utils/checkUnRead.js'
	export default {
		data() {
			return {
				isLoadMore: false, //是否加载中
				loadStatus: 'loading',
				current: 1,
				size: 10,
				sanyueMumber: true,
				isjoin: '加入',
				disabledJoin: false,
				fontSize: 28,
				emojiPosition01: 30,
				mcoverDisplay: 'none',
				pattern: {
					buttonColor: '#2a82e4',
					width: '50rpx',
					height: '50rpx'
				},
				marchCircleInfo: {
					people: 1222,
					articles: 2048,
					ismarch: 1,
					brief: "啊士大夫艰苦的萨拉就",
				},
				ideasList: [
					// {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 0,
					// 	likeTotal: 3,
					// 	favourTotal: 0,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg'
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',

					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg'
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg'
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg'
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg'
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg'
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, {
					// 	id: 6,
					// 	content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					// 	updateTime: '2020/12/12',
					// 	faceTotal: 2,
					// 	likeTotal: 3,
					// 	favourTotal: 42,
					// 	commentTotal: 0,
					// 	imgList: [
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg', '../../static/img/cat.jpg', '../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 		'../../static/img/cat.jpg',
					// 	],
					// 	create_by: 1,
					// 	nickname: "仙儿",
					// 	avatarPath: "",
					// 	isFollow: 0
					// }, 
				]


			}
		},
		computed: {
			reverseIdeaList() {
				return this.ideasList.reverse();
			}
		},
		components: {
			emojiControl,
			attentionAndFansCell,
			imageAdaptation,
			articleContent
		},
		//上拉触底函数
		onReachBottom() {
			if (!this.isLoadMore) { //此处判断，上锁，防止重复请求
				this.isLoadMore = true
				this.current += 1
				this.getCircleList();
			}
		},
		created() {
			// 获取三月基本信息接口
			getMarchCircleInfo().then(res => {
				this.marchCircleInfo = res.data;
			})
			if (this.marchCircleInfo.ismarch == 0) {
				this.sanyueMumber = false;
			}
			let params = {
				current: this.current,
				size: this.size
			}

			this.getCircleList(params);
		},
		onShow() {
			check()
		},
		methods: {

			// 获取三月圈列表
			// 获取想法列表接口
			getCircleList(params) {
				let _this = this;
				marchCircleList(params).then(res => {
					_this.ideasList = [...this.ideasList, ...res.data];
					if (res.data.length <= _this.size) {
						_this.loadStatus = 'nomore';
					}

				})

				if (this.ideasList.length > 16) {
					_this.loadStatus = "nomore";
				} else if (this.current === 1) {
					_this.isLoadMore = false;
				} else {
					setTimeout(function () {
						_this.isLoadMore = false;
					}, 2000);
				}
			},
			// 调用微信接口分享内容
			share() {
				this.mcoverDisplay = 'block';
				jwx.configWeiXin();
			},
			// 控制遮罩层显示
			isshow() {
				this.mcoverDisplay = 'none';
			},
			// 申请加入三月圈
			join() {
				this.$refs.uToast.show({
					title: '操作成功,请等待审核',
					type: 'primaty'

				})
				this.isjoin = '审核中';
				this.fontSize = 21;
				this.disabledJoin = true;
				// 调申请加入三月圈接口
				joinMarchCircle().then(res => {
					console.log(res)
				})
			},
			publish() {
				// 跳转到编辑页面
				uni.navigateTo({
					url: '../ideaPublish/index'
				})
			}
		}
	}
</script>

<style>
	/* 发布想法的按钮 */
	.publishbtn {
		border-radius: 50%;
		width: 100rpx;
		height: 100rpx;
		display: flex;
		position: fixed;
		z-index: 222;
		top: 1080rpx;
		right: 30rpx;
		/* background-image: linear-gradient(to top, #4481eb 0%, #04befe 100%); */
		background-image: linear-gradient(-225deg, #22E1FF 0%, #1D8FE1 48%, #625EB1 100%);
	}

	.addicon {
		margin-left: 8rpx;
		margin-top: 5rpx;
	}

	/* 微信分享 遮罩层*/
	.mcover {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		background: rgba(0, 0, 0, 0.7);
		z-index: 20;
	}

	/* 想法文字部分样式 */
	.wordscontent {
		margin-top: 20rpx;
	}

	/* 公共头像组件样式 */
	>>>.attention-cell .flex-item {
		border: none;
		padding: 0;
		margin-top: 17rpx;
	}

	.content {
		overflow: hidden;
		-webkit-line-clamp: 3;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		font-size: 32rpx;
		color: #404040;
		text-align: left;
		line-height: 1.5;

	}

	.ideacontent {
		background-color: #FFFFFF;
		padding: 26rpx;

	}

	.idea {
		/* margin-top: 20rpx; */
		width: 100%;
		/* height: 75rpx; */
		/* line-height: 20rpx; */
		color: #101010;
		font-size: 14rpx;
		text-align: center;
		font-family: Arial;
		display: flex;
		border: 1rpx solid rgba(255, 255, 255, 100);
	}

	.wrap {
		background-color: #F5F5F5;
		/* height: 100vh; */
	}

	/* 头部样式 */

	.marchcicrle {
		font-size: 36rpx;
		width: 108rpx;
		height: 46px;
	}

	.member,
	.total {
		font-size: 28rpx;
		margin: 0 10rpx;
	}

	.info {
		margin-top: 30rpx;
		font-size: 32rpx;
	}

	.titlebg {
		display: flex;
		width: 100%;
		height: 320rpx;
		background-image: url(../../static/img/marchcicrlebg.jpg);
		background-position: center;
		background-size: 100% 100%;
	}

	.title {
		margin: 72rpx 25rpx;
		width: 416rpx;
		color: #FFFFFF;
		font-weight: bold;

	}

	.wxlogo {
		width: 70rpx;
		height: 70rpx;
		left: 50rpx;
		top: 100rpx;
	}

	.join {
		left: 143rpx;
		top: 30rpx;
		width: 120rpx;
		height: 54rpx;

		line-height: 55rpx;
		border-radius: 28rpx 28rpx 28rpx 28rpx;
		background-color: #404040;
		color: #FFFFFF;
		font-size: 28rpx;
		text-align: center;
		font-family: Arial;
		border: 0.5rpx solid #101010 inset;
	}
</style>
