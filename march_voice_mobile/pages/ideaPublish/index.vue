<template>
	<view class="wrap">

		<view>
			<textarea class="input" placeholder="这一刻的想法" maxlength="-1" @input="getWords" />

			</view>
		<view class="imgs" >
			<view v-for="(item,index) in srcList">
			<image :src="item" class="chooseimg" mode="scaleToFill" @click="preview(item)" @longtap="todelete(index)"></image>
			</view>
			<button v-if="srcList.length!=9"  @click="upImg()" style="width: 220rpx;height: 220rpx;"><image src="../../static/img/add.png" style="width: 100rpx;height: 100rpx; margin-top: 50rpx;"></image></button>
		</view>
		<view class="publish">
			<button  :loading="btnLoading" class="publishbtn" @click="publish()">发布想法</button>
		</view>
		
	</view>
</template>

<script>
	
	import unloadImage from "../../utils/api.js"
	export default {
		data() {
			return {
				ideaWords: "",
				srcList:[],
				btnLoading:false,
			}
		},
		
		methods: {
			// 删除图片
			todelete(index){
				let _this=this;
				uni.showActionSheet({
				    itemList: ['删除'],
				    success: function (res) {
				        console.log('选中了第' + index + '个图片');
						_this.srcList.splice(index,1)
				    },
				    fail: function (res) {
				    }
				});
				
			},
		// 实时获取输入文字
			getWords(event) {
				this.ideaWords = event.target.value;
			},
			// 上传图片显示到页面
			upImg(){
				let _this =this;
				uni.chooseImage({
					    count: 6, //默认9
					    sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
					    sourceType: ['album'], //从相册选择
						success(res) {
						for(let value of res.tempFilePaths){
								_this.srcList.push(value);
						}
						
					}
				})
				
			},
			// 预览图片
			preview(item){
				
				let _this = this ;
				uni.previewImage({
				        current: item,
						urls:_this.srcList,
				        longPressActions: {
				            success: function(data) {
				                console.log('选中了第'  + '个按钮,第' + '张图片');
				            },
				            fail: function(err) {
				            }
				        }
				    });
				},
				// 发布想法，调接口
			publish(){
				this.btnLoading=true;
				let _this = this ;
				if(this.ideaWords==""){
					uni.showToast({
					    title: '请输入你的想法',
						icon:"none",
						position:"top",
					    duration: 2000
					});
				}else{
						let imageList=""
						// 调用接口转化imgurl
						let file=this.srcList
						
						unloadImage(file).then(res=>{
						imageList=	res.data.full_path
						})
						let params={
							content:this.ideaWords,
							imageList:imageList,
							image:imageList[0]
							}
						// 调用发布接口,
					// marchCircleList(params).then(res=>{
					// 		 uni.showToast({
					// 			title: '发布成功',
					// 			icon:"none",
					// 			position:"top",
					// 			 duration: 2000
					// 			});
					// })
					
					}
				
				setTimeout(function() {
					_this.btnLoading=false;
				}, 2000);
				
			}
		}
	
	}
</script>

<style scoped>
	/* 发布部分 */
	.publishbtn{
		border-radius: 40rpx;
		width: 500rpx;
		height: 80rpx;
		color: #FFFFFF;
		background-image: linear-gradient(140deg, #7ce1e6 0%, #5b97e5 100%);
		display: inline-block;
		text-align: center;
		line-height: 2.3;
		position: fixed;
		bottom: 10rpx;
		left: 120rpx;
	}
	.publishimg{
		width: 80rpx;
		height: 80rpx;
	}
	/* 图片部分 */
	uni-button{
		margin: 0 0;
	}
	.imgs{
		display: flex;
		flex-flow: row wrap;
	
	}
	.chooseimg{
		width:220rpx;
		height: 220rpx;
		margin-right: 10rpx;
		
	}
	/* 文字部分 */
	.input {
		width: 100%;
		height: 300rpx;
		text-align: left;
		overflow: hidden;
		-webkit-line-clamp: 2;
		text-overflow: ellipsis;
		display: -webkit-box;
		-webkit-box-orient: vertical;
		margin-bottom: 20rpx;
	}

	.wrap {
		padding: 30rpx;
	}
</style>
