import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
	state: {
		interactRead:false,
		attentionRead:false,
		otherRead:false,
	},
	mutations: {
		changeInteract(state,n){
			// 获取到未读消息
			if(n==0){
				state.interactRead=false;
			}
			// 消息已读
			else{
				state.interactRead=true;
			}
		
		},
		changeAttention(state,n){
			// 获取到未读消息
			if(n==0){
				state.attentionRead=false;
			}
			// 消息已读
			else{
				state.attentionRead=true;
			}
		},
		changeOther(state,n){
			// 获取到未读消息
			if(n==0){
				state.otherRead=false;
			}
			// 消息已读
			else{
				state.otherRead=true;
			}
		}
	},
	actions: {},
	modules: {}
})
