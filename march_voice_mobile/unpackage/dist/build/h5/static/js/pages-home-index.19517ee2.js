(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-home-index"],{"0391":function(t,e,n){"use strict";(function(t){var a=n("4ea4");n("99af"),n("d81d"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var o=a(n("2909")),i=a(n("11a2")),r=a(n("a2be")),s=a(n("8b18")),d=a(n("58de")),c=a(n("19bf")),u=n("ac2f"),l=n("77ac"),f=n("c076"),h=n("1ffe"),v=n("a4be"),m={components:{uniIcons:i.default,tabs:r.default,recommend:s.default,follow:d.default,uniLoadMore:c.default},data:function(){return{recommendCurrent:1,articleCurrent:1,followCurrent:1,size:4,tabIndex:"",tablist:[{index:0,value:"推荐",isActive:!0},{index:1,value:"关注",isActive:!1}],recommendList:[],followList:[],loadStatus:"loading",recommendLoadStatus:"loading",followLoadStatus:"loading",isLoadMore:!1,isRecommend:!0}},created:function(){t("log","homebeforeCreate"," at pages/home/index.vue:100"),this.recommend(),this.follow()},onShow:function(){(0,u.check)()},onLoad:function(){if(t("log","onload"," at pages/home/index.vue:116"),!(0,h.getToken)())if((0,h.getOpenId)()){t("log","没有token，有OpenId"," at pages/home/index.vue:151");var e=(0,f.returnWxcode)();if(e){t("log","成功拿到code"," at pages/home/index.vue:156");var n={code:e,status:1};(0,v.login)(n).then((function(e){t("log",e.data.token," at pages/home/index.vue:164"),(0,h.setToken)(e.data.token)})).catch((function(e){t("log",e,"err login"," at pages/home/index.vue:167")}))}else(0,f.getWxCode)()}else{t("log","没有token，没有OpenId"," at pages/home/index.vue:120");var a=(0,f.returnWxcode)();if(a){t("log","成功拿到code"," at pages/home/index.vue:125");var o={code:a,status:1};(0,v.login)(o).then((function(e){t("log",e,"注册"," at pages/home/index.vue:132"),(0,h.setOpenId)(e.data.openid),1==e.data.status?(t("log","未登录"," at pages/home/index.vue:137"),uni.navigateTo({url:"../login/login"})):(t("log",e.data.token," at pages/home/index.vue:143"),(0,h.setToken)(e.data.token))})).catch((function(e){t("log",e,"err login"," at pages/home/index.vue:147")}))}else(0,f.getWxCode)()}},onReachBottom:function(){this.isLoadMore||this.tabIndex?!this.isLoadMore&&this.tabIndex&&(this.isLoadMore=!0,this.followCurrent+=1,this.follow()):(this.isLoadMore=!0,this.isRecommend?(this.recommendCurrent+=1,this.recommend()):(this.articleCurrent+=1,this.getArticleList()))},methods:{tabActive:function(t){this.tablist.map((function(e,n){e.isActive=t==n})),t?t&&(this.loadStatus=this.followLoadStatus,this.isLoadMore=!1):(this.loadStatus=this.recommendLoadStatus,this.isLoadMore=!1),this.tabIndex=t},recommend:function(){var t=this,e=this,n={current:this.recommendCurrent,size:this.size};(0,l.getRecommend)(n).then((function(n){e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(n.data)),n.data.length<e.size?(e.loadStatus="nomore",e.isRecommend=!1,e.getArticleList()):1===t.recommendCurrent?(e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(n.data))):setTimeout((function(){e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(n.data))}),3e3)}))},getArticleList:function(){var t=this,e=this,n={current:this.articleCurrent,size:this.size};(0,l.getArticleList)(n).then((function(n){e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(n.data)),n.data.length<e.size?(e.loadStatus="nomore",e.recommendLoadStatus="nomore"):1===t.articleCurrent?(e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(n.data))):setTimeout((function(){e.isLoadMore=!1,e.recommendList=[].concat((0,o.default)(e.recommendList),(0,o.default)(n.data))}),3e3)}))},follow:function(){var t=this,e=this,n={current:this.followCurrent,size:this.size};(0,l.getFollow)(n).then((function(n){e.followList=[].concat((0,o.default)(e.followList),(0,o.default)(n.data.records)),n.data.records.length<e.size?(e.loadStatus="nomore",e.follLoadStatus="nomore"):1===t.followCurrent?(e.isLoadMore=!1,e.followList=[].concat((0,o.default)(e.followList),(0,o.default)(n.data.records))):setTimeout((function(){e.isLoadMore=!1,e.followList=[].concat((0,o.default)(e.followList),(0,o.default)(n.data.records))}),3e3)}))},search:function(){uni.navigateTo({url:"../search/index"})}}};e.default=m}).call(this,n("0de9")["log"])},"0529":function(t,e,n){var a=n("1b3a");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var o=n("4f06").default;o("33158e76",a,!0,{sourceMap:!1,shadowMode:!1})},"0599":function(t,e,n){"use strict";n.r(e);var a=n("2393"),o=n("6f34");for(var i in o)"default"!==i&&function(t){n.d(e,t,(function(){return o[t]}))}(i);n("97b9");var r,s=n("f0c5"),d=Object(s["a"])(o["default"],a["b"],a["c"],!1,null,"1467d174",null,!1,a["a"],r);e["default"]=d.exports},"1b3a":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".home[data-v-1467d174],\n.home .content[data-v-1467d174]{background-color:#f7f7f7;margin-bottom:%?30?%}.header[data-v-1467d174]{position:fixed;z-index:10;width:100%;height:%?100?%;line-height:%?100?%;background-color:#fff;margin-bottom:%?15?%}.tag-nav[data-v-1467d174]{display:inline-block;height:100%;width:90%}.search[data-v-1467d174]{display:inline-block;width:8%}.search .search-icon[data-v-1467d174]{position:absolute;top:%?5?%}.content[data-v-1467d174]{background-color:#fff;padding-top:%?115?%}.loading[data-v-1467d174]{background-color:#fff}[data-v-1467d174] .uni-load-more .uni-load-more__text{font-size:%?28?%}[data-v-1467d174] .uni-load-more .uni-load-more__img{height:%?30?%!important;width:%?30?%!important}",""]),t.exports=e},2393:function(t,e,n){"use strict";n.d(e,"b",(function(){return o})),n.d(e,"c",(function(){return i})),n.d(e,"a",(function(){return a}));var a={uniIcons:n("11a2").default,uniLoadMore:n("19bf").default},o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"home"},[n("v-uni-view",{staticClass:"header"},[n("tabs",{staticClass:"tag-nav",attrs:{tabs:t.tablist},on:{tabActive:function(e){arguments[0]=e=t.$handleEvent(e),t.tabActive.apply(void 0,arguments)}}}),n("v-uni-view",{staticClass:"search",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.search()}}},[n("v-uni-view",{staticClass:"search-icon"},[n("uni-icons",{attrs:{type:"search",size:"25",color:"#999"}})],1)],1)],1),n("v-uni-view",{staticClass:"content"},[t.tabIndex?t._e():n("v-uni-view",t._l(t.recommendList,(function(t,e){return n("v-uni-view",[n("recommend",{attrs:{articleInfo:t}})],1)})),1),t.tabIndex?n("v-uni-view",t._l(t.followList,(function(t,e){return n("v-uni-view",{key:e},[n("follow",{attrs:{articleInfo:t}})],1)})),1):t._e(),n("v-uni-view",{directives:[{name:"show",rawName:"v-show",value:t.isLoadMore,expression:"isLoadMore"}]},[n("uni-load-more",{staticClass:"loading",attrs:{status:t.loadStatus,iconType:"circle"}})],1)],1)],1)},i=[]},3803:function(t,e,n){"use strict";n.r(e);var a=n("80b4"),o=n("cccf");for(var i in o)"default"!==i&&function(t){n.d(e,t,(function(){return o[t]}))}(i);n("9211");var r,s=n("f0c5"),d=Object(s["a"])(o["default"],a["b"],a["c"],!1,null,"34077204",null,!1,a["a"],r);e["default"]=d.exports},"459d":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.head-nav[data-v-06dc7823]{display:-webkit-flex;display:-webkit-box;display:flex;overflow:hidden;height:100%}.head-nav .head-nav-item[data-v-06dc7823]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center;font-size:%?34?%;font-family:Microsoft YaHei;font-weight:700;letter-spacing:%?4?%;color:#999}.head-nav-item .center[data-v-06dc7823]{position:relative;margin:auto;width:%?228?%}.active .center[data-v-06dc7823]{color:#404040}.active .center .head-nav-bottom[data-v-06dc7823]{position:relative;margin:auto;width:50%;height:%?8?%;top:%?-8?%;background-color:#404040}',""]),t.exports=e},"4c7a":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={name:"u-badge",props:{type:{type:String,default:"error"},size:{type:String,default:"default"},isDot:{type:Boolean,default:!1},count:{type:[Number,String]},overflowCount:{type:Number,default:99},showZero:{type:Boolean,default:!1},offset:{type:Array,default:function(){return[20,20]}},absolute:{type:Boolean,default:!0},fontSize:{type:[String,Number],default:"24"},color:{type:String,default:"#ffffff"},bgColor:{type:String,default:""},isCenter:{type:Boolean,default:!1}},computed:{boxStyle:function(){var t={};return this.isCenter?(t.top=0,t.right=0,t.transform="translateY(-50%) translateX(50%)"):(t.top=this.offset[0]+"rpx",t.right=this.offset[1]+"rpx",t.transform="translateY(0) translateX(0)"),"mini"==this.size&&(t.transform=t.transform+" scale(0.8)"),t},showText:function(){return this.isDot?"":this.count>this.overflowCount?"".concat(this.overflowCount,"+"):this.count},show:function(){return 0!=this.count||0!=this.showZero}}};e.default=a},"51b5":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-badge[data-v-34077204]{display:-webkit-inline-box;display:-webkit-inline-flex;display:inline-flex;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;line-height:%?24?%;padding:%?4?% %?8?%;border-radius:%?100?%;z-index:9}.u-badge--bg--primary[data-v-34077204]{background-color:#2979ff}.u-badge--bg--error[data-v-34077204]{background-color:#fa3534}.u-badge--bg--success[data-v-34077204]{background-color:#19be6b}.u-badge--bg--info[data-v-34077204]{background-color:#909399}.u-badge--bg--warning[data-v-34077204]{background-color:#f90}.u-badge-dot[data-v-34077204]{height:%?16?%;width:%?16?%;border-radius:%?100?%;line-height:1}.u-badge-mini[data-v-34077204]{-webkit-transform:scale(.8);transform:scale(.8);-webkit-transform-origin:center center;transform-origin:center center}.u-info[data-v-34077204]{background-color:#909399;color:#fff}',""]),t.exports=e},6022:function(t,e,n){"use strict";n.d(e,"b",(function(){return o})),n.d(e,"c",(function(){return i})),n.d(e,"a",(function(){return a}));var a={uBadge:n("3803").default},o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"tab-card"},[n("v-uni-view",{staticClass:"head-nav"},t._l(t.tabs,(function(e,a){return n("v-uni-view",{key:e.index,staticClass:"head-nav-item",class:e.isActive?"active":"",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.handlerActive(e.index)}}},[n("v-uni-view",{staticClass:"center"},[t._v(t._s(e.value)),0==e.index&&0==t.interactRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),1==e.index&&0==t.attentionRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),2==e.index&&0==t.otherRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),n("v-uni-view",{staticClass:"head-nav-bottom"})],1)],1)})),1),t._t("default")],2)},i=[]},"6f34":function(t,e,n){"use strict";n.r(e);var a=n("0391"),o=n.n(a);for(var i in a)"default"!==i&&function(t){n.d(e,t,(function(){return a[t]}))}(i);e["default"]=o.a},"6f65":function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={data:function(){return{}},props:{tabs:{type:Array,default:[]},type:{type:Boolean}},methods:{handlerActive:function(t){0==t?this.$store.commit("changeInteract",1):1==t?this.$store.commit("changeAttention",1):2==t&&this.$store.commit("changeOther",1),this.$emit("tabActive",t)}},computed:{interactRead:function(){return this.$store.state.interactRead},attentionRead:function(){return this.$store.state.attentionRead},otherRead:function(){return this.$store.state.otherRead}},mounted:function(){},created:function(){}};e.default=a},"80b4":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return o})),n.d(e,"c",(function(){return i})),n.d(e,"a",(function(){return a}));var o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return t.show?n("v-uni-view",{staticClass:"u-badge",class:[t.isDot?"u-badge-dot":"","mini"==t.size?"u-badge-mini":"",t.type?"u-badge--bg--"+t.type:""],style:[{top:t.offset[0]+"rpx",right:t.offset[1]+"rpx",fontSize:t.fontSize+"rpx",position:t.absolute?"absolute":"static",color:t.color,backgroundColor:t.bgColor},t.boxStyle]},[t._v(t._s(t.showText))]):t._e()},i=[]},9211:function(t,e,n){"use strict";var a=n("b325"),o=n.n(a);o.a},"97b9":function(t,e,n){"use strict";var a=n("0529"),o=n.n(a);o.a},a2be:function(t,e,n){"use strict";n.r(e);var a=n("6022"),o=n("d0d3");for(var i in o)"default"!==i&&function(t){n.d(e,t,(function(){return o[t]}))}(i);n("f866");var r,s=n("f0c5"),d=Object(s["a"])(o["default"],a["b"],a["c"],!1,null,"06dc7823",null,!1,a["a"],r);e["default"]=d.exports},ac2f:function(t,e,n){"use strict";(function(t){var a=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.check=o;a(n("c2da"));function o(){t("log","有未读消息"," at utils/checkUnRead.js:6"),uni.showTabBarRedDot({index:3})}}).call(this,n("0de9")["log"])},aef2:function(t,e,n){var a=n("459d");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var o=n("4f06").default;o("987548c8",a,!0,{sourceMap:!1,shadowMode:!1})},b325:function(t,e,n){var a=n("51b5");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var o=n("4f06").default;o("e307bc0a",a,!0,{sourceMap:!1,shadowMode:!1})},c2da:function(t,e,n){"use strict";var a=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.interactList=i,e.attentionList=r,e.otherList=s,e.unreadMessage=d,e.readMessage=c;var o=a(n("7692"));function i(t){return o.default.get("/api/message/user",{params:t})}function r(t){return o.default.get("/api/message/follow-notice",{params:t})}function s(t){return o.default.get("/api/message/system",{params:t})}function d(t){return o.default.get("/api/message/user/unread-count",{params:t})}function c(t){return o.default.post("/api/message/user",{type:t.type})}},cccf:function(t,e,n){"use strict";n.r(e);var a=n("4c7a"),o=n.n(a);for(var i in a)"default"!==i&&function(t){n.d(e,t,(function(){return a[t]}))}(i);e["default"]=o.a},d0d3:function(t,e,n){"use strict";n.r(e);var a=n("6f65"),o=n.n(a);for(var i in a)"default"!==i&&function(t){n.d(e,t,(function(){return a[t]}))}(i);e["default"]=o.a},f866:function(t,e,n){"use strict";var a=n("aef2"),o=n.n(a);o.a}}]);