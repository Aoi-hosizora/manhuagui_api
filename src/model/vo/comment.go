package vo

// 漫画评论列表
type Comments struct {
	CommentIds []string            `json:"commentIds"`
	Comments   map[string]*Comment `json:"comments"`
	Total      int32               `json:"total"`
}

// 漫画评论
type Comment struct {
	Id            uint64            `json:"id"`            // 评论编号
	UserId        uint64            `json:"user_id"`       // 用户编号
	Username      string            `json:"user_name"`     // 用户名称
	Avatar        string            `json:"avatar"`        // 用户头像
	Sex           uint8             `json:"sex"`           // 用户性别
	Content       string            `json:"content"`       // 评论内同
	SupportCount  uint8             `json:"support_count"` // 被赞次数
	ReplyCount    uint8             `json:"reply_count"`   // 回复次数
	ReplyTimeline []*RepliedComment `json:"-"`             // 评论回复
	AddTime       string            `json:"add_time"`      // 评论时间
}

// 被回复的评论
type RepliedComment struct {
	Id           uint64 `json:"id"`            // 评论编号
	UserId       uint64 `json:"user_id"`       // 用户编号
	Username     string `json:"user_name"`     // 用户名称
	Avatar       string `json:"avatar"`        // 用户头像
	Sex          uint8  `json:"sex"`           // 用户性别
	Content      string `json:"content"`       // 评论内同
	SupportCount uint8  `json:"support_count"` // 被赞次数
	ReplyCount   uint8  `json:"reply_count"`   // 回复次数
	AddTime      string `json:"add_time"`      // 评论时间
}
