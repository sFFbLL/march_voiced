package dto

// MessageUnreadDto 是否有未读消息dto
type MessageUnreadDto struct {
	Type uint8 `form:"type"` //1动态，2关注，3其他 不传是全部
}

// UnReadMessageRead 未读消息已读dto
type UnReadMessageRead struct {
	Type uint8 `form:"type"` //1动态，2关注，3其他 不传是全部
}
