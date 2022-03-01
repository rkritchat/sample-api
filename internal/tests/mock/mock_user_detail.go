package mock

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"sample-api/internal/repository"
	"time"
)

type userDetail struct {
	mock.Mock
}

func (m *userDetail) Create(entity repository.UserDetailEntity) error {
	return m.Called(entity).Error(0)
}

func (m *userDetail) FindById(id int64) (*repository.UserDetailEntity, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*repository.UserDetailEntity), nil
	}
	return nil, args.Error(1)
}

func UserDetail(tc string) *userDetail {
	m := new(userDetail)
	switch tc {
	case "OK":
		m.On("Create", mock.AnythingOfType("repository.UserDetailEntity")).Return(nil)
		m.On("FindById", mock.AnythingOfType("int64")).Return(&repository.UserDetailEntity{
			Id:        1,
			Firstname: "aaa",
			Lastname:  "bbb",
			CreatedAt: time.Now(),
		}, nil)
	case "!OK":
		m.On("Create", mock.AnythingOfType("repository.UserDetailEntity")).Return(errors.New("mock err"))
		m.On("FindById", mock.AnythingOfType("int64")).Return(nil, errors.New("mock err"))
	}

	return m
}
