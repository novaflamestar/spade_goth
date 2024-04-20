package service

import "github.com/novaflamestar/spade_goth/storage"

type UserService struct {
	UserStore *storage.Store
}

func InitService(store *storage.Store) *UserService {
	return &UserService{UserStore: store}
}
