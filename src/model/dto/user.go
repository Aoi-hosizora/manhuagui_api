package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
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
				goapidoc.NewProperty("login_ip", "string", true, "user current login ip"),
				goapidoc.NewProperty("last_login_ip", "string", true, "user last login ip"),
				goapidoc.NewProperty("register_time", "string", true, "user register time"),
				goapidoc.NewProperty("last_login_time", "string", true, "user last login time"),
			),

		goapidoc.NewDefinition("ShelfStatusDto", "Shelf status response").
			Properties(
				goapidoc.NewProperty("in", "boolean", true, "manga is in the shelf"),
			),
	)
}

type TokenDto struct {
	Token string `json:"token"` // 登录令牌
}

// 用户信息 vo.User
type UserDto struct {
	Username      string `json:"username"`        // 用户名
	Avatar        string `json:"avatar"`          // 用户头像
	Class         string `json:"class"`           // 会员等级
	Score         int32  `json:"score"`           // 个人成长值
	LoginIP       string `json:"login_ip"`        // 本地登录IP
	LastLoginIP   string `json:"last_login_ip"`   // 上次登录IP
	RegisterTime  string `json:"register_time"`   // 注册时间
	LastLoginTime string `json:"last_login_time"` // 上次登录时间
}

func BuildUserDto(user *vo.User) *UserDto {
	return &UserDto{
		Username:      user.Username,
		Avatar:        user.Avatar,
		Class:         user.Class,
		Score:         user.Score,
		LoginIP:       user.LoginIP,
		LastLoginIP:   user.LastLoginIP,
		RegisterTime:  user.RegisterTime,
		LastLoginTime: user.LastLoginTime,
	}
}

func BuildBuildUserDtos(users []*vo.User) []*UserDto {
	out := make([]*UserDto, len(users))
	for idx, user := range users {
		out[idx] = BuildUserDto(user)
	}
	return out
}

// 书柜状态 vo.ShelfStatus
type ShelfStatusDto struct {
	In bool `json:"in"`
}

func BuildShelfStatusDto(status *vo.ShelfStatus) *ShelfStatusDto {
	return &ShelfStatusDto{
		In: status.In,
	}
}
