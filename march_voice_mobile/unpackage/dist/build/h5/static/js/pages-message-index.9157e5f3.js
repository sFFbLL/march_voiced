(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-message-index"],{"029d":function(t,e,n){"use strict";n.r(e);var a=n("c48b"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},"04ff":function(t,e,n){var a=n("4675");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("674bf75a",a,!0,{sourceMap:!1,shadowMode:!1})},"0ef7":function(t,e,n){"use strict";var a=n("c4fe"),i=n.n(a);i.a},1704:function(t,e,n){"use strict";var a=n("3842"),i=n.n(a);i.a},"1fc3":function(t,e,n){"use strict";n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var a={uBadge:n("cb97").default},i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"tab-card"},[n("v-uni-view",{staticClass:"head-nav"},t._l(t.tabs,(function(e,a){return n("v-uni-view",{key:e.index,staticClass:"head-nav-item",class:e.isActive?"active":"",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.handlerActive(e.index)}}},[n("v-uni-view",{staticClass:"center"},[t._v(t._s(e.value)),0==e.index&&0==t.interactRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),1==e.index&&0==t.attentionRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),2==e.index&&0==t.otherRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),n("v-uni-view",{staticClass:"head-nav-bottom"})],1)],1)})),1),t._t("default")],2)},r=[]},"26ab":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n/* 评论样式 */.comment[data-v-107e6a24]{margin-left:%?110?%;font-size:%?28?%;color:#505050;margin-top:%?10?%;margin-bottom:%?10?%;width:%?600?%;overflow:hidden;-webkit-line-clamp:1;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical}\n/* 头部插槽动作 */.slot[data-v-107e6a24]{margin-left:%?15?%;font-size:%?28?%;color:#505050;font-weight:300}\n/* 文章和想法的样式 */.thecontent[data-v-107e6a24]{margin-left:%?110?%;width:%?600?%;height:%?140?%;background-color:#f1f1f1;border-radius:%?8?%;display:-webkit-box;display:-webkit-flex;display:flex;padding:%?20?% %?20?%;margin-bottom:%?30?%;margin-top:%?6?%}.thewords[data-v-107e6a24]{width:%?1200?%}.theimg[data-v-107e6a24]{width:%?300?%;height:%?100?%;border-radius:%?4?%;margin-left:%?10?%}\n/* 公共头像组件样式 */[data-v-107e6a24] .attention-cell .flex-item{border:none;padding:%?10?%;margin-top:%?17?%}",""]),t.exports=e},"349b":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={data:function(){return{}},props:{articleId:{type:Number,default:null},articleTitle:{type:String,default:""}},components:{},methods:{articleDetails:function(){uni.navigateTo({url:"../articleDetails/index?id="+this.articleId})}},mounted:function(){}};e.default=a},"34e3":function(t,e,n){"use strict";n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var a={uGap:n("53ac").default},i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{staticClass:"wrap"},[a("attentionAndFansCell",{attrs:{aid:t.messageInfo.userId,nickname:t.messageInfo.nickname,avatarPath:t.messageInfo.avatarPath},scopedSlots:t._u([{key:"afterNicknameText",fn:function(){return[a("v-uni-text",{staticClass:"slot"},[t._v(t._s(t.tip))]),3==t.messageInfo.status&&1==t.messageInfo.type?a("v-uni-image",{staticStyle:{width:"44rpx",height:"30rpx","margin-left":"10rpx"},attrs:{src:n("61ee"),mode:""}}):t._e(),2==t.messageInfo.status&&1==t.messageInfo.type?a("v-uni-image",{staticStyle:{width:"34rpx",height:"30rpx","margin-left":"10rpx"},attrs:{src:n("4907"),mode:""}}):t._e(),4==t.messageInfo.status&&1==t.messageInfo.type?a("v-uni-image",{staticStyle:{width:"34rpx",height:"30rpx","margin-left":"10rpx"},attrs:{src:n("c7cd"),mode:""}}):t._e()]},proxy:!0},{key:"underText",fn:function(){return[a("v-uni-text",[t._v(t._s(t.createDate))])]},proxy:!0}])}),t.messageInfo.comment?a("v-uni-text",{staticClass:"comment"},[t._v(t._s(t.messageInfo.comment))]):t._e(),a("v-uni-view",{staticClass:"thecontent"},[0==t.messageInfo.type?a("articleTitle",{staticClass:"thewords",attrs:{articleTitle:t.messageInfo.title,articleId:t.messageInfo.articleId}}):t._e(),1==t.messageInfo.type?a("articleTitle",{staticClass:"thewords",attrs:{articleTitle:t.messageInfo.content,articleId:t.messageInfo.articleId}}):t._e(),t.messageInfo.image?a("v-uni-image",{staticClass:"theimg",attrs:{src:t.messageInfo.image,mode:"scaleToFill"}}):t._e()],1),a("u-gap",{attrs:{height:"10","bg-color":"#f5f5f5"}})],1)},r=[]},3842:function(t,e,n){var a=n("26ab");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("2ec1631c",a,!0,{sourceMap:!1,shadowMode:!1})},"3db2":function(t,e,n){"use strict";n.r(e);var a=n("9472"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},4023:function(t,e,n){"use strict";var a=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i=a(n("d50f")),r=a(n("40e6")),o={props:{messageInfo:{type:Object,default:function(){return{type:0,status:1,title:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",articleId:1,comment:"哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",content:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",image:n("f187"),createTime:"2020/12/12",userId:"1",avatarPath:n("b9dd"),nickname:"张三"}}}},components:{attentionAndFansCell:r.default,articleTitle:i.default},data:function(){return{tip:"赞了你的文章"}},computed:{createDate:function(){var t=new Date;return t.setTime(1e3*this.messageInfo.createTime),console.log(t.toLocaleDateString()),t.toLocaleDateString()}},created:function(){0==this.messageInfo.type?1==this.messageInfo.status?this.tip="评论了你的文章":2==this.messageInfo.status?this.tip="收藏了你的文章":3==this.messageInfo.status?this.tip="转发了你的文章":4==this.messageInfo.status&&(this.tip="赞了你的文章"):1==this.messageInfo.type&&(1==this.messageInfo.status?this.tip="评论了你的想法":0==this.messageInfo.status||2!=this.messageInfo.status&&3!=this.messageInfo.status&&4!=this.messageInfo.status||(this.tip="送了你一个"))}};e.default=o},4213:function(t,e,n){"use strict";n.r(e);var a=n("c21f"),i=n("cf2e");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("0ef7");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"a16dac8e",null,!1,a["a"],o);e["default"]=c.exports},"43df":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{staticClass:"wrap"},[0==t.otherList.status?a("v-uni-image",{staticStyle:{width:"30rpx",height:"30rpx",margin:"0 10rpx"},attrs:{src:n("6d64"),mode:"scaleToFill"}}):t._e(),1==t.otherList.status?a("v-uni-image",{staticStyle:{width:"30rpx",height:"30rpx",margin:"0 10rpx"},attrs:{src:n("f982"),mode:"scaleToFill"}}):t._e(),0==t.otherList.type?a("v-uni-text",[t._v("您的文章")]):t._e(),0==t.otherList.type?a("v-uni-navigator",{staticStyle:{color:"#007AFF",display:"inline-block"},attrs:{url:t.url}},[t._v("《"+t._s(t.otherList.title)+"》")]):t._e(),a("v-uni-text",{staticClass:"tipWords"},[t._v(t._s(t.tipWords))]),a("v-uni-view",{staticClass:"time"},[a("v-uni-text",[t._v(t._s(t.createDate))])],1)],1)},r=[]},"463f":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return t.show?n("v-uni-view",{staticClass:"u-badge",class:[t.isDot?"u-badge-dot":"","mini"==t.size?"u-badge-mini":"",t.type?"u-badge--bg--"+t.type:""],style:[{top:t.offset[0]+"rpx",right:t.offset[1]+"rpx",fontSize:t.fontSize+"rpx",position:t.absolute?"absolute":"static",color:t.color,backgroundColor:t.bgColor},t.boxStyle]},[t._v(t._s(t.showText))]):t._e()},r=[]},4675:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-badge[data-v-34077204]{display:-webkit-inline-box;display:-webkit-inline-flex;display:inline-flex;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;line-height:%?24?%;padding:%?4?% %?8?%;border-radius:%?100?%;z-index:9}.u-badge--bg--primary[data-v-34077204]{background-color:#2979ff}.u-badge--bg--error[data-v-34077204]{background-color:#fa3534}.u-badge--bg--success[data-v-34077204]{background-color:#19be6b}.u-badge--bg--info[data-v-34077204]{background-color:#909399}.u-badge--bg--warning[data-v-34077204]{background-color:#f90}.u-badge-dot[data-v-34077204]{height:%?16?%;width:%?16?%;border-radius:%?100?%;line-height:1}.u-badge-mini[data-v-34077204]{-webkit-transform:scale(.8);transform:scale(.8);-webkit-transform-origin:center center;transform-origin:center center}.u-info[data-v-34077204]{background-color:#909399;color:#fff}',""]),t.exports=e},4788:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".article-title .article-text[data-v-74809812]{color:#404040;font-size:%?34?%;font-family:微软雅黑;text-align:left;font-weight:700;overflow:hidden;-webkit-line-clamp:2;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;line-height:150%}.article-text[data-v-74809812]{display:inline-block;width:100%}",""]),t.exports=e},4907:function(t,e,n){t.exports=n.p+"static/img/xiaoku.e5dd8801.png"},"50fc":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.head-nav[data-v-06dc7823]{display:-webkit-flex;display:-webkit-box;display:flex;overflow:hidden;height:100%}.head-nav .head-nav-item[data-v-06dc7823]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center;font-size:%?34?%;font-family:Microsoft YaHei;font-weight:700;letter-spacing:%?4?%;color:#999}.head-nav-item .center[data-v-06dc7823]{position:relative;margin:auto;width:%?228?%}.active .center[data-v-06dc7823]{color:#404040}.active .center .head-nav-bottom[data-v-06dc7823]{position:relative;margin:auto;width:50%;height:%?8?%;top:%?-8?%;background-color:#404040}',""]),t.exports=e},"53ac":function(t,e,n){"use strict";n.r(e);var a=n("7e61"),i=n("7f39");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("c788");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"71495536",null,!1,a["a"],o);e["default"]=c.exports},5742:function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={name:"u-gap",props:{bgColor:{type:String,default:"transparent "},height:{type:[String,Number],default:30},marginTop:{type:[String,Number],default:0},marginBottom:{type:[String,Number],default:0}},computed:{gapStyle:function(){return{backgroundColor:this.bgColor,height:this.height+"rpx",marginTop:this.marginTop+"rpx",marginBottom:this.marginBottom+"rpx"}}}};e.default=a},"5e05":function(t,e,n){"use strict";var a=n("8199"),i=n.n(a);i.a},"61ee":function(t,e,n){t.exports=n.p+"static/img/aixin.39c2cdd4.png"},"6d64":function(t,e,n){t.exports=n.p+"static/img/success.74305a2d.png"},"7b92":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-content",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.toDetails.apply(void 0,arguments)}}},[n("v-uni-view",[n("v-uni-text",{ref:"articleContent",staticClass:"article-text",domProps:{innerHTML:t._s(t.textContent)}},[t._v(t._s(t.textContent)+"123")])],1),""!=t.articleImg?n("v-uni-view",{staticClass:"article-img"},[n("v-uni-image",{attrs:{src:t.articleImg,mode:"aspectFill"}})],1):t._e()],1)},r=[]},"7cb0":function(t,e,n){"use strict";n.r(e);var a=n("7b92"),i=n("029d");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("5e05");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"9100d782",null,!1,a["a"],o);e["default"]=c.exports},"7e61":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"u-gap",style:[t.gapStyle]})},r=[]},"7f39":function(t,e,n){"use strict";n.r(e);var a=n("5742"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},8199:function(t,e,n){var a=n("c10a");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("b58bd422",a,!0,{sourceMap:!1,shadowMode:!1})},"8a32":function(t,e,n){"use strict";n.r(e);var a=n("8c76"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},"8b2c":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-title",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.articleDetails.apply(void 0,arguments)}}},[n("v-uni-view",[n("v-uni-text",{staticClass:"article-text",domProps:{innerHTML:t._s(t.articleTitle)}})],1)],1)},r=[]},"8bc3":function(t,e,n){"use strict";n.r(e);var a=n("923a"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},"8c76":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={name:"u-badge",props:{type:{type:String,default:"error"},size:{type:String,default:"default"},isDot:{type:Boolean,default:!1},count:{type:[Number,String]},overflowCount:{type:Number,default:99},showZero:{type:Boolean,default:!1},offset:{type:Array,default:function(){return[20,20]}},absolute:{type:Boolean,default:!0},fontSize:{type:[String,Number],default:"24"},color:{type:String,default:"#ffffff"},bgColor:{type:String,default:""},isCenter:{type:Boolean,default:!1}},computed:{boxStyle:function(){var t={};return this.isCenter?(t.top=0,t.right=0,t.transform="translateY(-50%) translateX(50%)"):(t.top=this.offset[0]+"rpx",t.right=this.offset[1]+"rpx",t.transform="translateY(0) translateX(0)"),"mini"==this.size&&(t.transform=t.transform+" scale(0.8)"),t},showText:function(){return this.isDot?"":this.count>this.overflowCount?"".concat(this.overflowCount,"+"):this.count},show:function(){return 0!=this.count||0!=this.showZero}}};e.default=a},"8cbe":function(t,e,n){"use strict";var a=n("4ea4");n("99af"),n("d81d"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i=a(n("2909")),r=a(n("d6e3")),o=a(n("7cb0")),s=a(n("99f6")),c=a(n("9752")),u=a(n("40e6")),d=n("5642"),l={components:{tabs:r.default,articleContent:o.default,singleMessage:s.default,attentionAndFansCell:u.default,otherMessage:c.default},data:function(){return{interactCurrent:1,attentionCurrent:1,otherCurrent:1,size:10,isLoadMore:!1,loadStatus:"loading",interactLoadStatus:"loading",attentionLoadStatus:"loading",otherLoadStatus:"loading",isLoading:!1,isDisabled:!1,isAttention:!1,tabIndex:0,message:!0,attentionRead:!1,otherRead:!1,interactRead:!1,tablist:[{index:0,value:"互动消息",isActive:!0},{index:1,value:"关注消息",isActive:!1},{index:2,value:"系统消息",isActive:!1}],interactList:{pages:"22",total:"100",current:"1",size:"1",records:[]},attentionList:{pages:"22",total:"100",current:"1",size:"1",records:[]},otherList:{pages:"22",total:"100",current:"1",size:"1",records:[]}}},onReachBottom:function(){this.isLoadMore||0!=this.tabIndex?this.isLoadMore||1!=this.tabIndex?this.isLoadMore||2!=this.tabIndex||(this.isLoadMore=!0,this.otherCurrent+=1,this.other()):(this.isLoadMore=!0,this.attentionCurrent+=1,this.attention()):(this.isLoadMore=!0,this.interactCurrent+=1,this.interact())},onPullDownRefresh:function(){var t=this;(0,d.unreadMessage)(1).then((function(e){e.data.count>0&&t.$store.commit("changeInteract",0)})),(0,d.unreadMessage)(2).then((function(e){e.data.count>0&&t.$store.commit("changeAttention",0)})),(0,d.unreadMessage)(3).then((function(e){e.data.count>0&&t.$store.commit("changeOther",0)})),0==this.tabIndex?(this.interact(),this.$store.commit("changeInteract",1)):1==this.tabIndex?this.$store.commit("changeAttention",1):2==this.tabIndex&&this.$store.commit("changeOther",1),setTimeout((function(){uni.stopPullDownRefresh()}),1e3)},onShow:function(){uni.hideTabBarRedDot({index:3})},created:function(){var t=this;this.$store.commit("changeInteract",1),(0,d.readMessage)(1).then((function(t){})),this.interact(),(0,d.unreadMessage)(2).then((function(e){e.data.count>0&&t.$store.commit("changeAttention",0)})),(0,d.unreadMessage)(3).then((function(e){e.data.count>0&&t.$store.commit("changeOther",0)}))},methods:{tabActive:function(t){this.tablist.map((function(e,n){e.isActive=t==n})),0==t?(this.loadStatus=this.interactLoadStatus,this.interact()):1==t?(this.loadStatus=this.attentionLoadStatus,this.attention()):(this.loadStatus=this.otherLoadStatus,this.other()),this.tabIndex=t},interact:function(){var t=this,e={current:this.interactCurrent,size:this.size};(0,d.interactList)(e).then((function(e){t.interactList.records=[].concat((0,i.default)(t.interactList.records),(0,i.default)(e.data.records)),e.data.length<=t.size&&(t.loadStatus="nomore")})),this.interactList.length>16?(t.loadStatus="nomore",t.interactLoadStatus="nomore"):1===this.interactCurrent?t.isLoadMore=!1:setTimeout((function(){t.isLoadMore=!1}),2e3),(0,d.readMessage)(1).then((function(t){}))},attention:function(){var t=this,e={current:this.attentionCurrent,size:this.size};(0,d.attentionList)(e).then((function(e){t.attentionList.records=[].concat((0,i.default)(t.attentionList.records),(0,i.default)(e.data.records)),e.data.length<=t.size&&(t.loadStatus="nomore")})),this.attentionList.length>16?(t.loadStatus="nomore",t.attentionLoadStatus="nomore"):1===this.attentionCurrent?t.isLoadMore=!1:setTimeout((function(){t.isLoadMore=!1}),2e3),(0,d.readMessage)(2).then((function(t){}))},other:function(){var t=this,e={current:this.otherCurrent,size:this.size};(0,d.otherList)(e).then((function(e){t.otherList.records=[].concat((0,i.default)(t.otherList.records),(0,i.default)(e.data.records)),e.data.length<=t.size&&(t.loadStatus="nomore")})),this.otherList.length>16?(t.loadStatus="nomore",t.otherLoadStatus="nomore"):1===this.otherCurrent?t.isLoadMore=!1:setTimeout((function(){t.isLoadMore=!1}),2e3),(0,d.readMessage)(3).then((function(t){}))}}};e.default=l},"923a":function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={data:function(){return{}},props:{tabs:{type:Array,default:[]},type:{type:Boolean}},methods:{handlerActive:function(t){0==t?this.$store.commit("changeInteract",1):1==t?this.$store.commit("changeAttention",1):2==t&&this.$store.commit("changeOther",1),this.$emit("tabActive",t)}},computed:{interactRead:function(){return this.$store.state.interactRead},attentionRead:function(){return this.$store.state.attentionRead},otherRead:function(){return this.$store.state.otherRead}},mounted:function(){},created:function(){}};e.default=a},"936b":function(t,e,n){var a=n("ad20");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("7c85361e",a,!0,{sourceMap:!1,shadowMode:!1})},9472:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={props:{otherList:{type:Object}},data:function(){return{tipWords:"",url:"../../pages/articleDetails/index?id="}},computed:{createDate:function(){var t=new Date;return t.setTime(1e3*this.otherList.createTime),console.log(t.toLocaleDateString()),t.toLocaleDateString()}},created:function(){0==this.otherList.type?(this.url="../../pages/articleDetails/index?id="+this.otherList.articleId,0==this.otherList.status?this.tipWords="通过审核，已为您自动发布。":1==this.otherList.status&&(this.tipWords="含有敏感信息，请进行修正后重新发布。")):1==this.otherList.type&&(0==this.otherList.status?this.tipWords="恭喜你成功加入三月圈，快去看看吧。":1==this.otherList.status&&(this.tipWords="由于权限不足，未能成功加入三月圈。"))}};e.default=a},9752:function(t,e,n){"use strict";n.r(e);var a=n("43df"),i=n("3db2");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("d2ce");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"0ea76f1c",null,!1,a["a"],o);e["default"]=c.exports},"99f6":function(t,e,n){"use strict";n.r(e);var a=n("34e3"),i=n("b9b9");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("1704");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"107e6a24",null,!1,a["a"],o);e["default"]=c.exports},ac84:function(t,e,n){var a=n("c6f8");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("2166a651",a,!0,{sourceMap:!1,shadowMode:!1})},ad20:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/',""]),t.exports=e},b37c:function(t,e,n){var a=n("50fc");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("0ffdf5f4",a,!0,{sourceMap:!1,shadowMode:!1})},b9b9:function(t,e,n){"use strict";n.r(e);var a=n("4023"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},bea5:function(t,e,n){var a=n("4788");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("332ec73e",a,!0,{sourceMap:!1,shadowMode:!1})},c10a:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".article-content[data-v-9100d782]{margin-top:%?18?%;display:-webkit-box;display:-webkit-flex;display:flex;font-size:%?28?%;font-family:PingFangSC;color:#404040}.article-text[data-v-9100d782]{overflow:hidden;-webkit-line-clamp:3;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;line-height:150%}.article-img[data-v-9100d782]{display:inline-block;margin-left:%?30?%}.article-img uni-image[data-v-9100d782]{width:%?170?%;height:100%;background-color:#eee;border-radius:4px}",""]),t.exports=e},c21f:function(t,e,n){"use strict";n.d(e,"b",(function(){return i})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return a}));var a={uGap:n("53ac").default,uniLoadMore:n("d508").default},i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"message"},[n("v-uni-view",{staticClass:"header"},[n("tabs",{staticClass:"tag-nav",attrs:{type:t.message,tabs:t.tablist},on:{tabActive:function(e){arguments[0]=e=t.$handleEvent(e),t.tabActive.apply(void 0,arguments)}}})],1),n("v-uni-view",{staticClass:"content"},[t._l(t.interactList.records,(function(e){return 0==t.tabIndex?n("v-uni-view",[n("singleMessage",{attrs:{messageInfo:e}})],1):t._e()})),t._l(t.attentionList.records,(function(e){return 1==t.tabIndex?n("v-uni-view",[n("attentionAndFansCell",{staticClass:"attention",attrs:{aid:e.userId,nickname:e.nickname,avatarPath:e.avatarPath,isFollow:e.isFollow},scopedSlots:t._u([{key:"afterNicknameText",fn:function(){return[n("v-uni-text",{staticClass:"slot"},[t._v("关注了你")])]},proxy:!0}],null,!0)})],1):t._e()})),t._l(t.otherList.records,(function(e){return 2==t.tabIndex?n("v-uni-view",[n("otherMessage",{attrs:{otherList:e}}),n("u-gap",{attrs:{height:"2","bg-color":"#f5f5f5"}})],1):t._e()}))],2),n("v-uni-view",{directives:[{name:"show",rawName:"v-show",value:t.isLoadMore,expression:"isLoadMore"}]},[n("uni-load-more",{attrs:{status:t.loadStatus}})],1)],1)},r=[]},c48b:function(t,e,n){"use strict";n("a9e3"),n("ac1f"),n("5319"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={data:function(){return{textContent:""}},props:{id:{type:Number,default:null},articleContent:{type:String,default:""},articleImg:{type:String,default:""},isIdea:{type:Boolean,default:!1}},components:{},methods:{toDetails:function(){this.isIdea?uni.navigateTo({url:"../ideaDetails/index?id="+this.id}):uni.navigateTo({url:"../articleDetails/index?id="+this.id})}},created:function(){this.textContent=this.articleContent},mounted:function(){this.textContent=this.$refs.articleContent.$el.innerText.replace(/\ +/g,"").replace(/[\r\n]/g,"")}};e.default=a},c4fe:function(t,e,n){var a=n("fb4c");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=n("4f06").default;i("6ef5a7c0",a,!0,{sourceMap:!1,shadowMode:!1})},c6f8:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".wrap[data-v-0ea76f1c]{padding:%?50?% %?10?%}.time[data-v-0ea76f1c]{margin-top:%?10?%;font-size:%?20?%;color:#616367;margin-left:%?10?%}",""]),t.exports=e},c788:function(t,e,n){"use strict";var a=n("936b"),i=n.n(a);i.a},c7cd:function(t,e,n){t.exports=n.p+"static/img/qingzhu.a25da244.png"},c905:function(t,e,n){"use strict";var a=n("bea5"),i=n.n(a);i.a},cb97:function(t,e,n){"use strict";n.r(e);var a=n("463f"),i=n("8a32");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("d67e");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"34077204",null,!1,a["a"],o);e["default"]=c.exports},cf2e:function(t,e,n){"use strict";n.r(e);var a=n("8cbe"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},d040:function(t,e,n){"use strict";var a=n("b37c"),i=n.n(a);i.a},d2ce:function(t,e,n){"use strict";var a=n("ac84"),i=n.n(a);i.a},d50f:function(t,e,n){"use strict";n.r(e);var a=n("8b2c"),i=n("f708");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("c905");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"74809812",null,!1,a["a"],o);e["default"]=c.exports},d67e:function(t,e,n){"use strict";var a=n("04ff"),i=n.n(a);i.a},d6e3:function(t,e,n){"use strict";n.r(e);var a=n("1fc3"),i=n("8bc3");for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);n("d040");var o,s=n("f0c5"),c=Object(s["a"])(i["default"],a["b"],a["c"],!1,null,"06dc7823",null,!1,a["a"],o);e["default"]=c.exports},f708:function(t,e,n){"use strict";n.r(e);var a=n("349b"),i=n.n(a);for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);e["default"]=i.a},f982:function(t,e,n){t.exports=n.p+"static/img/fail.538da67e.png"},fb4c:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".header[data-v-a16dac8e]{position:fixed;top:0;background-color:#fff;z-index:99;width:100%;height:%?80?%;line-height:%?80?%}\n\n/* 公共头像组件样式 */[data-v-a16dac8e] .attention-cell .flex-item{padding:%?10?%;margin-top:%?17?%}\n\n/* 关注插槽 */.slot[data-v-a16dac8e]{margin-left:%?20?%;font-weight:400}[data-v-a16dac8e] .tab-card .head-nav .head-nav-bottom{width:70%}.flex-item[data-v-a16dac8e]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-flex-wrap:wrap;flex-wrap:wrap;box-sizing:border-box;border-bottom:%?6?% solid #f0f0f0;padding:%?20?% %?30?% %?20?% %?30?%\n\t/* margin-bottom: 20rpx; */}\n\n/* 左侧图片盒子样式 */.left-img-box[data-v-a16dac8e]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;\n\t/* flex: 1; */margin-right:%?18?%;\n\t/* width: 112rpx; */\n\t/* height: 112rpx; */vertical-align:middle}.inner-img[data-v-a16dac8e]{height:%?78?%;width:%?78?%;border-radius:50%}\n\n/* 中部文字盒子样式 */.middle-text-box[data-v-a16dac8e]{-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:5;-webkit-flex:5;flex:5;display:-webkit-box;display:-webkit-flex;display:flex}\n\n/* 灰色字体 */.greay-text[data-v-a16dac8e]{color:#969696}\n\n/* 昵称 */.inner-text-name[data-v-a16dac8e]{font-size:%?35?%;font-weight:700;margin-right:%?16?%}\n\n/* 关注数粉丝数文章数 */.inner-text-message[data-v-a16dac8e]{padding-top:%?4?%;font-size:%?30?%}.inner-text[data-v-a16dac8e]{font-size:%?35?%}[data-v-a16dac8e] .article-img uni-image{width:%?120?%;height:%?120?%}[data-v-a16dac8e] .article-text{font-size:%?32?%;font-weight:700;-webkit-line-clamp:2;line-height:180%}[data-v-a16dac8e] .article-content{background-color:#f1f1f1;border-radius:4px;padding:10px 12px}.comment-content[data-v-a16dac8e]{margin-left:%?96?%;margin-top:%?10?%}.text-comment[data-v-a16dac8e]{font-size:%?30?%;color:#505050;width:%?580?%;white-space:nowrap;text-overflow:ellipsis;overflow:hidden}.content[data-v-a16dac8e]{padding-left:%?10?%;margin-top:%?80?%}",""]),t.exports=e}}]);