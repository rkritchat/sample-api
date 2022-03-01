package repository

import (
	"errors"
	"fmt"
	"time"
)

type UserDetailEntity struct {
	Id        int64     `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
}

type UserDetail interface {
	Create(entity UserDetailEntity) error
	FindById(id int64) (*UserDetailEntity, error)
}

type userDetail struct {
	dbConnection string
}

func NewUserDetail(dbConnection string) UserDetail {
	return &userDetail{
		dbConnection: dbConnection,
	}
}

func (repo userDetail) Create(_ UserDetailEntity) error {
	fmt.Println("err while create entity")
	//implement here
	return errors.New("err while create entity")
}

func (repo userDetail) FindById(id int64) (*UserDetailEntity, error) {
	return &UserDetailEntity{
		Id:        id,
		Firstname: "kritchat",
		Lastname:  "rojaphtunk",
		CreatedAt: time.Now(),
	}, nil
}
