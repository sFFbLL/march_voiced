<template>
	<view>
		<view class="header">
			<view class="header-left-box">
				<view class="header-img-box">
					<u-image src="../../static/img/sanyue.png" height="96rpx" width="96rpx"></u-image>
				</view>
				<view class="header-text">
					<span>三月之声创作平台</span>
				</view>
			</view>
			<view class="header-right-box">
				<view class="header-img-box">
					<u-image class="inner-img" src="../../static/img/my.png" height="50rpx" width="50rpx" shape="circle"></u-image>
				</view>
			</view>
		</view>
		<view class="my-tool-box">
			<view class="my-tool">
				<jinIcon class="single" type="&#xe6e7;" font-size="44rpx" title="加粗" @click="setBold" :color="showBold ? activeColor : '#666666'"></jinIcon>
				<jinIcon class="single" type="&#xe6fe;" font-size="44rpx" title="斜体" @click="setItalic" :color="showItalic ? activeColor : '#666666'"></jinIcon>
				<jinIcon class="single" type="&#xe6f8;" font-size="44rpx" title="分割线" @click="setIns" :color="showIns ? activeColor : '#666666'"></jinIcon>
				<jinIcon class="single" type="&#xe6e3;" font-size="44rpx" title="标题" @click="setHeader" :color="showHeader ? activeColor : '#666666'"></jinIcon>
				<jinIcon class="single" type="&#xe6f1;" font-size="44rpx" title="居中" @click="setCenter" :color="showCenter ? activeColor : '#666666'"></jinIcon>
				<jinIcon class="single" type="&#xe6ed;" font-size="44rpx" title="居右" @click="setRight" :color="showRight ? activeColor : '#666666'"></jinIcon>
				<jinIcon class="single" type="&#xe6f3;" font-size="44rpx" title="插入图片" @click="insertImage"></jinIcon>
				<jinIcon class="single" type="&#xe6eb;" font-size="44rpx" title="分割线" @click="insertDivider"></jinIcon>
				<jinIcon class="single" type="&#xe6e8;" font-size="44rpx" title="撤销" @click="undo"></jinIcon>
				<jinIcon class="single" type="&#xe705;" font-size="44rpx" title="重做" @click="redo"></jinIcon>
			</view>
		</view>
		<view class="body-box">
			<view class="body">

				<view class="body-left-box">

				</view>
				<view class="body-right-box">
					<view class="title-box">
						<input class="title-text" type="text" v-model="title" placeholder="标题" placeholder-class="placeholder" />
					</view>
					<!-- http://www.kuntong.site/api/file/uploadImage -->
					<jinEdit ref="child" placeholder="请输入内容" :html="html" :readOnly="false" @editOk="editOk" uploadFileUrl="http://www.kuntong.site/api/file/uploadImage"></jinEdit>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import jinEdit from '../../components/jin-edit/jin-edit.vue';
	import jinIcon from '../../components/jin-edit/jin-icons.vue';
	import {
		publishArticle
	} from '@/utils/api/publish-api.js'
	export default {
		data() {
			return {
				option: null,
				title: null,
				html: '<p>爱搜IDno为阿斯兰的框架内</p><hr><p><br></p><hr><h2>我是test</h2><p>我是<strong>加粗</strong></p><p><em>我是斜体</em><span class="ql-cursor">﻿</span></p>', // 编辑时候跳过来绑定的html内容
				showBold: false,
				showItalic: false,
				showIns: false,
				showHeader: false,
				showCenter: false,
				showRight: false,
				showSettingLayer: false,
				activeColor: '#F56C6C'
			}
		},
		components: {
			jinEdit,
			jinIcon
		},
		onLoad: function(option) { //option为object类型，会序列化上个页面传递的参数
			console.log(option)
			this.option = option //打印出上个页面传递的参数。
		},
		methods: {
			// 点击发布
			editOk(res) {
				let url = null;
				for (let i = 0; i < res.delta.ops.length; i++) {
					if (res.delta.ops[i].insert.image) {
						url = res.delta.ops[i].insert.image
						break;
					}
				}
				let params = {
					title: this.title,
					contnet: res.html,
					image: url,
					kind: 1,
					tag: this.option.tag,
					type: 0,
					status: res.isPublic,
					wordCount: res.textLength
				}
				// console.log(params)
				publishArticle(params).then(res => {
					if (res.code === 200) {
						console.log("发布成功")
					}
				})

				// console.log(this.title)
				// console.log(res.html)
				// console.log(url)
				// console.log(this.option.tag)
				// console.log("kind:1")
				// console.log("type:0")
				// console.log(res.isPublic)
			},
			setBold() {
				this.$refs.child.setBold();
			},
			setItalic() {
				this.$refs.child.setItalic();
			},
			setIns() {
				this.$refs.child.setIns();
			},
			setHeader() {
				this.$refs.child.setHeader();
			},
			setCenter() {
				this.$refs.child.setCenter();
			},
			setRight() {
				this.$refs.child.setRight();
			},
			insertImage() {
				this.$refs.child.insertImage();
			},
			insertDivider() {
				this.$refs.child.insertDivider();
			},
			undo() {
				this.$refs.child.undo();
			},
			redo() {
				this.$refs.child.redo();
			}
		}
	}
</script>

<style scoped>
	/deep/ .tool-view {
		display: none;
	}

	/deep/.single {
		height: 80rpx;
		font-size: 32rpx;
		padding: 0 30rpx;
		display: flex;
		align-items: center;
		line-height: 80rpx;
		flex-direction: row;
		color: #666;
	}

	/deep/.single .icon {
		font-size: 32rpx !important;
	}

	/deep/.ql-container {
		max-height: 1640rpx;
		overflow-y: scroll;
		/* padding-bottom: 200rpx !important; */
	}

	/deep/.placeholder,
	.title-text {
		/* text-align: center; */
		/* padding: 20rpx 30rpx; */
		font-size: 45rpx;
	}

	.title-box {
		padding: 30rpx 30rpx 30rpx 30rpx;
	}

	.title-text {
		color: rgb(118, 118, 118);

	}

	/deep/.container {
		box-sizing: border-box;
		border-top: 2rpx #d2d5cc solid;
	}

	.header {
		width: 100%;
		height: 116rpx;
		color: rgba(80, 80, 80, 1);
		background-color: rgba(255, 255, 255, 1);
		font-size: 27rpx;
		line-height: 150%;
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 12rpx 24rpx 12rpx 10rpx;
	}

	.header-left-box,
	.header-text,
	.header-img-box {
		display: flex;
		align-items: center;
	}

	.my-tool-box {
		width: 100%;
		/* height: 90rpx; */
		/* color: rgba(80, 80, 80, 1); */
		background: #eee;
		padding: 0 240rpx 0 1120rpx;
	}

	.my-tool {
		height: 100rpx;
		display: flex;
		align-items: center;
		justify-content: space-around;
		width: 100%;
		background: #eee;
	}

	.body-box {
		padding: 48rpx 300rpx 0 300rpx;
		display: flex;
		justify-content: center;
		padding-bottom: 20rpx;
	}

	.body {
		width: 2500rpx;
		display: flex;
		justify-content: space-between;
	}

	.body-left-box,
	.body-right-box {
		color: rgba(80, 80, 80, 1);
		background-color: rgba(255, 255, 255, 1);
		border-radius: 10rpx;
		font-size: 28rpx;
		line-height: 150%;
		box-shadow: 0px 8rpx 8rpx 0rpx rgba(229, 228, 228, 1);
	}

	.body-left-box {
		width: 614rpx;
		height: 1662rpx;

	}

	.body-right-box {
		width: 1570rpx;
		height: 1794rpx;
		padding: 50rpx;
	}
</style>
