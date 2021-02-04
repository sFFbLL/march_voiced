<template>
	<view class="home">
		<view class="header">
			<tabs class="tag-nav"
			 :tabs='tablist'
			 v-on:tabActive='tabActive' />

			<view class="search">
				<view class="search-icon">
					<uni-icons type="search"
					 size="30"
					 color="#999"></uni-icons>
				</view>
			</view>
		</view>
		<view class="content">
			<view v-if="!tabIndex">
				<view v-for="(item,index) in recommendList">
					<recommend :articleInfo="item" />
				</view>
			</view>
			<view v-if="tabIndex">
				<view>
					<view v-for="(item,index) in attentionList">
						<recommend :articleInfo="item" />
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import uniIcons from "@/components/uni-icons/uni-icons.vue"
	import tabs from '../../marchVoiceComponents/tabCard.vue'
	import recommend from '../../marchVoiceComponents/showArticle/recommend.vue'

	export default {
		components: {
			uniIcons,
			tabs,
			recommend
		},
		data() {
			return {
				tabIndex: '',
				tablist: [{
						index: 0,
						value: '推荐',
						isActive: true
					},
					{
						index: 1,
						value: '关注',
						isActive: false
					}
				],
				recommendList: [{
					article: {
						id: 1,
						title: "我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程？",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹...",
						articleImg: require('static/img/2.jpg'),
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 1,
						userName: "张三",
						userImage: "",
						isFollow: 0
					}
				}, {
					article: {
						id: 2,
						title: "所以监听用户的截图操作，提示用户进行分",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						articleImg: "",
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 2,
						userName: "李四",
						userImage: "",
						isFollow: 1
					}
				}],
				attentionList: [{
					article: {
						id: 2,
						title: "所以监听用户的截图操作，提示用户进行分",
						content: "今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈...",
						articleImg: "",
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 2,
						userName: "李四",
						userImage: "",
						isFollow: 1
					}
				}]
			}
		},
		methods: {
			/* 切换选项卡选项 */
			tabActive(tabIndex) {
				this.tablist.map((value, index) => {
					value.isActive = tabIndex == index ? true : false;
				})
				if (!tabIndex) {
					this.recommend();
				} else if (tabIndex) {
					this.attention();
				}
				this.tabIndex = tabIndex;
				console.log(this.tabIndex);
			},
			recommend() {
				// console.log(1);
			},
			attention() {
				// console.log(2);
			}
		}
	}
</script>

<style>
	.home,
	.home .content {
		background-color: rgba(241, 241, 241, 1);
	}

	.header {
		height: 100rpx;
		line-height: 100rpx;
		background-color: #fff;
		margin-bottom: 15rpx;
	}

	.tag-nav {
		display: inline-block;
		height: 100%;
		width: 90%;
	}

	.search {
		display: inline-block;
		width: 8%;
	}

	.search .search-icon {
		position: absolute;
		top: 5rpx;
	}

	.content {
		background-color: #fff;
	}
</style>
