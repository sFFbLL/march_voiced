(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-search-index"],{"06b5":function(t,e,i){"use strict";i.r(e);var n=i("07c2"),a=i.n(n);for(var r in n)"default"!==r&&function(t){i.d(e,t,(function(){return n[t]}))}(r);e["default"]=a.a},"07c2":function(t,e,i){"use strict";i("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={name:"u-empty",props:{src:{type:String,default:""},text:{type:String,default:""},color:{type:String,default:"#c0c4cc"},iconColor:{type:String,default:"#c0c4cc"},iconSize:{type:[String,Number],default:120},fontSize:{type:[String,Number],default:26},mode:{type:String,default:"data"},imgWidth:{type:[String,Number],default:120},imgHeight:{type:[String,Number],default:"auto"},show:{type:Boolean,default:!0},marginTop:{type:[String,Number],default:0},iconStyle:{type:Object,default:function(){return{}}}},data:function(){return{icons:{car:"购物车为空",page:"页面不存在",search:"没有搜索结果",address:"没有收货地址",wifi:"没有WiFi",order:"订单为空",coupon:"没有优惠券",favor:"暂无收藏",permission:"无权限",history:"无历史记录",news:"无新闻列表",message:"消息列表为空",list:"列表为空",data:"数据为空"}}}};e.default=n},"098b":function(t,e,i){var n=i("67ea");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var a=i("4f06").default;a("531cf51e",n,!0,{sourceMap:!1,shadowMode:!1})},1589:function(t,e,i){"use strict";var n=i("69cc"),a=i.n(n);a.a},"17a9":function(t,e,i){"use strict";var n=i("9b75"),a=i.n(n);a.a},"17f2":function(t,e,i){"use strict";var n=i("88dd"),a=i.n(n);a.a},2297:function(t,e,i){"use strict";var n=i("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=n(i("d50f")),r=n(i("7cb0")),o=n(i("21b7")),c={data:function(){return{}},props:{articleInfo:{type:Object,default:{}}},components:{articleTitle:a.default,articleContent:r.default,articleInteract:o.default},methods:{},mounted:function(){}};e.default=c},"295f":function(t,e,i){"use strict";i("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={name:"u-search",props:{shape:{type:String,default:"round"},bgColor:{type:String,default:"#f2f2f2"},placeholder:{type:String,default:"请输入关键字"},clearabled:{type:Boolean,default:!0},focus:{type:Boolean,default:!1},showAction:{type:Boolean,default:!0},actionStyle:{type:Object,default:function(){return{}}},actionText:{type:String,default:"搜索"},inputAlign:{type:String,default:"left"},disabled:{type:Boolean,default:!1},animation:{type:Boolean,default:!1},borderColor:{type:String,default:"none"},value:{type:String,default:""},height:{type:[Number,String],default:64},inputStyle:{type:Object,default:function(){return{}}},maxlength:{type:[Number,String],default:"-1"},searchIconColor:{type:String,default:""},color:{type:String,default:"#606266"},placeholderColor:{type:String,default:"#909399"},margin:{type:String,default:"0"},searchIcon:{type:String,default:"search"}},data:function(){return{keyword:"",showClear:!1,show:!1,focused:this.focus}},watch:{keyword:function(t){this.$emit("input",t),this.$emit("change",t)},value:{immediate:!0,handler:function(t){this.keyword=t}}},computed:{showActionBtn:function(){return!(this.animation||!this.showAction)},borderStyle:function(){return this.borderColor?"1px solid ".concat(this.borderColor):"none"}},methods:{inputChange:function(t){this.keyword=t.detail.value},clear:function(){var t=this;this.keyword="",this.$nextTick((function(){t.$emit("clear")}))},search:function(t){this.$emit("search",t.detail.value);try{uni.hideKeyboard()}catch(t){}},custom:function(){this.$emit("custom",this.keyword);try{uni.hideKeyboard()}catch(t){}},getFocus:function(){this.focused=!0,this.animation&&this.showAction&&(this.show=!0),this.$emit("focus",this.keyword)},blur:function(){var t=this;setTimeout((function(){t.focused=!1}),100),this.show=!1,this.$emit("blur",this.keyword)},clickHandler:function(){this.disabled&&this.$emit("click")}}};e.default=n},3449:function(t,e,i){"use strict";i.r(e);var n=i("9d59"),a=i("ac5e");for(var r in a)"default"!==r&&function(t){i.d(e,t,(function(){return a[t]}))}(r);i("1589");var o,c=i("f0c5"),l=Object(c["a"])(a["default"],n["b"],n["c"],!1,null,"3783e340",null,!1,n["a"],o);e["default"]=l.exports},"3c05":function(t,e,i){"use strict";i.r(e);var n=i("e072"),a=i("d393");for(var r in a)"default"!==r&&function(t){i.d(e,t,(function(){return a[t]}))}(r);i("6772");var o,c=i("f0c5"),l=Object(c["a"])(a["default"],n["b"],n["c"],!1,null,"6d8f3acc",null,!1,n["a"],o);e["default"]=l.exports},"459f":function(t,e,i){var n=i("24fb");e=n(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-empty[data-v-21937aaa]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\nflex-direction:column;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;height:100%}.u-image[data-v-21937aaa]{margin-bottom:%?20?%}.u-slot-wrap[data-v-21937aaa]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;margin-top:%?20?%}',""]),t.exports=e},6772:function(t,e,i){"use strict";var n=i("098b"),a=i.n(n);a.a},"67a4":function(t,e,i){"use strict";i.d(e,"b",(function(){return a})),i.d(e,"c",(function(){return r})),i.d(e,"a",(function(){return n}));var n={uIcon:i("ac18").default},a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"u-section"},[i("v-uni-view",{staticClass:"u-section__title",class:{"u-section--line":t.showLine},style:{fontWeight:t.bold?"bold":"normal",color:t.color,fontSize:t.fontSize+"rpx",paddingLeft:t.showLine?.7*t.fontSize+"rpx":0}},[t.showLine?i("v-uni-view",{staticClass:"u-section__title__icon-wrap u-flex",style:[t.lineStyle]},[i("u-icon",{attrs:{top:"0",name:"column-line",size:1.25*t.fontSize,bold:!0,color:t.lineColor?t.lineColor:t.color}})],1):t._e(),i("v-uni-text",{staticClass:"u-flex u-section__title__text"},[t._v(t._s(t.title))])],1),t.right||t.$slots.right?i("v-uni-view",{staticClass:"u-section__right-info",style:{color:t.subColor},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.rightClick.apply(void 0,arguments)}}},[t.$slots.right?t._t("right"):[t._v(t._s(t.subTitle)),t.arrow?i("v-uni-view",{staticClass:"u-section__right-info__icon-arrow u-flex"},[i("u-icon",{attrs:{name:"arrow-right",size:"24",color:t.subColor}})],1):t._e()]],2):t._e()],1)},r=[]},"67ea":function(t,e,i){var n=i("24fb");e=n(!1),e.push([t.i,".search-article-item[data-v-6d8f3acc]{padding:%?20?% 0;overflow:hidden;border-bottom:1px solid #f0f0f0}[data-v-6d8f3acc] .article-title{line-height:100%}[data-v-6d8f3acc] .article-title .article-text{overflow:hidden;text-overflow:ellipsis;white-space:nowrap}[data-v-6d8f3acc] .article-content{color:#999;font-size:%?24?%;margin-top:%?10?%}[data-v-6d8f3acc] .article-interact{font-size:%?20?%;margin-top:%?10?%}[data-v-6d8f3acc] .interact-item uni-text{margin-left:%?4?%}[data-v-6d8f3acc] .point{visibility:hidden}.nickname[data-v-6d8f3acc]{display:inline-block;padding-right:%?20?%;max-width:%?200?%;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;line-height:100%}.time[data-v-6d8f3acc]{float:right;margin-right:%?20?%}",""]),t.exports=e},"69cc":function(t,e,i){var n=i("ea75");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var a=i("4f06").default;a("19950d46",n,!0,{sourceMap:!1,shadowMode:!1})},"6e3a":function(t,e,i){var n=i("24fb");e=n(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-section[data-v-38a44428]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;-webkit-box-align:center;-webkit-align-items:center;align-items:center;width:100%}.u-section__title[data-v-38a44428]{position:relative;font-size:%?28?%;padding-left:%?20?%;\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-align:center;-webkit-align-items:center;align-items:center}.u-section__title__icon-wrap[data-v-38a44428]{position:absolute}.u-section__title__text[data-v-38a44428]{line-height:1}.u-section__right-info[data-v-38a44428]{color:#909399;font-size:%?26?%;\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-align:center;-webkit-align-items:center;align-items:center}.u-section__right-info__icon-arrow[data-v-38a44428]{margin-left:%?6?%}',""]),t.exports=e},"6f0f":function(t,e,i){"use strict";var n=i("4ea4");i("99af"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=n(i("2909")),r=i("e220"),o=n(i("3c05")),c=n(i("d508")),l={data:function(){return{searchText:"",articleList:[],userList:[],loadStatus:"loading",isLoadMore:!1,articleCurrent:1,articleSize:10,userSize:4,articleList1:[{id:1,title:"所以监听用户的截图操作，提示<font color=red>用户</font>进行分，我还是个大学生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",content:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。<font color=red>在写作</font>一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",createTime:"2021-02-16",favourTotal:0,commentTotal:0,nickname:"张三"},{id:2,title:"个大学生啊，我该怎么学编程？我还是个大学生啊我该怎么学编程",content:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",createTime:"2021-02-16",favourTotal:0,commentTotal:0,nickname:"张三"},{id:3,title:"所以监听用户的截图操作，提示用户进生啊，我该怎么学编程？我还是个大学生啊，我该怎么学编程",content:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",createTime:"2021-02-16",favourTotal:0,commentTotal:0,nickname:"张三"},{id:4,title:"所以监听用户的截图操怎么学编程？我还是个大学生啊，我该怎么学编程",content:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",createTime:"2021-02-16",favourTotal:0,commentTotal:0,nickname:"张三"},{id:5,title:"所以监听用户的截图操怎么学编程？我还是个大学生啊，我该怎么学编程",content:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪我思存指责《甄嬛传》的作者流潋紫抄袭，另一件就是闹得沸沸扬扬的周冲洗稿六神磊磊今年春天在写作圈",createTime:"2021-02-16",favourTotal:0,commentTotal:0,nickname:"张三"}],userList1:[{id:1,nickname:"张三",avatarPath:i("f187")},{id:2,nickname:"张四",avatarPath:i("b9dd")},{id:3,nickname:"张五",avatarPath:i("f187")},{id:4,nickname:"张六",avatarPath:i("f187")}]}},components:{searchArticleItem:o.default,uniLoadMore:c.default},onReachBottom:function(){this.isLoadMore||(this.isLoadMore=!0,this.articleCurrent+=1,this.getList())},methods:{search:function(){this.searchText&&(this.userList=this.userList1,this.articleList=[].concat((0,a.default)(this.articleList),(0,a.default)(this.articleList1)))},clear:function(){this.searchText=""},getList:function(){if(this.searchText){var t=this;this.articleList.length>16?t.loadStatus="nomore":setTimeout((function(){t.isLoadMore=!1,t.articleList=[].concat((0,a.default)(t.articleList),(0,a.default)(t.articleList1))}),2e3)}},moreUser:function(){uni.navigateTo({url:"../search/searchUser"})},getArticleList:function(){var t=this;this.searchText,this.current,this.articleSize;(0,r.searchArticle)().then((function(e){t.articleList=[].concat((0,a.default)(t.articleList),(0,a.default)(e.data.articleList))}))},getUserList:function(){var t=this;this.searchText,this.userSize;(0,r.searchUser)().then((function(e){t.userList=e.data.userList}))}}};e.default=l},7454:function(t,e,i){var n=i("24fb");e=n(!1),e.push([t.i,".search[data-v-0fc645dc]{background-color:#f7f7f7}.nodate[data-v-0fc645dc],\n.search-user[data-v-0fc645dc],\n.search-article[data-v-0fc645dc],\n.search-header[data-v-0fc645dc]{background-color:#fff}.search-header[data-v-0fc645dc]{position:fixed;height:%?110?%;width:100%;line-height:%?100?%;background-color:#fff;border-bottom:%?1?% solid #f0f0f0;z-index:10}.search-input[data-v-0fc645dc]{width:90%;margin:auto}.content[data-v-0fc645dc]{padding-top:%?110?%}.nodate[data-v-0fc645dc]{min-height:%?800?%}.search-user[data-v-0fc645dc]{width:100%;padding:%?30?%}.user-list[data-v-0fc645dc]{margin-top:%?30?%}.user-item[data-v-0fc645dc]{display:inline-block;width:25%}.user-img[data-v-0fc645dc]{display:-webkit-box;display:-webkit-flex;display:flex;width:%?70?%;height:%?70?%;margin:auto;border-radius:50%}.divide[data-v-0fc645dc]{height:%?30?%;border-top:%?1?% solid #f0f0f0;border-bottom:%?1?% solid #f0f0f0}.search-article[data-v-0fc645dc]{padding:%?20?% %?30?% 0 %?30?%}.user-item .username[data-v-0fc645dc]{width:60%;margin:auto;margin-top:%?10?%;text-align:center;font-size:%?22?%;color:#303133;overflow:hidden;text-overflow:ellipsis;white-space:nowrap}[data-v-0fc645dc] .uni-load-more .uni-load-more__img{height:%?30?%!important;width:%?30?%!important}",""]),t.exports=e},7857:function(t,e,i){var n=i("7454");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var a=i("4f06").default;a("2ff7b1c1",n,!0,{sourceMap:!1,shadowMode:!1})},"88dd":function(t,e,i){var n=i("459f");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var a=i("4f06").default;a("ecf364d6",n,!0,{sourceMap:!1,shadowMode:!1})},"99a8":function(t,e,i){"use strict";i.r(e);var n=i("a5d1"),a=i("d9b9");for(var r in a)"default"!==r&&function(t){i.d(e,t,(function(){return a[t]}))}(r);i("cb58");var o,c=i("f0c5"),l=Object(c["a"])(a["default"],n["b"],n["c"],!1,null,"0fc645dc",null,!1,n["a"],o);e["default"]=l.exports},"9b75":function(t,e,i){var n=i("6e3a");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var a=i("4f06").default;a("604ebf22",n,!0,{sourceMap:!1,shadowMode:!1})},"9d59":function(t,e,i){"use strict";i.d(e,"b",(function(){return a})),i.d(e,"c",(function(){return r})),i.d(e,"a",(function(){return n}));var n={uIcon:i("ac18").default},a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"u-search",style:{margin:t.margin},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.clickHandler.apply(void 0,arguments)}}},[i("v-uni-view",{staticClass:"u-content",style:{backgroundColor:t.bgColor,borderRadius:"round"==t.shape?"100rpx":"10rpx",border:t.borderStyle,height:t.height+"rpx"}},[i("v-uni-view",{staticClass:"u-icon-wrap"},[i("u-icon",{staticClass:"u-clear-icon",attrs:{size:30,name:t.searchIcon,color:t.searchIconColor?t.searchIconColor:t.color}})],1),i("v-uni-input",{staticClass:"u-input",style:[{textAlign:t.inputAlign,color:t.color,backgroundColor:t.bgColor},t.inputStyle],attrs:{"confirm-type":"search",value:t.value,disabled:t.disabled,focus:t.focus,maxlength:t.maxlength,"placeholder-class":"u-placeholder-class",placeholder:t.placeholder,"placeholder-style":"color: "+t.placeholderColor,type:"text"},on:{blur:function(e){arguments[0]=e=t.$handleEvent(e),t.blur.apply(void 0,arguments)},confirm:function(e){arguments[0]=e=t.$handleEvent(e),t.search.apply(void 0,arguments)},input:function(e){arguments[0]=e=t.$handleEvent(e),t.inputChange.apply(void 0,arguments)},focus:function(e){arguments[0]=e=t.$handleEvent(e),t.getFocus.apply(void 0,arguments)}}}),t.keyword&&t.clearabled&&t.focused?i("v-uni-view",{staticClass:"u-close-wrap",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.clear.apply(void 0,arguments)}}},[i("u-icon",{staticClass:"u-clear-icon",attrs:{name:"close-circle-fill",size:"34",color:"#c0c4cc"}})],1):t._e()],1),i("v-uni-view",{staticClass:"u-action",class:[t.showActionBtn||t.show?"u-action-active":""],style:[t.actionStyle],on:{click:function(e){e.stopPropagation(),e.preventDefault(),arguments[0]=e=t.$handleEvent(e),t.custom.apply(void 0,arguments)}}},[t._v(t._s(t.actionText))])],1)},r=[]},a5d1:function(t,e,i){"use strict";i.d(e,"b",(function(){return a})),i.d(e,"c",(function(){return r})),i.d(e,"a",(function(){return n}));var n={uSearch:i("3449").default,uEmpty:i("f8b9").default,uSection:i("db13").default,uniLoadMore:i("d508").default},a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"search"},[i("v-uni-view",{staticClass:"search-header"},[i("v-uni-view",{staticClass:"search-input"},[i("u-search",{attrs:{"show-action":!1,placeholder:"搜索感兴趣的内容"},on:{search:function(e){arguments[0]=e=t.$handleEvent(e),t.search()},clear:function(e){arguments[0]=e=t.$handleEvent(e),t.clear()}},model:{value:t.searchText,callback:function(e){t.searchText=e},expression:"searchText"}})],1)],1),i("v-uni-view",{staticClass:"content"},[t.userList.length||t.articleList.length?t._e():i("v-uni-view",[i("u-empty",{staticClass:"nodate",attrs:{text:"没有数据",mode:"search"}})],1),t.userList.length?i("v-uni-view",{staticClass:"search-user"},[i("u-section",{attrs:{title:"相关用户","show-line":!1,bold:!1,color:"#969696","sub-color":"#969696"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.moreUser()}}}),i("v-uni-view",{staticClass:"user-list"},t._l(t.userList,(function(e,n){return i("v-uni-view",{staticClass:"user-item"},[i("v-uni-image",{staticClass:"user-img",attrs:{src:e.avatarPath,mode:"aspectFill"}}),i("v-uni-view",{staticClass:"username"},[t._v(t._s(e.nickname))])],1)})),1)],1):t._e(),t.userList.length&&t.articleList.length?i("v-uni-view",{staticClass:"divide"}):t._e(),t.articleList.length?i("v-uni-view",{staticClass:"search-article"},t._l(t.articleList,(function(t,e){return i("v-uni-view",{key:e},[i("searchArticleItem",{attrs:{articleInfo:t}})],1)})),1):t._e()],1),i("v-uni-view",{directives:[{name:"show",rawName:"v-show",value:t.isLoadMore,expression:"isLoadMore"}]},[i("uni-load-more",{staticClass:"loading",attrs:{status:t.loadStatus,iconType:"circle"}})],1)],1)},r=[]},ac5e:function(t,e,i){"use strict";i.r(e);var n=i("295f"),a=i.n(n);for(var r in n)"default"!==r&&function(t){i.d(e,t,(function(){return n[t]}))}(r);e["default"]=a.a},b9dd:function(t,e,i){t.exports=i.p+"static/img/1.e963bd74.jpg"},bfed:function(t,e,i){"use strict";i.d(e,"b",(function(){return a})),i.d(e,"c",(function(){return r})),i.d(e,"a",(function(){return n}));var n={uIcon:i("ac18").default},a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return t.show?i("v-uni-view",{staticClass:"u-empty",style:{marginTop:t.marginTop+"rpx"}},[i("u-icon",{attrs:{name:t.src?t.src:"empty-"+t.mode,"custom-style":t.iconStyle,label:t.text?t.text:t.icons[t.mode],"label-pos":"bottom","label-color":t.color,"label-size":t.fontSize,size:t.iconSize,color:t.iconColor,"margin-top":"14"}}),i("v-uni-view",{staticClass:"u-slot-wrap"},[t._t("bottom")],2)],1):t._e()},r=[]},cb58:function(t,e,i){"use strict";var n=i("7857"),a=i.n(n);a.a},d393:function(t,e,i){"use strict";i.r(e);var n=i("2297"),a=i.n(n);for(var r in n)"default"!==r&&function(t){i.d(e,t,(function(){return n[t]}))}(r);e["default"]=a.a},d9b9:function(t,e,i){"use strict";i.r(e);var n=i("6f0f"),a=i.n(n);for(var r in n)"default"!==r&&function(t){i.d(e,t,(function(){return n[t]}))}(r);e["default"]=a.a},db13:function(t,e,i){"use strict";i.r(e);var n=i("67a4"),a=i("faf5");for(var r in a)"default"!==r&&function(t){i.d(e,t,(function(){return a[t]}))}(r);i("17a9");var o,c=i("f0c5"),l=Object(c["a"])(a["default"],n["b"],n["c"],!1,null,"38a44428",null,!1,n["a"],o);e["default"]=l.exports},e072:function(t,e,i){"use strict";var n;i.d(e,"b",(function(){return a})),i.d(e,"c",(function(){return r})),i.d(e,"a",(function(){return n}));var a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"search-article-item"},[i("articleTitle",{attrs:{articleTitle:t.articleInfo.title,articleId:t.articleInfo.articleId}}),i("articleContent",{attrs:{articleContent:t.articleInfo.content,articleId:t.articleInfo.id}}),i("articleInteract",{attrs:{favourTotal:t.articleInfo.favourTotal,commentTotal:t.articleInfo.commentTotal}},[i("v-uni-text",{staticClass:"nickname",attrs:{slot:"before"},slot:"before"},[t._v(t._s(t.articleInfo.nickname))]),i("v-uni-text",{staticClass:"time",attrs:{slot:"after"},slot:"after"},[t._v(t._s(t.articleInfo.createTime))])],1)],1)},r=[]},e220:function(t,e,i){"use strict";var n=i("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.getRecommend=o,e.getFollow=c,e.searchArticle=l,e.searchUser=s;var a=n(i("367e")),r=n(i("e8a0"));function o(t){return a.default.get(r.default+"/api/article/top",{params:t})}function c(t){return a.default.get(r.default+"/api/base/follow",{params:t})}function l(t){return a.default.get(r.default+"/api/base/searchArticle",{params:t})}function s(t){return a.default.get(r.default+"/api/base/searchUser",{params:t})}},ea75:function(t,e,i){var n=i("24fb");e=n(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-search[data-v-3783e340]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:1;-webkit-flex:1;flex:1}.u-content[data-v-3783e340]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-align:center;-webkit-align-items:center;align-items:center;padding:0 %?18?%;-webkit-box-flex:1;-webkit-flex:1;flex:1}.u-clear-icon[data-v-3783e340]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-align:center;-webkit-align-items:center;align-items:center}.u-input[data-v-3783e340]{-webkit-box-flex:1;-webkit-flex:1;flex:1;font-size:%?28?%;line-height:1;margin:0 %?10?%;color:#909399}.u-close-wrap[data-v-3783e340]{width:%?40?%;height:100%;\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;border-radius:50%}.u-placeholder-class[data-v-3783e340]{color:#909399}.u-action[data-v-3783e340]{font-size:%?28?%;color:#303133;width:0;overflow:hidden;-webkit-transition:all .3s;transition:all .3s;white-space:nowrap;text-align:center}.u-action-active[data-v-3783e340]{width:%?80?%;margin-left:%?10?%}',""]),t.exports=e},f187:function(t,e,i){t.exports=i.p+"static/img/2.c15f2abc.jpg"},f66a:function(t,e,i){"use strict";i("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={name:"u-section",props:{title:{type:String,default:""},subTitle:{type:String,default:"更多"},right:{type:Boolean,default:!0},fontSize:{type:[Number,String],default:28},bold:{type:Boolean,default:!0},color:{type:String,default:"#303133"},subColor:{type:String,default:"#909399"},showLine:{type:Boolean,default:!0},lineColor:{type:String,default:""},arrow:{type:Boolean,default:!0}},computed:{lineStyle:function(){return{left:-.9*Number(this.fontSize)+"rpx",top:-Number(this.fontSize)*("ios"==this.$u.os()?.14:.15)+"rpx"}}},methods:{rightClick:function(){this.$emit("click")}}};e.default=n},f8b9:function(t,e,i){"use strict";i.r(e);var n=i("bfed"),a=i("06b5");for(var r in a)"default"!==r&&function(t){i.d(e,t,(function(){return a[t]}))}(r);i("17f2");var o,c=i("f0c5"),l=Object(c["a"])(a["default"],n["b"],n["c"],!1,null,"21937aaa",null,!1,n["a"],o);e["default"]=l.exports},faf5:function(t,e,i){"use strict";i.r(e);var n=i("f66a"),a=i.n(n);for(var r in n)"default"!==r&&function(t){i.d(e,t,(function(){return n[t]}))}(r);e["default"]=a.a}}]);