(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-fans-index"],{"6c3b":function(n,t,e){"use strict";e.d(t,"b",(function(){return a})),e.d(t,"c",(function(){return s})),e.d(t,"a",(function(){return i}));var i={uniList:e("0f04").default},a=function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("uni-list",{staticClass:"list-item"},[n._l(n.list,(function(t,i){return e("attentionAndFansCell",{key:i,staticClass:"item",attrs:{id:t.id,nickname:t.nickname,isFollow:t.isFollow,avatarPath:t.avatarPath},on:{change:function(t){arguments[0]=t=n.$handleEvent(t),n.change(i)},inToPageMine:function(e){arguments[0]=e=n.$handleEvent(e),n.inToPageMine(t.user.id)}}})})),e("uniLoadMore")],2)},s=[]},"7c72":function(n,t,e){"use strict";e.r(t);var i=e("6c3b"),a=e("aa8f");for(var s in a)"default"!==s&&function(n){e.d(t,n,(function(){return a[n]}))}(s);e("b805");var o,u=e("f0c5"),r=Object(u["a"])(a["default"],i["b"],i["c"],!1,null,"5b0fc177",null,!1,i["a"],o);t["default"]=r.exports},8012:function(n,t,e){"use strict";(function(n){var i=e("4ea4");Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var a=i(e("2afe")),s=i(e("19bf")),o=e("b62e"),u={data:function(){return{list:[],current:1,size:10}},components:{attentionAndFansCell:a.default,uniLoadMore:s.default},onLoad:function(n){this.list=[],uni.startPullDownRefresh()},onPullDownRefresh:function(){var n=this;n.list=[],setTimeout((function(){n.getList(),uni.stopPullDownRefresh()}),1e3)},onReachBottom:function(){this.getList()},methods:{getList:function(){var t=this,e={current:this.current,size:this.size};(0,o.getFansnList)(e).then((function(e){n("log",e," at pages/fans/index.vue:176"),e.data.Follow.length>0&&(t.list.push.apply(t.list,e.data.Follow),t.current++)}))},inToPageMine:function(n){uni.switchTab({url:"/pages/personalCenter/index"})},change:function(n){var t=this,e={id:this.list[n].id};(0,o.changeStatus)(e).then((function(e){switch(t.list[n].isFollow){case 0:t.list[n].isFollow=1;break;case 1:t.list[n].isFollow=0;break}}))}}};t.default=u}).call(this,e("0de9")["log"])},aa8f:function(n,t,e){"use strict";e.r(t);var i=e("8012"),a=e.n(i);for(var s in i)"default"!==s&&function(n){e.d(t,n,(function(){return i[n]}))}(s);t["default"]=a.a},b805:function(n,t,e){"use strict";var i=e("c529"),a=e.n(i);a.a},c529:function(n,t,e){var i=e("c772");"string"===typeof i&&(i=[[n.i,i,""]]),i.locals&&(n.exports=i.locals);var a=e("4f06").default;a("0b673ce9",i,!0,{sourceMap:!1,shadowMode:!1})},c772:function(n,t,e){var i=e("24fb");t=i(!1),t.push([n.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n/* 关注数粉丝数文章数去掉最后的 |  */.inner-text-message span[data-v-5b0fc177]:not(:last-child){margin-right:%?10?%;padding-right:%?10?%;border-right:%?1?% solid #969696}",""]),n.exports=t}}]);