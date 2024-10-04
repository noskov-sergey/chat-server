package converter

import (
	"github.com/noskov-sergey/chat-server/internal/model"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
)

func ToServiceFromChats(req *desc.CreateChatRequest) model.Users {
	return model.Users{
		Usernames: req.GetUsernames(),
	}
}
