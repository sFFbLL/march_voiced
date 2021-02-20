package bo

type ApplyMarchUser struct {
	Current uint                     `json:"current"` //页码
	Size    uint                     `json:"size"`    //页内大小
	Pages   int                      `json:"pages"`   //总页数
	Total   int64                    `json:"total"`   //总条数
	Records *[]ApplyMarchUserMessage `json:"records"` //消息
}

type ApplyMarchUserMessage struct {
	UserId     uint   `json:"userId" gorm:"column:id"`                    //用户id
	Nickname   string `json:"nickname" gorm:"column:nick_name"`           //昵称
	AvatarPath string `json:"avatarPath" gorm:"column:avatar_path"`       //头像
	Phone      string `json:"phone"`                                      //手机号
	CreateTime int64  `json:"createTime" gorm:"column:march_update_time"` //创建时间
	Dept       string `json:"dept" gorm:"column:name"`                    //部门
	Status     uint8  `json:"status" gorm:"column:is_march"`              //2审核中
}
