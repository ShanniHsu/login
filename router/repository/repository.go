package repository

import (
	"login/connect"
)

type Repo struct {
	UserRepository UserRepository
}

func NewRepository() Repo {
	db, _ := connect.GetDBConn()
	redis, _ := connect.Client()
	return Repo{
		UserRepository: NewUserRepository(db, redis),
	}
}
