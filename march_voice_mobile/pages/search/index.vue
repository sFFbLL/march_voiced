<template>
	<view class="search">
		<!-- 搜索页面头部 -->
		<view class="search-header">
			<view class="search-input">
				<u-search :show-action="false"
				 placeholder="搜索感兴趣的内容"
				 v-model="searchText"
				 @search="search()"
				 @clear="clear()"></u-search>
			</view>
		</view>
		<!-- 用户列表-->
		<view class="searchUser">
			<view class="userList"
			 v-for="(item,index) in userList">
				<image :src=item.image
				 mode="aspectFill"></image>
			</view>
		</view>
	</view>
</template>

<script>
	import {
		search
	} from '@/utils/api/home-api.js'
	export default {
		data() {
			return {
				searchText: "",
				articleList: [],
				userList: [],
				articleList1: [{}],
				userList1: [{
						userId: 1,
						name: "张三",
						image: require('static/img/2.jpg')
					},
					{
						userId: 2,
						name: "张四",
						image: require('static/img/1.jpg')
					},
					{
						userId: 3,
						name: "张五",
						image: require('static/img/2.jpg')
					},
					{
						userId: 4,
						name: "张六",
						image: require('static/img/2.jpg')
					},
					{
						userId: 5,
						name: "张七",
						image: require('static/img/1.jpg')
					},
					{
						userId: 6,
						name: "张八",
						image: require('static/img/2.jpg')
					},
					{
						userId: 7,
						name: "张九",
						image: require('static/img/1.jpg')
					},
					{
						usersId: 8,
						name: "张十",
						image: require('static/img/2.jpg')
					}
				]
			}
		},
		methods: {
			search() {
				// let _this = this;
				if (this.searchText) {
					let userList = this.userList1;
					/* let _this = this;
					let params = {
						searchText: this.searchText,
						current: this.current,
						size: this.size
					}
					search().then(res => {
						_this.userList = userList;
					}) */
					this.userList = userList;
					console.log(this.userList);
				}

			},
			clear() {
				this.searchText = "";
			}

		}
	}
</script>

<style>
	.search-header {
		height: 110rpx;
		line-height: 100rpx;
	}

	.search-input {
		width: 90%;
		margin: auto;
	}

	.searchUser {
		width: 100%;
		height: 100rpx;
	}

	.userList {
		display: inline-block;
	}
</style>
