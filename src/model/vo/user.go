package vo

// 用户状态
type UserStatus struct {
	Status   int32  `json:"status"`   // 返回状态
	Username string `json:"username"` // 用户名
	Message  int32  `json:"message"`  // 消息数量
	Shelf    int32  `json:"shelf"`    // 书架数量
}

// 用户信息
type User struct {
	Username      string // 用户名
	Avatar        string // 用户头像
	Class         string // 会员等级
	Score         int32  // 个人成长值
	LoginIP       string // 本地登录IP
	LastLoginIP   string // 上次登录IP
	RegisterTime  string // 注册时间
	LastLoginTime string // 上次登录时间
}

// 书柜状态
type ShelfStatus struct {
	In bool // 存在
}
