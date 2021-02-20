package bo

type ApplyMarchUser struct {
	Total   uint                     `json:"total"`   //总页数
	Count   int64                    `json:"count"`   //总条数
	Message *[]ApplyMarchUserMessage `json:"message"` //消息
}

type ApplyMarchUserMessage struct {
	UserId     uint   `json:"userId"`     //用户id
	Nickname   string `json:"nickname"`   //昵称
	AvatarPath string `json:"avatarPath"` //头像
	Phone      string `json:"phone"`      //手机号
	Dept       string `json:"dept"`       //部门
	CreateTime int64  `json:"createTime"` //创建时间
	Status     uint8  `json:"status"`     //2审核中
}
