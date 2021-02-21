<template>
 	<view class="content" :class="{'active':active}">
 		<view class="tabbar-box-wrap">
 			<view class="tabbar-box">
 				<view class="tabbar-box-item">
 					<publishTypeCard :tag="item.tag" :icon="item.icon"  v-for="(item,index) in list" :key="item.index">
 						<template v-slot:title>
 							写{{item.tag}}
 						</template>
 						<template v-slot:describe>
 							{{item.description}}
 						</template>
 					</publishTypeCard>
<!-- 					<publishTypeCard :kind="kind">
						<template v-slot:title>
							写笔记
						</template>
						<template v-slot:describe>
							记录我的生活，酸甜苦辣，喜怒哀乐
						</template>
					</publishTypeCard> -->
 				</view>
 			</view>
 		</view>
 	</view>
 </template> 
 
 <script>
	  import publishTypeCard from '../../marchVoiceComponents/publishTypeCard.vue';
	  import {
	  	getTags
	  } from '@/utils/api/publish-api.js';
	  import {check} from '../../utils/checkUnRead.js';
 export default {
 	data() {
 		return {
 			active: false,
			tag:"article",
			list:[
				{
					"tag": "文章",
					"icon": '../../static/img/1.jpg',
					"description":"1"
				},
				{
					"tag": "想法",
					"icon": '../../static/img/1.jpg',
					"description":"2"
				}
			]
 		};
 	},
	components:{
		publishTypeCard
	},
 	onLoad() {
		// 获取文章标签
		getTags().then(res => {
			this.list = res.data
			console.log("ASDASD")
		})
	},
 	onShow() {
 		// setTimeout(() => {
 		this.active = true;
 		// }, 500);
			check()
		
 	},
 	onHide() {
 		this.active = false;
 	},
 	methods: {
 		goToPage(url) {
 			if (!url) return;
 			uni.navigateTo({
 				url
 			});
 		}
 	}
 };
 </script>
 
 <style lang="scss" scoped>
 .content {
 	display: flex;
 	align-items: center;
 	justify-content: center;
 	width: 100%;
 	/* #ifdef H5 */
 	height: calc(100vh - var(--window-bottom) - var(--window-top));
 	/* #endif */
 	/* #ifndef H5 */
 	height: 100vh;
 	/* #endif */
 	transition: opacity 0.3s;

 	opacity: 0;
 	&.active {
 		opacity: 1;
 	}
 	.logo {
 		position: relative;
 		margin-top: -400upx;
 		width: 200upx;
 		height: 200upx;
 		// z-index: -1;
 		opacity: 0;
 		transition: opacity 0.3s;
 		&.active {
 			opacity: 1;
 		}
 	}
 }
 .tabbar-box-wrap {
 	position: absolute;
 	width: 100%;
 	// padding: 50upx;
 	box-sizing: border-box;
 	bottom: 0;
 	left: 0;
 	.tabbar-box {
 		position: relative;
 		display: flex;
 		width: 100%;
 		background: #fff;
 		border-radius: 20upx 20upx 0 0;
 		padding: 15upx 20upx;
 		box-sizing: border-box;
 		z-index: 2;
 		box-shadow: 0px 2px 5px 2px rgba(0, 0, 0, 0.1);
 		&:after {
 			content: '';
 			position: absolute;
 			bottom: -16upx;
 			left: 0;
 			right: 0;
 			margin: auto;
 			width: 50upx;
 			height: 50upx;
 			transform: rotate(45deg);
 			background: #fff;
 			z-index: 1;
 			box-shadow: 2px 2px 5px 1px rgba(0, 0, 0, 0.1);
 			border-radius: 2px;
 		}
 		&:before {
 			content: '';
 			position: absolute;
 			top: 0;
 			left: 0;
 			width: 100%;
 			height: 100%;
 			background: #ffffff;
 			border-radius: 20upx;
 			z-index: 2;
 		}
 		.tabbar-box-item {
 			// position: relative;
 			width: 100%;
 			z-index: 3;
 			margin: 10upx;
 			color: $uni-color-subtitle;
 			text-align: center;
 			font-size: $uni-font-size-base;
 			.box-image {
 				width: 100%;
 				height: $uni-img-size-lg;
 			}
 		}
 	}
 }
 </style>
 
