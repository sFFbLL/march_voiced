(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-fans-index"],{"0cad":function(n,t,e){"use strict";e.r(t);var i=e("6f44"),a=e("4e27");for(var o in a)"default"!==o&&function(n){e.d(t,n,(function(){return a[n]}))}(o);e("a208");var s,r=e("f0c5"),u=Object(r["a"])(a["default"],i["b"],i["c"],!1,null,"48f2b592",null,!1,i["a"],s);t["default"]=u.exports},"2dec":function(n,t,e){var i=e("31e2");"string"===typeof i&&(i=[[n.i,i,""]]),i.locals&&(n.exports=i.locals);var a=e("4f06").default;a("21cbf66c",i,!0,{sourceMap:!1,shadowMode:!1})},"31e2":function(n,t,e){var i=e("24fb");t=i(!1),t.push([n.i,"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n/* 关注数粉丝数文章数去掉最后的 |  */.inner-text-message span[data-v-48f2b592]:not(:last-child){margin-right:%?10?%;padding-right:%?10?%;border-right:%?1?% solid #969696}",""]),n.exports=t},"4e27":function(n,t,e){"use strict";e.r(t);var i=e("5087"),a=e.n(i);for(var o in i)"default"!==o&&function(n){e.d(t,n,(function(){return i[n]}))}(o);t["default"]=a.a},5087:function(n,t,e){"use strict";var i=e("4ea4");Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var a=i(e("40e6")),o=i(e("d508")),s=e("87d4"),r={data:function(){return{list:[],current:1,size:10}},components:{attentionAndFansCell:a.default,uniLoadMore:o.default},onLoad:function(n){uni.startPullDownRefresh()},onPullDownRefresh:function(){this.current=1;var n=this;n.list=[],setTimeout((function(){n.getList(),uni.stopPullDownRefresh()}),1e3)},onReachBottom:function(){this.getList()},methods:{getList:function(){var n=this,t={current:this.current,size:this.size};(0,s.getFansnList)(t).then((function(t){console.log(t),t.data.follow.length>0&&(n.list.push.apply(n.list,t.data.follow),n.current++)}))}}};t.default=r},"6f44":function(n,t,e){"use strict";e.d(t,"b",(function(){return a})),e.d(t,"c",(function(){return o})),e.d(t,"a",(function(){return i}));var i={uniList:e("2dfb").default},a=function(){var n=this,t=n.$createElement,e=n._self._c||t;return e("uni-list",{staticClass:"list-item"},[n._l(n.list,(function(n,t){return e("attentionAndFansCell",{key:t,staticClass:"item",attrs:{aid:n.id,nickname:n.nickname,isFollow:n.isFollow,avatarPath:n.avatarPath}})})),e("uniLoadMore")],2)},o=[]},a208:function(n,t,e){"use strict";var i=e("2dec"),a=e.n(i);a.a}}]);