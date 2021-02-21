package bo

//返回用户登录详细列表
type MarchLoginData struct {
	Token string     `json:"token"`
	Status int 		 `json:"status"`
	OpenId string    `json:"openid"`
}

//返回补全信息后的token
type LoginInfoData struct {
	Token string     `json:"token"`
}

//返回查询用户信息数据
type UserInfo struct {
	NickName string `json:"nickname"`
	Sex   int  `json:"sex"`
	Signature string `json:"signature"`
	AvatarPath string `json:"avatarPath"`
}

//返回用户登录详细列表
type LoginData struct {
	Token string     `json:"token"`
	User  *LoginUser `json:"user"`
}

type LoginUser struct {
	DataScopes []int       `json:"dataScopes"`
	Roles      []string    `json:"roles"`
	User       *RecordUser `json:"user"`
}

type RecordUser struct {
	*RecordUserHalf
	*RoleDeptJobBool
}

//order 排序规则
type Order struct {
	Column string `json:"column"`
	Asc    string `json:"asc"`
}

//返回用户详细列表
type UserInfoListBo struct {
	Current          int           `json:"current"`
	CountID          int           `json:"count_id"`
	MaxLimit         int           `json:"maxLimit"`
	Pages            int           `json:"pages"`
	SearchCount      bool          `json:"searchCount"`
	Size             int           `json:"size"`
	Total            int           `json:"total"`
	HitCount         bool          `json:"hitCount"`
	OptimizeCountSql bool          `json:"optimizeCountSql"`
	Orders           []Order       `json:"orders"`
	Records          []*RecordUser `json:"records"` //记录查询
}

//返回用户详细列表
type RecordUserHalf struct {
	Id           int    `json:"id"`
	DeptId       int    `json:"deptId"`
	CreateBy     int    `json:"createBy"`
	UpdateBy     int    `json:"updatedBy"`
	PwdResetTime int64  `json:"pwdResetTime"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	AvatarName   string `json:"avatarName"`
	AvatarPath   string `json:"avatarPath"`
	Email        string `json:"email"`
	NickName     string `json:"nickName"`
	Phone        string `json:"phone"`
	Username     string `json:"username"`
}
//返回用户详细列表
type RoleDeptJobBool struct {
	Enabled bool        `json:"enabled"`
	Gender  bool        `json:"gender"`
	Jobs    []*Job      `json:"jobs"`
	Role    []*Role     `json:"role"`
	Dept    *DeptCommon `json:"dept"`
}
type DeptCommon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Job struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID        int    `json:"id"`
	Level     int    `json:"level"`
	Name      string `json:"name"`
	DataScope string `json:"dataScope"`
}

// OnlineUser 用户线上数据
type OnlineUser struct {
	LoginTime     int64  `json:"loginTime"`     //登录时间
	LoginLocation string `json:"loginLocation"` // 归属地
	Browser       string `json:"browser"`       // 浏览器
	Dept          string `json:"dept"`          //部门
	Ip            string `json:"ip"`            //ip地址
	Nickname      string `json:"nickname"`      //昵称
	Username      string `json:"username"`      //用户名
	Token         string `json:"key"`           // token
}