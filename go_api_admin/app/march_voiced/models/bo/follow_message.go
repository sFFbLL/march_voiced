package bo

type GetFollowMessage struct {
	Current uint                    `json:"current"` //页码
	Size    uint                    `json:"size"`    //页内大小
	Pages   int                     `json:"pages"`   //总页数
	Total   int64                   `json:"total"`   //总条数
	Records *[]GetFollowMessageData `json:"records"` //消息
}

type GetFollowMessageData struct {
	AvatarPath string `json:"avatarPath" gorm:"column:avatar_path"` //头像
	Nickname   string `json:"nickname" gorm:"column:nick_name"`     //昵称
	IsFollow   uint8  `json:"isFollow"`                             //0未关注1已关注
	Id         int64  `json:"id"`                                   //id
	CreateTime int64  `json:"createTime" gorm:"column:create_time"` //创建时间
}
