(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-moreComments-moreComments"],{"079b":function(t,e,n){"use strict";var i=n("4ea4");n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var a=i(n("1e61")),o={data:function(){return{text:"展开更多回复"}},props:{id:{type:Number,default:null},commentList:{type:Array,default:[]},kidsCommentCount:{type:Array,default:[]},type:{type:Number,default:null},showKids:{type:Boolean,default:!0}},methods:{format:function(t){var e=new Date(t),n=(0,a.default)(e).format("YYYY-MM-DD HH:mm:ss");return n},toAllReply:function(t){this.text="",uni.navigateTo({url:"../../pages/moreComments/moreComments?commentList"+this.commentList[t]+"&id="+this.id})},addComment:function(t,e,n,i){var a={childComment:!0,index:n,replyId:t,follewId:e,replyName:i};this.$emit("childFn",a)}}};e.default=o},1845:function(t,e,n){var i=n("e4dc");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("30c747ff",i,!0,{sourceMap:!1,shadowMode:!1})},"3eaf":function(t,e,n){"use strict";n.r(e);var i=n("60cc"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},"564f":function(t,e,n){"use strict";n.r(e);var i=n("079b"),a=n.n(i);for(var o in i)"default"!==o&&function(t){n.d(e,t,(function(){return i[t]}))}(o);e["default"]=a.a},"60cc":function(t,e,n){"use strict";(function(t){var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;i(n("f4e3"));var a=n("ee60"),o={data:function(){return{type:3,commentList:[],showKids:!1,id:0,ChildCommentsList:[],artilceId:0,showAddComment:!0,isComment:!0}},onLoad:function(e){var n=JSON.stringify(e);t("log",n," at pages/moreComments/moreComments.vue:37"),this.artilceId=n.id,this.commentList=n.commentList,this.id=this.commentList.id,this.getMore(),this.addCommentArg={id:this.artilceId,replyId:0,follewId:0,childComment:!1}},methods:{getMore:function(){var t=this,e={id:this.id};(0,a.getChildCommentList)(e).then((function(e){t.ChildCommentsList=e.data}))},newComment:function(t){(0,a.addArticleComment)(t).then((function(t){}))},addComment:function(e){var n={createByName:getUserName(),idAvatar:getAvatarPath(),content:e,createTime:new Date,ChildComments:[]};this.commentList.unshift(n),this.isComment=!1,this.commentCount++;var i={id:this.addCommentArg.id,content:e,replyId:this.addCommentArg.replyId,followId:this.addCommentArg.follewId};t("log",i," at pages/moreComments/moreComments.vue:85"),this.newComment(i)}}};e.default=o}).call(this,n("0de9")["log"])},8583:function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",[n("v-uni-view",{staticClass:"comment-list"},[n("comment",{attrs:{showKids:t.showKids,type:t.type,commentList:t.commentList},on:{childFn:function(e){arguments[0]=e=t.$handleEvent(e),t.comment.apply(void 0,arguments)}}}),n("comment",{attrs:{showKids:t.showKids,type:t.type,commentList:t.ChildCommentsList},on:{childFn:function(e){arguments[0]=e=t.$handleEvent(e),t.comment.apply(void 0,arguments)}}}),t.isComment?n("commentInput",{attrs:{show:t.showAddComment,type:t.type,addCommentArg:t.addCommentArg},on:{addComment:function(e){arguments[0]=e=t.$handleEvent(e),t.addComment.apply(void 0,arguments)},addChildComment:function(e){arguments[0]=e=t.$handleEvent(e),t.addChildComment.apply(void 0,arguments)},childFn:function(e){arguments[0]=e=t.$handleEvent(e),t.parentFn.apply(void 0,arguments)}}}):t._e()],1)],1)},o=[]},"8f93":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"a",(function(){return i}));var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",t._l(t.commentList,(function(e,i){return n("v-uni-view",{staticClass:"comment"},[n("v-uni-view",{staticClass:"left"},[n("v-uni-image",{attrs:{src:e.idAvatar?e.idAvatar:null,mode:"aspectFill"}})],1),n("v-uni-view",{staticClass:"right"},[n("v-uni-view",{staticClass:"top"},[n("v-uni-view",{staticClass:"name"},[t._v(t._s(e.createByName))])],1),e.replyName?n("v-uni-text",{staticClass:"reply-name"},[t._v("@"+t._s(e.replyName))]):t._e(),n("v-uni-view",{staticClass:"content",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.addComment(e.createBy,e.id,i,e.createByName)}}},[t._v(t._s(e.content))]),n("v-uni-view",{staticClass:"bottom"},[t._v(t._s(t.format(e.createTime)))]),t.showKids?n("v-uni-view",{staticClass:"reply-box"},[t._l(e.ChildComments,(function(a,o){return n("v-uni-view",[n("v-uni-view",{staticClass:"item"},[n("v-uni-view",{staticClass:"left"},[n("v-uni-image",{attrs:{src:a.idAvatar,mode:"aspectFill"}})],1),n("v-uni-view",{staticClass:"right",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.addComment(a.replyId,e.id,i,a.replyName)}}},[n("v-uni-view",{staticClass:"name"},[t._v(t._s(a.createByName))]),n("v-uni-text",{staticClass:"reply-name"},[t._v("@"+t._s(a.replyName))]),n("v-uni-text",{staticClass:"kid-content"},[t._v(t._s(a.content))]),n("v-uni-text",{staticClass:"bottom"},[t._v(t._s(t.format(a.createTime)))])],1)],1)],1)})),e.ChildComments&&e.ChildComments.length>=3?n("v-uni-view",[n("v-uni-view",{staticClass:"all-reply",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.toAllReply(i)}}},[t._v(t._s(t.text))])],1):t._e()],2):t._e()],1)],1)})),1)},o=[]},b354:function(t,e,n){"use strict";n.r(e);var i=n("8583"),a=n("3eaf");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);var r,l=n("f0c5"),s=Object(l["a"])(a["default"],i["b"],i["c"],!1,null,"6c8549c7",null,!1,i["a"],r);e["default"]=s.exports},ca06:function(t,e,n){"use strict";var i=n("1845"),a=n.n(i);a.a},e4dc:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/.comment[data-v-224e58c8]:last-child{border:none}.name[data-v-224e58c8]{color:grey;font-size:%?26?%;line-height:150%}.comment[data-v-224e58c8]{display:-webkit-box;display:-webkit-flex;display:flex;padding:%?30?% 0;border-bottom:%?1?% solid #f0f0f0}.comment .left uni-image[data-v-224e58c8]{width:%?70?%;height:%?70?%;border-radius:50%}.comment .right[data-v-224e58c8]{-webkit-box-flex:1;-webkit-flex:1;flex:1;padding-left:%?20?%;font-size:%?30?%}.comment .right .top[data-v-224e58c8]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.comment .right .top .like[data-v-224e58c8]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;color:#9a9a9a;font-size:%?26?%}.comment .right .top .highlight[data-v-224e58c8]{color:#5677fc}.comment .right .content[data-v-224e58c8],\r\n.comment .right .kid-content[data-v-224e58c8]{margin-bottom:%?10?%;color:#505050;font-size:%?28?%;line-height:150%;text-align:left}.comment .right .reply-box .item[data-v-224e58c8]{display:-webkit-box;display:-webkit-flex;display:flex;padding:%?20?% %?20?% 0 %?20?%}.comment .right .reply-box .all-reply[data-v-224e58c8]{padding:%?20?%;color:#2a82e4;font-size:%?28?%;line-height:150%;text-align:left;font-weight:700}.comment .right .reply-box .all-reply .more[data-v-224e58c8]{margin-left:%?6?%}.comment .right .bottom[data-v-224e58c8]{margin-top:%?20?%;display:-webkit-box;display:-webkit-flex;display:flex;font-size:%?24?%;color:#9a9a9a}.comment .right .bottom .reply[data-v-224e58c8]{color:#5677fc;margin-left:%?10?%}.reply-name[data-v-224e58c8]{color:#2a82e4;font-size:%?28?%;line-height:150%}.item .left[data-v-224e58c8]{display:-webkit-box;display:-webkit-flex;display:flex}.item .left uni-image[data-v-224e58c8]{-webkit-box-flex:1;-webkit-flex:1;flex:1;width:%?50?%;height:%?50?%;line-height:%?80?%}',""]),t.exports=e},ee60:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.getArtileDetails=o,e.getArticleCommentList=r,e.addArticleComment=l,e.getChildCommentList=s,e.favour=c,e.collect=m,e.reprint=d;var a=i(n("367e"));function o(t){return a.default.get("/api/article/detail/"+t)}function r(t){return a.default.get("/api/comment/article",{params:t})}function l(t){return a.default.post("/api/comment/article",t)}function s(t){return a.default.get("/api/comment/children/article",{id:t.id})}function c(t){return a.default.post("/api/favour/article",t)}function m(t){return a.default.post("/api/collect/article",t)}function d(t){return a.default.post("/api/article/reprint",t)}},f4e3:function(t,e,n){"use strict";n.r(e);var i=n("8f93"),a=n("564f");for(var o in a)"default"!==o&&function(t){n.d(e,t,(function(){return a[t]}))}(o);n("ca06");var r,l=n("f0c5"),s=Object(l["a"])(a["default"],i["b"],i["c"],!1,null,"224e58c8",null,!1,i["a"],r);e["default"]=s.exports}}]);