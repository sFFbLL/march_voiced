import {
	returnWxcode,
	getWxCode,
	parseCode
} from "../utils/wxcode.js"
import {
	getToken,
	setToken,
	setOpenId,
	getOpenId,
	setUserName,
	setAvatarPath,
} from "../utils/auth.js"
import {
	login,
	creatNewUser
} from "../utils/login.js"
import {
	information
} from "./api/personInfo-api.js"
// 登录判断
export function forLogin() {
	if (!getOpenId()) {
		console.log("没有token，没有OpenId");
		let code = returnWxcode();
		if (!code) {
			getWxCode();
		} else {
			console.log("成功拿到code")
			let params = {
				code: code,
				status: 1
			}
			// 判断该用户是否注册
			login(params).then(res => {
				console.log(res, "注册")
				setOpenId(res.data.openid)

				if (res.data.status == 1) { //没有token
					// 跳转注册页
					console.log("未登录")
					uni.navigateTo({
						url: "../login/login"
					})
				} else {
					// 拿到token，登陆成功
					setToken(res.data.token);
					information().then(res => {
						setUserName(res.data.nickname);
						setAvatarPath(res.data.avatarPath);
						
					})

				}
			}).catch(err => {
				console.log(err, "err login")
			})
		}
	} else {
		let code = returnWxcode();
		if (!code) {
			getWxCode();
		} else {
			console.log("成功拿到code")
			let params = {
				code: code,
				status: 1
			}
			// 判断该用户是否注册
			login(params).then(res => {
				// 重新登陆成功
				setToken(res.data.token);
				information().then(res => {
					setUserName(res.data.nickname);
					setAvatarPath(res.data.avatarPath);
				})
			}).catch(err => {
				console.log(err, "err login")
			})
		}

	}
}
