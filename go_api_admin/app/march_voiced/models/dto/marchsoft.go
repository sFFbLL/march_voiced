package dto

type MarchPass struct {
	Id     uint   `json:"id" binding:"required"`
	Status *uint8 `json:"status" binding:"required,lte=1"`
}

type ApplyMarchPaginator struct {
	Current   uint   `form:"current"`
	Size      uint   `form:"size"`
	EndTime   uint   `form:"endTime"`   //结束时间
	StartTime uint   `form:"startTime"` //创建时间
	Nickname  string `form:"nickname"`
}

type InsertMarchSoft struct {
	Content   string   `json:"content" binding:"required"`
	Image     string   `json:"image"`
	ImageList []string `json:"imageList"`
}

type SelectMarchListById struct {
	Paging
	ID uint `form:"id"`
}
