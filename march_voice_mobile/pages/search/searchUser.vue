<template>
	<view class="search-user-list">
		<view class="line"></view>
		<view class="content">
			<view class="search-user-item"
			 v-for="(item,index) in userList">
				<attentionAndFansCell :id="item.id"
				 :nickname="item.nickname"
				 :avatarPath="item.avatarPath"
				 :isFollow="item.isFollow">
					<view slot="underText"
					 class="user-signature">{{item.signature}}</view>
				</attentionAndFansCell>
			</view>
			<!-- 下拉加载更多 -->
			<view v-show="isLoadMore">
				<uni-load-more class="loading"
				 :status="loadStatus"
				 iconType="circle"></uni-load-more>
			</view>
		</view>

	</view>
</template>

<script>
	import attentionAndFansCell from '../../marchVoiceComponents/attentionAndFansCell.vue'
	import {
		searchUser
	} from '@/utils/api/home-api.js'
	export default {
		data() {
			return {
				loadStatus: 'loading', //加载样式：more-加载前样式，loading-加载中样式，nomore-没有数据样式
				isLoadMore: false, //是否加载中
				userCurrent: 1, //当前页数
				size: 11,
				userList: [],
				userList1: [{
						id: 1,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 1
					},
					{
						id: 2,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 1
					}, {
						id: 3,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/2.jpg'),
						isFollow: 1
					}, {
						id: 4,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 0
					}, {
						id: 5,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/2.jpg'),
						isFollow: 0
					}, {
						id: 6,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/2.jpg'),
						isFollow: 1
					}, {
						id: 7,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/2.jpg'),
						isFollow: 1
					}, {
						id: 8,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 0
					}, {
						id: 9,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/2.jpg'),
						isFollow: 1
					}, {
						id: 10,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/2.jpg'),
						isFollow: 1
					}, {
						id: 11,
						nickname: "张三",
						signature: "我只是我，我没有个性签名，你凑活着看吧",
						avatarPath: require('../../static/img/1.jpg'),
						isFollow: 0
					}
				]
			};
		},
		props: {

		},
		components: {
			attentionAndFansCell
		},
		onLoad() {
			this.getSearchUserList();
		},

		onReachBottom() { //上拉触底函数
			if (!this.isLoadMore) {
				this.isLoadMore = true
				this.userCurrent += 1;
				this.getSearchUserList();
			}
		},
		methods: {
			getSearchUserList() {
				let _this = this;
				if (this.userList.length > 50) {
					_this.loadStatus = "nomore";
				} else {
					setTimeout(function () {
						_this.isLoadMore = false;
						// _this.getUserList();
						_this.userList = [..._this.userList, ..._this.userList1];
					}, 2000);
				}
			},
			getUserList() {
				let _this = this;
				let params = {
					searchText: this.searchText,
					size: this.userSize
				}
				searchUser().then(res => {
					_this.userList = [..._this.userList, ...res.data.userList];
				})
			}
		},
		mounted() {
			this.userList = [...this.userList, ...this.userList1];
		}
	}
</script>

<style scoped>
	.line {
		position: fixed;
		height: 1rpx;
		width: 100%;
		background-color: #f0f0f0;
	}

	.content {
		margin: 0 30rpx;
	}

	>>>.attention-cell .flex-item {
		padding: 30rpx;
		border-bottom: 1rpx solid #f0f0f0;
	}

	>>>.inner-img {
		height: 70rpx;
		width: 70rpx;
	}

	>>>.right-button {
		width: 130rpx;
		height: 54rpx;
		font-size: 24rpx;
		line-height: 54rpx;
	}


	.user-signature {
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
		width: 85%;
	}

	>>>.uni-load-more .uni-load-more__img {
		height: 30rpx !important;
		width: 30rpx !important;
	}
</style>
