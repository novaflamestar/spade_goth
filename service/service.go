package service

import "github.com/novaflamestar/spade_goth/storage"

type UserService struct {
	UserStore storage.UserStore
}

func InitService(store storage.UserStore) *UserService {
	return &UserService{UserStore: store}
}
