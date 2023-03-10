package dto

import (
	"github.com/Aoi-hosizora/ahlib/xpointer"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/vo"
	"time"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("MessageDto", "Message response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "message id"),
				goapidoc.NewProperty("title", "string", true, "message title"),
				goapidoc.NewProperty("notification", "NotificationContentDto", true, "notification content").AllowEmpty(true),
				goapidoc.NewProperty("new_version", "NewVersionContentDto", true, "new version content").AllowEmpty(true),
				goapidoc.NewProperty("created_at", "string#date-time", true, "created time"),
				goapidoc.NewProperty("updated_at", "string#date-time", true, "updated time"),
			),

		goapidoc.NewDefinition("NotificationContentDto", "Notification content response").
			Properties(
				goapidoc.NewProperty("content", "string", true, "notification content"),
				goapidoc.NewProperty("dismissible", "boolean", true, "flag if the message is dismissible"),
				goapidoc.NewProperty("link", "boolean", true, "link associated to the message"),
			),

		goapidoc.NewDefinition("NewVersionContentDto", "New version content response").
			Properties(
				goapidoc.NewProperty("version", "string", true, "version string"),
				goapidoc.NewProperty("must_upgrade", "boolean", true, "flag if must be upgrade"),
				goapidoc.NewProperty("change_logs", "string", true, "version change logs"),
				goapidoc.NewProperty("release_page", "string", true, "version release page"),
			),

		goapidoc.NewDefinition("LatestMessageDto", "Latest message response").
			Properties(
				goapidoc.NewProperty("notification", "MessageDto", true, "latest notification message"),
				goapidoc.NewProperty("new_version", "MessageDto", true, "latest new version message"),
				goapidoc.NewProperty("not_dismissible_notification", "MessageDto", true, "latest notification message (not dismissible)"),
				goapidoc.NewProperty("must_upgrade_new_version", "MessageDto", true, "latest new version message (must upgrade)"),
			),
	)
}

// 消息
type MessageDto struct {
	Mid          uint64                  `json:"mid"`          // 消息编号
	Title        string                  `json:"title"`        // 消息标题
	Notification *NotificationContentDto `json:"notification"` // 通知公告内容
	NewVersion   *NewVersionContentDto   `json:"new_version"`  // 新版本更新内容
	CreatedAt    time.Time               `json:"created_at"`   // 通知创建时间
	UpdatedAt    time.Time               `json:"updated_at"`   // 通知更新时间
}

func BuildMessageDto(message *vo.Message) *MessageDto {
	if message == nil {
		return nil
	}
	return &MessageDto{
		Mid:          message.Mid,
		Title:        message.Title,
		Notification: BuildNotificationContentDto(message.Notification),
		NewVersion:   BuildNewVersionContentDto(message.NewVersion),
		CreatedAt:    message.CreatedAt,
		UpdatedAt:    message.UpdatedAt,
	}
}

func BuildMessageDtos(messages []*vo.Message) []*MessageDto {
	out := make([]*MessageDto, len(messages))
	for i, message := range messages {
		out[i] = BuildMessageDto(message)
	}
	return out
}

// 消息 - 通知公告
type NotificationContentDto struct {
	Content     string `json:"content"`     // 通知内容
	Dismissible bool   `json:"dismissible"` // 可被跳过
	Link        string `json:"link"`        // 关联的链接
}

func BuildNotificationContentDto(notification *vo.NotificationContent) *NotificationContentDto {
	if notification == nil {
		return nil
	}
	return &NotificationContentDto{
		Content:     notification.Content,
		Dismissible: xpointer.BoolVal(notification.Dismissible, true),
		Link:        notification.Link,
	}
}

// 消息 - 新版本更新
type NewVersionContentDto struct {
	Version     string `json:"version"`      // 版本号
	MustUpgrade bool   `json:"must_upgrade"` // 必须更新
	ChangeLogs  string `json:"change_logs"`  // 更新日志
	ReleasePage string `json:"release_page"` // 发布页面
}

func BuildNewVersionContentDto(newVersion *vo.NewVersionContent) *NewVersionContentDto {
	if newVersion == nil {
		return nil
	}
	return &NewVersionContentDto{
		Version:     newVersion.Version,
		MustUpgrade: xpointer.BoolVal(newVersion.MustUpgrade, false),
		ChangeLogs:  newVersion.ChangeLogs,
		ReleasePage: newVersion.ReleasePage,
	}
}

// 最新消息
type LatestMessageDto struct {
	Notification               *MessageDto `json:"notification"`                 // 通知公告
	NewVersion                 *MessageDto `json:"new_version"`                  // 新版本更新
	NotDismissibleNotification *MessageDto `json:"not_dismissible_notification"` // 通知公告 (不得跳过)
	MustUpgradeNewVersion      *MessageDto `json:"must_upgrade_new_version"`     // 新版本更新 (必须更新)
}

func BuildLatestMessageDto(message *vo.LatestMessage) *LatestMessageDto {
	return &LatestMessageDto{
		Notification:               BuildMessageDto(message.Notification),
		NewVersion:                 BuildMessageDto(message.NewVersion),
		NotDismissibleNotification: BuildMessageDto(message.NotDismissibleNotification),
		MustUpgradeNewVersion:      BuildMessageDto(message.MustUpgradeNewVersion),
	}
}
