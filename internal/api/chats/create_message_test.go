package chats

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestImplementation_CreateMessage_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.uc.EXPECT().
		CreateMessage(gomock.Any(), gomock.Any()).
		Return(nil)

	_, err := td.client.CreateMessage(context.Background(), testCreateMessageRequest)

	require.NoError(t, err)
}

func TestImplementation_CreateMessage_Error(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		td := newTestDeps(t)

		td.uc.EXPECT().
			CreateMessage(gomock.Any(), gomock.Any()).
			Return(assert.AnError)

		resp, err := td.client.CreateMessage(context.Background(), testCreateMessageRequest)

		require.Error(t, err)
		require.Nil(t, resp)
	})
}
