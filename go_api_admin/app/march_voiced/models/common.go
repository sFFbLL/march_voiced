package models

// AddMessage
/*
t=0文章 t=1三月圈
articleId 文章或三月圈id
userId 消息创建人id
t 消息类型 0文章 1三月圈
status 消息状态 文章：0发布、1评论、2收藏、3转载、4赞 三月圈：0发布、1评论、2，3，4三种赞
comment 如果是评论需要传评论内容
*/
func AddMessage(t uint8, status uint8, articleId uint, userId uint, comment string)  {
	message := new(Message)
	message.Type = t
	message.Status = status
	message.Comment = comment
	_ = message.Add(articleId, userId)
}

// AddSysMessage
/*
t=0文章 t=1三月圈
status 0驳回 1待审核 2通过
*/
func AddSysMessage(t uint8, status uint8, userId uint)  {
	sysMessage := new(SysMessage)
	sysMessage.Type = t
	sysMessage.Status = status
	sysMessage.CreateBy = userId
	sysMessage.UpdateBy = userId
	sysMessage.Add()
}