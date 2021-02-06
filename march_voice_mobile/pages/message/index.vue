<template>
	<view class="message">
		<view class="header">
			<tabs class="tag-nav"
			 :tabs='tablist'
			 v-on:tabActive='tabActive' />
		</view>
		<view class="content">
			<view v-if="tabIndex==0">
				<view class="flex-item" v-for="(item,index) in interactList">
					<!-- 评论者头像 -->
					<a @click="inToPageMine"
					 class="left-img-box inner-box">
						<image class="inner-img" 
						:src="interactList[index].user.avatar" 
						 mode="aspectFill"></image>
					</a>
					<!-- 中部文字 -->
					<view class="middle-text-box inner-box"
					 @click="inToPageMine">
						<view class="inner-middle-box inner-box">
							<!-- 昵称盒子 -->
							<view>	
								<span class="inner-text-name">{{item.user.userName}}</span>
								<span class="inner-text" v-if="item.article.type==0">
									<span class="inner-text" v-if="item.article.status==1">评论</span>
									<span class="inner-text" v-if="item.article.status==2">收藏</span>
									<span class="inner-text" v-if="item.article.status==3">转发</span>
									<span class="inner-text" v-if="item.article.status==4">赞</span>了你的文章</span>
								<span class="inner-text" v-if="item.article.type==1">
										<span class="inner-text" v-if="item.article.status==1">评论了你的三月圈</span>
										<span class="inner-text" v-else>为你的三月圈贴了一个表情</span>
								</span>
							</view>
							<!-- 评论时间 -->
							<view class="inner-text-message greay-text">
								<span>{{item.article.createTime}}</span>
							</view>
							
						</view>
					</view>
					<!-- 评论，文章简略 -->
					<view class="comment-content">
						<!-- 评论内容 -->
						<view class="text-comment" v-if="item.article.status==1">{{item.article.commentId}}</view>
						<!-- 文章内容 -->
						<articleContent :articleContent="item.article.content"
						:articleImg="item.article.image" />
					</view>
				</view>
			</view>
			
			<view v-if="tabIndex==1">
				<view>
					<view class="flex-item" v-for="(item,index) in attentionList">
						
						
						
						
						
						
						
						
						<!-- 关注者头像 -->
						<a @click="inToPageMine"
						 class="left-img-box inner-box">
							<image class="inner-img" 
							:src="attentionList[index].user.avatar" 
							 mode="aspectFill"></image>
						</a>
						<!-- 中部文字 -->
						<view class="middle-text-box inner-box"
						 @click="inToPageMine">
							<view class="inner-middle-box inner-box">
								<!-- 昵称盒子 -->
								
								
								<view>	
									<span class="inner-text-name">{{item.user.username}}</span>
									<span class="inner-text">关注了你</span>
								</view>
								
								
								
								
								<!-- 评论时间 -->
								<view class="inner-text-message greay-text">
									<span>{{item.user.createTime}}</span>
								</view>
								
							</view>
						</view>
						
						<!-- 按钮盒子 -->
						<view class="right-button-box">
							<button v-show="isAttention === false"
							 class="right-button"
							 type="default-green"
							 :loading="isLoading"
							 :disabled="isDisabled"
							 @click.stop="changeBtn"><span>关 注</span></button>
							<button v-show="isAttention === true" class="right-button" type="default" :loading="isLoading" :disabled="isDisabled"
							 @click.stop="changeBtn">已关注</button>
						</view>
						
						
						
						
						
						
						
						
						
						
						
					</view>
				</view>
			</view>
			<view v-if="tabIndex==2">
				<view>
					<view v-for="(item,index) in otherList">
						
						<span class="inner-text" v-if="item.user.type==0">你投稿的文章审核
							<span class="inner-text" v-if="item.user.status==0">已通过。</span>
							<span class="inner-text" v-if="item.user.status==1">未通过！</span>
						</span>
						
						<span v-else>
							<span class="inner-text" v-if="item.user.status==0">您已加入三月圈</span>
							<span class="inner-text" v-if="item.user.status==1">您加入三月圈请求被驳回</span>
						</span>
				
				
				
					</view>
					
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import tabs from '../../marchVoiceComponents/tabCard.vue'
	import articleContent from "../../marchVoiceComponents/showArticle/childComponents/artilceContent.vue"
	
	export default {
		components: {
			tabs,
			articleContent
		},
		data() {
			return {
			//	isLoading: false, // 是否为加载中
			//	isDisabled: false // 是否禁用按钮点击
				tabIndex: '',
				tablist: [{
						index: 0,
						value: '互动消息',
						isActive: true
					},
					{
						index: 1,
						value: '关注',
						isActive: false
					},
					{
						index: 2,
						value: '系统消息',
						isActive: false
					}
				],
				// 互动消息
				interactList: [{
					article: {
						type: 0,
						status: 1,
						articleId: 1,
						commentId: "哈哈哈哈在法国夫是德国人头地方杠",
						image: require('static/img/2.jpg'),
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
						createTime: "更新时间"
					},
					user: {
						avatar: require('static/img/1.jpg'),
						userName: "张三",
						userId: 1,
					}
				},{
					article: {
						type: 0,
						status: 1,
						articleId: 1,
						commentId: "哈哈哈哈在法国夫是德国人头地方杠",
						image: '',
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",
						createTime: "更新时间"
					},
					user: {
						avatar: require('static/img/1.jpg'),
						userName: "张三",
						userId: 1,
					}
				}],
				
				
				//关注消息
				attentionList: [{
					user: {
						avatar:  require('static/img/1.jpg'),
						username: "李四士",
						userId: 2,
						createTime: "时间时间"
					}
				}],
				
				//其他消息
				otherList: [{
					user: {
						type: 1,
						status: 0,
						createTime: "时间时间"
					}
				}],
				
			}
		},
		methods: {
			
			
			/* 切换选项卡选项 */
			// tabActive(tabIndex) {
			// 	this.tablist.map((value, index) => {
			// 		value.isActive = tabIndex == index ? true : false;
			// 	})
			// 	if (tabIndex==0) {
			// 		this.recommend();
			// 	} else if (tabIndex==1) {
			// 		this.attention();
			// 	}else{
			// 		this.aaa();
			// 	}
			// 	this.tabIndex = tabIndex;
			// },
			
			
			recommend() {
				// console.log("互动");
			},
			attention() {
				// console.log("关注");
			},
			aaa(){
				// console.log("其他");
			}
		},
		
		// // 按钮切换显示
		// changeBtn() {
		// 	this.isDisabled = true
		// 	this.isLoading = true
		// 	let that = this;
		// 	setTimeout(function() {
		// 		that.isDisabled = false
		// 		that.isLoading = false
		// 		that.$emit('change');
		// 	}, 2000);
		// },
		// // 进入其他页面
		// inToPageMine() {
		// 	console.log(this.id)
		// }
	}
	
	
</script>

<style scoped>
	.header{
		height: 80rpx;
		line-height: 80rpx;
	}
	
	>>> .tab-card .head-nav .head-nav-bottom{
		width: 70%;
	}
	
	.flex-item {
		display: flex;
		flex-wrap: wrap;
		box-sizing: border-box;
		border-bottom: 6rpx solid #f0f0f0;
		padding: 20rpx 30rpx 20rpx 30rpx;
		/* margin-bottom: 20rpx; */
	}
	
	/* 左侧图片盒子样式 */
	.left-img-box {
		display: flex;
		align-items: center;
		/* flex: 1; */
		margin-right: 18rpx;
		/* width: 112rpx; */
		/* height: 112rpx; */
		vertical-align: middle;
	}
	
	.inner-img {
		height: 78rpx;
		width: 78rpx;
		border-radius: 50%;
	}
	
	/* 中部文字盒子样式 */
	.middle-text-box {
		align-items: center;
		flex: 5;
		display: flex;
	}
	
	/* 灰色字体 */
	.greay-text {
		color: #969696;
	}
	
	/* 昵称 */
	.inner-text-name {
		font-size: 35rpx;
		font-weight: bold;
		margin-right: 16rpx;
	}
	
	/* 关注数粉丝数文章数 */
	.inner-text-message {
		padding-top: 4rpx;
		font-size: 30rpx;
	}
	
	.inner-text{
		font-size: 35rpx;
	}
	
	>>> .article-img image{
		width: 120rpx;
		height: 120rpx;
	}
	
	>>> .article-text{
		font-size: 32rpx;
	    font-weight: bold;
	    -webkit-line-clamp: 2;
	    line-height: 180%;
	}
	
	>>> .article-content{
	    background-color: rgba(241, 241, 241, 1);
	    border-radius: 4px;
	    padding: 10px 12px;
	}
	.comment-content{
		margin-left: 96rpx;
		margin-top: 10rpx;
		}
	
	.text-comment{
		font-size: 30rpx;
		color: rgba(80, 80, 80, 1);
		width: 580rpx;
		white-space: nowrap;
		text-overflow:ellipsis;
		overflow:hidden;
	}
	
</style>
