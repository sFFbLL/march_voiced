(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-home-index"],{"0279":function(t,e,n){"use strict";n.r(e);var a=n("2c63"),r=n.n(a);for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);e["default"]=r.a},"04ff":function(t,e,n){var a=n("4675");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var r=n("4f06").default;r("674bf75a",a,!0,{sourceMap:!1,shadowMode:!1})},"105f":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".recommend-component[data-v-68d1538d]{background-color:#fff;margin:auto;margin-bottom:%?15?%;padding:%?30?% %?30?%}[data-v-68d1538d] .attention-cell .flex-item{border:none;padding:0;margin-top:%?17?%}",""]),t.exports=e},"1fc3":function(t,e,n){"use strict";n.d(e,"b",(function(){return r})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return a}));var a={uBadge:n("cb97").default},r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"tab-card"},[n("v-uni-view",{staticClass:"head-nav"},t._l(t.tabs,(function(e,a){return n("v-uni-view",{key:e.index,staticClass:"head-nav-item",class:e.isActive?"active":"",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.handlerActive(e.index)}}},[n("v-uni-view",{staticClass:"center"},[t._v(t._s(e.value)),0==e.index&&0==t.interactRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),1==e.index&&0==t.attentionRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),2==e.index&&0==t.otherRead&&t.type?n("u-badge",{attrs:{"is-dot":!0,type:"error"}}):t._e(),n("v-uni-view",{staticClass:"head-nav-bottom"})],1)],1)})),1),t._t("default")],2)},o=[]},"2c63":function(t,e,n){"use strict";var a=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var r=a(n("d50f")),o=a(n("7cb0")),i=a(n("21b7")),c=a(n("40e6")),s={data:function(){return{}},props:{articleInfo:{type:Object,default:{}}},components:{articleTitle:r.default,articleContent:o.default,articleInteract:i.default,attentionAndFansCell:c.default},methods:{},mounted:function(){}};e.default=s},3848:function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return r})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return a}));var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"recommend-component"},[n("attentionAndFansCell",{attrs:{aid:t.articleInfo.create_by,nickname:t.articleInfo.nickname,avatarPath:t.articleInfo.avatarPath}},[n("v-uni-view",{staticClass:"followMessage",attrs:{slot:"underText"},slot:"underText"},[0===t.articleInfo.status?n("v-uni-text",[t._v("发布")]):t._e(),1===t.articleInfo.status?n("v-uni-text",[t._v("评论")]):t._e(),2===t.articleInfo.status?n("v-uni-text",[t._v("收藏")]):t._e(),3===t.articleInfo.status?n("v-uni-text",[t._v("转载")]):t._e(),4===t.articleInfo.status?n("v-uni-text",[t._v("赞")]):t._e(),n("v-uni-text",[t._v("了该内容")])],1)],1),n("articleTitle",{attrs:{articleTitle:t.articleInfo.title,articleId:t.articleInfo.articleId}}),n("articleContent",{attrs:{articleContent:t.articleInfo.describe,articleImg:t.articleInfo.image,id:t.articleInfo.articleId}}),n("articleInteract",{attrs:{favourTotal:t.articleInfo.favourTotal,commentTotal:t.articleInfo.commentTotal,dateTime:t.articleInfo.createTime}})],1)},o=[]},"463f":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return r})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return a}));var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return t.show?n("v-uni-view",{staticClass:"u-badge",class:[t.isDot?"u-badge-dot":"","mini"==t.size?"u-badge-mini":"",t.type?"u-badge--bg--"+t.type:""],style:[{top:t.offset[0]+"rpx",right:t.offset[1]+"rpx",fontSize:t.fontSize+"rpx",position:t.absolute?"absolute":"static",color:t.color,backgroundColor:t.bgColor},t.boxStyle]},[t._v(t._s(t.showText))]):t._e()},o=[]},4675:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-badge[data-v-34077204]{display:-webkit-inline-box;display:-webkit-inline-flex;display:inline-flex;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;line-height:%?24?%;padding:%?4?% %?8?%;border-radius:%?100?%;z-index:9}.u-badge--bg--primary[data-v-34077204]{background-color:#2979ff}.u-badge--bg--error[data-v-34077204]{background-color:#fa3534}.u-badge--bg--success[data-v-34077204]{background-color:#19be6b}.u-badge--bg--info[data-v-34077204]{background-color:#909399}.u-badge--bg--warning[data-v-34077204]{background-color:#f90}.u-badge-dot[data-v-34077204]{height:%?16?%;width:%?16?%;border-radius:%?100?%;line-height:1}.u-badge-mini[data-v-34077204]{-webkit-transform:scale(.8);transform:scale(.8);-webkit-transform-origin:center center;transform-origin:center center}.u-info[data-v-34077204]{background-color:#909399;color:#fff}',""]),t.exports=e},"4dff":function(t,e,n){"use strict";n.r(e);var a=n("83b8"),r=n("5606");for(var o in r)"default"!==o&&function(t){n.d(e,t,(function(){return r[t]}))}(o);n("f237");var i,c=n("f0c5"),s=Object(c["a"])(r["default"],a["b"],a["c"],!1,null,"5235a0bc",null,!1,a["a"],i);e["default"]=s.exports},"50fc":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.head-nav[data-v-06dc7823]{display:-webkit-flex;display:-webkit-box;display:flex;overflow:hidden;height:100%}.head-nav .head-nav-item[data-v-06dc7823]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center;font-size:%?34?%;font-family:Microsoft YaHei;font-weight:700;letter-spacing:%?4?%;color:#999}.head-nav-item .center[data-v-06dc7823]{position:relative;margin:auto;width:%?228?%}.active .center[data-v-06dc7823]{color:#404040}.active .center .head-nav-bottom[data-v-06dc7823]{position:relative;margin:auto;width:50%;height:%?8?%;top:%?-8?%;background-color:#404040}',""]),t.exports=e},5606:function(t,e,n){"use strict";n.r(e);var a=n("fd0c"),r=n.n(a);for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);e["default"]=r.a},"72c4":function(t,e,n){"use strict";var a=n("ac56"),r=n.n(a);r.a},"7b34":function(t,e,n){"use strict";n.r(e);var a=n("86f5"),r=n("e2bb");for(var o in r)"default"!==o&&function(t){n.d(e,t,(function(){return r[t]}))}(o);n("f077");var i,c=n("f0c5"),s=Object(c["a"])(r["default"],a["b"],a["c"],!1,null,"68d1538d",null,!1,a["a"],i);e["default"]=s.exports},"83b8":function(t,e,n){"use strict";n.d(e,"b",(function(){return r})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return a}));var a={uniIcons:n("34f5").default,uniLoadMore:n("d508").default},r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"home"},[n("v-uni-view",{staticClass:"header"},[n("tabs",{staticClass:"tag-nav",attrs:{tabs:t.tablist},on:{tabActive:function(e){arguments[0]=e=t.$handleEvent(e),t.tabActive.apply(void 0,arguments)}}}),n("v-uni-view",{staticClass:"search",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.search()}}},[n("v-uni-view",{staticClass:"search-icon"},[n("uni-icons",{attrs:{type:"search",size:"25",color:"#999"}})],1)],1)],1),n("v-uni-view",{staticClass:"content"},[t.tabIndex?t._e():n("v-uni-view",t._l(t.recommendList,(function(e,a){return n("v-uni-view",[t.recommendList.length?n("recommend",{attrs:{articleInfo:e}}):t._e()],1)})),1),t.tabIndex?n("v-uni-view",t._l(t.followList,(function(t,e){return n("v-uni-view",{key:e},[n("follow",{attrs:{articleInfo:t}})],1)})),1):t._e(),n("v-uni-view",{directives:[{name:"show",rawName:"v-show",value:t.isLoadMore,expression:"isLoadMore"}]},[n("uni-load-more",{staticClass:"loading",attrs:{status:t.loadStatus,iconType:"circle"}})],1)],1)],1)},o=[]},"86f5":function(t,e,n){"use strict";var a;n.d(e,"b",(function(){return r})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return a}));var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"recommend-component"},[n("articleTitle",{attrs:{articleTitle:t.articleInfo.title,articleId:t.articleInfo.id}}),t.articleInfo.create_by?n("attentionAndFansCell",{attrs:{aid:t.articleInfo.create_by,nickname:t.articleInfo.nickname,avatarPath:t.articleInfo.avatarPath,isFollow:t.articleInfo.isFollow}}):t._e(),n("articleContent",{attrs:{articleContent:t.articleInfo.describe,articleImg:t.articleInfo.image,id:t.articleInfo.id}}),t.isArticleInteract?n("articleInteract",{attrs:{favourTotal:t.articleInfo.favourTotal,commentTotal:t.articleInfo.commentTotal,collectTotal:t.articleInfo.collectTotal}}):t._e()],1)},o=[]},"8a32":function(t,e,n){"use strict";n.r(e);var a=n("8c76"),r=n.n(a);for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);e["default"]=r.a},"8bb7":function(t,e,n){var a=n("90a3");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var r=n("4f06").default;r("5ecf7a84",a,!0,{sourceMap:!1,shadowMode:!1})},"8bc3":function(t,e,n){"use strict";n.r(e);var a=n("923a"),r=n.n(a);for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);e["default"]=r.a},"8c76":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={name:"u-badge",props:{type:{type:String,default:"error"},size:{type:String,default:"default"},isDot:{type:Boolean,default:!1},count:{type:[Number,String]},overflowCount:{type:Number,default:99},showZero:{type:Boolean,default:!1},offset:{type:Array,default:function(){return[20,20]}},absolute:{type:Boolean,default:!0},fontSize:{type:[String,Number],default:"24"},color:{type:String,default:"#ffffff"},bgColor:{type:String,default:""},isCenter:{type:Boolean,default:!1}},computed:{boxStyle:function(){var t={};return this.isCenter?(t.top=0,t.right=0,t.transform="translateY(-50%) translateX(50%)"):(t.top=this.offset[0]+"rpx",t.right=this.offset[1]+"rpx",t.transform="translateY(0) translateX(0)"),"mini"==this.size&&(t.transform=t.transform+" scale(0.8)"),t},showText:function(){return this.isDot?"":this.count>this.overflowCount?"".concat(this.overflowCount,"+"):this.count},show:function(){return 0!=this.count||0!=this.showZero}}};e.default=a},"901f":function(t,e,n){"use strict";n.r(e);var a=n("3848"),r=n("0279");for(var o in r)"default"!==o&&function(t){n.d(e,t,(function(){return r[t]}))}(o);n("72c4");var i,c=n("f0c5"),s=Object(c["a"])(r["default"],a["b"],a["c"],!1,null,"796498b4",null,!1,a["a"],i);e["default"]=s.exports},"90a3":function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".home[data-v-5235a0bc],\n.home .content[data-v-5235a0bc]{background-color:#f7f7f7;margin-bottom:%?30?%}.header[data-v-5235a0bc]{position:fixed;z-index:10;width:100%;height:%?100?%;line-height:%?100?%;background-color:#fff;margin-bottom:%?15?%}.tag-nav[data-v-5235a0bc]{display:inline-block;height:100%;width:90%}.search[data-v-5235a0bc]{display:inline-block;width:8%}.search .search-icon[data-v-5235a0bc]{position:absolute;top:%?5?%}.content[data-v-5235a0bc]{background-color:#fff;padding-top:%?115?%}.loading[data-v-5235a0bc]{background-color:#fff}[data-v-5235a0bc] .uni-load-more .uni-load-more__text{font-size:%?28?%}[data-v-5235a0bc] .uni-load-more .uni-load-more__img{height:%?30?%!important;width:%?30?%!important}",""]),t.exports=e},"923a":function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a={data:function(){return{}},props:{tabs:{type:Array,default:[]},type:{type:Boolean}},methods:{handlerActive:function(t){0==t?this.$store.commit("changeInteract",1):1==t?this.$store.commit("changeAttention",1):2==t&&this.$store.commit("changeOther",1),this.$emit("tabActive",t)}},computed:{interactRead:function(){return this.$store.state.interactRead},attentionRead:function(){return this.$store.state.attentionRead},otherRead:function(){return this.$store.state.otherRead}},mounted:function(){},created:function(){}};e.default=a},"98e0":function(t,e,n){var a=n("105f");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var r=n("4f06").default;r("74b4fba6",a,!0,{sourceMap:!1,shadowMode:!1})},ac56:function(t,e,n){var a=n("f988");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var r=n("4f06").default;r("6cb6776f",a,!0,{sourceMap:!1,shadowMode:!1})},b37c:function(t,e,n){var a=n("50fc");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var r=n("4f06").default;r("0ffdf5f4",a,!0,{sourceMap:!1,shadowMode:!1})},cb97:function(t,e,n){"use strict";n.r(e);var a=n("463f"),r=n("8a32");for(var o in r)"default"!==o&&function(t){n.d(e,t,(function(){return r[t]}))}(o);n("d67e");var i,c=n("f0c5"),s=Object(c["a"])(r["default"],a["b"],a["c"],!1,null,"34077204",null,!1,a["a"],i);e["default"]=s.exports},cdf3:function(t,e,n){"use strict";var a=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var r=a(n("d50f")),o=a(n("7cb0")),i=a(n("21b7")),c=a(n("40e6")),s={data:function(){return{}},props:{articleInfo:{type:Object,default:{}},isArticleInteract:{type:Boolean,default:!0}},components:{articleTitle:r.default,articleContent:o.default,articleInteract:i.default,attentionAndFansCell:c.default},methods:{},mounted:function(){}};e.default=s},d040:function(t,e,n){"use strict";var a=n("b37c"),r=n.n(a);r.a},d67e:function(t,e,n){"use strict";var a=n("04ff"),r=n.n(a);r.a},d6e3:function(t,e,n){"use strict";n.r(e);var a=n("1fc3"),r=n("8bc3");for(var o in r)"default"!==o&&function(t){n.d(e,t,(function(){return r[t]}))}(o);n("d040");var i,c=n("f0c5"),s=Object(c["a"])(r["default"],a["b"],a["c"],!1,null,"06dc7823",null,!1,a["a"],i);e["default"]=s.exports},e220:function(t,e,n){"use strict";var a=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.getRecommend=o,e.getArticleList=i,e.getFollow=c,e.searchArticle=s,e.searchUser=u;var r=a(n("367e"));function o(t){return r.default.get("/api/article/top",{params:t})}function i(t){return r.default.get("/api/article/index",{params:t})}function c(t){return r.default.get("/api/message/follow",{params:t})}function s(t){return r.default.get("/api/base/searchArticle",{params:t})}function u(t){return r.default.get("/api/base/searchUser",{params:t})}},e2bb:function(t,e,n){"use strict";n.r(e);var a=n("cdf3"),r=n.n(a);for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);e["default"]=r.a},f077:function(t,e,n){"use strict";var a=n("98e0"),r=n.n(a);r.a},f237:function(t,e,n){"use strict";var a=n("8bb7"),r=n.n(a);r.a},f988:function(t,e,n){var a=n("24fb");e=a(!1),e.push([t.i,".recommend-component[data-v-796498b4]{background-color:#fff;margin:auto;margin-bottom:%?15?%;padding:%?30?% %?30?%}[data-v-796498b4] .attention-cell .flex-item{border:none;padding:0}[data-v-796498b4] .article-title{margin-top:%?17?%}.followMessage[data-v-796498b4]{font-size:%?26?%;width:%?420?%;color:#969696}",""]),t.exports=e},fd0c:function(t,e,n){"use strict";var a=n("4ea4");n("99af"),n("d81d"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var r=a(n("2909")),o=a(n("34f5")),i=a(n("d6e3")),c=a(n("7b34")),s=a(n("901f")),u=a(n("d508")),l=(n("0e07"),n("e220")),d=n("4b28"),f=(n("dccd"),n("c98b")),v=(n("5c16"),{components:{uniIcons:o.default,tabs:i.default,recommend:c.default,follow:s.default,uniLoadMore:u.default},data:function(){return{recommendCurrent:1,articleCurrent:1,followCurrent:1,size:4,tabIndex:"",tablist:[{index:0,value:"推荐",isActive:!0},{index:1,value:"关注",isActive:!1}],recommendList:[],followList:[],loadStatus:"loading",recommendLoadStatus:"loading",followLoadStatus:"loading",isLoadMore:!1,isRecommend:!0}},created:function(){(0,f.getToken)()?(this.recommend(),this.follow()):((0,d.forLogin)(),this.recommend(),this.follow())},onPullDownRefresh:function(){this.recommendCurrent=1,this.articleCurrent=1,this.followCurrent=1;var t=this;t.recommendList=[],this.recommend(),this.follow(),setTimeout((function(){uni.stopPullDownRefresh()}),2e3)},onReachBottom:function(){this.isLoadMore||this.tabIndex?!this.isLoadMore&&this.tabIndex&&(this.isLoadMore=!0,this.followCurrent+=1,this.follow()):(this.isLoadMore=!0,this.isRecommend?(this.recommendCurrent+=1,this.recommend()):(this.articleCurrent+=1,this.getArticleList()))},methods:{tabActive:function(t){this.tablist.map((function(e,n){e.isActive=t==n})),t?t&&(this.loadStatus=this.followLoadStatus,this.isLoadMore=!1):(this.loadStatus=this.recommendLoadStatus,this.isLoadMore=!1),this.tabIndex=t},recommend:function(){var t=this,e=this,n={current:this.recommendCurrent,size:this.size};(0,l.getRecommend)(n).then((function(n){e.recommendList=[].concat((0,r.default)(e.recommendList),(0,r.default)(n.data)),n.data.length<e.size?(e.loadStatus="nomore",e.isRecommend=!1,e.getArticleList()):1===t.recommendCurrent?(e.isLoadMore=!1,e.recommendList=[].concat((0,r.default)(e.recommendList),(0,r.default)(n.data))):setTimeout((function(){e.isLoadMore=!1,e.recommendList=[].concat((0,r.default)(e.recommendList),(0,r.default)(n.data))}),3e3)}))},getArticleList:function(){var t=this,e=this,n={current:this.articleCurrent,size:this.size};(0,l.getArticleList)(n).then((function(n){e.recommendList=[].concat((0,r.default)(e.recommendList),(0,r.default)(n.data)),n.data.length<e.size?(e.loadStatus="nomore",e.recommendLoadStatus="nomore"):1===t.articleCurrent?(e.isLoadMore=!1,e.recommendList=[].concat((0,r.default)(e.recommendList),(0,r.default)(n.data))):setTimeout((function(){e.isLoadMore=!1,e.recommendList=[].concat((0,r.default)(e.recommendList),(0,r.default)(n.data))}),3e3)}))},follow:function(){var t=this,e=this,n={current:this.followCurrent,size:this.size};(0,l.getFollow)(n).then((function(n){0===n.code&&(e.followList=[].concat((0,r.default)(e.followList),(0,r.default)(n.data.records)),n.data.records.length<e.size?(e.loadStatus="nomore",e.follLoadStatus="nomore"):1===t.followCurrent?(e.isLoadMore=!1,e.followList=[].concat((0,r.default)(e.followList),(0,r.default)(n.data.records))):setTimeout((function(){e.isLoadMore=!1,e.followList=[].concat((0,r.default)(e.followList),(0,r.default)(n.data.records))}),3e3))}))},search:function(){uni.navigateTo({url:"../search/index"})}}});e.default=v}}]);