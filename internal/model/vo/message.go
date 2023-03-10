package vo

import (
	"time"
)

// GitHub Issue
type Issue struct {
	Number   uint64 `json:"number"`
	Comments uint64 `json:"comments"`
}

// GitHub Issue Comment
type IssueComment struct {
	Id        uint64    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GitHub Issue Comment Body
type IssueCommentBody struct {
	MidStr       string               `yaml:"mid"` // empty-able, for overwrite
	Title        string               `yaml:"title"`
	Notification *NotificationContent `yaml:"notification"`
	NewVersion   *NewVersionContent   `yaml:"new_version"`
	CreatedAt    string               `yaml:"created_at"` // empty-able, for overwrite
	UpdatedAt    string               `yaml:"updated_at"` // empty-able, for overwrite
}

// 消息
type Message struct {
	Mid          uint64               // 消息编号
	Title        string               // 消息标题
	Notification *NotificationContent // 通知公告
	NewVersion   *NewVersionContent   // 新版本更新
	CreatedAt    time.Time            // 通知创建时间
	UpdatedAt    time.Time            // 通知更新时间
}

// 消息 - 通知公告
type NotificationContent struct {
	Content     string `yaml:"content"`     // 通知内容
	Dismissible *bool  `yaml:"dismissible"` // 可被跳过
	Link        string `yaml:"link"`        // 关联的链接
}

// 消息 - 新版本更新
type NewVersionContent struct {
	Version     string `yaml:"version"`      // 版本号
	MustUpgrade *bool  `yaml:"must_upgrade"` // 必须更新
	ChangeLogs  string `yaml:"change_logs"`  // 更新日志
	ReleasePage string `yaml:"release_page"` // 发布页面
}

// 最新消息
type LatestMessage struct {
	Notification               *Message // 通知公告
	NewVersion                 *Message // 新版本更新
	NotDismissibleNotification *Message // 通知公告 (不得跳过)
	MustUpgradeNewVersion      *Message // 新版本更新 (必须更新)
}
