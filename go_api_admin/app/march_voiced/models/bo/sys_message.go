package bo

type GetSysMessage struct {
	Current uint                 `json:"current"` //页码
	Size    uint                 `json:"size"`    //页内大小
	Pages   int                  `json:"pages"`   //总页数
	Total   int64                `json:"total"`   //总条数
	Records *[]GetSysMessageData `json:"records"` //消息
}

type GetSysMessageData struct {
	Type       uint8 `json:"type"`       //0文章审核/1三月圈申请
	Status     uint8 `json:"status"`     //0同意/1驳回
	CreateTime int64 `json:"createTime"` //创建时间
}
