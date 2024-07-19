package converter

import (
	"github.com/noskov-sergey/chat-server/internal/model"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
)

func ToServiceFromMessages(req *desc.SendMessageRequest) model.Message {
	return model.Message{
		Username: req.GetFrom(),
		ChatId:   int(req.GetChatId()),
		Text:     req.GetText(),
	}
}
