import Vue from 'vue'
import App from './App'
//引入vuex
import store from './store'
import uView from "uview-ui";
//在main.js内

import {Icon} from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import Vconsole from 'vconsole'
let vConsole = new Vconsole()
Vue.use(vConsole);
Vue.use(uView);
Vue.use(Icon);
//把vuex定义成全局组件
Vue.prototype.$store = store
Vue.config.productionTip = false

App.mpType = 'app'

const app = new Vue({
	...App,
	//挂载
	store
})
app.$mount()