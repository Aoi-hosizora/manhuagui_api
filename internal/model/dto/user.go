package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("TokenDto", "Token response").
			Properties(
				goapidoc.NewProperty("token", "string", true, "access token"),
			),

		goapidoc.NewDefinition("UserDto", "User response").
			Properties(
				goapidoc.NewProperty("username", "string", true, "user name"),
				goapidoc.NewProperty("avatar", "string", true, "user avatar"),
				goapidoc.NewProperty("class", "string", true, "user class (会员等级)"),
				goapidoc.NewProperty("score", "integer#int32", true, "user score (个人成长值)"),
				goapidoc.NewProperty("account_point", "integer#int32", true, "user account point (账户积分)"),
				goapidoc.NewProperty("unread_message_count", "integer#int32", true, "unread message count)"),
				goapidoc.NewProperty("login_ip", "string", true, "user current login ip"),
				goapidoc.NewProperty("last_login_ip", "string", true, "user last login ip"),
				goapidoc.NewProperty("register_time", "string", true, "user register time"),
				goapidoc.NewProperty("last_login_time", "string", true, "user last login time"),
				goapidoc.NewProperty("cumulative_day_count", "string", true, "user cumulative logined day count"),
				goapidoc.NewProperty("total_comment_count", "string", true, "user total comment count"),
			),

		goapidoc.NewDefinition("UsernameDto", "Username response").
			Properties(
				goapidoc.NewProperty("username", "string", true, "username"),
			),

		goapidoc.NewDefinition("ShelfStatusDto", "Shelf status response").
			Properties(
				goapidoc.NewProperty("in", "boolean", true, "manga is in the shelf"),
				goapidoc.NewProperty("count", "integer#int32", true, "manga starred count"),
			),
	)
}

type TokenDto struct {
	Token string `json:"token"` // 登录令牌
}

// 用户信息 vo.User
type UserDto struct {
	Username           string `json:"username"`             // 用户名
	Avatar             string `json:"avatar"`               // 用户头像
	Class              string `json:"class"`                // 会员等级
	Score              int32  `json:"score"`                // 个人成长值
	AccountPoint       int32  `json:"account_point"`        // 账户积分
	UnreadMessageCount int32  `json:"unread_message_count"` // 未读消息条数
	LoginIP            string `json:"login_ip"`             // 本地登录IP
	LastLoginIP        string `json:"last_login_ip"`        // 上次登录IP
	RegisterTime       string `json:"register_time"`        // 注册时间
	LastLoginTime      string `json:"last_login_time"`      // 上次登录时间
	CumulativeDayCount int32  `json:"cumulative_day_count"` // 累计登录天数
	TotalCommentCount  int32  `json:"total_comment_count"`  // 累计评论总数
}

func BuildUserDto(user *vo.User) *UserDto {
	return &UserDto{
		Username:           user.Username,
		Avatar:             user.Avatar,
		Class:              user.Class,
		Score:              user.Score,
		AccountPoint:       user.AccountPoint,
		UnreadMessageCount: user.UnreadMessageCount,
		LoginIP:            user.LoginIP,
		LastLoginIP:        user.LastLoginIP,
		RegisterTime:       user.RegisterTime,
		LastLoginTime:      user.LastLoginTime,
		CumulativeDayCount: user.CumulativeDayCount,
		TotalCommentCount:  user.TotalCommentCount,
	}
}

func BuildBuildUserDtos(users []*vo.User) []*UserDto {
	out := make([]*UserDto, len(users))
	for idx, user := range users {
		out[idx] = BuildUserDto(user)
	}
	return out
}

type UsernameDto struct {
	Username string `json:"username"`
}

// 书柜状态 vo.ShelfStatus
type ShelfStatusDto struct {
	In    bool  `json:"in"`    // 已收藏
	Count int32 `json:"count"` // 收藏用户数
}

func BuildShelfStatusDto(status *vo.ShelfStatus) *ShelfStatusDto {
	return &ShelfStatusDto{
		In:    status.Status == 1,
		Count: status.Total,
	}
}
