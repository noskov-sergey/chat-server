package chats

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUseCase_CreateMessage_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.mRep.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil)

	uc := td.newUseCase()

	err := uc.CreateMessage(context.Background(), testModelMessage)

	require.NoError(t, err)
}

func TestUseCase_CreateMessage_Error(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.mRep.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(assert.AnError)

	uc := td.newUseCase()

	err := uc.CreateMessage(context.Background(), testModelMessage)

	require.Error(t, err)
}
