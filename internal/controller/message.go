package controller

import (
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/dto"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/errno"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/result"
	"github.com/Aoi-hosizora/manhuagui-api/internal/service"
	"github.com/gin-gonic/gin"
)

func init() {
	goapidoc.AddOperations(
		goapidoc.NewGetOperation("/v1/message", "Get all messages").
			Tags("Message").
			Responses(goapidoc.NewResponse(200, "_Result<_Page<MessageDto>>")),

		goapidoc.NewGetOperation("/v1/message/latest", "Get latest message").
			Tags("Message").
			Responses(goapidoc.NewResponse(200, "_Result<LatestMessageDto>")),
	)
}

type MessageController struct {
	messageService *service.MessageService
}

func NewMessageController() *MessageController {
	return &MessageController{
		messageService: xmodule.MustGetByName(sn.SMessageService).(*service.MessageService),
	}
}

// GET /v1/message
func (m *MessageController) GetMessages(c *gin.Context) *result.Result {
	messages, err := m.messageService.GetAllMessages()
	if err != nil {
		return result.Error(errno.GetMessageError).SetError(err, c)
	}

	res := dto.BuildMessageDtos(messages)
	return result.Ok().SetPage(0, int32(len(messages)), int32(len(messages)), res)
}

// GET /v1/message/latest
func (m *MessageController) GetLatestMessage(c *gin.Context) *result.Result {
	latestMessage, err := m.messageService.GetLatestMessage()
	if err != nil {
		return result.Error(errno.GetMessageError).SetError(err, c)
	}

	res := dto.BuildLatestMessageDto(latestMessage)
	return result.Ok().SetData(res)
}
