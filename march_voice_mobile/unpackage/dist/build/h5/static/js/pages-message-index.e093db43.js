(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-message-index"],{"029d":function(t,e,n){"use strict";n.r(e);var i=n("c48b"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},"04ff":function(t,e,n){var i=n("4675");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("674bf75a",i,!0,{sourceMap:!1,shadowMode:!1})},"06b5":function(t,e,n){"use strict";n.r(e);var i=n("07c2"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},"07c2":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={name:"u-empty",props:{src:{type:String,default:""},text:{type:String,default:""},color:{type:String,default:"#c0c4cc"},iconColor:{type:String,default:"#c0c4cc"},iconSize:{type:[String,Number],default:120},fontSize:{type:[String,Number],default:26},mode:{type:String,default:"data"},imgWidth:{type:[String,Number],default:120},imgHeight:{type:[String,Number],default:"auto"},show:{type:Boolean,default:!0},marginTop:{type:[String,Number],default:0},iconStyle:{type:Object,default:function(){return{}}}},data:function(){return{icons:{car:"购物车为空",page:"页面不存在",search:"没有搜索结果",address:"没有收货地址",wifi:"没有WiFi",order:"订单为空",coupon:"没有优惠券",favor:"暂无收藏",permission:"无权限",history:"无历史记录",news:"无新闻列表",message:"消息列表为空",list:"列表为空",data:"数据为空"}}}};e.default=i},"17f2":function(t,e,n){"use strict";var i=n("88dd"),a=n.n(i);a.a},"1b01":function(t,e,n){"use strict";var i=n("2c46"),a=n.n(i);a.a},"1fc3":function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var i={uBadge:n("cb97").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"tab-card"},[n("v-uni-view",{staticClass:"head-nav"},t._l(t.tabs,(function(e,i){return n("v-uni-view",{key:e.index,staticClass:"head-nav-item",class:e.isActive?"active":"",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.handlerActive(e.index)}}},[n("v-uni-view",{staticClass:"center"},[t._v(t._s(e.value)),0==e.index&&0==t.interactRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),1==e.index&&0==t.attentionRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),2==e.index&&0==t.otherRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),n("v-uni-view",{staticClass:"head-nav-bottom"})],1)],1)})),1),t._t("default")],2)},r=[]},"2c46":function(t,e,n){var i=n("b7b0");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("0e0e6fda",i,!0,{sourceMap:!1,shadowMode:!1})},"349b":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={data:function(){return{}},props:{articleId:{type:Number,default:null},articleStatus:{type:Number,default:null},articleTitle:{type:String,default:""},commentTotal:{type:Number,default:0}},methods:{articleDetails:function(){uni.navigateTo({url:"../articleDetails/index?id="+this.articleId+"&commentTotal="+this.commentTotal})}},mounted:function(){}};e.default=i},3761:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".wrap[data-v-a0d0d1bc]{padding:%?50?% %?10?%}.time[data-v-a0d0d1bc]{margin-top:%?10?%;font-size:%?20?%;color:#616367;margin-left:%?10?%}",""]),t.exports=e},3775:function(t,e,n){var i=n("5564");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("62782ad1",i,!0,{sourceMap:!1,shadowMode:!1})},"3db2":function(t,e,n){"use strict";n.r(e);var i=n("9472"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},4023:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("d50f")),r=i(n("1e61")),o=i(n("40e6")),s={props:{messageInfo:{type:Object,default:function(){return{type:0,status:1,title:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件今年春天在写作圈发生了几件不大不小的抄袭洗稿事件",articleId:1,comment:"哈哈哈哈在哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠哈哈哈哈在法国夫是德国人头地方杠法国夫是德国人头地方杠",content:"今年春天在写作圈发生了几件不大不小的抄袭洗稿事件。一件是言情大神匪",image:n("f187"),createTime:"2020/12/12",userId:"1",avatarPath:n("b9dd"),nickname:"张三"}}}},components:{attentionAndFansCell:o.default,articleTitle:a.default},data:function(){return{tip:"赞了你的文章"}},methods:{format:function(t){var e=new Date(t),n=(0,r.default)(e).format("YYYY-MM-DD HH:mm:ss");return n}},created:function(){0==this.messageInfo.type?1==this.messageInfo.status?this.tip="评论了你的文章":2==this.messageInfo.status?this.tip="收藏了你的文章":3==this.messageInfo.status?this.tip="转发了你的文章":4==this.messageInfo.status&&(this.tip="赞了你的文章"):1==this.messageInfo.type&&(1==this.messageInfo.status?this.tip="评论了你的想法":0==this.messageInfo.status||2!=this.messageInfo.status&&3!=this.messageInfo.status&&4!=this.messageInfo.status||(this.tip="送了你一个"))}};e.default=s},4213:function(t,e,n){"use strict";n.r(e);var i=n("8494"),a=n("cf2e");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("b26e");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"07f2de9a",null,!1,i["a"],o);e["default"]=c.exports},"459f":function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-empty[data-v-21937aaa]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\nflex-direction:column;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;height:100%}.u-image[data-v-21937aaa]{margin-bottom:%?20?%}.u-slot-wrap[data-v-21937aaa]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\n-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;margin-top:%?20?%}',""]),t.exports=e},"463f":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return t.show?n("v-uni-view",{staticClass:"u-badge",class:[t.isDot?"u-badge-dot":"","mini"==t.size?"u-badge-mini":"",t.type?"u-badge--bg--"+t.type:""],style:[{top:t.offset[0]+"rpx",right:t.offset[1]+"rpx",fontSize:t.fontSize+"rpx",position:t.absolute?"absolute":"static",color:t.color,backgroundColor:t.bgColor},t.boxStyle]},[t._v(t._s(t.showText))]):t._e()},r=[]},4675:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-badge[data-v-34077204]{display:-webkit-inline-box;display:-webkit-inline-flex;display:inline-flex;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;line-height:%?24?%;padding:%?4?% %?8?%;border-radius:%?100?%;z-index:9}.u-badge--bg--primary[data-v-34077204]{background-color:#2979ff}.u-badge--bg--error[data-v-34077204]{background-color:#fa3534}.u-badge--bg--success[data-v-34077204]{background-color:#19be6b}.u-badge--bg--info[data-v-34077204]{background-color:#909399}.u-badge--bg--warning[data-v-34077204]{background-color:#f90}.u-badge-dot[data-v-34077204]{height:%?16?%;width:%?16?%;border-radius:%?100?%;line-height:1}.u-badge-mini[data-v-34077204]{-webkit-transform:scale(.8);transform:scale(.8);-webkit-transform-origin:center center;transform-origin:center center}.u-info[data-v-34077204]{background-color:#909399;color:#fff}',""]),t.exports=e},4907:function(t,e,n){t.exports=n.p+"static/img/xiaoku.e5dd8801.png"},"50fc":function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.head-nav[data-v-06dc7823]{display:-webkit-flex;display:-webkit-box;display:flex;overflow:hidden;height:100%}.head-nav .head-nav-item[data-v-06dc7823]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center;font-size:%?34?%;font-family:Microsoft YaHei;font-weight:700;letter-spacing:%?4?%;color:#999}.head-nav-item .center[data-v-06dc7823]{position:relative;margin:auto;width:%?228?%}.active .center[data-v-06dc7823]{color:#404040}.active .center .head-nav-bottom[data-v-06dc7823]{position:relative;margin:auto;width:50%;height:%?8?%;top:%?-8?%;background-color:#404040}',""]),t.exports=e},"53ac":function(t,e,n){"use strict";n.r(e);var i=n("7e61"),a=n("7f39");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("c788");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"71495536",null,!1,i["a"],o);e["default"]=c.exports},5564:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".header[data-v-07f2de9a]{position:fixed;top:0;background-color:#fff;z-index:99;width:100%;height:%?80?%;line-height:%?80?%}\n\n/* 公共头像组件样式 */[data-v-07f2de9a] .attention-cell .flex-item{padding:%?10?%;margin-top:%?17?%}\n\n/* 关注插槽 */.slot[data-v-07f2de9a]{margin-left:%?20?%;font-weight:400}[data-v-07f2de9a] .tab-card .head-nav .head-nav-bottom{width:70%}[data-v-07f2de9a] .head-nav .head-nav-item .center{top:%?10?%}.flex-item[data-v-07f2de9a]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-flex-wrap:wrap;flex-wrap:wrap;box-sizing:border-box;border-bottom:%?6?% solid #f0f0f0;padding:%?20?% %?30?% %?20?% %?30?%\n\t/* margin-bottom: 20rpx; */}\n\n/* 左侧图片盒子样式 */.left-img-box[data-v-07f2de9a]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;\n\t/* flex: 1; */margin-right:%?18?%;\n\t/* width: 112rpx; */\n\t/* height: 112rpx; */vertical-align:middle}.inner-img[data-v-07f2de9a]{height:%?78?%;width:%?78?%;border-radius:50%}\n\n/* 中部文字盒子样式 */.middle-text-box[data-v-07f2de9a]{-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:5;-webkit-flex:5;flex:5;display:-webkit-box;display:-webkit-flex;display:flex}\n\n/* 灰色字体 */.greay-text[data-v-07f2de9a]{color:#969696}\n\n/* 昵称 */.inner-text-name[data-v-07f2de9a]{font-size:%?35?%;font-weight:700;margin-right:%?16?%}\n\n/* 关注数粉丝数文章数 */.inner-text-message[data-v-07f2de9a]{padding-top:%?4?%;font-size:%?30?%}.inner-text[data-v-07f2de9a]{font-size:%?35?%}[data-v-07f2de9a] .article-img uni-image{width:%?120?%;height:%?120?%}[data-v-07f2de9a] .article-text{font-size:%?32?%;font-weight:700;-webkit-line-clamp:2;line-height:180%}[data-v-07f2de9a] .article-content{background-color:#f1f1f1;border-radius:4px;padding:10px 12px}.comment-content[data-v-07f2de9a]{margin-left:%?96?%;margin-top:%?10?%}.text-comment[data-v-07f2de9a]{font-size:%?30?%;color:#505050;width:%?580?%;white-space:nowrap;text-overflow:ellipsis;overflow:hidden}.content[data-v-07f2de9a]{padding-left:%?10?%;margin-top:%?90?%}.nodate[data-v-07f2de9a]{background-color:#fff;min-height:%?800?%}",""]),t.exports=e},5742:function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={name:"u-gap",props:{bgColor:{type:String,default:"transparent "},height:{type:[String,Number],default:30},marginTop:{type:[String,Number],default:0},marginBottom:{type:[String,Number],default:0}},computed:{gapStyle:function(){return{backgroundColor:this.bgColor,height:this.height+"rpx",marginTop:this.marginTop+"rpx",marginBottom:this.marginBottom+"rpx"}}}};e.default=i},"5bfe":function(t,e,n){var i=n("3761");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("414e5b35",i,!0,{sourceMap:!1,shadowMode:!1})},"61ee":function(t,e,n){t.exports=n.p+"static/img/aixin.39c2cdd4.png"},"6d64":function(t,e,n){t.exports=n.p+"static/img/success.74305a2d.png"},"7cb0":function(t,e,n){"use strict";n.r(e);var i=n("d058"),a=n("029d");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("ff1b");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"41f451a6",null,!1,i["a"],o);e["default"]=c.exports},"7e61":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"u-gap",style:[t.gapStyle]})},r=[]},"7f39":function(t,e,n){"use strict";n.r(e);var i=n("5742"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},8494:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var i={uEmpty:n("f8b9").default,uGap:n("53ac").default,uniLoadMore:n("d508").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"message"},[n("v-uni-view",{staticClass:"header"},[n("tabs",{staticClass:"tag-nav",attrs:{type:t.message,tabs:t.tablist},on:{tabActive:function(e){arguments[0]=e=t.$handleEvent(e),t.tabActive.apply(void 0,arguments)}}})],1),n("v-uni-view",{staticClass:"content"},[0==t.tabIndex?n("v-uni-view",[t._l(t.interactListCopy,(function(t){return n("v-uni-view",[n("singleMessage",{attrs:{messageInfo:t}})],1)})),t.interactListCopy.length<1?n("v-uni-view",[n("u-empty",{staticClass:"nodate",attrs:{text:"没有数据",mode:"search"}})],1):t._e()],2):t._e(),1==t.tabIndex?n("v-uni-view",[t._l(t.attentionListCopy,(function(e){return n("v-uni-view",[n("attentionAndFansCell",{staticClass:"attention",attrs:{aid:e.userId,nickname:e.nickname,avatarPath:e.avatarPath,isFollow:e.isFollow},scopedSlots:t._u([{key:"afterNicknameText",fn:function(){return[n("v-uni-text",{staticClass:"slot"},[t._v("关注了你")])]},proxy:!0},{key:"underText",fn:function(){return[t._v(t._s(t.format(e.createTime)))]},proxy:!0}],null,!0)})],1)})),t.attentionListCopy.length<1?n("v-uni-view",[n("u-empty",{staticClass:"nodate",attrs:{text:"没有数据",mode:"search"}})],1):t._e()],2):t._e(),2==t.tabIndex?n("v-uni-view",[t._l(t.otherListCopy,(function(e){return n("v-uni-view",[n("otherMessage",{attrs:{otherList:t.otherListCopy}}),n("u-gap",{attrs:{height:"1","bg-color":"#f5f5f5"}})],1)})),t.otherListCopy.length<1?n("v-uni-view",[n("u-empty",{staticClass:"nodate",attrs:{text:"没有数据",mode:"search"}})],1):t._e()],2):t._e()],1),n("v-uni-view",{directives:[{name:"show",rawName:"v-show",value:t.isLoadMore,expression:"isLoadMore"}]},[n("uni-load-more",{attrs:{status:t.loadStatus}})],1)],1)},r=[]},"84a8":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",t._l(t.otherList,(function(e,a){return i("v-uni-view",{staticClass:"wrap"},[0==e.status?i("v-uni-image",{staticStyle:{width:"30rpx",height:"30rpx",margin:"0 10rpx"},attrs:{src:n("6d64"),mode:"scaleToFill"}}):t._e(),1==e.status?i("v-uni-image",{staticStyle:{width:"30rpx",height:"30rpx",margin:"0 10rpx"},attrs:{src:n("f982"),mode:"scaleToFill"}}):t._e(),0==e.type?i("v-uni-text",[t._v("您的文章")]):t._e(),0==e.type?i("v-uni-navigator",{staticStyle:{color:"#007AFF",display:"inline-block"},attrs:{url:t.url}},[t._v("《"+t._s(t.otherList.title)+"》")]):t._e(),i("v-uni-text",{staticClass:"tipWords"},[t._v(t._s(t.tipWords))]),i("v-uni-view",{staticClass:"time"},[i("v-uni-text",[t._v(t._s(t.format(e.createTime)))])],1)],1)})),1)},r=[]},"88dd":function(t,e,n){var i=n("459f");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("ecf364d6",i,!0,{sourceMap:!1,shadowMode:!1})},"8a32":function(t,e,n){"use strict";n.r(e);var i=n("8c76"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},"8bc3":function(t,e,n){"use strict";n.r(e);var i=n("923a"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},"8c76":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={name:"u-badge",props:{type:{type:String,default:"error"},size:{type:String,default:"default"},isDot:{type:Boolean,default:!1},count:{type:[Number,String]},overflowCount:{type:Number,default:99},showZero:{type:Boolean,default:!1},offset:{type:Array,default:function(){return[20,20]}},absolute:{type:Boolean,default:!0},fontSize:{type:[String,Number],default:"24"},color:{type:String,default:"#ffffff"},bgColor:{type:String,default:""},isCenter:{type:Boolean,default:!1}},computed:{boxStyle:function(){var t={};return this.isCenter?(t.top=0,t.right=0,t.transform="translateY(-50%) translateX(50%)"):(t.top=this.offset[0]+"rpx",t.right=this.offset[1]+"rpx",t.transform="translateY(0) translateX(0)"),"mini"==this.size&&(t.transform=t.transform+" scale(0.8)"),t},showText:function(){return this.isDot?"":this.count>this.overflowCount?"".concat(this.overflowCount,"+"):this.count},show:function(){return 0!=this.count||0!=this.showZero}}};e.default=i},"8cbe":function(t,e,n){"use strict";var i=n("4ea4");n("99af"),n("d81d"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("2909")),r=i(n("d6e3")),o=i(n("7cb0")),s=i(n("99f6")),c=i(n("9752")),u=i(n("40e6")),d=i(n("1e61")),l=n("5642"),f={components:{tabs:r.default,articleContent:o.default,singleMessage:s.default,attentionAndFansCell:u.default,otherMessage:c.default},data:function(){return{interactCurrent:1,attentionCurrent:1,otherCurrent:1,size:10,isLoadMore:!1,loadStatus:"loading",interactLoadStatus:"loading",attentionLoadStatus:"loading",otherLoadStatus:"loading",isLoading:!1,isDisabled:!1,isAttention:!1,tabIndex:0,message:!0,attentionRead:!1,otherRead:!1,interactRead:!1,attentionClick:!1,otherionClick:!1,interactClick:!1,tablist:[{index:0,value:"互动消息",isActive:!0},{index:1,value:"关注消息",isActive:!1},{index:2,value:"系统消息",isActive:!1}],interactList:{pages:"22",total:"100",current:"1",size:"1",records:[]},interactListCopy:[],attentionList:{pages:"22",total:"100",current:"1",size:"1",records:[]},attentionListCopy:[],otherList:{pages:"22",total:"100",current:"1",size:"1",records:[]},otherListCopy:[],isNextPage:!0}},onReachBottom:function(){this.isLoadMore||0!=this.tabIndex?this.isLoadMore||1!=this.tabIndex?this.isLoadMore||2!=this.tabIndex||(this.isLoadMore=!0,this.isNextPage&&(this.otherCurrent+=1),this.other()):(this.isLoadMore=!0,this.isNextPage&&(this.attentionCurrent+=1),this.attention()):(this.isLoadMore=!0,this.isNextPage&&(this.interactCurrent+=1),this.interact())},onPullDownRefresh:function(){var t=this;setTimeout((function(){uni.stopPullDownRefresh()}),2e3),0==this.tabIndex?(this.interactCurrent=1,this.interactList=[],this.interact(),this.$store.commit("changeInteract",1)):1==this.tabIndex?(this.attentionCurrent=1,this.attentionList=[],this.attention(),this.$store.commit("changeAttention",1)):2==this.tabIndex&&(this.otherCurrent=1,this.otherList=[],this.other(),this.$store.commit("changeOther",1)),(0,l.unreadMessage)(1).then((function(e){e.data.count>0&&t.$store.commit("changeInteract",0)})),(0,l.unreadMessage)(2).then((function(e){e.data.count>0&&t.$store.commit("changeAttention",0)})),(0,l.unreadMessage)(3).then((function(e){e.data.count>0&&t.$store.commit("changeOther",0)}))},onShow:function(){uni.hideTabBarRedDot({index:3})},created:function(){var t=this;this.$store.commit("changeInteract",1),(0,l.readMessage)(1).then((function(t){})),this.interact(),(0,l.unreadMessage)(2).then((function(e){e.data.count>0&&t.$store.commit("changeAttention",0)})),(0,l.unreadMessage)(3).then((function(e){e.data.count>0&&t.$store.commit("changeOther",0)}))},methods:{tabActive:function(t){var e=this;if(this.tablist.map((function(e,n){e.isActive=t==n})),0==t){if(this.loadStatus=this.interactLoadStatus,1===this.interactCurrent)while(!this.interactClick)this.interactClick=!0,this.interactList.records=[],this.interact(),setTimeout((function(){e.interactClick=!1}),1e3)}else if(1==t){if(this.loadStatus=this.attentionLoadStatus,1===this.attentionCurrent)while(!this.attentionClick)this.attentionList.records=[],this.attentionClick=!0,this.attention(),setTimeout((function(){e.attentionClick=!1}),1e3)}else if(this.loadStatus=this.otherLoadStatus,1===this.otherCurrent)while(!this.otherClick)this.otherClick=!0,this.otherList.records=[],this.other(),setTimeout((function(){e.otherClick=!1}),1e3);this.tabIndex=t},interact:function(){var t=this,e={current:this.interactCurrent,size:this.size};(0,l.interactList)(e).then((function(e){0!=e.data.records?(t.interactCurrent,t.isLoadMore=!1,t.interactList.records=[].concat((0,a.default)(t.interactList.records),(0,a.default)(e.data.records)),t.interactListCopy=t.interactList.records,t.isNextPage=!0):(t.isLoadMore=!1,t.loadStatus="nomore",t.isNextPage=!1)})),(0,l.readMessage)(1).then((function(t){}))},attention:function(){var t=this,e={current:this.attentionCurrent,size:this.size};(0,l.attentionList)(e).then((function(e){e.data.records?(t.attentionCurrent,t.isLoadMore=!1,t.attentionList.records=[].concat((0,a.default)(t.attentionList.records),(0,a.default)(e.data.records)),t.attentionListCopy=t.attentionList.records,t.isNextPage=!0):(t.loadStatus="nomore",t.isLoadMore=!1,t.isNextPage=!1)})),(0,l.readMessage)(2).then((function(t){}))},other:function(){var t=this,e={current:this.otherCurrent,size:this.size};(0,l.otherList)(e).then((function(e){e.data.records?(t.otherCurrent,t.isLoadMore=!1,t.otherList.records=[].concat((0,a.default)(t.otherList.records),(0,a.default)(e.data.records)),t.otherListCopy=t.otherList.records,t.isNextPage=!0):(t.loadStatus="nomore",t.isLoadMore=!1,t.isNextPage=!1)})),(0,l.readMessage)(3).then((function(t){}))},format:function(t){var e=new Date(t),n=(0,d.default)(e).format("YYYY-MM-DD HH:mm:ss");return n}}};e.default=f},"8dc0":function(t,e,n){var i=n("fbce");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("8df927d0",i,!0,{sourceMap:!1,shadowMode:!1})},"923a":function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={data:function(){return{}},props:{tabs:{type:Array,default:[]},type:{type:Boolean}},methods:{handlerActive:function(t){0==t?this.$store.commit("changeInteract",1):1==t?this.$store.commit("changeAttention",1):2==t&&this.$store.commit("changeOther",1),this.$emit("tabActive",t)}},computed:{interactRead:function(){return this.$store.state.interactRead},attentionRead:function(){return this.$store.state.attentionRead},otherRead:function(){return this.$store.state.otherRead}},mounted:function(){},created:function(){}};e.default=i},"936b":function(t,e,n){var i=n("ad20");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("7c85361e",i,!0,{sourceMap:!1,shadowMode:!1})},"93e5":function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n/* 评论样式 */.comment[data-v-98dd1344]{margin-left:%?110?%;font-size:%?28?%;color:#505050;margin-top:%?10?%;margin-bottom:%?10?%;width:%?600?%;overflow:hidden;-webkit-line-clamp:1;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical}\n/* 头部插槽动作 */.slot[data-v-98dd1344]{margin-left:%?15?%;font-size:%?28?%;color:#505050;font-weight:300}\n/* 文章和想法的样式 */.thecontent[data-v-98dd1344]{margin-left:%?110?%;width:%?600?%;\n\t/* height: 140rpx; */background-color:#f1f1f1;border-radius:%?8?%;display:-webkit-box;display:-webkit-flex;display:flex;padding:%?20?% %?20?%;margin-bottom:%?30?%;margin-top:%?6?%}.thewords[data-v-98dd1344]{width:%?1200?%}.theimg[data-v-98dd1344]{width:%?300?%;height:%?100?%;border-radius:%?4?%;margin-left:%?10?%}\n/* 公共头像组件样式 */[data-v-98dd1344] .attention-cell .flex-item{border:none;padding:%?10?%;margin-top:%?17?%}",""]),t.exports=e},9472:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("1e61")),r={props:{otherList:{type:Array}},data:function(){return{tipWords:"",url:"../../pages/articleDetails/index?id="}},created:function(){0==this.otherList.type?(this.url="../../pages/articleDetails/index?id="+this.otherList.articleId,0==this.otherList.status?this.tipWords="通过审核，已为您自动发布。":1==this.otherList.status&&(this.tipWords="含有敏感信息，请进行修正后重新发布。")):1==this.otherList.type&&(0==this.otherList.status?this.tipWords="恭喜你成功加入三月圈，快去看看吧。":1==this.otherList.status&&(this.tipWords="由于权限不足，未能成功加入三月圈。"))},methods:{format:function(t){var e=new Date(t),n=(0,a.default)(e).format("YYYY-MM-DD HH:mm:ss");return n}}};e.default=r},9752:function(t,e,n){"use strict";n.r(e);var i=n("84a8"),a=n("3db2");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("abc3");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"a0d0d1bc",null,!1,i["a"],o);e["default"]=c.exports},"99f6":function(t,e,n){"use strict";n.r(e);var i=n("a87d"),a=n("b9b9");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("b90e");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"98dd1344",null,!1,i["a"],o);e["default"]=c.exports},a87d:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var i={uGap:n("53ac").default},a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"wrap"},[i("attentionAndFansCell",{attrs:{aid:t.messageInfo.userId,nickname:t.messageInfo.nickname,avatarPath:t.messageInfo.avatarPath},scopedSlots:t._u([{key:"afterNicknameText",fn:function(){return[i("v-uni-text",{staticClass:"slot"},[t._v(t._s(t.tip))]),3==t.messageInfo.status&&1==t.messageInfo.type?i("v-uni-image",{staticStyle:{width:"44rpx",height:"30rpx","margin-left":"10rpx"},attrs:{src:n("61ee"),mode:""}}):t._e(),2==t.messageInfo.status&&1==t.messageInfo.type?i("v-uni-image",{staticStyle:{width:"34rpx",height:"30rpx","margin-left":"10rpx"},attrs:{src:n("4907"),mode:""}}):t._e(),4==t.messageInfo.status&&1==t.messageInfo.type?i("v-uni-image",{staticStyle:{width:"34rpx",height:"30rpx","margin-left":"10rpx"},attrs:{src:n("c7cd"),mode:""}}):t._e()]},proxy:!0},{key:"underText",fn:function(){return[i("v-uni-text",[t._v(t._s(t.format(t.messageInfo.createTime)))])]},proxy:!0}])}),t.messageInfo.comment?i("v-uni-text",{staticClass:"comment"},[t._v(t._s(t.messageInfo.comment))]):t._e(),i("v-uni-view",{staticClass:"thecontent"},[0==t.messageInfo.type?i("articleTitle",{staticClass:"thewords",attrs:{articleTitle:t.messageInfo.title,articleId:t.messageInfo.articleId}}):t._e(),1==t.messageInfo.type?i("articleTitle",{staticClass:"thewords",attrs:{articleTitle:t.messageInfo.content,articleId:t.messageInfo.articleId}}):t._e(),t.messageInfo.image?i("v-uni-image",{staticClass:"theimg",attrs:{src:t.messageInfo.image,mode:"scaleToFill"}}):t._e()],1),i("u-gap",{attrs:{height:"10","bg-color":"#f5f5f5"}})],1)},r=[]},abc3:function(t,e,n){"use strict";var i=n("5bfe"),a=n.n(i);a.a},ad20:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/',""]),t.exports=e},b26e:function(t,e,n){"use strict";var i=n("3775"),a=n.n(i);a.a},b37c:function(t,e,n){var i=n("50fc");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("0ffdf5f4",i,!0,{sourceMap:!1,shadowMode:!1})},b7b0:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".article-title .article-text[data-v-51c1e28d]{color:#404040;font-size:%?34?%;font-family:微软雅黑;text-align:left;font-weight:700;overflow:hidden;-webkit-line-clamp:2;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;line-height:150%}.article-text[data-v-51c1e28d]{display:inline-block;width:100%}",""]),t.exports=e},b90e:function(t,e,n){"use strict";var i=n("e0fd"),a=n.n(i);a.a},b9b9:function(t,e,n){"use strict";n.r(e);var i=n("4023"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},b9dd:function(t,e,n){t.exports=n.p+"static/img/1.e963bd74.jpg"},bfed:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var i={uIcon:n("ac18").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return t.show?n("v-uni-view",{staticClass:"u-empty",style:{marginTop:t.marginTop+"rpx"}},[n("u-icon",{attrs:{name:t.src?t.src:"empty-"+t.mode,"custom-style":t.iconStyle,label:t.text?t.text:t.icons[t.mode],"label-pos":"bottom","label-color":t.color,"label-size":t.fontSize,size:t.iconSize,color:t.iconColor,"margin-top":"14"}}),n("v-uni-view",{staticClass:"u-slot-wrap"},[t._t("bottom")],2)],1):t._e()},r=[]},c48b:function(t,e,n){"use strict";n("a9e3"),n("ac1f"),n("5319"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={data:function(){return{textContent:""}},props:{id:{type:Number,default:null},articleContent:{type:String,default:""},articleImg:{type:String,default:""},isIdea:{type:Boolean,default:!1},commentTotal:{type:Number,default:0},type:{type:String,default:null}},components:{},methods:{toDetails:function(){"draft"!==this.type?this.isIdea?uni.navigateTo({url:"../ideaDetails/index?id="+this.id+"&commentTotal="+this.commentTotal}):uni.navigateTo({url:"../articleDetails/index?id="+this.id+"&commentTotal="+this.commentTotal}):uni.navigateTo({url:"../../pages/publish/richTextEditor?type=draft&id="+this.id+"&commentTotal="+this.commentTotal})}},created:function(){this.textContent=this.articleContent},mounted:function(){this.textContent=this.$refs.articleContent.$el.innerText.replace(/\ +/g,"").replace(/[\r\n]/g,"")}};e.default=i},c788:function(t,e,n){"use strict";var i=n("936b"),a=n.n(i);a.a},c7cd:function(t,e,n){t.exports=n.p+"static/img/qingzhu.a25da244.png"},cb97:function(t,e,n){"use strict";n.r(e);var i=n("463f"),a=n("8a32");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("d67e");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"34077204",null,!1,i["a"],o);e["default"]=c.exports},cf2e:function(t,e,n){"use strict";n.r(e);var i=n("8cbe"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},d040:function(t,e,n){"use strict";var i=n("b37c"),a=n.n(i);a.a},d058:function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-content",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.toDetails.apply(void 0,arguments)}}},[n("v-uni-view",[n("v-uni-text",{ref:"articleContent",staticClass:"article-text",domProps:{innerHTML:t._s(t.textContent)}},[t._v(t._s(t.textContent))])],1),""!=t.articleImg?n("v-uni-view",{staticClass:"article-img"},[n("v-uni-image",{attrs:{src:t.articleImg,mode:"aspectFill"}})],1):t._e()],1)},r=[]},d50f:function(t,e,n){"use strict";n.r(e);var i=n("e8e2"),a=n("f708");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("1b01");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"51c1e28d",null,!1,i["a"],o);e["default"]=c.exports},d67e:function(t,e,n){"use strict";var i=n("04ff"),a=n.n(i);a.a},d6e3:function(t,e,n){"use strict";n.r(e);var i=n("1fc3"),a=n("8bc3");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("d040");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"06dc7823",null,!1,i["a"],o);e["default"]=c.exports},e0fd:function(t,e,n){var i=n("93e5");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("6ae71043",i,!0,{sourceMap:!1,shadowMode:!1})},e8e2:function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-title",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.articleDetails.apply(void 0,arguments)}}},[n("v-uni-view",{staticStyle:{display:"flex"}},[n("v-uni-text",{staticClass:"article-text",domProps:{innerHTML:t._s(t.articleTitle)}}),2==t.articleStatus?n("v-uni-text",{staticStyle:{width:"51px",color:"gray"}},[t._v("审核中")]):t._e()],1)],1)},r=[]},f187:function(t,e,n){t.exports=n.p+"static/img/2.c15f2abc.jpg"},f708:function(t,e,n){"use strict";n.r(e);var i=n("349b"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},f8b9:function(t,e,n){"use strict";n.r(e);var i=n("bfed"),a=n("06b5");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("17f2");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"21937aaa",null,!1,i["a"],o);e["default"]=c.exports},f982:function(t,e,n){t.exports=n.p+"static/img/fail.538da67e.png"},fbce:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".article-content[data-v-41f451a6]{margin-top:%?18?%;display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;font-size:%?32?%;font-family:PingFangSC;color:#404040;min-height:%?140?%}.article-text[data-v-41f451a6]{overflow:hidden;-webkit-line-clamp:3;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;line-height:150%}.article-img[data-v-41f451a6]{display:inline-block\n\t/* margin-right: 30rpx; */}.article-img uni-image[data-v-41f451a6]{width:%?180?%;height:%?140?%;background-color:#eee;border-radius:4px}",""]),t.exports=e},ff1b:function(t,e,n){"use strict";var i=n("8dc0"),a=n.n(i);a.a}}]);