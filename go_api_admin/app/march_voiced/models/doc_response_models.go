package models

import (
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
