<template>
	<view>
		<!-- 三月圈头部 -->
		<view class="titlebg">
			<view class="title">
				<span class="marchcicrle">三月圈</span>
				<span class="member">成员:{{marchCircleInfo.personTotal}}</span><span class="total">条数:{{marchCircleInfo.articleTotal}}</span>

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
			<!-- 间隔槽 -->
			<view v-for="item in ideasListCopy">
				<u-gap height="30" bg-color="#f5f5f5"></u-gap>
				<view class="ideacontent">
					<!-- 用户头像公共组件 -->
					<attentionAndFansCell :aid="item.create_by" :nickname="item.nickname" :avatarPath="item.avatarPath" :isFollow="item.isFollow"></attentionAndFansCell>
					<!-- 想法的文字部分 -->
					<articleContent :articleContent="item.content" :commentTotal="item.commentTotal" :isIdea="true" :id="item.id"></articleContent>
					<!-- 想法的图片部分组件 -->
					<imageAdaptation :imgList="item.imageList"></imageAdaptation>
					<!-- 点赞表情组件+评论 -->
					<emojiControl :faceTotals="item.faceTotal" :likeTotals="item.likeTotal" :favourTotals="item.favourTotal"
					 :commentTotals="item.commentTotal" :id="item.id"></emojiControl>
				</view>
			</view>
		</view>
		<view v-if="isLoadMore">
			<uni-load-more :status="loadStatus"></uni-load-more>
		</view>
		<!-- 微信分享遮罩层 -->
		<view class="mcover" @click="isshow()" :style="{display:mcoverDisplay}">
			<image src="https://oscimg.oschina.net/oscnet/fd2170a448e37826ae9f4d7088f287b8f24.jpg" />
		</view>
		<!-- 发布三月圈悬浮按钮 -->
		<view v-if="!sanyueMumber" @click="publish()" class="publishbtn">
			<uni-icons class="addicon" type="plusempty" size="43" color="white"></uni-icons>
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
				marchCircleInfo: {},
				ideasList: [],
				ideasListCopy: [],
				pullDownRefresh:false
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
				this.isLoadMore = true;
				this.current += 1;
				this.getCircleList();
			}
		},
		// 下拉刷新
		onPullDownRefresh() {
			this.current = 1; //推荐当前页数，
			this.ideasList = [];
			setTimeout(function() {
				uni.stopPullDownRefresh();
			}, 2000);
			this.getCircleList();
			this.getMarchCircleInfo()
		},
		created() {
			// 获取三月基本信息接口
			getMarchCircleInfo().then(res => {
				this.marchCircleInfo = res.data;
				if (this.marchCircleInfo.is_march == 0) {
					this.isjoin = '加入';
				} else if (this.marchCircleInfo.is_march == 1) {
					this.isjoin = '已加入'
				} else if (this.marchCircleInfo.is_march == 2) {
					this.isjoin = '审核中'
					this.disabledJoin = true;
				}
			})

			this.getCircleList();
		},

		methods: {
			// 获取三月圈列表
			// 获取想法列表接口
			getCircleList() {
				let params = {
					current: this.current,
					size: this.size
				}
				marchCircleList(params).then(res => {
					if (res.data) {
						if (this.current === 1) {
							this.isLoadMore = false;
						} else {
							setTimeout(function() {
								this.loadStatus = 'loading';
							}, 2000);
							this.isLoadMore = false;
						}
						this.ideasList = [...this.ideasList, ...res.data];
						this.ideasListCopy=this.ideasList;
					} else {
						this.loadStatus = 'nomore';
					}

				})


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

	>>>.article-content {
		min-height: 40rpx;
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
		left: 134rpx;
		top: 30rpx;
		width: 144rpx;
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
