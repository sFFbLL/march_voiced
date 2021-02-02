<template>
	<view>
		<view class="head-nav">
			<view class="head-nav-item" v-for="(item,index) in tabs" :key="item.index" :class="item.isActive ? 'active' : ''"
			 @tap="handlerActive(item.index)">
				<view class="center">{{item.value}}</view>
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

		.head-nav-item {
			flex: 1;
			text-align: center;
			font-size: 36rpx;
			font-family: PingFangSC;
			color: #999;
		}
	}

	.head-nav-item .center {
		margin: auto;
		width: 40%;
		height: 90%;
	}


	.active .center {
		color: #404040;
		border-bottom: 8rpx solid #404040;
	}
</style>
