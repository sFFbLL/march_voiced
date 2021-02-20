package bo

type GetSysMessage struct {
	Total   uint                 `json:"total"`   //总页数
	Count   uint                 `json:"count"`   //总条数
	Message *[]GetSysMessageData `json:"message"` //消息
}

type GetSysMessageData struct {
	Type       uint8 `json:"type"`       //0文章审核/1三月圈申请
	Status     uint8 `json:"status"`     //0同意/1驳回
	CreateTime int64 `json:"createTime"` //创建时间
}
