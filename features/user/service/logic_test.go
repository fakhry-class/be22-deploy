package service

import (
	"be22/clean-arch/features/user"
	"be22/clean-arch/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetById(t *testing.T) {
	repoUserMock := new(mocks.UserData)
	hashMock := new(mocks.Hash)
	returnDataMock := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}

	t.Run("Success Get By Id", func(t *testing.T) {
		repoUserMock.On("SelectById", uint(1)).Return(&returnDataMock, nil).Once()

		srv := New(repoUserMock, hashMock)

		res, err := srv.GetById(uint(1))
		assert.NoError(t, err)
		assert.Equal(t, returnDataMock.ID, res.ID)
		repoUserMock.AssertExpectations(t)
	})
}
