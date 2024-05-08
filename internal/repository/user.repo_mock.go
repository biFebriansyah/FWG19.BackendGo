package repository

import (
	"biFebriansyah/back/config"
	"biFebriansyah/back/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) CreateUser(data *models.User) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) GetAllUser() (*config.Result, error) {
	args := r.Mock.Called()
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) GetAuthData(user string) (*models.User, error) {
	args := r.Mock.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}
