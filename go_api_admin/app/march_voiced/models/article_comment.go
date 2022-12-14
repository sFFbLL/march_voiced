package models

import (
	"project/app/march_voiced/models/dto"
	"project/common/global"

	"go.uber.org/zap"
)

type ArticleComment struct {
	ArticleId uint   `json:"article_id" gorm:""`
	Pid       uint   `json:"pid" gorm:""`
	ReplyId   uint   `json:"reply_id" gorm:""`
	CreateBy  uint   `json:"creat_by" gorm:""`
	UpdateBy  uint   `json:"update_by" gorm:""`
	Content   string `json:"content" gorm:""`
	BaseModel
}

type UserInfo struct {
	NickName   string `json:"nick_name"`
	AvatarPath string `json:"avatar_path"`
}

func (co *ArticleComment) TableName() string {
	return `article_comment`
}

func (co *ArticleComment) ArticleCommentCount() (count int64, err error) {
	err = global.Eloquent.Table(co.TableName()).Where("article_id = ? AND is_deleted = 0", co.ArticleId).Count(&count).Error
	if err != nil {
		zap.L().Error("ArticleCommentCount Select failed", zap.Error(err))
		return
	}
	return
}

// AddArticleComment 新增详情页的评论
func (co *ArticleComment) AddArticleComment() (err error) {
	err = global.Eloquent.Create(co).Error
	return
}

//AddArticleCommentMessage 新增详情页的评论消息
func (co *ArticleComment) AddArticleCommentMessage(userId uint, objectId uint) {
	if userId != objectId {
		AddMessage(0, 1, co.ArticleId, userId, co.Content)
	}
}

// DeleteArticleComment 删除详情页的评论
func (co *ArticleComment) DeleteArticleComment() (err error) {
	// 修改is_deleted为1
	err = global.Eloquent.Table(co.TableName()).Where("id in (?) AND is_deleted = ?", co.ID, []byte{0}).Updates(map[string]interface{}{"is_deleted": 1, "update_by": co.UpdateBy}).Error
	return
}

// GetCommentList 先查询size个一级评论的相关信息
func (co *ArticleComment) GetCommentList(p *dto.GetArticleComment) (commentList []*ArticleComment, err error) {
	table := global.Eloquent.Table(co.TableName()).Where("is_deleted=? AND article_id=?", []byte{0}, p.ID)
	err = table.Where("pid=? AND reply_id=?", 0, 0).
		Offset(int((p.Current - 1) * p.Size)).Limit(int(p.Size)).Order("create_time desc").
		Find(&commentList).Error
	return
}

// GetChildSum 查询一级评论下有多少个二级评论
func (co *ArticleComment) GetChildSum(id int) (sum int64, err error) {
	table := global.Eloquent.Table(co.TableName())
	err = table.Where("pid = ? AND is_deleted = ?", id, []byte{0}).Count(&sum).Error
	return
}

// GetChildCommentList 查询父评论下属的childSize个子评论
func (co *ArticleComment) GetChildCommentList(articleId int, size int, pid int) (commentList []*ArticleComment, err error) {

	table := global.Eloquent.Table(co.TableName()).Where("is_deleted=? AND article_id=?", []byte{0}, articleId)
	err = table.Where("pid=?", pid).Limit(size).Order("create_time asc").Find(&commentList).Error
	return

}

// GetUserInfo 获取与评论有关的用户的信息
func (co *ArticleComment) GetUserInfo(id uint) (userInfo *UserInfo, err error) {
	if id == 0 {
		return &UserInfo{
			NickName:   "",
			AvatarPath: "",
		}, nil
	}
	userInfo = new(UserInfo)
	idI := int(id)
	table := global.Eloquent.Table("sys_user").Where("is_deleted=?", []byte{0})
	err = table.Where("id=?", idI).First(userInfo).Error
	return
}

// GetArticleChildComment 查询指定文章下某一条一级评论的size條二级评论
func (co *ArticleComment) GetArticleChildComment(p *dto.GetArticleChildComment) (commentList *[]ArticleComment, err error) {
	commentList = new([]ArticleComment)
	table := global.Eloquent.Table(co.TableName()).Where("is_deleted=?", []byte{0})
	err = table.Where("pid=?", p.ID).Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).Order("create_time asc").Find(commentList).Error
	return
}
