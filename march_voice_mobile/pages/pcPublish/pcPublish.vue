<template>
	<view class="page-warp">
		<view class="header top-warp">
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
		<view class="my-tool-box top-warp">
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
		<view class="body-box center-warp">
			<view class="body">
				<view class="body-left-box right-warp" ref="test">
					<view class="body-left-header">
						<mescroll-uni :fixed="false" ref="mescrollRef" @init="mescrollInit" @down="downCallback" @up="upCallback">
							<view class="add-article-btn" @click="newArticle">
								<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 48 48" fill="rgba(153, 153, 153, 1)">
									<path d="M0 0h48v48H0z" fill="none"></path>
									<path d="M24 4C12.95 4 4 12.95 4 24s8.95 20 20 20 20-8.95 20-20S35.05 4 24 4zm10 22h-8v8h-4v-8h-8v-4h8v-8h4v8h8v4z"></path>
								</svg>
								<span>新建文章</span>
							</view>

							<view class="article-card" v-for="item,index in list" :key="index" @click="showArticleInEditor(item.title,item.tag,item.content,item.word_count,item)">
								<u-image src="../../static/img/aiticleLogo.png" mode="widthFix" width="50rpx"></u-image>
								<view class="article-card-text-box">
									<span class="article-card-title">{{item.title}}</span>
									<span class="article-card-tag" style="color: rgba(128, 128, 128, 1);" v-if="item.status === 0">草稿</span>
									<span class="article-card-tag" style="color: #19be6b" v-if="item.status === 1">已发布</span>
									<span class="article-card-tag" style="color: rgba(42, 130, 228, 1);" v-if="item.status === 2">审核中</span>
								</view>
							</view>
						</mescroll-uni>
					</view>
				</view>
				<!-- <scroll-view class="left-warp" :scroll-y="true"> -->
				<view class="body-right-box">
					<view class="title-box">
						<input class="title-text" type="text" v-model="title" placeholder="标题" placeholder-class="placeholder" />
					</view>
					<view  class="tags-box">
						<view class="uni-list">
							<view class="uni-list-cell">
								<view class="uni-list-cell-db">
									<picker @change="bindPickerChange" :value="index" :range="tags">
										<view class="uni-input">{{tags[index]}}</view>
									</picker>
								</view>
							</view>
						</view>
					</view>
					<!-- http://www.kuntong.site/api/file/uploadImage -->
					<jinEdit ref="child" placeholder="请输入内容" :readOnly="false" @editOk="editOk" uploadFileUrl="http://linbolun.cn/api/file/uploadImage" @eidtorChange="eidtorChange"></jinEdit>
					<view class="body-right-foot-box">
						<view class="text-box">
							<span>正文字数 {{word_count}}</span>
						</view>
						<view class="btn-box">
							<u-button :custom-style="save" type="primary" size="medium " @click="release(0)">保存到草稿箱</u-button>
							<u-button :custom-style="publish" type="info" size="medium " @click="release(2)">发布</u-button>
						</view>
					</view>
				</view>
				<!-- </scroll-view> -->
			</view>
		</view>
	</view>
</template>

<script>
	import jinEdit from '../../components/jin-edit/jin-edit.vue';
	import jinIcon from '../../components/jin-edit/jin-icons.vue';
	import {
		publishArticle,
		getArticleList,
		upDateArticle,
		getTags
	} from '@/utils/api/publish-api.js'
	import MescrollMixin from "@/components/mescroll-uni/mescroll-mixins.js";

	export default {
		mixins: [MescrollMixin], // 使用mixin
		data() {
			return {
				title: null,
				html: `
				<div class="image-package  ">\n<div class="image-container" style="max-width: 700px; max-height: 807px;">\n<div class="image-container-fill" style="padding-bottom: 115.21%;"></div>\n<div class="image-view" data-width="743" data-height="856"><img data-original-src="//upload-images.jianshu.io/upload_images/14261054-7b08f2243c713c38.png" data-original-width="743" data-original-height="856" data-original-format="image/png" data-original-filesize="670480"></div>\n</div>\n<div class="image-caption"></div>\n</div><p>最近可谓是在瓜田上蹿下跳的一周，感觉这个社会戏子的事天下皆知，穷困的人却还难以维持正常的生活</p><p>一直以来感觉这些事情距离我们还是比较遥远的，但是实际上居然就每天在发生着。</p><p>一方面明星各种黑料爆出甚至违反社会底线，另一方面却有和我们相仿的年纪的人却从来没有听说过冬至要吃饺子。明明在努力生活却难以继续活下去...</p><div class="image-package  ">\n<div class="image-container" style="max-width: 700px; max-height: 143px;">\n<div class="image-container-fill" style="padding-bottom: 15.329999999999998%;"></div>\n<div class="image-view" data-width="933" data-height="143"><img data-original-src="//upload-images.jianshu.io/upload_images/14261054-a60d4c10a81dc1d5.png" data-original-width="933" data-original-height="143" data-original-format="image/png" data-original-filesize="21945"></div>\n</div>\n<div class="image-caption"></div>\n</div><div class="image-package  ">\n<div class="image-container" style="max-width: 700px; max-height: 456px;">\n<div class="image-container-fill" style="padding-bottom: 40.71%;"></div>\n<div class="image-view" data-width="1120" data-height="456"><img data-original-src="//upload-images.jianshu.io/upload_images/14261054-c5bd0780256289b6.png" data-original-width="1120" data-original-height="456" data-original-format="image/png" data-original-filesize="155833"></div>\n</div>\n<div class="image-caption"></div>\n</div><div class="image-package  ">\n<div class="image-container" style="max-width: 592px; max-height: 198px;">\n<div class="image-container-fill" style="padding-bottom: 33.45%;"></div>\n<div class="image-view" data-width="592" data-height="198"><img data-original-src="//upload-images.jianshu.io/upload_images/14261054-0d03a5e91df77209.png" data-original-width="592" data-original-height="198" data-original-format="image/png" data-original-filesize="18879"></div>\n</div>\n<div class="image-caption"></div>\n</div><p>看到新闻的时候就眼泪唰唰往下流\xa0 明明只是一个努力向上想活下去的人 最后却是被房东发现离世的..</p><p>然后突然就觉得，我们每天还有很多剩饭剩菜，以后吃饭一定尽最大努力做到光盘行动。多考虑一下还有很多人想吃这些,,,同时想要做一些什么，但是仔细想了想又觉得很无力，一个人的力量太渺小了，况且我还只是个学生</p><p>可能我是个共情能力很强的人，所以总是会被很多事情困扰，总是想要每个人都好，可是现实却往往做不到这样，小许说他没什么感觉，他会觉得授人以鱼不如授人以渔，而且这种事情很多，帮不过来。我觉得能帮助一点是一点，至少现在的我们还可以拿出自己的一部分去帮助别人。所以选择了春蕾计划，每周少吃两顿好吃的，就可以帮助一点~！</p><p>立一个flag，到工作的时候资助一个贫穷孩子的高中学习，貌似1800一年，加油吧！</p>
				`,
				word_count:null,
				
				tags: [],
				tagsId:[],
				index: 0, // tags默认选中的下标
				
				showBold: false,
				showItalic: false,
				showIns: false,
				showHeader: false,
				showCenter: false,
				showRight: false,
				showSettingLayer: false,
				activeColor: '#F56C6C',

				save: {
					borderRadius: "6rpx"
				},
				publish: {
					borderRadius: "6rpx",
					color: "rgba(0, 145, 255, 1)"
				},
				
				list:[],
				resList:[],
				isUpdate:false, // 是否是编辑文章，用来调不同的接口
				num: 1,
				size: 15,
				current:1
			}
		},
		components: {
			jinEdit,
			jinIcon
		},
		beforeMount(){
			getTags().then(res =>{
				for(var i =0; i < res.data.length; i++){
					this.tags.push(res.data[i].tag)
					this.tagsId.push(res.data[i].id)
				}
			})
		},
		methods: {
			newArticle(){
				this.isUpdate = false // 表示是否编辑文章
				this.index = 0
				let params = {
					html:'',
					word_count:0
				}

				// this.$refs.child.html = ``
				this.$refs.child.reLoadEditor(params);
			},
			bindPickerChange: function(e) {
				console.log('picker发送选择改变，携带值为', e.target.value)
				this.index = e.target.value
				console.log(this.tagsId[this.index])
			},
			open(){
			         // 通过组件定义的ref调用uni-popup方法
			         this.$refs.popup.open()
			      },
			upCallback() {
				let params = {
					current: this.current,
					size: 10,
					kind:2
				}
				getArticleList(params).then(res => {
					
					this.resList = res.data
					this.current++;
					//联网成功的回调,隐藏下拉刷新和上拉加载的状态;
					this.mescroll.endSuccess(this.resList.length);
					//设置列表数据
					if (this.num == 1) this.list = []; //如果是第一页需手动制空列表
					this.list = this.list.concat(this.resList); //追加新数据
				}).catch(() => {
					this.mescroll.endErr();
				})
			},
			// upCallback(page) {
			// 	let current = 1; // 页码, 默认从1开始
			// 	let size = 16; // 页长, 默认每页10条

			// 	uni.request({
			// 		url: '/api/article/user?current='+this.current+'&size='+this.size,
			// 		success: (data) => {

			// 			// 接口返回的当前页数据列表 (数组)
			// 			let curPageData = this.list1; 
			// 			// 接口返回的当前页数据长度 (如列表有26个数据,当前页返回8个,则curPageLen=8)
			// 			let curPageLen = 15; 
			// 			// 接口返回的总页数 (如列表有26个数据,每页10条,共3页; 则totalPage=3)
			// 			// let totalPage = 2; 
			// 			// 接口返回的总数据量(如列表有26个数据,每页10条,共3页; 则totalSize=26)
			// 			// let totalSize = data.xxx; 
			// 			// 接口返回的是否有下一页 (true/false)
			// 			// let hasNext = data.xxx; 

			// 			//设置列表数据
			// 			if(page.num == 1) this.list = []; //如果是第一页需手动置空列表
			// 			this.list = this.list.concat(curPageData); //追加新数据

			// 			// 请求成功,隐藏加载状态
			// 			//方法一(推荐): 后台接口有返回列表的总页数 totalPage
			// 			// this.mescroll.endByPage(curPageLen, totalPage); 

			// 			//方法二(推荐): 后台接口有返回列表的总数据量 totalSize
			// 			//this.mescroll.endBySize(curPageLen, totalSize); 

			// 			//方法三(推荐): 您有其他方式知道是否有下一页 hasNext
			// 			// this.mescroll.endSuccess(curPageLen, hasNext); 
			// 			//方法四 (不推荐),会存在一个小问题:比如列表共有20条数据,每页加载10条,共2页.
			// 			//如果只根据当前页的数据个数判断,则需翻到第三页才会知道无更多数据
			// 			//如果传了hasNext,则翻到第二页即可显示无更多数据.
			// 			if(this.size >= this.curPageLen){
			// 				setTimeout(()=>{
			// 					this.mescroll.endSuccess(this.curPageLen)
			// 				},20)
			// 			}
			// 			// 如果数据较复杂,可等到渲染完成之后再隐藏下拉加载状态: 如
			// 			// 建议使用setTimeout,因为this.$nextTick某些情况某些机型不触发
			// 			// setTimeout(()=>{
			// 			// 	this.mescroll.endSuccess(curPageLen)
			// 			// },20)

			// 			// 上拉加载的 curPageLen 必传, 否则会一直显示'加载中...':
			// 			// 1. 使配置的noMoreSize 和 empty生效
			// 			// 2. 判断是否有下一页的首要依据: 
			// 			//    当传的值小于page.size时(说明不满页了),则一定会认为无更多数据;
			// 			//    比传入的totalPage, totalSize, hasNext具有更高的判断优先级;
			// 			// 3. 当传的值等于page.size时(满页),才取totalPage, totalSize, hasNext判断是否有下一页
			// 			// 传totalPage, totalSize, hasNext目的是避免方法四描述的小问题

			// 			// 提示: 您无需额外维护页码和判断显示空布局,mescroll已自动处理好.
			// 			// 当您发现结果和预期不一样时, 建议再认真检查以上参数是否传正确
			// 			// current++;

			// 			this.curPageLen--;

			// 			fail: () => {
			// 				//  请求失败,隐藏加载状态
			// 				alert("Faile!")
			// 				this.mescroll.endErr()
			// 			}
			// 		}
			// 	})
			// 	this.current += 1;
			// 	// 此处仍可以继续写其他接口请求...
			// 	// 调用其他方法...
			// },
			
			showArticleInEditor(title,tag,content,word_count,item){
				console.log(item)
				this.isUpdate = true // 表示编辑文章
				this.index = tag
				this.title = title
				let params = {
					html:content,
					word_count:word_count
				}
				// this.$refs.child.html = content;
				this.$refs.child.reLoadEditor(params);
			},
			// 点击发布
			editOk(res) {
				let url = null;
				let describe = res.text.substr(0,20)
				for (let i = 0; i < res.delta.ops.length; i++) {
					if (res.delta.ops[i].insert.image) {
						url = res.delta.ops[i].insert.image
						return;
					}
				}
				var params = {
					title: this.title,
					content: res.html,
					image: url,
					kind: 1,
					tag: this.tagsId[this.index],
					type: 0,
					status: res.isPublic,
					word_count: res.textLength,
					describe:describe
				}
				switch(this.isUpdate){
					case false: 
						publishArticle(params).then( _res => {
							if (_res.code === 200) {
								console.log("发布成功")
							}
						})
						break;
					case true:
						upDateArticle(params).then( _res => {
							if(_res.code === 200){
								console.log("发布成功")
							}
						})
				}
				

			},
			eidtorChange(textLength) {
				this.word_count = textLength
				this.$refs.child.reLoadEditor();
			},
			setBold() {
				this.showBold = !this.showBold;
				this.$refs.child.setBold();
			},
			setItalic() {
				this.showItalic = !this.showItalic;
				this.$refs.child.setItalic();
			},
			setIns() {
				this.showIns = !this.showIns;
				this.$refs.child.setIns();
			},
			setHeader() {
				this.showHeader = !this.showHeader;
				this.$refs.child.setHeader();
			},
			setCenter() {
				this.showCenter = !this.showCenter;
				this.$refs.child.setCenter();
			},
			setRight() {
				this.showRight = !this.showRight;
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
			},
			release(num) {
				this.$refs.child.release(num)
			}
		}
	}
</script>

<style lang="scss" scoped>
	.tags-box{
		padding: 20rpx 0px 0rpx 120rpx;
	}
	/deep/.uni-list{
		background: rgba(191, 182, 172, 1.0) ;
		color: rgba(232, 232, 232, 1.0);
		border: 1px rgba(159, 159, 159, 1.0) solid;
		width: 50px;
		border-radius: 10%;
		text-align: center;
	}
	.page-warp {
		display: flex;
		flex-direction: column;
	}

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
		height: 1490rpx;
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
		padding: 100rpx 120rpx 0rpx 120rpx;
	}

	.title-text {
		color: rgb(118, 118, 118);

	}

	/deep/.container {
		box-sizing: border-box;
		/* border-top: 2rpx #d2d5cc solid; */
		padding: 30rpx 100rpx 0 100rpx !important;
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
		background: #dddddd;
		padding: 0 7% 0 30.3%;
	}

	.my-tool {
		height: 100rpx;
		display: flex;
		align-items: center;
		justify-content: space-around;
		width: 100%;
		background: #dddddd;
	}

	.body-box {
		background-color: rgb(241, 241, 241);
		padding: 48rpx 300rpx 0 300rpx;
		display: flex;
		flex: 1;
		justify-content: center;
		padding-bottom: 200rpx;
	}

	.body {
		width: 2500rpx;
		// flex: 1;
		// width: 100%;
		// min-width: 0;
		// min-height: 0;
		// display: -webkit-box;
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

	.body-left-header {
		height: 100%;
	}

	.body-left-box {
		/* width: 614rpx; */
		// flex: 1;
		height: 100%;
		overflow-y: scroll;
		width: 20%;
		height: 1662rpx;
	}

	.body-right-box {
		position: relative;
		max-height: inherit;
		// overflow: hidden auto;
		/* width: 1570rpx; */
		width: 71%;
		height: 1794rpx;

	}

	.body-right-foot-box {
		width: 100%;
		padding: 0 50rpx;
		height: 132rpx;
		color: rgba(80, 80, 80, 1);
		background-color: rgba(255, 255, 255, 1);
		font-size: 28rpx;
		box-shadow: 0px -2px 4px 0px rgba(241, 241, 241, 1);
		border-radius: 10rpx;
		display: flex;
		justify-content: space-between;
	}

	.btn-box,
	.text-box {
		display: flex;
		align-items: center;
	}

	.btn-box .u-btn {
		margin-right: 30rpx;
		flex: 1;
	}

	.add-article-btn {
		display: flex;
		align-items: center;
		padding-left: 50rpx;
		height: 100rpx;
		border-bottom: 1rpx solid #C0C0C0;
		/* background-color: #000000; */
	}

	.add-article-btn span {
		margin-left: 2rpx;
	}

	.article-card {
		height: 90prx;
		position: relative;
		/* display: flex; */
		/* align-items: center; */
		/* background-color: #55ffff; */
	}

	.article-card {
		padding: 30rpx 20rpx 10rpx 40rpx;
		border-bottom: 1rpx solid #d2d5cc;
		display: flex;
	}

	/deep/.u-image {
		display: flex;
		align-items: center;
	}

	.article-card-title {
		top: 10rpx;
		line-height: 150%;
	}

	.article-card-text-box {
		padding: 0 0 0 38rpx;
		display: flex;
		flex-direction: column;
		width: 100%;
	}

	.article-card-tag {
		color: rgba(128, 128, 128, 1);
		font-size: 22rpx;
		display: flex;
		align-self: flex-end;
	}

	/*根元素需要有固定的高度*/
	page {
		height: 100%;
		// 支付宝小程序,钉钉小程序需添加绝对定位,否则height:100%失效: https://opendocs.alipay.com/mini/framework/acss#%E5%B8%B8%E8%A7%81%E9%97%AE%E9%A2%98
		/* #ifdef MP-ALIPAY || MP-DINGTALK*/
		position: absolute;
		top: 0;
		left: 0;
		width: 100%;
		/* #endif */

		/*需给父元素设置height:100%*/
		.page-warp {
			height: 100%;
			display: flex;
			flex-direction: column;

			/* 顶部区域 */
			.top-warp {
				font-size: 28rpx;
				// padding: 20rpx;
				// text-align: center;
			}

			/* 中间 */
			.center-warp {
				flex: 1;
				// min-width: 0;
				// min-height: 0;/* 需给flex:1的元素加上最小高,否则内容超过会溢出容器 (如:小程序Android真机) */
				// border: 4px solid red;
				display: flex;

				// 左边
				.left-warp {
					// width: 180rpx;
					height: 100%;
					background-color: #eee;

					.tab {
						font-size: 28rpx;
						padding: 28rpx;

						&.active {
							background-color: #fff;
						}
					}
				}

				// 右边
				.right-warp {
					// flex: 1;
					// min-width: 0;
				}
			}

			/* 底部区域 */
			.bottom-warp {
				padding: 20rpx;
				text-align: center;
				background-color: #FF6990;
			}
		}
	}
</style>
