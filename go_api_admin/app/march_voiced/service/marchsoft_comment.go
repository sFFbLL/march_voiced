package service

import (
	"errors"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
)

type MarchsoftComment struct {
}

// AddMarchsoftComment 新增文章详情页的评论
func (mc *MarchsoftComment) AddMarchsoftComment(userId uint, p *dto.AddMarchsoftComment) (err error) {
	comment := new(models.MarchsoftComment)
	comment.Pid = p.Pid
	comment.Content = p.Content
	comment.MarchsoftId = p.ID
	comment.ReplyId = p.ReplyId
	comment.UpdateBy = userId
	comment.CreateBy = userId
	err = comment.AddMarchsoftComment()
	go comment.AddMarchsoftCommentMessage(int(userId))
	return
}

// DeleteMarchsoftComment 删除文章详情页的评论
func (mc *MarchsoftComment) DeleteMarchsoftComment(userId int, id int) (err error) {
	comment := new(models.MarchsoftComment)
	comment.ID = id
	comment.UpdateBy = uint(userId)
	err = comment.DeleteMarchsoftComment()
	return
}

// GetMarchsoftComment 查询文章详情页的评论，一次只返回N条一级评论和3条二级评论
func (mc *MarchsoftComment) GetMarchsoftComment(p *dto.GetMarchsoftComment) (res *bo.GetMarchsoftComment, err error) {
	comment := new(models.MarchsoftComment)

	if p.Size == 0 || p.Current == 0 {
		return nil, errors.New("参数缺失")
	} else {

		res = new(bo.GetMarchsoftComment)
		var SignalComments []bo.SignalMarchsoftComment
		var ChildComments []bo.MarchsoftComment

		commentList, _ := comment.GetCommentList(p)
		// 这里面获取的是二级评论的信息和一级评论的用户的信息
		for i := 0; i < len(commentList); i++ {
			userInfo, err := comment.GetUserInfo(commentList[i].CreateBy)
			userReplyInfo, err := comment.GetUserInfo(commentList[i].ReplyId)
			if err != nil {
				return nil, errors.New("获取用户信息失败")
			}
			commentChild, err := comment.GetChildCommentList(p.ID, p.ChildSize, commentList[i].ID)
			if err != nil {
				return nil, errors.New("查询子评论失败")
			}
			// 这里面获取的是二级评论的用户的信息
			for _, commentChild := range commentChild {
				userChildInfo, err := comment.GetUserInfo(commentChild.CreateBy)
				userChildReplyInfo, err := comment.GetUserInfo(commentChild.ReplyId)
				if err != nil {
					return nil, errors.New("获取参数失败")
				}
				ChildComments = append(ChildComments, bo.MarchsoftComment{
					Id:         commentChild.ID,
					Content:    commentChild.Content,
					CreateTime: commentChild.CreateTime,
					MarchsoftCreate: bo.MarchsoftCreate{
						CreateByName: userChildInfo.NickName,
						IdAvatar:     userChildInfo.AvatarPath,
						CreateBy:     commentChild.CreateBy,
					},
					MarchsoftReply: bo.MarchsoftReply{
						ReplyId:     commentChild.ReplyId,
						ReplyName:   userChildReplyInfo.NickName,
						ReplyAvatar: userChildReplyInfo.AvatarPath,
					},
				})
			}
			SignalComments = append(SignalComments, bo.SignalMarchsoftComment{
				MarchsoftComment: bo.MarchsoftComment{
					Id:         commentList[i].ID,
					Content:    commentList[i].Content,
					CreateTime: commentList[i].CreateTime,
					MarchsoftCreate: bo.MarchsoftCreate{
						CreateBy:     commentList[i].CreateBy,
						CreateByName: userInfo.NickName,
						IdAvatar:     userInfo.AvatarPath,
					},
					MarchsoftReply: bo.MarchsoftReply{
						ReplyId:     commentList[i].ReplyId,
						ReplyName:   userReplyInfo.NickName,
						ReplyAvatar: userReplyInfo.AvatarPath,
					},
				},
				ChildComments: ChildComments,
			})
		}
		res.CommentSum = SignalComments
	}
	return res, nil
}

// GetMarchsoftChildComment 查询文章详情页的某个一级评论下的size条子评论
func (mc *MarchsoftComment) GetMarchsoftChildComment(p *dto.GetMarchsoftChildComment) (*bo.ChildrenMarchsoftComment, error) {
	comment := new(models.MarchsoftComment)
	res := new(bo.ChildrenMarchsoftComment)
	if p.Size == 0 || p.Current == 0 {
		return nil, errors.New("参数缺失")
	} else {
		commentList, err := comment.GetMarchsoftChildComment(p)
		if err != nil {
			return nil, err
		}
		var getChildrenComment []bo.MarchsoftComment
		for _, childComment := range *commentList {

			userChildInfo, err := comment.GetUserInfo(childComment.CreateBy)
			replyUserChildInfo, err := comment.GetUserInfo(childComment.ReplyId)
			if err != nil {
				return nil, errors.New("获取用户信息失败")
			}
			getChildrenComment = append(getChildrenComment, bo.MarchsoftComment{
				Id:         childComment.ID,
				Content:    childComment.Content,
				CreateTime: childComment.CreateTime,
				MarchsoftCreate: bo.MarchsoftCreate{
					CreateBy:     childComment.CreateBy,
					CreateByName: userChildInfo.NickName,
					IdAvatar:     userChildInfo.AvatarPath,
				},
				MarchsoftReply: bo.MarchsoftReply{
					ReplyId:     childComment.ReplyId,
					ReplyName:   replyUserChildInfo.NickName,
					ReplyAvatar: replyUserChildInfo.AvatarPath,
				},
			})
		}
		res.Comment = getChildrenComment
	}
	return res, nil
}
