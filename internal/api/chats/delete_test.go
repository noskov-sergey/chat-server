package chats

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestImplementation_Delete_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.uc.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	_, err := td.client.Delete(context.Background(), testDeleteChatReq)

	require.NoError(t, err)
}

func TestImplementation_Delete_Error(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		td := newTestDeps(t)

		td.uc.EXPECT().
			Delete(gomock.Any(), gomock.Any()).
			Return(assert.AnError)

		_, err := td.client.Delete(context.Background(), testDeleteChatReq)

		require.Error(t, err)
	})
}
