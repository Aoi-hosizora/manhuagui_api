package controller

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/exception"
	"github.com/Aoi-hosizora/manhuagui-api/src/common/result"
	"github.com/Aoi-hosizora/manhuagui-api/src/config"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/src/provide/sn"
	"github.com/Aoi-hosizora/manhuagui-api/src/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddRoutePaths(
		goapidoc.NewRoutePath("GET", "/v1/message", "Get all messages").
			Tags("Message").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MessageDto>>")),

		goapidoc.NewRoutePath("GET", "/v1/message/latest", "Get latest message").
			Tags("Message").
			Responses(goapidoc.NewResponse(200, "_Result<LatestMessageDto>")),
	)
}

type MessageController struct {
	config         *config.Config
	messageService *service.MessageService
}

func NewMessageController() *MessageController {
	return &MessageController{
		config:         xdi.GetByNameForce(sn.SConfig).(*config.Config),
		messageService: xdi.GetByNameForce(sn.SMessageService).(*service.MessageService),
	}
}

// GET /v1/message
func (m *MessageController) GetMessages(c *gin.Context) *result.Result {
	messages, err := m.messageService.GetAllMessages(m.config.Message.GitHubToken)
	if err != nil {
		return result.Error(exception.GetMessageError).SetError(err, c)
	}
	res := dto.BuildMessageDtos(messages)
	return result.Ok().SetPage(0, int32(len(messages)), int32(len(messages)), res)
}

// GET /v1/message/latest
func (m *MessageController) GetLatestMessage(c *gin.Context) *result.Result {
	latestMessage, err := m.messageService.GetLatestMessage(m.config.Message.GitHubToken)
	if err != nil {
		return result.Error(exception.GetMessageError).SetError(err, c)
	}
	res := dto.BuildLatestMessageDto(latestMessage)
	return result.Ok().SetData(res)
}
