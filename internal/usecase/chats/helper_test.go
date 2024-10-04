package chats

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/noskov-sergey/chat-server/internal/model"
	mock_file "github.com/noskov-sergey/chat-server/internal/usecase/chats/mocks"
	"github.com/noskov-sergey/platform-common/pkg/db"
)

var (
	testId      = 12
	testUser    = "testName"
	testIdError = 0

	testModelUsers = model.Users{
		Usernames: []string{"testName", "testName2", "testName3"},
	}

	testModelMessage = model.Message{
		Username: testUser,
		ChatId:   999,
		Text:     "test message",
	}
)

type testDeps struct {
	cRep      *mock_file.MockChatRepository
	uRep      *mock_file.MockUserRepository
	mRep      *mock_file.MockMessageRepository
	TxManager db.TxManager
}

func newTestDeps(t *testing.T) *testDeps {
	ctrl := gomock.NewController(t)

	return &testDeps{
		cRep:      mock_file.NewMockChatRepository(ctrl),
		uRep:      mock_file.NewMockUserRepository(ctrl),
		mRep:      mock_file.NewMockMessageRepository(ctrl),
		TxManager: db.NewTestTxManager(t),
	}
}

func (d *testDeps) newUseCase() *UseCase {
	return New(d.cRep, d.uRep, d.mRep, d.TxManager)
}
