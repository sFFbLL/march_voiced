import {unreadMessage} from "./api/message-api.js"

export function check() {
	unreadMessage().then(res => {
		if (res.data.count > 0) {
			uni.showTabBarRedDot({
				index: 3
			})
		}
	})
}
