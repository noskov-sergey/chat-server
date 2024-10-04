package chats

import (
	"context"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"testing"

	mock_file "github.com/noskov-sergey/chat-server/internal/api/chats/mocks"
	chat "github.com/noskov-sergey/chat-server/pkg/chat_v1"
)

var (
	testChatId      int64 = 77
	testChatIdError       = 0
	testUsernames         = []string{"testName1", "testName2"}
	testUserName          = "testName1"

	testCreateChatReq = &chat.CreateChatRequest{
		Usernames: testUsernames,
	}

	testDeleteChatReq = &chat.DeleteRequest{
		Id: testChatId,
	}

	testCreateMessageRequest = &chat.CreateMessageRequest{
		ChatId:    testChatId,
		From:      testUserName,
		Text:      "testText",
		Timestamp: timestamppb.Now(),
	}
)

type testDeps struct {
	uc     *mock_file.MockUsecase
	client chat.ChatV1Client
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	uc := mock_file.NewMockUsecase(ctrl)

	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	server := grpc.NewServer(grpc.MaxSendMsgSize(500))
	chat.RegisterChatV1Server(server, New(uc))
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalln("error serving server", err)
		}
	}()

	t.Cleanup(func() {
		err := lis.Close()
		if err != nil {
			log.Fatalln("error closing listener", err)
		}
		server.Stop()
	})

	conn, err := grpc.Dial(
		"buf dial",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("error connecting to server", err)
	}

	return &testDeps{
		uc:     uc,
		client: chat.NewChatV1Client(conn),
	}
}
