<template>
	<view class="wrap">
		<text class="title">欢迎登录三月之声</text>
		<view class="inputs">
			<input class="input" placeholder="请输入简书链接(必填项)"   @input="janInput"/>
			<input class="input" placeholder="请输入CSDN链接"  @input="csdnInput"/></view>
				<button :loading="btnloading" class="submitbtn" @click="toSubmit">提交</button>
	</view>
</template>

<script>
	export default{
		data() {
			return {
				btnloading: false,
				janValue:"",
				csdnValue:"",
			}
		},
		methods:{
			// 获取链接
			janInput(event) {
				this.janValue = event.target.value;
			},
			csdnInput(event) {
				this.csdnValue = event.target.value;
			},
			toSubmit(){
				if(this.janValue==""){
					uni.showToast({
						title: '请完善简书链接',
						icon: "none",
						position: "top",
						duration: 2000
					});
				}else{
					// 调用接口
				}
			}
		}
	}
</script>

<style scoped>
	.wrap{
		margin: 100rpx auto;
	}
	.title {
		font-size: 50rpx;
		margin: 0 40rpx;
	}
	.inputs{
		margin: 70rpx 40rpx;
	}
	.input{
		margin: 60rpx 0rpx;
		width: 600rpx;
		border-bottom: 2rpx solid #f1eeec;
	}
	.submitbtn {
		width: 300rpx;
		margin-top: 250rpx;
		border-radius: 40rpx;
		color: #FFFFFF;
		background-image: linear-gradient(140deg, #7ce1e6 0%, #5b97e5 100%);
	}
	
</style>
