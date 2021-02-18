import unreadMessage from "./api/message-api.js"

export function check(){
	// unreadMessage().then(res=>{
	// 	if(res.data.count>0){
		console.log("有未读消息")
			uni.showTabBarRedDot({
				index:3
			})
	// 	}
	// })
}