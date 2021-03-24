package models

import (
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type MarchsoftComment struct {
	BaseModel
	Content     string `json:"content" gorm:"not null"`
	CreateBy    uint   `json:"create_by" gorm:"not null;index"`    //创建人
	UpdateBy    uint   `json:"update_by" gorm:"not null"`          //更新人
	MarchsoftId uint   `json:"marchsoft_id" gorm:"not null;index"` //三月圈id
	Pid         uint   `json:"pid" gorm:"not null"`                //上级评论id
	ReplyId     uint   `json:"reply_id" gorm:"not null;index"`     //回复人id
}

func (co *MarchsoftComment) TableName() string {
	return `marchsoft_comment`
}

// AddMarchsoftComment 新增评论
func (co *MarchsoftComment) AddMarchsoftComment() (err error) {
	err = global.Eloquent.Create(co).Error
	return
}

// AddMarchsoftCommentMessage 新增评论消息
func (co *MarchsoftComment) AddMarchsoftCommentMessage(SubjectId uint, objectId uint) {
	if SubjectId != objectId {
		AddMessage(0, 1, co.MarchsoftId, SubjectId, co.Content)
	}
}

// DeleteMarchsoftComment 删除详情页的评论
func (co *MarchsoftComment) DeleteMarchsoftComment() (err error) {
	// 修改is_deleted为1
	err = global.Eloquent.Table(co.TableName()).Where("id in (?) AND is_deleted = ?", co.ID, []byte{0}).Updates(map[string]interface{}{"is_deleted": 1, "update_by": co.UpdateBy}).Error
	return
}

// GetCommentList 先查询size个一级评论的相关信息
func (co *MarchsoftComment) GetCommentList(p *dto.GetMarchsoftComment) (commentList []*MarchsoftComment, err error) {
	table := global.Eloquent.Table(co.TableName()).Where("is_deleted=? AND marchsoft_id=?", []byte{0}, p.ID)
	err = table.Where("pid=? AND reply_id=?", 0, 0).
		Offset(int((p.Current - 1) * p.Size)).Limit(int(p.Size)).Order("create_time desc").
		Find(&commentList).Error
	return
}

// GetChildSum 查询一级评论下有多少个二级评论
func (co *MarchsoftComment) GetChildSum(id int) (sum int64, err error) {
	table := global.Eloquent.Table(co.TableName())
	err = table.Where("pid = ? AND is_deleted = ?", id, []byte{0}).Count(&sum).Error
	return
}

// GetChildCommentList 查询父评论下属的childSize个子评论
func (co *MarchsoftComment) GetChildCommentList(marchsoftId uint, size uint, pid int) (commentList []*MarchsoftComment, err error) {
	table := global.Eloquent.Table(co.TableName()).Where("is_deleted=? AND marchsoft_id=?", []byte{0}, marchsoftId)
	err = table.Where("pid=?", pid).Limit(int(size)).Order("create_time asc").Find(&commentList).Error
	return
}

// GetUserInfo 获取与评论有关的用户的信息
func (co *MarchsoftComment) GetUserInfo(id uint) (userInfo *UserInfo, err error) {
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

// GetMarchsoftChildComment 查询指定文章下某一条一级评论的size条二级评论
func (co *MarchsoftComment) GetMarchsoftChildComment(p *dto.GetMarchsoftChildComment) (commentList *[]MarchsoftComment, err error) {
	commentList = new([]MarchsoftComment)
	table := global.Eloquent.Table(co.TableName()).Where("is_deleted=?", []byte{0})
	err = table.Where("pid=?", p.ID).Limit(int(p.Size)).Offset(int((p.Current - 1) * p.Size)).Order("create_time asc").Find(commentList).Error
	return
}

func (co *MarchsoftComment) GetMarchSoftCommentCount() (count int64, err error) {
	err = global.Eloquent.Table(co.TableName()).Where("marchsoft_id = ?", co.MarchsoftId).Count(&count).Error
	return
}
