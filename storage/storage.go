package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/novaflamestar/spade_goth/model"
)

type UserStore interface {
	CreateUser(string) error
	DeleteUser(string) error
	GetUser(string) (model.User, error)
	GetUsers() ([]model.User, error)
}
type Store struct {
	Users []model.User
}

func InitStore() *Store {
	store := new(Store)
	store.CreateUser("Karen")
	store.CreateUser("Bob")
	store.CreateUser("Ted")
	return store
}

func (s *Store) CreateUser(name string) error {
	user := model.User{}
	user.Name = name
	user.Id = uuid.NewString()
	s.Users = append(s.Users, user)
	return nil
}

func (s *Store) DeleteUser(id string) error {
	var filtered_list []model.User
	for _, user := range s.Users {
		if user.Id != id {
			filtered_list = append(filtered_list, user)
		}
	}
	if len(s.Users) == len(filtered_list) {
		return fmt.Errorf("no user to delete with id: %s", id)
	}
	return nil
}

func (s *Store) GetUser(id string) (model.User, error) {
	return model.User{}, nil
}

func (s *Store) GetUsers() ([]model.User, error) {
	return s.Users, nil
}
