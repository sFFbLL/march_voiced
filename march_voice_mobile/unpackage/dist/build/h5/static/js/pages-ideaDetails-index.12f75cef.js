(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-ideaDetails-index"],{"029d":function(n,t,e){"use strict";e.r(t);var i=e("c48b"),a=e.n(i);for(var o in i)"default"!==o&&function(n){e.d(t,n,(function(){return i[n]}))}(o);t["default"]=a.a},"0fbe":function(n,t,e){"use strict";var i;e.d(t,"b",(function(){return a})),e.d(t,"c",(function(){return o})),e.d(t,"a",(function(){return i}));var a=function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("v-uni-view",{staticClass:"article-content",on:{click:function(t){arguments[0]=t=n.$handleEvent(t),n.toDetails.apply(void 0,arguments)}}},[e("v-uni-view",[e("v-uni-text",{ref:"articleContent",staticClass:"article-text",domProps:{innerHTML:n._s(n.textContent)}},[n._v(n._s(n.textContent))])],1),""!=n.articleImg?e("v-uni-view",{staticClass:"article-img"},[e("v-uni-image",{attrs:{src:n.articleImg,mode:"aspectFill"}})],1):n._e()],1)},o=[]},"2c0f":function(n,t,e){"use strict";var i=e("4ea4");Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var a=e("1874"),o=i(e("40e6")),r=i(e("8830")),d=i(e("e1f8")),c=i(e("7cb0")),s=i(e("f4e3")),l=i(e("2184")),u={components:{emojiControl:r.default,attentionAndFansCell:o.default,imageAdaptation:d.default,articleContent:c.default,commentInput:l.default,comment:s.default},data:function(){return{ideaInfoList:{},commentList:[],commentCount:3,current:1,size:5,childSize:3,type:1,addCommentArg:{},showAddComment:!1,isComment:!0}},onLoad:function(n){this.ideaId=n.id},created:function(){var n=this,t=this.ideaId;(0,a.ideaDetail)(t).then((function(t){n.ideaInfoList=t.data}));var e={id:t,current:this.current,size:this.size,childSize:this.childSize};console.log(e+33333333333333),(0,a.ideaCommentList)(e).then((function(t){0===t.code&&(n.commentList=t.data,n.commentCount=t.data.length)})),this.addCommentArg={id:t,replyId:0,follewId:0,childComment:!1}},methods:{comment:function(n){var t=this.ideaId;this.showAddComment=!0,n&&(this.addCommentArg={childComment:n.childComment,id:t,index:n.index,replyId:n.replyId,follewId:n.follewId,replyName:n.replyName})},parentFn:function(n){this.showAddComment=!1},addComment:function(n){this.commentList.unshift(n),this.showAddComment=!1},addChildComment:function(n){this.commentList[n.index].commentKids.push(n),this.showAddComment=!1}}};t.default=u},"47f1":function(n,t,e){"use strict";e.r(t);var i=e("2c0f"),a=e.n(i);for(var o in i)"default"!==o&&function(n){e.d(t,n,(function(){return i[n]}))}(o);t["default"]=a.a},"53ac":function(n,t,e){"use strict";e.r(t);var i=e("7e61"),a=e("7f39");for(var o in a)"default"!==o&&function(n){e.d(t,n,(function(){return a[n]}))}(o);e("c788");var r,d=e("f0c5"),c=Object(d["a"])(a["default"],i["b"],i["c"],!1,null,"71495536",null,!1,i["a"],r);t["default"]=c.exports},5742:function(n,t,e){"use strict";e("a9e3"),Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var i={name:"u-gap",props:{bgColor:{type:String,default:"transparent "},height:{type:[String,Number],default:30},marginTop:{type:[String,Number],default:0},marginBottom:{type:[String,Number],default:0}},computed:{gapStyle:function(){return{backgroundColor:this.bgColor,height:this.height+"rpx",marginTop:this.marginTop+"rpx",marginBottom:this.marginBottom+"rpx"}}}};t.default=i},"6f5e":function(n,t,e){"use strict";var i=e("b489"),a=e.n(i);a.a},"7cb0":function(n,t,e){"use strict";e.r(t);var i=e("0fbe"),a=e("029d");for(var o in a)"default"!==o&&function(n){e.d(t,n,(function(){return a[n]}))}(o);e("b30a");var r,d=e("f0c5"),c=Object(d["a"])(a["default"],i["b"],i["c"],!1,null,"cb3085ce",null,!1,i["a"],r);t["default"]=c.exports},"7da5":function(n,t,e){var i=e("24fb");t=i(!1),t.push([n.i,".article-content[data-v-cb3085ce]{margin-top:%?18?%;display:-webkit-box;display:-webkit-flex;display:flex;font-size:%?28?%;font-family:PingFangSC;color:#404040}.article-text[data-v-cb3085ce]{overflow:hidden;-webkit-line-clamp:3;text-overflow:ellipsis;display:-webkit-box;-webkit-box-orient:vertical;line-height:150%}.article-img[data-v-cb3085ce]{display:inline-block;margin-left:%?30?%}.article-img uni-image[data-v-cb3085ce]{width:%?170?%;height:100%;background-color:#eee;border-radius:4px}",""]),n.exports=t},"7e61":function(n,t,e){"use strict";var i;e.d(t,"b",(function(){return a})),e.d(t,"c",(function(){return o})),e.d(t,"a",(function(){return i}));var a=function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("v-uni-view",{staticClass:"u-gap",style:[n.gapStyle]})},o=[]},"7f39":function(n,t,e){"use strict";e.r(t);var i=e("5742"),a=e.n(i);for(var o in i)"default"!==o&&function(n){e.d(t,n,(function(){return i[n]}))}(o);t["default"]=a.a},"936b":function(n,t,e){var i=e("ad20");"string"===typeof i&&(i=[[n.i,i,""]]),i.locals&&(n.exports=i.locals);var a=e("4f06").default;a("7c85361e",i,!0,{sourceMap:!1,shadowMode:!1})},"95ec":function(n,t,e){var i=e("7da5");"string"===typeof i&&(i=[[n.i,i,""]]),i.locals&&(n.exports=i.locals);var a=e("4f06").default;a("5d69678e",i,!0,{sourceMap:!1,shadowMode:!1})},ad20:function(n,t,e){var i=e("24fb");t=i(!1),t.push([n.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */\r\n/*引用uview*/',""]),n.exports=t},b30a:function(n,t,e){"use strict";var i=e("95ec"),a=e.n(i);a.a},b489:function(n,t,e){var i=e("dfbd");"string"===typeof i&&(i=[[n.i,i,""]]),i.locals&&(n.exports=i.locals);var a=e("4f06").default;a("120ac105",i,!0,{sourceMap:!1,shadowMode:!1})},c48b:function(n,t,e){"use strict";e("a9e3"),e("ac1f"),e("5319"),Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var i={data:function(){return{textContent:""}},props:{id:{type:Number,default:null},articleContent:{type:String,default:""},articleImg:{type:String,default:""},isIdea:{type:Boolean,default:!1}},components:{},methods:{toDetails:function(){this.isIdea?uni.navigateTo({url:"../ideaDetails/index?id="+this.id}):uni.navigateTo({url:"../articleDetails/index?id="+this.id})}},created:function(){this.textContent=this.articleContent},mounted:function(){this.textContent=this.$refs.articleContent.$el.innerText.replace(/\ +/g,"").replace(/[\r\n]/g,"")}};t.default=i},c788:function(n,t,e){"use strict";var i=e("936b"),a=e.n(i);a.a},dfbd:function(n,t,e){var i=e("24fb");t=i(!1),t.push([n.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n/* 公共头像组件样式 */[data-v-11110da9] .attention-cell .flex-item{border:none;padding:0;margin-top:%?17?%}\n/* 想法大模块 */.ideacontent[data-v-11110da9]{background-color:#fff;padding:%?26?%;line-height:1.5}\n/* 想法文字部分样式 */.wordscontent[data-v-11110da9]{margin-top:%?20?%}",""]),n.exports=t},e0eb:function(n,t,e){"use strict";e.d(t,"b",(function(){return a})),e.d(t,"c",(function(){return o})),e.d(t,"a",(function(){return i}));var i={uGap:e("53ac").default},a=function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("v-uni-view",[e("u-gap",{attrs:{height:"20","bg-color":"#f5f5f5"}}),e("v-uni-view",{staticClass:"ideacontent"},[e("attentionAndFansCell",{attrs:{nickname:n.ideaInfoList.nickname,avatarPath:n.ideaInfoList.avatarPath,isFollow:n.ideaInfoList.isFollow}}),e("articleContent",{attrs:{articleContent:n.ideaInfoList.content}}),e("imageAdaptation",{attrs:{imgList:n.ideaInfoList.imgList}}),e("emojiControl",{attrs:{isShow:!1,faceTotals:n.ideaInfoList.faceTotal,likeTotals:n.ideaInfoList.likeTotal,favourTotals:n.ideaInfoList.favourTotal,commentTotals:n.ideaInfoList.commentTotal,id:n.ideaInfoList.id}})],1),e("v-uni-view",{staticClass:"comment-view"},[e("v-uni-view",{staticClass:"top-title"},[e("v-uni-text",[n._v("评论")]),e("v-uni-text",[n._v("("+n._s(n.commentCount)+")")])],1),e("v-uni-view",{staticClass:"comment-list"},[e("comment",{attrs:{commentList:n.commentList},on:{childFn:function(t){arguments[0]=t=n.$handleEvent(t),n.comment.apply(void 0,arguments)}}})],1)],1),n.isComment?e("commentInput",{attrs:{show:n.showAddComment,type:n.type,addCommentArg:n.addCommentArg},on:{addComment:function(t){arguments[0]=t=n.$handleEvent(t),n.addComment.apply(void 0,arguments)},addChildComment:function(t){arguments[0]=t=n.$handleEvent(t),n.addChildComment.apply(void 0,arguments)},childFn:function(t){arguments[0]=t=n.$handleEvent(t),n.parentFn.apply(void 0,arguments)}}}):n._e()],1)},o=[]},f5dfd:function(n,t,e){"use strict";e.r(t);var i=e("e0eb"),a=e("47f1");for(var o in a)"default"!==o&&function(n){e.d(t,n,(function(){return a[n]}))}(o);e("6f5e");var r,d=e("f0c5"),c=Object(d["a"])(a["default"],i["b"],i["c"],!1,null,"11110da9",null,!1,i["a"],r);t["default"]=c.exports}}]);