import Config from '../settings.js'

//向微信发送请求，获取code（在url上）
export function getWxCode() {
	// window.location.href =
		// "https://open.weixin.qq.com/connect/oauth2/authorize?" +
		// "appid=" + Config.appId + "&redirect_uri=" +
		// encodeURIComponent(location.href) +
		// "&response_type=code&scope=snsapi_base&state=123#wechat_redirect";


	window.location.href =
		"https://open.weixin.qq.com/connect/oauth2/authorize?" +
		"appid=" + Config.appId + "&redirect_uri=" +
		"http://linbolun.cn/h5/" +
		"&response_type=code&scope=snsapi_base&state=123#wechat_redirect";


}
//解析url，拿到code
export function parseCode() {
	let parmas = decodeURIComponent(window.location.search);
	let startIndex = parmas.indexOf("=") + 1;
	let endIndex = parmas.indexOf("&");
	let code = parmas.substring(startIndex, endIndex);
	return code;
}

export function returnWxcode() {
	//拿code
	let wxCode = parseCode();
	if (!wxCode) {
		// 如果code为空，向微信发送请求用于获取code
		getWxCode();
	}
	return wxCode;

}
