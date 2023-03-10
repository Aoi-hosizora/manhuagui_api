package object

// 用户状态
type UserStatus struct {
	Status   int32  `json:"status"`   // 返回状态
	Username string `json:"username"` // 用户名
	Message  int32  `json:"message"`  // 消息数量
	Shelf    int32  `json:"shelf"`    // 书架数量
}

// 用户信息
type User struct {
	Username           string // 用户名
	Avatar             string // 用户头像
	Class              string // 会员等级
	Score              int32  // 个人成长值
	AccountPoint       int32  // 账户积分
	UnreadMessageCount int32  // 未读消息条数
	LoginIP            string // 本地登录IP
	LastLoginIP        string // 上次登录IP
	RegisterTime       string // 注册时间
	LastLoginTime      string // 上次登录时间
	CumulativeDayCount int32  // 累计登录天数
	TotalCommentCount  int32  // 累计评论总数
}

// 书柜状态
type ShelfStatus struct {
	Status int32 `json:"status"` // 已收藏
	Total  int32 `json:"total"`  // 收藏用户数
}
