package chats

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUseCase_Delete_Success(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.cRep.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(nil)

	uc := td.newUseCase()

	err := uc.Delete(context.Background(), testId)

	require.NoError(t, err)
}

func TestUseCase_Delete_Error(t *testing.T) {
	t.Parallel()

	td := newTestDeps(t)

	td.cRep.EXPECT().
		Delete(gomock.Any(), gomock.Any()).
		Return(assert.AnError)

	uc := td.newUseCase()

	err := uc.Delete(context.Background(), testId)

	require.Error(t, err)
}
