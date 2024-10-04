package chats

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestImplementation_CreateChat_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.uc.EXPECT().
		CreateChat(gomock.Any(), gomock.Any()).
		Return(int(testChatId), nil)

	resp, err := td.client.CreateChat(context.Background(), testCreateChatReq)

	require.NoError(t, err)
	require.Equal(t, testChatId, resp.Id)
}

func TestImplementation_CreateChat_Error(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		td := newTestDeps(t)

		td.uc.EXPECT().
			CreateChat(gomock.Any(), gomock.Any()).
			Return(testChatIdError, assert.AnError)

		resp, err := td.client.CreateChat(context.Background(), testCreateChatReq)

		require.Error(t, err)
		require.Nil(t, resp)
	})
}
