package models

import (
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/utils/app"
)

// _ResponseSuccess 成功响应结构体
type _ResponseSuccess struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
}

type _ResponseMessageUnRead struct {
	Code    app.ResCode          `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    dto.MessageUnreadDto `json:"data"`    // 数据
}

type _ResponseMatchSensitiveWord struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
	Data    bo.WordRes  `json:"data"`    // 数据
}

type _ResponseGetArticleComment struct {
	Code    app.ResCode          `json:"code"`    // 业务响应状态码
	Message string               `json:"message"` // 提示信息
	Data    bo.GetArticleComment `json:"data"`    // 数据
}

type _ResponseGetArticleChildComment struct {
	Code    app.ResCode        `json:"code"`    // 业务响应状态码
	Message string             `json:"message"` // 提示信息
	Data    bo.ChildrenComment `json:"code"`    // 数据
}

type _ResponseGetFollowList struct {
	Code    app.ResCode      `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
	Data    bo.GetFollowList `json:"data"`    // 数据
}

type _ResponseGetFansList struct {
	Code    app.ResCode      `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
	Data    bo.GetFollowList `json:"data"`    // 数据
}

type _ResponseGetMarchsoftComment struct {
	Code    app.ResCode            `json:"code"`    // 业务响应状态码
	Message string                 `json:"message"` // 提示信息
	Data    bo.GetMarchsoftComment `json:"data"`    // 数据
}

type _ResponseGetMarchsoftChildComment struct {
	Code    app.ResCode                 `json:"code"`    // 业务响应状态码
	Message string                      `json:"message"` // 提示信息
	Data    bo.ChildrenMarchsoftComment `json:"code"`    // 数据
}

type _ResponseSelectUserInfo struct {
	Code    app.ResCode       `json:"code"`    // 业务响应状态码
	Message string            `json:"message"` // 提示信息
	Data    bo.SelectUserInfo `json:"code"`    // 数据
}
