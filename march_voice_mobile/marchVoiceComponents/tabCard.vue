<template>
	<view class="tab-card">
		<!-- 选项卡导航 -->
		<view class="head-nav">
			<view class="head-nav-item" v-for="(item,index) in tabs" :key="item.index" :class="item.isActive ? 'active' : ''"
			 @click="handlerActive(item.index)">
				<view class="center">{{item.value}}
					<u-badge v-if="(item.index==0&&interactRead==false)" :is-dot="true" type="error"></u-badge>
					<u-badge v-if="(item.index==1&&attentionRead==false)" :is-dot="true" type="error"></u-badge>
					<u-badge v-if="item.index==2&&otherRead==false" :is-dot="true"  type="error"></u-badge>
					<view class="head-nav-bottom"></view>
				</view>
			</view>

		</view>
		<slot></slot>
	</view>
</template>

<script>
	import readMessage from "../utils/api/message-api.js"
	export default {
		data() {
			return {
			
			};
		},
		props: {
			tabs: {
				type: Array,
				default: []
			},

		},
		methods: {
			handlerActive(index) {
				if(index==0){
					// 消除红点
					this.$store.commit('changeInteract',1);
					// 已读接口
					// readMessage(1).then(res=>{
						
					// })
				}else if(index==1){
					// 消除红点
					this.$store.commit('changeAttention',1);
					// 已读接口
					// readMessage(2).then(res=>{
						
					// })
				}else if(index==2){
					// 消除红点
					this.$store.commit('changeOther',1);
					// 已读接口
					// readMessage(3).then(res=>{
						
					// })
				}
				
				this.$emit('tabActive', index);
			}
		},
		computed:{
			interactRead(){
				return this.$store.state.interactRead;
			},
			attentionRead(){
				return this.$store.state.attentionRead;
			},
			otherRead(){
				return this.$store.state.otherRead;
			}
		},
		mounted() {},
		created() {

		}
	}
</script>

<style lang="scss">
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
		width: 228rpx;
	}


	.active .center {
		color: #404040;

		.head-nav-bottom {
			position: relative;
			margin: auto;
			width: 50%;
			height: 8rpx;
			top: -8rpx;
			background-color: #404040;
		}
	}
</style>
