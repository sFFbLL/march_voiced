(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-home-index"],{"04ff":function(t,e,a){var n=a("4675");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var o=a("4f06").default;o("674bf75a",n,!0,{sourceMap:!1,shadowMode:!1})},"0e07":function(t,e,a){"use strict";(function(t){var n=a("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.check=o;n(a("5642"));function o(){t("log","有未读消息"," at utils/checkUnRead.js:6"),uni.showTabBarRedDot({index:3})}}).call(this,a("0de9")["log"])},"1fc3":function(t,e,a){"use strict";a.d(e,"b",(function(){return o})),a.d(e,"c",(function(){return i})),a.d(e,"a",(function(){return n}));var n={uBadge:a("cb97").default},o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{staticClass:"tab-card"},[a("v-uni-view",{staticClass:"head-nav"},t._l(t.tabs,(function(e,n){return a("v-uni-view",{key:e.index,staticClass:"head-nav-item",class:e.isActive?"active":"",on:{click:function(a){arguments[0]=a=t.$handleEvent(a),t.handlerActive(e.index)}}},[a("v-uni-view",{staticClass:"center"},[t._v(t._s(e.value)),0==e.index&&0==t.interactRead&&t.type?a("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),1==e.index&&0==t.attentionRead&&t.type?a("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),2==e.index&&0==t.otherRead&&t.type?a("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),a("v-uni-view",{staticClass:"head-nav-bottom"})],1)],1)})),1),t._t("default")],2)},i=[]},3801:function(t,e,a){var n=a("24fb");e=n(!1),e.push([t.i,".home[data-v-79698aa2],\n.home .content[data-v-79698aa2]{background-color:#f7f7f7;margin-bottom:%?30?%}.header[data-v-79698aa2]{position:fixed;z-index:10;width:100%;height:%?100?%;line-height:%?100?%;background-color:#fff;margin-bottom:%?15?%}.tag-nav[data-v-79698aa2]{display:inline-block;height:100%;width:90%}.search[data-v-79698aa2]{display:inline-block;width:8%}.search .search-icon[data-v-79698aa2]{position:absolute;top:%?5?%}.content[data-v-79698aa2]{background-color:#fff;padding-top:%?115?%}.loading[data-v-79698aa2]{background-color:#fff}[data-v-79698aa2] .uni-load-more .uni-load-more__text{font-size:%?28?%}[data-v-79698aa2] .uni-load-more .uni-load-more__img{height:%?30?%!important;width:%?30?%!important}",""]),t.exports=e},"463f":function(t,e,a){"use strict";var n;a.d(e,"b",(function(){return o})),a.d(e,"c",(function(){return i})),a.d(e,"a",(function(){return n}));var o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return t.show?a("v-uni-view",{staticClass:"u-badge",class:[t.isDot?"u-badge-dot":"","mini"==t.size?"u-badge-mini":"",t.type?"u-badge--bg--"+t.type:""],style:[{top:t.offset[0]+"rpx",right:t.offset[1]+"rpx",fontSize:t.fontSize+"rpx",position:t.absolute?"absolute":"static",color:t.color,backgroundColor:t.bgColor},t.boxStyle]},[t._v(t._s(t.showText))]):t._e()},i=[]},4675:function(t,e,a){var n=a("24fb");e=n(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-badge[data-v-34077204]{display:-webkit-inline-box;display:-webkit-inline-flex;display:inline-flex;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;line-height:%?24?%;padding:%?4?% %?8?%;border-radius:%?100?%;z-index:9}.u-badge--bg--primary[data-v-34077204]{background-color:#2979ff}.u-badge--bg--error[data-v-34077204]{background-color:#fa3534}.u-badge--bg--success[data-v-34077204]{background-color:#19be6b}.u-badge--bg--info[data-v-34077204]{background-color:#909399}.u-badge--bg--warning[data-v-34077204]{background-color:#f90}.u-badge-dot[data-v-34077204]{height:%?16?%;width:%?16?%;border-radius:%?100?%;line-height:1}.u-badge-mini[data-v-34077204]{-webkit-transform:scale(.8);transform:scale(.8);-webkit-transform-origin:center center;transform-origin:center center}.u-info[data-v-34077204]{background-color:#909399;color:#fff}',""]),t.exports=e},"4dff":function(t,e,a){"use strict";a.r(e);var n=a("fd1c"),o=a("5606");for(var i in o)"default"!==i&&function(t){a.d(e,t,(function(){return o[t]}))}(i);a("61f6");var r,s=a("f0c5"),c=Object(s["a"])(o["default"],n["b"],n["c"],!1,null,"79698aa2",null,!1,n["a"],r);e["default"]=c.exports},"50fc":function(t,e,a){var n=a("24fb");e=n(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.head-nav[data-v-06dc7823]{display:-webkit-flex;display:-webkit-box;display:flex;overflow:hidden;height:100%}.head-nav .head-nav-item[data-v-06dc7823]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center;font-size:%?34?%;font-family:Microsoft YaHei;font-weight:700;letter-spacing:%?4?%;color:#999}.head-nav-item .center[data-v-06dc7823]{position:relative;margin:auto;width:%?228?%}.active .center[data-v-06dc7823]{color:#404040}.active .center .head-nav-bottom[data-v-06dc7823]{position:relative;margin:auto;width:50%;height:%?8?%;top:%?-8?%;background-color:#404040}',""]),t.exports=e},5606:function(t,e,a){"use strict";a.r(e);var n=a("fd0c"),o=a.n(n);for(var i in n)"default"!==i&&function(t){a.d(e,t,(function(){return n[t]}))}(i);e["default"]=o.a},5642:function(t,e,a){"use strict";var n=a("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.interactList=r,e.attentionList=s,e.otherList=c,e.unreadMessage=d,e.readMessage=u;var o=n(a("367e")),i=n(a("e8a0"));function r(t){return o.default.get(i.default+"/api/message/user",{params:t})}function s(t){return o.default.get(i.default+"/api/message/follow-notice",{params:t})}function c(t){return o.default.get(i.default+"/api/message/system",{params:t})}function d(t){return o.default.get(i.default+"/api/message/user/unread-count",{params:t})}function u(t){return o.default.post(i.default+"/api/message/user",{type:t.type})}},"61f6":function(t,e,a){"use strict";var n=a("a6a0"),o=a.n(n);o.a},"8a32":function(t,e,a){"use strict";a.r(e);var n=a("8c76"),o=a.n(n);for(var i in n)"default"!==i&&function(t){a.d(e,t,(function(){return n[t]}))}(i);e["default"]=o.a},"8bc3":function(t,e,a){"use strict";a.r(e);var n=a("923a"),o=a.n(n);for(var i in n)"default"!==i&&function(t){a.d(e,t,(function(){return n[t]}))}(i);e["default"]=o.a},"8c76":function(t,e,a){"use strict";a("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={name:"u-badge",props:{type:{type:String,default:"error"},size:{type:String,default:"default"},isDot:{type:Boolean,default:!1},count:{type:[Number,String]},overflowCount:{type:Number,default:99},showZero:{type:Boolean,default:!1},offset:{type:Array,default:function(){return[20,20]}},absolute:{type:Boolean,default:!0},fontSize:{type:[String,Number],default:"24"},color:{type:String,default:"#ffffff"},bgColor:{type:String,default:""},isCenter:{type:Boolean,default:!1}},computed:{boxStyle:function(){var t={};return this.isCenter?(t.top=0,t.right=0,t.transform="translateY(-50%) translateX(50%)"):(t.top=this.offset[0]+"rpx",t.right=this.offset[1]+"rpx",t.transform="translateY(0) translateX(0)"),"mini"==this.size&&(t.transform=t.transform+" scale(0.8)"),t},showText:function(){return this.isDot?"":this.count>this.overflowCount?"".concat(this.overflowCount,"+"):this.count},show:function(){return 0!=this.count||0!=this.showZero}}};e.default=n},"923a":function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var n={data:function(){return{}},props:{tabs:{type:Array,default:[]},type:{type:Boolean}},methods:{handlerActive:function(t){0==t?this.$store.commit("changeInteract",1):1==t?this.$store.commit("changeAttention",1):2==t&&this.$store.commit("changeOther",1),this.$emit("tabActive",t)}},computed:{interactRead:function(){return this.$store.state.interactRead},attentionRead:function(){return this.$store.state.attentionRead},otherRead:function(){return this.$store.state.otherRead}},mounted:function(){},created:function(){}};e.default=n},a6a0:function(t,e,a){var n=a("3801");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var o=a("4f06").default;o("3bfa08e4",n,!0,{sourceMap:!1,shadowMode:!1})},b37c:function(t,e,a){var n=a("50fc");"string"===typeof n&&(n=[[t.i,n,""]]),n.locals&&(t.exports=n.locals);var o=a("4f06").default;o("0ffdf5f4",n,!0,{sourceMap:!1,shadowMode:!1})},cb97:function(t,e,a){"use strict";a.r(e);var n=a("463f"),o=a("8a32");for(var i in o)"default"!==i&&function(t){a.d(e,t,(function(){return o[t]}))}(i);a("d67e");var r,s=a("f0c5"),c=Object(s["a"])(o["default"],n["b"],n["c"],!1,null,"34077204",null,!1,n["a"],r);e["default"]=c.exports},d040:function(t,e,a){"use strict";var n=a("b37c"),o=a.n(n);o.a},d67e:function(t,e,a){"use strict";var n=a("04ff"),o=a.n(n);o.a},d6e3:function(t,e,a){"use strict";a.r(e);var n=a("1fc3"),o=a("8bc3");for(var i in o)"default"!==i&&function(t){a.d(e,t,(function(){return o[t]}))}(i);a("d040");var r,s=a("f0c5"),c=Object(s["a"])(o["default"],n["b"],n["c"],!1,null,"06dc7823",null,!1,n["a"],r);e["default"]=c.exports},fd0c:function(t,e,a){"use strict";(function(t){var n=a("4ea4");a("99af"),a("d81d"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var o=n(a("2909")),i=n(a("34f5")),r=n(a("d6e3")),s=n(a("7b34")),c=n(a("901f")),d=n(a("d508")),u=a("0e07"),l=a("e220"),f=a("dccd"),h=a("c98b"),v=a("5c16"),m={components:{uniIcons:i.default,tabs:r.default,recommend:s.default,follow:c.default,uniLoadMore:d.default},data:function(){return{recommendCurrent:1,articleCurrent:1,followCurrent:1,size:4,tabIndex:"",tablist:[{index:0,value:"推荐",isActive:!0},{index:1,value:"关注",isActive:!1}],recommendList:[],followList:[],loadStatus:"loading",recommendLoadStatus:"loading",followLoadStatus:"loading",isLoadMore:!1,isRecommend:!0}},beforeCreate:function(){if(t("log","homebeforeCreate"," at pages/home/index.vue:106"),!(0,h.getToken)()){var e;t("log","没有token"," at pages/home/index.vue:108"),(0,f.parseCode)()?e=(0,f.parseCode)():(t("log","当前没有wxCode，去获取"," at pages/home/index.vue:111"),e=(0,f.returnWxcode)()),t("log","成功拿到code"," at pages/home/index.vue:118");var a={code:e,status:1};(0,v.login)(a).then((function(e){t("log",e,"注册"," at pages/home/index.vue:125"),1==e.data.status?(t("log","未登录"," at pages/home/index.vue:128"),uni.navigateTo({url:"../login/login"})):(t("log",e.data.token," at pages/home/index.vue:134"),(0,h.setToken)(e.data.token),setOpenId(e.data.openid))})).catch((function(e){t("log",e,"err login"," at pages/home/index.vue:139")}))}},onShow:function(){(0,u.check)()},onLoad:function(){this.recommend(),this.follow()},onReachBottom:function(){this.isLoadMore||this.tabIndex?!this.isLoadMore&&this.tabIndex&&(this.isLoadMore=!0,this.followCurrent+=1,this.follow()):(this.isLoadMore=!0,this.isRecommend?(this.recommendCurrent+=1,this.recommend()):(this.articleCurrent+=1,this.getArticleList()))},methods:{tabActive:function(t){this.tablist.map((function(e,a){e.isActive=t==a})),t?t&&(this.loadStatus=this.followLoadStatus,this.isLoadMore=!1):(this.loadStatus=this.recommendLoadStatus,this.isLoadMore=!1),this.tabIndex=t},recommend:function(){var t=this,e=this,a={current:this.recommendCurrent,size:this.size};(0,l.getRecommend)(a).then((function(a){e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(a.data)),a.data.length<e.size?(e.loadStatus="nomore",e.isRecommend=!1,e.getArticleList()):1===t.recommendCurrent?(e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(a.data))):setTimeout((function(){e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(a.data))}),3e3)}))},getArticleList:function(){var t=this,e=this,a={current:this.articleCurrent,size:this.size};(0,l.getArticleList)(a).then((function(a){e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(a.data)),a.data.length<e.size?(e.loadStatus="nomore",e.recommendLoadStatus="nomore"):1===t.articleCurrent?(e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(a.data))):setTimeout((function(){e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(a.data))}),3e3)}))},follow:function(){var t=this,e=this,a={current:this.followCurrent,size:this.size};(0,l.getFollow)(a).then((function(a){e.followList=[].concat((0,o.default)(e.followList),(0,o.default)(a.data.records)),a.data.records.length<e.size?(e.loadStatus="nomore",e.follLoadStatus="nomore"):1===t.followCurrent?(e.isLoadMore=!1,e.followList=[].concat((0,o.default)(e.followList),(0,o.default)(a.data.records))):setTimeout((function(){e.isLoadMore=!1,e.followList=[].concat((0,o.default)(e.followList),(0,o.default)(a.data.records))}),3e3)}))},search:function(){uni.navigateTo({url:"../search/index"})}}};e.default=m}).call(this,a("0de9")["log"])},fd1c:function(t,e,a){"use strict";a.d(e,"b",(function(){return o})),a.d(e,"c",(function(){return i})),a.d(e,"a",(function(){return n}));var n={uniIcons:a("34f5").default,uniLoadMore:a("d508").default},o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{staticClass:"home"},[a("v-uni-view",{staticClass:"header"},[a("tabs",{staticClass:"tag-nav",attrs:{tabs:t.tablist},on:{tabActive:function(e){arguments[0]=e=t.$handleEvent(e),t.tabActive.apply(void 0,arguments)}}}),a("v-uni-view",{staticClass:"search",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.search()}}},[a("v-uni-view",{staticClass:"search-icon"},[a("uni-icons",{attrs:{type:"search",size:"25",color:"#999"}})],1)],1)],1),a("v-uni-view",{staticClass:"content"},[t.tabIndex?t._e():a("v-uni-view",t._l(t.recommendList,(function(t,e){return a("v-uni-view",[a("recommend",{attrs:{articleInfo:t}})],1)})),1),t.tabIndex?a("v-uni-view",t._l(t.followList,(function(t,e){return a("v-uni-view",{key:e},[a("follow",{attrs:{articleInfo:t}})],1)})),1):t._e(),a("v-uni-view",{directives:[{name:"show",rawName:"v-show",value:t.isLoadMore,expression:"isLoadMore"}]},[a("uni-load-more",{staticClass:"loading",attrs:{status:t.loadStatus,iconType:"circle"}})],1)],1)],1)},i=[]}}]);