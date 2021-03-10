package dto

type WxTokenMessages struct {
	Ticket string	`json:"Ticket"`
	Url string		`json:"url"`
	Strdata string   `json:"strdata"`
	Errmsg string `json:"errmsg"`
	Errcode int	`json:"errcode"`
}
//InsertUserDto 新增用户
type InsertUserDto struct {
	//DeptId   int    `json:"deptid" binding:"required"`   //部门
	//ID       int    `json:"id"`                          //id (目前不传)
	//Phone    string  `json:"phone" `    //手机号
	//Email    string `json:"email"`                       //邮箱 (目前不传)
	//NickName string `json:"nickName" ` //昵称
	UserName string `json:"username" binding:"required"` //用户名
	//Gender   string `json:"gender" `   //性别
	OpenId   string `json:"openid" binding:"required"`
	Password string `json:"password" binging:"required"`
	//Enabled  string `json:"enabled" binding:"required"`  //是否激活
	//Jobs     []int  `json:"jobs" binding:"required"`     //职位
	//Roles    []int  `json:"roles" binding:"required"`    //角色
}

//modInformation 修改用户信息
type ModInformationDto struct {
	NickName string `json:"nickname"`
	Sex int `json:"sex"`
	Signature string `json:"signature"`
	AvatarPath string `json:"avatarPath"`
}






