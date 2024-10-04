package chats

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUseCase_CreateChat_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.cRep.EXPECT().
		Create(gomock.Any()).
		Return(testId, nil)

	td.uRep.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	uc := td.newUseCase()

	rid, err := uc.CreateChat(context.Background(), testModelUsers)

	require.NoError(t, err)
	require.Equal(t, testId, rid)
}

func TestUseCase_Create_Error(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.cRep.EXPECT().
		Create(gomock.Any()).
		Return(testIdError, assert.AnError)

	td.uRep.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	uc := td.newUseCase()

	id, err := uc.CreateChat(context.Background(), testModelUsers)

	require.Error(t, err)
	require.Equal(t, testIdError, id)
}
