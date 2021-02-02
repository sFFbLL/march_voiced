<template>
	<view class="home">
		<view class="header">
			<view class="head-nav-items">
				<tabs :tabs='tablist' v-on:tabActive='tabActive' />
			</view>
			<view class="search">
				<view class="search-icon">
					<uni-icons type="search" size="30" color="#999"></uni-icons>
				</view>
			</view>
		</view>
		<view class="content">
			<view v-for="(index,item) in articleList">
				<!-- <recommend :article /> -->
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
				contentIndex: '',
				articleList: [{
					article: {
						id: 1,
						title: "文章1",
						content: "文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容",
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 1,
						userName: "张三",
						userImage: "",
						isFollow: "是否关注"
					}
				}, {
					article: {
						id: 1,
						title: "文章1",
						content: "文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容文章内容",
						upDateTime: "更新时间",
						favourTotal: 1,
						collectTotal: 1,
						commentTotal: 1
					},
					user: {
						userId: 1,
						userName: "张三",
						userImage: "",
						isFollow: "是否关注"
					}
				}]

			}
		},
		methods: {
			tabActive(tabIndex) {
				this.tablist.map((value, index) => {
					value.isActive = tabIndex == index ? true : false;
				})
				if (!tabIndex) {
					this.recommend();
				} else if (tabIndex) {
					this.attention();
				}
				this.contentIndex = tabIndex;
			},
			recommend() {
				console.log(1);
			},
			attention() {
				console.log(2);
			}
		}
	}
</script>

<style>
	.home {
		background-color: rgba(241, 241, 241, 1);
	}

	.header {
		height: 100rpx;
		line-height: 100rpx;
		background-color: #fff;
		margin-bottom: 15rpx;
	}

	.head-nav-items {
		display: inline-block;
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
	
	.content{
		min-height: 300rpx;
		background-color: #fff;
	}
</style>
