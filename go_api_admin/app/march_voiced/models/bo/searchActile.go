package bo


// SearchActileBo 返回文章搜索数据
type SearchActileBo struct {
	Content		string	`json:"content"`
	Title		string	`json:"title"`
	CreatTime	string	`json:"creatTime"`
	CommentTotal string `json:"commentTotal"`
	FavourTotal	string	`json:"favourTotal"`
	ContentId	string  `json:"contentId"`
	Nickname	string	`json:"nickname"`
	CreateID    int 	`json:"createBy"`
}

//
type SearchUserBo struct {
	AvatarPath string `json:"avatarPath"`
	NickName string `json:"nickName"`
	UserId int 	`json:"userId"`
	Signature string `json:"signature"`
	IsFollow int `json:"isFollow"`
}
