package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("CommentDto", "Comment response").
			Properties(
				goapidoc.NewProperty("cid", "integer#int64", true, "comment id"),
				goapidoc.NewProperty("uid", "integer#int64", true, "user id"),
				goapidoc.NewProperty("username", "string", true, "username"),
				goapidoc.NewProperty("avatar", "string", true, "user avatar"),
				goapidoc.NewProperty("gender", "integer#int32", true, "user gender, 1: Male, 2: Female"),
				goapidoc.NewProperty("content", "string", true, "comment content"),
				goapidoc.NewProperty("like_count", "integer#int32", true, "comment liked count"),
				goapidoc.NewProperty("reply_count", "integer#int32", true, "comment reply count"),
				goapidoc.NewProperty("reply_timeline", "RepliedCommentDto[]", true, "comment reply timeline"),
				goapidoc.NewProperty("comment_time", "string", true, "comment create time"),
			),

		goapidoc.NewDefinition("RepliedCommentDto", "Replied comment response").
			Properties(
				goapidoc.NewProperty("cid", "integer#int64", true, "comment id"),
				goapidoc.NewProperty("uid", "integer#int64", true, "user id"),
				goapidoc.NewProperty("username", "string", true, "username"),
				goapidoc.NewProperty("avatar", "string", true, "user avatar"),
				goapidoc.NewProperty("gender", "integer#int32", true, "user gender, 1: Male, 2: Female"),
				goapidoc.NewProperty("content", "string", true, "comment content"),
				goapidoc.NewProperty("like_count", "integer#int32", true, "comment liked count"),
				goapidoc.NewProperty("reply_count", "integer#int32", true, "comment reply count"),
				goapidoc.NewProperty("comment_time", "string", true, "comment create time"),
			),
	)
}

// 漫画评论 vo.Comment
type CommentDto struct {
	Cid           uint64               `json:"cid"`            // 评论编号
	Uid           uint64               `json:"uid"`            // 用户编号
	Username      string               `json:"username"`       // 用户名
	Avatar        string               `json:"avatar"`         // 用户头像
	Gender        uint8                `json:"gender"`         // 用户性别
	Content       string               `json:"content"`        // 评论内同
	LikeCount     uint8                `json:"like_count"`     // 被赞次数
	ReplyCount    uint8                `json:"reply_count"`    // 回复次数
	ReplyTimeline []*RepliedCommentDto `json:"reply_timeline"` // 回复列表
	CommentTime   string               `json:"comment_time"`   // 评论时间
}

func BuildCommentDto(comment *vo.Comment) *CommentDto {
	return &CommentDto{
		Cid:           comment.Id,
		Uid:           comment.UserId,
		Username:      comment.Username,
		Avatar:        comment.Avatar,
		Gender:        comment.Sex,
		Content:       comment.Content,
		LikeCount:     comment.SupportCount,
		ReplyCount:    comment.ReplyCount,
		ReplyTimeline: BuildRepliedCommentDtos(comment.ReplyTimeline),
		CommentTime:   comment.AddTime,
	}
}

func BuildCommentDtos(comments []*vo.Comment) []*CommentDto {
	out := make([]*CommentDto, len(comments))
	for idx, comment := range comments {
		out[idx] = BuildCommentDto(comment)
	}
	return out
}

// 被回复的评论 vo.RepliedComment
type RepliedCommentDto struct {
	Cid         uint64 `json:"cid"`          // 评论编号
	Uid         uint64 `json:"uid"`          // 用户编号
	Username    string `json:"username"`     // 用户名
	Avatar      string `json:"avatar"`       // 用户头像
	Gender      uint8  `json:"gender"`       // 用户性别
	Content     string `json:"content"`      // 评论内同
	LikeCount   uint8  `json:"like_count"`   // 被赞次数
	ReplyCount  uint8  `json:"reply_count"`  // 回复次数
	CommentTime string `json:"comment_time"` // 评论时间
}

func BuildRepliedCommentDto(comment *vo.RepliedComment) *RepliedCommentDto {
	return &RepliedCommentDto{
		Cid:         comment.Id,
		Uid:         comment.UserId,
		Username:    comment.Username,
		Avatar:      comment.Avatar,
		Gender:      comment.Sex,
		Content:     comment.Content,
		LikeCount:   comment.SupportCount,
		ReplyCount:  comment.ReplyCount,
		CommentTime: comment.AddTime,
	}
}

func BuildRepliedCommentDtos(comments []*vo.RepliedComment) []*RepliedCommentDto {
	if comments == nil {
		return []*RepliedCommentDto{}
	}
	out := make([]*RepliedCommentDto, len(comments))
	for idx, comment := range comments {
		out[idx] = BuildRepliedCommentDto(comment)
	}
	return out
}
