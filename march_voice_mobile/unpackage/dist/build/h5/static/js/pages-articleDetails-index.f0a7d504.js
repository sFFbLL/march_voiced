(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-articleDetails-index"],{1874:function(t,n,e){"use strict";var i=e("4ea4");Object.defineProperty(n,"__esModule",{value:!0}),n.getMarchCircleInfo=o,n.marchCircleList=l,n.publishIdea=c,n.joinMarchCircle=s,n.changeFavour=r,n.ideaDetail=d,n.publishComment=u,n.ideaCommentList=f,n.ideaChildCommentList=m;var a=i(e("367e"));function o(t){return a.default.get("/api/march/msg",{params:t})}function l(t){return a.default.get("/api/march",{params:t})}function c(t){return a.default.post("/api/march",{content:t.content,imageList:t.imageList,image:t.image})}function s(t){return a.default.post("/api/apply/march",{params:t})}function r(t){return a.default.post("/api/favour/march",{id:t.id,type:t.type})}function d(t){return a.default.get("/api/march/detail/".concat(t))}function u(t){return a.default.post("/api/march/comment",{params:t})}function f(t){return a.default.get("/api/comment/marchsoft",{params:t})}function m(t){return a.default.get("/api/comment/children/article",{id:t.id,current:t.current,size:t.size})}},2071:function(t,n,e){"use strict";var i=e("4ea4");e("a9e3"),Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var a=i(e("2208")),o=e("87d4"),l={props:{aid:{type:Number,default:null},nickname:{type:String,default:"昵称"},avatarPath:{type:String,default:null},isFollow:{type:Number,default:null}},data:function(){return{isDisabled:!1,openid:null}},methods:{changeBtn:function(){var t=this;this.isDisabled=!0;var n=this;setTimeout((function(){n.isDisabled=!1}),1e3);var e={id:this.aid};(0,o.changeStatus)(e).then((function(n){switch(t.isFollow){case 0:t.isFollow=1;break;case 1:t.isFollow=0;break}}))},inToMine:function(){console.log(this.aid),uni.navigateTo({url:"/pages/universalPersonalCenter/index?id="+this.aid})},getImgUrl:function(t){return a.default.imgUrl+t}}};n.default=l},"27ca":function(t,n,e){"use strict";var i;e.d(n,"b",(function(){return a})),e.d(n,"c",(function(){return o})),e.d(n,"a",(function(){return i}));var a=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-uni-view",{staticClass:"attention-cell"},[e("v-uni-view",{staticClass:"flex-item"},[e("a",{staticClass:"left-img-box inner-box",on:{click:function(n){n.stopPropagation(),arguments[0]=n=t.$handleEvent(n),t.inToMine.apply(void 0,arguments)}}},[e("v-uni-image",{staticClass:"inner-img",attrs:{src:t.getImgUrl(t.avatarPath),mode:"aspectFill"}})],1),e("v-uni-view",{staticClass:"middle-text-box inner-box",on:{click:function(n){n.stopPropagation(),arguments[0]=n=t.$handleEvent(n),t.inToMine.apply(void 0,arguments)}}},[e("v-uni-view",{staticClass:"inner-middle-box inner-box"},[e("v-uni-view",{staticClass:"inner-text-name"},[t._v(t._s(t.nickname)),t._t("afterNicknameText")],2),e("v-uni-view",{staticClass:"inner-text-message greay-text"},[t._t("middleText")],2),e("v-uni-view",{staticClass:"inner-text-selfintroduce greay-text"},[t._t("underText")],2)],1)],1),e("v-uni-view",{staticClass:"right-button-box"},[null!=t.isFollow?e("v-uni-view",[e("v-uni-button",{directives:[{name:"show",rawName:"v-show",value:0===t.isFollow,expression:"isFollow === 0"}],staticClass:"right-button",attrs:{type:"default-green",disabled:t.isDisabled,iconType:"circle"},on:{click:function(n){n.stopPropagation(),arguments[0]=n=t.$handleEvent(n),t.changeBtn.apply(void 0,arguments)}}},[e("span",[t._v("关 注")])]),e("v-uni-button",{directives:[{name:"show",rawName:"v-show",value:1===t.isFollow,expression:"isFollow === 1"}],staticClass:"right-button",attrs:{type:"default",disabled:t.isDisabled,iconType:"circle"},on:{click:function(n){n.stopPropagation(),arguments[0]=n=t.$handleEvent(n),t.changeBtn.apply(void 0,arguments)}}},[e("span",[t._v("已关注")])])],1):t._e()],1)],1)],1)},o=[]},"40e6":function(t,n,e){"use strict";e.r(n);var i=e("27ca"),a=e("bd53");for(var o in a)"default"!==o&&function(t){e.d(n,t,(function(){return a[t]}))}(o);e("dee4");var l,c=e("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"8142f390",null,!1,i["a"],l);n["default"]=s.exports},"533c":function(t,n,e){"use strict";e.d(n,"b",(function(){return a})),e.d(n,"c",(function(){return o})),e.d(n,"a",(function(){return i}));var i={uIcon:e("ac18").default},a=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-uni-view",{staticClass:"article-details"},[e("v-uni-view",{staticClass:"view-content"},[e("v-uni-view",{staticClass:"article-title"},[t._v(t._s(t.articleInfo.title))]),e("attentionAndFansCell",{staticClass:"user-info",attrs:{aid:t.articleInfo.create_by,nickname:t.articleInfo.nickname,avatarPath:t.articleInfo.avatarPath,isFollow:t.articleInfo.isFollow}}),e("v-uni-view",{staticClass:"article-message"},[e("v-uni-text",{staticClass:"word-num"},[e("v-uni-text",[t._v("字数")]),e("v-uni-text",{staticClass:"word-num-num"},[t._v(t._s(t.articleInfo.word_count))])],1),e("v-uni-text",{staticClass:"time"},[e("v-uni-text",[t._v(t._s(t.format(t.articleInfo.update_time)))])],1)],1),e("v-uni-view",{staticClass:"article-content"},[e("v-uni-text",{domProps:{innerHTML:t._s(t.articleInfo.content)}})],1)],1),e("v-uni-view",{staticClass:"comment-view"},[e("v-uni-view",{staticClass:"top-title"},[e("v-uni-text",[t._v("评论")]),e("v-uni-text",[t._v("("+t._s(t.commentCount)+")")])],1),e("v-uni-view",{staticClass:"comment-list"},[e("comment",{attrs:{kidsCommentCount:t.kidsCommentCount,commentList:t.commentList},on:{childFn:function(n){arguments[0]=n=t.$handleEvent(n),t.comment.apply(void 0,arguments)}}})],1)],1),e("v-uni-view",{staticClass:"bottom-nav"},[e("v-uni-view",{staticClass:"bottom-nav-centre"},[e("v-uni-view",{staticClass:"attention-icon comment"},[e("v-uni-view",{staticClass:"attention-item",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.comment()}}},[e("u-icon",{attrs:{name:"pinglun","custom-prefix":"custom-icon"}})],1)],1),e("v-uni-view",{staticClass:"attention-icon like",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.like()}}},[e("v-uni-view",{staticClass:"attention-item"},[e("u-icon",{attrs:{name:t.likeIcon,color:t.likeColor}}),e("v-uni-text",{staticClass:"attention-num"},[t._v(t._s(t.articleInfo.favourTotal))])],1)],1),e("v-uni-view",{staticClass:"attention-icon collect",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.collent()}}},[e("v-uni-view",{staticClass:"attention-item"},[e("u-icon",{attrs:{name:t.collentIcon,"custom-prefix":"custom-icon"}}),e("v-uni-text",{staticClass:"attention-num"},[t._v(t._s(t.articleInfo.collectTotal))])],1)],1),e("v-uni-view",{staticClass:"attention-icon",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.transpond()}}},[e("u-icon",{attrs:{name:"zhuanfa"}})],1)],1)],1),t.isComment?e("commentInput",{attrs:{show:t.showAddComment,type:t.type,addCommentArg:t.addCommentArg},on:{addComment:function(n){arguments[0]=n=t.$handleEvent(n),t.addComment.apply(void 0,arguments)},addChildComment:function(n){arguments[0]=n=t.$handleEvent(n),t.addChildComment.apply(void 0,arguments)},childFn:function(n){arguments[0]=n=t.$handleEvent(n),t.parentFn.apply(void 0,arguments)}}}):t._e()],1)},o=[]},"55d9":function(t,n,e){var i=e("24fb");n=i(!1),n.push([t.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\t/* 大盒子样式 */.flex-item[data-v-8142f390]{display:-webkit-box;display:-webkit-flex;display:flex;box-sizing:border-box;border-bottom:%?6?% solid #f0f0f0;padding:%?20?% %?30?% %?20?% %?30?%\n\t\t/* margin-bottom: 20rpx; */}\n\t/* 左侧图片盒子样式 */.left-img-box[data-v-8142f390]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:1;-webkit-flex:1;flex:1;margin-right:%?10?%;\n\t\t/* width: 112rpx; */\n\t\t/* height: 112rpx; */vertical-align:middle}\n\t/* 内部盒子样式 */\n\t/* .inner-box {\n\t\tdisplay: inline-block;\n\t} */\n\t/* .inner-box {\n    display: inline-block;\n  } */\n\t/* 图片样式 */.inner-img[data-v-8142f390]{height:%?70?%;width:%?70?%;border-radius:50%;box-shadow:0 0 1px 1px rgba(0,0,0,.2)}\n\t/* 中部文字盒子样式 */.middle-text-box[data-v-8142f390]{-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:6;-webkit-flex:6;flex:6;display:-webkit-box;display:-webkit-flex;display:flex}\n\t/* 灰色字体 */.greay-text[data-v-8142f390]{color:#969696}\n\t/* 昵称 */.inner-text-name[data-v-8142f390]{font-size:%?30?%;font-weight:700;vertical-align:middle}\n\t/* 关注数粉丝数文章数 */.inner-text-message[data-v-8142f390]{padding-top:%?4?%;font-size:%?24?%}\n\t/* 自我介绍 */.inner-text-selfintroduce[data-v-8142f390]{padding-top:%?4?%;font-size:%?24?%;width:%?420?%;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}\n\t/* 右侧按钮盒子 */.right-button-box[data-v-8142f390]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-flex:1;-webkit-flex:1;flex:1;-webkit-box-align:center;-webkit-align-items:center;align-items:center}\n\t/* 右侧按钮 */.right-button[data-v-8142f390]{border-radius:%?58?%;width:%?130?%;height:%?56?%;font-size:%?27?%;font-weight:400;line-height:%?56?%;color:#969696}\n\t/* 修改按钮样式 */uni-button[type=default][data-v-8142f390]{border:%?1?% #969696 solid;box-sizing:border-box}\n\t/* 修改按钮样式 */uni-button[type=default-green][data-v-8142f390]{color:#fff;background-color:#404040}\n\t/* 修改点击按钮后的样式 */.button-hover[type=default-green][data-v-8142f390]{background-color:#1a1a1a}\n\t/* 修改按钮样式 */uni-button[data-v-8142f390]{margin:0;padding:0}",""]),t.exports=n},6872:function(t,n,e){var i=e("7c15");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=e("4f06").default;a("5aede26b",i,!0,{sourceMap:!1,shadowMode:!1})},"784d":function(t,n,e){"use strict";e.r(n);var i=e("533c"),a=e("8140");for(var o in a)"default"!==o&&function(t){e.d(n,t,(function(){return a[t]}))}(o);e("b4d5");var l,c=e("f0c5"),s=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"b3a23d84",null,!1,i["a"],l);n["default"]=s.exports},"7c15":function(t,n,e){var i=e("24fb");n=i(!1),n.push([t.i,'.article-details[data-v-b3a23d84]{background-color:#f7f7f7}.view-content[data-v-b3a23d84]{padding:0 %?30?%;background-color:#fff}.article-title[data-v-b3a23d84]{color:#404040;font-size:%?36?%;line-height:150%;font-weight:700}[data-v-b3a23d84] .attention-cell .flex-item{padding:%?20?% 0;border:none}[data-v-b3a23d84] .attention-cell .left-img-box{margin-right:0}[data-v-b3a23d84] .attention-cell .right-button{background-color:#fff}.article-message[data-v-b3a23d84]{color:#999797;font-size:%?24?%}.word-num-num[data-v-b3a23d84]{margin-left:%?4?%}.time[data-v-b3a23d84]{float:right}.article-content[data-v-b3a23d84]{padding:%?20?% %?10?%}[data-v-b3a23d84] .article-content .ql-container{margin:0;width:100%;color:#404040;font-size:%?32?%;line-height:150%}[data-v-b3a23d84] .article-content .container{padding-top:%?20?%}[data-v-b3a23d84] .article-content .ql-container .ql-editor{text-align:justify}[data-v-b3a23d84] .article-content .tool-view{display:none}\n\n/* 底部导航:点赞，评论等 */.bottom-nav[data-v-b3a23d84]{position:fixed;bottom:0;width:100%;height:%?120?%;border:%?1?% solid #f0f0f0;border-width:%?1?% 0;background-color:#fff}.bottom-nav-centre[data-v-b3a23d84]{display:-webkit-box;display:-webkit-flex;display:flex;margin:0 %?50?%;height:100%}.attention-icon[data-v-b3a23d84]::before{content:"";display:inline-block;vertical-align:middle;height:100%}.attention-icon[data-v-b3a23d84]{-webkit-box-flex:1;-webkit-flex:1;flex:1;text-align:center}[data-v-b3a23d84] .attention-icon .u-icon{font-size:%?50?%!important;color:#404040!important}.attention-item[data-v-b3a23d84]{position:relative;display:inline-block}.attention-icon .attention-num[data-v-b3a23d84]{position:absolute;top:%?-8?%}\n\n/* 评论样式 */.top-title[data-v-b3a23d84]{height:%?88?%;color:#505050;font-size:%?34?%;line-height:150%;font-weight:700}.top-title uni-text[data-v-b3a23d84]:first-child{display:inline-block;margin:0 %?10?% 0 %?30?%;line-height:%?88?%}.comment-view[data-v-b3a23d84]{margin-bottom:%?120?%}.comment-list[data-v-b3a23d84]{padding:0 %?30?%;background-color:#fff}',""]),t.exports=n},8140:function(t,n,e){"use strict";e.r(n);var i=e("a7f0"),a=e.n(i);for(var o in i)"default"!==o&&function(t){e.d(n,t,(function(){return i[t]}))}(o);n["default"]=a.a},"85a7":function(t,n,e){var i=e("55d9");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=e("4f06").default;a("1fc44d7a",i,!0,{sourceMap:!1,shadowMode:!1})},"87d4":function(t,n,e){"use strict";var i=e("4ea4");Object.defineProperty(n,"__esModule",{value:!0}),n.getAttentionList=o,n.getFansnList=l,n.changeStatus=c;var a=i(e("367e"));function o(t){return a.default.get("/api/follow",{params:t})}function l(t){return a.default.get("/api/follow/fans",{params:t})}function c(t){return a.default.post("/api/follow",{id:t.id})}},a7f0:function(t,n,e){"use strict";var i=e("4ea4");e("4160"),e("a9e3"),e("159b"),Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var a=i(e("b85c")),o=i(e("ade3")),l=i(e("d50f")),c=i(e("40e6")),s=i(e("6edb")),r=i(e("f4e3")),d=i(e("1e61")),u=i(e("2184")),f=e("ee60"),m={data:function(){var t;return t={showAddComment:!0,id:0,current:1,size:5,childSize:3,articleInfo:{},addCommentArg:{},commentCount:0,commentList:[],isComment:!1,type:0},(0,o.default)(t,"addCommentArg",{}),(0,o.default)(t,"clickChild",!1),(0,o.default)(t,"likeIcon","heart"),(0,o.default)(t,"collentIcon","shoucang1"),(0,o.default)(t,"likeColor","black"),(0,o.default)(t,"clickLike",!1),(0,o.default)(t,"clickCollect",!1),(0,o.default)(t,"kidsCommentCount",[]),t},components:{articleTitle:l.default,attentionAndFansCell:c.default,jinEdit:s.default,comment:r.default,commentInput:u.default},onLoad:function(t){var n=this;this.id=Number(t.id);var e=this.id;(0,f.getArtileDetails)(e).then((function(t){0===t.code&&(n.articleInfo=t.data)})).catch((function(t){console.log(t,"err login")})),console.log(this.articleInfo);var i={id:e,current:this.current,size:this.size,childSize:this.childSize};(0,f.getArticleCommentList)(i).then((function(t){if(0===t.code){console.log(t.data),n.commentList.CommentSum=t.data,n.commentCount=t.data.CommentSum.length;var e,i=n.commentList.CommentSum,o=(0,a.default)(i);try{for(o.s();!(e=o.n()).done;){var l=e.value;console.log(i),l.ChildComments?n.kidsCommentCount.push(l.ChildComments.length):n.kidsCommentCount.push(0)}}catch(c){o.e(c)}finally{o.f()}}}))},created:function(){this.id;this.addCommentArg={id:this.id,replyId:0,follewId:0,childComment:!1}},methods:{like:function(){if(this.clickLike)this.articleInfo.favourTotal--,this.likeIcon="heart",this.likeColor="#080100";else{var t=this.id;this.likeIcon="heart-fill",this.likeColor="#d81e06",this.articleInfo.favourTotal++,(0,f.favour)(t).then((function(t){}))}this.clickLike=!this.clickLike},collent:function(){if(this.clickCollect)this.collentIcon="shoucang1",this.articleInfo.collectTotal--;else{var t=this.id;this.collentIcon="shoucang2",this.articleInfo.collectTotal++,(0,f.collect)(t).then((function(t){}))}this.clickCollect=!this.clickCollect},transpond:function(){var t=this.id;(0,f.reprint)(t).then((function(t){})),uni.showLoading({title:"加载中"}),setTimeout((function(){uni.hideLoading()}),2e3),setTimeout((function(){uni.showToast({title:"转发成功",duration:2e3})}),2e3)},comment:function(t){this.id;this.isComment=!0,t&&(this.addCommentArg={childComment:t.childComment,id:this.id,index:t.index,replyId:t.replyId,follewId:t.follewId,replyName:t.replyName})},parentFn:function(t){this.isComment=!1},addComment:function(t){this.commentList.unshift(t),this.isComment=!1,this.commentCount++},addChildComment:function(t){this.commentList[t.index].commentKids.push(t),this.isComment=!1,this.commentCount++},format:function(t){var n=new Date(t),e=(0,d.default)(n).format("YYYY-MM-DD HH:mm:ss");return e}},mounted:function(){var t=0;this.commentList.forEach((function(n){t++,n.commentKids.forEach((function(n){t++}))})),this.commentCount=t}};n.default=m},b4d5:function(t,n,e){"use strict";var i=e("6872"),a=e.n(i);a.a},bd53:function(t,n,e){"use strict";e.r(n);var i=e("2071"),a=e.n(i);for(var o in i)"default"!==o&&function(t){e.d(n,t,(function(){return i[t]}))}(o);n["default"]=a.a},dee4:function(t,n,e){"use strict";var i=e("85a7"),a=e.n(i);a.a}}]);