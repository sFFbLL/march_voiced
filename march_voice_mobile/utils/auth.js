/**
 * token处理
 */
import Cookies from 'js-cookie'
import Config from '@/settings' 

const TokenKey = Config.TokenKey;
const OpenKey = Config.OpenKey;

export function getToken() {
    return Cookies.get(TokenKey)
}
export function getOpenId() {
    return Cookies.get(OpenKey)
}

export function setToken(token) {
    // if (rememberMe) {
        return Cookies.set(TokenKey, token)
    // } else {
    //     let token_temp = Cookies.set(TokenKey, token)
    //     return token_temp
    // }
}

export function setOpenId(OpenId) {
    return Cookies.set(OpenKey, OpenId)
}



export function removeToken() {
    return Cookies.remove(TokenKey)
}