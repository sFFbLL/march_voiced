(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-pcLogin-index"],{"12a9":function(t,n,e){var a=e("24fb");n=a(!1),n.push([t.i,"[data-v-38c4b421] uni-page-wrapper{background-color:#000!important}",""]),t.exports=n},"36db":function(t,n,e){var a=e("12a9");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var i=e("4f06").default;i("31143356",a,!0,{sourceMap:!1,shadowMode:!1})},"7a8d":function(t,n,e){"use strict";var a=e("36db"),i=e.n(a);i.a},8089:function(t,n,e){"use strict";e.r(n);var a=e("f6d1"),i=e("df81");for(var r in i)"default"!==r&&function(t){e.d(n,t,(function(){return i[t]}))}(r);e("7a8d");var c,o=e("f0c5"),s=Object(o["a"])(i["default"],a["b"],a["c"],!1,null,"38c4b421",null,!1,a["a"],c);n["default"]=s.exports},bbbf:function(t,n,e){"use strict";(function(t){Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var a=e("5c16"),i=e("c98b"),r={data:function(){return{ticket:null,strData:null,baseSrc:"https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=",src:""}},created:function(){var t=this;this.login(),this.interval=setInterval((function(){t.isLogin(t.strData)}),2e3)},beforeDestroy:function(){clearInterval(this.interval)},methods:{login:function(){var n=this;(0,a.pcLogin)().then((function(e){t("log","res",e," at pages/pcLogin/index.vue:31"),n.strData=e.data.strdata,n.src=n.baseSrc+e.data.Ticket})).catch((function(n){t("log","errlogin"," at pages/pcLogin/index.vue:35")}))},isLogin:function(n){var e={strData:n};(0,a.pcIsLogin)(e).then((function(n){switch(n.code){case 0:break;case 1:(0,i.setOpenId)(n.data.openid),uni.redirectTo({url:"../pcPublish/registered"});break;case 2:(0,i.setToken)(n.token),t("log","ASDADADA"," at pages/pcLogin/index.vue:55"),uni.redirectTo({url:"../pcPublish/pcPublish"});break}}))}}};n.default=r}).call(this,e("0de9")["log"])},df81:function(t,n,e){"use strict";e.r(n);var a=e("bbbf"),i=e.n(a);for(var r in a)"default"!==r&&function(t){e.d(n,t,(function(){return a[t]}))}(r);n["default"]=i.a},f6d1:function(t,n,e){"use strict";var a;e.d(n,"b",(function(){return i})),e.d(n,"c",(function(){return r})),e.d(n,"a",(function(){return a}));var i=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-uni-view",{staticClass:"background"},[e("v-uni-image",{attrs:{src:t.src,mode:""}})],1)},r=[]}}]);