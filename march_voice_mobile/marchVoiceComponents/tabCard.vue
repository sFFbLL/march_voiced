<template>
	<view class="tab-card">
		<!-- 选项卡导航 -->
		<view class="head-nav">
			<view class="head-nav-item"
			 v-for="(item,index) in tabs"
			 :key="item.index"
			 :class="item.isActive ? 'active' : ''"
			 @click="handlerActive(item.index)">
				<view class="center">{{item.value}}
					<u-badge v-if="(item.index==0&&iread!=0)"
					 :is-dot="true"
					 type="error"></u-badge>
					<u-badge v-if="(item.index==1&&aread!=1)"
					 :is-dot="true"
					 type="error"></u-badge>
					<u-badge v-if="item.index==2&&oread!=2"
					 :is-dot="true"
					 type="error"></u-badge>
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
				aread: -1,
				oread: -1,
				iread: -1

			};
		},
		props: {
			tabs: {
				type: Array,
				default: []
			},
			otherRead: {
				type: Boolean,
				default: false
			},
			attentionRead: {
				type: Boolean,
				default: false
			},
			interactRead: {
				type: Boolean,
				default: false
			},

		},
		methods: {
			handlerActive(index) {
				if (index == 1 && this.attentionRead == false) {
					this.aread = 1;
					this.$emit('attentionRead', true)
				} else if (index == 2) {
					this.oread = 2;
					this.$emit('otherRead', true)
				} else if (index == 0) {
					this.iread = 0;
					this.$emit('interactRead', true)
				}
				this.$emit('tabActive', index);
			}
		},
		mounted() {},
		created() {

		}
	}
</script>

<style lang="scss"
 scoped>
	.tab-card {
		height: 100%;
	}

	.head-nav {
		display: -webkit-flex;
		display: flex;
		overflow: hidden;
		height: 100%;

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
		height: 100%;
	}


	.active .center {
		color: #404040;

		.head-nav-bottom {
			position: relative;
			margin: auto;
			width: 50%;
			height: 6rpx;
			top: -6rpx;
			background-color: #404040;
		}
	}
</style>
