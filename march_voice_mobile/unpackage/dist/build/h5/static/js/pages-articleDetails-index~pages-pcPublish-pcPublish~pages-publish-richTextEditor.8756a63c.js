(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-articleDetails-index~pages-pcPublish-pcPublish~pages-publish-richTextEditor"],{"06c5":function(t,e,n){"use strict";n("a630"),n("fb6a"),n("d3b7"),n("25f0"),n("3ca3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=a;var i=o(n("6b75"));function o(t){return t&&t.__esModule?t:{default:t}}function a(t,e){if(t){if("string"===typeof t)return(0,i.default)(t,e);var n=Object.prototype.toString.call(t).slice(8,-1);return"Object"===n&&t.constructor&&(n=t.constructor.name),"Map"===n||"Set"===n?Array.from(t):"Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)?(0,i.default)(t,e):void 0}}},"19e6":function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,".container[data-v-02c1b132]{padding:%?30?% 0;box-sizing:border-box;padding-bottom:%?120?%}.ql-container[data-v-02c1b132]{line-height:160%;font-size:%?34?%;width:calc(100% - %?60?%);height:auto;margin:0 auto}.tool-view[data-v-02c1b132]{width:100vw;position:fixed;bottom:0;left:0}.tool[data-v-02c1b132]{height:%?100?%;display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-justify-content:space-around;justify-content:space-around;width:100%;background:#eee}.font-more[data-v-02c1b132]{position:absolute;left:0;bottom:%?100?%;display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-justify-content:space-around;justify-content:space-around;width:100%;background:#ebebeb;overflow:hidden;-webkit-transition:all .15s;transition:all .15s}.setting-layer[data-v-02c1b132]{position:absolute;bottom:%?100?%;background:#fff;width:%?250?%;right:%?20?%;box-shadow:0 2px 8px rgba(0,0,0,.15);border-radius:%?8?%}.setting-layer .single[data-v-02c1b132]{height:%?80?%;font-size:%?32?%;padding:0 %?30?%;display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;line-height:%?80?%;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;color:#666}.setting-layer .single .icon[data-v-02c1b132]{margin-right:%?20?%}.setting-layer-mask[data-v-02c1b132]{position:fixed;left:0;top:0;width:100vw;height:100vh;background:transparent}.content-length-tip-box[data-v-02c1b132]{background:#eee;text-align:right;padding:0 %?10?%\r\n\t/* display: flex; */}.content-length-tip-text[data-v-02c1b132]{font-size:%?30?%;color:#999\r\n\t/* align-items: flex-end; */\r\n\t/* right: 0; */}",""]),t.exports=e},"1dc6":function(t,e,n){"use strict";n.r(e);var i=n("c7a2"),o=n.n(i);for(var a in i)"default"!==a&&function(t){n.d(e,t,(function(){return i[t]}))}(a);e["default"]=o.a},"1fbd":function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return o})),n.d(e,"c",(function(){return a})),n.d(e,"a",(function(){return i}));var o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"container",style:{paddingBottom:t.showMoreTool?"220rpx":"120rpx"}},[n("v-uni-editor",{ref:"editot",staticClass:"ql-container",attrs:{placeholder:t.placeholder,"show-img-size":!0,"show-img-toolbar":!0,"show-img-resize":!0,"read-only":t.readOnly,id:"editor"},on:{ready:function(e){arguments[0]=e=t.$handleEvent(e),t.onEditorReady.apply(void 0,arguments)},statuschange:function(e){arguments[0]=e=t.$handleEvent(e),t.statuschange.apply(void 0,arguments)},focus:function(e){arguments[0]=e=t.$handleEvent(e),t.editFocus.apply(void 0,arguments)},blur:function(e){arguments[0]=e=t.$handleEvent(e),t.editBlur.apply(void 0,arguments)},input:function(e){arguments[0]=e=t.$handleEvent(e),t.editChange.apply(void 0,arguments)}}}),n("v-uni-view",{staticClass:"tool-view"},[n("v-uni-view",{staticClass:"content-length-tip-box"},[n("span",{staticClass:"content-length-tip-text"},[t._v("当前字数："+t._s(t.textLength))])]),n("v-uni-view",{staticClass:"tool"},[n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6f3;","font-size":"44rpx",title:"插入图片"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.insertImage.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6f9;","font-size":"44rpx",title:"修改文字样式",color:t.showMoreTool?t.activeColor:"#666666"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.showMore.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6eb;","font-size":"44rpx",title:"分割线"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.insertDivider.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6e8;","font-size":"44rpx",title:"撤销"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.undo.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe705;","font-size":"44rpx",title:"重做"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.redo.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xeb8a;","font-size":"44rpx",title:"设置"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.showSetting.apply(void 0,arguments)}}})],1),n("v-uni-view",{staticClass:"font-more",style:{height:t.showMoreTool?"100rpx":0}},[n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6e7;","font-size":"44rpx",title:"加粗",color:t.showBold?t.activeColor:"#666666"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.setBold.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6fe;","font-size":"44rpx",title:"斜体",color:t.showItalic?t.activeColor:"#666666"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.setItalic.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6f8;","font-size":"44rpx",title:"分割线",color:t.showIns?t.activeColor:"#666666"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.setIns.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6e3;","font-size":"44rpx",title:"标题",color:t.showHeader?t.activeColor:"#666666"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.setHeader.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6f1;","font-size":"44rpx",title:"居中",color:t.showCenter?t.activeColor:"#666666"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.setCenter.apply(void 0,arguments)}}}),n("jinIcon",{staticClass:"single",attrs:{type:"&#xe6ed;","font-size":"44rpx",title:"居右",color:t.showRight?t.activeColor:"#666666"},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.setRight.apply(void 0,arguments)}}})],1),t.showSettingLayer?n("v-uni-view",{staticClass:"setting-layer-mask",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.showSetting.apply(void 0,arguments)}}}):t._e(),t.showSettingLayer?n("v-uni-view",{staticClass:"setting-layer"},[n("v-uni-view",{staticClass:"single",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.release(2)}}},[n("v-uni-view",[t._v("发布")])],1),n("v-uni-view",{staticClass:"single",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.release(0)}}},[n("v-uni-view",[t._v("保存草稿")])],1)],1):t._e()],1)],1)},a=[]},"26a4":function(t,e,n){var i=n("e31a");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var o=n("4f06").default;o("b9193c20",i,!0,{sourceMap:!1,shadowMode:!1})},"33b8":function(t,e,n){"use strict";n.r(e);var i=n("e500"),o=n.n(i);for(var a in i)"default"!==a&&function(t){n.d(e,t,(function(){return i[t]}))}(a);e["default"]=o.a},"3ac6":function(t,e,n){"use strict";var i=n("26a4"),o=n.n(i);o.a},"6b75":function(t,e,n){"use strict";function i(t,e){(null==e||e>t.length)&&(e=t.length);for(var n=0,i=new Array(e);n<e;n++)i[n]=t[n];return i}Object.defineProperty(e,"__esModule",{value:!0}),e.default=i},"6edb":function(t,e,n){"use strict";n.r(e);var i=n("1fbd"),o=n("33b8");for(var a in o)"default"!==a&&function(t){n.d(e,t,(function(){return o[t]}))}(a);n("93b4");var s,r=n("f0c5"),c=Object(r["a"])(o["default"],i["b"],i["c"],!1,null,"02c1b132",null,!1,i["a"],s);e["default"]=c.exports},"75b0":function(t,e,n){var i=n("19e6");"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var o=n("4f06").default;o("3cac7b84",i,!0,{sourceMap:!1,shadowMode:!1})},9394:function(t,e,n){"use strict";n.r(e);var i=n("efca"),o=n("1dc6");for(var a in o)"default"!==a&&function(t){n.d(e,t,(function(){return o[t]}))}(a);n("3ac6");var s,r=n("f0c5"),c=Object(r["a"])(o["default"],i["b"],i["c"],!1,null,"8e53f232",null,!1,i["a"],s);e["default"]=c.exports},"93b4":function(t,e,n){"use strict";var i=n("75b0"),o=n.n(i);o.a},b85c:function(t,e,n){"use strict";n("a4d3"),n("e01a"),n("d28b"),n("d3b7"),n("3ca3"),n("ddb0"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=a;var i=o(n("06c5"));function o(t){return t&&t.__esModule?t:{default:t}}function a(t,e){var n;if("undefined"===typeof Symbol||null==t[Symbol.iterator]){if(Array.isArray(t)||(n=(0,i.default)(t))||e&&t&&"number"===typeof t.length){n&&(t=n);var o=0,a=function(){};return{s:a,n:function(){return o>=t.length?{done:!0}:{done:!1,value:t[o++]}},e:function(t){throw t},f:a}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var s,r=!0,c=!1;return{s:function(){n=t[Symbol.iterator]()},n:function(){var t=n.next();return r=t.done,t},e:function(t){c=!0,s=t},f:function(){try{r||null==n["return"]||n["return"]()}finally{if(c)throw s}}}}},c7a2:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={props:{type:{type:String,default:"&#xe644;"},color:{type:String,default:"#666666"},fontSize:{type:String,default:"34rpx"}},methods:{toclick:function(){this.$emit("click")}}};e.default=i},e31a:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'.content[data-v-8e53f232]{display:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center}@font-face{font-family:jin;\n\t/** 阿里巴巴矢量图标库的字体库地址，可以替换自己的字体库地址 **/src:url(https://at.alicdn.com/t/font_1491431_6m7ltjo8wi.ttf) format("truetype")}.icon[data-v-8e53f232]{font-family:jin!important;font-size:%?34?%}',""]),t.exports=e},e500:function(t,e,n){"use strict";var i=n("4ea4");n("ac1f"),n("5319"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var o=i(n("b85c"));n("96cf");var a=i(n("1da1")),s=i(n("9394")),r={props:{showImgSize:{type:Boolean,default:!1},showImgToolbar:{type:Boolean,default:!1},showImgResize:{type:Boolean,default:!1},placeholder:{type:String,default:"开始输入..."},uploadFileUrl:{type:String,default:"#"},fileKeyName:{type:String,default:"file"},header:{type:Object},html:{type:String},readOnly:{type:null,default:!1}},computed:{},data:function(){return{textLength:0,showMoreTool:!1,showBold:!1,showItalic:!1,showIns:!1,showHeader:!1,showCenter:!1,showRight:!1,showSettingLayer:!1,activeColor:"#F56C6C"}},components:{jinIcon:s.default},methods:{onEditorReady:function(t){var e=this;uni.createSelectorQuery().in(this).select(".ql-container").fields({size:!0,context:!0},(function(t){e.editorCtx=t.context,e.editorCtx.setContents({html:e.html})})).exec()},undo:function(){this.editorCtx.undo()},insertImage:function(){var t=this;uni.chooseImage({count:9,sizeType:["original","compressed"],sourceType:["album","camera"],success:function(){var e=(0,a.default)(regeneratorRuntime.mark((function e(n){var i,a,s,r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:i=n.tempFilePaths,uni.showLoading({title:"正在上传中..."}),a=(0,o.default)(i),e.prev=3,a.s();case 5:if((s=a.n()).done){e.next=11;break}return r=s.value,e.next=9,uni.uploadFile({url:t.uploadFileUrl,filePath:r,name:t.fileKeyName,header:t.header,success:function(e){var n=JSON.parse(e.data);t.editorCtx.insertImage({src:"http://linbolun.cn/api/file/download/"+n.data.path,alt:"图片",success:function(t){}}),uni.hideLoading()}});case 9:e.next=5;break;case 11:e.next=16;break;case 13:e.prev=13,e.t0=e["catch"](3),a.e(e.t0);case 16:return e.prev=16,a.f(),e.finish(16);case 19:case"end":return e.stop()}}),e,null,[[3,13,16,19]])})));function n(t){return e.apply(this,arguments)}return n}()})},insertDivider:function(){this.editorCtx.insertDivider()},redo:function(){this.editorCtx.redo()},showMore:function(){this.showMoreTool=!this.showMoreTool,this.editorCtx.setContents()},setBold:function(){this.showBold=!this.showBold,this.editorCtx.format("bold")},setItalic:function(){this.showItalic=!this.showItalic,this.editorCtx.format("italic")},checkStatus:function(t,e,n){e.hasOwnProperty(t)?this[n]=!0:this[n]=!1},statuschange:function(t){var e=t.detail;this.checkStatus("bold",e,"showBold"),this.checkStatus("italic",e,"showItalic"),this.checkStatus("ins",e,"showIns"),this.checkStatus("header",e,"showHeader"),e.hasOwnProperty("align")?"center"==e.align?(this.showCenter=!0,this.showRight=!1):"right"==e.align?(this.showCenter=!1,this.showRight=!0):(this.showCenter=!1,this.showRight=!1):(this.showCenter=!1,this.showRight=!1)},setIns:function(){this.showIns=!this.showIns,this.editorCtx.format("ins")},setHeader:function(){this.showHeader=!this.showHeader,this.editorCtx.format("header",!!this.showHeader&&"H2")},setCenter:function(){this.showCenter=!this.showCenter,this.editorCtx.format("align",!!this.showCenter&&"center")},setRight:function(){this.showRight=!this.showRight,this.editorCtx.format("align",!!this.showRight&&"right")},showSetting:function(){this.showSettingLayer=!this.showSettingLayer},editFocus:function(){return(0,a.default)(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:case"end":return t.stop()}}),t)})))()},editBlur:function(){},editChange:function(t){this.textLength=t.detail.text.replace(/[\r\n\s+]/g,"").length,this.$emit("eidtorChange",this.textLength)},reLoadEditor:function(t){t&&(this.html=t.html,this.textLength=t.textLength),this.onEditorReady()},release:function(t){var e=this;this.showSettingLayer=!1,this.editorCtx.getContents({success:function(n){Object.assign(n,{isPublic:t,textLength:e.textLength}),e.$emit("editOk",n)}})}}};e.default=r},efca:function(t,e,n){"use strict";var i;n.d(e,"b",(function(){return o})),n.d(e,"c",(function(){return a})),n.d(e,"a",(function(){return i}));var o=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"content"},[n("v-uni-view",{staticClass:"icon",style:{color:t.color,fontSize:t.fontSize},domProps:{innerHTML:t._s(t.type)},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.toclick.apply(void 0,arguments)}}})],1)},a=[]}}]);