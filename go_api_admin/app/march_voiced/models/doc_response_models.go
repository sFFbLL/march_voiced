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

type _ResponseGetSysMessage struct {
	Code    app.ResCode      `json:"code"`    // 业务响应状态码
	Message string           `json:"message"` // 提示信息
	Data    bo.GetSysMessage `json:"data"`    // 数据
}

// _ResponseGetMessage 关注页面(动态)列表页响应
type _ResponseGetMessage struct {
	Code    app.ResCode   `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    bo.GetMessage `json:"data"`    // 数据
}

// _ResponseApplyMarchUser 申请三月圈审核中用户
type _ResponseApplyMarchUser struct {
	Code    app.ResCode       `json:"code"`    // 业务响应状态码
	Message string            `json:"message"` // 提示信息
	Data    bo.ApplyMarchUser `json:"data"`    // 数据
}

// 申请三月圈审核中用户
type _ResponseGetMarchApplyUser struct {
	Code    app.ResCode   `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    bo.GetMessage `json:"data"`    // 数据
}

type _ResponseTopArticleListHandler struct {
	Code    app.ResCode  `json:"code"`    // 业务响应状态码
	Message string       `json:"message"` // 提示信息
	Data    []bo.Article `json:"data"`    // 数据
}

type _Article struct {
	Code    app.ResCode `json:"code"`    // 业务响应状态码
	Message string      `json:"message"` // 提示信息
	Data    bo.Article  `json:"data"`    // 数据
}

type _ArticleUser struct {
	Code    app.ResCode    `json:"code"`    // 业务响应状态码
	Message string         `json:"message"` // 提示信息
	Data    bo.ArticleUser `json:"data"`    // 数据
}

type _IsFavourCollectByArticleId struct {
	Code    app.ResCode                   `json:"code"`    // 业务响应状态码
	Message string                        `json:"message"` // 提示信息
	Data    bo.IsFavourCollectByArticleId `json:"data"`    // 数据
}

type _ArticleTagList struct {
	Code    app.ResCode       `json:"code"`    // 业务响应状态码
	Message string            `json:"message"` // 提示信息
	Data    bo.ArticleTagList `json:"data"`    // 数据
}

type _ArticleReprint struct {
	Code    app.ResCode   `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    bo.ArticleMsg `json:"data"`    // 数据
}

// _ResponseApplyArticleList （后台）文章审核列表页响应
type _ResponseApplyArticleList struct {
	Code    app.ResCode           `json:"code"`    // 业务响应状态码
	Message string                `json:"message"` // 提示信息
	Data    bo.ApplyArticleList `json:"data"`    // 数据
}
