/**
 * token处理
 */
import Cookies from 'js-cookie'
import Config from '@/settings'

const TokenKey = Config.TokenKey;
const OpenKey = Config.OpenKey;
const UserName = Config.UserName;
const AvatarPath = Config.AvatarPath;
export function getToken() {
	return Cookies.get(TokenKey)
}
export function getOpenId() {
	return Cookies.get(OpenKey)
}
export function getUserName() {
	return Cookies.get(unescape(UserName));
}
export function getAvatarPath() {
	return Cookies.get(AvatarPath)
}




export function setToken(token) {
	return Cookies.set(TokenKey, token)

}
export function setOpenId(OpenId) {
	return Cookies.set(OpenKey, OpenId)
}
export function setUserName(userName) {
	return Cookies.set(UserName, userName)

}
export function setAvatarPath(path) {
	let avatarPath=Config.imgUrl +path;
	console.log(avatarPath)
	return Cookies.set(AvatarPath, avatarPath)

}



export function removeToken() {
	return Cookies.remove(TokenKey)
}
