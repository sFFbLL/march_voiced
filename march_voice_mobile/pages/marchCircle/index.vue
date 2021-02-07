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
				<image @click="share()" class="wxlogo" src="../../static/img/wxlogo.png"></image>
				<button v-if="sanyueMumber" class="join" :disabled="disabledJoin" :style="{fontSize:fontSize+'rpx'}" @click="join()">{{isjoin}}</button>
				<u-toast ref="uToast" />

			</view>


		</view>

		<!-- 三月圈内容 -->
		<view class="wrap">

			<u-gap height="30" bg-color="#f5f5f5"></u-gap>
			<view class="ideacontent">

				<!-- 用户头像公共组件 -->
				<attentionAndFansCell :showDteial="false"></attentionAndFansCell>
				<!-- 想法的文字部分 -->
				<view class="wordscontent">
					<text class="content" space="ensp" v-html="ideasList.content"></text>

				</view>
				<!-- 想法的图片部分组件 -->
				<imageAdaptation :imgList="imgList"></imageAdaptation>
		<!-- 		<view class="allImage">
					<view class="images" v-for="(item,index) in ideasList.imgList" :key="index">
						<image class="oneimg" :src="item" mode="aspectFill" :style="{width:imgWidth+'rpx',height:imgHeight+'rpx'}"></image>
					</view>
				</view> -->
				
				<!-- 点赞表情组件 -->
				<emojiControl :emojiList="emojiList"></emojiControl>
			</view>
		</view>

		<!-- 发布三月圈悬浮按钮 -->
		<uni-fab v-if="sanyueMumber" :pattern="pattern" horizontal="right" @fabClick="publish()" class="publishbtn"></uni-fab>
	</view>
</template>

<script>
	import {
		getMarchCircleInfo,
		marchCircleList,
		joinMarchCircle
	} from '../../utils/api/marchCircle-api.js'
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import emojiControl from '../../marchVoiceComponents/marchCircle/emojiControl.vue'
	import imageAdaptation from '../../marchVoiceComponents/marchCircle/imageAdaptation.vue'
	import *as jwx from '../../utils/jws.js'
	export default {
		data() {
			return {
				sanyueMumber: true,
				isjoin: '加入',
				disabledJoin: false,
				fontSize: 28,
				emojiPosition01: 30,
				pattern: {
					buttonColor: '#2a82e4',
					width: '100rpx',
					height: '100rpx'
				},
				marchCircleInfo: {
					people: 1222,
					articles: 2048,
					ismarch: 1,
					brief: "啊士大夫艰苦的萨拉就",
				},
				emojiList: {
					faceTotal:3,
					likeTotal: 6,
					favourTotal: 10,
					commentTotal: 0,
				},
				imgList:[
					'../../static/img/cat.jpg', '../../static/img/cat.jpg'
				],
				ideasList: {
					content: "<span>今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六...</span>",
					upDateTime: '2020/12/12',
					faceTotal: 2,
					likeTotal: 3,
					favourTotal: 42,
					commentTotal: 0,
					imgList:[
						'../../static/img/cat.jpg', '../../static/img/cat.jpg'
					],
					user: {
						userName: "xianer",
						userId: 0,
						userImage: '',
						isFollow: 0
					}
				}


			}
		},
		components: {
			emojiControl,
			attentionAndFansCell,
			imageAdaptation
		},
		created() {
			// 获取三月基本信息接口
			// getMarchCircleInfo().then(res=>{
			// 	this.marchCircleInfo=res.data;
			// })
			if (this.marchCircleInfo.ismarch == 0) {
				this.sanyueMumber = false;
			}

			// 获取想法列表接口
			// marchCircleList().then(res=>{
			// 	this.ideasList=res.data;
			// 	this.emojiList.faceTotal=this.ideasList.faceTotal;
			// 	this.emojiList.favourTotal=this.ideasList.favourTotal;
			// 	this.emojiList.likeTotal=this.ideasList.likeTotal;
			// 	this.emojiList.commentTotal=this.ideasList.commentTotal;
			// 	this.imgList=this.ideasList.imgList;

			// })

		

		},
		mounted() {
		},
		methods: {
			// 调用微信接口分享内容
			share(){
				jwx.configWeiXin()
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
				// 调接口
				joinMarchCircle().then(res=>{
					console.log(res)
				})
			},
			publish() {
				// 跳转到编辑页面
			}
		}
	}
</script>

<style>
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
		height: 100vh;
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
