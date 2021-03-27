<template>
	<view>
		<view class="title-box">
			<input class="title-text" type="text" v-model="title" placeholder="标题" placeholder-class="placeholder"/>
		</view>
		<!-- http://www.kuntong.site/api/file/uploadImage -->
		<jinEdit placeholder="请输入内容" :html="html" :readOnly="false" @editOk="editOk" uploadFileUrl="http://linbolun.cn/api/file/uploadImage"></jinEdit>
	</view>
</template>

<script>
	import jinEdit from '../../components/jin-edit/jin-edit.vue';
	import { publishArticle,upDateArticle } from '@/utils/api/publish-api.js';
	import { getArtileDetails } from '../../utils/api/articleDetails-api.js';
	export default {
		data() {
			return {
				option:null,
				title:null,
				html:'' ,// 编辑时候跳过来绑定的html内容
				draftInfo:{},
				isUpdateArticle:false,
			}
		},
		components: {
			jinEdit
		},
		onLoad: function(option) { //option为object类型，会序列化上个页面传递的参数
			this.option = option //打印出上个页面传递的参数。
		},
		created() {
			if(this.option.type === "draft"){
				let id = this.option.id;
				// 获取文章详情内容
				getArtileDetails(id).then(res => {
						this.draftInfo = res.data
						this.html = this.draftInfo.content
						this.title = this.draftInfo.title
						// this.option.tag = this.draftInfo.tag
						this.option.tag = 1
						this.isUpdateArticle = true
				}).catch(err => {
					console.log(err, "err login")
				})
			}
		},
		methods: {
			// 点击发布
			editOk(res) {
				let url = null;
				let describe = res.text.substr(0,20)
				for(let i = 0; i < res.delta.ops.length; i++){
					if(res.delta.ops[i].insert.image){
						url = res.delta.ops[i].insert.image
						break;
					}
				}
				let params = {
					title:this.title,
					content:res.html,
					image:url,
					kind:1,
					tag:parseInt(this.option.tag),
					type:0,
					status:res.isPublic,
					word_count:res.textLength,
					describe:describe
				}
				switch (this.isUpdateArticle){
					case false:
						publishArticle(params).then(res => {
							uni.switchTab({
								url:'/pages/publish/index'
							})
						})
						break;
					case true:
						params.id = parseInt(this.option.id)
						upDateArticle(params).then(res =>{
							uni.switchTab({
								url:'/pages/publish/index'
							})
						})
					default:
						break;
				}
				
				
				// console.log(this.title)
				// console.log(res.html)
				// console.log(url)
				// console.log(this.option.tag)
				// console.log("kind:1")
				// console.log("type:0")
				// console.log(res.isPublic)
			}
		}
	}
</script>

<style scoped>
view /deep/.ql-container{
	max-height: 1050rpx;
	overflow-y: scroll;
	/* padding-bottom: 200rpx !important; */
}
/deep/.placeholder,.title-text{
	/* text-align: center; */
	/* padding: 20rpx 30rpx; */
	font-size: 45rpx;
}
.title-box{
	padding: 30rpx 30rpx 30rpx 30rpx;
}
.title-text{
	color: rgb(118, 118, 118);

}
/deep/.container{
	box-sizing: border-box;
	border-top: 1px #d2d5cc solid;
}
</style>
