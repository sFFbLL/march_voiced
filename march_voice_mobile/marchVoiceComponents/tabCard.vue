<template>
	<view class="tab-card">
		<!-- 选项卡导航 -->
		<view class="head-nav">
			<view class="head-nav-item"
			 v-for="(item,index) in tabs"
			 :key="item.index"
			 :class="item.isActive ? 'active' : ''"
			 @tap="handlerActive(item.index)">
				<view class="center">{{item.value}}
					<view class="head-nav-bottom"></view>
				</view>
			</view>
		</view>
		<slot></slot>
	</view>
</template>

<script>
	export default {
		data() {
			return {

			};
		},
		props: {
			tabs: {
				type: Array,
				default: []
			}
		},
		methods: {
			handlerActive(index) {
				this.$emit('tabActive', index)
			}
		},
		mounted() {}
	}
</script>

<style lang="scss">
	.head-nav {
		display: -webkit-flex;
		display: flex;
		overflow: hidden;

		.head-nav-item {
			flex: 1;
			text-align: center;
			font-size: 34rpx;
			font-family: Microsoft YaHei;
			font-weight: bold;
			letter-spacing: 4rpx;
			color: #999;
		}
	}

	.head-nav-item .center {
		position: relative;
		margin: auto;
	}


	.active .center {
		color: #404040;

		.head-nav-bottom {
			position: relative;
			margin: auto;
			width: 50%;
			background-color: #404040;
		}
	}

	.head-nav-bottom {
		position: absolute;
		bottom: 0;
		height: 8rpx;
		width: 100%;
	}
</style>
