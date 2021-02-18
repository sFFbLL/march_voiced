package service

import (
	"errors"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
)

type ArticleComment struct {
}

// AddArticleComment 新增文章详情页的评论
func (co *ArticleComment) AddArticleComment(userId int, p *dto.AddArticleComment) (err error) {
	comment := new(models.ArticleComment)
	comment.Pid = p.Pid
	comment.Content = p.Content
	comment.ArticleId = p.ID
	comment.ReplyId = p.ReplyId
	comment.UpdateBy = uint(userId)
	comment.CreateBy = uint(userId)
	err = comment.AddArticleComment()
	return
}

// DeleteArticleComment 删除文章详情页的评论
func (co *ArticleComment) DeleteArticleComment(userId int, id int) (err error) {
	comment := new(models.ArticleComment)
	comment.ID = id
	comment.UpdateBy = uint(userId)
	err = comment.DeleteArticleComment()
	return
}

// GetArticleComment 查询文章详情页的评论，一次只返回N条一级评论和3条二级评论
func (co *ArticleComment) GetArticleComment(p *dto.GetArticleComment) (res *bo.GetArticleComment, err error) {
	comment := new(models.ArticleComment)

	if p.Size == 0 || p.Current == 0 {
		return nil, errors.New("参数缺失")
	} else {

		res = new(bo.GetArticleComment)
		var SignalComments []bo.SignalArticleComment
		var ChildComments []bo.ArticleComment

		commentList, _ := comment.GetCommentList(p)
		// 这里面获取的是二级评论的信息和一级评论的用户的信息
		for i := 0; i < len(commentList); i++ {
			userInfo, err := comment.GetUserInfo(commentList[i].CreateBy)
			userReplyInfo, err := comment.GetUserInfo(commentList[i].ReplyId)
			if err != nil {
				return nil, errors.New("获取用户信息失败")
			}
			commentChild, err := comment.GetChildCommentList(p.ID, commentList[i].ID)
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
				ChildComments = append(ChildComments, bo.ArticleComment{
					Id:         commentChild.ID,
					Content:    commentChild.Content,
					CreateTime: commentChild.CreateTime,
					ArticleCreate: bo.ArticleCreate{
						CreateByName: userChildInfo.NickName,
						IdAvatar:     userChildInfo.AvatarPath,
						CreateBy:     int(commentChild.CreateBy),
					},
					ArticleReply: bo.ArticleReply{
						ReplyId:     int(commentChild.ReplyId),
						ReplyName:   userChildReplyInfo.NickName,
						ReplyAvatar: userChildReplyInfo.AvatarPath,
					},
				})
			}
			SignalComments = append(SignalComments, bo.SignalArticleComment{
				ArticleComment: bo.ArticleComment{
					Id:         commentList[i].ID,
					Content:    commentList[i].Content,
					CreateTime: commentList[i].CreateTime,
					ArticleCreate: bo.ArticleCreate{
						CreateBy:     int(commentList[i].CreateBy),
						CreateByName: userInfo.NickName,
						IdAvatar:     userInfo.AvatarPath,
					},
					ArticleReply: bo.ArticleReply{
						ReplyId:     int(commentList[i].ReplyId),
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

// GetArticleChildComment 查询文章详情页的某个一级评论下的所有子评论
func (co *ArticleComment) GetArticleChildComment(p *dto.GetArticleChildComment) (*bo.ChildrenComment, error) {
	comment := new(models.ArticleComment)
	res := new(bo.ChildrenComment)
	if p.Size == 0 || p.Current == 0 {
		return nil, errors.New("参数缺失")
	} else {
		commentList, err := comment.GetArticleChildComment(p)
		if err != nil {
			return nil, err
		}
		var getChildrenComment []bo.ArticleComment
		for _, childComment := range *commentList {

			userChildInfo, err := comment.GetUserInfo(childComment.CreateBy)
			replyUserChildInfo, err := comment.GetUserInfo(childComment.ReplyId)
			if err != nil {
				return nil, errors.New("获取用户信息失败")
			}
			getChildrenComment = append(getChildrenComment, bo.ArticleComment{
				Id:         childComment.ID,
				Content:    childComment.Content,
				CreateTime: childComment.CreateTime,
				ArticleCreate: bo.ArticleCreate{
					CreateBy:     int(childComment.CreateBy),
					CreateByName: userChildInfo.NickName,
					IdAvatar:     userChildInfo.AvatarPath,
				},
				ArticleReply: bo.ArticleReply{
					ReplyId:     int(childComment.ReplyId),
					ReplyName:   replyUserChildInfo.NickName,
					ReplyAvatar: replyUserChildInfo.AvatarPath,
				},
			})
		}
		res.Comment = getChildrenComment
	}
	return res, nil
}