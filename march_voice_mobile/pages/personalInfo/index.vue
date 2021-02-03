<template>
	<view class="content"><!---->
		<view class="base-info-tabbar">
			<text>基本信息</text>
		</view>
		
		<view class="base-info-content">
			<view class="base-info-head-img" style="height: 130rpx;">
				<text clas="base-info-area">头像</text>
				<view class="base-info-change-img">
			            <image style="width: 100rpx; height: 100rpx;" mode="aspectFill" src="../../static/img/my1.png"
						@click="changeHeadImg()"></image>
					<uni-icons type="arrowright" size="15"></uni-icons>
				</view>
					
			</view>
			
			<view class="base-info-name">
				<text clas="base-info-area">昵称</text>
				<view class="base-info-change-name">
					<view class="limit-lengh">
						
						<text class="title" @click="changeNameOpen()">{{baseUserName}}</text>
					</view>
					
					<uni-popup ref="popup" type="dialog">
					    <uni-popup-dialog mode="input" title="请修改昵称" placeholder="请输入2-15个字符" :duration="2000"
						:before-close="true" @close="close" @confirm="changeNameConfirm"></uni-popup-dialog>
					</uni-popup>
					
					<uni-popup ref="popupUp" type="dialog">
					    <uni-popup-dialog mode="base" type="error" title="警告" content="请输入2-15个字符" :duration="2000"
						:before-close="true" @close="close" @confirm="close"></uni-popup-dialog>
					</uni-popup>
					<uni-icons type="arrowright" size="15"></uni-icons>
				</view>			
			</view>
		
			<view class="base-info-gender">
				<text clas="base-info-area">性别</text>
				<view class="base-info-gender-change">
				    <picker @change="bindPickerChange" :value="index" :range="array">
						<text class="uni-input">{{array[index]}}</text>
				    </picker>
					<uni-icons type="arrowright" size="15"></uni-icons>
				</view>
			</view>
		
			<view class="base-info-person-info">
				<text clas="base-info-area">个人简介</text>
				<view class="base-info-change-info">
					<view class="limit-lengh">
						<!-- -->
						<text class="title"  @click="changeInfoOpen()">{{baseUserInfo}}</text>
					</view>
					
							<uni-popup ref="popup1" type="dialog">
							    <uni-popup-dialog mode="input" title="请修改个人简介" placeholder="用一句话介绍一下自己吧(50字以内)" :duration="2000"
								:before-close="true" @close="close" @confirm="changeInfoConfirm"></uni-popup-dialog>
							</uni-popup>
							
							<uni-popup ref="popup1Up" type="dialog">
							    <uni-popup-dialog mode="base" type="error" title="警告" content="请输入2-15个字符" :duration="2000"
								:before-close="true" @close="close" @confirm="close"></uni-popup-dialog>
							</uni-popup>
						<uni-icons type="arrowright" size="15" @click="changeInfoOpen()"></uni-icons>
				</view>	
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				array: ['女', '男'],
				index: 0,
				baseUserName:'三月强化班',
				baseUserInfo:'这个人很懒，什么也没有写'
			}
		},
		onLoad() {},
		methods: {
			bindPickerChange: function(e) {
	            // console.log('picker发送选择改变，携带值为', e.target.value)
	            this.index = e.target.value
	        },
			changeHeadImg: function(e) {
			    uni.chooseImage({
			        count: 1, //默认9
			        sizeType: ['compressed'], //可以指定是原图还是压缩图，默认二者都有
			        sourceType: ['album','camera'], //从相册选择
			        success: function (res) {
			            console.log(JSON.stringify(res.tempFilePaths));
			        }
			    });
			},
			
			
			changeNameOpen(){
			    this.$refs.popup.open()
			},
			/**
	         * 点击取消按钮触发
	         * @param {Object} done
	         */
	        close(done){
	            // TODO 做一些其他的事情，before-close 为true的情况下，手动执行 done 才会关闭对话框
	            // ...
	            done()
	        },
			
	        /**
	         * 点击确认按钮触发
	         * @param {Object} done
	         * @param {Object} value
	         */
	        changeNameConfirm(done,value){
	            // 输入框的值2-15个字符
				if(/^.{2,15}$/.test(value)){
					//true
					this.baseUserName = value;
				}else{
					this.$refs.popupUp.open();
					return;
				}
	            console.log(value)
	            // TODO 做一些其他的事情，手动执行 done 才会关闭对话框
	            // ...
	            done()
	        },
			
			changeInfoOpen(){
			    this.$refs.popup1.open()
			},
			
			changeInfoConfirm(done,value){
			    // 输入框的值0-50个字符
				if(/^.{0,50}$/.test(value)){
					//true
					this.baseUserInfo = value;
				}else{
					this.$refs.popup1Up.open();
					return;
				}
			    console.log(value)
			    // TODO 做一些其他的事情，手动执行 done 才会关闭对话框
			    // ...
			    done()
			},
		}
	}
</script>

<style>
	.limit-lengh{
		width: 300rpx;
        text-align:right;
		white-space: nowrap;
		text-overflow:ellipsis;
		overflow:hidden;
	}
	.base-info-tabbar{
		font-size: 30rpx;
		font-weight: bold;
		display: flex;
		flex-direction: column;
		border-bottom: 1px solid #F1F1F1;
		padding: 0rpx 28rpx;
	}
	
	.base-info-content{
		color: #404040;
		padding: 0rpx 28rpx;
		font-size: 32rpx;
		display: flex;
		flex-direction: column;
	}
	.content>.base-info-tabbar{
		height: 86rpx;
		display: flex;
		flex-direction: column;
		justify-content: center;
	} 
	.base-info-gender-change,.base-info-change-img,.base-info-change-name,.base-info-change-info{
		display: flex;
		align-items: center;
	}
	.base-info-content>view{
		height: 86rpx;
		line-height: 86rpx;
		display: flex; 
		flex-direction: row;
		justify-content:space-between;
		align-items: center;
	} 
	.base-info-content>*{
		border-bottom: 1px dashed #F1F1F1;
	}
</style>
