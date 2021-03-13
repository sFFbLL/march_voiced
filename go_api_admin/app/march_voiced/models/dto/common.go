package dto

type Paginator struct {
	Current uint `form:"current"`
	Size    uint `form:"size" binding:"lte=100"`
}
