(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-articleDetails-index"],{"0187":function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-drawer[data-v-740af5f8]{display:block;position:fixed;top:0;left:0;right:0;bottom:0;overflow:hidden}.u-drawer-content[data-v-740af5f8]{display:block;position:absolute;z-index:1003;-webkit-transition:all .25s linear;transition:all .25s linear}.u-drawer__scroll-view[data-v-740af5f8]{width:100%;height:100%}.u-drawer-left[data-v-740af5f8]{top:0;bottom:0;left:0;background-color:#fff}.u-drawer-right[data-v-740af5f8]{right:0;top:0;bottom:0;background-color:#fff}.u-drawer-top[data-v-740af5f8]{top:0;left:0;right:0;background-color:#fff}.u-drawer-bottom[data-v-740af5f8]{bottom:0;left:0;right:0;background-color:#fff}.u-drawer-center[data-v-740af5f8]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;\r\nflex-direction:column;bottom:0;left:0;right:0;top:0;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;opacity:0;z-index:99999}.u-mode-center-box[data-v-740af5f8]{min-width:%?100?%;min-height:%?100?%;display:block;position:relative;background-color:#fff}.u-drawer-content-visible.u-drawer-center[data-v-740af5f8]{-webkit-transform:scale(1);transform:scale(1);opacity:1}.u-animation-zoom[data-v-740af5f8]{-webkit-transform:scale(1.15);transform:scale(1.15)}.u-drawer-content-visible[data-v-740af5f8]{-webkit-transform:translateZ(0)!important;transform:translateZ(0)!important}.u-close[data-v-740af5f8]{position:absolute;z-index:3}.u-close--top-left[data-v-740af5f8]{top:%?30?%;left:%?30?%}.u-close--top-right[data-v-740af5f8]{top:%?30?%;right:%?30?%}.u-close--bottom-left[data-v-740af5f8]{bottom:%?30?%;left:%?30?%}.u-close--bottom-right[data-v-740af5f8]{right:%?30?%;bottom:%?30?%}',""]),t.exports=e},"079b":function(t,e,n){"use strict";(function(t){var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("1e61")),o={data:function(){return{}},props:{commentList:{type:Array,default:[]},kidsCommentCount:{type:Array,default:[]}},created:function(){this.commentList.ChildComments.length},methods:{format:function(t){var e=new Date(t),n=(0,a.default)(e).format("YYYY-MM-DD HH:mm:ss");return n},toAllReply:function(){t("log","查看更多"," at marchVoiceComponents/comment/index.vue:72")},addComment:function(t,e,n,i){var a={childComment:!0,index:n,replyId:t,follewId:e,replyName:i};this.$emit("childFn",a)}}};e.default=o}).call(this,n("0de9")["log"])},"15ed":function(t,e,n){"use strict";n.r(e);var i=n("6429"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},1874:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.getMarchCircleInfo=o,e.marchCircleList=r,e.publishIdea=c,e.joinMarchCircle=s,e.changeFavour=l,e.ideaDetail=d,e.publishComment=u,e.ideaCommentList=f,e.ideaChildCommentList=m;var a=i(n("367e"));function o(t){return a.default.get("/api/march/msg",{params:t})}function r(t){return a.default.get("/api/march",{params:t})}function c(t){return a.default.post("/api/march",{content:t.content,imageList:t.imageList,image:t.image})}function s(t){return a.default.post("/api/apply/march",{params:t})}function l(t){return a.default.post("/api/favour/march",{id:t.id,type:t.type})}function d(t){return a.default.get("/api/march/detail/".concat(t))}function u(t){return a.default.post("/api/march/comment",{params:t})}function f(t){return a.default.get("/api/comment/marchsoft",{params:t})}function m(t){return a.default.get("/api/comment/children/article",{id:t.id,current:t.current,size:t.size})}},"18f1":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"u-mask",class:{"u-mask-zoom":t.zoom,"u-mask-show":t.show},style:[t.maskStyle,t.zoomStyle],attrs:{"hover-stop-propagation":!0},on:{touchmove:function(e){e.stopPropagation(),e.preventDefault(),arguments[0]=e=t.$handleEvent(e),function(){}.apply(void 0,arguments)},click:function(e){arguments[0]=e=t.$handleEvent(e),t.click.apply(void 0,arguments)}}},[t._t("default")],2)},o=[]},"1c0ba":function(t,e,n){"use strict";var i=n("c9ba"),a=n.n(i);a.a},2184:function(t,e,n){"use strict";n.r(e);var i=n("3cd2"),a=n("dc97");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("7a90");var r,c=n("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"057bfa41",null,!1,i["a"],r);e["default"]=s.exports},2909:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=s;var i=c(n("6005")),a=c(n("db90")),o=c(n("06c5")),r=c(n("3427"));function c(t){return t&&t.__esModule?t:{default:t}}function s(t){return(0,i.default)(t)||(0,a.default)(t)||(0,o.default)(t)||(0,r.default)()}},"2c25":function(t,e,n){"use strict";var i=n("cf82"),a=n.n(i);a.a},3427:function(t,e,n){"use strict";function i(){throw new TypeError("Invalid attempt to spread non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}Object.defineProperty(e,"__esModule",{value:!0}),e.default=i},"349b":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={data:function(){return{}},props:{articleId:{type:Number,default:null},articleTitle:{type:String,default:""}},components:{},methods:{articleDetails:function(){uni.navigateTo({url:"../articleDetails/index?id="+this.articleId})}},mounted:function(){}};e.default=i},"3cd2":function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var i={uPopup:n("43dd").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",[t.show?t._e():n("v-uni-view",{staticClass:"commentwarp"},[n("v-uni-view",{staticClass:"comment"},[n("v-uni-view",{staticClass:"comment-view"},[n("v-uni-textarea",{staticClass:"commentinput",attrs:{"auto-height":!0,maxlength:"-1",placeholder:"请输入评论"},on:{input:function(e){arguments[0]=e=t.$handleEvent(e),t.onkeyinput.apply(void 0,arguments)}}})],1),n("v-uni-button",{staticClass:"sendbtn",attrs:{size:"mini",disabled:t.sendDisabled},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.sendComment()}}},[t._v("发布")])],1)],1),n("u-popup",{attrs:{mode:"bottom","border-radius":"10"},on:{close:function(e){arguments[0]=e=t.$handleEvent(e),t.close.apply(void 0,arguments)}},model:{value:t.show,callback:function(e){t.show=e},expression:"show"}},[n("v-uni-view",{staticClass:"comment"},[n("v-uni-view",{staticClass:"comment-view"},[n("v-uni-textarea",{staticClass:"commentinput",attrs:{"auto-height":!0,maxlength:"-1",placeholder:"请输入评论"},on:{input:function(e){arguments[0]=e=t.$handleEvent(e),t.onkeyinput.apply(void 0,arguments)}}})],1),n("v-uni-button",{staticClass:"sendbtn",attrs:{size:"mini",disabled:t.sendDisabled},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.sendComment()}}},[t._v("发布")])],1)],1)],1)},o=[]},"3e8a":function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".comment[data-v-057bfa41]{display:-webkit-box;display:-webkit-flex;display:flex;margin-top:%?30?%}\n\n/* 评论 */.commentinput[data-v-057bfa41]{line-height:%?62?%;width:%?550?%}.sendbtn[data-v-057bfa41]{height:%?68?%;width:%?140?%;background-color:#fff;color:#007aff;font-size:%?28?%;margin-bottom:%?30?%}[data-v-057bfa41] .sendbtn::after{border:none}.comment-view[data-v-057bfa41]{max-height:%?126?%;overflow:scroll}.commentwarp[data-v-057bfa41]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;width:100%;font-size:%?28?%;padding:%?15?% %?28?%;text-align:left;position:fixed;bottom:0;color:#303133;border-top:%?2?% solid #efefef;border-radius:%?20?%;background-color:#fff}",""]),t.exports=e},"43dd":function(t,e,n){"use strict";n.r(e);var i=n("b23f"),a=n("56b1");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("6186");var r,c=n("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"740af5f8",null,!1,i["a"],r);e["default"]=s.exports},4788:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".article-title .article-text[data-v-74809812]{color:#404040;font-size:%?34?%;font-family:微软雅黑;text-align:left;font-weight:700;overflow:hidden;-webkit-line-clamp:2;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;line-height:150%}.article-text[data-v-74809812]{display:inline-block;width:100%}",""]),t.exports=e},"4c6b":function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'.article-details[data-v-90dcf2e0]{background-color:#f7f7f7}.view-content[data-v-90dcf2e0]{padding:0 %?30?%;background-color:#fff}.article-title[data-v-90dcf2e0]{color:#404040;font-size:%?36?%;line-height:150%;font-weight:700}[data-v-90dcf2e0] .attention-cell .flex-item{padding:%?20?% 0;border:none}[data-v-90dcf2e0] .attention-cell .left-img-box{margin-right:0}[data-v-90dcf2e0] .attention-cell .right-button{background-color:#fff}.article-message[data-v-90dcf2e0]{color:#999797;font-size:%?24?%}.word-num-num[data-v-90dcf2e0]{margin-left:%?4?%}.time[data-v-90dcf2e0]{float:right}.article-content[data-v-90dcf2e0]{padding:%?20?% %?10?%}[data-v-90dcf2e0] .article-content .ql-container{margin:0;width:100%;color:#404040;font-size:%?32?%;line-height:150%}[data-v-90dcf2e0] .article-content .container{padding-top:%?20?%}[data-v-90dcf2e0] .article-content .ql-container .ql-editor{text-align:justify}[data-v-90dcf2e0] .article-content .tool-view{display:none}\n\n/* 底部导航:点赞，评论等 */.bottom-nav[data-v-90dcf2e0]{position:fixed;bottom:0;width:100%;height:%?120?%;border:%?1?% solid #f0f0f0;border-width:%?1?% 0;background-color:#fff}.bottom-nav-centre[data-v-90dcf2e0]{display:-webkit-box;display:-webkit-flex;display:flex;margin:0 %?50?%;height:100%}.attention-icon[data-v-90dcf2e0]::before{content:"";display:inline-block;vertical-align:middle;height:100%}.attention-icon[data-v-90dcf2e0]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center}[data-v-90dcf2e0] .attention-icon .u-icon{font-size:%?50?%!important;color:#404040!important}.attention-item[data-v-90dcf2e0]{position:relative;display:inline-block}.attention-icon .attention-num[data-v-90dcf2e0]{position:absolute;top:%?-8?%}\n\n/* 评论样式 */.top-title[data-v-90dcf2e0]{height:%?88?%;color:#505050;font-size:%?34?%;line-height:150%;font-weight:700}.top-title uni-text[data-v-90dcf2e0]:first-child{display:inline-block;margin:0 %?10?% 0 %?30?%;line-height:%?88?%}.comment-view[data-v-90dcf2e0]{margin-bottom:%?120?%}.comment-list[data-v-90dcf2e0]{padding:0 %?30?%;background-color:#fff}',""]),t.exports=e},"4edc":function(t,e,n){"use strict";n("99af"),n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={name:"u-popup",props:{show:{type:Boolean,default:!1},mode:{type:String,default:"left"},mask:{type:Boolean,default:!0},length:{type:[Number,String],default:"auto"},zoom:{type:Boolean,default:!0},safeAreaInsetBottom:{type:Boolean,default:!1},maskCloseAble:{type:Boolean,default:!0},customStyle:{type:Object,default:function(){return{}}},value:{type:Boolean,default:!1},popup:{type:Boolean,default:!0},borderRadius:{type:[Number,String],default:0},zIndex:{type:[Number,String],default:""},closeable:{type:Boolean,default:!1},closeIcon:{type:String,default:"close"},closeIconPos:{type:String,default:"top-right"},closeIconColor:{type:String,default:"#909399"},closeIconSize:{type:[String,Number],default:"30"},width:{type:String,default:""},height:{type:String,default:""},negativeTop:{type:[String,Number],default:0},maskCustomStyle:{type:Object,default:function(){return{}}},duration:{type:[String,Number],default:250}},data:function(){return{visibleSync:!1,showDrawer:!1,timer:null,closeFromInner:!1}},computed:{style:function(){var t={};if("left"==this.mode||"right"==this.mode?t={width:this.width?this.getUnitValue(this.width):this.getUnitValue(this.length),height:"100%",transform:"translate3D(".concat("left"==this.mode?"-100%":"100%",",0px,0px)")}:"top"!=this.mode&&"bottom"!=this.mode||(t={width:"100%",height:this.height?this.getUnitValue(this.height):this.getUnitValue(this.length),transform:"translate3D(0px,".concat("top"==this.mode?"-100%":"100%",",0px)")}),t.zIndex=this.uZindex,this.borderRadius){switch(this.mode){case"left":t.borderRadius="0 ".concat(this.borderRadius,"rpx ").concat(this.borderRadius,"rpx 0");break;case"top":t.borderRadius="0 0 ".concat(this.borderRadius,"rpx ").concat(this.borderRadius,"rpx");break;case"right":t.borderRadius="".concat(this.borderRadius,"rpx 0 0 ").concat(this.borderRadius,"rpx");break;case"bottom":t.borderRadius="".concat(this.borderRadius,"rpx ").concat(this.borderRadius,"rpx 0 0");break;default:}t.overflow="hidden"}return this.duration&&(t.transition="all ".concat(this.duration/1e3,"s linear")),t},centerStyle:function(){var t={};return t.width=this.width?this.getUnitValue(this.width):this.getUnitValue(this.length),t.height=this.height?this.getUnitValue(this.height):"auto",t.zIndex=this.uZindex,t.marginTop="-".concat(this.$u.addUnit(this.negativeTop)),this.borderRadius&&(t.borderRadius="".concat(this.borderRadius,"rpx"),t.overflow="hidden"),t},uZindex:function(){return this.zIndex?this.zIndex:this.$u.zIndex.popup}},watch:{value:function(t){t?this.open():this.closeFromInner||this.close(),this.closeFromInner=!1}},mounted:function(){this.value&&this.open()},methods:{getUnitValue:function(t){return/(%|px|rpx|auto)$/.test(t)?t:t+"rpx"},maskClick:function(){this.close()},close:function(){this.closeFromInner=!0,this.change("showDrawer","visibleSync",!1)},modeCenterClose:function(t){"center"==t&&this.maskCloseAble&&this.close()},open:function(){this.change("visibleSync","showDrawer",!0)},change:function(t,e,n){var i=this;1==this.popup&&this.$emit("input",n),this[t]=n,this.timer=n?setTimeout((function(){i[e]=n,i.$emit(n?"open":"close")}),50):setTimeout((function(){i[e]=n,i.$emit(n?"open":"close")}),this.duration)}}};e.default=i},"564f":function(t,e,n){"use strict";n.r(e);var i=n("079b"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},"56b1":function(t,e,n){"use strict";n.r(e);var i=n("4edc"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},6005:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=o;var i=a(n("6b75"));function a(t){return t&&t.__esModule?t:{default:t}}function o(t){if(Array.isArray(t))return(0,i.default)(t)}},6186:function(t,e,n){"use strict";var i=n("7f9d"),a=n.n(i);a.a},6429:function(t,e,n){"use strict";var i=n("4ea4");n("a9e3"),n("b64b"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("5530")),o={name:"u-mask",props:{show:{type:Boolean,default:!1},zIndex:{type:[Number,String],default:""},customStyle:{type:Object,default:function(){return{}}},zoom:{type:Boolean,default:!0},duration:{type:[Number,String],default:300},maskClickAble:{type:Boolean,default:!0}},data:function(){return{zoomStyle:{transform:""},scale:"scale(1.2, 1.2)"}},watch:{show:function(t){t&&this.zoom?this.zoomStyle.transform="scale(1, 1)":!t&&this.zoom&&(this.zoomStyle.transform=this.scale)}},computed:{maskStyle:function(){var t={backgroundColor:"rgba(0, 0, 0, 0.6)"};return this.show?t.zIndex=this.zIndex?this.zIndex:this.$u.zIndex.mask:t.zIndex=-1,t.transition="all ".concat(this.duration/1e3,"s ease-in-out"),Object.keys(this.customStyle).length&&(t=(0,a.default)((0,a.default)({},t),this.customStyle)),t}},methods:{click:function(){this.maskClickAble&&this.$emit("click")}}};e.default=o},6782:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.information=o,e.modInformation=r;var a=i(n("367e"));function o(t){return a.default.get("/api/user/information",{id:t})}function r(t){return a.default.put("/api/user/modInformation",{nickname:nickname,sex:sex,signature:signature,avatarPath:avatarPath})}},"784d":function(t,e,n){"use strict";n.r(e);var i=n("9c76"),a=n("8140");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("2c25");var r,c=n("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"90dcf2e0",null,!1,i["a"],r);e["default"]=s.exports},"7a90":function(t,e,n){"use strict";var i=n("c8db"),a=n.n(i);a.a},"7a94":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",t._l(t.commentList,(function(e,i){return n("v-uni-view",{staticClass:"comment"},[n("v-uni-view",{staticClass:"left"},[n("v-uni-image",{attrs:{src:e.idAvatar?e.idAvatar:null,mode:"aspectFill"}})],1),n("v-uni-view",{staticClass:"right"},[n("v-uni-view",{staticClass:"top"},[n("v-uni-view",{staticClass:"name"},[t._v(t._s(e.createByName))])],1),n("v-uni-view",{staticClass:"content",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.addComment(0,e.id,i)}}},[t._v(t._s(e.content))]),n("v-uni-view",{staticClass:"bottom"},[t._v(t._s(t.format(e.createTime)))]),n("v-uni-view",{staticClass:"reply-box"},[t._l(e.ChildComments,(function(a,o){return n("v-uni-view",[n("v-uni-view",{staticClass:"item"},[n("v-uni-view",{staticClass:"left"},[n("v-uni-image",{attrs:{src:a.idAvatar,mode:"aspectFill"}})],1),n("v-uni-view",{staticClass:"right",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.addComment(a.replyId,e.id,i,a.replyName)}}},[n("v-uni-view",{staticClass:"name"},[t._v(t._s(a.createByName))]),a.replyName?n("v-uni-text",{staticClass:"reply-name"},[t._v("@"+t._s(a.replyName))]):t._e(),n("v-uni-text",{staticClass:"kid-content"},[t._v(t._s(a.content))])],1)],1)],1)})),t._l(t.kidsCommentCount,(function(e,i){return n("v-uni-view",[e>=3?n("v-uni-view",{staticClass:"all-reply",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.toAllReply.apply(void 0,arguments)}}},[t._v("展开更多回复")]):t._e()],1)}))],2)],1)],1)})),1)},o=[]},"7f9d":function(t,e,n){var i=n("0187");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("ea6714f4",i,!0,{sourceMap:!1,shadowMode:!1})},8140:function(t,e,n){"use strict";n.r(e);var i=n("a7f0"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},8472:function(t,e,n){"use strict";var i=n("8c2a"),a=n.n(i);a.a},"8b2c":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-title",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.articleDetails.apply(void 0,arguments)}}},[n("v-uni-view",[n("v-uni-text",{staticClass:"article-text",domProps:{innerHTML:t._s(t.articleTitle)}})],1)],1)},o=[]},"8c2a":function(t,e,n){var i=n("e104");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("5fc0f25a",i,!0,{sourceMap:!1,shadowMode:!1})},"9c76":function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var i={uIcon:n("ac18").default,uniLoadMore:n("d508").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-details"},[n("v-uni-view",{staticClass:"view-content"},[n("v-uni-view",{staticClass:"article-title"},[t._v(t._s(t.articleInfo.title))]),n("attentionAndFansCell",{staticClass:"user-info",attrs:{aid:t.articleInfo.create_by,nickname:t.articleInfo.nickname,avatarPath:t.articleInfo.avatarPath,isFollow:t.articleInfo.isFollow}}),n("v-uni-view",{staticClass:"article-message"},[n("v-uni-text",{staticClass:"word-num"},[n("v-uni-text",[t._v("字数")]),n("v-uni-text",{staticClass:"word-num-num"},[t._v(t._s(t.articleInfo.word_count))])],1),n("v-uni-text",{staticClass:"time"},[n("v-uni-text",[t._v(t._s(t.format(t.articleInfo.update_time)))])],1)],1),n("v-uni-view",{staticClass:"article-content"},[n("v-uni-text",{domProps:{innerHTML:t._s(t.articleInfo.content)}})],1)],1),n("v-uni-view",{staticClass:"comment-view"},[n("v-uni-view",{staticClass:"top-title"},[n("v-uni-text",[t._v("评论")]),n("v-uni-text",[t._v("("+t._s(t.commentCount)+")")])],1),n("v-uni-view",{staticClass:"comment-list"},[n("comment",{attrs:{kidsCommentCount:t.kidsCommentCount,commentList:t.commentList},on:{childFn:function(e){arguments[0]=e=t.$handleEvent(e),t.comment.apply(void 0,arguments)}}})],1)],1),n("v-uni-view",{staticClass:"bottom-nav"},[n("v-uni-view",{staticClass:"bottom-nav-centre"},[n("v-uni-view",{staticClass:"attention-icon comment"},[n("v-uni-view",{staticClass:"attention-item",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.comment()}}},[n("u-icon",{attrs:{name:"pinglun","custom-prefix":"custom-icon"}})],1)],1),n("v-uni-view",{staticClass:"attention-icon like",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.like()}}},[n("v-uni-view",{staticClass:"attention-item"},[n("u-icon",{attrs:{name:t.likeIcon,color:t.likeColor}}),n("v-uni-text",{staticClass:"attention-num"},[t._v(t._s(t.articleInfo.favourTotal))])],1)],1),n("v-uni-view",{staticClass:"attention-icon collect",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.collent()}}},[n("v-uni-view",{staticClass:"attention-item"},[n("u-icon",{attrs:{name:t.collentIcon,"custom-prefix":"custom-icon"}}),n("v-uni-text",{staticClass:"attention-num"},[t._v(t._s(t.articleInfo.collectTotal))])],1)],1),n("v-uni-view",{staticClass:"attention-icon",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.transpond()}}},[n("u-icon",{attrs:{name:"zhuanfa"}})],1)],1)],1),t.isComment?n("commentInput",{attrs:{show:t.showAddComment,type:t.type,addCommentArg:t.addCommentArg},on:{addComment:function(e){arguments[0]=e=t.$handleEvent(e),t.addComment.apply(void 0,arguments)},addChildComment:function(e){arguments[0]=e=t.$handleEvent(e),t.addChildComment.apply(void 0,arguments)},childFn:function(e){arguments[0]=e=t.$handleEvent(e),t.parentFn.apply(void 0,arguments)}}}):t._e(),n("v-uni-view",{directives:[{name:"show",rawName:"v-show",value:t.isLoadMore,expression:"isLoadMore"}]},[n("uni-load-more",{staticClass:"loading",attrs:{status:t.loadStatus,iconType:"circle"}})],1)],1)},o=[]},a0cf:function(t,e,n){"use strict";(function(t){n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i=n("1874"),a=n("ee60"),o=n("6782"),r={data:function(){return{sendDisabled:!0,comment:null}},props:{type:{type:Number,default:null},addCommentArg:{type:Object},show:{type:Boolean}},methods:{onkeyinput:function(t){this.comment=t.target.value,this.sendDisabled=!1},close:function(){this.$emit("childFn")},sendComment:function(){var e=this,n=this,r={id:this.addCommentArg.id,content:this.comment,replyId:this.addCommentArg.replyId,followId:this.addCommentArg.follewId};if(1===this.type)(0,i.publishIdea)(r).then((function(t){return e.close(),t.code})).then((function(t){var i;if((0,o.information)().then((function(t){i=t.data})),e.addCommentArg.childComment){var a={createByName:i.nickname,idAvatar:i.avatarPath,content:e.comment,createTime:new Date,replyName:e.addCommentArg.replyName,index:e.addCommentArg.index,commentKids:[]};n.$emit("addChildComment",a)}else{var r={createByName:i.nickname,idAvatar:i.avatarPath,content:e.comment,createTime:new Date,commentKids:[]};e.$emit("addComment",r)}}));else if(0===this.type){var c=this;(0,a.addArticleComment)(r).then((function(t){return e.close(),t.code})).then((function(e){(0,o.information)().then((function(t){return t.data})).then((function(e){c.addComment(e),t("log",127127," at marchVoiceComponents/comment/commentInput.vue:127")}))}))}},addComment:function(e){var n=this;if(this.addCommentArg.childComment){var i={createByName:user.nickname,idAvatar:user.avatarPath,content:this.comment,createTime:new Date,replyName:this.addCommentArg.replyName,index:this.addCommentArg.index,commentKids:[]};t("log",i," at marchVoiceComponents/comment/commentInput.vue:164"),this.$emit("addChildComment",i)}else{t("log",131," at marchVoiceComponents/comment/commentInput.vue:141");var a={createByName:e.nickname,idAvatar:e.avatarPath,content:this.comment,createTime:new Date,commentKids:[]};t("log",a," at marchVoiceComponents/comment/commentInput.vue:150"),n.$emit("addComment",a),t("log",152," at marchVoiceComponents/comment/commentInput.vue:152")}}}};e.default=r}).call(this,n("0de9")["log"])},a7f0:function(t,e,n){"use strict";(function(t){var i=n("4ea4");n("99af"),n("4160"),n("a9e3"),n("159b"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("b85c")),o=i(n("2909")),r=i(n("d50f")),c=i(n("40e6")),s=i(n("6edb")),l=i(n("f4e3")),d=i(n("1e61")),u=i(n("2184")),f=n("ee60"),m={data:function(){return{isLoadMore:!1,loadStatus:"loading",showAddComment:!0,id:0,current:1,size:5,childSize:3,articleInfo:{},commentCount:0,commentList:[],isComment:!1,type:0,addCommentArg:{},clickChild:!1,likeIcon:"heart",collentIcon:"shoucang1",likeColor:"black",clickLike:!1,clickCollect:!1,kidsCommentCount:[]}},components:{articleTitle:r.default,attentionAndFansCell:c.default,jinEdit:s.default,comment:l.default,commentInput:u.default},onLoad:function(t){this.id=Number(t.id),this.getArtile(),this.getComments()},created:function(){this.id;this.addCommentArg={id:this.id,replyId:0,follewId:0,childComment:!1}},methods:{like:function(){if(this.clickLike)this.articleInfo.favourTotal--,this.likeIcon="heart",this.likeColor="#080100";else{var t=this.id;this.likeIcon="heart-fill",this.likeColor="#d81e06",this.articleInfo.favourTotal++,(0,f.favour)(t).then((function(t){}))}this.clickLike=!this.clickLike},collent:function(){if(this.clickCollect)this.collentIcon="shoucang1",this.articleInfo.collectTotal--;else{var t=this.id;this.collentIcon="shoucang2",this.articleInfo.collectTotal++,(0,f.collect)(t).then((function(t){}))}this.clickCollect=!this.clickCollect},transpond:function(){var t=this.id;(0,f.reprint)(t).then((function(t){})),uni.showLoading({title:"加载中"}),setTimeout((function(){uni.hideLoading()}),2e3),setTimeout((function(){uni.showToast({title:"转发成功",duration:2e3})}),2e3)},comment:function(t){this.id;this.isComment=!0,t&&(this.addCommentArg={childComment:t.childComment,id:this.id,index:t.index,replyId:t.replyId,follewId:t.follewId,replyName:t.replyName})},parentFn:function(t){this.isComment=!1},addComment:function(e){t("log",212," at pages/articleDetails/index.vue:212"),t("log",e," at pages/articleDetails/index.vue:213"),this.commentList.unshift(e),this.isComment=!1,this.commentCount++},addChildComment:function(t){this.commentList[t.index].commentKids.push(t),this.isComment=!1,this.commentCount++},format:function(t){var e=new Date(t),n=(0,d.default)(e).format("YYYY-MM-DD HH:mm:ss");return n},getComments:function(){var t=this,e={id:this.id,current:this.current,size:this.size,childSize:this.childSize};(0,f.getArticleCommentList)(e).then((function(e){if(0===e.code){t.commentList=[].concat((0,o.default)(t.commentList),(0,o.default)(e.data.CommentSum));var n,i=t.commentList,r=(0,a.default)(i);try{for(r.s();!(n=r.n()).done;){var c=n.value;c.ChildComments?t.kidsCommentCount.push(c.ChildComments.length):t.kidsCommentCount.push(0)}}catch(s){r.e(s)}finally{r.f()}1===t.current?(t.isLoadMore=!1,t.commentList=[].concat((0,o.default)(t.commentList),(0,o.default)(e.data.CommentSum))):setTimeout((function(){this.isLoadMore=!1,this.commentList=[].concat((0,o.default)(this.commentList),(0,o.default)(e.data.CommentSum))}),3e3)}}))},getArtile:function(){var e=this,n=this.id;(0,f.getArtileDetails)(n).then((function(t){0===t.code&&(e.articleInfo=t.data)})).catch((function(e){t("log",e,"err login"," at pages/articleDetails/index.vue:277")}))}},mounted:function(){var t=this.articleInfo.commentTotal;this.commentList.forEach((function(e){t++,e.commentKids.forEach((function(e){t++}))})),this.commentCount=t},onReachBottom:function(){this.isLoadMore||(this.isLoadMore=!0,this.current+=1,this.getComments())},onPullDownRefresh:function(){this.current=1,this.getArtile(),this.getComments(),setTimeout((function(){uni.stopPullDownRefresh()}),1e3)}};e.default=m}).call(this,n("0de9")["log"])},b23f:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var i={uMask:n("c0e8").default,uIcon:n("ac18").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return t.visibleSync?n("v-uni-view",{staticClass:"u-drawer",style:[t.customStyle,{zIndex:t.uZindex-1}],attrs:{"hover-stop-propagation":!0}},[n("u-mask",{attrs:{duration:t.duration,"custom-style":t.maskCustomStyle,maskClickAble:t.maskCloseAble,"z-index":t.uZindex-2,show:t.showDrawer&&t.mask},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.maskClick.apply(void 0,arguments)}}}),n("v-uni-view",{staticClass:"u-drawer-content",class:[t.safeAreaInsetBottom?"safe-area-inset-bottom":"","u-drawer-"+t.mode,t.showDrawer?"u-drawer-content-visible":"",t.zoom&&"center"==t.mode?"u-animation-zoom":""],style:[t.style],on:{touchmove:function(e){e.stopPropagation(),e.preventDefault(),arguments[0]=e=t.$handleEvent(e)},click:[function(e){arguments[0]=e=t.$handleEvent(e),t.modeCenterClose(t.mode)},function(e){e.stopPropagation(),e.preventDefault(),arguments[0]=e=t.$handleEvent(e)}]}},["center"==t.mode?n("v-uni-view",{staticClass:"u-mode-center-box",style:[t.centerStyle],on:{touchmove:function(e){e.stopPropagation(),e.preventDefault(),arguments[0]=e=t.$handleEvent(e)},click:function(e){e.stopPropagation(),e.preventDefault(),arguments[0]=e=t.$handleEvent(e)}}},[t.closeable?n("u-icon",{staticClass:"u-close",class:["u-close--"+t.closeIconPos],attrs:{name:t.closeIcon,color:t.closeIconColor,size:t.closeIconSize},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.close.apply(void 0,arguments)}}}):t._e(),n("v-uni-scroll-view",{staticClass:"u-drawer__scroll-view",attrs:{"scroll-y":"true"}},[t._t("default")],2)],1):n("v-uni-scroll-view",{staticClass:"u-drawer__scroll-view",attrs:{"scroll-y":"true"}},[t._t("default")],2),n("v-uni-view",{staticClass:"u-close",class:["u-close--"+t.closeIconPos],on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.close.apply(void 0,arguments)}}},["center"!=t.mode&&t.closeable?n("u-icon",{attrs:{name:t.closeIcon,color:t.closeIconColor,size:t.closeIconSize}}):t._e()],1)],1)],1):t._e()},o=[]},bea5:function(t,e,n){var i=n("4788");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("332ec73e",i,!0,{sourceMap:!1,shadowMode:!1})},c0e8:function(t,e,n){"use strict";n.r(e);var i=n("18f1"),a=n("15ed");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("1c0ba");var r,c=n("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"ff9dd99c",null,!1,i["a"],r);e["default"]=s.exports},c8db:function(t,e,n){var i=n("3e8a");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("6a9888a6",i,!0,{sourceMap:!1,shadowMode:!1})},c905:function(t,e,n){"use strict";var i=n("bea5"),a=n.n(i);a.a},c9ba:function(t,e,n){var i=n("d1c5");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("28bd3649",i,!0,{sourceMap:!1,shadowMode:!1})},cf82:function(t,e,n){var i=n("4c6b");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("3ec8a329",i,!0,{sourceMap:!1,shadowMode:!1})},d1c5:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.u-mask[data-v-ff9dd99c]{position:fixed;top:0;left:0;right:0;bottom:0;opacity:0;-webkit-transition:-webkit-transform .3s;transition:-webkit-transform .3s;transition:transform .3s;transition:transform .3s,-webkit-transform .3s}.u-mask-show[data-v-ff9dd99c]{opacity:1}.u-mask-zoom[data-v-ff9dd99c]{-webkit-transform:scale(1.2);transform:scale(1.2)}',""]),t.exports=e},d50f:function(t,e,n){"use strict";n.r(e);var i=n("8b2c"),a=n("f708");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("c905");var r,c=n("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"74809812",null,!1,i["a"],r);e["default"]=s.exports},db90:function(t,e,n){"use strict";function i(t){if("undefined"!==typeof Symbol&&Symbol.iterator in Object(t))return Array.from(t)}n("a4d3"),n("e01a"),n("d28b"),n("a630"),n("d3b7"),n("3ca3"),n("ddb0"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=i},dc97:function(t,e,n){"use strict";n.r(e);var i=n("a0cf"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},e104:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.comment[data-v-dd03f188]:last-child{border:none}.name[data-v-dd03f188]{color:grey;font-size:%?26?%;line-height:150%}.comment[data-v-dd03f188]{display:-webkit-box;display:-webkit-flex;display:flex;padding:%?30?% 0;border-bottom:%?1?% solid #f0f0f0}.comment .left uni-image[data-v-dd03f188]{width:%?70?%;height:%?70?%;border-radius:50%}.comment .right[data-v-dd03f188]{-webkit-box-flex:1;-webkit-flex:1;flex:1;padding-left:%?20?%;font-size:%?30?%}.comment .right .top[data-v-dd03f188]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.comment .right .top .like[data-v-dd03f188]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;color:#9a9a9a;font-size:%?26?%}.comment .right .top .highlight[data-v-dd03f188]{color:#5677fc}.comment .right .content[data-v-dd03f188],\r\n.comment .right .kid-content[data-v-dd03f188]{margin-bottom:%?10?%;color:#505050;font-size:%?28?%;line-height:150%;text-align:left}.comment .right .reply-box .item[data-v-dd03f188]{display:-webkit-box;display:-webkit-flex;display:flex;padding:%?20?% %?20?% 0 %?20?%}.comment .right .reply-box .all-reply[data-v-dd03f188]{padding:%?20?%;color:#2a82e4;font-size:%?28?%;line-height:150%;text-align:left;font-weight:700}.comment .right .reply-box .all-reply .more[data-v-dd03f188]{margin-left:%?6?%}.comment .right .bottom[data-v-dd03f188]{margin-top:%?20?%;display:-webkit-box;display:-webkit-flex;display:flex;font-size:%?24?%;color:#9a9a9a}.comment .right .bottom .reply[data-v-dd03f188]{color:#5677fc;margin-left:%?10?%}.reply-name[data-v-dd03f188]{color:#2a82e4;font-size:%?28?%;line-height:150%}.item .left[data-v-dd03f188]{display:-webkit-box;display:-webkit-flex;display:flex}.item .left uni-image[data-v-dd03f188]{-webkit-box-flex:1;-webkit-flex:1;flex:1;width:%?50?%;height:%?50?%;line-height:%?80?%}',""]),t.exports=e},ee60:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.getArtileDetails=o,e.getArticleCommentList=r,e.addArticleComment=c,e.getChildCommentList=s,e.favour=l,e.collect=d,e.reprint=u;var a=i(n("367e"));function o(t){return a.default.get("/api/article/detail/"+t)}function r(t){return a.default.get("/api/comment/article",{params:t})}function c(t){return a.default.post("/api/comment/article",{id:t.id,content:t.content,replyId:t.replyId,follewId:t.follewId})}function s(t){return a.default.get("/api/comment/children/article",{params:t})}function l(t){return a.default.post("/api/favour/article",{id:t})}function d(t){return a.default.post("/api/collect/article",{id:t})}function u(t){return a.default.post("/api/article/reprint",{id:t})}},f4e3:function(t,e,n){"use strict";n.r(e);var i=n("7a94"),a=n("564f");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("8472");var r,c=n("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"dd03f188",null,!1,i["a"],r);e["default"]=s.exports},f708:function(t,e,n){"use strict";n.r(e);var i=n("349b"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a}}]);