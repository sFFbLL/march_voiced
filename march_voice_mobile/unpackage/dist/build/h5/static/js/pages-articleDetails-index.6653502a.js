(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-articleDetails-index"],{"079b":function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={data:function(){return{}},props:{commentList:{type:Array,default:[]}},methods:{toAllReply:function(){console.log("查看更多")}}};e.default=i},1146:function(t,e,n){"use strict";var i=n("29fa"),a=n.n(i);a.a},"131c":function(t,e,n){var i=n("ee6f");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("d2e6c738",i,!0,{sourceMap:!1,shadowMode:!1})},1874:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.getMarchCircleInfo=o,e.marchCircleList=r,e.publishIdea=l,e.joinMarchCircle=c,e.changeFavour=s,e.ideaDetail=d,e.publishComment=u,e.ideaCommentList=f,e.ideaChildCommentList=v;var a=i(n("367e"));function o(t){return a.default.get("/api/march/msg",{params:t})}function r(t){return a.default.get("/api/march",{params:t})}function l(t){return a.default.post("/api/march",{content:t.content,imageList:t.imageList,image:t.image})}function c(t){return a.default.post("/api/apply/march",{params:t})}function s(t){return a.default.post("/api/favour/march",{id:t.id,type:t.type})}function d(t){return a.default.get("/api/march/detail/id",{id:t.id})}function u(t){return a.default.post("/api/march/comment",{params:t})}function f(t){return a.default.post("/api/comment/march",{id:t.id,current:t.current,size:t.size,childSize:t.childSize})}function v(t){return a.default.get("/api/comment/children/article",{id:t.id,current:t.current,size:t.size})}},2071:function(t,e,n){"use strict";var i=n("4ea4");n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("2208")),o={props:{id:{type:Number,default:null},nickname:{type:String,default:"昵称"},avatarPath:{type:String,default:null},isFollow:{type:Number,default:null}},data:function(){return{isDisabled:!1,openid:null}},methods:{changeBtn:function(){this.isDisabled=!0;var t=this;setTimeout((function(){t.isDisabled=!1,t.$emit("change")}),1e3)},inToMine:function(){this.$emit("inToPageMine")},getImgUrl:function(t){return a.default.imgUrl+t}}};e.default=o},2184:function(t,e,n){"use strict";n.r(e);var i=n("e1fc"),a=n("dc97");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("9ff3");var r,l=n("f0c5"),c=Object(l["a"])(a["default"],i["b"],i["c"],!1,null,"3722ba4d",null,!1,i["a"],r);e["default"]=c.exports},2739:function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",t._l(t.commentList,(function(e,i){return n("v-uni-view",{key:e.id,staticClass:"comment"},[n("v-uni-view",{staticClass:"left"},[n("v-uni-image",{attrs:{src:e.idAvatar,mode:"aspectFill"}})],1),n("v-uni-view",{staticClass:"right"},[n("v-uni-view",{staticClass:"top"},[n("v-uni-view",{staticClass:"name"},[t._v(t._s(e.createByName))])],1),n("v-uni-view",{staticClass:"content"},[t._v(t._s(e.content))]),n("v-uni-view",{staticClass:"bottom"},[t._v(t._s(e.createTime))]),n("v-uni-view",{staticClass:"reply-box"},[t._l(e.commentKids,(function(e,i){return n("v-uni-view",{key:e.index},[i<3?n("v-uni-view",{staticClass:"item"},[n("v-uni-view",{staticClass:"left"},[n("v-uni-image",{attrs:{src:e.idAvatar,mode:"aspectFill"}})],1),n("v-uni-view",{staticClass:"right"},[n("v-uni-view",{staticClass:"name"},[t._v(t._s(e.createByName))]),e.replyName?n("v-uni-text",{staticClass:"reply-name"},[t._v("@"+t._s(e.replyName))]):t._e(),n("v-uni-text",{staticClass:"kid-content"},[t._v(t._s(e.content))])],1)],1):t._e()],1)})),e.commentKids.length>3?n("v-uni-view",{staticClass:"all-reply",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.toAllReply.apply(void 0,arguments)}}},[t._v("展开更多回复")]):t._e()],2)],1)],1)})),1)},o=[]},"29fa":function(t,e,n){var i=n("b191");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("e821e2e2",i,!0,{sourceMap:!1,shadowMode:!1})},"349b":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={data:function(){return{}},props:{articleId:{type:Number,default:null},articleTitle:{type:String,default:""}},components:{},methods:{articleDetails:function(){uni.navigateTo({url:"../articleDetails/index?id="+this.articleId})}},mounted:function(){}};e.default=i},"40e6":function(t,e,n){"use strict";n.r(e);var i=n("f142"),a=n("bd53");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("7af7");var r,l=n("f0c5"),c=Object(l["a"])(a["default"],i["b"],i["c"],!1,null,"7be5294c",null,!1,i["a"],r);e["default"]=c.exports},4788:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".article-title .article-text[data-v-74809812]{color:#404040;font-size:%?34?%;font-family:微软雅黑;text-align:left;font-weight:700;overflow:hidden;-webkit-line-clamp:2;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;line-height:150%}.article-text[data-v-74809812]{display:inline-block;width:100%}",""]),t.exports=e},5333:function(t,e,n){"use strict";var i=n("6d84"),a=n.n(i);a.a},"564f":function(t,e,n){"use strict";n.r(e);var i=n("079b"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},"6b8b":function(t,e,n){var i=n("cfa8");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("11f328bb",i,!0,{sourceMap:!1,shadowMode:!1})},"6d84":function(t,e,n){var i=n("aa47");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("0625dc16",i,!0,{sourceMap:!1,shadowMode:!1})},"784d":function(t,e,n){"use strict";n.r(e);var i=n("ae2b"),a=n("8140");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("5333");var r,l=n("f0c5"),c=Object(l["a"])(a["default"],i["b"],i["c"],!1,null,"d2456c20",null,!1,i["a"],r);e["default"]=c.exports},"7af7":function(t,e,n){"use strict";var i=n("131c"),a=n.n(i);a.a},8140:function(t,e,n){"use strict";n.r(e);var i=n("a7f0"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},"8b2c":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-title",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.articleDetails.apply(void 0,arguments)}}},[n("v-uni-view",[n("v-uni-text",{staticClass:"article-text",domProps:{innerHTML:t._s(t.articleTitle)}})],1)],1)},o=[]},"9ff3":function(t,e,n){"use strict";var i=n("6b8b"),a=n.n(i);a.a},a0cf:function(t,e,n){"use strict";var i=n("4ea4");n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;i(n("1874")),i(n("ee60"));var a={data:function(){return{sendDisabled:!0,comment:null}},props:{type:{type:Number,default:null}},methods:{onkeyinput:function(t){this.comment=t.target.value,this.sendDisabled=!1},sendComment:function(){}}};e.default=a},a7f0:function(t,e,n){"use strict";var i=n("4ea4");n("4160"),n("159b"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("d50f")),o=i(n("40e6")),r=i(n("6edb")),l=i(n("f4e3")),c=i(n("2184")),s=n("ee60"),d={data:function(){return{id:"",current:1,size:5,childSize:3,articleInfo:{},commentCount:0,commentList:[],isComment:!1,type:0}},components:{articleTitle:a.default,attentionAndFansCell:o.default,jinEdit:r.default,comment:l.default,commentInput:c.default},onLoad:function(t){var e=this;this.id=t,(0,s.getArtileDetails)(t).then((function(t){0===t.code&&(e.articleInfo=t.data)})).catch((function(t){console.log(t,"err login")}));var n={id:t,current:this.current,size:this.size,childSize:this.childSize};(0,s.getArticleCommentList)(n).then((function(t){0===t.code&&(e.commentList=t.data,e.commentCount=t.data.length)}))},created:function(){},methods:{comment:function(){this.isComment=!0},sendComment:function(){console.log("发布评论"),this.isComment=!1}},mounted:function(){var t=0;this.commentList.forEach((function(e){t++,e.commentKids.forEach((function(e){t++}))})),this.commentCount=t}};e.default=d},aa47:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'.article-details[data-v-d2456c20]{background-color:#f7f7f7}.view-content[data-v-d2456c20]{padding:0 %?30?%;background-color:#fff}.article-title[data-v-d2456c20]{color:#404040;font-size:%?36?%;line-height:150%;font-weight:700}[data-v-d2456c20] .attention-cell .flex-item{padding:%?20?% 0;border:none}[data-v-d2456c20] .attention-cell .left-img-box{margin-right:0}[data-v-d2456c20] .attention-cell .right-button{background-color:#fff}.article-message[data-v-d2456c20]{color:#999797;font-size:%?24?%}.word-num-num[data-v-d2456c20]{margin-left:%?4?%}.time[data-v-d2456c20]{float:right}.article-content[data-v-d2456c20]{padding:0 %?10?%}[data-v-d2456c20] .article-content .ql-container{margin:0;width:100%;color:#404040;font-size:%?32?%;line-height:150%}[data-v-d2456c20] .article-content .container{padding-top:%?20?%}[data-v-d2456c20] .article-content .ql-container .ql-editor{text-align:justify}[data-v-d2456c20] .article-content .tool-view{display:none}\n\n/* 底部导航:点赞，评论等 */.bottom-nav[data-v-d2456c20]{position:fixed;bottom:0;width:100%;height:%?120?%;border:%?1?% solid #f0f0f0;border-width:%?1?% 0;background-color:#fff}.bottom-nav-centre[data-v-d2456c20]{display:-webkit-box;display:-webkit-flex;display:flex;margin:0 %?50?%;height:100%}.attention-icon[data-v-d2456c20]::before{content:"";display:inline-block;vertical-align:middle;height:100%}.attention-icon[data-v-d2456c20]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center}[data-v-d2456c20] .attention-icon .u-icon{font-size:%?50?%!important;color:#404040!important}.attention-item[data-v-d2456c20]{position:relative;display:inline-block}.attention-icon .attention-num[data-v-d2456c20]{position:absolute;top:%?-8?%}\n\n/* 评论样式 */.top-title[data-v-d2456c20]{height:%?88?%;color:#505050;font-size:%?34?%;line-height:150%;font-weight:700}.top-title uni-text[data-v-d2456c20]:first-child{display:inline-block;margin:0 %?10?% 0 %?30?%;line-height:%?88?%}.comment-view[data-v-d2456c20]{margin-bottom:%?120?%}.comment-list[data-v-d2456c20]{padding:0 %?30?%;background-color:#fff}',""]),t.exports=e},ae2b:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var i={uIcon:n("ac18").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"article-details"},[n("v-uni-view",{staticClass:"view-content"},[n("v-uni-view",{staticClass:"article-title"},[t._v(t._s(t.articleInfo.title))]),n("attentionAndFansCell",{staticClass:"user-info",attrs:{nickname:t.articleInfo.nickname,avatarPath:t.articleInfo.avatarPath,isFollow:t.articleInfo.isFollow}}),n("v-uni-view",{staticClass:"article-message"},[n("v-uni-text",{staticClass:"word-num"},[n("v-uni-text",[t._v("字数")]),n("v-uni-text",{staticClass:"word-num-num"},[t._v(t._s(t.articleInfo.word_count))])],1),n("v-uni-text",{staticClass:"time"},[n("v-uni-text",[t._v(t._s(t.articleInfo.update_time))])],1)],1),n("v-uni-view",{staticClass:"article-content"},[n("jinEdit",{attrs:{html:t.articleInfo.content,readOnly:!0}})],1)],1),n("v-uni-view",{staticClass:"comment-view"},[n("v-uni-view",{staticClass:"top-title"},[n("v-uni-text",[t._v("评论")]),n("v-uni-text",[t._v("("+t._s(t.commentCount)+")")])],1),n("v-uni-view",{staticClass:"comment-list"},[n("comment",{attrs:{commentList:t.commentList}})],1)],1),t.isComment?t._e():n("v-uni-view",{staticClass:"bottom-nav"},[n("v-uni-view",{staticClass:"bottom-nav-centre"},[n("v-uni-view",{staticClass:"attention-icon comment"},[n("v-uni-view",{staticClass:"attention-item",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.comment()}}},[n("u-icon",{attrs:{name:"pinglun","custom-prefix":"custom-icon"}}),n("v-uni-text",{staticClass:"attention-num"},[t._v(t._s(t.articleInfo.commentTotal))])],1)],1),n("v-uni-view",{staticClass:"attention-icon like"},[n("v-uni-view",{staticClass:"attention-item"},[n("u-icon",{attrs:{name:"heart"}}),n("v-uni-text",{staticClass:"attention-num"},[t._v(t._s(t.articleInfo.favourTotal))])],1)],1),n("v-uni-view",{staticClass:"attention-icon collect"},[n("v-uni-view",{staticClass:"attention-item"},[n("u-icon",{attrs:{name:"shoucang1","custom-prefix":"custom-icon"}}),n("v-uni-text",{staticClass:"attention-num"},[t._v(t._s(t.articleInfo.collectTotal))])],1)],1),n("v-uni-view",{staticClass:"attention-icon"},[n("u-icon",{attrs:{name:"zhuanfa"}})],1)],1)],1),t.isComment?n("commentInput",{attrs:{type:t.type},on:{sendComment:function(e){arguments[0]=e=t.$handleEvent(e),t.sendComment.apply(void 0,arguments)}}}):t._e()],1)},o=[]},b191:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.comment[data-v-13b71f4a]:last-child{border:none}.name[data-v-13b71f4a]{color:grey;font-size:%?26?%;line-height:150%}.comment[data-v-13b71f4a]{display:-webkit-box;display:-webkit-flex;display:flex;padding:%?30?% 0;border-bottom:%?1?% solid #f0f0f0}.comment .left uni-image[data-v-13b71f4a]{width:%?70?%;height:%?70?%;border-radius:50%}.comment .right[data-v-13b71f4a]{-webkit-box-flex:1;-webkit-flex:1;flex:1;padding-left:%?20?%;font-size:%?30?%}.comment .right .top[data-v-13b71f4a]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.comment .right .top .like[data-v-13b71f4a]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;color:#9a9a9a;font-size:%?26?%}.comment .right .top .highlight[data-v-13b71f4a]{color:#5677fc}.comment .right .content[data-v-13b71f4a],\r\n.comment .right .kid-content[data-v-13b71f4a]{margin-bottom:%?10?%;color:#505050;font-size:%?28?%;line-height:150%;text-align:left}.comment .right .reply-box .item[data-v-13b71f4a]{display:-webkit-box;display:-webkit-flex;display:flex;padding:%?20?% %?20?% 0 %?20?%}.comment .right .reply-box .all-reply[data-v-13b71f4a]{padding:%?20?%;color:#2a82e4;font-size:%?28?%;line-height:150%;text-align:left;font-weight:700}.comment .right .reply-box .all-reply .more[data-v-13b71f4a]{margin-left:%?6?%}.comment .right .bottom[data-v-13b71f4a]{margin-top:%?20?%;display:-webkit-box;display:-webkit-flex;display:flex;font-size:%?24?%;color:#9a9a9a}.comment .right .bottom .reply[data-v-13b71f4a]{color:#5677fc;margin-left:%?10?%}.reply-name[data-v-13b71f4a]{color:#2a82e4;font-size:%?28?%;line-height:150%}.item .left[data-v-13b71f4a]{display:-webkit-box;display:-webkit-flex;display:flex}.item .left uni-image[data-v-13b71f4a]{-webkit-box-flex:1;-webkit-flex:1;flex:1;width:%?50?%;height:%?50?%;line-height:%?80?%}',""]),t.exports=e},bd53:function(t,e,n){"use strict";n.r(e);var i=n("2071"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},bea5:function(t,e,n){var i=n("4788");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("332ec73e",i,!0,{sourceMap:!1,shadowMode:!1})},c905:function(t,e,n){"use strict";var i=n("bea5"),a=n.n(i);a.a},cfa8:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".commentwarp[data-v-3722ba4d]{background-color:#fff}\n\n/* 评论 */.commentinput[data-v-3722ba4d]{line-height:%?42?%}.sendbtn[data-v-3722ba4d]{height:%?68?%;width:%?140?%;background-color:#fff;color:#007aff;font-size:%?28?%}[data-v-3722ba4d] .sendbtn::after{border:none}.comment-view[data-v-3722ba4d]{max-height:%?126?%;overflow:scroll}.commentwarp[data-v-3722ba4d]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;\n\t/*定义body的元素垂直居中*/width:100%;font-size:%?28?%;padding:%?15?% %?28?%;text-align:left;position:fixed;bottom:0;border-top:%?1?% solid #d1d6de;color:#303133}",""]),t.exports=e},d50f:function(t,e,n){"use strict";n.r(e);var i=n("8b2c"),a=n("f708");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("c905");var r,l=n("f0c5"),c=Object(l["a"])(a["default"],i["b"],i["c"],!1,null,"74809812",null,!1,i["a"],r);e["default"]=c.exports},dc97:function(t,e,n){"use strict";n.r(e);var i=n("a0cf"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},e1fc:function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"commentwarp"},[n("v-uni-view",{staticClass:"comment-view"},[n("v-uni-textarea",{staticClass:"commentinput",attrs:{"auto-height":!0,maxlength:"-1",placeholder:"请输入评论"},on:{input:function(e){arguments[0]=e=t.$handleEvent(e),t.onkeyinput.apply(void 0,arguments)}}})],1),n("v-uni-button",{staticClass:"sendbtn",attrs:{size:"mini",disabled:t.sendDisabled},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.sendComment()}}},[t._v("发布")])],1)},o=[]},ee60:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.getArtileDetails=o,e.getArticleCommentList=r,e.addArticleComment=l,e.getChildCommentList=c,e.favour=s,e.collect=d,e.reprint=u;var a=i(n("367e"));function o(t){return a.default.get("/api/article/detail",{params:t})}function r(t){return a.default.get("/api/comment/article",{id:t.id,current:t.current,size:t.size,childSize:t.childSize})}function l(t){return a.default.post("/api/comment/article",{id:t.id,content:t.content,replyId:t.replyId,follewId:t.follewId})}function c(t){return a.default.get("/api/comment/children/article",{params:t})}function s(t){return a.default.post("/api/favour/article",{params:t})}function d(t){return a.default.post("/api/collect/article",{params:t})}function u(t){return a.default.post("/api/article/reprint",{params:t})}},ee6f:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\t/* 大盒子样式 */.flex-item[data-v-7be5294c]{display:-webkit-box;display:-webkit-flex;display:flex;box-sizing:border-box;border-bottom:%?6?% solid #f0f0f0;padding:%?20?% %?30?% %?20?% %?30?%\n\t\t/* margin-bottom: 20rpx; */}\n\t/* 左侧图片盒子样式 */.left-img-box[data-v-7be5294c]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:1;-webkit-flex:1;flex:1;margin-right:%?10?%;\n\t\t/* width: 112rpx; */\n\t\t/* height: 112rpx; */vertical-align:middle}\n\t/* 内部盒子样式 */\n\t/* .inner-box {\n\t\tdisplay: inline-block;\n\t} */\n\t/* .inner-box {\n    display: inline-block;\n  } */\n\t/* 图片样式 */.inner-img[data-v-7be5294c]{height:%?70?%;width:%?70?%;border-radius:50%;box-shadow:0 0 1px 1px rgba(0,0,0,.2)}\n\t/* 中部文字盒子样式 */.middle-text-box[data-v-7be5294c]{-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:6;-webkit-flex:6;flex:6;display:-webkit-box;display:-webkit-flex;display:flex}\n\t/* 灰色字体 */.greay-text[data-v-7be5294c]{color:#969696}\n\t/* 昵称 */.inner-text-name[data-v-7be5294c]{font-size:%?30?%;font-weight:700;vertical-align:middle}\n\t/* 关注数粉丝数文章数 */.inner-text-message[data-v-7be5294c]{padding-top:%?4?%;font-size:%?24?%}\n\t/* 自我介绍 */.inner-text-selfintroduce[data-v-7be5294c]{padding-top:%?4?%;font-size:%?24?%;width:%?420?%;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}\n\t/* 右侧按钮盒子 */.right-button-box[data-v-7be5294c]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-flex:1;-webkit-flex:1;flex:1;-webkit-box-align:center;-webkit-align-items:center;align-items:center}\n\t/* 右侧按钮 */.right-button[data-v-7be5294c]{border-radius:%?58?%;width:%?130?%;height:%?56?%;font-size:%?27?%;font-weight:400;line-height:%?56?%;color:#969696}\n\t/* 修改按钮样式 */uni-button[type=default][data-v-7be5294c]{border:%?1?% #969696 solid;box-sizing:border-box}\n\t/* 修改按钮样式 */uni-button[type=default-green][data-v-7be5294c]{color:#fff;background-color:#404040}\n\t/* 修改点击按钮后的样式 */.button-hover[type=default-green][data-v-7be5294c]{background-color:#1a1a1a}\n\t/* 修改按钮样式 */uni-button[data-v-7be5294c]{margin:0;padding:0}",""]),t.exports=e},f142:function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"attention-cell"},[n("v-uni-view",{staticClass:"flex-item"},[n("a",{staticClass:"left-img-box inner-box",on:{click:function(e){e.stopPropagation(),arguments[0]=e=t.$handleEvent(e),t.inToMine.apply(void 0,arguments)}}},[n("v-uni-image",{staticClass:"inner-img",attrs:{src:t.getImgUrl(t.avatarPath),mode:"aspectFill"}})],1),n("v-uni-view",{staticClass:"middle-text-box inner-box",on:{click:function(e){e.stopPropagation(),arguments[0]=e=t.$handleEvent(e),t.inToMine.apply(void 0,arguments)}}},[n("v-uni-view",{staticClass:"inner-middle-box inner-box"},[n("v-uni-view",{staticClass:"inner-text-name"},[t._v(t._s(t.nickname)),t._t("afterNicknameText")],2),n("v-uni-view",{staticClass:"inner-text-message greay-text"},[t._t("middleText")],2),n("v-uni-view",{staticClass:"inner-text-selfintroduce greay-text"},[t._t("underText")],2)],1)],1),n("v-uni-view",{staticClass:"right-button-box"},[null!=t.isFollow?n("v-uni-view",[n("v-uni-button",{directives:[{name:"show",rawName:"v-show",value:0===t.isFollow,expression:"isFollow === 0"}],staticClass:"right-button",attrs:{type:"default-green",disabled:t.isDisabled,iconType:"circle"},on:{click:function(e){e.stopPropagation(),arguments[0]=e=t.$handleEvent(e),t.changeBtn.apply(void 0,arguments)}}},[n("span",[t._v("关 注")])]),n("v-uni-button",{directives:[{name:"show",rawName:"v-show",value:1===t.isFollow,expression:"isFollow === 1"}],staticClass:"right-button",attrs:{type:"default",disabled:t.isDisabled,iconType:"circle"},on:{click:function(e){e.stopPropagation(),arguments[0]=e=t.$handleEvent(e),t.changeBtn.apply(void 0,arguments)}}},[n("span",[t._v("已关注")])])],1):t._e()],1)],1)],1)},o=[]},f4e3:function(t,e,n){"use strict";n.r(e);var i=n("2739"),a=n("564f");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("1146");var r,l=n("f0c5"),c=Object(l["a"])(a["default"],i["b"],i["c"],!1,null,"13b71f4a",null,!1,i["a"],r);e["default"]=c.exports},f708:function(t,e,n){"use strict";n.r(e);var i=n("349b"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a}}]);